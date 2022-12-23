package main

import (
	"context"

	"github.com/SnoozeHub/snoozehub-backend/grpc_gen"
)

type publicService struct {
	grpc_gen.UnimplementedPublicServiceServer
}

func newPublicService() *publicService {
	service := publicService{}
	return &service
}

func (s *publicService) GetNonce(context.Context, *grpc_gen.Empty) (*grpc_gen.GetNonceResponse, error) {
	return nil, nil
}
func (s *publicService) Auth(context.Context, *grpc_gen.AuthRequest) (*grpc_gen.AuthResponse, error) {
	return nil, nil
}
func (s *publicService) GetBeds(context.Context, *grpc_gen.GetBedsRequest) (*grpc_gen.BedList, error) {
	return nil, nil
}
func (s *publicService) GetReview(context.Context, *grpc_gen.GetReviewsRequest) (*grpc_gen.GetReviewsResponse, error) {
	return nil, nil
}