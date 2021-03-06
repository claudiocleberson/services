// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/posts.proto

package posts

import (
	fmt "fmt"
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

type Post struct {
	Id                   string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title                string            `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Slug                 string            `protobuf:"bytes,3,opt,name=slug,proto3" json:"slug,omitempty"`
	Content              string            `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
	Created              int64             `protobuf:"varint,5,opt,name=created,proto3" json:"created,omitempty"`
	Updated              int64             `protobuf:"varint,6,opt,name=updated,proto3" json:"updated,omitempty"`
	Author               string            `protobuf:"bytes,7,opt,name=author,proto3" json:"author,omitempty"`
	Tags                 []string          `protobuf:"bytes,8,rep,name=tags,proto3" json:"tags,omitempty"`
	Metadata             map[string]string `protobuf:"bytes,9,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Image                string            `protobuf:"bytes,19,opt,name=image,proto3" json:"image,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Post) Reset()         { *m = Post{} }
func (m *Post) String() string { return proto.CompactTextString(m) }
func (*Post) ProtoMessage()    {}
func (*Post) Descriptor() ([]byte, []int) {
	return fileDescriptor_e93dc7d934d9dc10, []int{0}
}

func (m *Post) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Post.Unmarshal(m, b)
}
func (m *Post) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Post.Marshal(b, m, deterministic)
}
func (m *Post) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Post.Merge(m, src)
}
func (m *Post) XXX_Size() int {
	return xxx_messageInfo_Post.Size(m)
}
func (m *Post) XXX_DiscardUnknown() {
	xxx_messageInfo_Post.DiscardUnknown(m)
}

var xxx_messageInfo_Post proto.InternalMessageInfo

func (m *Post) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Post) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Post) GetSlug() string {
	if m != nil {
		return m.Slug
	}
	return ""
}

func (m *Post) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *Post) GetCreated() int64 {
	if m != nil {
		return m.Created
	}
	return 0
}

func (m *Post) GetUpdated() int64 {
	if m != nil {
		return m.Updated
	}
	return 0
}

func (m *Post) GetAuthor() string {
	if m != nil {
		return m.Author
	}
	return ""
}

func (m *Post) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *Post) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Post) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

type IndexRequest struct {
	Limit                int64    `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset               int64    `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IndexRequest) Reset()         { *m = IndexRequest{} }
func (m *IndexRequest) String() string { return proto.CompactTextString(m) }
func (*IndexRequest) ProtoMessage()    {}
func (*IndexRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e93dc7d934d9dc10, []int{1}
}

func (m *IndexRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IndexRequest.Unmarshal(m, b)
}
func (m *IndexRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IndexRequest.Marshal(b, m, deterministic)
}
func (m *IndexRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IndexRequest.Merge(m, src)
}
func (m *IndexRequest) XXX_Size() int {
	return xxx_messageInfo_IndexRequest.Size(m)
}
func (m *IndexRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_IndexRequest.DiscardUnknown(m)
}

var xxx_messageInfo_IndexRequest proto.InternalMessageInfo

func (m *IndexRequest) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *IndexRequest) GetOffset() int64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

type IndexResponse struct {
	Posts                []*Post  `protobuf:"bytes,1,rep,name=posts,proto3" json:"posts,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IndexResponse) Reset()         { *m = IndexResponse{} }
func (m *IndexResponse) String() string { return proto.CompactTextString(m) }
func (*IndexResponse) ProtoMessage()    {}
func (*IndexResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e93dc7d934d9dc10, []int{2}
}

func (m *IndexResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IndexResponse.Unmarshal(m, b)
}
func (m *IndexResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IndexResponse.Marshal(b, m, deterministic)
}
func (m *IndexResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IndexResponse.Merge(m, src)
}
func (m *IndexResponse) XXX_Size() int {
	return xxx_messageInfo_IndexResponse.Size(m)
}
func (m *IndexResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_IndexResponse.DiscardUnknown(m)
}

var xxx_messageInfo_IndexResponse proto.InternalMessageInfo

func (m *IndexResponse) GetPosts() []*Post {
	if m != nil {
		return m.Posts
	}
	return nil
}

// Query posts. Acts as a listing when no id or slug provided.
// Gets a single post by id or slug if any of them provided.
type QueryRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Slug                 string   `protobuf:"bytes,2,opt,name=slug,proto3" json:"slug,omitempty"`
	Tag                  string   `protobuf:"bytes,3,opt,name=tag,proto3" json:"tag,omitempty"`
	Offset               int64    `protobuf:"varint,4,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit                int64    `protobuf:"varint,5,opt,name=limit,proto3" json:"limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueryRequest) Reset()         { *m = QueryRequest{} }
func (m *QueryRequest) String() string { return proto.CompactTextString(m) }
func (*QueryRequest) ProtoMessage()    {}
func (*QueryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e93dc7d934d9dc10, []int{3}
}

func (m *QueryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryRequest.Unmarshal(m, b)
}
func (m *QueryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryRequest.Marshal(b, m, deterministic)
}
func (m *QueryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryRequest.Merge(m, src)
}
func (m *QueryRequest) XXX_Size() int {
	return xxx_messageInfo_QueryRequest.Size(m)
}
func (m *QueryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryRequest proto.InternalMessageInfo

func (m *QueryRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *QueryRequest) GetSlug() string {
	if m != nil {
		return m.Slug
	}
	return ""
}

func (m *QueryRequest) GetTag() string {
	if m != nil {
		return m.Tag
	}
	return ""
}

func (m *QueryRequest) GetOffset() int64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *QueryRequest) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

type QueryResponse struct {
	Posts                []*Post  `protobuf:"bytes,1,rep,name=posts,proto3" json:"posts,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueryResponse) Reset()         { *m = QueryResponse{} }
func (m *QueryResponse) String() string { return proto.CompactTextString(m) }
func (*QueryResponse) ProtoMessage()    {}
func (*QueryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e93dc7d934d9dc10, []int{4}
}

func (m *QueryResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryResponse.Unmarshal(m, b)
}
func (m *QueryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryResponse.Marshal(b, m, deterministic)
}
func (m *QueryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryResponse.Merge(m, src)
}
func (m *QueryResponse) XXX_Size() int {
	return xxx_messageInfo_QueryResponse.Size(m)
}
func (m *QueryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryResponse proto.InternalMessageInfo

func (m *QueryResponse) GetPosts() []*Post {
	if m != nil {
		return m.Posts
	}
	return nil
}

type SaveRequest struct {
	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title     string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Slug      string `protobuf:"bytes,3,opt,name=slug,proto3" json:"slug,omitempty"`
	Content   string `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
	Timestamp int64  `protobuf:"varint,5,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// When updating a post and wanting to delete all tags,
	// send a list of tags with only one member being an empty string [""]
	Tags                 []string          `protobuf:"bytes,6,rep,name=tags,proto3" json:"tags,omitempty"`
	Metadata             map[string]string `protobuf:"bytes,7,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Image                string            `protobuf:"bytes,8,opt,name=image,proto3" json:"image,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *SaveRequest) Reset()         { *m = SaveRequest{} }
func (m *SaveRequest) String() string { return proto.CompactTextString(m) }
func (*SaveRequest) ProtoMessage()    {}
func (*SaveRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e93dc7d934d9dc10, []int{5}
}

func (m *SaveRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SaveRequest.Unmarshal(m, b)
}
func (m *SaveRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SaveRequest.Marshal(b, m, deterministic)
}
func (m *SaveRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SaveRequest.Merge(m, src)
}
func (m *SaveRequest) XXX_Size() int {
	return xxx_messageInfo_SaveRequest.Size(m)
}
func (m *SaveRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SaveRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SaveRequest proto.InternalMessageInfo

func (m *SaveRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *SaveRequest) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *SaveRequest) GetSlug() string {
	if m != nil {
		return m.Slug
	}
	return ""
}

func (m *SaveRequest) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *SaveRequest) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *SaveRequest) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *SaveRequest) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *SaveRequest) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

type SaveResponse struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SaveResponse) Reset()         { *m = SaveResponse{} }
func (m *SaveResponse) String() string { return proto.CompactTextString(m) }
func (*SaveResponse) ProtoMessage()    {}
func (*SaveResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e93dc7d934d9dc10, []int{6}
}

func (m *SaveResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SaveResponse.Unmarshal(m, b)
}
func (m *SaveResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SaveResponse.Marshal(b, m, deterministic)
}
func (m *SaveResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SaveResponse.Merge(m, src)
}
func (m *SaveResponse) XXX_Size() int {
	return xxx_messageInfo_SaveResponse.Size(m)
}
func (m *SaveResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SaveResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SaveResponse proto.InternalMessageInfo

func (m *SaveResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type DeleteRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteRequest) Reset()         { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()    {}
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e93dc7d934d9dc10, []int{7}
}

func (m *DeleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteRequest.Unmarshal(m, b)
}
func (m *DeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteRequest.Marshal(b, m, deterministic)
}
func (m *DeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteRequest.Merge(m, src)
}
func (m *DeleteRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteRequest.Size(m)
}
func (m *DeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteRequest proto.InternalMessageInfo

func (m *DeleteRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type DeleteResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteResponse) Reset()         { *m = DeleteResponse{} }
func (m *DeleteResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteResponse) ProtoMessage()    {}
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e93dc7d934d9dc10, []int{8}
}

func (m *DeleteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteResponse.Unmarshal(m, b)
}
func (m *DeleteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteResponse.Marshal(b, m, deterministic)
}
func (m *DeleteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteResponse.Merge(m, src)
}
func (m *DeleteResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteResponse.Size(m)
}
func (m *DeleteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Post)(nil), "posts.Post")
	proto.RegisterMapType((map[string]string)(nil), "posts.Post.MetadataEntry")
	proto.RegisterType((*IndexRequest)(nil), "posts.IndexRequest")
	proto.RegisterType((*IndexResponse)(nil), "posts.IndexResponse")
	proto.RegisterType((*QueryRequest)(nil), "posts.QueryRequest")
	proto.RegisterType((*QueryResponse)(nil), "posts.QueryResponse")
	proto.RegisterType((*SaveRequest)(nil), "posts.SaveRequest")
	proto.RegisterMapType((map[string]string)(nil), "posts.SaveRequest.MetadataEntry")
	proto.RegisterType((*SaveResponse)(nil), "posts.SaveResponse")
	proto.RegisterType((*DeleteRequest)(nil), "posts.DeleteRequest")
	proto.RegisterType((*DeleteResponse)(nil), "posts.DeleteResponse")
}

func init() { proto.RegisterFile("proto/posts.proto", fileDescriptor_e93dc7d934d9dc10) }

var fileDescriptor_e93dc7d934d9dc10 = []byte{
	// 496 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0x4b, 0x8e, 0xd3, 0x40,
	0x10, 0x1d, 0x7f, 0x93, 0x54, 0x92, 0x51, 0xe8, 0x0c, 0xa8, 0xb1, 0x10, 0x18, 0xaf, 0xb2, 0x0a,
	0x22, 0x80, 0x40, 0x30, 0x4b, 0x58, 0xb0, 0x40, 0x02, 0x73, 0x82, 0x06, 0xf7, 0x04, 0x0b, 0xff,
	0x70, 0x97, 0x47, 0xe4, 0x3e, 0x1c, 0x85, 0x7b, 0x70, 0x15, 0xd4, 0x1f, 0x7b, 0xda, 0x41, 0x23,
	0x66, 0x31, 0xbb, 0x7e, 0xd5, 0x7e, 0x5d, 0xef, 0x55, 0x3d, 0x19, 0xee, 0x34, 0x6d, 0x8d, 0xf5,
	0x93, 0xa6, 0x16, 0x28, 0xb6, 0xea, 0x4c, 0x02, 0x05, 0x92, 0xdf, 0x2e, 0xf8, 0x1f, 0x6b, 0x81,
	0xe4, 0x14, 0xdc, 0x3c, 0xa3, 0x4e, 0xec, 0x6c, 0x66, 0xa9, 0x9b, 0x67, 0xe4, 0x0c, 0x02, 0xcc,
	0xb1, 0xe0, 0xd4, 0x55, 0x25, 0x0d, 0x08, 0x01, 0x5f, 0x14, 0xdd, 0x9e, 0x7a, 0xaa, 0xa8, 0xce,
	0x84, 0xc2, 0xe4, 0x6b, 0x5d, 0x21, 0xaf, 0x90, 0xfa, 0xaa, 0xdc, 0x43, 0x75, 0xd3, 0x72, 0x86,
	0x3c, 0xa3, 0x41, 0xec, 0x6c, 0xbc, 0xb4, 0x87, 0xf2, 0xa6, 0x6b, 0x32, 0x75, 0x13, 0xea, 0x1b,
	0x03, 0xc9, 0x3d, 0x08, 0x59, 0x87, 0xdf, 0xea, 0x96, 0x4e, 0xd4, 0x63, 0x06, 0xc9, 0xce, 0xc8,
	0xf6, 0x82, 0x4e, 0x63, 0x4f, 0x76, 0x96, 0x67, 0xf2, 0x02, 0xa6, 0x25, 0x47, 0x96, 0x31, 0x64,
	0x74, 0x16, 0x7b, 0x9b, 0xf9, 0xee, 0xfe, 0x56, 0x7b, 0x94, 0x96, 0xb6, 0x1f, 0xcc, 0xdd, 0xbb,
	0x0a, 0xdb, 0x43, 0x3a, 0x7c, 0x2a, 0xad, 0xe5, 0x25, 0xdb, 0x73, 0xba, 0xd6, 0xd6, 0x14, 0x88,
	0xde, 0xc0, 0x72, 0x44, 0x20, 0x2b, 0xf0, 0xbe, 0xf3, 0x83, 0x19, 0x89, 0x3c, 0x4a, 0xe2, 0x25,
	0x2b, 0xba, 0x61, 0x26, 0x0a, 0xbc, 0x76, 0x5f, 0x39, 0xc9, 0x39, 0x2c, 0xde, 0x57, 0x19, 0xff,
	0x99, 0xf2, 0x1f, 0x1d, 0x17, 0x28, 0xbf, 0x2c, 0xf2, 0x32, 0x47, 0xc5, 0xf6, 0x52, 0x0d, 0xa4,
	0xb7, 0xfa, 0xe2, 0x42, 0x70, 0x54, 0x0f, 0x78, 0xa9, 0x41, 0xc9, 0x0e, 0x96, 0x86, 0x2d, 0x9a,
	0xba, 0x12, 0x9c, 0x3c, 0x06, 0xbd, 0x1e, 0xea, 0x28, 0x57, 0x73, 0xcb, 0x55, 0x6a, 0x16, 0xd7,
	0xc2, 0xe2, 0x53, 0xc7, 0xdb, 0x43, 0xdf, 0xf1, 0x78, 0x7f, 0xfd, 0xa6, 0x5c, 0x6b, 0x53, 0x2b,
	0xf0, 0x90, 0xf5, 0xcb, 0x93, 0x47, 0x4b, 0x91, 0x6f, 0x2b, 0xba, 0xd2, 0x1f, 0x58, 0xfa, 0xa5,
	0x4e, 0xd3, 0xf3, 0xe6, 0x3a, 0x7f, 0xb9, 0x30, 0xff, 0xcc, 0x2e, 0xf9, 0x75, 0x3a, 0x6f, 0x23,
	0x67, 0x0f, 0x60, 0x86, 0x79, 0xc9, 0x05, 0xb2, 0xb2, 0x31, 0x8a, 0xaf, 0x0a, 0x43, 0x72, 0x42,
	0x2b, 0x39, 0xe7, 0x56, 0x72, 0x26, 0x4a, 0x7b, 0x6c, 0xb4, 0x5b, 0x5a, 0xff, 0x1f, 0xa0, 0xe9,
	0xad, 0x05, 0xe8, 0x21, 0x2c, 0x74, 0x67, 0x33, 0xd9, 0xa3, 0x31, 0x25, 0x8f, 0x60, 0xf9, 0x96,
	0x17, 0x1c, 0xaf, 0x9b, 0x63, 0xb2, 0x82, 0xd3, 0xfe, 0x03, 0xfd, 0xc4, 0xee, 0x8f, 0x03, 0x81,
	0xdc, 0x84, 0x20, 0xcf, 0x21, 0x50, 0xf9, 0x22, 0x6b, 0x63, 0xd2, 0xce, 0x6a, 0x74, 0x36, 0x2e,
	0x6a, 0x76, 0x72, 0x22, 0x59, 0x6a, 0xdb, 0x03, 0xcb, 0xce, 0xdb, 0xc0, 0x1a, 0x05, 0x22, 0x39,
	0x21, 0x4f, 0xc1, 0x97, 0x46, 0x08, 0xf9, 0x77, 0x9e, 0xd1, 0x7a, 0x54, 0x1b, 0x28, 0x2f, 0x21,
	0xd4, 0xd2, 0x49, 0xff, 0xe8, 0xc8, 0x6a, 0x74, 0xf7, 0xa8, 0xda, 0x13, 0xbf, 0x84, 0xea, 0x57,
	0xf6, 0xec, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa7, 0xee, 0x3f, 0xc4, 0xdf, 0x04, 0x00, 0x00,
}
