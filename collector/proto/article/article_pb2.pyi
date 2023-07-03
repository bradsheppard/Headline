from google.protobuf import empty_pb2 as _empty_pb2
from google.protobuf import timestamp_pb2 as _timestamp_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class Article(_message.Message):
    __slots__ = ["date", "description", "imageUrl", "source", "title", "url"]
    DATE_FIELD_NUMBER: _ClassVar[int]
    DESCRIPTION_FIELD_NUMBER: _ClassVar[int]
    IMAGEURL_FIELD_NUMBER: _ClassVar[int]
    SOURCE_FIELD_NUMBER: _ClassVar[int]
    TITLE_FIELD_NUMBER: _ClassVar[int]
    URL_FIELD_NUMBER: _ClassVar[int]
    date: _timestamp_pb2.Timestamp
    description: str
    imageUrl: str
    source: str
    title: str
    url: str
    def __init__(self, title: _Optional[str] = ..., description: _Optional[str] = ..., url: _Optional[str] = ..., imageUrl: _Optional[str] = ..., source: _Optional[str] = ..., date: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ...) -> None: ...

class Articles(_message.Message):
    __slots__ = ["articles"]
    ARTICLES_FIELD_NUMBER: _ClassVar[int]
    articles: _containers.RepeatedCompositeFieldContainer[Article]
    def __init__(self, articles: _Optional[_Iterable[_Union[Article, _Mapping]]] = ...) -> None: ...

class TopicArticles(_message.Message):
    __slots__ = ["topicArticles"]
    class TopicArticlesEntry(_message.Message):
        __slots__ = ["key", "value"]
        KEY_FIELD_NUMBER: _ClassVar[int]
        VALUE_FIELD_NUMBER: _ClassVar[int]
        key: str
        value: Articles
        def __init__(self, key: _Optional[str] = ..., value: _Optional[_Union[Articles, _Mapping]] = ...) -> None: ...
    TOPICARTICLES_FIELD_NUMBER: _ClassVar[int]
    topicArticles: _containers.MessageMap[str, Articles]
    def __init__(self, topicArticles: _Optional[_Mapping[str, Articles]] = ...) -> None: ...

class TopicNames(_message.Message):
    __slots__ = ["topics"]
    TOPICS_FIELD_NUMBER: _ClassVar[int]
    topics: _containers.RepeatedScalarFieldContainer[str]
    def __init__(self, topics: _Optional[_Iterable[str]] = ...) -> None: ...
