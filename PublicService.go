package main

import (
	"context"
	"errors"
	"time"

	"github.com/SnoozeHub/snoozehub-backend/grpc_gen"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/patrickmn/go-cache"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type publicService struct {
	grpc_gen.UnimplementedPublicServiceServer
	nonces *cache.Cache
	authTokens *cache.Cache
	db *mongo.Database
}

func newPublicService(db *mongo.Database, authTokens *cache.Cache) *publicService {
	service := publicService{
		nonces: cache.New(20*time.Second, 20*time.Second),
		authTokens: authTokens,
		db: db,
	}
	return &service
}

func (s *publicService) GetNonce(context.Context, *grpc_gen.Empty) (*grpc_gen.GetNonceResponse, error) {
	nonce := GenRandomString(20)
	s.nonces.Add(nonce, struct{}{}, cache.DefaultExpiration)
	return &grpc_gen.GetNonceResponse{Nonce: nonce}, nil
}
func (s *publicService) Auth(_ context.Context, req *grpc_gen.AuthRequest) (*grpc_gen.AuthResponse, error) {
	_, found := s.nonces.Get(req.Nonce)
	if !found {
		return nil, errors.New("expired or non existing nonce")
	}

	nonceHashed := crypto.Keccak256Hash([]byte(req.Nonce))
	publicKeyBytes, err := crypto.Ecrecover(nonceHashed.Bytes(), []byte(req.SignedNonce))
	if err != nil {
		return nil, errors.New("invalid signedNonce")
	}
	publicKey := string(publicKeyBytes)
	token := GenRandomString(20)
	s.authTokens.Add(token, publicKey, cache.DefaultExpiration)

	res, _ := s.db.Collection("accounts").Find(
		context.TODO(),
		bson.D{{"publicKey", publicKey}},
	)
	accountExist := res.Next(context.TODO())
	return &grpc_gen.AuthResponse{
		AuthToken: token,
		AccountExist: accountExist,
	}, nil
}
func (s *publicService) GetBeds(_ context.Context, req *grpc_gen.GetBedsRequest) (*grpc_gen.BedList, error) {
	// Prepare the mondadory features
	if !allDistinct(req.FeaturesMondadory) {
		return nil, errors.New("not distinct featuresMondatory")
	}
	featuresRequested := featuresToInts(req.FeaturesMondadory)

	// Filter the place
	if len(req.Place) < 1 || len(req.Place) > 100 {
		return nil, errors.New("invalid place")
	}
	filter := bson.D{
		{Key: "place", Value: req.Place},
	}

	// Filter the mondadory features
	tmp := bson.A{}
	for _, v := range featuresRequested {
		tmp = append(tmp, v)
	}
	filter = append(filter, bson.E{Key: "features", Value: bson.M{"$all": tmp}})
	
	// Filter the date range
	if !isDateValidAndFromTomorrow(req.DateRangeLow) || !isDateValidAndFromTomorrow(req.DateRangeHigh) {
		return nil, errors.New("invalid dateRangeLow or dateRangeHigh")
	}
	dateRangeLowFlatterized := flatterizeDate(req.DateRangeLow)
	dateRangeHighFlatterized := flatterizeDate(req.DateRangeHigh)
	if dateRangeHighFlatterized < dateRangeLowFlatterized {
		return nil, errors.New("dateRangeHigh is before dateRangeLow")
	}
	filter = append(filter, bson.E{Key: "dateAvailables", Value: bson.M{"$elemMatch": bson.A{bson.M{"$gte": dateRangeLowFlatterized}, bson.M{"$lte": dateRangeHighFlatterized}}}})
	
	res, _ := s.db.Collection("beds").Find(
		context.TODO(),
		filter,
	)
	
	beds := make([]*grpc_gen.BedList_Bed, 0)
	i := req.FromIndex
	for res.Next(context.TODO()) && i > 0{
		i--
	}
	i = 15
	for res.Next(context.TODO()) && i > 0{
		var tmp bed
		err := bson.Unmarshal(res.Current, tmp)
		if err != nil {
			return nil, err
		}
		beds = append(beds, &grpc_gen.BedList_Bed{
			Id: &grpc_gen.BedId{BedId: tmp.Id},
			BedMutableInfo: &grpc_gen.BedMutableInfo{
				Place: tmp.Place,
				Images: tmp.Images,
				Description: tmp.Description,
				Features: intsTofeatures(tmp.Features),
			},
			DateAvailables: tmp.DateAvailables,
		})
		i--
	}
	return nil, nil
}
func (s *publicService) GetReview(context.Context, *grpc_gen.GetReviewsRequest) (*grpc_gen.GetReviewsResponse, error) {
	return nil, nil
}