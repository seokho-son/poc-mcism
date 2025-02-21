// Code generated by protoc-gen-go. DO NOT EDIT.
// source: mcism_agent.proto

/*
Package resourcestat is a generated protocol buffer package.

Proof of Concepts for the Cloud-Barista Multi-Cloud Project.
	* Cloud-Barista: https://github.com/cloud-barista

This Definition is interfaces between MCISM Master and Agent.

by powerkim@powerkim.co.kr, 2019.03.15

It is generated from these files:
	mcism_agent.proto

It has these top-level messages:
	ResourceStatRequest
	ResourceStatReply
*/
package resourcestat

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// The request message with empty
type ResourceStatRequest struct {
}

func (m *ResourceStatRequest) Reset()                    { *m = ResourceStatRequest{} }
func (m *ResourceStatRequest) String() string            { return proto.CompactTextString(m) }
func (*ResourceStatRequest) ProtoMessage()               {}
func (*ResourceStatRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// The response message containing the resouce status
type ResourceStatReply struct {
	Servername string `protobuf:"bytes,1,opt,name=servername" json:"servername,omitempty"`
	Cpu        string `protobuf:"bytes,2,opt,name=cpu" json:"cpu,omitempty"`
	Mem        string `protobuf:"bytes,3,opt,name=mem" json:"mem,omitempty"`
	Dsk        string `protobuf:"bytes,4,opt,name=dsk" json:"dsk,omitempty"`
}

func (m *ResourceStatReply) Reset()                    { *m = ResourceStatReply{} }
func (m *ResourceStatReply) String() string            { return proto.CompactTextString(m) }
func (*ResourceStatReply) ProtoMessage()               {}
func (*ResourceStatReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ResourceStatReply) GetServername() string {
	if m != nil {
		return m.Servername
	}
	return ""
}

func (m *ResourceStatReply) GetCpu() string {
	if m != nil {
		return m.Cpu
	}
	return ""
}

func (m *ResourceStatReply) GetMem() string {
	if m != nil {
		return m.Mem
	}
	return ""
}

func (m *ResourceStatReply) GetDsk() string {
	if m != nil {
		return m.Dsk
	}
	return ""
}

func init() {
	proto.RegisterType((*ResourceStatRequest)(nil), "resourcestat.ResourceStatRequest")
	proto.RegisterType((*ResourceStatReply)(nil), "resourcestat.ResourceStatReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for ResourceStat service

type ResourceStatClient interface {
	// Sends a request of Resource status
	GetResourceStat(ctx context.Context, in *ResourceStatRequest, opts ...grpc.CallOption) (*ResourceStatReply, error)
}

type resourceStatClient struct {
	cc *grpc.ClientConn
}

func NewResourceStatClient(cc *grpc.ClientConn) ResourceStatClient {
	return &resourceStatClient{cc}
}

func (c *resourceStatClient) GetResourceStat(ctx context.Context, in *ResourceStatRequest, opts ...grpc.CallOption) (*ResourceStatReply, error) {
	out := new(ResourceStatReply)
	err := grpc.Invoke(ctx, "/resourcestat.ResourceStat/GetResourceStat", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ResourceStat service

type ResourceStatServer interface {
	// Sends a request of Resource status
	GetResourceStat(context.Context, *ResourceStatRequest) (*ResourceStatReply, error)
}

func RegisterResourceStatServer(s *grpc.Server, srv ResourceStatServer) {
	s.RegisterService(&_ResourceStat_serviceDesc, srv)
}

func _ResourceStat_GetResourceStat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResourceStatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceStatServer).GetResourceStat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/resourcestat.ResourceStat/GetResourceStat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceStatServer).GetResourceStat(ctx, req.(*ResourceStatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ResourceStat_serviceDesc = grpc.ServiceDesc{
	ServiceName: "resourcestat.ResourceStat",
	HandlerType: (*ResourceStatServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetResourceStat",
			Handler:    _ResourceStat_GetResourceStat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mcism_agent.proto",
}

func init() { proto.RegisterFile("mcism_agent.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 210 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4e, 0x4b, 0x2c, 0xca,
	0xcd, 0xcf, 0xcb, 0x8c, 0x4f, 0x4c, 0x4f, 0xcd, 0x2b, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0xe2, 0x29, 0x4a, 0x2d, 0xce, 0x2f, 0x2d, 0x4a, 0x4e, 0x2d, 0x2e, 0x49, 0x2c, 0x51, 0x12, 0xe5,
	0x12, 0x0e, 0x82, 0xf2, 0x83, 0x4b, 0x12, 0x4b, 0x82, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x94,
	0x32, 0xb9, 0x04, 0x51, 0x85, 0x0b, 0x72, 0x2a, 0x85, 0xe4, 0xb8, 0xb8, 0x8a, 0x53, 0x8b, 0xca,
	0x52, 0x8b, 0xf2, 0x12, 0x73, 0x53, 0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0x90, 0x44, 0x84,
	0x04, 0xb8, 0x98, 0x93, 0x0b, 0x4a, 0x25, 0x98, 0xc0, 0x12, 0x20, 0x26, 0x48, 0x24, 0x37, 0x35,
	0x57, 0x82, 0x19, 0x22, 0x92, 0x9b, 0x9a, 0x0b, 0x12, 0x49, 0x29, 0xce, 0x96, 0x60, 0x81, 0x88,
	0xa4, 0x14, 0x67, 0x1b, 0xa5, 0x73, 0xf1, 0x20, 0x5b, 0x25, 0x14, 0xce, 0xc5, 0xef, 0x9e, 0x5a,
	0x82, 0x22, 0xa4, 0xa8, 0x87, 0xec, 0x66, 0x3d, 0x2c, 0x0e, 0x96, 0x92, 0xc7, 0xa7, 0xa4, 0x20,
	0xa7, 0x52, 0x89, 0xc1, 0xc9, 0x80, 0x4b, 0x32, 0x29, 0xb1, 0x28, 0xb3, 0xb8, 0x24, 0x51, 0x2f,
	0x37, 0x3f, 0x0f, 0xa4, 0x1e, 0xa1, 0xc1, 0x09, 0xc5, 0xbb, 0x01, 0xa0, 0x80, 0x0a, 0x60, 0x4c,
	0x62, 0x03, 0x87, 0x98, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x08, 0xaa, 0x7c, 0xc6, 0x48, 0x01,
	0x00, 0x00,
}
