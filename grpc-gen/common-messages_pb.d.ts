// package: 
// file: common-messages.proto

/* tslint:disable */
/* eslint-disable */

import * as jspb from "google-protobuf";

export class Empty extends jspb.Message { 

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): Empty.AsObject;
    static toObject(includeInstance: boolean, msg: Empty): Empty.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: Empty, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): Empty;
    static deserializeBinaryFromReader(message: Empty, reader: jspb.BinaryReader): Empty;
}

export namespace Empty {
    export type AsObject = {
    }
}

export class AccountInfo extends jspb.Message { 
    getName(): string;
    setName(value: string): AccountInfo;
    getSurname(): string;
    setSurname(value: string): AccountInfo;
    getMail(): string;
    setMail(value: string): AccountInfo;
    getTelegramusername(): string;
    setTelegramusername(value: string): AccountInfo;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): AccountInfo.AsObject;
    static toObject(includeInstance: boolean, msg: AccountInfo): AccountInfo.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: AccountInfo, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): AccountInfo;
    static deserializeBinaryFromReader(message: AccountInfo, reader: jspb.BinaryReader): AccountInfo;
}

export namespace AccountInfo {
    export type AsObject = {
        name: string,
        surname: string,
        mail: string,
        telegramusername: string,
    }
}

export class ProfilePic extends jspb.Message { 

    hasImage(): boolean;
    clearImage(): void;
    getImage(): Uint8Array | string;
    getImage_asU8(): Uint8Array;
    getImage_asB64(): string;
    setImage(value: Uint8Array | string): ProfilePic;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ProfilePic.AsObject;
    static toObject(includeInstance: boolean, msg: ProfilePic): ProfilePic.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: ProfilePic, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ProfilePic;
    static deserializeBinaryFromReader(message: ProfilePic, reader: jspb.BinaryReader): ProfilePic;
}

export namespace ProfilePic {
    export type AsObject = {
        image: Uint8Array | string,
    }
}

export class Features extends jspb.Message { 
    getInternetconnection(): boolean;
    setInternetconnection(value: boolean): Features;
    getBathroom(): boolean;
    setBathroom(value: boolean): Features;
    getHeating(): boolean;
    setHeating(value: boolean): Features;
    getAirconditioner(): boolean;
    setAirconditioner(value: boolean): Features;
    getElectricaloutlet(): boolean;
    setElectricaloutlet(value: boolean): Features;
    getTap(): boolean;
    setTap(value: boolean): Features;
    getBedlinens(): boolean;
    setBedlinens(value: boolean): Features;
    getPillows(): boolean;
    setPillows(value: boolean): Features;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): Features.AsObject;
    static toObject(includeInstance: boolean, msg: Features): Features.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: Features, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): Features;
    static deserializeBinaryFromReader(message: Features, reader: jspb.BinaryReader): Features;
}

export namespace Features {
    export type AsObject = {
        internetconnection: boolean,
        bathroom: boolean,
        heating: boolean,
        airconditioner: boolean,
        electricaloutlet: boolean,
        tap: boolean,
        bedlinens: boolean,
        pillows: boolean,
    }
}

export class Booking extends jspb.Message { 

    hasBedid(): boolean;
    clearBedid(): void;
    getBedid(): BedId | undefined;
    setBedid(value?: BedId): Booking;
    getDate(): number;
    setDate(value: number): Booking;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): Booking.AsObject;
    static toObject(includeInstance: boolean, msg: Booking): Booking.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: Booking, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): Booking;
    static deserializeBinaryFromReader(message: Booking, reader: jspb.BinaryReader): Booking;
}

export namespace Booking {
    export type AsObject = {
        bedid?: BedId.AsObject,
        date: number,
    }
}

export class BedMutableInfo extends jspb.Message { 
    getPlace(): string;
    setPlace(value: string): BedMutableInfo;
    clearImagesList(): void;
    getImagesList(): Array<Uint8Array | string>;
    getImagesList_asU8(): Array<Uint8Array>;
    getImagesList_asB64(): Array<string>;
    setImagesList(value: Array<Uint8Array | string>): BedMutableInfo;
    addImages(value: Uint8Array | string, index?: number): Uint8Array | string;

    hasDescription(): boolean;
    clearDescription(): void;
    getDescription(): string | undefined;
    setDescription(value: string): BedMutableInfo;

    hasFeatures(): boolean;
    clearFeatures(): void;
    getFeatures(): Features | undefined;
    setFeatures(value?: Features): BedMutableInfo;
    getMinimumdaysnotice(): number;
    setMinimumdaysnotice(value: number): BedMutableInfo;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): BedMutableInfo.AsObject;
    static toObject(includeInstance: boolean, msg: BedMutableInfo): BedMutableInfo.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: BedMutableInfo, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): BedMutableInfo;
    static deserializeBinaryFromReader(message: BedMutableInfo, reader: jspb.BinaryReader): BedMutableInfo;
}

export namespace BedMutableInfo {
    export type AsObject = {
        place: string,
        imagesList: Array<Uint8Array | string>,
        description?: string,
        features?: Features.AsObject,
        minimumdaysnotice: number,
    }
}

export class BedList extends jspb.Message { 
    clearBedsList(): void;
    getBedsList(): Array<BedList.Bed>;
    setBedsList(value: Array<BedList.Bed>): BedList;
    addBeds(value?: BedList.Bed, index?: number): BedList.Bed;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): BedList.AsObject;
    static toObject(includeInstance: boolean, msg: BedList): BedList.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: BedList, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): BedList;
    static deserializeBinaryFromReader(message: BedList, reader: jspb.BinaryReader): BedList;
}

export namespace BedList {
    export type AsObject = {
        bedsList: Array<BedList.Bed.AsObject>,
    }


    export class Bed extends jspb.Message { 

        hasId(): boolean;
        clearId(): void;
        getId(): BedId | undefined;
        setId(value?: BedId): Bed;

        hasBedmutableinfo(): boolean;
        clearBedmutableinfo(): void;
        getBedmutableinfo(): BedMutableInfo | undefined;
        setBedmutableinfo(value?: BedMutableInfo): Bed;
        clearDateavailablesList(): void;
        getDateavailablesList(): Array<number>;
        setDateavailablesList(value: Array<number>): Bed;
        addDateavailables(value: number, index?: number): number;
        getReviewcount(): number;
        setReviewcount(value: number): Bed;

        hasAverageevaluation(): boolean;
        clearAverageevaluation(): void;
        getAverageevaluation(): number | undefined;
        setAverageevaluation(value: number): Bed;

        serializeBinary(): Uint8Array;
        toObject(includeInstance?: boolean): Bed.AsObject;
        static toObject(includeInstance: boolean, msg: Bed): Bed.AsObject;
        static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
        static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
        static serializeBinaryToWriter(message: Bed, writer: jspb.BinaryWriter): void;
        static deserializeBinary(bytes: Uint8Array): Bed;
        static deserializeBinaryFromReader(message: Bed, reader: jspb.BinaryReader): Bed;
    }

    export namespace Bed {
        export type AsObject = {
            id?: BedId.AsObject,
            bedmutableinfo?: BedMutableInfo.AsObject,
            dateavailablesList: Array<number>,
            reviewcount: number,
            averageevaluation?: number,
        }
    }

}

export class BedId extends jspb.Message { 
    getBedid(): string;
    setBedid(value: string): BedId;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): BedId.AsObject;
    static toObject(includeInstance: boolean, msg: BedId): BedId.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: BedId, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): BedId;
    static deserializeBinaryFromReader(message: BedId, reader: jspb.BinaryReader): BedId;
}

export namespace BedId {
    export type AsObject = {
        bedid: string,
    }
}

export class Review extends jspb.Message { 
    getEvaluation(): number;
    setEvaluation(value: number): Review;

    hasComment(): boolean;
    clearComment(): void;
    getComment(): string | undefined;
    setComment(value: string): Review;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): Review.AsObject;
    static toObject(includeInstance: boolean, msg: Review): Review.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: Review, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): Review;
    static deserializeBinaryFromReader(message: Review, reader: jspb.BinaryReader): Review;
}

export namespace Review {
    export type AsObject = {
        evaluation: number,
        comment?: string,
    }
}
