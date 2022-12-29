package main

type account struct {
	PublicKey string `bson:"publicKey"`
}

type review struct {
	Evaluation int32  `bson:"evaluation"`
	Comment    string `bson:"comment"`
}

type bed struct {
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
