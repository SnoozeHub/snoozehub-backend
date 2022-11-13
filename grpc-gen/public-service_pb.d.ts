// package: 
// file: public-service.proto

/* tslint:disable */
/* eslint-disable */

import * as jspb from "google-protobuf";
import * as common_messages_pb from "./common-messages_pb";

export class GetNonceResponse extends jspb.Message { 
    getNonce(): string;
    setNonce(value: string): GetNonceResponse;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): GetNonceResponse.AsObject;
    static toObject(includeInstance: boolean, msg: GetNonceResponse): GetNonceResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: GetNonceResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): GetNonceResponse;
    static deserializeBinaryFromReader(message: GetNonceResponse, reader: jspb.BinaryReader): GetNonceResponse;
}

export namespace GetNonceResponse {
    export type AsObject = {
        nonce: string,
    }
}

export class AuthRequest extends jspb.Message { 
    getAddress(): string;
    setAddress(value: string): AuthRequest;
    getNoncesigned(): string;
    setNoncesigned(value: string): AuthRequest;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): AuthRequest.AsObject;
    static toObject(includeInstance: boolean, msg: AuthRequest): AuthRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: AuthRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): AuthRequest;
    static deserializeBinaryFromReader(message: AuthRequest, reader: jspb.BinaryReader): AuthRequest;
}

export namespace AuthRequest {
    export type AsObject = {
        address: string,
        noncesigned: string,
    }
}

export class AuthResponse extends jspb.Message { 
    getAuthtoken(): string;
    setAuthtoken(value: string): AuthResponse;
    getAccountdoesnotexist(): boolean;
    setAccountdoesnotexist(value: boolean): AuthResponse;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): AuthResponse.AsObject;
    static toObject(includeInstance: boolean, msg: AuthResponse): AuthResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: AuthResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): AuthResponse;
    static deserializeBinaryFromReader(message: AuthResponse, reader: jspb.BinaryReader): AuthResponse;
}

export namespace AuthResponse {
    export type AsObject = {
        authtoken: string,
        accountdoesnotexist: boolean,
    }
}

export class GetBedsRequest extends jspb.Message { 
    getDaterangelow(): number;
    setDaterangelow(value: number): GetBedsRequest;
    getDaterangehigh(): number;
    setDaterangehigh(value: number): GetBedsRequest;
    getPlace(): string;
    setPlace(value: string): GetBedsRequest;

    hasFeaturesmondadory(): boolean;
    clearFeaturesmondadory(): void;
    getFeaturesmondadory(): common_messages_pb.Features | undefined;
    setFeaturesmondadory(value?: common_messages_pb.Features): GetBedsRequest;
    getFromindex(): number;
    setFromindex(value: number): GetBedsRequest;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): GetBedsRequest.AsObject;
    static toObject(includeInstance: boolean, msg: GetBedsRequest): GetBedsRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: GetBedsRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): GetBedsRequest;
    static deserializeBinaryFromReader(message: GetBedsRequest, reader: jspb.BinaryReader): GetBedsRequest;
}

export namespace GetBedsRequest {
    export type AsObject = {
        daterangelow: number,
        daterangehigh: number,
        place: string,
        featuresmondadory?: common_messages_pb.Features.AsObject,
        fromindex: number,
    }
}

export class GetReviewsRequest extends jspb.Message { 

    hasBedid(): boolean;
    clearBedid(): void;
    getBedid(): common_messages_pb.BedId | undefined;
    setBedid(value?: common_messages_pb.BedId): GetReviewsRequest;
    getFromindex(): number;
    setFromindex(value: number): GetReviewsRequest;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): GetReviewsRequest.AsObject;
    static toObject(includeInstance: boolean, msg: GetReviewsRequest): GetReviewsRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: GetReviewsRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): GetReviewsRequest;
    static deserializeBinaryFromReader(message: GetReviewsRequest, reader: jspb.BinaryReader): GetReviewsRequest;
}

export namespace GetReviewsRequest {
    export type AsObject = {
        bedid?: common_messages_pb.BedId.AsObject,
        fromindex: number,
    }
}

export class GetReviewsResponse extends jspb.Message { 
    clearReviewsList(): void;
    getReviewsList(): Array<common_messages_pb.Review>;
    setReviewsList(value: Array<common_messages_pb.Review>): GetReviewsResponse;
    addReviews(value?: common_messages_pb.Review, index?: number): common_messages_pb.Review;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): GetReviewsResponse.AsObject;
    static toObject(includeInstance: boolean, msg: GetReviewsResponse): GetReviewsResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: GetReviewsResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): GetReviewsResponse;
    static deserializeBinaryFromReader(message: GetReviewsResponse, reader: jspb.BinaryReader): GetReviewsResponse;
}

export namespace GetReviewsResponse {
    export type AsObject = {
        reviewsList: Array<common_messages_pb.Review.AsObject>,
    }
}
