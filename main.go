package main

import (
	"context"
	"log"
	"net"

	"github.com/SnoozeHub/snoozehub-backend/grpc_gen"
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
	_, err := db.Collection("accounts").Indexes().CreateOne(
		context.TODO(),
		mongo.IndexModel{
			Keys:    bson.D{{Key: "publicKey", Value: 1}},
			Options: options.Index().SetUnique(true),
		},
	)
	if err != nil {
		return err
	}
	_, err = db.Collection("beds").Indexes().CreateOne(
		context.TODO(),
		mongo.IndexModel{
			Keys: bson.D{{Key: "place", Value: 1}},
		},
	)
	if err != nil {
		return err
	}
	return nil
}
func runGrpc() error {
	lis, err := net.Listen("tcp", ":9090")
	if err != nil {
		return err
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://mongodb:27017"))
	if err != nil {
		return err
	}
	defer client.Disconnect(context.TODO())
	err = client.Ping(context.TODO(), nil)
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
	return s.Serve(lis)
}
func main() {
	log.Println("Server has started!")
	defer log.Println("Server stopped!")
	if err := runGrpc(); err != nil {
		log.Fatal(err)
	}
}
