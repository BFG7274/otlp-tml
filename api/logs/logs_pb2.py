# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: logs.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\nlogs.proto\x1a\x1bgoogle/protobuf/empty.proto\"l\n\x06LogMsg\x12\x11\n\ttimestamp\x18\x01 \x01(\x03\x12\x0b\n\x03msg\x18\x02 \x01(\t\x12\r\n\x05level\x18\x03 \x01(\t\"3\n\x04Type\x12\t\n\x05\x64\x65\x62ug\x10\x00\x12\x08\n\x04info\x10\x01\x12\x0b\n\x07warning\x10\x02\x12\t\n\x05\x65rror\x10\x03\x32\x39\n\x08LogsData\x12-\n\x08WriteLog\x12\x07.LogMsg\x1a\x16.google.protobuf.Empty\"\x00\x62\x06proto3')



_LOGMSG = DESCRIPTOR.message_types_by_name['LogMsg']
_LOGMSG_TYPE = _LOGMSG.enum_types_by_name['Type']
LogMsg = _reflection.GeneratedProtocolMessageType('LogMsg', (_message.Message,), {
  'DESCRIPTOR' : _LOGMSG,
  '__module__' : 'logs_pb2'
  # @@protoc_insertion_point(class_scope:LogMsg)
  })
_sym_db.RegisterMessage(LogMsg)

_LOGSDATA = DESCRIPTOR.services_by_name['LogsData']
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  _LOGMSG._serialized_start=43
  _LOGMSG._serialized_end=151
  _LOGMSG_TYPE._serialized_start=100
  _LOGMSG_TYPE._serialized_end=151
  _LOGSDATA._serialized_start=153
  _LOGSDATA._serialized_end=210
# @@protoc_insertion_point(module_scope)
