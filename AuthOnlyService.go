package main

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"strings"
	"sync"
	"time"

	"github.com/SnoozeHub/snoozehub-backend/dev_vs_prod"
	"github.com/SnoozeHub/snoozehub-backend/grpc_gen"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/patrickmn/go-cache"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type authOnlyService struct {
	grpc_gen.UnimplementedAuthOnlyServiceServer
	authTokens        *cache.Cache
	db                *mongo.Database
	attendingBookings map[booking]bool
	mu                *sync.Mutex
}

func newAuthOnlyService(db *mongo.Database) *authOnlyService {
	dev_vs_prod.Init()
	service := authOnlyService{
		authTokens: cache.New(24*time.Hour, 24*time.Hour),
		db:         db,
		mu:         &sync.Mutex{},
	}
	return &service
}
func (s *authOnlyService) GetAuthTokens() *cache.Cache {
	return s.authTokens
}
func (s *authOnlyService) GetMutex() *sync.Mutex {
	return s.mu
}

func (s *authOnlyService) SignUp(ctx context.Context, req *grpc_gen.AccountInfo) (*grpc_gen.Empty, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	publicKey, err := s.auth(ctx)
	if err != nil {
		return nil, err
	}

	if s.accountExist(publicKey) {
		return nil, errors.New("account already exist")
	}

	if !isAccountInfoValid(req) {
		return nil, errors.New("invalid account info")
	}

	verificationCode := GenRandomString(5)

	account := account{
		Id:               primitive.NewObjectID(),
		PublicKey:        publicKey,
		Name:             req.Name,
		Mail:             req.Mail,
		TelegramUsername: req.TelegramUsername,
		ProfilePic:       nil,
		VerificationCode: &verificationCode,
		BedIdBookings:    []primitive.ObjectID{},
	}
	accountMarsheled, _ := bson.Marshal(account)

	s.db.Collection("accounts").InsertOne(
		context.Background(),
		accountMarsheled,
	)

	dev_vs_prod.Send(req.Mail, "Verify your mail", "Verification code: "+verificationCode)

	return &grpc_gen.Empty{}, nil
}
func (s *authOnlyService) VerifyMail(ctx context.Context, req *grpc_gen.VerifyMailRequest) (*grpc_gen.VerifyMailResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	publicKey, err := s.authAndExist(ctx)
	if err != nil {
		return nil, err
	}

	if len(req.VerificationCode) != 5 {
		return nil, errors.New("verification code wrong lenght")
	}

	filter := bson.D{{Key: "publicKey", Value: publicKey}, {Key: "verificationCode", Value: req.VerificationCode}}
	var tmp *string = nil
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "verificationCode", Value: tmp}}}}

	res, err := s.db.Collection("accounts").UpdateOne(
		context.Background(),
		filter,
		update,
	)
	if err != nil {
		return nil, err
	} else if res.ModifiedCount == 0 {
		return &grpc_gen.VerifyMailResponse{Ok: false}, nil
	} else {
		return &grpc_gen.VerifyMailResponse{Ok: true}, nil
	}
}
func (s *authOnlyService) GetAccountInfo(ctx context.Context, req *grpc_gen.Empty) (*grpc_gen.AccountInfo, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	publicKey, err := s.authAndExist(ctx)
	if err != nil {
		return nil, err
	}

	res := s.db.Collection("accounts").FindOne(
		context.Background(),
		bson.D{{Key: "publicKey", Value: publicKey}},
	)
	if res.Err() != nil {
		return nil, errors.New("error getting the account")
	}
	var acc account
	res.Decode(&acc)

	accInfo := grpc_gen.AccountInfo{
		Name:             acc.Name,
		Mail:             acc.Mail,
		TelegramUsername: acc.TelegramUsername,
	}

	return &accInfo, nil
}
func (s *authOnlyService) GetProfilePic(ctx context.Context, req *grpc_gen.Empty) (*grpc_gen.ProfilePic, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	publicKey, err := s.authAndExist(ctx)
	if err != nil {
		return nil, err
	}

	res := s.db.Collection("accounts").FindOne(
		context.Background(),
		bson.D{{Key: "publicKey", Value: publicKey}},
	)

	if res.Err() != nil {
		return nil, errors.New("error getting the account")
	}
	var acc account
	res.Decode(&acc)

	profPic := grpc_gen.ProfilePic{
		Image: acc.ProfilePic,
	}
	return &profPic, nil
}
func (s *authOnlyService) SetProfilePic(ctx context.Context, req *grpc_gen.ProfilePic) (*grpc_gen.Empty, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	publicKey, err := s.authAndExist(ctx)
	if err != nil {
		return nil, err
	}

	if !isImageValid(req.Image) {
		return nil, errors.New("too large image")
	}

	filter := bson.D{{Key: "publicKey", Value: publicKey}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "profilePic", Value: req.Image}}}}

	_, err = s.db.Collection("accounts").UpdateOne(
		context.Background(),
		filter,
		update,
	)
	if err != nil {
		return nil, err
	}
	return &grpc_gen.Empty{}, nil
}
func (s *authOnlyService) DeleteAccount(ctx context.Context, req *grpc_gen.Empty) (*grpc_gen.Empty, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	publicKey, err := s.authAndExist(ctx)
	if err != nil {
		return nil, err
	}

	filter := bson.D{{Key: "publicKey", Value: publicKey}}
	res := s.db.Collection("accounts").FindOne(context.Background(), filter)

	var host account
	res.Decode(&host)

	// Remove my reviews
	for _, id := range host.BedIdBookings {
		update := bson.M{"$pull": bson.M{"reviews": bson.M{"reviewer": host.Id}}}
		s.db.Collection("beds").UpdateOne(context.Background(), bson.D{}, update)

		s.adjustAverageEvaluation(id.Hex())
	}

	// remove my beds
	cur, _ := s.db.Collection("beds").Find(context.Background(), bson.D{{Key: "host", Value: host.Id}})

	var myBeds []bed
	cur.All(context.Background(), &myBeds)

	for _, b := range myBeds {
		// Remove bedIdBookings of other accounts that points to b
		s.db.Collection("accounts").UpdateMany(context.Background(), bson.M{}, bson.M{"$pull": bson.M{"bedIdBookings": bson.M{"$eq": b.Id}}})

		// actual remove b
		s.db.Collection("beds").DeleteOne(context.Background(), bson.D{{Key: "_id", Value: b.Id}})
	}

	// actually remove account
	_, err = s.db.Collection("accounts").DeleteOne(
		context.Background(),
		filter,
	)
	if err != nil {
		return nil, err
	}

	return &grpc_gen.Empty{}, nil
}
func (s *authOnlyService) UpdateAccountInfo(ctx context.Context, req *grpc_gen.AccountInfo) (*grpc_gen.Empty, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	publicKey, err := s.authAndExist(ctx)
	if err != nil {
		return nil, err
	}

	if !isAccountInfoValid(req) {
		return nil, errors.New("invalid account info")
	}

	res := s.db.Collection("accounts").FindOne(
		context.Background(),
		bson.D{{Key: "publicKey", Value: publicKey}},
	)

	var acc account
	res.Decode(&acc)

	var verificationCode *string = nil

	if acc.Mail != req.Mail {
		tmp := GenRandomString(5)
		verificationCode = &tmp

		dev_vs_prod.Send(acc.Mail, "Verify your mail", "Verification code: "+*verificationCode)
	}

	s.db.Collection("accounts").UpdateOne(
		context.Background(),
		bson.D{{Key: "publicKey", Value: publicKey}},
		bson.D{{Key: "$set", Value: bson.D{{Key: "verificationCode", Value: verificationCode},
			{Key: "name", Value: req.Name},
			{Key: "mail", Value: req.Mail},
			{Key: "telegramUsername", Value: req.TelegramUsername}}}},
	)

	return &grpc_gen.Empty{}, nil
}
func (s *authOnlyService) Logout(ctx context.Context, empt *grpc_gen.Empty) (*grpc_gen.Empty, error) {
	return nil, nil
}

func (s *authOnlyService) Book(ctx context.Context, req *grpc_gen.Booking) (*grpc_gen.BookResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	publicKey, err := s.authAndExistAndVerified(ctx)
	if err != nil {
		return nil, err
	}

	if !s.isBookingValid(req) {
		return nil, errors.New("invalid booking")
	}

	dates := dateIntervalToDateSlice(req.DateInterval)
	bookings := make([]booking, len(dates))
	for i, v := range dates {
		bookings[i] = booking{
			BedId: req.BedId.BedId,
			Date:  flatterizeDate(timeToGrpcDate(&v)),
		}
	}

	// check if another bookings is running
	for _, book := range bookings {
		_, exist := s.attendingBookings[book]
		if exist {
			return &grpc_gen.BookResponse{IsBookingUnlocked: false}, nil
		}
	}

	// Set this bookings as running
	for _, book := range bookings {
		s.attendingBookings[book] = true
	}

	// get host public key
	res := s.db.Collection("beds").FindOne(
		context.Background(),
		bson.D{{Key: "_id", Value: hexToObjectId(bookings[0].BedId)}},
	)
	var b bed
	res.Decode(&b)
	res = s.db.Collection("accounts").FindOne(
		context.Background(),
		bson.D{{Key: "_id", Value: b.Host}},
	)

	var host account
	res.Decode(&host)

	// Check if transfer is done
	client, err := ethclient.Dial("https://goerli.infura.io/v3/9aa3d95b3bc440fa88ea12eaa4456161")
	if err != nil {
		return nil, errors.New("error connecting to ethereum rpc")
	}
	query := ethereum.FilterQuery{
		Addresses: []common.Address{common.HexToAddress("0xc2Cd631A73E0D94dE392F686D22E9e792E426000")},
	}

	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		return nil, errors.New("error subscribing to contract")
	}

	go func() {

		timeout := time.After(60 * time.Second)

		contractAbi, _ := abi.JSON(strings.NewReader(restAbiJson))

		for {
			select {
			case <-sub.Err():
				return
			case <-timeout:
				return
			case log := <-logs:
				var amount struct{ Value *big.Int }
				contractAbi.UnpackIntoInterface(&amount, "Transfer", log.Data)
				from := common.BytesToAddress(log.Topics[1].Bytes()).String()
				to := common.BytesToAddress(log.Topics[2].Bytes()).String()

				// actual check
				if from == publicKey && to == host.PublicKey && amount.Value.Cmp(big.NewInt(1)) == len(bookings) {
					s.mu.Lock()
					defer s.mu.Unlock()

					// check if the guest account still exist and the booking is still valid, otherwise return
					publicKey, err := s.authAndExistAndVerified(ctx)
					if err != nil {
						return
					}
					if !s.isBookingValid(req) {
						return
					}

					humanProofToken := GenRandomString(4)

					// get latest host informations
					res = s.db.Collection("accounts").FindOne(
						context.Background(),
						bson.D{{Key: "_id", Value: b.Host}},
					)
					res.Decode(&host)

					// send mails
					bookingInfo :=
						`Booking info:
						Bed id: ` + req.BedId.BedId + `
						Start date: ` + fmt.Sprint(req.DateInterval.StartDate.Day, '/', req.DateInterval.StartDate.Month, '/', req.DateInterval.StartDate.Year) + `
						End date: ` + fmt.Sprint(req.DateInterval.EndDate.Day, '/', req.DateInterval.EndDate.Month, '/', req.DateInterval.EndDate.Year)
					dev_vs_prod.Send(host.Mail, "You have a new guest!", bookingInfo+"\nIn order to authenticate him, use the following snooze token: "+humanProofToken)

					res = s.db.Collection("accounts").FindOne(
						context.Background(),
						bson.D{{Key: "publicKey", Value: publicKey}},
					)

					var guest account
					res.Decode(&guest)

					dev_vs_prod.Send(guest.Mail, "You have booked a bed!", bookingInfo+"\nIn order to authenticate you with the host, use the following snooze token: "+humanProofToken)

					// update db
					filter := bson.M{"PublicKey": host.PublicKey}
					update := bson.M{"$addToSet": bson.M{"bedIdBookings": req.BedId.BedId}}
					s.db.Collection("accounts").UpdateOne(context.Background(), filter, update)

					filter = bson.M{"_id": hexToObjectId(req.BedId.BedId)}
					tmp := make([]int32, 0)
					for _, b2 := range bookings {
						tmp = append(tmp, b2.Date)
					}
					update = bson.M{"$pull": bson.M{"dateAvailables": bson.M{"$in": tmp}}}
					s.db.Collection("beds").UpdateOne(context.Background(), filter, update)

					return
				}
			}
		}
	}()

	return &grpc_gen.BookResponse{IsBookingUnlocked: true}, nil
}
func (s *authOnlyService) Review(ctx context.Context, req *grpc_gen.ReviewRequest) (*grpc_gen.Empty, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	publicKey, err := s.authAndExistAndVerified(ctx)
	if err != nil {
		return nil, err
	}

	publicKeyHasBooked := func(pubKey string, bedId string) bool {
		res := s.db.Collection("accounts").FindOne(
			context.Background(),
			bson.D{{Key: "publicKey", Value: pubKey}},
		)
		var a account
		res.Decode(&a)
		for _, v := range a.BedIdBookings {
			if v.Hex() == bedId {
				return true
			}
		}
		return false
	}

	if !s.doesBedIdExist(req.BedId) || !publicKeyHasBooked(publicKey, req.BedId.BedId) || s.publicKeyHasReviewed(publicKey, req.BedId.BedId) || len(req.Review.Comment) > 200 || req.Review.Evaluation > 50 {
		return nil, errors.New("invalid request")
	}

	res := s.db.Collection("accounts").FindOne(
		context.Background(),
		bson.D{{Key: "publicKey", Value: publicKey}},
	)
	var a account
	res.Decode(&a)

	r := review{
		Reviewer:   a.Id,
		Evaluation: int32(req.Review.Evaluation),
		Comment:    req.Review.Comment,
	}
	rMarsheled, _ := bson.Marshal(r)
	filter := bson.M{"_id": hexToObjectId(req.BedId.BedId)}
	update := bson.M{"$addToSet": bson.M{"reviews": rMarsheled}}
	s.db.Collection("beds").UpdateOne(context.Background(), filter, update)

	s.adjustAverageEvaluation(req.BedId.BedId)

	return &grpc_gen.Empty{}, nil
}
func (s *authOnlyService) GetMyReview(ctx context.Context, req *grpc_gen.BedId) (*grpc_gen.GetMyReviewResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	publicKey, err := s.authAndExistAndVerified(ctx)
	if err != nil {
		return nil, err
	}

	if !s.doesBedIdExist(req) {
		return nil, errors.New("invalid request")
	}

	res := s.db.Collection("accounts").FindOne(
		context.Background(),
		bson.D{{Key: "publicKey", Value: publicKey}},
	)
	var a account
	res.Decode(&a)

	var r *grpc_gen.Review = nil
	b := s.getBed(req.BedId)
	for _, v := range b.Reviews {
		if v.Reviewer == a.Id {
			r = &grpc_gen.Review{Evaluation: uint32(v.Evaluation), Comment: v.Comment}
			break
		}
	}

	return &grpc_gen.GetMyReviewResponse{Review: r}, nil
}
func (s *authOnlyService) RemoveReview(ctx context.Context, req *grpc_gen.BedId) (*grpc_gen.Empty, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	publicKey, err := s.authAndExistAndVerified(ctx)
	if err != nil {
		return nil, err
	}

	if !s.doesBedIdExist(req) || !s.publicKeyHasReviewed(publicKey, req.BedId) {
		return nil, errors.New("invalid request")
	}

	res := s.db.Collection("accounts").FindOne(
		context.Background(),
		bson.D{{Key: "publicKey", Value: publicKey}},
	)
	var a account
	res.Decode(&a)

	filter := bson.M{"_id": hexToObjectId(req.BedId)}
	update := bson.M{"$pull": bson.M{"reviews": bson.M{"reviewer": a.Id}}}
	s.db.Collection("beds").UpdateOne(context.Background(), filter, update)

	s.adjustAverageEvaluation(req.BedId)

	return &grpc_gen.Empty{}, nil
}
func (s *authOnlyService) AddBed(ctx context.Context, req *grpc_gen.BedMutableInfo) (*grpc_gen.Empty, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	publicKey, err := s.authAndExistAndVerified(ctx)
	if err != nil {
		return nil, err
	}

	validImages := func(images [][]byte) bool {
		for _, v := range images {
			if !isImageValid(v) {
				return false
			}
		}
		return true
	}
	if len(req.Address) < 1 || len(req.Address) > 100 || req.Coordinates.Latitude < -90 || req.Coordinates.Latitude > 90 || req.Coordinates.Longitude < -180 || req.Coordinates.Longitude > 180 ||
		len(req.Description) > 200 || !allDistinct(req.Features) || len(req.Images) < 1 || len(req.Images) > 5 || !validImages(req.Images) || req.MinimumDaysNotice < 1 || req.MinimumDaysNotice > 30 {
		return nil, errors.New("invalid request")
	}

	// insert

	res := s.db.Collection("accounts").FindOne(
		context.Background(),
		bson.D{{Key: "publicKey", Value: publicKey}},
	)
	if res.Err() != nil {
		return nil, res.Err()
	}
	var a account
	res.Decode(&a)

	b, _ := bson.Marshal(bed{
		Id:                primitive.NewObjectID(),
		Host:              a.Id,
		Address:           req.Address,
		Latitude:          req.Coordinates.Latitude,
		Longitude:         req.Coordinates.Longitude,
		Images:            req.Images,
		Description:       req.Description,
		Features:          req.Features,
		MinimumDaysNotice: int32(req.MinimumDaysNotice),
		DateAvailables:    []int32{},
		AverageEvaluation: nil,
		Reviews:           []review{},
	})
	s.db.Collection("beds").InsertOne(context.Background(), b)

	return &grpc_gen.Empty{}, nil
}
func (s *authOnlyService) ModifyMyBed(ctx context.Context, req *grpc_gen.ModifyBedRequest) (*grpc_gen.Empty, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	publicKey, err := s.authAndExistAndVerified(ctx)
	if err != nil {
		return nil, err
	}

	validImages := func(images [][]byte) bool {
		for _, v := range images {
			if !isImageValid(v) {
				return false
			}
		}
		return true
	}

	publicKeyOwnBed := func() bool {
		res := s.db.Collection("accounts").FindOne(context.Background(), bson.D{{Key: "publicKey", Value: publicKey}})
		var a account
		res.Decode(&a)
		return a.Id == s.getBed(req.BedId.BedId).Host
	}

	if !s.doesBedIdExist(req.BedId) || !publicKeyOwnBed() || len(req.BedMutableInfo.Address) < 1 || len(req.BedMutableInfo.Address) > 100 || req.BedMutableInfo.Coordinates.Latitude < -90 || req.BedMutableInfo.Coordinates.Latitude > 90 || req.BedMutableInfo.Coordinates.Longitude < -180 || req.BedMutableInfo.Coordinates.Longitude > 180 ||
		len(req.BedMutableInfo.Description) > 200 || !allDistinct(req.BedMutableInfo.Features) || len(req.BedMutableInfo.Images) < 1 || len(req.BedMutableInfo.Images) > 5 || !validImages(req.BedMutableInfo.Images) || req.BedMutableInfo.MinimumDaysNotice < 1 || req.BedMutableInfo.MinimumDaysNotice > 30 {
		return nil, errors.New("invalid request")
	}

	filter := bson.D{{Key: "_id", Value: hexToObjectId(req.BedId.BedId)}}
	update := bson.D{{Key: "$set", Value: bson.D{
		{Key: "address", Value: req.BedMutableInfo.Address},
		{Key: "latitude", Value: req.BedMutableInfo.Coordinates.Latitude},
		{Key: "longitude", Value: req.BedMutableInfo.Coordinates.Longitude},
		{Key: "images", Value: req.BedMutableInfo.Images},
		{Key: "description", Value: req.BedMutableInfo.Description},
		{Key: "features", Value: req.BedMutableInfo.Features},
		{Key: "minimumDaysNotice", Value: req.BedMutableInfo.MinimumDaysNotice}}}}
	s.db.Collection("beds").UpdateOne(context.Background(), filter, update)

	return &grpc_gen.Empty{}, nil
}
func (s *authOnlyService) RemoveMyBed(ctx context.Context, req *grpc_gen.BedId) (*grpc_gen.Empty, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	publicKey, err := s.authAndExistAndVerified(ctx)
	if err != nil {
		return nil, err
	}

	publicKeyOwnBed := func() bool {
		res := s.db.Collection("accounts").FindOne(context.Background(), bson.D{{Key: "publicKey", Value: publicKey}})
		var a account
		res.Decode(&a)
		return a.Id == s.getBed(req.BedId).Host
	}

	if !s.doesBedIdExist(req) || !publicKeyOwnBed() {
		return nil, errors.New("invalid request")
	}

	// Remove related BedIdBookings
	update := bson.M{"$pull": bson.M{"bedIdBookings": bson.M{"$eq": req.BedId}}}
	s.db.Collection("accounts").UpdateMany(context.Background(), bson.M{}, update)

	// Actual remove
	filter := bson.D{{Key: "_id", Value: hexToObjectId(req.BedId)}}
	s.db.Collection("beds").DeleteOne(context.Background(), filter)

	return &grpc_gen.Empty{}, nil
}
func (s *authOnlyService) GetMyBeds(ctx context.Context, _ *grpc_gen.Empty) (*grpc_gen.BedList, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	publicKey, err := s.authAndExistAndVerified(ctx)
	if err != nil {
		return nil, err
	}

	res := s.db.Collection("accounts").FindOne(context.Background(), bson.D{{Key: "publicKey", Value: publicKey}})
	var a account
	res.Decode(&a)

	cursor, _ := s.db.Collection("beds").Find(context.Background(), bson.D{{Key: "host", Value: a.Id}})

	var beds []bed
	cursor.All(context.Background(), &beds)
	var grpcBeds []*grpc_gen.Bed
	for _, b := range beds {
		grpcBeds = append(grpcBeds, bedToGrpcBed(s.db, b))
	}

	return &grpc_gen.BedList{Beds: grpcBeds}, nil
}
func (s *authOnlyService) AddBookingAvailability(ctx context.Context, req *grpc_gen.BookingAvailability) (*grpc_gen.Empty, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	publicKey, err := s.authAndExistAndVerified(ctx)
	if err != nil {
		return nil, err
	}

	// get bed
	res := s.db.Collection("beds").FindOne(
		context.Background(),
		bson.D{{Key: "_id", Value: hexToObjectId(req.BedId.BedId)}},
	)
	if res.Err() != nil {
		return nil, errors.New("invalid bedId")
	}
	var b bed
	res.Decode(&b)

	// Check if the caller own the bed
	res = s.db.Collection("accounts").FindOne(context.Background(), bson.D{{Key: "publicKey", Value: publicKey}})
	var a account
	res.Decode(&a)

	if b.Host != a.Id {
		return nil, errors.New("caller doesn't own the bed with the specified bedId")
	}

	// Check if DateInterval is valid
	if !isDateIntervalValid(req.DateInterval) {
		return nil, errors.New("invalid date interval")
	}

	// Check if dates are not available
	datesSlice := dateIntervalToDateSlice(req.DateInterval)
	for _, date := range datesSlice {
		for _, ava := range b.DateAvailables {
			avaAsDate := grpcDateToTime(deflatterizeDate(ava))
			if datesAreSameDay(&date, avaAsDate) {
				return nil, errors.New("invalid date interval")
			}
		}
	}
	dates := dateSliceToFlatSlice(datesSlice)
	// Check dates vailidty
	days := numDaysUntil(req.DateInterval.StartDate)
	if days < 1 {
		return nil, errors.New("invalid date interval")
	}
	days = numDaysUntil(req.DateInterval.EndDate)
	if days > 90 {
		return nil, errors.New("invalid date interval")
	}

	// Add availability
	filter := bson.M{"_id": hexToObjectId(req.BedId.BedId)}
	update := bson.M{"$addToSet": bson.M{"dateAvailables": bson.M{
		"$each": dates,
	}}}
	s.db.Collection("beds").UpdateOne(context.Background(), filter, update)
	return &grpc_gen.Empty{}, nil
}
func (s *authOnlyService) RemoveBookAvailability(ctx context.Context, req *grpc_gen.BookingAvailability) (*grpc_gen.Empty, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	publicKey, err := s.authAndExistAndVerified(ctx)
	if err != nil {
		return nil, err
	}

	// get bed
	res := s.db.Collection("beds").FindOne(
		context.Background(),
		bson.D{{Key: "_id", Value: hexToObjectId(req.BedId.BedId)}},
	)
	if res.Err() != nil {
		return nil, errors.New("invalid bedId")
	}
	var b bed
	res.Decode(&b)

	// Check if the caller own the bed
	res = s.db.Collection("accounts").FindOne(context.Background(), bson.D{{Key: "publicKey", Value: publicKey}})
	var a account
	res.Decode(&a)

	if b.Host != a.Id {
		return nil, errors.New("caller doesn't own the bed with the specified bedId")
	}

	// Check if DateInterval is valid
	if !isDateIntervalValid(req.DateInterval) {
		return nil, errors.New("invalid date interval")
	}

	// Check if date is not available
	datesSlice := dateIntervalToDateSlice(req.DateInterval)
	for _, date := range datesSlice {
		for _, ava := range b.DateAvailables {
			avaAsDate := grpcDateToTime(deflatterizeDate(ava))
			if datesAreSameDay(&date, avaAsDate) {
				goto next
			}
		}
		return nil, errors.New("invalid date")
	next:
	}
	dates := dateSliceToFlatSlice(datesSlice)

	// Remove availability
	filter := bson.M{"_id": hexToObjectId(req.BedId.BedId)}
	update := bson.M{"$pull": bson.M{"dateAvailables": bson.M{"$in": dates}}}
	s.db.Collection("beds").UpdateOne(context.Background(), filter, update)

	return &grpc_gen.Empty{}, nil
}
