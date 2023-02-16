package main

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"strings"
	"sync"
	"time"

	"github.com/SnoozeHub/snoozehub-backend/grpc_gen"
	"github.com/SnoozeHub/snoozehub-backend/mail"
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
	mail.Init()
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

	mail.Send(req.Mail, "Verify your mail", "Verification code: "+verificationCode)

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
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "profilePicture", Value: req.Image}}}}

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

		mail.Send(acc.Mail, "Verify your mail", "Verification code: "+*verificationCode)
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

	book := booking{
		BedId: req.BedId.BedId,
		Date:  flatterizeDate(req.Date),
	}

	// check if another booking is running
	_, exist := s.attendingBookings[book]
	if exist {
		return &grpc_gen.BookResponse{IsBookingUnlocked: false}, nil
	}

	s.attendingBookings[book] = true

	// get host public key
	res := s.db.Collection("beds").FindOne(
		context.Background(),
		bson.D{{Key: "_id", Value: book.BedId}},
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
				if from == publicKey && to == host.PublicKey && amount.Value.Cmp(big.NewInt(1)) >= 0 {
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
					Bed id: ` + req.BedId.BedId +
							`Date: ` + fmt.Sprint(req.Date.Day, '/', req.Date.Month, '/', req.Date.Year)
					mail.Send(host.Mail, "You have a new guest!", bookingInfo+"\nIn order to authenticate him, use the following snooze token: "+humanProofToken)

					res = s.db.Collection("accounts").FindOne(
						context.Background(),
						bson.D{{Key: "publicKey", Value: publicKey}},
					)

					var guest account
					res.Decode(&guest)

					mail.Send(guest.Mail, "You have booked a bed!", bookingInfo+"\nIn order to authenticate you with the host, use the following snooze token: "+humanProofToken)

					// update db
					filter := bson.M{"PublicKey": host.PublicKey}
					update := bson.M{"$addToSet": bson.M{"BedIdsBookings": req.BedId}}
					s.db.Collection("accounts").UpdateOne(context.Background(), filter, update)

					filter = bson.M{"_id": req.BedId}
					update = bson.M{"$pull": bson.M{"DateAvailables": bson.M{"$eq": book.Date}}}
					s.db.Collection("beds").UpdateOne(context.Background(), filter, update)

					return
				}
			}
		}
	}()

	return &grpc_gen.BookResponse{IsBookingUnlocked: true}, nil
}
func (s *authOnlyService) Review(context.Context, *grpc_gen.ReviewRequest) (*grpc_gen.Empty, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	return nil, nil
}
func (s *authOnlyService) RemoveReview(context.Context, *grpc_gen.BedId) (*grpc_gen.Empty, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	return nil, nil
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
		len(req.Description) > 200 || !allDistinct(req.Features) || len(req.Images) < 1 || len(req.Images) > 5 || validImages(req.Images) || req.MinimumDaysNotice < 1 || req.MinimumDaysNotice > 30 {
		return nil, errors.New("invalid request")
	}

	publicKey = publicKey // #todo

	return nil, nil
}
func (s *authOnlyService) ModifyMyBed(context.Context, *grpc_gen.ModifyBedRequest) (*grpc_gen.Empty, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	return nil, nil
}
func (s *authOnlyService) RemoveMyBed(context.Context, *grpc_gen.BedId) (*grpc_gen.Empty, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	return nil, nil
}
func (s *authOnlyService) GetMyBeds(context.Context, *grpc_gen.Empty) (*grpc_gen.BedList, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	return nil, nil
}
func (s *authOnlyService) AddBookingAvaiability(context.Context, *grpc_gen.Booking) (*grpc_gen.Empty, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	return nil, nil
}
func (s *authOnlyService) RemoveBookAvaiability(context.Context, *grpc_gen.Booking) (*grpc_gen.Empty, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	return nil, nil
}
