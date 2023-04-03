from google.protobuf import empty_pb2 as _empty_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class Article(_message.Message):
    __slots__ = ["link", "summary", "title"]
    LINK_FIELD_NUMBER: _ClassVar[int]
    SUMMARY_FIELD_NUMBER: _ClassVar[int]
    TITLE_FIELD_NUMBER: _ClassVar[int]
    link: str
    summary: str
    title: str
    def __init__(self, title: _Optional[str] = ..., summary: _Optional[str] = ..., link: _Optional[str] = ...) -> None: ...

class User(_message.Message):
    __slots__ = ["userId"]
    USERID_FIELD_NUMBER: _ClassVar[int]
    userId: int
    def __init__(self, userId: _Optional[int] = ...) -> None: ...

class UserArticles(_message.Message):
    __slots__ = ["articles", "userId"]
    ARTICLES_FIELD_NUMBER: _ClassVar[int]
    USERID_FIELD_NUMBER: _ClassVar[int]
    articles: _containers.RepeatedCompositeFieldContainer[Article]
    userId: int
    def __init__(self, articles: _Optional[_Iterable[_Union[Article, _Mapping]]] = ..., userId: _Optional[int] = ...) -> None: ...
