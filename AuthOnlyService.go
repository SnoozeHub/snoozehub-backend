package main

import (
	"context"
	"time"

	"github.com/SnoozeHub/snoozehub-backend/grpc_gen"
	"github.com/patrickmn/go-cache"
	"go.mongodb.org/mongo-driver/mongo"
)

type authOnlyService struct {
	grpc_gen.UnimplementedAuthOnlyServiceServer
	authTokens *cache.Cache
	db *mongo.Database
}

func newAuthOnlyService(db *mongo.Database) *authOnlyService {
	service := authOnlyService{
		authTokens: cache.New(24*time.Hour, 24*time.Hour),
		db: db,
	}
	return &service
}
func (s *authOnlyService) GetAuthTokens() *cache.Cache {
	return s.authTokens
}

func (s *authOnlyService) SignUp(context.Context, *grpc_gen.AccountInfo) (*grpc_gen.Empty, error) {
	return nil, nil
}
func (s *authOnlyService) VerifyMail(context.Context, *grpc_gen.VerifyMailRequest) (*grpc_gen.VerifyMailResponse, error) {
	return nil, nil
}
func (s *authOnlyService) GetAccountInfo(context.Context, *grpc_gen.Empty) (*grpc_gen.AccountInfo, error) {
	return nil, nil
}
func (s *authOnlyService) GetProfilePic(context.Context, *grpc_gen.Empty) (*grpc_gen.ProfilePic, error) {
	return nil, nil
}
func (s *authOnlyService) SetProfilePic(context.Context, *grpc_gen.ProfilePic) (*grpc_gen.Empty, error) {
	return nil, nil
}
func (s *authOnlyService) DeleteAccount(context.Context, *grpc_gen.Empty) (*grpc_gen.Empty, error) {
	return nil, nil
}
func (s *authOnlyService) UpdateAccountInfo(context.Context, *grpc_gen.AccountInfo) (*grpc_gen.Empty, error) {
	return nil, nil
}
func (s *authOnlyService) Book(context.Context, *grpc_gen.Booking) (*grpc_gen.BookResponse, error) {
	return nil, nil
}
func (s *authOnlyService) GetMyBookings(context.Context, *grpc_gen.Empty) (*grpc_gen.GetBookingsResponse, error) {
	return nil, nil
}
func (s *authOnlyService) Review(context.Context, *grpc_gen.ReviewRequest) (*grpc_gen.Empty, error) {
	return nil, nil
}
func (s *authOnlyService) RemoveReview(context.Context, *grpc_gen.BedId) (*grpc_gen.Empty, error) {
	return nil, nil
}
func (s *authOnlyService) AddBed(context.Context, *grpc_gen.BedMutableInfo) (*grpc_gen.Empty, error) {
	return nil, nil
}
func (s *authOnlyService) ModifyMyBed(context.Context, *grpc_gen.ModifyBedRequest) (*grpc_gen.Empty, error) {
	return nil, nil
}
func (s *authOnlyService) RemoveMyBed(context.Context, *grpc_gen.BedId) (*grpc_gen.Empty, error) {
	return nil, nil
}
func (s *authOnlyService) GetMyBeds(context.Context, *grpc_gen.Empty) (*grpc_gen.BedList, error) {
	return nil, nil
}
func (s *authOnlyService) AddBookingAvaiability(context.Context, *grpc_gen.Booking) (*grpc_gen.Empty, error) {
	return nil, nil
}
func (s *authOnlyService) RemoveBookAvaiability(context.Context, *grpc_gen.Booking) (*grpc_gen.Empty, error) {
	return nil, nil
}