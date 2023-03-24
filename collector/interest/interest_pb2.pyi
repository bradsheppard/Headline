from google.protobuf import empty_pb2 as _empty_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class AddInterestsRequest(_message.Message):
    __slots__ = ["interests", "userId"]
    INTERESTS_FIELD_NUMBER: _ClassVar[int]
    USERID_FIELD_NUMBER: _ClassVar[int]
    interests: _containers.RepeatedCompositeFieldContainer[CreateInterest]
    userId: int
    def __init__(self, interests: _Optional[_Iterable[_Union[CreateInterest, _Mapping]]] = ..., userId: _Optional[int] = ...) -> None: ...

class CreateInterest(_message.Message):
    __slots__ = ["name", "userId"]
    NAME_FIELD_NUMBER: _ClassVar[int]
    USERID_FIELD_NUMBER: _ClassVar[int]
    name: str
    userId: int
    def __init__(self, name: _Optional[str] = ..., userId: _Optional[int] = ...) -> None: ...

class DeleteInterestsRequest(_message.Message):
    __slots__ = ["ids"]
    IDS_FIELD_NUMBER: _ClassVar[int]
    ids: _containers.RepeatedScalarFieldContainer[int]
    def __init__(self, ids: _Optional[_Iterable[int]] = ...) -> None: ...

class GetInterestsRequest(_message.Message):
    __slots__ = ["userId"]
    USERID_FIELD_NUMBER: _ClassVar[int]
    userId: int
    def __init__(self, userId: _Optional[int] = ...) -> None: ...

class Interest(_message.Message):
    __slots__ = ["id", "name", "userId"]
    ID_FIELD_NUMBER: _ClassVar[int]
    NAME_FIELD_NUMBER: _ClassVar[int]
    USERID_FIELD_NUMBER: _ClassVar[int]
    id: int
    name: str
    userId: int
    def __init__(self, id: _Optional[int] = ..., name: _Optional[str] = ..., userId: _Optional[int] = ...) -> None: ...

class InterestResponse(_message.Message):
    __slots__ = ["interests"]
    INTERESTS_FIELD_NUMBER: _ClassVar[int]
    interests: _containers.RepeatedCompositeFieldContainer[Interest]
    def __init__(self, interests: _Optional[_Iterable[_Union[Interest, _Mapping]]] = ...) -> None: ...
