// package: 
// file: auth-only-service.proto

/* tslint:disable */
/* eslint-disable */

import * as grpc from "grpc";
import * as auth_only_service_pb from "./auth-only-service_pb";
import * as common_messages_pb from "./common-messages_pb";

interface IAuthOnlyServiceService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
    signUp: IAuthOnlyServiceService_ISignUp;
    verifyMail: IAuthOnlyServiceService_IVerifyMail;
    getAccountInfo: IAuthOnlyServiceService_IGetAccountInfo;
    getProfilePic: IAuthOnlyServiceService_IGetProfilePic;
    deleteAccount: IAuthOnlyServiceService_IDeleteAccount;
    setProfilePic: IAuthOnlyServiceService_ISetProfilePic;
    updateAccountInfo: IAuthOnlyServiceService_IUpdateAccountInfo;
    book: IAuthOnlyServiceService_IBook;
    getMyBookings: IAuthOnlyServiceService_IGetMyBookings;
    review: IAuthOnlyServiceService_IReview;
    removeReview: IAuthOnlyServiceService_IRemoveReview;
    addBed: IAuthOnlyServiceService_IAddBed;
    modifyMyBed: IAuthOnlyServiceService_IModifyMyBed;
    removeMyBed: IAuthOnlyServiceService_IRemoveMyBed;
    getMyBeds: IAuthOnlyServiceService_IGetMyBeds;
    addBookingAvaiability: IAuthOnlyServiceService_IAddBookingAvaiability;
    removeBookAvaiability: IAuthOnlyServiceService_IRemoveBookAvaiability;
}

interface IAuthOnlyServiceService_ISignUp extends grpc.MethodDefinition<common_messages_pb.AccountInfo, common_messages_pb.Empty> {
    path: "/AuthOnlyService/SignUp";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<common_messages_pb.AccountInfo>;
    requestDeserialize: grpc.deserialize<common_messages_pb.AccountInfo>;
    responseSerialize: grpc.serialize<common_messages_pb.Empty>;
    responseDeserialize: grpc.deserialize<common_messages_pb.Empty>;
}
interface IAuthOnlyServiceService_IVerifyMail extends grpc.MethodDefinition<auth_only_service_pb.VerifyMailRequest, auth_only_service_pb.VerifyMailResponse> {
    path: "/AuthOnlyService/VerifyMail";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<auth_only_service_pb.VerifyMailRequest>;
    requestDeserialize: grpc.deserialize<auth_only_service_pb.VerifyMailRequest>;
    responseSerialize: grpc.serialize<auth_only_service_pb.VerifyMailResponse>;
    responseDeserialize: grpc.deserialize<auth_only_service_pb.VerifyMailResponse>;
}
interface IAuthOnlyServiceService_IGetAccountInfo extends grpc.MethodDefinition<common_messages_pb.Empty, common_messages_pb.AccountInfo> {
    path: "/AuthOnlyService/GetAccountInfo";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<common_messages_pb.Empty>;
    requestDeserialize: grpc.deserialize<common_messages_pb.Empty>;
    responseSerialize: grpc.serialize<common_messages_pb.AccountInfo>;
    responseDeserialize: grpc.deserialize<common_messages_pb.AccountInfo>;
}
interface IAuthOnlyServiceService_IGetProfilePic extends grpc.MethodDefinition<common_messages_pb.Empty, common_messages_pb.ProfilePic> {
    path: "/AuthOnlyService/GetProfilePic";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<common_messages_pb.Empty>;
    requestDeserialize: grpc.deserialize<common_messages_pb.Empty>;
    responseSerialize: grpc.serialize<common_messages_pb.ProfilePic>;
    responseDeserialize: grpc.deserialize<common_messages_pb.ProfilePic>;
}
interface IAuthOnlyServiceService_IDeleteAccount extends grpc.MethodDefinition<common_messages_pb.Empty, common_messages_pb.Empty> {
    path: "/AuthOnlyService/DeleteAccount";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<common_messages_pb.Empty>;
    requestDeserialize: grpc.deserialize<common_messages_pb.Empty>;
    responseSerialize: grpc.serialize<common_messages_pb.Empty>;
    responseDeserialize: grpc.deserialize<common_messages_pb.Empty>;
}
interface IAuthOnlyServiceService_ISetProfilePic extends grpc.MethodDefinition<common_messages_pb.ProfilePic, common_messages_pb.Empty> {
    path: "/AuthOnlyService/SetProfilePic";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<common_messages_pb.ProfilePic>;
    requestDeserialize: grpc.deserialize<common_messages_pb.ProfilePic>;
    responseSerialize: grpc.serialize<common_messages_pb.Empty>;
    responseDeserialize: grpc.deserialize<common_messages_pb.Empty>;
}
interface IAuthOnlyServiceService_IUpdateAccountInfo extends grpc.MethodDefinition<common_messages_pb.AccountInfo, common_messages_pb.Empty> {
    path: "/AuthOnlyService/UpdateAccountInfo";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<common_messages_pb.AccountInfo>;
    requestDeserialize: grpc.deserialize<common_messages_pb.AccountInfo>;
    responseSerialize: grpc.serialize<common_messages_pb.Empty>;
    responseDeserialize: grpc.deserialize<common_messages_pb.Empty>;
}
interface IAuthOnlyServiceService_IBook extends grpc.MethodDefinition<common_messages_pb.Booking, auth_only_service_pb.BookResponse> {
    path: "/AuthOnlyService/Book";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<common_messages_pb.Booking>;
    requestDeserialize: grpc.deserialize<common_messages_pb.Booking>;
    responseSerialize: grpc.serialize<auth_only_service_pb.BookResponse>;
    responseDeserialize: grpc.deserialize<auth_only_service_pb.BookResponse>;
}
interface IAuthOnlyServiceService_IGetMyBookings extends grpc.MethodDefinition<common_messages_pb.Empty, auth_only_service_pb.GetBookingsResponse> {
    path: "/AuthOnlyService/GetMyBookings";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<common_messages_pb.Empty>;
    requestDeserialize: grpc.deserialize<common_messages_pb.Empty>;
    responseSerialize: grpc.serialize<auth_only_service_pb.GetBookingsResponse>;
    responseDeserialize: grpc.deserialize<auth_only_service_pb.GetBookingsResponse>;
}
interface IAuthOnlyServiceService_IReview extends grpc.MethodDefinition<auth_only_service_pb.ReviewRequest, common_messages_pb.Empty> {
    path: "/AuthOnlyService/Review";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<auth_only_service_pb.ReviewRequest>;
    requestDeserialize: grpc.deserialize<auth_only_service_pb.ReviewRequest>;
    responseSerialize: grpc.serialize<common_messages_pb.Empty>;
    responseDeserialize: grpc.deserialize<common_messages_pb.Empty>;
}
interface IAuthOnlyServiceService_IRemoveReview extends grpc.MethodDefinition<common_messages_pb.BedId, common_messages_pb.Empty> {
    path: "/AuthOnlyService/RemoveReview";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<common_messages_pb.BedId>;
    requestDeserialize: grpc.deserialize<common_messages_pb.BedId>;
    responseSerialize: grpc.serialize<common_messages_pb.Empty>;
    responseDeserialize: grpc.deserialize<common_messages_pb.Empty>;
}
interface IAuthOnlyServiceService_IAddBed extends grpc.MethodDefinition<common_messages_pb.BedMutableInfo, common_messages_pb.Empty> {
    path: "/AuthOnlyService/AddBed";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<common_messages_pb.BedMutableInfo>;
    requestDeserialize: grpc.deserialize<common_messages_pb.BedMutableInfo>;
    responseSerialize: grpc.serialize<common_messages_pb.Empty>;
    responseDeserialize: grpc.deserialize<common_messages_pb.Empty>;
}
interface IAuthOnlyServiceService_IModifyMyBed extends grpc.MethodDefinition<auth_only_service_pb.ModifyBedRequest, common_messages_pb.Empty> {
    path: "/AuthOnlyService/ModifyMyBed";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<auth_only_service_pb.ModifyBedRequest>;
    requestDeserialize: grpc.deserialize<auth_only_service_pb.ModifyBedRequest>;
    responseSerialize: grpc.serialize<common_messages_pb.Empty>;
    responseDeserialize: grpc.deserialize<common_messages_pb.Empty>;
}
interface IAuthOnlyServiceService_IRemoveMyBed extends grpc.MethodDefinition<common_messages_pb.BedId, common_messages_pb.Empty> {
    path: "/AuthOnlyService/RemoveMyBed";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<common_messages_pb.BedId>;
    requestDeserialize: grpc.deserialize<common_messages_pb.BedId>;
    responseSerialize: grpc.serialize<common_messages_pb.Empty>;
    responseDeserialize: grpc.deserialize<common_messages_pb.Empty>;
}
interface IAuthOnlyServiceService_IGetMyBeds extends grpc.MethodDefinition<common_messages_pb.Empty, common_messages_pb.BedList> {
    path: "/AuthOnlyService/GetMyBeds";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<common_messages_pb.Empty>;
    requestDeserialize: grpc.deserialize<common_messages_pb.Empty>;
    responseSerialize: grpc.serialize<common_messages_pb.BedList>;
    responseDeserialize: grpc.deserialize<common_messages_pb.BedList>;
}
interface IAuthOnlyServiceService_IAddBookingAvaiability extends grpc.MethodDefinition<common_messages_pb.Booking, common_messages_pb.Empty> {
    path: "/AuthOnlyService/AddBookingAvaiability";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<common_messages_pb.Booking>;
    requestDeserialize: grpc.deserialize<common_messages_pb.Booking>;
    responseSerialize: grpc.serialize<common_messages_pb.Empty>;
    responseDeserialize: grpc.deserialize<common_messages_pb.Empty>;
}
interface IAuthOnlyServiceService_IRemoveBookAvaiability extends grpc.MethodDefinition<common_messages_pb.Booking, common_messages_pb.Empty> {
    path: "/AuthOnlyService/RemoveBookAvaiability";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<common_messages_pb.Booking>;
    requestDeserialize: grpc.deserialize<common_messages_pb.Booking>;
    responseSerialize: grpc.serialize<common_messages_pb.Empty>;
    responseDeserialize: grpc.deserialize<common_messages_pb.Empty>;
}

export const AuthOnlyServiceService: IAuthOnlyServiceService;

export interface IAuthOnlyServiceServer {
    signUp: grpc.handleUnaryCall<common_messages_pb.AccountInfo, common_messages_pb.Empty>;
    verifyMail: grpc.handleUnaryCall<auth_only_service_pb.VerifyMailRequest, auth_only_service_pb.VerifyMailResponse>;
    getAccountInfo: grpc.handleUnaryCall<common_messages_pb.Empty, common_messages_pb.AccountInfo>;
    getProfilePic: grpc.handleUnaryCall<common_messages_pb.Empty, common_messages_pb.ProfilePic>;
    deleteAccount: grpc.handleUnaryCall<common_messages_pb.Empty, common_messages_pb.Empty>;
    setProfilePic: grpc.handleUnaryCall<common_messages_pb.ProfilePic, common_messages_pb.Empty>;
    updateAccountInfo: grpc.handleUnaryCall<common_messages_pb.AccountInfo, common_messages_pb.Empty>;
    book: grpc.handleUnaryCall<common_messages_pb.Booking, auth_only_service_pb.BookResponse>;
    getMyBookings: grpc.handleUnaryCall<common_messages_pb.Empty, auth_only_service_pb.GetBookingsResponse>;
    review: grpc.handleUnaryCall<auth_only_service_pb.ReviewRequest, common_messages_pb.Empty>;
    removeReview: grpc.handleUnaryCall<common_messages_pb.BedId, common_messages_pb.Empty>;
    addBed: grpc.handleUnaryCall<common_messages_pb.BedMutableInfo, common_messages_pb.Empty>;
    modifyMyBed: grpc.handleUnaryCall<auth_only_service_pb.ModifyBedRequest, common_messages_pb.Empty>;
    removeMyBed: grpc.handleUnaryCall<common_messages_pb.BedId, common_messages_pb.Empty>;
    getMyBeds: grpc.handleUnaryCall<common_messages_pb.Empty, common_messages_pb.BedList>;
    addBookingAvaiability: grpc.handleUnaryCall<common_messages_pb.Booking, common_messages_pb.Empty>;
    removeBookAvaiability: grpc.handleUnaryCall<common_messages_pb.Booking, common_messages_pb.Empty>;
}

export interface IAuthOnlyServiceClient {
    signUp(request: common_messages_pb.AccountInfo, callback: (error: grpc.ServiceError | null, response: common_messages_pb.Empty) => void): grpc.ClientUnaryCall;
    signUp(request: common_messages_pb.AccountInfo, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: common_messages_pb.Empty) => void): grpc.ClientUnaryCall;
    signUp(request: common_messages_pb.AccountInfo, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: common_messages_pb.Empty) => void): grpc.ClientUnaryCall;
    verifyMail(request: auth_only_service_pb.VerifyMailRequest, callback: (error: grpc.ServiceError | null, response: auth_only_service_pb.VerifyMailResponse) => void): grpc.ClientUnaryCall;
    verifyMail(request: auth_only_service_pb.VerifyMailRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: auth_only_service_pb.VerifyMailResponse) => void): grpc.ClientUnaryCall;
    verifyMail(request: auth_only_service_pb.VerifyMailRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: auth_only_service_pb.VerifyMailResponse) => void): grpc.ClientUnaryCall;
    getAccountInfo(request: common_messages_pb.Empty, callback: (error: grpc.ServiceError | null, response: common_messages_pb.AccountInfo) => void): grpc.ClientUnaryCall;
    getAccountInfo(request: common_messages_pb.Empty, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: common_messages_pb.AccountInfo) => void): grpc.ClientUnaryCall;
    getAccountInfo(request: common_messages_pb.Empty, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: common_messages_pb.AccountInfo) => void): grpc.ClientUnaryCall;
    getProfilePic(request: common_messages_pb.Empty, callback: (error: grpc.ServiceError | null, response: common_messages_pb.ProfilePic) => void): grpc.ClientUnaryCall;
    getProfilePic(request: common_messages_pb.Empty, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: common_messages_pb.ProfilePic) => void): grpc.ClientUnaryCall;
    getProfilePic(request: common_messages_pb.Empty, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: common_messages_pb.ProfilePic) => void): grpc.ClientUnaryCall;
    deleteAccount(request: common_messages_pb.Empty, callback: (error: grpc.ServiceError | null, response: common_messages_pb.Empty) => void): grpc.ClientUnaryCall;
    deleteAccount(request: common_messages_pb.Empty, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: common_messages_pb.Empty) => void): grpc.ClientUnaryCall;
    deleteAccount(request: common_messages_pb.Empty, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: common_messages_pb.Empty) => void): grpc.ClientUnaryCall;
    setProfilePic(request: common_messages_pb.ProfilePic, callback: (error: grpc.ServiceError | null, response: common_messages_pb.Empty) => void): grpc.ClientUnaryCall;
    setProfilePic(request: common_messages_pb.ProfilePic, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: common_messages_pb.Empty) => void): grpc.ClientUnaryCall;
    setProfilePic(request: common_messages_pb.ProfilePic, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: common_messages_pb.Empty) => void): grpc.ClientUnaryCall;
    updateAccountInfo(request: common_messages_pb.AccountInfo, callback: (error: grpc.ServiceError | null, response: common_messages_pb.Empty) => void): grpc.ClientUnaryCall;
    updateAccountInfo(request: common_messages_pb.AccountInfo, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: common_messages_pb.Empty) => void): grpc.ClientUnaryCall;
    updateAccountInfo(request: common_messages_pb.AccountInfo, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: common_messages_pb.Empty) => void): grpc.ClientUnaryCall;
    book(request: common_messages_pb.Booking, callback: (error: grpc.ServiceError | null, response: auth_only_service_pb.BookResponse) => void): grpc.ClientUnaryCall;
    book(request: common_messages_pb.Booking, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: auth_only_service_pb.BookResponse) => void): grpc.ClientUnaryCall;
    book(request: common_messages_pb.Booking, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: auth_only_service_pb.BookResponse) => void): grpc.ClientUnaryCall;
    getMyBookings(request: common_messages_pb.Empty, callback: (error: grpc.ServiceError | null, response: auth_only_service_pb.GetBookingsResponse) => void): grpc.ClientUnaryCall;
    getMyBookings(request: common_messages_pb.Empty, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: auth_only_service_pb.GetBookingsResponse) => void): grpc.ClientUnaryCall;
    getMyBookings(request: common_messages_pb.Empty, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: auth_only_service_pb.GetBookingsResponse) => void): grpc.ClientUnaryCall;
    review(request: auth_only_service_pb.ReviewRequest, callback: (error: grpc.ServiceError | null, response: common_messages_pb.Empty) => void): grpc.ClientUnaryCall;
    review(request: auth_only_service_pb.ReviewRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: common_messages_pb.Empty) => void): grpc.ClientUnaryCall;
    review(request: auth_only_service_pb.ReviewRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: common_messages_pb.Empty) => void): grpc.ClientUnaryCall;
    removeReview(request: common_messages_pb.BedId, callback: (error: grpc.ServiceError | null, response: common_messages_pb.Empty) => void): grpc.ClientUnaryCall;
    removeReview(request: common_messages_pb.BedId, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: common_messages_pb.Empty) => void): grpc.ClientUnaryCall;
    removeReview(request: common_messages_pb.BedId, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: common_messages_pb.Empty) => void): grpc.ClientUnaryCall;
    addBed(request: common_messages_pb.BedMutableInfo, callback: (error: grpc.ServiceError | null, response: common_messages_pb.Empty) => void): grpc.ClientUnaryCall;
    addBed(request: common_messages_pb.BedMutableInfo, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: common_messages_pb.Empty) => void): grpc.ClientUnaryCall;
    addBed(request: common_messages_pb.BedMutableInfo, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: common_messages_pb.Empty) => void): grpc.ClientUnaryCall;
    modifyMyBed(request: auth_only_service_pb.ModifyBedRequest, callback: (error: grpc.ServiceError | null, response: common_messages_pb.Empty) => void): grpc.ClientUnaryCall;
    modifyMyBed(request: auth_only_service_pb.ModifyBedRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: common_messages_pb.Empty) => void): grpc.ClientUnaryCall;
    modifyMyBed(request: auth_only_service_pb.ModifyBedRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: common_messages_pb.Empty) => void): grpc.ClientUnaryCall;
    removeMyBed(request: common_messages_pb.BedId, callback: (error: grpc.ServiceError | null, response: common_messages_pb.Empty) => void): grpc.ClientUnaryCall;
    removeMyBed(request: common_messages_pb.BedId, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: common_messages_pb.Empty) => void): grpc.ClientUnaryCall;
    removeMyBed(request: common_messages_pb.BedId, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: common_messages_pb.Empty) => void): grpc.ClientUnaryCall;
    getMyBeds(request: common_messages_pb.Empty, callback: (error: grpc.ServiceError | null, response: common_messages_pb.BedList) => void): grpc.ClientUnaryCall;
    getMyBeds(request: common_messages_pb.Empty, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: common_messages_pb.BedList) => void): grpc.ClientUnaryCall;
    getMyBeds(request: common_messages_pb.Empty, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: common_messages_pb.BedList) => void): grpc.ClientUnaryCall;
    addBookingAvaiability(request: common_messages_pb.Booking, callback: (error: grpc.ServiceError | null, response: common_messages_pb.Empty) => void): grpc.ClientUnaryCall;
    addBookingAvaiability(request: common_messages_pb.Booking, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: common_messages_pb.Empty) => void): grpc.ClientUnaryCall;
    addBookingAvaiability(request: common_messages_pb.Booking, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: common_messages_pb.Empty) => void): grpc.ClientUnaryCall;
    removeBookAvaiability(request: common_messages_pb.Booking, callback: (error: grpc.ServiceError | null, response: common_messages_pb.Empty) => void): grpc.ClientUnaryCall;
    removeBookAvaiability(request: common_messages_pb.Booking, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: common_messages_pb.Empty) => void): grpc.ClientUnaryCall;
    removeBookAvaiability(request: common_messages_pb.Booking, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: common_messages_pb.Empty) => void): grpc.ClientUnaryCall;
}

export class AuthOnlyServiceClient extends grpc.Client implements IAuthOnlyServiceClient {
    constructor(address: string, credentials: grpc.ChannelCredentials, options?: object);
    public signUp(request: common_messages_pb.AccountInfo, callback: (error: grpc.ServiceError | null, response: common_messages_pb.Empty) => void): grpc.ClientUnaryCall;
    public signUp(request: common_messages_pb.AccountInfo, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: common_messages_pb.Empty) => void): grpc.ClientUnaryCall;
    public signUp(request: common_messages_pb.AccountInfo, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: common_messages_pb.Empty) => void): grpc.ClientUnaryCall;
    public verifyMail(request: auth_only_service_pb.VerifyMailRequest, callback: (error: grpc.ServiceError | null, response: auth_only_service_pb.VerifyMailResponse) => void): grpc.ClientUnaryCall;
    public verifyMail(request: auth_only_service_pb.VerifyMailRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: auth_only_service_pb.VerifyMailResponse) => void): grpc.ClientUnaryCall;
    public verifyMail(request: auth_only_service_pb.VerifyMailRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: auth_only_service_pb.VerifyMailResponse) => void): grpc.ClientUnaryCall;
    public getAccountInfo(request: common_messages_pb.Empty, callback: (error: grpc.ServiceError | null, response: common_messages_pb.AccountInfo) => void): grpc.ClientUnaryCall;
    public getAccountInfo(request: common_messages_pb.Empty, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: common_messages_pb.AccountInfo) => void): grpc.ClientUnaryCall;
    public getAccountInfo(request: common_messages_pb.Empty, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: common_messages_pb.AccountInfo) => void): grpc.ClientUnaryCall;
    public getProfilePic(request: common_messages_pb.Empty, callback: (error: grpc.ServiceError | null, response: common_messages_pb.ProfilePic) => void): grpc.ClientUnaryCall;
    public getProfilePic(request: common_messages_pb.Empty, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: common_messages_pb.ProfilePic) => void): grpc.ClientUnaryCall;
    public getProfilePic(request: common_messages_pb.Empty, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: common_messages_pb.ProfilePic) => void): grpc.ClientUnaryCall;
    public deleteAccount(request: common_messages_pb.Empty, callback: (error: grpc.ServiceError | null, response: common_messages_pb.Empty) => void): grpc.ClientUnaryCall;
    public deleteAccount(request: common_messages_pb.Empty, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: common_messages_pb.Empty) => void): grpc.ClientUnaryCall;
    public deleteAccount(request: common_messages_pb.Empty, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: common_messages_pb.Empty) => void): grpc.ClientUnaryCall;
    public setProfilePic(request: common_messages_pb.ProfilePic, callback: (error: grpc.ServiceError | null, response: common_messages_pb.Empty) => void): grpc.ClientUnaryCall;
    public setProfilePic(request: common_messages_pb.ProfilePic, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: common_messages_pb.Empty) => void): grpc.ClientUnaryCall;
    public setProfilePic(request: common_messages_pb.ProfilePic, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: common_messages_pb.Empty) => void): grpc.ClientUnaryCall;
    public updateAccountInfo(request: common_messages_pb.AccountInfo, callback: (error: grpc.ServiceError | null, response: common_messages_pb.Empty) => void): grpc.ClientUnaryCall;
    public updateAccountInfo(request: common_messages_pb.AccountInfo, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: common_messages_pb.Empty) => void): grpc.ClientUnaryCall;
    public updateAccountInfo(request: common_messages_pb.AccountInfo, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: common_messages_pb.Empty) => void): grpc.ClientUnaryCall;
    public book(request: common_messages_pb.Booking, callback: (error: grpc.ServiceError | null, response: auth_only_service_pb.BookResponse) => void): grpc.ClientUnaryCall;
    public book(request: common_messages_pb.Booking, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: auth_only_service_pb.BookResponse) => void): grpc.ClientUnaryCall;
    public book(request: common_messages_pb.Booking, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: auth_only_service_pb.BookResponse) => void): grpc.ClientUnaryCall;
    public getMyBookings(request: common_messages_pb.Empty, callback: (error: grpc.ServiceError | null, response: auth_only_service_pb.GetBookingsResponse) => void): grpc.ClientUnaryCall;
    public getMyBookings(request: common_messages_pb.Empty, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: auth_only_service_pb.GetBookingsResponse) => void): grpc.ClientUnaryCall;
    public getMyBookings(request: common_messages_pb.Empty, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: auth_only_service_pb.GetBookingsResponse) => void): grpc.ClientUnaryCall;
    public review(request: auth_only_service_pb.ReviewRequest, callback: (error: grpc.ServiceError | null, response: common_messages_pb.Empty) => void): grpc.ClientUnaryCall;
    public review(request: auth_only_service_pb.ReviewRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: common_messages_pb.Empty) => void): grpc.ClientUnaryCall;
    public review(request: auth_only_service_pb.ReviewRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: common_messages_pb.Empty) => void): grpc.ClientUnaryCall;
    public removeReview(request: common_messages_pb.BedId, callback: (error: grpc.ServiceError | null, response: common_messages_pb.Empty) => void): grpc.ClientUnaryCall;
    public removeReview(request: common_messages_pb.BedId, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: common_messages_pb.Empty) => void): grpc.ClientUnaryCall;
    public removeReview(request: common_messages_pb.BedId, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: common_messages_pb.Empty) => void): grpc.ClientUnaryCall;
    public addBed(request: common_messages_pb.BedMutableInfo, callback: (error: grpc.ServiceError | null, response: common_messages_pb.Empty) => void): grpc.ClientUnaryCall;
    public addBed(request: common_messages_pb.BedMutableInfo, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: common_messages_pb.Empty) => void): grpc.ClientUnaryCall;
    public addBed(request: common_messages_pb.BedMutableInfo, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: common_messages_pb.Empty) => void): grpc.ClientUnaryCall;
    public modifyMyBed(request: auth_only_service_pb.ModifyBedRequest, callback: (error: grpc.ServiceError | null, response: common_messages_pb.Empty) => void): grpc.ClientUnaryCall;
    public modifyMyBed(request: auth_only_service_pb.ModifyBedRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: common_messages_pb.Empty) => void): grpc.ClientUnaryCall;
    public modifyMyBed(request: auth_only_service_pb.ModifyBedRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: common_messages_pb.Empty) => void): grpc.ClientUnaryCall;
    public removeMyBed(request: common_messages_pb.BedId, callback: (error: grpc.ServiceError | null, response: common_messages_pb.Empty) => void): grpc.ClientUnaryCall;
    public removeMyBed(request: common_messages_pb.BedId, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: common_messages_pb.Empty) => void): grpc.ClientUnaryCall;
    public removeMyBed(request: common_messages_pb.BedId, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: common_messages_pb.Empty) => void): grpc.ClientUnaryCall;
    public getMyBeds(request: common_messages_pb.Empty, callback: (error: grpc.ServiceError | null, response: common_messages_pb.BedList) => void): grpc.ClientUnaryCall;
    public getMyBeds(request: common_messages_pb.Empty, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: common_messages_pb.BedList) => void): grpc.ClientUnaryCall;
    public getMyBeds(request: common_messages_pb.Empty, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: common_messages_pb.BedList) => void): grpc.ClientUnaryCall;
    public addBookingAvaiability(request: common_messages_pb.Booking, callback: (error: grpc.ServiceError | null, response: common_messages_pb.Empty) => void): grpc.ClientUnaryCall;
    public addBookingAvaiability(request: common_messages_pb.Booking, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: common_messages_pb.Empty) => void): grpc.ClientUnaryCall;
    public addBookingAvaiability(request: common_messages_pb.Booking, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: common_messages_pb.Empty) => void): grpc.ClientUnaryCall;
    public removeBookAvaiability(request: common_messages_pb.Booking, callback: (error: grpc.ServiceError | null, response: common_messages_pb.Empty) => void): grpc.ClientUnaryCall;
    public removeBookAvaiability(request: common_messages_pb.Booking, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: common_messages_pb.Empty) => void): grpc.ClientUnaryCall;
    public removeBookAvaiability(request: common_messages_pb.Booking, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: common_messages_pb.Empty) => void): grpc.ClientUnaryCall;
}
