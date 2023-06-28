import * as jspb from 'google-protobuf'

import * as google_protobuf_empty_pb from 'google-protobuf/google/protobuf/empty_pb';
import * as google_protobuf_timestamp_pb from 'google-protobuf/google/protobuf/timestamp_pb';


export class Topic extends jspb.Message {
  getName(): string;
  setName(value: string): Topic;

  getLastupdated(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setLastupdated(value?: google_protobuf_timestamp_pb.Timestamp): Topic;
  hasLastupdated(): boolean;
  clearLastupdated(): Topic;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Topic.AsObject;
  static toObject(includeInstance: boolean, msg: Topic): Topic.AsObject;
  static serializeBinaryToWriter(message: Topic, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Topic;
  static deserializeBinaryFromReader(message: Topic, reader: jspb.BinaryReader): Topic;
}

export namespace Topic {
  export type AsObject = {
    name: string,
    lastupdated?: google_protobuf_timestamp_pb.Timestamp.AsObject,
  }
}

export class TopicResponse extends jspb.Message {
  getTopicsList(): Array<Topic>;
  setTopicsList(value: Array<Topic>): TopicResponse;
  clearTopicsList(): TopicResponse;
  addTopics(value?: Topic, index?: number): Topic;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): TopicResponse.AsObject;
  static toObject(includeInstance: boolean, msg: TopicResponse): TopicResponse.AsObject;
  static serializeBinaryToWriter(message: TopicResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): TopicResponse;
  static deserializeBinaryFromReader(message: TopicResponse, reader: jspb.BinaryReader): TopicResponse;
}

export namespace TopicResponse {
  export type AsObject = {
    topicsList: Array<Topic.AsObject>,
  }
}

export class GetTopicsRequest extends jspb.Message {
  getUserid(): number;
  setUserid(value: number): GetTopicsRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetTopicsRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetTopicsRequest): GetTopicsRequest.AsObject;
  static serializeBinaryToWriter(message: GetTopicsRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetTopicsRequest;
  static deserializeBinaryFromReader(message: GetTopicsRequest, reader: jspb.BinaryReader): GetTopicsRequest;
}

export namespace GetTopicsRequest {
  export type AsObject = {
    userid: number,
  }
}

export class GetPendingTopicsRequest extends jspb.Message {
  getLastupdated(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setLastupdated(value?: google_protobuf_timestamp_pb.Timestamp): GetPendingTopicsRequest;
  hasLastupdated(): boolean;
  clearLastupdated(): GetPendingTopicsRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetPendingTopicsRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetPendingTopicsRequest): GetPendingTopicsRequest.AsObject;
  static serializeBinaryToWriter(message: GetPendingTopicsRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetPendingTopicsRequest;
  static deserializeBinaryFromReader(message: GetPendingTopicsRequest, reader: jspb.BinaryReader): GetPendingTopicsRequest;
}

export namespace GetPendingTopicsRequest {
  export type AsObject = {
    lastupdated?: google_protobuf_timestamp_pb.Timestamp.AsObject,
  }
}

export class AddTopicsRequest extends jspb.Message {
  getTopicsList(): Array<Topic>;
  setTopicsList(value: Array<Topic>): AddTopicsRequest;
  clearTopicsList(): AddTopicsRequest;
  addTopics(value?: Topic, index?: number): Topic;

  getUserid(): number;
  setUserid(value: number): AddTopicsRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): AddTopicsRequest.AsObject;
  static toObject(includeInstance: boolean, msg: AddTopicsRequest): AddTopicsRequest.AsObject;
  static serializeBinaryToWriter(message: AddTopicsRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): AddTopicsRequest;
  static deserializeBinaryFromReader(message: AddTopicsRequest, reader: jspb.BinaryReader): AddTopicsRequest;
}

export namespace AddTopicsRequest {
  export type AsObject = {
    topicsList: Array<Topic.AsObject>,
    userid: number,
  }
}

export class RemoveTopicsRequest extends jspb.Message {
  getTopicnamesList(): Array<string>;
  setTopicnamesList(value: Array<string>): RemoveTopicsRequest;
  clearTopicnamesList(): RemoveTopicsRequest;
  addTopicnames(value: string, index?: number): RemoveTopicsRequest;

  getUserid(): number;
  setUserid(value: number): RemoveTopicsRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): RemoveTopicsRequest.AsObject;
  static toObject(includeInstance: boolean, msg: RemoveTopicsRequest): RemoveTopicsRequest.AsObject;
  static serializeBinaryToWriter(message: RemoveTopicsRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): RemoveTopicsRequest;
  static deserializeBinaryFromReader(message: RemoveTopicsRequest, reader: jspb.BinaryReader): RemoveTopicsRequest;
}

export namespace RemoveTopicsRequest {
  export type AsObject = {
    topicnamesList: Array<string>,
    userid: number,
  }
}

