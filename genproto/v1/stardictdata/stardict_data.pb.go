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
	Name         string                     `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Code         string                     `protobuf:"bytes,9,opt,name=code" json:"code,omitempty"`
	Version      string                     `protobuf:"bytes,2,opt,name=version" json:"version,omitempty"`
	Description  string                     `protobuf:"bytes,3,opt,name=description" json:"description,omitempty"`
	Author       string                     `protobuf:"bytes,4,opt,name=author" json:"author,omitempty"`
	Email        string                     `protobuf:"bytes,5,opt,name=email" json:"email,omitempty"`
	Website      string                     `protobuf:"bytes,6,opt,name=website" json:"website,omitempty"`
	Date         *google_protobuf.Timestamp `protobuf:"bytes,7,opt,name=date" json:"date,omitempty"`
	Type         string                     `protobuf:"bytes,8,opt,name=type" json:"type,omitempty"`
	WordCount    int32                      `protobuf:"varint,10,opt,name=word_count,json=wordCount" json:"word_count,omitempty"`
	SynonymCount int32                      `protobuf:"varint,11,opt,name=synonym_count,json=synonymCount" json:"synonym_count,omitempty"`
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

func (m *Info) GetCode() string {
	if m != nil {
		return m.Code
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

func (m *Info) GetWordCount() int32 {
	if m != nil {
		return m.WordCount
	}
	return 0
}

func (m *Info) GetSynonymCount() int32 {
	if m != nil {
		return m.SynonymCount
	}
	return 0
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
	// 551 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x93, 0x51, 0x6f, 0xd3, 0x40,
	0x0c, 0xc7, 0x49, 0x9b, 0x34, 0xad, 0x33, 0x50, 0x66, 0x10, 0x9c, 0x2a, 0x01, 0x55, 0x79, 0x99,
	0x98, 0x94, 0x68, 0x45, 0xe2, 0x85, 0xa7, 0xae, 0x44, 0x2c, 0x12, 0xac, 0x25, 0xcb, 0x18, 0xf0,
	0x32, 0xa5, 0xc9, 0xb5, 0x8b, 0xb4, 0xdc, 0x55, 0xc9, 0xb5, 0xa5, 0xdf, 0x81, 0x4f, 0xc1, 0xc7,
	0xe0, 0xd3, 0xa1, 0xbb, 0x5c, 0xa6, 0x22, 0xc6, 0x9b, 0xfd, 0xf7, 0xcf, 0x76, 0x6c, 0xe7, 0xe0,
	0x71, 0x25, 0x92, 0x32, 0xcb, 0x53, 0x71, 0x9d, 0x25, 0x22, 0xf1, 0x56, 0x25, 0x17, 0x1c, 0x0f,
	0xb7, 0x94, 0xd1, 0xd2, 0x6b, 0x42, 0xde, 0xe6, 0xa4, 0xff, 0x72, 0xc9, 0xf9, 0xf2, 0x96, 0xfa,
	0x0a, 0x98, 0xaf, 0x17, 0xbe, 0xc8, 0x0b, 0x5a, 0x89, 0xa4, 0x58, 0xd5, 0x39, 0x43, 0x0e, 0x07,
	0x17, 0x9a, 0x7f, 0x9f, 0x88, 0x04, 0x8f, 0xc1, 0xcc, 0xd9, 0x82, 0x13, 0x63, 0x60, 0x1c, 0x39,
	0xa3, 0x67, 0xde, 0x3f, 0x25, 0xbd, 0x90, 0x2d, 0x78, 0xa4, 0x20, 0x1c, 0x81, 0x4d, 0x99, 0x28,
	0x73, 0x5a, 0x91, 0xd6, 0xa0, 0x7d, 0xe4, 0x8c, 0xc8, 0x3d, 0x7c, 0xc0, 0x44, 0xb9, 0x8b, 0x1a,
	0x70, 0xf8, 0xbb, 0x05, 0xa6, 0x2c, 0x81, 0x08, 0x26, 0x4b, 0x0a, 0xaa, 0x3a, 0xf5, 0x22, 0x65,
	0x4b, 0x2d, 0xe5, 0x19, 0x25, 0xbd, 0x5a, 0x93, 0x36, 0x12, 0xb0, 0x37, 0xb4, 0xac, 0x72, 0xce,
	0x48, 0x4b, 0xc9, 0x8d, 0x8b, 0x03, 0x70, 0x32, 0x5a, 0xa5, 0x65, 0xbe, 0x12, 0x32, 0xda, 0x56,
	0xd1, 0x7d, 0x09, 0x9f, 0x42, 0x27, 0x59, 0x8b, 0x1b, 0x5e, 0x12, 0x53, 0x05, 0xb5, 0x87, 0x4f,
	0xc0, 0xa2, 0x45, 0x92, 0xdf, 0x12, 0x4b, 0xc9, 0xb5, 0x23, 0x3b, 0x6d, 0xe9, 0xbc, 0xca, 0x05,
	0x25, 0x9d, 0xba, 0x93, 0x76, 0xd1, 0x03, 0x33, 0x4b, 0x04, 0x25, 0xb6, 0xda, 0x4a, 0xdf, 0xab,
	0xb7, 0xea, 0x35, 0x5b, 0xf5, 0xe2, 0x66, 0xab, 0x91, 0xe2, 0xe4, 0x1c, 0x62, 0xb7, 0xa2, 0xa4,
	0x5b, 0xcf, 0x21, 0x6d, 0x7c, 0x0e, 0xb0, 0xe5, 0x65, 0x76, 0x9d, 0xf2, 0x35, 0x13, 0x04, 0x06,
	0xc6, 0x91, 0x15, 0xf5, 0xa4, 0x32, 0x91, 0x02, 0xbe, 0x82, 0x87, 0xd5, 0x8e, 0x71, 0xb6, 0x2b,
	0x34, 0xe1, 0x28, 0xe2, 0x40, 0x8b, 0x0a, 0x1a, 0x72, 0xb0, 0xd4, 0x3a, 0x65, 0x03, 0x99, 0xda,
	0x2c, 0x4f, 0xda, 0xd8, 0x87, 0xae, 0x86, 0xeb, 0x73, 0xf4, 0xa2, 0x3b, 0x1f, 0xdf, 0x42, 0x37,
	0xe5, 0x4c, 0x50, 0x26, 0x2a, 0xd2, 0x56, 0xa7, 0xea, 0xdf, 0x73, 0xaa, 0x49, 0x8d, 0x44, 0x77,
	0xec, 0xf0, 0x33, 0xd8, 0x5a, 0xc4, 0x91, 0x9e, 0x49, 0xb6, 0x7c, 0x34, 0x7a, 0xf1, 0xff, 0xf4,
	0x78, 0xb7, 0xa2, 0x7a, 0x66, 0xb9, 0x07, 0xfa, 0x43, 0xe8, 0xc3, 0x29, 0xfb, 0x75, 0x0e, 0xce,
	0x1e, 0x88, 0x5d, 0x30, 0xe3, 0xe0, 0x6b, 0xec, 0x3e, 0x40, 0x80, 0xce, 0xb7, 0xf0, 0x34, 0x1c,
	0x4f, 0x5d, 0x03, 0x1d, 0xb0, 0x67, 0x67, 0xd3, 0xf3, 0x20, 0x0e, 0xdd, 0x16, 0xf6, 0xc0, 0x9a,
	0x8d, 0xcf, 0x3f, 0x4c, 0xdd, 0xb6, 0xa4, 0xcf, 0xe2, 0x4f, 0x1f, 0x5d, 0x13, 0x0f, 0xa0, 0x1b,
	0x05, 0x17, 0xd3, 0xcb, 0x68, 0x12, 0xb8, 0x16, 0xda, 0xd0, 0xbe, 0x1a, 0x7f, 0x71, 0x3b, 0x2a,
	0x31, 0x9c, 0xc4, 0x97, 0x51, 0xe0, 0xda, 0xa7, 0x3f, 0x0d, 0x38, 0x4e, 0x79, 0xe1, 0x2d, 0x73,
	0x71, 0xb3, 0x9e, 0xd7, 0x5f, 0x5c, 0xd0, 0xbf, 0xbe, 0xb9, 0xb1, 0xe5, 0x33, 0x3a, 0x3d, 0xdc,
	0x7f, 0x0a, 0x33, 0x79, 0xdc, 0x99, 0xf1, 0x7d, 0xac, 0xb3, 0x53, 0x5e, 0xf8, 0xba, 0x82, 0xdf,
	0x64, 0xf9, 0x4b, 0xca, 0xd4, 0x4f, 0xe0, 0x6f, 0x4e, 0xfc, 0xfd, 0x52, 0xef, 0xf6, 0x9d, 0x5f,
	0x2d, 0xf3, 0x6a, 0x3c, 0x0b, 0xe7, 0x1d, 0x85, 0xbe, 0xf9, 0x13, 0x00, 0x00, 0xff, 0xff, 0x80,
	0x0a, 0x17, 0x5d, 0xbd, 0x03, 0x00, 0x00,
}
