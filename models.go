package main

type account struct {
	PublicKey string`bson:"publicKey"`
}

type review struct {
	Evaluation int32  `bson:"evaluation"`
	Comment    string `bson:"comment"`
}

type bed struct {
	Id                string   `bson:"id"`
	Place             string   `bson:"place"`
	Images            [][]byte `bson:"images"`
	Description       string   `bson:"description"`
	Features          []int32  `bson:"features"`
	MinimumDaysNotice int32    `bson:"minimumDaysNotice"`
	DateAvailables    []int32  `bson:"dateAvailables"`
	AverageEvaluation int32    `bson:"averageEvaluation"`
	Reviews           []review `bson:"reviews"`
}