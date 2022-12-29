package main

type account struct {
	PublicKey string
	Name string
	Mail string
	TelegramUsername string
	ProfilePic *[]byte
	VerificationCode *string // If doesn't exist, the account is verified
}

type review struct {
	Evaluation int32
	Comment    string
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
