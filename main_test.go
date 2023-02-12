package main

import (
	"context"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	"math/big"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Test(t *testing.T) {
	client, _ := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://mongodb:27017"))
	db := client.Database("main")
	filter := bson.D{}
	filter = append(filter, bson.E{Key: "test", Value: bson.M{"$all": bson.A{2, 3}}})
	res, err := db.Collection("accounts").Find(
		context.TODO(),
		filter,
	)
	if err != nil {
		t.Fatal("ERROR: ", err)
	} else {
		for res.Next(context.TODO()) {
			fmt.Println(res.Current.String())
		}
	}
}




func TestMail243(t *testing.T) {
	client, _ := ethclient.Dial("wss://goerli.infura.io/ws/v3/9aa3d95b3bc440fa88ea12eaa4456161")
	query := ethereum.FilterQuery{
		Addresses: []common.Address{common.HexToAddress("0xc2Cd631A73E0D94dE392F686D22E9e792E426000")},
	}

	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)

	if err != nil {
		t.Fatal(err)
	}

	timeout := time.After(10 * time.Second)

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
			from := common.BytesToAddress(log.Topics[1].Bytes())
			to := common.BytesToAddress(log.Topics[2].Bytes())
			t.Log(from)
			t.Log(to)
			t.Log(amount)
			return
		}
	}
}
