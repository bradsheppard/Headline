from google.protobuf import empty_pb2 as _empty_pb2
from google.protobuf import timestamp_pb2 as _timestamp_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class Topic(_message.Message):
    __slots__ = ("name", "lastUpdated")
    NAME_FIELD_NUMBER: _ClassVar[int]
    LASTUPDATED_FIELD_NUMBER: _ClassVar[int]
    name: str
    lastUpdated: _timestamp_pb2.Timestamp
    def __init__(self, name: _Optional[str] = ..., lastUpdated: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ...) -> None: ...

class TopicResponse(_message.Message):
    __slots__ = ("topics",)
    TOPICS_FIELD_NUMBER: _ClassVar[int]
    topics: _containers.RepeatedCompositeFieldContainer[Topic]
    def __init__(self, topics: _Optional[_Iterable[_Union[Topic, _Mapping]]] = ...) -> None: ...

class GetPendingTopicsRequest(_message.Message):
    __slots__ = ("lastUpdated",)
    LASTUPDATED_FIELD_NUMBER: _ClassVar[int]
    lastUpdated: _timestamp_pb2.Timestamp
    def __init__(self, lastUpdated: _Optional[_Union[_timestamp_pb2.Timestamp, _Mapping]] = ...) -> None: ...

class AddTopicsRequest(_message.Message):
    __slots__ = ("topics",)
    TOPICS_FIELD_NUMBER: _ClassVar[int]
    topics: _containers.RepeatedCompositeFieldContainer[Topic]
    def __init__(self, topics: _Optional[_Iterable[_Union[Topic, _Mapping]]] = ...) -> None: ...

class RemoveTopicsRequest(_message.Message):
    __slots__ = ("topicNames",)
    TOPICNAMES_FIELD_NUMBER: _ClassVar[int]
    topicNames: _containers.RepeatedScalarFieldContainer[str]
    def __init__(self, topicNames: _Optional[_Iterable[str]] = ...) -> None: ...
