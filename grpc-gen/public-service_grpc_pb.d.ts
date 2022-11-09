// package: 
// file: public-service.proto

/* tslint:disable */
/* eslint-disable */

import * as grpc from "grpc";
import * as public_service_pb from "./public-service_pb";
import * as common_messages_pb from "./common-messages_pb";

interface IPublicServiceService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
    getNonce: IPublicServiceService_IGetNonce;
    auth: IPublicServiceService_IAuth;
    getBeds: IPublicServiceService_IGetBeds;
    getReview: IPublicServiceService_IGetReview;
}

interface IPublicServiceService_IGetNonce extends grpc.MethodDefinition<common_messages_pb.Empty, public_service_pb.GetNonceResponse> {
    path: "/PublicService/GetNonce";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<common_messages_pb.Empty>;
    requestDeserialize: grpc.deserialize<common_messages_pb.Empty>;
    responseSerialize: grpc.serialize<public_service_pb.GetNonceResponse>;
    responseDeserialize: grpc.deserialize<public_service_pb.GetNonceResponse>;
}
interface IPublicServiceService_IAuth extends grpc.MethodDefinition<public_service_pb.AuthRequest, public_service_pb.AuthResponse> {
    path: "/PublicService/Auth";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<public_service_pb.AuthRequest>;
    requestDeserialize: grpc.deserialize<public_service_pb.AuthRequest>;
    responseSerialize: grpc.serialize<public_service_pb.AuthResponse>;
    responseDeserialize: grpc.deserialize<public_service_pb.AuthResponse>;
}
interface IPublicServiceService_IGetBeds extends grpc.MethodDefinition<public_service_pb.GetBedsRequest, common_messages_pb.BedList> {
    path: "/PublicService/GetBeds";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<public_service_pb.GetBedsRequest>;
    requestDeserialize: grpc.deserialize<public_service_pb.GetBedsRequest>;
    responseSerialize: grpc.serialize<common_messages_pb.BedList>;
    responseDeserialize: grpc.deserialize<common_messages_pb.BedList>;
}
interface IPublicServiceService_IGetReview extends grpc.MethodDefinition<public_service_pb.GetReviewsRequest, public_service_pb.GetReviewsResponse> {
    path: "/PublicService/GetReview";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<public_service_pb.GetReviewsRequest>;
    requestDeserialize: grpc.deserialize<public_service_pb.GetReviewsRequest>;
    responseSerialize: grpc.serialize<public_service_pb.GetReviewsResponse>;
    responseDeserialize: grpc.deserialize<public_service_pb.GetReviewsResponse>;
}

export const PublicServiceService: IPublicServiceService;

export interface IPublicServiceServer {
    getNonce: grpc.handleUnaryCall<common_messages_pb.Empty, public_service_pb.GetNonceResponse>;
    auth: grpc.handleUnaryCall<public_service_pb.AuthRequest, public_service_pb.AuthResponse>;
    getBeds: grpc.handleUnaryCall<public_service_pb.GetBedsRequest, common_messages_pb.BedList>;
    getReview: grpc.handleUnaryCall<public_service_pb.GetReviewsRequest, public_service_pb.GetReviewsResponse>;
}

export interface IPublicServiceClient {
    getNonce(request: common_messages_pb.Empty, callback: (error: grpc.ServiceError | null, response: public_service_pb.GetNonceResponse) => void): grpc.ClientUnaryCall;
    getNonce(request: common_messages_pb.Empty, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: public_service_pb.GetNonceResponse) => void): grpc.ClientUnaryCall;
    getNonce(request: common_messages_pb.Empty, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: public_service_pb.GetNonceResponse) => void): grpc.ClientUnaryCall;
    auth(request: public_service_pb.AuthRequest, callback: (error: grpc.ServiceError | null, response: public_service_pb.AuthResponse) => void): grpc.ClientUnaryCall;
    auth(request: public_service_pb.AuthRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: public_service_pb.AuthResponse) => void): grpc.ClientUnaryCall;
    auth(request: public_service_pb.AuthRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: public_service_pb.AuthResponse) => void): grpc.ClientUnaryCall;
    getBeds(request: public_service_pb.GetBedsRequest, callback: (error: grpc.ServiceError | null, response: common_messages_pb.BedList) => void): grpc.ClientUnaryCall;
    getBeds(request: public_service_pb.GetBedsRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: common_messages_pb.BedList) => void): grpc.ClientUnaryCall;
    getBeds(request: public_service_pb.GetBedsRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: common_messages_pb.BedList) => void): grpc.ClientUnaryCall;
    getReview(request: public_service_pb.GetReviewsRequest, callback: (error: grpc.ServiceError | null, response: public_service_pb.GetReviewsResponse) => void): grpc.ClientUnaryCall;
    getReview(request: public_service_pb.GetReviewsRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: public_service_pb.GetReviewsResponse) => void): grpc.ClientUnaryCall;
    getReview(request: public_service_pb.GetReviewsRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: public_service_pb.GetReviewsResponse) => void): grpc.ClientUnaryCall;
}

export class PublicServiceClient extends grpc.Client implements IPublicServiceClient {
    constructor(address: string, credentials: grpc.ChannelCredentials, options?: object);
    public getNonce(request: common_messages_pb.Empty, callback: (error: grpc.ServiceError | null, response: public_service_pb.GetNonceResponse) => void): grpc.ClientUnaryCall;
    public getNonce(request: common_messages_pb.Empty, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: public_service_pb.GetNonceResponse) => void): grpc.ClientUnaryCall;
    public getNonce(request: common_messages_pb.Empty, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: public_service_pb.GetNonceResponse) => void): grpc.ClientUnaryCall;
    public auth(request: public_service_pb.AuthRequest, callback: (error: grpc.ServiceError | null, response: public_service_pb.AuthResponse) => void): grpc.ClientUnaryCall;
    public auth(request: public_service_pb.AuthRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: public_service_pb.AuthResponse) => void): grpc.ClientUnaryCall;
    public auth(request: public_service_pb.AuthRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: public_service_pb.AuthResponse) => void): grpc.ClientUnaryCall;
    public getBeds(request: public_service_pb.GetBedsRequest, callback: (error: grpc.ServiceError | null, response: common_messages_pb.BedList) => void): grpc.ClientUnaryCall;
    public getBeds(request: public_service_pb.GetBedsRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: common_messages_pb.BedList) => void): grpc.ClientUnaryCall;
    public getBeds(request: public_service_pb.GetBedsRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: common_messages_pb.BedList) => void): grpc.ClientUnaryCall;
    public getReview(request: public_service_pb.GetReviewsRequest, callback: (error: grpc.ServiceError | null, response: public_service_pb.GetReviewsResponse) => void): grpc.ClientUnaryCall;
    public getReview(request: public_service_pb.GetReviewsRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: public_service_pb.GetReviewsResponse) => void): grpc.ClientUnaryCall;
    public getReview(request: public_service_pb.GetReviewsRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: public_service_pb.GetReviewsResponse) => void): grpc.ClientUnaryCall;
}
