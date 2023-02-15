package main

import "go.mongodb.org/mongo-driver/bson/primitive"

type account struct {
	PublicKey        string               `bson:"publicKey"`
	Name             string               `bson:"name"`
	Mail             string               `bson:"mail"`
	TelegramUsername string               `bson:"telegramUsername"`
	ProfilePic       []byte               `bson:"profilePic"`
	VerificationCode *string              `bson:"verificationCode"` // If doesn't exist, the account is verified
	BedIdBookings    []primitive.ObjectID `bson:"bedIdBookings"`
}

type review struct {
	Reviewer   primitive.ObjectID `bson:"reviewer"`
	Evaluation int32              `bson:"evaluation"`
	Comment    string             `bson:"comment"`
}

type bed struct {
	Id                primitive.ObjectID `bson:"_id"`
	Host              primitive.ObjectID `bson:"host"`
	Address           string             `bson:"address"`
	Latitude          float64            `bson:"latitude"`
	Longitude         float64            `bson:"longitude"`
	Images            [][]byte           `bson:"images"`
	Description       string             `bson:"description"`
	Features          []int32            `bson:"features"`
	MinimumDaysNotice int32              `bson:"minimumDaysNotice"`
	DateAvailables    []int32            `bson:"dateAvailables"`
	AverageEvaluation *int32             `bson:"averageEvaluation"`
	Reviews           []review           `bson:"reviews"`
}
