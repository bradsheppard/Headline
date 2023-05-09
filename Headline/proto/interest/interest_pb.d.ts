import * as jspb from 'google-protobuf'

import * as google_protobuf_empty_pb from 'google-protobuf/google/protobuf/empty_pb';


export class CreateInterest extends jspb.Message {
  getName(): string;
  setName(value: string): CreateInterest;

  getUserid(): number;
  setUserid(value: number): CreateInterest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateInterest.AsObject;
  static toObject(includeInstance: boolean, msg: CreateInterest): CreateInterest.AsObject;
  static serializeBinaryToWriter(message: CreateInterest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateInterest;
  static deserializeBinaryFromReader(message: CreateInterest, reader: jspb.BinaryReader): CreateInterest;
}

export namespace CreateInterest {
  export type AsObject = {
    name: string,
    userid: number,
  }
}

export class Interest extends jspb.Message {
  getId(): number;
  setId(value: number): Interest;

  getName(): string;
  setName(value: string): Interest;

  getUserid(): number;
  setUserid(value: number): Interest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Interest.AsObject;
  static toObject(includeInstance: boolean, msg: Interest): Interest.AsObject;
  static serializeBinaryToWriter(message: Interest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Interest;
  static deserializeBinaryFromReader(message: Interest, reader: jspb.BinaryReader): Interest;
}

export namespace Interest {
  export type AsObject = {
    id: number,
    name: string,
    userid: number,
  }
}

export class InterestResponse extends jspb.Message {
  getInterestsList(): Array<Interest>;
  setInterestsList(value: Array<Interest>): InterestResponse;
  clearInterestsList(): InterestResponse;
  addInterests(value?: Interest, index?: number): Interest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): InterestResponse.AsObject;
  static toObject(includeInstance: boolean, msg: InterestResponse): InterestResponse.AsObject;
  static serializeBinaryToWriter(message: InterestResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): InterestResponse;
  static deserializeBinaryFromReader(message: InterestResponse, reader: jspb.BinaryReader): InterestResponse;
}

export namespace InterestResponse {
  export type AsObject = {
    interestsList: Array<Interest.AsObject>,
  }
}

export class GetInterestsRequest extends jspb.Message {
  getUserid(): number;
  setUserid(value: number): GetInterestsRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetInterestsRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetInterestsRequest): GetInterestsRequest.AsObject;
  static serializeBinaryToWriter(message: GetInterestsRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetInterestsRequest;
  static deserializeBinaryFromReader(message: GetInterestsRequest, reader: jspb.BinaryReader): GetInterestsRequest;
}

export namespace GetInterestsRequest {
  export type AsObject = {
    userid: number,
  }
}

export class AddInterestsRequest extends jspb.Message {
  getInterestsList(): Array<CreateInterest>;
  setInterestsList(value: Array<CreateInterest>): AddInterestsRequest;
  clearInterestsList(): AddInterestsRequest;
  addInterests(value?: CreateInterest, index?: number): CreateInterest;

  getUserid(): number;
  setUserid(value: number): AddInterestsRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): AddInterestsRequest.AsObject;
  static toObject(includeInstance: boolean, msg: AddInterestsRequest): AddInterestsRequest.AsObject;
  static serializeBinaryToWriter(message: AddInterestsRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): AddInterestsRequest;
  static deserializeBinaryFromReader(message: AddInterestsRequest, reader: jspb.BinaryReader): AddInterestsRequest;
}

export namespace AddInterestsRequest {
  export type AsObject = {
    interestsList: Array<CreateInterest.AsObject>,
    userid: number,
  }
}

export class DeleteInterestsRequest extends jspb.Message {
  getIdsList(): Array<number>;
  setIdsList(value: Array<number>): DeleteInterestsRequest;
  clearIdsList(): DeleteInterestsRequest;
  addIds(value: number, index?: number): DeleteInterestsRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): DeleteInterestsRequest.AsObject;
  static toObject(includeInstance: boolean, msg: DeleteInterestsRequest): DeleteInterestsRequest.AsObject;
  static serializeBinaryToWriter(message: DeleteInterestsRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): DeleteInterestsRequest;
  static deserializeBinaryFromReader(message: DeleteInterestsRequest, reader: jspb.BinaryReader): DeleteInterestsRequest;
}

export namespace DeleteInterestsRequest {
  export type AsObject = {
    idsList: Array<number>,
  }
}

