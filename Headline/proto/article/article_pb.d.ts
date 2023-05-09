import * as jspb from 'google-protobuf'

import * as google_protobuf_empty_pb from 'google-protobuf/google/protobuf/empty_pb';


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

  getInterest(): string;
  setInterest(value: string): Article;

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
    interest: string,
  }
}

export class UserArticles extends jspb.Message {
  getArticlesList(): Array<Article>;
  setArticlesList(value: Array<Article>): UserArticles;
  clearArticlesList(): UserArticles;
  addArticles(value?: Article, index?: number): Article;

  getUserid(): number;
  setUserid(value: number): UserArticles;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UserArticles.AsObject;
  static toObject(includeInstance: boolean, msg: UserArticles): UserArticles.AsObject;
  static serializeBinaryToWriter(message: UserArticles, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UserArticles;
  static deserializeBinaryFromReader(message: UserArticles, reader: jspb.BinaryReader): UserArticles;
}

export namespace UserArticles {
  export type AsObject = {
    articlesList: Array<Article.AsObject>,
    userid: number,
  }
}

export class User extends jspb.Message {
  getUserid(): number;
  setUserid(value: number): User;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): User.AsObject;
  static toObject(includeInstance: boolean, msg: User): User.AsObject;
  static serializeBinaryToWriter(message: User, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): User;
  static deserializeBinaryFromReader(message: User, reader: jspb.BinaryReader): User;
}

export namespace User {
  export type AsObject = {
    userid: number,
  }
}

