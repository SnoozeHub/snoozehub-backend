package main

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"math/big"
	"sort"
	"strconv"
	"sync"
	"time"

	"github.com/SnoozeHub/snoozehub-backend/grpc_gen"
	"github.com/ethereum/go-ethereum/crypto"
	geo "github.com/kellydunn/golang-geo"
	"github.com/patrickmn/go-cache"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type publicService struct {
	grpc_gen.UnimplementedPublicServiceServer
	nonces     *cache.Cache
	authTokens *cache.Cache
	db         *mongo.Database
	mu         *sync.Mutex
}

func newPublicService(db *mongo.Database, authTokens *cache.Cache, mu *sync.Mutex) *publicService {
	service := publicService{
		nonces:     cache.New(1*time.Minute, 1*time.Second),
		authTokens: authTokens,
		db:         db,
		mu:         mu,
	}
	return &service
}

func (s *publicService) GetNonce(context.Context, *grpc_gen.Empty) (*grpc_gen.GetNonceResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	nonce := GenRandomString(20)
	s.nonces.Add(nonce, struct{}{}, cache.DefaultExpiration)
	return &grpc_gen.GetNonceResponse{Nonce: nonce}, nil
}
func (s *publicService) Auth(_ context.Context, req *grpc_gen.AuthRequest) (*grpc_gen.AuthResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, found := s.nonces.Get(req.Nonce)
	if !found {
		return nil, errors.New("expired or non existing nonce")
	}

	nonceHashed := crypto.Keccak256Hash([]byte("\x19Ethereum Signed Message:\n" + strconv.Itoa(len(req.Nonce)) + req.Nonce))
	UncompressedPublicKeyBytes, err := crypto.Ecrecover(nonceHashed.Bytes(), req.SignedNonce)
	if err != nil {
		return nil, errors.New("invalid signedNonce")
	}

	x := new(big.Int).SetBytes(UncompressedPublicKeyBytes[1:33])
	y := new(big.Int).SetBytes(UncompressedPublicKeyBytes[33:])
	publicKey := crypto.PubkeyToAddress(ecdsa.PublicKey{Curve: crypto.S256(), X: x, Y: y}).String()

	token := GenRandomString(20)

	s.authTokens.Add(token, publicKey, cache.DefaultExpiration)

	res, err := s.db.Collection("accounts").Find(
		context.Background(),
		bson.D{{Key: "publicKey", Value: publicKey}},
	)
	if err != nil {
		return nil, err
	}
	accountExist := res.Next(context.Background())
	return &grpc_gen.AuthResponse{
		AuthToken:    token,
		AccountExist: accountExist,
	}, nil
}
func (s *publicService) GetBeds(_ context.Context, req *grpc_gen.GetBedsRequest) (*grpc_gen.BedList, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Prepare the mondadory features
	if !allDistinct(req.FeaturesMandatory) {
		return nil, errors.New("not distinct featuresMandatory")
	}

	// Check if coordinates are valid
	if req.Coordinates.Latitude < -90 || req.Coordinates.Latitude > 90 || req.Coordinates.Longitude < -180 || req.Coordinates.Latitude > 180 {
		return nil, errors.New("invalid coordinates")
	}

	// Filter the mondadory features
	tmp := bson.A{}
	for _, v := range req.FeaturesMandatory {
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
		context.Background(),
		filter,
	)
	if err != nil {
		return nil, err
	}

	// sort beds based on coordinates
	var resSorted []bed
	err = res.All(context.Background(), &resSorted)
	if err != nil {
		return nil, err
	}
	requestedLocation := geo.NewPoint(req.Coordinates.Latitude, req.Coordinates.Longitude)
	sort.Slice(resSorted, func(i, j int) bool {
		return requestedLocation.GreatCircleDistance(geo.NewPoint(resSorted[i].Latitude, resSorted[i].Longitude)) < requestedLocation.GreatCircleDistance(geo.NewPoint(resSorted[j].Latitude, resSorted[j].Longitude))
	})

	// Get first N result from req.FromIndex
	beds := make([]*grpc_gen.Bed, 0)
	for i := int(req.FromIndex); i < len(resSorted); i++ {
		beds = append(beds, bedToGrpcBed(s.db, resSorted[i]))
	}
	return &grpc_gen.BedList{Beds: beds}, nil
}
func (s *publicService) GetBed(_ context.Context, req *grpc_gen.BedId) (*grpc_gen.GetBedResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	res, err := s.db.Collection("beds").Find(
		context.Background(),
		bson.D{{Key: "_id", Value: req.BedId}},
	)
	if err != nil {
		return nil, err
	}
	if res.RemainingBatchLength() != 1 {
		return &grpc_gen.GetBedResponse{Bed: nil}, nil
	}
	res.Next(context.Background())
	var currentBed bed
	err = bson.Unmarshal(res.Current, currentBed)
	if err != nil {
		return nil, err
	}
	
	return &grpc_gen.GetBedResponse{Bed: bedToGrpcBed(s.db, currentBed)}, nil
}
func (s *publicService) GetReview(_ context.Context, req *grpc_gen.GetReviewsRequest) (*grpc_gen.GetReviewsResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	res, err := s.db.Collection("beds").Find(
		context.Background(),
		bson.D{{Key: "_id", Value: req.BedId.BedId}},
	)
	if err != nil {
		return nil, err
	}
	if res.RemainingBatchLength() != 1 {
		return nil, errors.New("invalid id")
	}
	res.Next(context.Background())
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
