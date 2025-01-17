package main

import (
	"context"
	"strconv"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/SnoozeHub/snoozehub-backend/dev_vs_prod"
	"github.com/SnoozeHub/snoozehub-backend/grpc_gen"
	asserter "github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc/metadata"
)

func TestDbConection(t *testing.T) {
	_, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://root:root@mongodb:27017"))
	if err != nil {
		t.Fatal(err)
	}
}
func TestRESTTokenSubscription(t *testing.T) {
	assert := asserter.New(t)
	client, err := ethclient.Dial("wss://goerli.infura.io/ws/v3/9aa3d95b3bc440fa88ea12eaa4456161")
	assert.Nil(err)
	query := ethereum.FilterQuery{
		Addresses: []common.Address{common.HexToAddress("0xc2Cd631A73E0D94dE392F686D22E9e792E426000")},
	}

	logs := make(chan types.Log)
	_, err = client.SubscribeFilterLogs(context.Background(), query, logs)

	assert.Nil(err)
}

func TestRpcs(t *testing.T) {
	// Restore default state of db
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://root:root@mongodb:27017"))
	if err != nil {
		t.Fatal(err)
	}
	err = client.Database("main").Drop(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	client.Disconnect(context.Background())

	assert := asserter.New(t)
	go main()
	time.Sleep(1 * time.Second) // Give time to start the server
	conn, err := grpc.Dial("localhost:9090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	assert.Nil(err)

	publicService := grpc_gen.NewPublicServiceClient(conn)
	authOnlyService := grpc_gen.NewAuthOnlyServiceClient(conn)

	var token string

	t.Run("TestAuthentication", func(t *testing.T) {
		assert := asserter.New(t)

		getNonceResponse, err := publicService.GetNonce(context.Background(), &grpc_gen.Empty{})
		if err != nil {
			t.Log(err.Error())
			t.Log(err)
		}
		assert.Nil(err)

		privateKey, _ := crypto.GenerateKey()

		t.Log("public key: " + crypto.PubkeyToAddress(privateKey.PublicKey).String())

		hash := crypto.Keccak256([]byte("\x19Ethereum Signed Message:\n" + strconv.Itoa(len(getNonceResponse.Nonce)) + getNonceResponse.Nonce))
		signature, _ := crypto.Sign(hash, privateKey)

		authResponse, err := publicService.Auth(context.Background(), &grpc_gen.AuthRequest{Nonce: getNonceResponse.Nonce, SignedNonce: signature})

		assert.Nil(err)
		assert.False(authResponse.AccountExist)
		token = authResponse.AuthToken
	})
	ctx := metadata.NewOutgoingContext(
		context.Background(),
		metadata.Pairs("authtoken", token),
	)

	t.Run("TestSignUpAndVerify", func(t *testing.T) {
		assert := asserter.New(t)

		_, err := authOnlyService.SignUp(ctx, &grpc_gen.AccountInfo{
			Name:             "user",
			Mail:             "user2@example.com",
			TelegramUsername: "username",
		},
		)
		assert.Nil(err)

		t.Log(dev_vs_prod.LatestMessage[19:])
		verifyResponse, err := authOnlyService.VerifyMail(ctx, &grpc_gen.VerifyMailRequest{
			VerificationCode: dev_vs_prod.LatestMessage[19:],
		},
		)
		assert.Nil(err)
		assert.True(verifyResponse.Ok)
	})

	t.Run("TestGetAccountInfo", func(t *testing.T) {
		assert := asserter.New(t)

		accountInfo, err := authOnlyService.GetAccountInfo(ctx, &grpc_gen.Empty{})
		assert.Nil(err)

		assert.Equal(accountInfo.Mail, "user2@example.com")
		assert.Equal(accountInfo.Name, "user")
		assert.Equal(accountInfo.TelegramUsername, "username")
	})

	t.Run("TestUpdateAccountInfo", func(t *testing.T) {
		assert := asserter.New(t)

		_, err := authOnlyService.UpdateAccountInfo(ctx, &grpc_gen.AccountInfo{
			Name:             "user2",
			Mail:             "user@example.com",
			TelegramUsername: "username2",
		},
		)
		assert.Nil(err)

		t.Log(dev_vs_prod.LatestMessage[19:])
		verifyResponse, err := authOnlyService.VerifyMail(ctx, &grpc_gen.VerifyMailRequest{
			VerificationCode: dev_vs_prod.LatestMessage[19:],
		},
		)
		assert.Nil(err)
		assert.True(verifyResponse.Ok)
		verifyResponse, err = authOnlyService.VerifyMail(ctx, &grpc_gen.VerifyMailRequest{
			VerificationCode: dev_vs_prod.LatestMessage[19:],
		},
		)
		assert.Nil(err)
		assert.False(verifyResponse.Ok)
	})

	t.Run("TestGetAccountInfo2", func(t *testing.T) {
		assert := asserter.New(t)

		accountInfo, err := authOnlyService.GetAccountInfo(ctx, &grpc_gen.Empty{})
		assert.Nil(err)

		assert.Equal(accountInfo.Mail, "user@example.com")
		assert.Equal(accountInfo.Name, "user2")
		assert.Equal(accountInfo.TelegramUsername, "username2")
	})
	t.Run("TestUpdateAccountInfo2", func(t *testing.T) {
		assert := asserter.New(t)

		_, err := authOnlyService.UpdateAccountInfo(ctx, &grpc_gen.AccountInfo{
			Name:             "user",
			Mail:             "user@example.com",
			TelegramUsername: "username",
		},
		)
		assert.Nil(err)

		verifyResponse, err := authOnlyService.VerifyMail(ctx, &grpc_gen.VerifyMailRequest{
			VerificationCode: dev_vs_prod.LatestMessage[19:],
		},
		)
		assert.Nil(err)
		assert.False(verifyResponse.Ok)
	})

	t.Run("TestGetAccountInfo3", func(t *testing.T) {
		assert := asserter.New(t)

		accountInfo, err := authOnlyService.GetAccountInfo(ctx, &grpc_gen.Empty{})
		assert.Nil(err)

		assert.Equal(accountInfo.Mail, "user@example.com")
		assert.Equal(accountInfo.Name, "user")
		assert.Equal(accountInfo.TelegramUsername, "username")
	})
	t.Run("TestProfilePic", func(t *testing.T) {
		assert := asserter.New(t)

		profilePic, err := authOnlyService.GetProfilePic(ctx, &grpc_gen.Empty{})
		assert.Nil(err)
		assert.Condition(func() bool { return len(profilePic.Image) == 0 })

		_, err = authOnlyService.SetProfilePic(ctx, &grpc_gen.ProfilePic{Image: make([]byte, 512*1024+1)})
		assert.NotNil(err)
		_, err = authOnlyService.SetProfilePic(ctx, &grpc_gen.ProfilePic{Image: make([]byte, 512*1024)})
		assert.Nil(err)
		profilePic, err = authOnlyService.GetProfilePic(ctx, &grpc_gen.Empty{})
		assert.Nil(err)
		assert.Condition(func() bool { return len(profilePic.Image) == 512*1024 })
	})

	var bedId *grpc_gen.BedId

	t.Run("TestAddBedAndGetMyBeds", func(t *testing.T) {
		assert := asserter.New(t)

		_, err := authOnlyService.AddBed(ctx, &grpc_gen.BedMutableInfo{
			Address:           "addr",
			Coordinates:       &grpc_gen.Coordinates{Latitude: 10, Longitude: 10},
			Images:            [][]byte{make([]byte, 100)},
			Description:       "descr",
			Features:          []grpc_gen.Feature{grpc_gen.Feature_bathroom},
			MinimumDaysNotice: 5,
		})
		assert.Nil(err)

		bedList, err := authOnlyService.GetMyBeds(ctx, &grpc_gen.Empty{})
		assert.Nil(err)

		bedId = bedList.Beds[0].Id
	})
	t.Run("TestModifyBed", func(t *testing.T) {
		assert := asserter.New(t)

		_, err := authOnlyService.ModifyMyBed(ctx, &grpc_gen.ModifyBedRequest{
			BedId: bedId,
			BedMutableInfo: &grpc_gen.BedMutableInfo{
				Address:           "address",
				Coordinates:       &grpc_gen.Coordinates{Latitude: 0, Longitude: 0},
				Images:            [][]byte{make([]byte, 1000)},
				Description:       "description",
				Features:          []grpc_gen.Feature{grpc_gen.Feature_airConditioner, grpc_gen.Feature_bathroom},
				MinimumDaysNotice: 10,
			},
		})
		assert.Nil(err)
	})
	t.Run("TestAddAvailability", func(t *testing.T) {
		assert := asserter.New(t)

		ti := time.Now()

		_, err := authOnlyService.AddBookingAvailability(ctx, &grpc_gen.BookingAvailability{
			BedId:        bedId,
			DateInterval: &grpc_gen.DateInterval{StartDate: &grpc_gen.Date{Day: uint32(ti.Day()), Month: uint32(ti.Month()), Year: uint32(ti.Year())}, EndDate: &grpc_gen.Date{Day: uint32(ti.Day()), Month: uint32(ti.Month()), Year: uint32(ti.Year())}},
		})
		assert.NotNil(err)

		ti = ti.Add(24 * time.Hour)

		_, err = authOnlyService.AddBookingAvailability(ctx, &grpc_gen.BookingAvailability{
			BedId:        bedId,
			DateInterval: &grpc_gen.DateInterval{StartDate: &grpc_gen.Date{Day: uint32(ti.Day()), Month: uint32(ti.Month()), Year: uint32(ti.Year())}, EndDate: &grpc_gen.Date{Day: uint32(ti.Day()), Month: uint32(ti.Month()), Year: uint32(ti.Year())}},
		})
		assert.Nil(err)

		ti = ti.Add(24 * 89 * time.Hour)

		_, err = authOnlyService.AddBookingAvailability(ctx, &grpc_gen.BookingAvailability{
			BedId:        bedId,
			DateInterval: &grpc_gen.DateInterval{StartDate: &grpc_gen.Date{Day: uint32(ti.Day()), Month: uint32(ti.Month()), Year: uint32(ti.Year())}, EndDate: &grpc_gen.Date{Day: uint32(ti.Day()), Month: uint32(ti.Month()), Year: uint32(ti.Year())}},
		})
		assert.Nil(err)

		ti = ti.Add(24 * time.Hour)

		_, err = authOnlyService.AddBookingAvailability(ctx, &grpc_gen.BookingAvailability{
			BedId:        bedId,
			DateInterval: &grpc_gen.DateInterval{StartDate: &grpc_gen.Date{Day: uint32(ti.Day()), Month: uint32(ti.Month()), Year: uint32(ti.Year())}, EndDate: &grpc_gen.Date{Day: uint32(ti.Day()), Month: uint32(ti.Month()), Year: uint32(ti.Year())}},
		})
		assert.NotNil(err)
	})
	t.Run("TestGetBeds", func(t *testing.T) {
		assert := asserter.New(t)

		ti := time.Now()
		ti = ti.Add(24 * time.Hour)

		res, err := publicService.GetBeds(ctx, &grpc_gen.GetBedsRequest{
			DateRangeLow:  &grpc_gen.Date{Day: uint32(ti.Day()), Month: uint32(ti.Month()), Year: uint32(ti.Year())},
			DateRangeHigh: &grpc_gen.Date{Day: uint32(ti.Day()), Month: uint32(ti.Month()), Year: uint32(ti.Year())},
			Coordinates: &grpc_gen.Coordinates{
				Latitude:  0,
				Longitude: 0,
			},
			FeaturesMandatory: []grpc_gen.Feature{grpc_gen.Feature_airConditioner},
			FromIndex:         0,
		})
		assert.Nil(err)
		assert.NotEmpty(res.Beds)

		res, err = publicService.GetBeds(ctx, &grpc_gen.GetBedsRequest{
			DateRangeLow:  &grpc_gen.Date{Day: uint32(ti.Day()), Month: uint32(ti.Month()), Year: uint32(ti.Year())},
			DateRangeHigh: &grpc_gen.Date{Day: uint32(ti.Day()), Month: uint32(ti.Month()), Year: uint32(ti.Year())},
			Coordinates: &grpc_gen.Coordinates{
				Latitude:  0,
				Longitude: 0,
			},
			FeaturesMandatory: []grpc_gen.Feature{grpc_gen.Feature_bedLinens},
			FromIndex:         0,
		})

		assert.Nil(err)
		assert.Empty(res.Beds)

		res, err = publicService.GetBeds(ctx, &grpc_gen.GetBedsRequest{
			DateRangeLow:  &grpc_gen.Date{Day: uint32(ti.Day()), Month: uint32(ti.Month()), Year: uint32(ti.Year())},
			DateRangeHigh: &grpc_gen.Date{Day: uint32(ti.Day()), Month: uint32(ti.Month()), Year: uint32(ti.Year())},
			Coordinates: &grpc_gen.Coordinates{
				Latitude:  0,
				Longitude: 0,
			},
			FeaturesMandatory: []grpc_gen.Feature{},
			FromIndex:         0,
		})

		assert.Nil(err)
		assert.NotEmpty(res.Beds)
	})
	t.Run("TestGetBed", func(t *testing.T) {
		assert := asserter.New(t)

		_, err := publicService.GetBed(ctx, bedId)
		assert.Nil(err)
	})
	t.Run("TestRemoveAvailability", func(t *testing.T) {
		assert := asserter.New(t)

		ti := time.Now().Add(24 * time.Hour)
		_, err := authOnlyService.RemoveBookAvailability(ctx,
			&grpc_gen.BookingAvailability{
				BedId:        bedId,
				DateInterval: &grpc_gen.DateInterval{StartDate: &grpc_gen.Date{Day: uint32(ti.Day()), Month: uint32(ti.Month()), Year: uint32(ti.Year())}, EndDate: &grpc_gen.Date{Day: uint32(ti.Day()), Month: uint32(ti.Month()), Year: uint32(ti.Year())}}})

		assert.Nil(err)
	})
	t.Run("TestDeleteAccount", func(t *testing.T) {
		assert := asserter.New(t)

		_, err := authOnlyService.DeleteAccount(ctx, &grpc_gen.Empty{})
		assert.Nil(err)
		_, err = authOnlyService.GetAccountInfo(ctx, &grpc_gen.Empty{})
		assert.NotNil(err)
	})
}
