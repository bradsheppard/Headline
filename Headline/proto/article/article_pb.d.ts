import * as jspb from 'google-protobuf'

import * as google_protobuf_empty_pb from 'google-protobuf/google/protobuf/empty_pb';
import * as google_protobuf_timestamp_pb from 'google-protobuf/google/protobuf/timestamp_pb';


export class TopicArticles extends jspb.Message {
  getTopicarticlesMap(): jspb.Map<string, Articles>;
  clearTopicarticlesMap(): TopicArticles;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): TopicArticles.AsObject;
  static toObject(includeInstance: boolean, msg: TopicArticles): TopicArticles.AsObject;
  static serializeBinaryToWriter(message: TopicArticles, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): TopicArticles;
  static deserializeBinaryFromReader(message: TopicArticles, reader: jspb.BinaryReader): TopicArticles;
}

export namespace TopicArticles {
  export type AsObject = {
    topicarticlesMap: Array<[string, Articles.AsObject]>,
  }
}

export class Articles extends jspb.Message {
  getArticlesList(): Array<Article>;
  setArticlesList(value: Array<Article>): Articles;
  clearArticlesList(): Articles;
  addArticles(value?: Article, index?: number): Article;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Articles.AsObject;
  static toObject(includeInstance: boolean, msg: Articles): Articles.AsObject;
  static serializeBinaryToWriter(message: Articles, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Articles;
  static deserializeBinaryFromReader(message: Articles, reader: jspb.BinaryReader): Articles;
}

export namespace Articles {
  export type AsObject = {
    articlesList: Array<Article.AsObject>,
  }
}

export class Article extends jspb.Message {
  getTitle(): string;
  setTitle(value: string): Article;

  getDescription(): string;
  setDescription(value: string): Article;

  getUrl(): string;
  setUrl(value: string): Article;

  getImageurl(): string;
  setImageurl(value: string): Article;

  getSource(): string;
  setSource(value: string): Article;

  getDate(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setDate(value?: google_protobuf_timestamp_pb.Timestamp): Article;
  hasDate(): boolean;
  clearDate(): Article;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Article.AsObject;
  static toObject(includeInstance: boolean, msg: Article): Article.AsObject;
  static serializeBinaryToWriter(message: Article, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Article;
  static deserializeBinaryFromReader(message: Article, reader: jspb.BinaryReader): Article;
}

export namespace Article {
  export type AsObject = {
    title: string,
    description: string,
    url: string,
    imageurl: string,
    source: string,
    date?: google_protobuf_timestamp_pb.Timestamp.AsObject,
  }
}

export class SetTopicArticlesRequest extends jspb.Message {
  getTopicarticlesMap(): jspb.Map<string, Articles>;
  clearTopicarticlesMap(): SetTopicArticlesRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): SetTopicArticlesRequest.AsObject;
  static toObject(includeInstance: boolean, msg: SetTopicArticlesRequest): SetTopicArticlesRequest.AsObject;
  static serializeBinaryToWriter(message: SetTopicArticlesRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): SetTopicArticlesRequest;
  static deserializeBinaryFromReader(message: SetTopicArticlesRequest, reader: jspb.BinaryReader): SetTopicArticlesRequest;
}

export namespace SetTopicArticlesRequest {
  export type AsObject = {
    topicarticlesMap: Array<[string, Articles.AsObject]>,
  }
}

export class GetTopicArticlesRequest extends jspb.Message {
  getTopicsList(): Array<string>;
  setTopicsList(value: Array<string>): GetTopicArticlesRequest;
  clearTopicsList(): GetTopicArticlesRequest;
  addTopics(value: string, index?: number): GetTopicArticlesRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetTopicArticlesRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetTopicArticlesRequest): GetTopicArticlesRequest.AsObject;
  static serializeBinaryToWriter(message: GetTopicArticlesRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetTopicArticlesRequest;
  static deserializeBinaryFromReader(message: GetTopicArticlesRequest, reader: jspb.BinaryReader): GetTopicArticlesRequest;
}

export namespace GetTopicArticlesRequest {
  export type AsObject = {
    topicsList: Array<string>,
  }
}

