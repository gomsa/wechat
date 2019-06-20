// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/wechat/wechat.proto

package aliyun

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "golang.org/x/net/context"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Request struct {
	Domain               string            `protobuf:"bytes,1,opt,name=Domain,proto3" json:"Domain,omitempty"`
	ApiName              string            `protobuf:"bytes,2,opt,name=ApiName,proto3" json:"ApiName,omitempty"`
	QueryParams          map[string]string `protobuf:"bytes,3,rep,name=QueryParams,proto3" json:"QueryParams,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_c33a8144c2244122, []int{0}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

func (m *Request) GetApiName() string {
	if m != nil {
		return m.ApiName
	}
	return ""
}

func (m *Request) GetQueryParams() map[string]string {
	if m != nil {
		return m.QueryParams
	}
	return nil
}

type Response struct {
	Content              string   `protobuf:"bytes,1,opt,name=Content,proto3" json:"Content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_c33a8144c2244122, []int{1}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func init() {
	proto.RegisterType((*Request)(nil), "aliyun.Request")
	proto.RegisterMapType((map[string]string)(nil), "aliyun.Request.QueryParamsEntry")
	proto.RegisterType((*Response)(nil), "aliyun.Response")
}

func init() { proto.RegisterFile("proto/wechat/wechat.proto", fileDescriptor_c33a8144c2244122) }

var fileDescriptor_c33a8144c2244122 = []byte{
	// 240 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2c, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x2f, 0x4f, 0x4d, 0xce, 0x48, 0x2c, 0x81, 0x52, 0x7a, 0x60, 0x31, 0x21, 0xb6, 0xc4,
	0x9c, 0xcc, 0xca, 0xd2, 0x3c, 0xa5, 0xfd, 0x8c, 0x5c, 0xec, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5,
	0x25, 0x42, 0x62, 0x5c, 0x6c, 0x2e, 0xf9, 0xb9, 0x89, 0x99, 0x79, 0x12, 0x8c, 0x0a, 0x8c, 0x1a,
	0x9c, 0x41, 0x50, 0x9e, 0x90, 0x04, 0x17, 0xbb, 0x63, 0x41, 0xa6, 0x5f, 0x62, 0x6e, 0xaa, 0x04,
	0x13, 0x58, 0x02, 0xc6, 0x15, 0x72, 0xe2, 0xe2, 0x0e, 0x2c, 0x4d, 0x2d, 0xaa, 0x0c, 0x48, 0x2c,
	0x4a, 0xcc, 0x2d, 0x96, 0x60, 0x56, 0x60, 0xd6, 0xe0, 0x36, 0x52, 0xd0, 0x83, 0x98, 0xad, 0x07,
	0x35, 0x57, 0x0f, 0x49, 0x89, 0x6b, 0x5e, 0x49, 0x51, 0x65, 0x10, 0xb2, 0x26, 0x29, 0x3b, 0x2e,
	0x01, 0x74, 0x05, 0x42, 0x02, 0x5c, 0xcc, 0xd9, 0xa9, 0x95, 0x50, 0x67, 0x80, 0x98, 0x42, 0x22,
	0x5c, 0xac, 0x65, 0x89, 0x39, 0xa5, 0x30, 0x17, 0x40, 0x38, 0x56, 0x4c, 0x16, 0x8c, 0x4a, 0x2a,
	0x5c, 0x1c, 0x41, 0xa9, 0xc5, 0x05, 0xf9, 0x79, 0xc5, 0xa9, 0x20, 0x97, 0x3a, 0xe7, 0xe7, 0x95,
	0xa4, 0xe6, 0x95, 0x40, 0xf5, 0xc2, 0xb8, 0x46, 0xae, 0x5c, 0x6c, 0xe1, 0x60, 0xff, 0x0b, 0x59,
	0x73, 0x89, 0x04, 0x14, 0xe5, 0x27, 0xa7, 0x16, 0x17, 0x3b, 0xe7, 0xe7, 0xe6, 0xe6, 0xe7, 0xc1,
	0x7c, 0xcf, 0x8f, 0xe6, 0x6c, 0x29, 0x01, 0x84, 0x00, 0xc4, 0x78, 0x25, 0x86, 0x24, 0x36, 0x70,
	0xe8, 0x19, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xc0, 0x44, 0x26, 0x95, 0x5a, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Wechat service

type WechatClient interface {
	// 处理请求
	ProcessCommonRequest(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
}

type wechatClient struct {
	c           client.Client
	serviceName string
}

func NewWechatClient(serviceName string, c client.Client) WechatClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "aliyun"
	}
	return &wechatClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *wechatClient) ProcessCommonRequest(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "Wechat.ProcessCommonRequest", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Wechat service

type WechatHandler interface {
	// 处理请求
	ProcessCommonRequest(context.Context, *Request, *Response) error
}

func RegisterWechatHandler(s server.Server, hdlr WechatHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&Wechat{hdlr}, opts...))
}

type Wechat struct {
	WechatHandler
}

func (h *Wechat) ProcessCommonRequest(ctx context.Context, in *Request, out *Response) error {
	return h.WechatHandler.ProcessCommonRequest(ctx, in, out)
}