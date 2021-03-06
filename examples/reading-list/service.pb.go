// Code generated by protoc-gen-go.
// source: service.proto
// DO NOT EDIT!

/*
Package readinglist is a generated protocol buffer package.

It is generated from these files:
	service.proto

It has these top-level messages:
	PutLinkProtoJSONRequest
	GetListProtoJSONRequest
	Link
	LinkRequest
	Links
	Message
*/
package readinglist

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type PutLinkProtoJSONRequest struct {
	// The transport to respond with (Protobuf or JSON)
	ProtoJSON string `protobuf:"bytes,1,opt,name=protoJSON" json:"protoJSON,omitempty"`
	// The Link to save
	Request *LinkRequest `protobuf:"bytes,2,opt,name=request" json:"request,omitempty"`
}

func (m *PutLinkProtoJSONRequest) Reset()                    { *m = PutLinkProtoJSONRequest{} }
func (m *PutLinkProtoJSONRequest) String() string            { return proto.CompactTextString(m) }
func (*PutLinkProtoJSONRequest) ProtoMessage()               {}
func (*PutLinkProtoJSONRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *PutLinkProtoJSONRequest) GetProtoJSON() string {
	if m != nil {
		return m.ProtoJSON
	}
	return ""
}

func (m *PutLinkProtoJSONRequest) GetRequest() *LinkRequest {
	if m != nil {
		return m.Request
	}
	return nil
}

type GetListProtoJSONRequest struct {
	// The limit of links to fetch
	Limit int32 `protobuf:"varint,2,opt,name=limit" json:"limit,omitempty"`
	// The transport to respond with (Protobuf or JSON).
	ProtoJSON string `protobuf:"bytes,1,opt,name=protoJSON" json:"protoJSON,omitempty"`
}

func (m *GetListProtoJSONRequest) Reset()                    { *m = GetListProtoJSONRequest{} }
func (m *GetListProtoJSONRequest) String() string            { return proto.CompactTextString(m) }
func (*GetListProtoJSONRequest) ProtoMessage()               {}
func (*GetListProtoJSONRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *GetListProtoJSONRequest) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *GetListProtoJSONRequest) GetProtoJSON() string {
	if m != nil {
		return m.ProtoJSON
	}
	return ""
}

type Link struct {
	Url string `protobuf:"bytes,1,opt,name=url" json:"url,omitempty"`
}

func (m *Link) Reset()                    { *m = Link{} }
func (m *Link) String() string            { return proto.CompactTextString(m) }
func (*Link) ProtoMessage()               {}
func (*Link) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Link) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

type LinkRequest struct {
	// True to remove this link from the user's list.
	Delete bool  `protobuf:"varint,1,opt,name=delete" json:"delete,omitempty"`
	Link   *Link `protobuf:"bytes,2,opt,name=link" json:"link,omitempty"`
}

func (m *LinkRequest) Reset()                    { *m = LinkRequest{} }
func (m *LinkRequest) String() string            { return proto.CompactTextString(m) }
func (*LinkRequest) ProtoMessage()               {}
func (*LinkRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *LinkRequest) GetDelete() bool {
	if m != nil {
		return m.Delete
	}
	return false
}

func (m *LinkRequest) GetLink() *Link {
	if m != nil {
		return m.Link
	}
	return nil
}

type Links struct {
	Links []*Link `protobuf:"bytes,1,rep,name=links" json:"links,omitempty"`
}

func (m *Links) Reset()                    { *m = Links{} }
func (m *Links) String() string            { return proto.CompactTextString(m) }
func (*Links) ProtoMessage()               {}
func (*Links) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Links) GetLinks() []*Link {
	if m != nil {
		return m.Links
	}
	return nil
}

type Message struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
}

func (m *Message) Reset()                    { *m = Message{} }
func (m *Message) String() string            { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()               {}
func (*Message) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *Message) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*PutLinkProtoJSONRequest)(nil), "readinglist.PutLinkProtoJSONRequest")
	proto.RegisterType((*GetListProtoJSONRequest)(nil), "readinglist.GetListProtoJSONRequest")
	proto.RegisterType((*Link)(nil), "readinglist.Link")
	proto.RegisterType((*LinkRequest)(nil), "readinglist.LinkRequest")
	proto.RegisterType((*Links)(nil), "readinglist.Links")
	proto.RegisterType((*Message)(nil), "readinglist.Message")
}

func init() { proto.RegisterFile("service.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 292 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x7c, 0x51, 0xed, 0x4e, 0xc2, 0x30,
	0x14, 0x65, 0xf2, 0x25, 0x97, 0x98, 0xcc, 0x1b, 0x22, 0x0d, 0xf1, 0xc7, 0x52, 0x35, 0xf2, 0x6b,
	0x31, 0xf3, 0x21, 0x4c, 0xcc, 0xc0, 0xa5, 0x3c, 0x01, 0xca, 0x0d, 0x69, 0x36, 0x06, 0xae, 0x9d,
	0x8f, 0xe6, 0xf3, 0x99, 0xae, 0x5d, 0x84, 0x4d, 0xfc, 0xd7, 0xd3, 0x73, 0x7a, 0xee, 0xb9, 0x3d,
	0x70, 0xa5, 0xa8, 0xf8, 0x92, 0x1f, 0x14, 0x1e, 0x8a, 0xbd, 0xde, 0xe3, 0xb8, 0xa0, 0xf5, 0x46,
	0xe6, 0xdb, 0x4c, 0x2a, 0xcd, 0x53, 0x98, 0x26, 0xa5, 0x8e, 0x65, 0x9e, 0x26, 0x86, 0x7c, 0x5d,
	0xbd, 0x2d, 0x05, 0x7d, 0x96, 0xa4, 0x34, 0xde, 0xc2, 0xe8, 0x50, 0xdf, 0x31, 0x2f, 0xf0, 0xe6,
	0x23, 0xf1, 0x7b, 0x81, 0x11, 0x0c, 0x0b, 0x2b, 0x64, 0x17, 0x81, 0x37, 0x1f, 0x47, 0x2c, 0x3c,
	0xf2, 0x0d, 0x8d, 0xa3, 0x33, 0x12, 0xb5, 0x90, 0x2f, 0x60, 0xfa, 0x42, 0x3a, 0x96, 0x4a, 0xb7,
	0x86, 0x4d, 0xa0, 0x9f, 0xc9, 0x9d, 0xb4, 0x66, 0x7d, 0x61, 0xc1, 0xff, 0x11, 0x38, 0x83, 0x9e,
	0x19, 0x83, 0x3e, 0x74, 0xcb, 0x22, 0x73, 0xbc, 0x39, 0xf2, 0x18, 0xc6, 0x47, 0x01, 0xf0, 0x06,
	0x06, 0x1b, 0xca, 0x48, 0x53, 0xa5, 0xb9, 0x14, 0x0e, 0xe1, 0x03, 0xf4, 0x32, 0x99, 0xa7, 0x6e,
	0x81, 0xeb, 0xf6, 0x02, 0x15, 0xcd, 0x9f, 0xa0, 0x6f, 0x90, 0xc2, 0x47, 0x13, 0x32, 0x4f, 0x15,
	0xf3, 0x82, 0xee, 0xdf, 0x0f, 0x2c, 0xcf, 0xef, 0x60, 0xb8, 0x20, 0xa5, 0xd6, 0x5b, 0x42, 0x06,
	0xc3, 0x9d, 0x3d, 0xba, 0x80, 0x35, 0x8c, 0xbe, 0x3d, 0x40, 0x61, 0x0d, 0xcc, 0x97, 0xac, 0x6c,
	0x49, 0x98, 0x80, 0xdf, 0x6c, 0x04, 0xef, 0x4f, 0x26, 0x9d, 0x29, 0x6c, 0x36, 0x39, 0x51, 0xb9,
	0x00, 0xbc, 0x83, 0x4b, 0xf0, 0x9b, 0xdf, 0xde, 0x70, 0x3c, 0xd3, 0xca, 0x0c, 0x5b, 0x1b, 0x2a,
	0xde, 0x79, 0x1f, 0x54, 0x15, 0x3c, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0xa0, 0x68, 0xe6, 0xce,
	0x58, 0x02, 0x00, 0x00,
}
