package main

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"time"

	"github.com/SnoozeHub/snoozehub-backend/grpc_gen"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/patrickmn/go-cache"
)

type publicService struct {
	grpc_gen.UnimplementedPublicServiceServer
	nonces *cache.Cache
}

func newPublicService() *publicService {
	service := publicService{
		nonces: cache.New(20*time.Second, 5*time.Second),
	}
	return &service
}

func (s *publicService) GetNonce(context.Context, *grpc_gen.Empty) (*grpc_gen.GetNonceResponse, error) {
	nonce := GenRandomString(20)
	s.nonces.Add(nonce, struct{}{}, cache.DefaultExpiration)
	return &grpc_gen.GetNonceResponse{Nonce: nonce}, nil
}
func (s *publicService) Auth(_ context.Context, authRequest *grpc_gen.AuthRequest) (*grpc_gen.AuthResponse, error) {
	_, found := s.nonces.Get(authRequest.Nonce)
	if !found {
		return nil, errors.New("Expired or non existing nonce")
	}

	nonceHashed := crypto.Keccak256Hash([]byte(authRequest.Nonce))
	publicKey, err := crypto.Ecrecover(nonceHashed.Bytes(), []byte(authRequest.SignedNonce))
	if err != nil {
		return nil, errors.New("Invalid signedNonce")
	}
	return nil, nil
}
func (s *publicService) GetBeds(context.Context, *grpc_gen.GetBedsRequest) (*grpc_gen.BedList, error) {
	return nil, nil
}
func (s *publicService) GetReview(context.Context, *grpc_gen.GetReviewsRequest) (*grpc_gen.GetReviewsResponse, error) {
	return nil, nil
}