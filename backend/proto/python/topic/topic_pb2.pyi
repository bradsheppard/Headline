from google.protobuf import empty_pb2 as _empty_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class AddTopicsRequest(_message.Message):
    __slots__ = ["topics", "userId"]
    TOPICS_FIELD_NUMBER: _ClassVar[int]
    USERID_FIELD_NUMBER: _ClassVar[int]
    topics: _containers.RepeatedCompositeFieldContainer[Topic]
    userId: int
    def __init__(self, topics: _Optional[_Iterable[_Union[Topic, _Mapping]]] = ..., userId: _Optional[int] = ...) -> None: ...

class GetTopicsRequest(_message.Message):
    __slots__ = ["userId"]
    USERID_FIELD_NUMBER: _ClassVar[int]
    userId: int
    def __init__(self, userId: _Optional[int] = ...) -> None: ...

class RemoveTopicsRequest(_message.Message):
    __slots__ = ["topicNames", "userId"]
    TOPICNAMES_FIELD_NUMBER: _ClassVar[int]
    USERID_FIELD_NUMBER: _ClassVar[int]
    topicNames: _containers.RepeatedScalarFieldContainer[str]
    userId: int
    def __init__(self, topicNames: _Optional[_Iterable[str]] = ..., userId: _Optional[int] = ...) -> None: ...

class Topic(_message.Message):
    __slots__ = ["name"]
    NAME_FIELD_NUMBER: _ClassVar[int]
    name: str
    def __init__(self, name: _Optional[str] = ...) -> None: ...

class TopicResponse(_message.Message):
    __slots__ = ["topics"]
    TOPICS_FIELD_NUMBER: _ClassVar[int]
    topics: _containers.RepeatedCompositeFieldContainer[Topic]
    def __init__(self, topics: _Optional[_Iterable[_Union[Topic, _Mapping]]] = ...) -> None: ...
