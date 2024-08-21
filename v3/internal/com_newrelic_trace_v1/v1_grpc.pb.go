// Copyright 2020 New Relic Corporation. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.27.3
// source: v3/internal/com_newrelic_trace_v1/v1.proto

package com_newrelic_trace_v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	IngestService_RecordSpan_FullMethodName      = "/com.newrelic.trace.v1.IngestService/RecordSpan"
	IngestService_RecordSpanBatch_FullMethodName = "/com.newrelic.trace.v1.IngestService/RecordSpanBatch"
)

// IngestServiceClient is the client API for IngestService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type IngestServiceClient interface {
	// Accepts a stream of Span messages, and returns an irregular stream of
	// RecordStatus messages.
	RecordSpan(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[Span, RecordStatus], error)
	// Accepts a stream of SpanBatch messages, and returns an irregular
	// stream of RecordStatus messages. This endpoint can be used to improve
	// throughput when Span messages are small
	RecordSpanBatch(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[SpanBatch, RecordStatus], error)
}

type ingestServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewIngestServiceClient(cc grpc.ClientConnInterface) IngestServiceClient {
	return &ingestServiceClient{cc}
}

func (c *ingestServiceClient) RecordSpan(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[Span, RecordStatus], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &IngestService_ServiceDesc.Streams[0], IngestService_RecordSpan_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[Span, RecordStatus]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type IngestService_RecordSpanClient = grpc.BidiStreamingClient[Span, RecordStatus]

func (c *ingestServiceClient) RecordSpanBatch(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[SpanBatch, RecordStatus], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &IngestService_ServiceDesc.Streams[1], IngestService_RecordSpanBatch_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[SpanBatch, RecordStatus]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type IngestService_RecordSpanBatchClient = grpc.BidiStreamingClient[SpanBatch, RecordStatus]

// IngestServiceServer is the server API for IngestService service.
// All implementations must embed UnimplementedIngestServiceServer
// for forward compatibility.
type IngestServiceServer interface {
	// Accepts a stream of Span messages, and returns an irregular stream of
	// RecordStatus messages.
	RecordSpan(grpc.BidiStreamingServer[Span, RecordStatus]) error
	// Accepts a stream of SpanBatch messages, and returns an irregular
	// stream of RecordStatus messages. This endpoint can be used to improve
	// throughput when Span messages are small
	RecordSpanBatch(grpc.BidiStreamingServer[SpanBatch, RecordStatus]) error
	mustEmbedUnimplementedIngestServiceServer()
}

// UnimplementedIngestServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedIngestServiceServer struct{}

func (UnimplementedIngestServiceServer) RecordSpan(grpc.BidiStreamingServer[Span, RecordStatus]) error {
	return status.Errorf(codes.Unimplemented, "method RecordSpan not implemented")
}
func (UnimplementedIngestServiceServer) RecordSpanBatch(grpc.BidiStreamingServer[SpanBatch, RecordStatus]) error {
	return status.Errorf(codes.Unimplemented, "method RecordSpanBatch not implemented")
}
func (UnimplementedIngestServiceServer) mustEmbedUnimplementedIngestServiceServer() {}
func (UnimplementedIngestServiceServer) testEmbeddedByValue()                       {}

// UnsafeIngestServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to IngestServiceServer will
// result in compilation errors.
type UnsafeIngestServiceServer interface {
	mustEmbedUnimplementedIngestServiceServer()
}

func RegisterIngestServiceServer(s grpc.ServiceRegistrar, srv IngestServiceServer) {
	// If the following call pancis, it indicates UnimplementedIngestServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&IngestService_ServiceDesc, srv)
}

func _IngestService_RecordSpan_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(IngestServiceServer).RecordSpan(&grpc.GenericServerStream[Span, RecordStatus]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type IngestService_RecordSpanServer = grpc.BidiStreamingServer[Span, RecordStatus]

func _IngestService_RecordSpanBatch_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(IngestServiceServer).RecordSpanBatch(&grpc.GenericServerStream[SpanBatch, RecordStatus]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type IngestService_RecordSpanBatchServer = grpc.BidiStreamingServer[SpanBatch, RecordStatus]

// IngestService_ServiceDesc is the grpc.ServiceDesc for IngestService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var IngestService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "com.newrelic.trace.v1.IngestService",
	HandlerType: (*IngestServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "RecordSpan",
			Handler:       _IngestService_RecordSpan_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "RecordSpanBatch",
			Handler:       _IngestService_RecordSpanBatch_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "v3/internal/com_newrelic_trace_v1/v1.proto",
}