package main

import (
	"context"
	"log"
	"net"
	"time"

	"github.com/SnoozeHub/snoozehub-backend/grpc_gen"
	"github.com/go-co-op/gocron"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

type Trainer struct {
	Name string
	Age  int
	City string
}

func setupDb(db *mongo.Database) error {
	/*
		_, err := db.Collection("accounts").Indexes().CreateOne(
			context.Background(),
			mongo.IndexModel{
				Keys:    bson.D{{Key: "publicKey", Value: 1}},
				Options: options.Index().SetUnique(true),
			},
		)
		if err != nil {
			return err
		}
	*/
	return nil
}
func KeepRemovingInvalidAvailabilities(db *mongo.Database) {

}
func runGrpc() error {
	lis, err := net.Listen("tcp", ":9090")

	if err != nil {
		return err
	}

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://root:root@mongodb:27017"))
	if err != nil {
		return err
	}
	defer client.Disconnect(context.Background())
	err = client.Ping(context.Background(), nil)
	if err != nil {
		return err
	}

	db := client.Database("main")

	err = setupDb(db)
	if err != nil {
		return err
	}

	s := grpc.NewServer()
	tmp := newAuthOnlyService(db)
	grpc_gen.RegisterPublicServiceServer(s, newPublicService(db, tmp.GetAuthTokens(), tmp.GetMutex()))
	grpc_gen.RegisterAuthOnlyServiceServer(s, tmp)

	// keep removing invalid (old) availabilities
	scheduler := gocron.NewScheduler(time.Local)
	scheduler.Every(1).Day().Do(func() {
		now := time.Now()
		todayFlat := flatterizeDate(&grpc_gen.Date{Day: uint32(now.Day()), Month: uint32(now.Month()), Year: uint32(now.Year())})
		update := bson.M{"$pull": bson.M{"dateAvailables": bson.M{"$lte": todayFlat}}}
		db.Collection("beds").UpdateMany(context.Background(), bson.D{}, update)
	})
	now := time.Now()
	scheduler.StartAt(time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())).StartAsync()

	return s.Serve(lis)
}
func main() {
	log.Println("Server has started!")
	defer log.Println("Server stopped!")
	if err := runGrpc(); err != nil {
		log.Fatal(err)
	}
}
