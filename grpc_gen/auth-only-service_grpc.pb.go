// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.6
// source: auth-only-service.proto

package grpc_gen

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// AuthOnlyServiceClient is the client API for AuthOnlyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthOnlyServiceClient interface {
	SignUp(ctx context.Context, in *AccountInfo, opts ...grpc.CallOption) (*Empty, error)
	Logout(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
	// Returns ok=false also if the account is already verified
	VerifyMail(ctx context.Context, in *VerifyMailRequest, opts ...grpc.CallOption) (*VerifyMailResponse, error)
	GetAccountInfo(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*AccountInfo, error)
	GetProfilePic(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ProfilePic, error)
	SetProfilePic(ctx context.Context, in *ProfilePic, opts ...grpc.CallOption) (*Empty, error)
	// It also delete all his beds and booking availabilities, review to other's beds (adjusting averageEvaluation)
	DeleteAccount(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
	// If the mail is changed, is needed to verify the account (and the eventual old verification code becomed invalid), you can call this function even if
	// the account is not verified, because for example the mail was wrong.
	UpdateAccountInfo(ctx context.Context, in *AccountInfo, opts ...grpc.CallOption) (*Empty, error)
	// GUEST RPCs
	// If guest can pay, he must do it within 1 minute
	// "Human proof token" are then sent through mail to both guest, and host
	// If the guest account will be deleted in the following minute or the booking becomes invalid, the will be no booking
	Book(ctx context.Context, in *Booking, opts ...grpc.CallOption) (*BookResponse, error)
	Review(ctx context.Context, in *ReviewRequest, opts ...grpc.CallOption) (*Empty, error)
	// It returns the optional own review for the BedId
	GetMyReview(ctx context.Context, in *BedId, opts ...grpc.CallOption) (*GetMyReviewResponse, error)
	//rpc GetMyBookings(Empty) returns (BookingInfoList); // NOT GOING TO BE IMPLEMENTED
	RemoveReview(ctx context.Context, in *BedId, opts ...grpc.CallOption) (*Empty, error)
	// HOST RPCs
	AddBed(ctx context.Context, in *BedMutableInfo, opts ...grpc.CallOption) (*Empty, error)
	ModifyMyBed(ctx context.Context, in *ModifyBedRequest, opts ...grpc.CallOption) (*Empty, error)
	RemoveMyBed(ctx context.Context, in *BedId, opts ...grpc.CallOption) (*Empty, error)
	GetMyBeds(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*BedList, error)
	AddBookingAvailability(ctx context.Context, in *BookingAvailability, opts ...grpc.CallOption) (*Empty, error)
	//rpc ModifyBookingAvailability(BookingAvailability) returns (Empty); // NOT GOING TO BE IMPLEMENTED!!
	RemoveBookAvailability(ctx context.Context, in *BookingAvailability, opts ...grpc.CallOption) (*Empty, error)
}

type authOnlyServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthOnlyServiceClient(cc grpc.ClientConnInterface) AuthOnlyServiceClient {
	return &authOnlyServiceClient{cc}
}

func (c *authOnlyServiceClient) SignUp(ctx context.Context, in *AccountInfo, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/AuthOnlyService/SignUp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authOnlyServiceClient) Logout(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/AuthOnlyService/logout", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authOnlyServiceClient) VerifyMail(ctx context.Context, in *VerifyMailRequest, opts ...grpc.CallOption) (*VerifyMailResponse, error) {
	out := new(VerifyMailResponse)
	err := c.cc.Invoke(ctx, "/AuthOnlyService/VerifyMail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authOnlyServiceClient) GetAccountInfo(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*AccountInfo, error) {
	out := new(AccountInfo)
	err := c.cc.Invoke(ctx, "/AuthOnlyService/GetAccountInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authOnlyServiceClient) GetProfilePic(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ProfilePic, error) {
	out := new(ProfilePic)
	err := c.cc.Invoke(ctx, "/AuthOnlyService/GetProfilePic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authOnlyServiceClient) SetProfilePic(ctx context.Context, in *ProfilePic, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/AuthOnlyService/SetProfilePic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authOnlyServiceClient) DeleteAccount(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/AuthOnlyService/DeleteAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authOnlyServiceClient) UpdateAccountInfo(ctx context.Context, in *AccountInfo, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/AuthOnlyService/UpdateAccountInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authOnlyServiceClient) Book(ctx context.Context, in *Booking, opts ...grpc.CallOption) (*BookResponse, error) {
	out := new(BookResponse)
	err := c.cc.Invoke(ctx, "/AuthOnlyService/Book", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authOnlyServiceClient) Review(ctx context.Context, in *ReviewRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/AuthOnlyService/Review", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authOnlyServiceClient) GetMyReview(ctx context.Context, in *BedId, opts ...grpc.CallOption) (*GetMyReviewResponse, error) {
	out := new(GetMyReviewResponse)
	err := c.cc.Invoke(ctx, "/AuthOnlyService/GetMyReview", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authOnlyServiceClient) RemoveReview(ctx context.Context, in *BedId, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/AuthOnlyService/RemoveReview", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authOnlyServiceClient) AddBed(ctx context.Context, in *BedMutableInfo, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/AuthOnlyService/AddBed", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authOnlyServiceClient) ModifyMyBed(ctx context.Context, in *ModifyBedRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/AuthOnlyService/ModifyMyBed", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authOnlyServiceClient) RemoveMyBed(ctx context.Context, in *BedId, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/AuthOnlyService/RemoveMyBed", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authOnlyServiceClient) GetMyBeds(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*BedList, error) {
	out := new(BedList)
	err := c.cc.Invoke(ctx, "/AuthOnlyService/GetMyBeds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authOnlyServiceClient) AddBookingAvailability(ctx context.Context, in *BookingAvailability, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/AuthOnlyService/AddBookingAvailability", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authOnlyServiceClient) RemoveBookAvailability(ctx context.Context, in *BookingAvailability, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/AuthOnlyService/RemoveBookAvailability", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthOnlyServiceServer is the server API for AuthOnlyService service.
// All implementations must embed UnimplementedAuthOnlyServiceServer
// for forward compatibility
type AuthOnlyServiceServer interface {
	SignUp(context.Context, *AccountInfo) (*Empty, error)
	Logout(context.Context, *Empty) (*Empty, error)
	// Returns ok=false also if the account is already verified
	VerifyMail(context.Context, *VerifyMailRequest) (*VerifyMailResponse, error)
	GetAccountInfo(context.Context, *Empty) (*AccountInfo, error)
	GetProfilePic(context.Context, *Empty) (*ProfilePic, error)
	SetProfilePic(context.Context, *ProfilePic) (*Empty, error)
	// It also delete all his beds and booking availabilities, review to other's beds (adjusting averageEvaluation)
	DeleteAccount(context.Context, *Empty) (*Empty, error)
	// If the mail is changed, is needed to verify the account (and the eventual old verification code becomed invalid), you can call this function even if
	// the account is not verified, because for example the mail was wrong.
	UpdateAccountInfo(context.Context, *AccountInfo) (*Empty, error)
	// GUEST RPCs
	// If guest can pay, he must do it within 1 minute
	// "Human proof token" are then sent through mail to both guest, and host
	// If the guest account will be deleted in the following minute or the booking becomes invalid, the will be no booking
	Book(context.Context, *Booking) (*BookResponse, error)
	Review(context.Context, *ReviewRequest) (*Empty, error)
	// It returns the optional own review for the BedId
	GetMyReview(context.Context, *BedId) (*GetMyReviewResponse, error)
	//rpc GetMyBookings(Empty) returns (BookingInfoList); // NOT GOING TO BE IMPLEMENTED
	RemoveReview(context.Context, *BedId) (*Empty, error)
	// HOST RPCs
	AddBed(context.Context, *BedMutableInfo) (*Empty, error)
	ModifyMyBed(context.Context, *ModifyBedRequest) (*Empty, error)
	RemoveMyBed(context.Context, *BedId) (*Empty, error)
	GetMyBeds(context.Context, *Empty) (*BedList, error)
	AddBookingAvailability(context.Context, *BookingAvailability) (*Empty, error)
	//rpc ModifyBookingAvailability(BookingAvailability) returns (Empty); // NOT GOING TO BE IMPLEMENTED!!
	RemoveBookAvailability(context.Context, *BookingAvailability) (*Empty, error)
	mustEmbedUnimplementedAuthOnlyServiceServer()
}

// UnimplementedAuthOnlyServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAuthOnlyServiceServer struct {
}

func (UnimplementedAuthOnlyServiceServer) SignUp(context.Context, *AccountInfo) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignUp not implemented")
}
func (UnimplementedAuthOnlyServiceServer) Logout(context.Context, *Empty) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Logout not implemented")
}
func (UnimplementedAuthOnlyServiceServer) VerifyMail(context.Context, *VerifyMailRequest) (*VerifyMailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyMail not implemented")
}
func (UnimplementedAuthOnlyServiceServer) GetAccountInfo(context.Context, *Empty) (*AccountInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccountInfo not implemented")
}
func (UnimplementedAuthOnlyServiceServer) GetProfilePic(context.Context, *Empty) (*ProfilePic, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProfilePic not implemented")
}
func (UnimplementedAuthOnlyServiceServer) SetProfilePic(context.Context, *ProfilePic) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetProfilePic not implemented")
}
func (UnimplementedAuthOnlyServiceServer) DeleteAccount(context.Context, *Empty) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAccount not implemented")
}
func (UnimplementedAuthOnlyServiceServer) UpdateAccountInfo(context.Context, *AccountInfo) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAccountInfo not implemented")
}
func (UnimplementedAuthOnlyServiceServer) Book(context.Context, *Booking) (*BookResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Book not implemented")
}
func (UnimplementedAuthOnlyServiceServer) Review(context.Context, *ReviewRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Review not implemented")
}
func (UnimplementedAuthOnlyServiceServer) GetMyReview(context.Context, *BedId) (*GetMyReviewResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMyReview not implemented")
}
func (UnimplementedAuthOnlyServiceServer) RemoveReview(context.Context, *BedId) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveReview not implemented")
}
func (UnimplementedAuthOnlyServiceServer) AddBed(context.Context, *BedMutableInfo) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddBed not implemented")
}
func (UnimplementedAuthOnlyServiceServer) ModifyMyBed(context.Context, *ModifyBedRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ModifyMyBed not implemented")
}
func (UnimplementedAuthOnlyServiceServer) RemoveMyBed(context.Context, *BedId) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveMyBed not implemented")
}
func (UnimplementedAuthOnlyServiceServer) GetMyBeds(context.Context, *Empty) (*BedList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMyBeds not implemented")
}
func (UnimplementedAuthOnlyServiceServer) AddBookingAvailability(context.Context, *BookingAvailability) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddBookingAvailability not implemented")
}
func (UnimplementedAuthOnlyServiceServer) RemoveBookAvailability(context.Context, *BookingAvailability) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveBookAvailability not implemented")
}
func (UnimplementedAuthOnlyServiceServer) mustEmbedUnimplementedAuthOnlyServiceServer() {}

// UnsafeAuthOnlyServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthOnlyServiceServer will
// result in compilation errors.
type UnsafeAuthOnlyServiceServer interface {
	mustEmbedUnimplementedAuthOnlyServiceServer()
}

func RegisterAuthOnlyServiceServer(s grpc.ServiceRegistrar, srv AuthOnlyServiceServer) {
	s.RegisterService(&AuthOnlyService_ServiceDesc, srv)
}

func _AuthOnlyService_SignUp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthOnlyServiceServer).SignUp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/AuthOnlyService/SignUp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthOnlyServiceServer).SignUp(ctx, req.(*AccountInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthOnlyService_Logout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthOnlyServiceServer).Logout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/AuthOnlyService/logout",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthOnlyServiceServer).Logout(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthOnlyService_VerifyMail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyMailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthOnlyServiceServer).VerifyMail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/AuthOnlyService/VerifyMail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthOnlyServiceServer).VerifyMail(ctx, req.(*VerifyMailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthOnlyService_GetAccountInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthOnlyServiceServer).GetAccountInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/AuthOnlyService/GetAccountInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthOnlyServiceServer).GetAccountInfo(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthOnlyService_GetProfilePic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthOnlyServiceServer).GetProfilePic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/AuthOnlyService/GetProfilePic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthOnlyServiceServer).GetProfilePic(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthOnlyService_SetProfilePic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProfilePic)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthOnlyServiceServer).SetProfilePic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/AuthOnlyService/SetProfilePic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthOnlyServiceServer).SetProfilePic(ctx, req.(*ProfilePic))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthOnlyService_DeleteAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthOnlyServiceServer).DeleteAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/AuthOnlyService/DeleteAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthOnlyServiceServer).DeleteAccount(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthOnlyService_UpdateAccountInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthOnlyServiceServer).UpdateAccountInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/AuthOnlyService/UpdateAccountInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthOnlyServiceServer).UpdateAccountInfo(ctx, req.(*AccountInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthOnlyService_Book_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Booking)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthOnlyServiceServer).Book(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/AuthOnlyService/Book",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthOnlyServiceServer).Book(ctx, req.(*Booking))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthOnlyService_Review_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReviewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthOnlyServiceServer).Review(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/AuthOnlyService/Review",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthOnlyServiceServer).Review(ctx, req.(*ReviewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthOnlyService_GetMyReview_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BedId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthOnlyServiceServer).GetMyReview(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/AuthOnlyService/GetMyReview",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthOnlyServiceServer).GetMyReview(ctx, req.(*BedId))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthOnlyService_RemoveReview_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BedId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthOnlyServiceServer).RemoveReview(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/AuthOnlyService/RemoveReview",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthOnlyServiceServer).RemoveReview(ctx, req.(*BedId))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthOnlyService_AddBed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BedMutableInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthOnlyServiceServer).AddBed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/AuthOnlyService/AddBed",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthOnlyServiceServer).AddBed(ctx, req.(*BedMutableInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthOnlyService_ModifyMyBed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModifyBedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthOnlyServiceServer).ModifyMyBed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/AuthOnlyService/ModifyMyBed",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthOnlyServiceServer).ModifyMyBed(ctx, req.(*ModifyBedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthOnlyService_RemoveMyBed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BedId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthOnlyServiceServer).RemoveMyBed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/AuthOnlyService/RemoveMyBed",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthOnlyServiceServer).RemoveMyBed(ctx, req.(*BedId))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthOnlyService_GetMyBeds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthOnlyServiceServer).GetMyBeds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/AuthOnlyService/GetMyBeds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthOnlyServiceServer).GetMyBeds(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthOnlyService_AddBookingAvailability_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BookingAvailability)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthOnlyServiceServer).AddBookingAvailability(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/AuthOnlyService/AddBookingAvailability",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthOnlyServiceServer).AddBookingAvailability(ctx, req.(*BookingAvailability))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthOnlyService_RemoveBookAvailability_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BookingAvailability)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthOnlyServiceServer).RemoveBookAvailability(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/AuthOnlyService/RemoveBookAvailability",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthOnlyServiceServer).RemoveBookAvailability(ctx, req.(*BookingAvailability))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthOnlyService_ServiceDesc is the grpc.ServiceDesc for AuthOnlyService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthOnlyService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "AuthOnlyService",
	HandlerType: (*AuthOnlyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SignUp",
			Handler:    _AuthOnlyService_SignUp_Handler,
		},
		{
			MethodName: "logout",
			Handler:    _AuthOnlyService_Logout_Handler,
		},
		{
			MethodName: "VerifyMail",
			Handler:    _AuthOnlyService_VerifyMail_Handler,
		},
		{
			MethodName: "GetAccountInfo",
			Handler:    _AuthOnlyService_GetAccountInfo_Handler,
		},
		{
			MethodName: "GetProfilePic",
			Handler:    _AuthOnlyService_GetProfilePic_Handler,
		},
		{
			MethodName: "SetProfilePic",
			Handler:    _AuthOnlyService_SetProfilePic_Handler,
		},
		{
			MethodName: "DeleteAccount",
			Handler:    _AuthOnlyService_DeleteAccount_Handler,
		},
		{
			MethodName: "UpdateAccountInfo",
			Handler:    _AuthOnlyService_UpdateAccountInfo_Handler,
		},
		{
			MethodName: "Book",
			Handler:    _AuthOnlyService_Book_Handler,
		},
		{
			MethodName: "Review",
			Handler:    _AuthOnlyService_Review_Handler,
		},
		{
			MethodName: "GetMyReview",
			Handler:    _AuthOnlyService_GetMyReview_Handler,
		},
		{
			MethodName: "RemoveReview",
			Handler:    _AuthOnlyService_RemoveReview_Handler,
		},
		{
			MethodName: "AddBed",
			Handler:    _AuthOnlyService_AddBed_Handler,
		},
		{
			MethodName: "ModifyMyBed",
			Handler:    _AuthOnlyService_ModifyMyBed_Handler,
		},
		{
			MethodName: "RemoveMyBed",
			Handler:    _AuthOnlyService_RemoveMyBed_Handler,
		},
		{
			MethodName: "GetMyBeds",
			Handler:    _AuthOnlyService_GetMyBeds_Handler,
		},
		{
			MethodName: "AddBookingAvailability",
			Handler:    _AuthOnlyService_AddBookingAvailability_Handler,
		},
		{
			MethodName: "RemoveBookAvailability",
			Handler:    _AuthOnlyService_RemoveBookAvailability_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth-only-service.proto",
}
