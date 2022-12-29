package main

import (
	"context"
	"errors"
	"sort"
	"time"
	geo "github.com/kellydunn/golang-geo"
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

	res, err := s.db.Collection("accounts").Find(
		context.TODO(),
		bson.D{{Key: "publicKey", Value: publicKey}},
	)
	if err != nil {
		return nil, err
	}
	accountExist := res.Next(context.TODO())
	return &grpc_gen.AuthResponse{
		AuthToken: token,
		AccountExist: accountExist,
	}, nil
}
func (s *publicService) GetBeds(_ context.Context, req *grpc_gen.GetBedsRequest) (*grpc_gen.BedList, error) {
	// Prepare the mondadory features
	if !allDistinct(req.FeaturesMandatory) {
		return nil, errors.New("not distinct featuresMandatory")
	}
	featuresRequested := featuresToInts(req.FeaturesMandatory)

	// Check if coordinates are valid
	if req.Coordinates.Latitude < -90 || req.Coordinates.Latitude > 90 || req.Coordinates.Longitude < -180 || req.Coordinates.Latitude > 180 {
		return nil, errors.New("invalid coordinates")
	}
	
	// Filter the mondadory features
	tmp := bson.A{}
	for _, v := range featuresRequested {
		tmp = append(tmp, v)
	}
	filter := bson.D{bson.E{Key: "features", Value: bson.M{"$all": tmp}}}
	
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
	
	// get beds
	res, err := s.db.Collection("beds").Find(
		context.TODO(),
		filter,
	)
	if err != nil {
		return nil, err
	}

	// sort beds based on coordinates
	resSorted := make([]bed, res.RemainingBatchLength())
	err = res.All(context.TODO(), resSorted)
	if err != nil {
		return nil, err
	}
	requestedLocation := geo.NewPoint(req.Coordinates.Latitude, req.Coordinates.Longitude)
	sort.Slice(resSorted, func(i, j int) bool {
		return requestedLocation.GreatCircleDistance(geo.NewPoint(resSorted[i].Latitude, resSorted[i].Longitude)) < requestedLocation.GreatCircleDistance(geo.NewPoint(resSorted[j].Latitude, resSorted[j].Longitude))
	})
	
	// Get first N result from req.FromIndex
	beds := make([]*grpc_gen.BedList_Bed, 0)
	for i := int(req.FromIndex); i < len(resSorted); i++ {
		tmp := resSorted[i]
		dateAvailables := make([]*grpc_gen.Date, len(tmp.DateAvailables))
		for _, v := range tmp.DateAvailables {
			dateAvailables = append(dateAvailables, deflatterizeDate(v))
		}
		var averageEvaluation *uint32 = nil
		if tmp.AverageEvaluation != nil {
			tmp2 := uint32(*tmp.AverageEvaluation)
			averageEvaluation = &tmp2
		}
		beds = append(beds, &grpc_gen.BedList_Bed{
			Id: &grpc_gen.BedId{BedId: tmp.Id},
			BedMutableInfo: &grpc_gen.BedMutableInfo{
				Address: tmp.Address,
				Coordinates: &grpc_gen.Coordinates{Latitude: tmp.Latitude, Longitude: tmp.Longitude},
				Images: tmp.Images,
				Description: tmp.Description,
				Features: intsTofeatures(tmp.Features),
				MinimumDaysNotice: uint32(tmp.MinimumDaysNotice),
			},
			DateAvailables: dateAvailables,
			ReviewCount: uint32(len(tmp.Reviews)),
			AverageEvaluation: averageEvaluation,
		})
	}
	return &grpc_gen.BedList{Beds: beds}, nil
}
func (s *publicService) GetReview(_ context.Context, req *grpc_gen.GetReviewsRequest) (*grpc_gen.GetReviewsResponse, error) {
	res, err := s.db.Collection("beds").Find(
		context.TODO(),
		bson.D{{Key: "id", Value: req.BedId.BedId}},
	)
	if err != nil {
		return nil, err
	}
	if res.RemainingBatchLength() != 1 {
		return nil, errors.New("invalid id")
	}
	res.Next(context.TODO())
	var currentBed bed
	err = bson.Unmarshal(res.Current, currentBed)
	if err != nil {
		return nil, err
	}

	reviews := make([]*grpc_gen.Review, 0)
	for i := int(req.FromIndex); i < len(currentBed.Reviews); i++ {
		reviews = append(reviews, &grpc_gen.Review{Evaluation: uint32(currentBed.Reviews[i].Evaluation), Comment: currentBed.Reviews[i].Comment})
	}
	return &grpc_gen.GetReviewsResponse{Reviews: reviews}, nil
}