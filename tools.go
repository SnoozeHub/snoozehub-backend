package main

import (
	"context"
	_ "embed"
	"errors"
	"math/rand"
	"time"

	"github.com/SnoozeHub/snoozehub-backend/grpc_gen"
	"github.com/badoux/checkmail"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/metadata"
)

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func GenRandomString(length int) string {
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func isDateValidAndFromTomorrow(d *grpc_gen.Date) bool {
	if d.Year < 1000 && d.Year >= 3000 && d.Month < 1 && d.Month > 12 && d.Day < 1 && d.Day > 31 {
		return false
	}
	date := time.Date(int(d.Year), time.Month(d.Month), int(d.Day), 0, 0, 0, 0, time.Local)
	if date.Year() != int(d.Year) || int(date.Month()) != int(d.Month) || date.Day() != int(d.Day) {
		return false
	}
	return time.Now().Before(date)
}

// Assumed that isDateValidAndFromTomorrow(d) is true
func flatterizeDate(d *grpc_gen.Date) int32 {
	return int32(d.Year*10000 + d.Month*100 + d.Day)
}
func deflatterizeDate(s int32) *grpc_gen.Date {
	d := s % 100
	m := (s % 10000) / 100
	y := s / 10000
	return &grpc_gen.Date{
		Day:   uint32(d),
		Month: uint32(m),
		Year:  uint32(y),
	}
}

func allDistinct[T comparable](s []T) bool {
	m := make(map[T]struct{})
	for _, v := range s {
		_, found := m[v]
		if found {
			return false
		}
		m[v] = struct{}{}
	}
	return true
}

func bedToGrpcBed(db *mongo.Database, b bed) *grpc_gen.Bed {
	dateAvailables := make([]*grpc_gen.Date, len(b.DateAvailables))
	for _, v := range b.DateAvailables {
		dateAvailables = append(dateAvailables, deflatterizeDate(v))
	}
	var averageEvaluation *uint32 = nil
	if b.AverageEvaluation != nil {
		tmp2 := uint32(*b.AverageEvaluation)
		averageEvaluation = &tmp2
	}

	res := db.Collection("accounts").FindOne(
		context.Background(),
		bson.D{{Key: "_id", Value: b.Host}},
	)
	var host account
	res.Decode(&host)

	return &grpc_gen.Bed{
		Id:                   &grpc_gen.BedId{BedId: b.Id.Hex()},
		HostPublicKey:        host.PublicKey,
		HostTelegramUsername: host.TelegramUsername,
		BedMutableInfo: &grpc_gen.BedMutableInfo{
			Address:           b.Address,
			Coordinates:       &grpc_gen.Coordinates{Latitude: b.Latitude, Longitude: b.Longitude},
			Images:            b.Images,
			Description:       b.Description,
			Features:          b.Features,
			MinimumDaysNotice: uint32(b.MinimumDaysNotice),
		},
		DateAvailables:    dateAvailables,
		ReviewCount:       uint32(len(b.Reviews)),
		AverageEvaluation: averageEvaluation,
	}
}

func isAccountInfoValid(ai *grpc_gen.AccountInfo) bool {
	return len(ai.Name) >= 1 && len(ai.Name) <= 40 && checkmail.ValidateFormat(ai.Mail) == nil && len(ai.Mail) <= 60
}

// returns the publicKey only if is valid
func (s *authOnlyService) auth(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", errors.New("can't get metadata")
	}

	authTokenArr := md["authtoken"]
	if len(authTokenArr) != 1 {
		return "", errors.New("invalid authtoken metadata")
	}

	authToken := authTokenArr[0]

	tmp, exist := s.authTokens.Get(authToken)
	if !exist {
		return "", errors.New("invalid or expired authtoken")
	}

	publicKey, _ := tmp.(string)

	return publicKey, nil
}

// publicKey is valid
func (s *authOnlyService) accountExist(publicKey string) bool {
	return s.db.Collection("accounts").FindOne(
		context.Background(),
		bson.D{{Key: "publicKey", Value: publicKey}},
	).Err() == nil
}

func (s *authOnlyService) authAndExist(ctx context.Context) (publicKey string, err error) {
	publicKey, err = s.auth(ctx)
	if !s.accountExist(publicKey) {
		publicKey = ""
		err = errors.New("account doesn't exist")
	}
	return
}
func (s *authOnlyService) authAndExistAndVerified(ctx context.Context) (publicKey string, err error) {
	publicKey, err = s.auth(ctx)
	if err != nil {
		return
	}
	var tmp *string = nil
	res := s.db.Collection("accounts").FindOne(
		context.Background(),
		bson.D{{Key: "publicKey", Value: publicKey}, {Key: "verificationCode", Value: tmp}},
	)
	if res.Err() != nil {
		err = errors.New("account doesn't exist or not verified")
	}
	return
}

func (s *authOnlyService) isBookingValid(book *grpc_gen.Booking) bool {
	res := s.db.Collection("beds").FindOne(
		context.Background(),
		bson.D{{Key: "_id", Value: hexToObjectId(book.BedId.BedId)}},
	)
	if res.Err() != nil {
		return false
	}
	var b bed
	res.Decode(&b)

	days := numDaysUntil(book.Date)
	if days < 1 || days > int(b.MinimumDaysNotice) {
		return false
	}

	date := flatterizeDate(book.Date)
	for _, v := range b.DateAvailables {
		if v == date {
			return true
		}
	}
	return false
}

func (s *authOnlyService) doesBedIdExist(bedId *grpc_gen.BedId) bool {
	res := s.db.Collection("beds").FindOne(
		context.Background(),
		bson.D{{Key: "_id", Value: hexToObjectId(bedId.BedId)}},
	)
	return res.Err() == nil
}
func (s *authOnlyService) getBed(bedId string) bed {
	res := s.db.Collection("beds").FindOne(
		context.Background(),
		bson.D{{Key: "_id", Value: hexToObjectId(bedId)}},
	)
	var b bed
	res.Decode(&b)
	return b
}

func (s *authOnlyService) publicKeyHasReviewed(pubKey string, bedId string) bool {
	b := s.getBed(bedId)
	for _, v := range b.Reviews {
		res := s.db.Collection("accounts").FindOne(
			context.Background(),
			bson.D{{Key: "_id", Value: v.Reviewer}},
		)
		var a account
		res.Decode(&a)
		if a.PublicKey == pubKey {
			return true
		}
	}
	return false
}

type booking struct {
	BedId string
	Date  int32
}

//go:embed assets/rest-abi.json
var restAbiJson string

func isImageValid(image []byte) bool {
	return len(image) <= 512*1024
}

func numDaysUntil(date *grpc_gen.Date) int {
	toDate := func(d *grpc_gen.Date) time.Time {
		return time.Date(int(d.Year), time.Month(d.Month), int(d.Day), 0, 0, 0, 0, time.Local)
	}

	tmp := time.Now()
	now := toDate(&grpc_gen.Date{Day: uint32(tmp.Day()), Month: uint32(tmp.Month()), Year: uint32(tmp.Year())})
	return int(toDate(date).Sub(now).Hours()) / 24
}

func (s *authOnlyService) adjustAverageEvaluation(bedId string) {
	b := s.getBed(bedId)
	sum := int32(0)
	for _, r := range b.Reviews {
		sum += int32(r.Evaluation)
	}
	var eval *int32 = nil
	if len(b.Reviews) > 0 {
		*eval = sum / int32(len(b.Reviews))
	}

	filter := bson.D{{Key: "_id", Value: hexToObjectId(bedId)}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "averageEvaluation", Value: eval}}}}
	s.db.Collection("beds").UpdateOne(context.Background(), filter, update)
}

func hexToObjectId(hex string) primitive.ObjectID {
	tmp, _ := primitive.ObjectIDFromHex(hex)
	return tmp
}