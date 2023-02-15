package main

import (
	"context"
	"strconv"
	"testing"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/SnoozeHub/snoozehub-backend/grpc_gen"
	"github.com/SnoozeHub/snoozehub-backend/mail"
	asserter "github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc/metadata"
)

func TestDbConection(t *testing.T) {
	_, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://mongodb:27017"))
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
	assert := asserter.New(t)
	go main()
	conn, err := grpc.Dial("localhost:9090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	assert.Nil(err)

	publicService := grpc_gen.NewPublicServiceClient(conn)
	authOnlyService := grpc_gen.NewAuthOnlyServiceClient(conn)

	var token string

	t.Run("TestAuthentication", func(t *testing.T) {
		assert := asserter.New(t)

		getNonceResponse, err := publicService.GetNonce(context.Background(), &grpc_gen.Empty{})
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
		metadata.Pairs("authToken", token),
	)

	t.Run("TestSignUpAndVerify", func(t *testing.T) {
		assert := asserter.New(t)

		_, err := authOnlyService.SignUp(ctx, &grpc_gen.AccountInfo{
			Name:             "user",
			Mail:             "user@example.com",
			TelegramUsername: "username",
		},
		)
		assert.Nil(err)

		t.Log(mail.LatestMessage[19:])
		_, err = authOnlyService.VerifyMail(ctx, &grpc_gen.VerifyMailRequest{
			VerificationCode: mail.LatestMessage[19:],
		},
		)
		assert.Nil(err)
	})
}
