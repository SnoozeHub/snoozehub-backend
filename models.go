package main

import "go.mongodb.org/mongo-driver/bson/primitive"

type account struct {
	PublicKey        string
	Name             string
	Mail             string
	TelegramUsername string
	ProfilePic       []byte
	VerificationCode *string // If doesn't exist, the account is verified
	BedIdsBookings   []string
}

type review struct {
	Reviewer   primitive.ObjectID
	Evaluation int32
	Comment    string
}

type bed struct {
	Host              primitive.ObjectID
	Id                string
	Address           string
	Latitude          float64
	Longitude         float64
	Images            [][]byte
	Description       string
	Features          []int32
	MinimumDaysNotice int32
	DateAvailables    []int32
	AverageEvaluation *int32
	Reviews           []review
}
