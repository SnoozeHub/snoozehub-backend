package main

import (
	"context"
	_ "embed"
	"errors"
	"math/rand"
	"regexp"
	"time"

	"github.com/SnoozeHub/snoozehub-backend/grpc_gen"
	"github.com/badoux/checkmail"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/metadata"
)

const (
	internetConnectionFeature = 1
	bathroomFeature           = 2
	heatingFeature            = 3
	airConditionerFeature     = 4
	electricalOutletFeature   = 5
	tapFeature                = 6
	bedLinensFeature          = 7
	pillowsFeature            = 8
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
func featuresToInts(s []grpc_gen.Feature) []int32 {
	res := make([]int32, len(s))
	for i, f := range s {
		res[i] = int32(f)
	}
	return res
}
func intsTofeatures(s []int32) []grpc_gen.Feature {
	res := make([]grpc_gen.Feature, len(s))
	for i, f := range s {
		res[i] = grpc_gen.Feature(f)
	}
	return res
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
		context.TODO(),
		bson.D{{Key: "_id", Value: b.Host}},
	)
	var host account
	res.Decode(&host)

	return &grpc_gen.Bed{
		Id:                   &grpc_gen.BedId{BedId: b.Id},
		HostPublicKey:        host.PublicKey,
		HostTelegramUsername: host.TelegramUsername,
		BedMutableInfo: &grpc_gen.BedMutableInfo{
			Address:           b.Address,
			Coordinates:       &grpc_gen.Coordinates{Latitude: b.Latitude, Longitude: b.Longitude},
			Images:            b.Images,
			Description:       b.Description,
			Features:          intsTofeatures(b.Features),
			MinimumDaysNotice: uint32(b.MinimumDaysNotice),
		},
		DateAvailables:    dateAvailables,
		ReviewCount:       uint32(len(b.Reviews)),
		AverageEvaluation: averageEvaluation,
	}
}

func isAccountInfoValid(ai *grpc_gen.AccountInfo) bool {
	return len(ai.Name) >= 1 && len(ai.Name) <= 40 && checkmail.ValidateFormat(ai.Mail) == nil && len(ai.Mail) <= 60 &&
		regexp.MustCompile("^(?=.{5,32}$)(?!.*__)(?!^(telegram|admin|support))[a-z][a-z0-9_]*[a-z0-9]$").MatchString(ai.TelegramUsername)
}

// returns the publicKey only if is valid
func (s *authOnlyService) auth(ctx context.Context) (publicKey string, _ error) {

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", errors.New("can't get metadata")
	}

	authTokenArr := md["authToken"]
	if len(authTokenArr) != 1 {
		return "", errors.New("invalid authToken metadata")
	}

	authToken := authTokenArr[0]

	tmp, exist := s.authTokens.Get(authToken)
	if !exist {
		return "", errors.New("invalid or expired authToken")
	}

	publicKey, _ = tmp.(string)

	return publicKey, nil
}

// publicKey is valid
func (s *authOnlyService) accountExist(publicKey string) bool {
	res := s.db.Collection("accounts").FindOne(
		context.TODO(),
		bson.D{{Key: "publicKey", Value: publicKey}},
	)
	return res == nil
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
		context.TODO(),
		bson.D{{Key: "publicKey", Value: publicKey}, {Key: "verificationCode", Value: tmp}},
	)
	if res != nil {
		err = errors.New("account doesn't exist or not verified")
	}
	return
}

func (s *authOnlyService) isBookingValid(book *grpc_gen.Booking) bool {
	res := s.db.Collection("beds").FindOne(
		context.TODO(),
		bson.D{{Key: "id", Value: book.BedId.BedId}},
	)
	if res == nil {
		return false
	}
	var b bed
	res.Decode(&b)

	date := flatterizeDate(book.Date)
	contained := false
	for _, v := range b.DateAvailables {
		if v == date {
			contained = true
			break
		}
	}
	return contained
}

type booking struct {
	BedId string
	Date  int32
}

//go:embed assets/rest-abi.json
var restAbiJson string
