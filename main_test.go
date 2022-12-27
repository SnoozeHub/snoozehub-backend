package main

import (
	"context"
	"fmt"
	"testing"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Test(t *testing.T) {
	client, _ := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://mongodb:27017"))
	db := client.Database("main")
	filter := bson.D{}
	filter = append(filter, bson.E{"test", bson.M{"$all": bson.A{2, 3}}})
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
