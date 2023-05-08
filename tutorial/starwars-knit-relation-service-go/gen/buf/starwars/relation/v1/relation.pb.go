// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: buf/starwars/relation/v1/relation.proto

package relationv1

import (
	_ "buf.build/gen/go/bufbuild/knit/protocolbuffers/go/buf/knit/v1alpha1"
	v1 "github.com/bufbuild/knit/tutorial/starwars-knit-relation-service-go/gen/buf/starwars/film/v1"
	v11 "github.com/bufbuild/knit/tutorial/starwars-knit-relation-service-go/gen/buf/starwars/quote/v1"
	v12 "github.com/bufbuild/knit/tutorial/starwars-knit-relation-service-go/gen/buf/starwars/starship/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetFilmStarshipsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bases []*v1.Film `protobuf:"bytes,1,rep,name=bases,proto3" json:"bases,omitempty"`
}

func (x *GetFilmStarshipsRequest) Reset() {
	*x = GetFilmStarshipsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buf_starwars_relation_v1_relation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFilmStarshipsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFilmStarshipsRequest) ProtoMessage() {}

func (x *GetFilmStarshipsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_buf_starwars_relation_v1_relation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFilmStarshipsRequest.ProtoReflect.Descriptor instead.
func (*GetFilmStarshipsRequest) Descriptor() ([]byte, []int) {
	return file_buf_starwars_relation_v1_relation_proto_rawDescGZIP(), []int{0}
}

func (x *GetFilmStarshipsRequest) GetBases() []*v1.Film {
	if x != nil {
		return x.Bases
	}
	return nil
}

type GetFilmStarshipsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Values []*GetFilmStarshipsResponse_Value `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty"`
}

func (x *GetFilmStarshipsResponse) Reset() {
	*x = GetFilmStarshipsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buf_starwars_relation_v1_relation_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFilmStarshipsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFilmStarshipsResponse) ProtoMessage() {}

func (x *GetFilmStarshipsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_buf_starwars_relation_v1_relation_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFilmStarshipsResponse.ProtoReflect.Descriptor instead.
func (*GetFilmStarshipsResponse) Descriptor() ([]byte, []int) {
	return file_buf_starwars_relation_v1_relation_proto_rawDescGZIP(), []int{1}
}

func (x *GetFilmStarshipsResponse) GetValues() []*GetFilmStarshipsResponse_Value {
	if x != nil {
		return x.Values
	}
	return nil
}

type GetQuoteFilmRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bases []*v11.Quote `protobuf:"bytes,1,rep,name=bases,proto3" json:"bases,omitempty"`
}

func (x *GetQuoteFilmRequest) Reset() {
	*x = GetQuoteFilmRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buf_starwars_relation_v1_relation_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetQuoteFilmRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetQuoteFilmRequest) ProtoMessage() {}

func (x *GetQuoteFilmRequest) ProtoReflect() protoreflect.Message {
	mi := &file_buf_starwars_relation_v1_relation_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetQuoteFilmRequest.ProtoReflect.Descriptor instead.
func (*GetQuoteFilmRequest) Descriptor() ([]byte, []int) {
	return file_buf_starwars_relation_v1_relation_proto_rawDescGZIP(), []int{2}
}

func (x *GetQuoteFilmRequest) GetBases() []*v11.Quote {
	if x != nil {
		return x.Bases
	}
	return nil
}

type GetQuoteFilmResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Values []*GetQuoteFilmResponse_Value `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty"`
}

func (x *GetQuoteFilmResponse) Reset() {
	*x = GetQuoteFilmResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buf_starwars_relation_v1_relation_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetQuoteFilmResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetQuoteFilmResponse) ProtoMessage() {}

func (x *GetQuoteFilmResponse) ProtoReflect() protoreflect.Message {
	mi := &file_buf_starwars_relation_v1_relation_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetQuoteFilmResponse.ProtoReflect.Descriptor instead.
func (*GetQuoteFilmResponse) Descriptor() ([]byte, []int) {
	return file_buf_starwars_relation_v1_relation_proto_rawDescGZIP(), []int{3}
}

func (x *GetQuoteFilmResponse) GetValues() []*GetQuoteFilmResponse_Value {
	if x != nil {
		return x.Values
	}
	return nil
}

type GetFilmStarshipsResponse_Value struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Starships []*v12.Starship `protobuf:"bytes,1,rep,name=starships,proto3" json:"starships,omitempty"`
}

func (x *GetFilmStarshipsResponse_Value) Reset() {
	*x = GetFilmStarshipsResponse_Value{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buf_starwars_relation_v1_relation_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFilmStarshipsResponse_Value) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFilmStarshipsResponse_Value) ProtoMessage() {}

func (x *GetFilmStarshipsResponse_Value) ProtoReflect() protoreflect.Message {
	mi := &file_buf_starwars_relation_v1_relation_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFilmStarshipsResponse_Value.ProtoReflect.Descriptor instead.
func (*GetFilmStarshipsResponse_Value) Descriptor() ([]byte, []int) {
	return file_buf_starwars_relation_v1_relation_proto_rawDescGZIP(), []int{1, 0}
}

func (x *GetFilmStarshipsResponse_Value) GetStarships() []*v12.Starship {
	if x != nil {
		return x.Starships
	}
	return nil
}

type GetQuoteFilmResponse_Value struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Film *v1.Film `protobuf:"bytes,1,opt,name=film,proto3" json:"film,omitempty"`
}

func (x *GetQuoteFilmResponse_Value) Reset() {
	*x = GetQuoteFilmResponse_Value{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buf_starwars_relation_v1_relation_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetQuoteFilmResponse_Value) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetQuoteFilmResponse_Value) ProtoMessage() {}

func (x *GetQuoteFilmResponse_Value) ProtoReflect() protoreflect.Message {
	mi := &file_buf_starwars_relation_v1_relation_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetQuoteFilmResponse_Value.ProtoReflect.Descriptor instead.
func (*GetQuoteFilmResponse_Value) Descriptor() ([]byte, []int) {
	return file_buf_starwars_relation_v1_relation_proto_rawDescGZIP(), []int{3, 0}
}

func (x *GetQuoteFilmResponse_Value) GetFilm() *v1.Film {
	if x != nil {
		return x.Film
	}
	return nil
}

var File_buf_starwars_relation_v1_relation_proto protoreflect.FileDescriptor

var file_buf_starwars_relation_v1_relation_proto_rawDesc = []byte{
	0x0a, 0x27, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x61, 0x72, 0x77, 0x61, 0x72, 0x73, 0x2f, 0x72,
	0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x62, 0x75, 0x66, 0x2e, 0x73,
	0x74, 0x61, 0x72, 0x77, 0x61, 0x72, 0x73, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x61, 0x72, 0x77, 0x61, 0x72,
	0x73, 0x2f, 0x66, 0x69, 0x6c, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x69, 0x6c, 0x6d, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x61, 0x72, 0x77, 0x61,
	0x72, 0x73, 0x2f, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x71, 0x75, 0x6f, 0x74,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x27, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x61,
	0x72, 0x77, 0x61, 0x72, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x72, 0x73, 0x68, 0x69, 0x70, 0x2f, 0x76,
	0x31, 0x2f, 0x73, 0x74, 0x61, 0x72, 0x73, 0x68, 0x69, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1f, 0x62, 0x75, 0x66, 0x2f, 0x6b, 0x6e, 0x69, 0x74, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x4b, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x6d, 0x53, 0x74, 0x61, 0x72,
	0x73, 0x68, 0x69, 0x70, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x30, 0x0a, 0x05,
	0x62, 0x61, 0x73, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x62, 0x75,
	0x66, 0x2e, 0x73, 0x74, 0x61, 0x72, 0x77, 0x61, 0x72, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x6d, 0x2e,
	0x76, 0x31, 0x2e, 0x46, 0x69, 0x6c, 0x6d, 0x52, 0x05, 0x62, 0x61, 0x73, 0x65, 0x73, 0x22, 0xb7,
	0x01, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x6d, 0x53, 0x74, 0x61, 0x72, 0x73, 0x68,
	0x69, 0x70, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x50, 0x0a, 0x06, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x38, 0x2e, 0x62, 0x75,
	0x66, 0x2e, 0x73, 0x74, 0x61, 0x72, 0x77, 0x61, 0x72, 0x73, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x6d, 0x53, 0x74,
	0x61, 0x72, 0x73, 0x68, 0x69, 0x70, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x1a, 0x49, 0x0a,
	0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x40, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x72, 0x73, 0x68,
	0x69, 0x70, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x62, 0x75, 0x66, 0x2e,
	0x73, 0x74, 0x61, 0x72, 0x77, 0x61, 0x72, 0x73, 0x2e, 0x73, 0x74, 0x61, 0x72, 0x73, 0x68, 0x69,
	0x70, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x72, 0x73, 0x68, 0x69, 0x70, 0x52, 0x09, 0x73,
	0x74, 0x61, 0x72, 0x73, 0x68, 0x69, 0x70, 0x73, 0x22, 0x49, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x51,
	0x75, 0x6f, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x32, 0x0a, 0x05, 0x62, 0x61, 0x73, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x62, 0x75, 0x66, 0x2e, 0x73, 0x74, 0x61, 0x72, 0x77, 0x61, 0x72, 0x73, 0x2e, 0x71, 0x75,
	0x6f, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x52, 0x05, 0x62, 0x61,
	0x73, 0x65, 0x73, 0x22, 0x9d, 0x01, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x51, 0x75, 0x6f, 0x74, 0x65,
	0x46, 0x69, 0x6c, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4c, 0x0a, 0x06,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x62,
	0x75, 0x66, 0x2e, 0x73, 0x74, 0x61, 0x72, 0x77, 0x61, 0x72, 0x73, 0x2e, 0x72, 0x65, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x51, 0x75, 0x6f, 0x74, 0x65,
	0x46, 0x69, 0x6c, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x1a, 0x37, 0x0a, 0x05, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x12, 0x2e, 0x0a, 0x04, 0x66, 0x69, 0x6c, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x73, 0x74, 0x61, 0x72, 0x77, 0x61, 0x72, 0x73,
	0x2e, 0x66, 0x69, 0x6c, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x6c, 0x6d, 0x52, 0x04, 0x66,
	0x69, 0x6c, 0x6d, 0x32, 0x9d, 0x02, 0x0a, 0x0f, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x8c, 0x01, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x46,
	0x69, 0x6c, 0x6d, 0x53, 0x74, 0x61, 0x72, 0x73, 0x68, 0x69, 0x70, 0x73, 0x12, 0x31, 0x2e, 0x62,
	0x75, 0x66, 0x2e, 0x73, 0x74, 0x61, 0x72, 0x77, 0x61, 0x72, 0x73, 0x2e, 0x72, 0x65, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x6d, 0x53,
	0x74, 0x61, 0x72, 0x73, 0x68, 0x69, 0x70, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x32, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x73, 0x74, 0x61, 0x72, 0x77, 0x61, 0x72, 0x73, 0x2e, 0x72,
	0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x69,
	0x6c, 0x6d, 0x53, 0x74, 0x61, 0x72, 0x73, 0x68, 0x69, 0x70, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x11, 0xaa, 0x48, 0x0b, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x72, 0x73, 0x68,
	0x69, 0x70, 0x73, 0x90, 0x02, 0x01, 0x12, 0x7b, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x51, 0x75, 0x6f,
	0x74, 0x65, 0x46, 0x69, 0x6c, 0x6d, 0x12, 0x2d, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x73, 0x74, 0x61,
	0x72, 0x77, 0x61, 0x72, 0x73, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x6d, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x73, 0x74, 0x61, 0x72,
	0x77, 0x61, 0x72, 0x73, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x6d, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x0c, 0xaa, 0x48, 0x06, 0x0a, 0x04, 0x66, 0x69, 0x6c, 0x6d,
	0x90, 0x02, 0x01, 0x42, 0x9d, 0x02, 0x0a, 0x1c, 0x63, 0x6f, 0x6d, 0x2e, 0x62, 0x75, 0x66, 0x2e,
	0x73, 0x74, 0x61, 0x72, 0x77, 0x61, 0x72, 0x73, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x76, 0x31, 0x42, 0x0d, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x6b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x62, 0x75, 0x66, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2f, 0x6b, 0x6e, 0x69, 0x74, 0x2f,
	0x74, 0x75, 0x74, 0x6f, 0x72, 0x69, 0x61, 0x6c, 0x2f, 0x73, 0x74, 0x61, 0x72, 0x77, 0x61, 0x72,
	0x73, 0x2d, 0x6b, 0x6e, 0x69, 0x74, 0x2d, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2d,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2d, 0x67, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x62,
	0x75, 0x66, 0x2f, 0x73, 0x74, 0x61, 0x72, 0x77, 0x61, 0x72, 0x73, 0x2f, 0x72, 0x65, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x3b, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x76, 0x31, 0xa2, 0x02, 0x03, 0x42, 0x53, 0x52, 0xaa, 0x02, 0x18, 0x42, 0x75, 0x66, 0x2e, 0x53,
	0x74, 0x61, 0x72, 0x77, 0x61, 0x72, 0x73, 0x2e, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x56, 0x31, 0xca, 0x02, 0x18, 0x42, 0x75, 0x66, 0x5c, 0x53, 0x74, 0x61, 0x72, 0x77, 0x61,
	0x72, 0x73, 0x5c, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5c, 0x56, 0x31, 0xe2, 0x02,
	0x24, 0x42, 0x75, 0x66, 0x5c, 0x53, 0x74, 0x61, 0x72, 0x77, 0x61, 0x72, 0x73, 0x5c, 0x52, 0x65,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x1b, 0x42, 0x75, 0x66, 0x3a, 0x3a, 0x53, 0x74, 0x61,
	0x72, 0x77, 0x61, 0x72, 0x73, 0x3a, 0x3a, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x3a,
	0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_buf_starwars_relation_v1_relation_proto_rawDescOnce sync.Once
	file_buf_starwars_relation_v1_relation_proto_rawDescData = file_buf_starwars_relation_v1_relation_proto_rawDesc
)

func file_buf_starwars_relation_v1_relation_proto_rawDescGZIP() []byte {
	file_buf_starwars_relation_v1_relation_proto_rawDescOnce.Do(func() {
		file_buf_starwars_relation_v1_relation_proto_rawDescData = protoimpl.X.CompressGZIP(file_buf_starwars_relation_v1_relation_proto_rawDescData)
	})
	return file_buf_starwars_relation_v1_relation_proto_rawDescData
}

var file_buf_starwars_relation_v1_relation_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_buf_starwars_relation_v1_relation_proto_goTypes = []interface{}{
	(*GetFilmStarshipsRequest)(nil),        // 0: buf.starwars.relation.v1.GetFilmStarshipsRequest
	(*GetFilmStarshipsResponse)(nil),       // 1: buf.starwars.relation.v1.GetFilmStarshipsResponse
	(*GetQuoteFilmRequest)(nil),            // 2: buf.starwars.relation.v1.GetQuoteFilmRequest
	(*GetQuoteFilmResponse)(nil),           // 3: buf.starwars.relation.v1.GetQuoteFilmResponse
	(*GetFilmStarshipsResponse_Value)(nil), // 4: buf.starwars.relation.v1.GetFilmStarshipsResponse.Value
	(*GetQuoteFilmResponse_Value)(nil),     // 5: buf.starwars.relation.v1.GetQuoteFilmResponse.Value
	(*v1.Film)(nil),                        // 6: buf.starwars.film.v1.Film
	(*v11.Quote)(nil),                      // 7: buf.starwars.quote.v1.Quote
	(*v12.Starship)(nil),                   // 8: buf.starwars.starship.v1.Starship
}
var file_buf_starwars_relation_v1_relation_proto_depIdxs = []int32{
	6, // 0: buf.starwars.relation.v1.GetFilmStarshipsRequest.bases:type_name -> buf.starwars.film.v1.Film
	4, // 1: buf.starwars.relation.v1.GetFilmStarshipsResponse.values:type_name -> buf.starwars.relation.v1.GetFilmStarshipsResponse.Value
	7, // 2: buf.starwars.relation.v1.GetQuoteFilmRequest.bases:type_name -> buf.starwars.quote.v1.Quote
	5, // 3: buf.starwars.relation.v1.GetQuoteFilmResponse.values:type_name -> buf.starwars.relation.v1.GetQuoteFilmResponse.Value
	8, // 4: buf.starwars.relation.v1.GetFilmStarshipsResponse.Value.starships:type_name -> buf.starwars.starship.v1.Starship
	6, // 5: buf.starwars.relation.v1.GetQuoteFilmResponse.Value.film:type_name -> buf.starwars.film.v1.Film
	0, // 6: buf.starwars.relation.v1.RelationService.GetFilmStarships:input_type -> buf.starwars.relation.v1.GetFilmStarshipsRequest
	2, // 7: buf.starwars.relation.v1.RelationService.GetQuoteFilm:input_type -> buf.starwars.relation.v1.GetQuoteFilmRequest
	1, // 8: buf.starwars.relation.v1.RelationService.GetFilmStarships:output_type -> buf.starwars.relation.v1.GetFilmStarshipsResponse
	3, // 9: buf.starwars.relation.v1.RelationService.GetQuoteFilm:output_type -> buf.starwars.relation.v1.GetQuoteFilmResponse
	8, // [8:10] is the sub-list for method output_type
	6, // [6:8] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_buf_starwars_relation_v1_relation_proto_init() }
func file_buf_starwars_relation_v1_relation_proto_init() {
	if File_buf_starwars_relation_v1_relation_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_buf_starwars_relation_v1_relation_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFilmStarshipsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_buf_starwars_relation_v1_relation_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFilmStarshipsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_buf_starwars_relation_v1_relation_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetQuoteFilmRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_buf_starwars_relation_v1_relation_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetQuoteFilmResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_buf_starwars_relation_v1_relation_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFilmStarshipsResponse_Value); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_buf_starwars_relation_v1_relation_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetQuoteFilmResponse_Value); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_buf_starwars_relation_v1_relation_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_buf_starwars_relation_v1_relation_proto_goTypes,
		DependencyIndexes: file_buf_starwars_relation_v1_relation_proto_depIdxs,
		MessageInfos:      file_buf_starwars_relation_v1_relation_proto_msgTypes,
	}.Build()
	File_buf_starwars_relation_v1_relation_proto = out.File
	file_buf_starwars_relation_v1_relation_proto_rawDesc = nil
	file_buf_starwars_relation_v1_relation_proto_goTypes = nil
	file_buf_starwars_relation_v1_relation_proto_depIdxs = nil
}
