// Code generated by protoc-gen-go.
// source: github.com/micro/monitoring-srv/proto/monitor/monitor.proto
// DO NOT EDIT!

/*
Package monitor is a generated protocol buffer package.

It is generated from these files:
	github.com/micro/monitoring-srv/proto/monitor/monitor.proto

It has these top-level messages:
	HealthChecksRequest
	HealthChecksResponse
*/
package monitor

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import monitor1 "github.com/micro/go-platform/monitor/proto"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "golang.org/x/net/context"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type HealthChecksRequest struct {
	Id     string                      `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Status monitor1.HealthCheck_Status `protobuf:"varint,2,opt,name=status,enum=HealthCheck_Status" json:"status,omitempty"`
	Limit  uint64                      `protobuf:"varint,3,opt,name=limit" json:"limit,omitempty"`
	Offset uint64                      `protobuf:"varint,4,opt,name=offset" json:"offset,omitempty"`
}

func (m *HealthChecksRequest) Reset()                    { *m = HealthChecksRequest{} }
func (m *HealthChecksRequest) String() string            { return proto.CompactTextString(m) }
func (*HealthChecksRequest) ProtoMessage()               {}
func (*HealthChecksRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type HealthChecksResponse struct {
	HealthChecks []*monitor1.HealthCheck `protobuf:"bytes,1,rep,name=healthChecks" json:"healthChecks,omitempty"`
}

func (m *HealthChecksResponse) Reset()                    { *m = HealthChecksResponse{} }
func (m *HealthChecksResponse) String() string            { return proto.CompactTextString(m) }
func (*HealthChecksResponse) ProtoMessage()               {}
func (*HealthChecksResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *HealthChecksResponse) GetHealthChecks() []*monitor1.HealthCheck {
	if m != nil {
		return m.HealthChecks
	}
	return nil
}

func init() {
	proto.RegisterType((*HealthChecksRequest)(nil), "HealthChecksRequest")
	proto.RegisterType((*HealthChecksResponse)(nil), "HealthChecksResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Monitor service

type MonitorClient interface {
	HealthChecks(ctx context.Context, in *HealthChecksRequest, opts ...client.CallOption) (*HealthChecksResponse, error)
}

type monitorClient struct {
	c           client.Client
	serviceName string
}

func NewMonitorClient(serviceName string, c client.Client) MonitorClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "monitor"
	}
	return &monitorClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *monitorClient) HealthChecks(ctx context.Context, in *HealthChecksRequest, opts ...client.CallOption) (*HealthChecksResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Monitor.HealthChecks", in)
	out := new(HealthChecksResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Monitor service

type MonitorHandler interface {
	HealthChecks(context.Context, *HealthChecksRequest, *HealthChecksResponse) error
}

func RegisterMonitorHandler(s server.Server, hdlr MonitorHandler) {
	s.Handle(s.NewHandler(&Monitor{hdlr}))
}

type Monitor struct {
	MonitorHandler
}

func (h *Monitor) HealthChecks(ctx context.Context, in *HealthChecksRequest, out *HealthChecksResponse) error {
	return h.MonitorHandler.HealthChecks(ctx, in, out)
}

var fileDescriptor0 = []byte{
	// 234 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x8f, 0x3d, 0x4f, 0xc3, 0x30,
	0x10, 0x86, 0xe9, 0x07, 0x41, 0x1c, 0xa1, 0x83, 0x5b, 0xa4, 0xa8, 0x13, 0x32, 0x4b, 0x97, 0x3a,
	0x52, 0x58, 0x10, 0x88, 0x89, 0xa5, 0x0b, 0x0b, 0xfc, 0x82, 0x34, 0x38, 0xb1, 0x45, 0x9c, 0x0b,
	0xbe, 0x0b, 0xbf, 0x9f, 0x90, 0x8a, 0xca, 0x15, 0x99, 0xac, 0x7b, 0xed, 0xe7, 0x9e, 0xd7, 0xf0,
	0x54, 0x59, 0x36, 0xdd, 0x5e, 0x15, 0xe8, 0x52, 0x67, 0x0b, 0x8f, 0xa9, 0xc3, 0xc6, 0x32, 0x7a,
	0xdb, 0x54, 0x5b, 0xf2, 0xdf, 0x69, 0xeb, 0x91, 0x8f, 0xe1, 0xdf, 0xa9, 0x86, 0x74, 0xfd, 0xf0,
	0x0f, 0xae, 0x70, 0xdb, 0xd6, 0x39, 0x97, 0xe8, 0xdd, 0x91, 0x39, 0xd9, 0x70, 0x20, 0xa5, 0x86,
	0xe5, 0x4e, 0xe7, 0x35, 0x9b, 0x17, 0xa3, 0x8b, 0x4f, 0x7a, 0xd3, 0x5f, 0x9d, 0x26, 0x16, 0x00,
	0x53, 0xfb, 0x91, 0x4c, 0x6e, 0x27, 0x9b, 0x4b, 0x71, 0x07, 0x11, 0x71, 0xce, 0x1d, 0x25, 0xd3,
	0x7e, 0x5e, 0x64, 0x4b, 0x15, 0x10, 0xea, 0x7d, 0xb8, 0x12, 0xd7, 0x70, 0x5e, 0x5b, 0x67, 0x39,
	0x99, 0xf5, 0x6f, 0xe6, 0x62, 0x01, 0x11, 0x96, 0x25, 0x69, 0x4e, 0xe6, 0xbf, 0xb3, 0x7c, 0x84,
	0xd5, 0xa9, 0x86, 0x5a, 0x6c, 0x48, 0x0b, 0x09, 0xb1, 0x09, 0xf2, 0xde, 0x38, 0xdb, 0x5c, 0x65,
	0x71, 0x68, 0xc8, 0x76, 0x70, 0xf1, 0x7a, 0xe8, 0x2c, 0x9e, 0x21, 0x0e, 0xd7, 0x88, 0x95, 0x1a,
	0x29, 0xbf, 0xbe, 0x51, 0x63, 0x2e, 0x79, 0xb6, 0x8f, 0x86, 0x3f, 0xdf, 0xff, 0x04, 0x00, 0x00,
	0xff, 0xff, 0x4b, 0x9e, 0x23, 0x69, 0x6c, 0x01, 0x00, 0x00,
}
