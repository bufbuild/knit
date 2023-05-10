// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: buf/starwars/quote/v1/quote.proto

package quotev1

import (
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

type Quote struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	QuoteId  string `protobuf:"bytes,1,opt,name=quote_id,json=quoteId,proto3" json:"quote_id,omitempty"`
	Text     string `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
	PersonId string `protobuf:"bytes,3,opt,name=person_id,json=personId,proto3" json:"person_id,omitempty"`
	FilmId   string `protobuf:"bytes,4,opt,name=film_id,json=filmId,proto3" json:"film_id,omitempty"`
}

func (x *Quote) Reset() {
	*x = Quote{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buf_starwars_quote_v1_quote_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Quote) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Quote) ProtoMessage() {}

func (x *Quote) ProtoReflect() protoreflect.Message {
	mi := &file_buf_starwars_quote_v1_quote_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Quote.ProtoReflect.Descriptor instead.
func (*Quote) Descriptor() ([]byte, []int) {
	return file_buf_starwars_quote_v1_quote_proto_rawDescGZIP(), []int{0}
}

func (x *Quote) GetQuoteId() string {
	if x != nil {
		return x.QuoteId
	}
	return ""
}

func (x *Quote) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *Quote) GetPersonId() string {
	if x != nil {
		return x.PersonId
	}
	return ""
}

func (x *Quote) GetFilmId() string {
	if x != nil {
		return x.FilmId
	}
	return ""
}

type StreamQuotesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit uint64 `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *StreamQuotesRequest) Reset() {
	*x = StreamQuotesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buf_starwars_quote_v1_quote_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamQuotesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamQuotesRequest) ProtoMessage() {}

func (x *StreamQuotesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_buf_starwars_quote_v1_quote_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamQuotesRequest.ProtoReflect.Descriptor instead.
func (*StreamQuotesRequest) Descriptor() ([]byte, []int) {
	return file_buf_starwars_quote_v1_quote_proto_rawDescGZIP(), []int{1}
}

func (x *StreamQuotesRequest) GetLimit() uint64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type StreamQuotesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Quote *Quote `protobuf:"bytes,1,opt,name=quote,proto3" json:"quote,omitempty"`
}

func (x *StreamQuotesResponse) Reset() {
	*x = StreamQuotesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buf_starwars_quote_v1_quote_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamQuotesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamQuotesResponse) ProtoMessage() {}

func (x *StreamQuotesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_buf_starwars_quote_v1_quote_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamQuotesResponse.ProtoReflect.Descriptor instead.
func (*StreamQuotesResponse) Descriptor() ([]byte, []int) {
	return file_buf_starwars_quote_v1_quote_proto_rawDescGZIP(), []int{2}
}

func (x *StreamQuotesResponse) GetQuote() *Quote {
	if x != nil {
		return x.Quote
	}
	return nil
}

var File_buf_starwars_quote_v1_quote_proto protoreflect.FileDescriptor

var file_buf_starwars_quote_v1_quote_proto_rawDesc = []byte{
	0x0a, 0x21, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x61, 0x72, 0x77, 0x61, 0x72, 0x73, 0x2f, 0x71,
	0x75, 0x6f, 0x74, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x15, 0x62, 0x75, 0x66, 0x2e, 0x73, 0x74, 0x61, 0x72, 0x77, 0x61, 0x72,
	0x73, 0x2e, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x22, 0x6c, 0x0a, 0x05, 0x51, 0x75,
	0x6f, 0x74, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x49, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65,
	0x78, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x49, 0x64, 0x12,
	0x17, 0x0a, 0x07, 0x66, 0x69, 0x6c, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x66, 0x69, 0x6c, 0x6d, 0x49, 0x64, 0x22, 0x2b, 0x0a, 0x13, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x4a, 0x0a, 0x14, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x51,
	0x75, 0x6f, 0x74, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a,
	0x05, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x62,
	0x75, 0x66, 0x2e, 0x73, 0x74, 0x61, 0x72, 0x77, 0x61, 0x72, 0x73, 0x2e, 0x71, 0x75, 0x6f, 0x74,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x52, 0x05, 0x71, 0x75, 0x6f, 0x74,
	0x65, 0x32, 0x7b, 0x0a, 0x0c, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x6b, 0x0a, 0x0c, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x51, 0x75, 0x6f, 0x74, 0x65,
	0x73, 0x12, 0x2a, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x73, 0x74, 0x61, 0x72, 0x77, 0x61, 0x72, 0x73,
	0x2e, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x51, 0x75, 0x6f, 0x74, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e,
	0x62, 0x75, 0x66, 0x2e, 0x73, 0x74, 0x61, 0x72, 0x77, 0x61, 0x72, 0x73, 0x2e, 0x71, 0x75, 0x6f,
	0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x51, 0x75, 0x6f, 0x74,
	0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x42, 0xfd,
	0x01, 0x0a, 0x19, 0x63, 0x6f, 0x6d, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x73, 0x74, 0x61, 0x72, 0x77,
	0x61, 0x72, 0x73, 0x2e, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x42, 0x0a, 0x51, 0x75,
	0x6f, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x5d, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x75, 0x66, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2f,
	0x6b, 0x6e, 0x69, 0x74, 0x2f, 0x74, 0x75, 0x74, 0x6f, 0x72, 0x69, 0x61, 0x6c, 0x2f, 0x73, 0x74,
	0x61, 0x72, 0x77, 0x61, 0x72, 0x73, 0x2d, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x2d, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2d, 0x67, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x62, 0x75, 0x66, 0x2f,
	0x73, 0x74, 0x61, 0x72, 0x77, 0x61, 0x72, 0x73, 0x2f, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x2f, 0x76,
	0x31, 0x3b, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x42, 0x53, 0x51, 0xaa,
	0x02, 0x15, 0x42, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x61, 0x72, 0x77, 0x61, 0x72, 0x73, 0x2e, 0x51,
	0x75, 0x6f, 0x74, 0x65, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x15, 0x42, 0x75, 0x66, 0x5c, 0x53, 0x74,
	0x61, 0x72, 0x77, 0x61, 0x72, 0x73, 0x5c, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x5c, 0x56, 0x31, 0xe2,
	0x02, 0x21, 0x42, 0x75, 0x66, 0x5c, 0x53, 0x74, 0x61, 0x72, 0x77, 0x61, 0x72, 0x73, 0x5c, 0x51,
	0x75, 0x6f, 0x74, 0x65, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0xea, 0x02, 0x18, 0x42, 0x75, 0x66, 0x3a, 0x3a, 0x53, 0x74, 0x61, 0x72, 0x77,
	0x61, 0x72, 0x73, 0x3a, 0x3a, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_buf_starwars_quote_v1_quote_proto_rawDescOnce sync.Once
	file_buf_starwars_quote_v1_quote_proto_rawDescData = file_buf_starwars_quote_v1_quote_proto_rawDesc
)

func file_buf_starwars_quote_v1_quote_proto_rawDescGZIP() []byte {
	file_buf_starwars_quote_v1_quote_proto_rawDescOnce.Do(func() {
		file_buf_starwars_quote_v1_quote_proto_rawDescData = protoimpl.X.CompressGZIP(file_buf_starwars_quote_v1_quote_proto_rawDescData)
	})
	return file_buf_starwars_quote_v1_quote_proto_rawDescData
}

var file_buf_starwars_quote_v1_quote_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_buf_starwars_quote_v1_quote_proto_goTypes = []interface{}{
	(*Quote)(nil),                // 0: buf.starwars.quote.v1.Quote
	(*StreamQuotesRequest)(nil),  // 1: buf.starwars.quote.v1.StreamQuotesRequest
	(*StreamQuotesResponse)(nil), // 2: buf.starwars.quote.v1.StreamQuotesResponse
}
var file_buf_starwars_quote_v1_quote_proto_depIdxs = []int32{
	0, // 0: buf.starwars.quote.v1.StreamQuotesResponse.quote:type_name -> buf.starwars.quote.v1.Quote
	1, // 1: buf.starwars.quote.v1.QuoteService.StreamQuotes:input_type -> buf.starwars.quote.v1.StreamQuotesRequest
	2, // 2: buf.starwars.quote.v1.QuoteService.StreamQuotes:output_type -> buf.starwars.quote.v1.StreamQuotesResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_buf_starwars_quote_v1_quote_proto_init() }
func file_buf_starwars_quote_v1_quote_proto_init() {
	if File_buf_starwars_quote_v1_quote_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_buf_starwars_quote_v1_quote_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Quote); i {
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
		file_buf_starwars_quote_v1_quote_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamQuotesRequest); i {
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
		file_buf_starwars_quote_v1_quote_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamQuotesResponse); i {
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
			RawDescriptor: file_buf_starwars_quote_v1_quote_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_buf_starwars_quote_v1_quote_proto_goTypes,
		DependencyIndexes: file_buf_starwars_quote_v1_quote_proto_depIdxs,
		MessageInfos:      file_buf_starwars_quote_v1_quote_proto_msgTypes,
	}.Build()
	File_buf_starwars_quote_v1_quote_proto = out.File
	file_buf_starwars_quote_v1_quote_proto_rawDesc = nil
	file_buf_starwars_quote_v1_quote_proto_goTypes = nil
	file_buf_starwars_quote_v1_quote_proto_depIdxs = nil
}
