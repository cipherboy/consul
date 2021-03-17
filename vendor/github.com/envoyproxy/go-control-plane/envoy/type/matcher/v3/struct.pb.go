// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/type/matcher/v3/struct.proto

package envoy_type_matcher_v3

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	math "math"
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

type StructMatcher struct {
	Path                 []*StructMatcher_PathSegment `protobuf:"bytes,2,rep,name=path,proto3" json:"path,omitempty"`
	Value                *ValueMatcher                `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *StructMatcher) Reset()         { *m = StructMatcher{} }
func (m *StructMatcher) String() string { return proto.CompactTextString(m) }
func (*StructMatcher) ProtoMessage()    {}
func (*StructMatcher) Descriptor() ([]byte, []int) {
	return fileDescriptor_2eb06d41ab351944, []int{0}
}

func (m *StructMatcher) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StructMatcher.Unmarshal(m, b)
}
func (m *StructMatcher) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StructMatcher.Marshal(b, m, deterministic)
}
func (m *StructMatcher) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StructMatcher.Merge(m, src)
}
func (m *StructMatcher) XXX_Size() int {
	return xxx_messageInfo_StructMatcher.Size(m)
}
func (m *StructMatcher) XXX_DiscardUnknown() {
	xxx_messageInfo_StructMatcher.DiscardUnknown(m)
}

var xxx_messageInfo_StructMatcher proto.InternalMessageInfo

func (m *StructMatcher) GetPath() []*StructMatcher_PathSegment {
	if m != nil {
		return m.Path
	}
	return nil
}

func (m *StructMatcher) GetValue() *ValueMatcher {
	if m != nil {
		return m.Value
	}
	return nil
}

type StructMatcher_PathSegment struct {
	// Types that are valid to be assigned to Segment:
	//	*StructMatcher_PathSegment_Key
	Segment              isStructMatcher_PathSegment_Segment `protobuf_oneof:"segment"`
	XXX_NoUnkeyedLiteral struct{}                            `json:"-"`
	XXX_unrecognized     []byte                              `json:"-"`
	XXX_sizecache        int32                               `json:"-"`
}

func (m *StructMatcher_PathSegment) Reset()         { *m = StructMatcher_PathSegment{} }
func (m *StructMatcher_PathSegment) String() string { return proto.CompactTextString(m) }
func (*StructMatcher_PathSegment) ProtoMessage()    {}
func (*StructMatcher_PathSegment) Descriptor() ([]byte, []int) {
	return fileDescriptor_2eb06d41ab351944, []int{0, 0}
}

func (m *StructMatcher_PathSegment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StructMatcher_PathSegment.Unmarshal(m, b)
}
func (m *StructMatcher_PathSegment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StructMatcher_PathSegment.Marshal(b, m, deterministic)
}
func (m *StructMatcher_PathSegment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StructMatcher_PathSegment.Merge(m, src)
}
func (m *StructMatcher_PathSegment) XXX_Size() int {
	return xxx_messageInfo_StructMatcher_PathSegment.Size(m)
}
func (m *StructMatcher_PathSegment) XXX_DiscardUnknown() {
	xxx_messageInfo_StructMatcher_PathSegment.DiscardUnknown(m)
}

var xxx_messageInfo_StructMatcher_PathSegment proto.InternalMessageInfo

type isStructMatcher_PathSegment_Segment interface {
	isStructMatcher_PathSegment_Segment()
}

type StructMatcher_PathSegment_Key struct {
	Key string `protobuf:"bytes,1,opt,name=key,proto3,oneof"`
}

func (*StructMatcher_PathSegment_Key) isStructMatcher_PathSegment_Segment() {}

func (m *StructMatcher_PathSegment) GetSegment() isStructMatcher_PathSegment_Segment {
	if m != nil {
		return m.Segment
	}
	return nil
}

func (m *StructMatcher_PathSegment) GetKey() string {
	if x, ok := m.GetSegment().(*StructMatcher_PathSegment_Key); ok {
		return x.Key
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*StructMatcher_PathSegment) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*StructMatcher_PathSegment_Key)(nil),
	}
}

func init() {
	proto.RegisterType((*StructMatcher)(nil), "envoy.type.matcher.v3.StructMatcher")
	proto.RegisterType((*StructMatcher_PathSegment)(nil), "envoy.type.matcher.v3.StructMatcher.PathSegment")
}

func init() { proto.RegisterFile("envoy/type/matcher/v3/struct.proto", fileDescriptor_2eb06d41ab351944) }

var fileDescriptor_2eb06d41ab351944 = []byte{
	// 335 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0x3f, 0x4b, 0xc3, 0x40,
	0x18, 0xc6, 0xbd, 0x4b, 0xff, 0x79, 0x41, 0x29, 0x01, 0xb1, 0x54, 0x94, 0xb4, 0x1d, 0xec, 0x20,
	0x77, 0xd2, 0x6c, 0xc5, 0xe9, 0x5c, 0x5c, 0x94, 0xd2, 0x82, 0xfb, 0xd9, 0x1e, 0x6d, 0xb0, 0xbd,
	0x0b, 0xc9, 0x9b, 0x60, 0x36, 0x47, 0x71, 0x74, 0xf4, 0xa3, 0x38, 0xb8, 0x09, 0xae, 0x7e, 0x1b,
	0xe9, 0x24, 0xc9, 0x45, 0x50, 0x4c, 0x71, 0x3b, 0xee, 0x7d, 0x9e, 0xdf, 0xfb, 0x3c, 0x77, 0xa4,
	0x2b, 0x55, 0xa2, 0x53, 0x06, 0x69, 0x20, 0xd9, 0x4a, 0xc0, 0x74, 0x21, 0x43, 0x96, 0x78, 0x2c,
	0x82, 0x30, 0x9e, 0x02, 0x0d, 0x42, 0x0d, 0xda, 0xd9, 0xcb, 0x35, 0x34, 0xd3, 0xd0, 0x42, 0x43,
	0x13, 0xaf, 0xdd, 0x29, 0xb7, 0x26, 0x62, 0x19, 0x4b, 0xe3, 0x6c, 0x1f, 0xc6, 0xb3, 0x40, 0x30,
	0xa1, 0x94, 0x06, 0x01, 0xbe, 0x56, 0x11, 0x8b, 0x40, 0x40, 0x1c, 0x15, 0xe3, 0xce, 0x9f, 0x71,
	0x22, 0xc3, 0xc8, 0xd7, 0xca, 0x57, 0xf3, 0x42, 0xb2, 0x9f, 0x88, 0xa5, 0x3f, 0x13, 0x20, 0xd9,
	0xf7, 0xc1, 0x0c, 0xba, 0xaf, 0x98, 0xec, 0x4c, 0xf2, 0x94, 0x97, 0x66, 0xb7, 0x73, 0x45, 0x2a,
	0x81, 0x80, 0x45, 0x0b, 0xbb, 0x56, 0xdf, 0x1e, 0x9c, 0xd2, 0xd2, 0xd4, 0xf4, 0x97, 0x87, 0x8e,
	0x04, 0x2c, 0x26, 0x72, 0xbe, 0x92, 0x0a, 0x78, 0x63, 0xcd, 0xab, 0x4f, 0x08, 0x37, 0xd0, 0x38,
	0xe7, 0x38, 0xe7, 0xa4, 0x9a, 0x77, 0x69, 0x59, 0x2e, 0xea, 0xdb, 0x83, 0xde, 0x06, 0xe0, 0x75,
	0xa6, 0x29, 0x78, 0x39, 0xe3, 0x11, 0xe1, 0x26, 0x1a, 0x1b, 0x6f, 0x5b, 0x13, 0xfb, 0xc7, 0x0e,
	0xe7, 0x80, 0x58, 0xb7, 0x32, 0x6d, 0x21, 0x17, 0xf5, 0xb7, 0x79, 0x7d, 0xcd, 0x2b, 0x21, 0x76,
	0xd1, 0xc5, 0xd6, 0x38, 0xbb, 0x1d, 0x7a, 0xcf, 0x6f, 0x0f, 0x47, 0x94, 0x9c, 0x94, 0xec, 0xd9,
	0x9c, 0x7a, 0x97, 0xd4, 0xa3, 0x02, 0x6e, 0x7d, 0x72, 0x34, 0x3c, 0xce, 0x20, 0x5d, 0xe2, 0xfe,
	0x07, 0xe1, 0x67, 0x2f, 0xf7, 0xef, 0x1f, 0x35, 0xdc, 0xc4, 0xa4, 0xe7, 0x6b, 0xd3, 0x2d, 0x08,
	0xf5, 0x5d, 0x5a, 0x5e, 0x93, 0xdb, 0xc6, 0x3d, 0xca, 0x1e, 0x7f, 0x84, 0x6e, 0x6a, 0xf9, 0x2f,
	0x78, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xb3, 0xae, 0xa1, 0xab, 0x40, 0x02, 0x00, 0x00,
}