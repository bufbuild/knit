// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: buf/starwars/starship/v1/starship.proto

package starshipv1

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

type Starship struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StarshipId           string  `protobuf:"bytes,1,opt,name=starship_id,json=starshipId,proto3" json:"starship_id,omitempty"`
	Name                 string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Model                string  `protobuf:"bytes,3,opt,name=model,proto3" json:"model,omitempty"`
	Manufacturer         string  `protobuf:"bytes,4,opt,name=manufacturer,proto3" json:"manufacturer,omitempty"`
	CostInCredits        uint64  `protobuf:"varint,5,opt,name=cost_in_credits,json=costInCredits,proto3" json:"cost_in_credits,omitempty"`
	Length               uint64  `protobuf:"varint,6,opt,name=length,proto3" json:"length,omitempty"`
	MaxAtmospheringSpeed uint64  `protobuf:"varint,7,opt,name=max_atmosphering_speed,json=maxAtmospheringSpeed,proto3" json:"max_atmosphering_speed,omitempty"`
	Crew                 uint64  `protobuf:"varint,8,opt,name=crew,proto3" json:"crew,omitempty"`
	Passengers           uint64  `protobuf:"varint,9,opt,name=passengers,proto3" json:"passengers,omitempty"`
	CargoCapacity        uint64  `protobuf:"varint,10,opt,name=cargo_capacity,json=cargoCapacity,proto3" json:"cargo_capacity,omitempty"`
	Consumables          string  `protobuf:"bytes,11,opt,name=consumables,proto3" json:"consumables,omitempty"`
	HyperdriveRating     float32 `protobuf:"fixed32,12,opt,name=hyperdrive_rating,json=hyperdriveRating,proto3" json:"hyperdrive_rating,omitempty"`
	Mglt                 uint64  `protobuf:"varint,13,opt,name=mglt,proto3" json:"mglt,omitempty"`
	StarshipClass        string  `protobuf:"bytes,14,opt,name=starship_class,json=starshipClass,proto3" json:"starship_class,omitempty"`
	// Relations
	PilotIds []string `protobuf:"bytes,15,rep,name=pilot_ids,json=pilotIds,proto3" json:"pilot_ids,omitempty"`
	FilmIds  []string `protobuf:"bytes,16,rep,name=film_ids,json=filmIds,proto3" json:"film_ids,omitempty"`
}

func (x *Starship) Reset() {
	*x = Starship{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buf_starwars_starship_v1_starship_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Starship) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Starship) ProtoMessage() {}

func (x *Starship) ProtoReflect() protoreflect.Message {
	mi := &file_buf_starwars_starship_v1_starship_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Starship.ProtoReflect.Descriptor instead.
func (*Starship) Descriptor() ([]byte, []int) {
	return file_buf_starwars_starship_v1_starship_proto_rawDescGZIP(), []int{0}
}

func (x *Starship) GetStarshipId() string {
	if x != nil {
		return x.StarshipId
	}
	return ""
}

func (x *Starship) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Starship) GetModel() string {
	if x != nil {
		return x.Model
	}
	return ""
}

func (x *Starship) GetManufacturer() string {
	if x != nil {
		return x.Manufacturer
	}
	return ""
}

func (x *Starship) GetCostInCredits() uint64 {
	if x != nil {
		return x.CostInCredits
	}
	return 0
}

func (x *Starship) GetLength() uint64 {
	if x != nil {
		return x.Length
	}
	return 0
}

func (x *Starship) GetMaxAtmospheringSpeed() uint64 {
	if x != nil {
		return x.MaxAtmospheringSpeed
	}
	return 0
}

func (x *Starship) GetCrew() uint64 {
	if x != nil {
		return x.Crew
	}
	return 0
}

func (x *Starship) GetPassengers() uint64 {
	if x != nil {
		return x.Passengers
	}
	return 0
}

func (x *Starship) GetCargoCapacity() uint64 {
	if x != nil {
		return x.CargoCapacity
	}
	return 0
}

func (x *Starship) GetConsumables() string {
	if x != nil {
		return x.Consumables
	}
	return ""
}

func (x *Starship) GetHyperdriveRating() float32 {
	if x != nil {
		return x.HyperdriveRating
	}
	return 0
}

func (x *Starship) GetMglt() uint64 {
	if x != nil {
		return x.Mglt
	}
	return 0
}

func (x *Starship) GetStarshipClass() string {
	if x != nil {
		return x.StarshipClass
	}
	return ""
}

func (x *Starship) GetPilotIds() []string {
	if x != nil {
		return x.PilotIds
	}
	return nil
}

func (x *Starship) GetFilmIds() []string {
	if x != nil {
		return x.FilmIds
	}
	return nil
}

type StarshipsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StarshipIds []string `protobuf:"bytes,1,rep,name=starship_ids,json=starshipIds,proto3" json:"starship_ids,omitempty"`
	Limit       uint32   `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *StarshipsRequest) Reset() {
	*x = StarshipsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buf_starwars_starship_v1_starship_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StarshipsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StarshipsRequest) ProtoMessage() {}

func (x *StarshipsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_buf_starwars_starship_v1_starship_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StarshipsRequest.ProtoReflect.Descriptor instead.
func (*StarshipsRequest) Descriptor() ([]byte, []int) {
	return file_buf_starwars_starship_v1_starship_proto_rawDescGZIP(), []int{1}
}

func (x *StarshipsRequest) GetStarshipIds() []string {
	if x != nil {
		return x.StarshipIds
	}
	return nil
}

func (x *StarshipsRequest) GetLimit() uint32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type StarshipsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Starships []*Starship `protobuf:"bytes,1,rep,name=starships,proto3" json:"starships,omitempty"`
}

func (x *StarshipsResponse) Reset() {
	*x = StarshipsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buf_starwars_starship_v1_starship_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StarshipsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StarshipsResponse) ProtoMessage() {}

func (x *StarshipsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_buf_starwars_starship_v1_starship_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StarshipsResponse.ProtoReflect.Descriptor instead.
func (*StarshipsResponse) Descriptor() ([]byte, []int) {
	return file_buf_starwars_starship_v1_starship_proto_rawDescGZIP(), []int{2}
}

func (x *StarshipsResponse) GetStarships() []*Starship {
	if x != nil {
		return x.Starships
	}
	return nil
}

var File_buf_starwars_starship_v1_starship_proto protoreflect.FileDescriptor

var file_buf_starwars_starship_v1_starship_proto_rawDesc = []byte{
	0x0a, 0x27, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x61, 0x72, 0x77, 0x61, 0x72, 0x73, 0x2f, 0x73,
	0x74, 0x61, 0x72, 0x73, 0x68, 0x69, 0x70, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x61, 0x72, 0x73,
	0x68, 0x69, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x62, 0x75, 0x66, 0x2e, 0x73,
	0x74, 0x61, 0x72, 0x77, 0x61, 0x72, 0x73, 0x2e, 0x73, 0x74, 0x61, 0x72, 0x73, 0x68, 0x69, 0x70,
	0x2e, 0x76, 0x31, 0x22, 0x8c, 0x04, 0x0a, 0x08, 0x53, 0x74, 0x61, 0x72, 0x73, 0x68, 0x69, 0x70,
	0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x72, 0x73, 0x68, 0x69, 0x70, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x73, 0x68, 0x69, 0x70, 0x49,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x22, 0x0a, 0x0c, 0x6d,
	0x61, 0x6e, 0x75, 0x66, 0x61, 0x63, 0x74, 0x75, 0x72, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x6d, 0x61, 0x6e, 0x75, 0x66, 0x61, 0x63, 0x74, 0x75, 0x72, 0x65, 0x72, 0x12,
	0x26, 0x0a, 0x0f, 0x63, 0x6f, 0x73, 0x74, 0x5f, 0x69, 0x6e, 0x5f, 0x63, 0x72, 0x65, 0x64, 0x69,
	0x74, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0d, 0x63, 0x6f, 0x73, 0x74, 0x49, 0x6e,
	0x43, 0x72, 0x65, 0x64, 0x69, 0x74, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x65, 0x6e, 0x67, 0x74,
	0x68, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x12,
	0x34, 0x0a, 0x16, 0x6d, 0x61, 0x78, 0x5f, 0x61, 0x74, 0x6d, 0x6f, 0x73, 0x70, 0x68, 0x65, 0x72,
	0x69, 0x6e, 0x67, 0x5f, 0x73, 0x70, 0x65, 0x65, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x14, 0x6d, 0x61, 0x78, 0x41, 0x74, 0x6d, 0x6f, 0x73, 0x70, 0x68, 0x65, 0x72, 0x69, 0x6e, 0x67,
	0x53, 0x70, 0x65, 0x65, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x72, 0x65, 0x77, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x04, 0x63, 0x72, 0x65, 0x77, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x61, 0x73,
	0x73, 0x65, 0x6e, 0x67, 0x65, 0x72, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x70,
	0x61, 0x73, 0x73, 0x65, 0x6e, 0x67, 0x65, 0x72, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x61, 0x72,
	0x67, 0x6f, 0x5f, 0x63, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x0d, 0x63, 0x61, 0x72, 0x67, 0x6f, 0x43, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79,
	0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x61, 0x62, 0x6c,
	0x65, 0x73, 0x12, 0x2b, 0x0a, 0x11, 0x68, 0x79, 0x70, 0x65, 0x72, 0x64, 0x72, 0x69, 0x76, 0x65,
	0x5f, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x02, 0x52, 0x10, 0x68,
	0x79, 0x70, 0x65, 0x72, 0x64, 0x72, 0x69, 0x76, 0x65, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x12,
	0x12, 0x0a, 0x04, 0x6d, 0x67, 0x6c, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x6d,
	0x67, 0x6c, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x74, 0x61, 0x72, 0x73, 0x68, 0x69, 0x70, 0x5f,
	0x63, 0x6c, 0x61, 0x73, 0x73, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x74, 0x61,
	0x72, 0x73, 0x68, 0x69, 0x70, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x69,
	0x6c, 0x6f, 0x74, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x0f, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x70,
	0x69, 0x6c, 0x6f, 0x74, 0x49, 0x64, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x6d, 0x5f,
	0x69, 0x64, 0x73, 0x18, 0x10, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x66, 0x69, 0x6c, 0x6d, 0x49,
	0x64, 0x73, 0x22, 0x4b, 0x0a, 0x10, 0x53, 0x74, 0x61, 0x72, 0x73, 0x68, 0x69, 0x70, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x74, 0x61, 0x72, 0x73, 0x68,
	0x69, 0x70, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x74,
	0x61, 0x72, 0x73, 0x68, 0x69, 0x70, 0x49, 0x64, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x22,
	0x55, 0x0a, 0x11, 0x53, 0x74, 0x61, 0x72, 0x73, 0x68, 0x69, 0x70, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x72, 0x73, 0x68, 0x69, 0x70,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x73, 0x74,
	0x61, 0x72, 0x77, 0x61, 0x72, 0x73, 0x2e, 0x73, 0x74, 0x61, 0x72, 0x73, 0x68, 0x69, 0x70, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x72, 0x73, 0x68, 0x69, 0x70, 0x52, 0x09, 0x73, 0x74, 0x61,
	0x72, 0x73, 0x68, 0x69, 0x70, 0x73, 0x32, 0x7a, 0x0a, 0x0f, 0x53, 0x74, 0x61, 0x72, 0x73, 0x68,
	0x69, 0x70, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x67, 0x0a, 0x0c, 0x47, 0x65, 0x74,
	0x53, 0x74, 0x61, 0x72, 0x73, 0x68, 0x69, 0x70, 0x73, 0x12, 0x2a, 0x2e, 0x62, 0x75, 0x66, 0x2e,
	0x73, 0x74, 0x61, 0x72, 0x77, 0x61, 0x72, 0x73, 0x2e, 0x73, 0x74, 0x61, 0x72, 0x73, 0x68, 0x69,
	0x70, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x72, 0x73, 0x68, 0x69, 0x70, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x73, 0x74, 0x61, 0x72,
	0x77, 0x61, 0x72, 0x73, 0x2e, 0x73, 0x74, 0x61, 0x72, 0x73, 0x68, 0x69, 0x70, 0x2e, 0x76, 0x31,
	0x2e, 0x53, 0x74, 0x61, 0x72, 0x73, 0x68, 0x69, 0x70, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x98, 0x02, 0x0a, 0x1c, 0x63, 0x6f, 0x6d, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x73,
	0x74, 0x61, 0x72, 0x77, 0x61, 0x72, 0x73, 0x2e, 0x73, 0x74, 0x61, 0x72, 0x73, 0x68, 0x69, 0x70,
	0x2e, 0x76, 0x31, 0x42, 0x0d, 0x53, 0x74, 0x61, 0x72, 0x73, 0x68, 0x69, 0x70, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x66, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x62, 0x75, 0x66, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2f, 0x6b, 0x6e, 0x69, 0x74, 0x2f, 0x74,
	0x75, 0x74, 0x6f, 0x72, 0x69, 0x61, 0x6c, 0x2f, 0x73, 0x74, 0x61, 0x72, 0x77, 0x61, 0x72, 0x73,
	0x2d, 0x73, 0x74, 0x61, 0x72, 0x73, 0x68, 0x69, 0x70, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2d, 0x67, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x61,
	0x72, 0x77, 0x61, 0x72, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x72, 0x73, 0x68, 0x69, 0x70, 0x2f, 0x76,
	0x31, 0x3b, 0x73, 0x74, 0x61, 0x72, 0x73, 0x68, 0x69, 0x70, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x42,
	0x53, 0x53, 0xaa, 0x02, 0x18, 0x42, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x61, 0x72, 0x77, 0x61, 0x72,
	0x73, 0x2e, 0x53, 0x74, 0x61, 0x72, 0x73, 0x68, 0x69, 0x70, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x18,
	0x42, 0x75, 0x66, 0x5c, 0x53, 0x74, 0x61, 0x72, 0x77, 0x61, 0x72, 0x73, 0x5c, 0x53, 0x74, 0x61,
	0x72, 0x73, 0x68, 0x69, 0x70, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x24, 0x42, 0x75, 0x66, 0x5c, 0x53,
	0x74, 0x61, 0x72, 0x77, 0x61, 0x72, 0x73, 0x5c, 0x53, 0x74, 0x61, 0x72, 0x73, 0x68, 0x69, 0x70,
	0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea,
	0x02, 0x1b, 0x42, 0x75, 0x66, 0x3a, 0x3a, 0x53, 0x74, 0x61, 0x72, 0x77, 0x61, 0x72, 0x73, 0x3a,
	0x3a, 0x53, 0x74, 0x61, 0x72, 0x73, 0x68, 0x69, 0x70, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_buf_starwars_starship_v1_starship_proto_rawDescOnce sync.Once
	file_buf_starwars_starship_v1_starship_proto_rawDescData = file_buf_starwars_starship_v1_starship_proto_rawDesc
)

func file_buf_starwars_starship_v1_starship_proto_rawDescGZIP() []byte {
	file_buf_starwars_starship_v1_starship_proto_rawDescOnce.Do(func() {
		file_buf_starwars_starship_v1_starship_proto_rawDescData = protoimpl.X.CompressGZIP(file_buf_starwars_starship_v1_starship_proto_rawDescData)
	})
	return file_buf_starwars_starship_v1_starship_proto_rawDescData
}

var file_buf_starwars_starship_v1_starship_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_buf_starwars_starship_v1_starship_proto_goTypes = []interface{}{
	(*Starship)(nil),          // 0: buf.starwars.starship.v1.Starship
	(*StarshipsRequest)(nil),  // 1: buf.starwars.starship.v1.StarshipsRequest
	(*StarshipsResponse)(nil), // 2: buf.starwars.starship.v1.StarshipsResponse
}
var file_buf_starwars_starship_v1_starship_proto_depIdxs = []int32{
	0, // 0: buf.starwars.starship.v1.StarshipsResponse.starships:type_name -> buf.starwars.starship.v1.Starship
	1, // 1: buf.starwars.starship.v1.StarshipService.GetStarships:input_type -> buf.starwars.starship.v1.StarshipsRequest
	2, // 2: buf.starwars.starship.v1.StarshipService.GetStarships:output_type -> buf.starwars.starship.v1.StarshipsResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_buf_starwars_starship_v1_starship_proto_init() }
func file_buf_starwars_starship_v1_starship_proto_init() {
	if File_buf_starwars_starship_v1_starship_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_buf_starwars_starship_v1_starship_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Starship); i {
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
		file_buf_starwars_starship_v1_starship_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StarshipsRequest); i {
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
		file_buf_starwars_starship_v1_starship_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StarshipsResponse); i {
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
			RawDescriptor: file_buf_starwars_starship_v1_starship_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_buf_starwars_starship_v1_starship_proto_goTypes,
		DependencyIndexes: file_buf_starwars_starship_v1_starship_proto_depIdxs,
		MessageInfos:      file_buf_starwars_starship_v1_starship_proto_msgTypes,
	}.Build()
	File_buf_starwars_starship_v1_starship_proto = out.File
	file_buf_starwars_starship_v1_starship_proto_rawDesc = nil
	file_buf_starwars_starship_v1_starship_proto_goTypes = nil
	file_buf_starwars_starship_v1_starship_proto_depIdxs = nil
}
