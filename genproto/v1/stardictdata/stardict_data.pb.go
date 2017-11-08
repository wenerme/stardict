// Code generated by protoc-gen-go. DO NOT EDIT.
// source: stardict_data.proto

/*
Package stardictdata is a generated protocol buffer package.

It is generated from these files:
	stardict_data.proto

It has these top-level messages:
	StardictData
	Info
	Entry
	Content
*/
package stardictdata

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ContentType int32

const (
	ContentType_TEXT     ContentType = 0
	ContentType_YIBIAO   ContentType = 1
	ContentType_PHONETI  ContentType = 2
	ContentType_PANGO    ContentType = 3
	ContentType_HTML     ContentType = 4
	ContentType_RESOURCE ContentType = 5
	ContentType_WAV      ContentType = 6
	ContentType_PICTURE  ContentType = 7
)

var ContentType_name = map[int32]string{
	0: "TEXT",
	1: "YIBIAO",
	2: "PHONETI",
	3: "PANGO",
	4: "HTML",
	5: "RESOURCE",
	6: "WAV",
	7: "PICTURE",
}
var ContentType_value = map[string]int32{
	"TEXT":     0,
	"YIBIAO":   1,
	"PHONETI":  2,
	"PANGO":    3,
	"HTML":     4,
	"RESOURCE": 5,
	"WAV":      6,
	"PICTURE":  7,
}

func (x ContentType) String() string {
	return proto.EnumName(ContentType_name, int32(x))
}
func (ContentType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type StardictData struct {
	Info    *Info    `protobuf:"bytes,1,opt,name=info" json:"info,omitempty"`
	Entries []*Entry `protobuf:"bytes,2,rep,name=entries" json:"entries,omitempty"`
}

func (m *StardictData) Reset()                    { *m = StardictData{} }
func (m *StardictData) String() string            { return proto.CompactTextString(m) }
func (*StardictData) ProtoMessage()               {}
func (*StardictData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *StardictData) GetInfo() *Info {
	if m != nil {
		return m.Info
	}
	return nil
}

func (m *StardictData) GetEntries() []*Entry {
	if m != nil {
		return m.Entries
	}
	return nil
}

type Info struct {
	Name        string                     `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Version     string                     `protobuf:"bytes,2,opt,name=version" json:"version,omitempty"`
	Description string                     `protobuf:"bytes,3,opt,name=description" json:"description,omitempty"`
	Author      string                     `protobuf:"bytes,4,opt,name=author" json:"author,omitempty"`
	Email       string                     `protobuf:"bytes,5,opt,name=email" json:"email,omitempty"`
	Website     string                     `protobuf:"bytes,6,opt,name=website" json:"website,omitempty"`
	Date        *google_protobuf.Timestamp `protobuf:"bytes,7,opt,name=date" json:"date,omitempty"`
	Type        string                     `protobuf:"bytes,8,opt,name=type" json:"type,omitempty"`
}

func (m *Info) Reset()                    { *m = Info{} }
func (m *Info) String() string            { return proto.CompactTextString(m) }
func (*Info) ProtoMessage()               {}
func (*Info) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Info) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Info) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *Info) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Info) GetAuthor() string {
	if m != nil {
		return m.Author
	}
	return ""
}

func (m *Info) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *Info) GetWebsite() string {
	if m != nil {
		return m.Website
	}
	return ""
}

func (m *Info) GetDate() *google_protobuf.Timestamp {
	if m != nil {
		return m.Date
	}
	return nil
}

func (m *Info) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

type Entry struct {
	Word     string     `protobuf:"bytes,1,opt,name=word" json:"word,omitempty"`
	Synonyms []string   `protobuf:"bytes,2,rep,name=synonyms" json:"synonyms,omitempty"`
	Contents []*Content `protobuf:"bytes,3,rep,name=contents" json:"contents,omitempty"`
}

func (m *Entry) Reset()                    { *m = Entry{} }
func (m *Entry) String() string            { return proto.CompactTextString(m) }
func (*Entry) ProtoMessage()               {}
func (*Entry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Entry) GetWord() string {
	if m != nil {
		return m.Word
	}
	return ""
}

func (m *Entry) GetSynonyms() []string {
	if m != nil {
		return m.Synonyms
	}
	return nil
}

func (m *Entry) GetContents() []*Content {
	if m != nil {
		return m.Contents
	}
	return nil
}

type Content struct {
	Type ContentType `protobuf:"varint,1,opt,name=type,enum=wener.stardict.v1.ContentType" json:"type,omitempty"`
	Text string      `protobuf:"bytes,2,opt,name=text" json:"text,omitempty"`
}

func (m *Content) Reset()                    { *m = Content{} }
func (m *Content) String() string            { return proto.CompactTextString(m) }
func (*Content) ProtoMessage()               {}
func (*Content) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Content) GetType() ContentType {
	if m != nil {
		return m.Type
	}
	return ContentType_TEXT
}

func (m *Content) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func init() {
	proto.RegisterType((*StardictData)(nil), "wener.stardict.v1.StardictData")
	proto.RegisterType((*Info)(nil), "wener.stardict.v1.Info")
	proto.RegisterType((*Entry)(nil), "wener.stardict.v1.Entry")
	proto.RegisterType((*Content)(nil), "wener.stardict.v1.Content")
	proto.RegisterEnum("wener.stardict.v1.ContentType", ContentType_name, ContentType_value)
}

func init() { proto.RegisterFile("stardict_data.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 508 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0xdd, 0x6e, 0xd3, 0x30,
	0x14, 0xc7, 0x49, 0x9b, 0x36, 0xdd, 0xe9, 0x84, 0x32, 0x83, 0xc0, 0xea, 0x05, 0x54, 0xbd, 0x9a,
	0x98, 0xe4, 0x68, 0x45, 0xe2, 0x86, 0xab, 0xb6, 0x44, 0x2c, 0x12, 0xac, 0x25, 0xcb, 0x18, 0x70,
	0x83, 0xdc, 0xc4, 0xed, 0x2c, 0x2d, 0x76, 0x94, 0xb8, 0x2d, 0x79, 0x07, 0x9e, 0x82, 0x47, 0xe3,
	0x49, 0x90, 0x9d, 0x64, 0x0a, 0x62, 0xdc, 0x9d, 0x8f, 0xdf, 0xf9, 0xfa, 0xdb, 0xf0, 0xa4, 0x50,
	0x34, 0x4f, 0x78, 0xac, 0xbe, 0x27, 0x54, 0x51, 0x92, 0xe5, 0x52, 0x49, 0x74, 0x72, 0x60, 0x82,
	0xe5, 0xa4, 0x49, 0x91, 0xfd, 0xf9, 0xe8, 0xe5, 0x56, 0xca, 0xed, 0x1d, 0xf3, 0x0c, 0xb0, 0xde,
	0x6d, 0x3c, 0xc5, 0x53, 0x56, 0x28, 0x9a, 0x66, 0x55, 0xcd, 0x44, 0xc2, 0xf1, 0x55, 0xcd, 0xbf,
	0xa3, 0x8a, 0xa2, 0x33, 0xb0, 0xb9, 0xd8, 0x48, 0x6c, 0x8d, 0xad, 0xd3, 0xe1, 0xf4, 0x39, 0xf9,
	0xa7, 0x25, 0x09, 0xc4, 0x46, 0x86, 0x06, 0x42, 0x53, 0x70, 0x98, 0x50, 0x39, 0x67, 0x05, 0xee,
	0x8c, 0xbb, 0xa7, 0xc3, 0x29, 0x7e, 0x80, 0xf7, 0x85, 0xca, 0xcb, 0xb0, 0x01, 0x27, 0xbf, 0x2d,
	0xb0, 0x75, 0x0b, 0x84, 0xc0, 0x16, 0x34, 0x65, 0x66, 0xd2, 0x51, 0x68, 0x6c, 0x84, 0xc1, 0xd9,
	0xb3, 0xbc, 0xe0, 0x52, 0xe0, 0x8e, 0x09, 0x37, 0x2e, 0x1a, 0xc3, 0x30, 0x61, 0x45, 0x9c, 0xf3,
	0x4c, 0xe9, 0x6c, 0xd7, 0x64, 0xdb, 0x21, 0xf4, 0x0c, 0xfa, 0x74, 0xa7, 0x6e, 0x65, 0x8e, 0x6d,
	0x93, 0xac, 0x3d, 0xf4, 0x14, 0x7a, 0x2c, 0xa5, 0xfc, 0x0e, 0xf7, 0x4c, 0xb8, 0x72, 0xf4, 0xa4,
	0x03, 0x5b, 0x17, 0x5c, 0x31, 0xdc, 0xaf, 0x26, 0xd5, 0x2e, 0x22, 0x60, 0x27, 0x54, 0x31, 0xec,
	0x18, 0x05, 0x46, 0xa4, 0x52, 0x90, 0x34, 0x0a, 0x92, 0xa8, 0x51, 0x30, 0x34, 0x9c, 0xbe, 0x43,
	0x95, 0x19, 0xc3, 0x83, 0xea, 0x0e, 0x6d, 0x4f, 0x24, 0xf4, 0xcc, 0xd9, 0x3a, 0x79, 0x90, 0x79,
	0xd2, 0x1c, 0xa9, 0x6d, 0x34, 0x82, 0x41, 0x51, 0x0a, 0x29, 0xca, 0xb4, 0x92, 0xed, 0x28, 0xbc,
	0xf7, 0xd1, 0x1b, 0x18, 0xc4, 0x52, 0x28, 0x26, 0x54, 0x81, 0xbb, 0x46, 0xd2, 0xd1, 0x03, 0x92,
	0x2e, 0x2a, 0x24, 0xbc, 0x67, 0x27, 0x9f, 0xc0, 0xa9, 0x83, 0x68, 0x5a, 0xef, 0xa3, 0x47, 0x3e,
	0x9e, 0xbe, 0xf8, 0x7f, 0x79, 0x54, 0x66, 0xac, 0xda, 0xd7, 0xdc, 0xc0, 0x7e, 0xa8, 0x5a, 0x74,
	0x63, 0xbf, 0xe2, 0x30, 0x6c, 0x81, 0x68, 0x00, 0x76, 0xe4, 0x7f, 0x89, 0xdc, 0x47, 0x08, 0xa0,
	0xff, 0x35, 0x98, 0x07, 0xb3, 0xa5, 0x6b, 0xa1, 0x21, 0x38, 0xab, 0x8b, 0xe5, 0xa5, 0x1f, 0x05,
	0x6e, 0x07, 0x1d, 0x41, 0x6f, 0x35, 0xbb, 0x7c, 0xbf, 0x74, 0xbb, 0x9a, 0xbe, 0x88, 0x3e, 0x7e,
	0x70, 0x6d, 0x74, 0x0c, 0x83, 0xd0, 0xbf, 0x5a, 0x5e, 0x87, 0x0b, 0xdf, 0xed, 0x21, 0x07, 0xba,
	0x37, 0xb3, 0xcf, 0x6e, 0xdf, 0x14, 0x06, 0x8b, 0xe8, 0x3a, 0xf4, 0x5d, 0x67, 0xfe, 0xd3, 0x82,
	0xb3, 0x58, 0xa6, 0x64, 0xcb, 0xd5, 0xed, 0x6e, 0x5d, 0x6d, 0x9c, 0xb2, 0xbf, 0x76, 0x6e, 0x6c,
	0xfd, 0xdd, 0xe7, 0x27, 0xed, 0x2f, 0xbb, 0xd2, 0x0f, 0xb3, 0xb2, 0xbe, 0xcd, 0xea, 0xea, 0x58,
	0xa6, 0x5e, 0xdd, 0xc1, 0x6b, 0xaa, 0xbc, 0x2d, 0x13, 0xe6, 0x01, 0xbd, 0xfd, 0xb9, 0xd7, 0x6e,
	0xf5, 0xb6, 0xed, 0xfc, 0xea, 0xd8, 0x37, 0xb3, 0x55, 0xb0, 0xee, 0x1b, 0xf4, 0xf5, 0x9f, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x44, 0x84, 0x03, 0xd1, 0x65, 0x03, 0x00, 0x00,
}
