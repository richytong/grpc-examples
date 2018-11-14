// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('grpc');
var cache_pb = require('./cache_pb.js');

function serialize_cache_GetReq(arg) {
  if (!(arg instanceof cache_pb.GetReq)) {
    throw new Error('Expected argument of type cache.GetReq');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_cache_GetReq(buffer_arg) {
  return cache_pb.GetReq.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_cache_GetResp(arg) {
  if (!(arg instanceof cache_pb.GetResp)) {
    throw new Error('Expected argument of type cache.GetResp');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_cache_GetResp(buffer_arg) {
  return cache_pb.GetResp.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_cache_SetReq(arg) {
  if (!(arg instanceof cache_pb.SetReq)) {
    throw new Error('Expected argument of type cache.SetReq');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_cache_SetReq(buffer_arg) {
  return cache_pb.SetReq.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_cache_SetResp(arg) {
  if (!(arg instanceof cache_pb.SetResp)) {
    throw new Error('Expected argument of type cache.SetResp');
  }
  return new Buffer(arg.serializeBinary());
}

function deserialize_cache_SetResp(buffer_arg) {
  return cache_pb.SetResp.deserializeBinary(new Uint8Array(buffer_arg));
}


var CacheService = exports.CacheService = {
  set: {
    path: '/cache.Cache/Set',
    requestStream: false,
    responseStream: false,
    requestType: cache_pb.SetReq,
    responseType: cache_pb.SetResp,
    requestSerialize: serialize_cache_SetReq,
    requestDeserialize: deserialize_cache_SetReq,
    responseSerialize: serialize_cache_SetResp,
    responseDeserialize: deserialize_cache_SetResp,
  },
  get: {
    path: '/cache.Cache/Get',
    requestStream: false,
    responseStream: false,
    requestType: cache_pb.GetReq,
    responseType: cache_pb.GetResp,
    requestSerialize: serialize_cache_GetReq,
    requestDeserialize: deserialize_cache_GetReq,
    responseSerialize: serialize_cache_GetResp,
    responseDeserialize: deserialize_cache_GetResp,
  },
};

exports.CacheClient = grpc.makeGenericClientConstructor(CacheService);
