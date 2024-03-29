// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.2
// source: peer_service.proto

package pb

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

type GetPeersListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SubnetworkID          []byte `protobuf:"bytes,1,opt,name=subnetworkID,proto3" json:"subnetworkID,omitempty"`
	IncludeAllSubnetworks bool   `protobuf:"varint,2,opt,name=includeAllSubnetworks,proto3" json:"includeAllSubnetworks,omitempty"`
}

func (x *GetPeersListRequest) Reset() {
	*x = GetPeersListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_peer_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPeersListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPeersListRequest) ProtoMessage() {}

func (x *GetPeersListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_peer_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPeersListRequest.ProtoReflect.Descriptor instead.
func (*GetPeersListRequest) Descriptor() ([]byte, []int) {
	return file_peer_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetPeersListRequest) GetSubnetworkID() []byte {
	if x != nil {
		return x.SubnetworkID
	}
	return nil
}

func (x *GetPeersListRequest) GetIncludeAllSubnetworks() bool {
	if x != nil {
		return x.IncludeAllSubnetworks
	}
	return false
}

type GetPeersListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Addresses []*NetAddress `protobuf:"bytes,1,rep,name=addresses,proto3" json:"addresses,omitempty"`
}

func (x *GetPeersListResponse) Reset() {
	*x = GetPeersListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_peer_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPeersListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPeersListResponse) ProtoMessage() {}

func (x *GetPeersListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_peer_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPeersListResponse.ProtoReflect.Descriptor instead.
func (*GetPeersListResponse) Descriptor() ([]byte, []int) {
	return file_peer_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetPeersListResponse) GetAddresses() []*NetAddress {
	if x != nil {
		return x.Addresses
	}
	return nil
}

type NetAddress struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Timestamp int64  `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	IP        []byte `protobuf:"bytes,2,opt,name=IP,proto3" json:"IP,omitempty"`
	Port      uint32 `protobuf:"varint,3,opt,name=port,proto3" json:"port,omitempty"`
}

func (x *NetAddress) Reset() {
	*x = NetAddress{}
	if protoimpl.UnsafeEnabled {
		mi := &file_peer_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NetAddress) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetAddress) ProtoMessage() {}

func (x *NetAddress) ProtoReflect() protoreflect.Message {
	mi := &file_peer_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetAddress.ProtoReflect.Descriptor instead.
func (*NetAddress) Descriptor() ([]byte, []int) {
	return file_peer_service_proto_rawDescGZIP(), []int{2}
}

func (x *NetAddress) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *NetAddress) GetIP() []byte {
	if x != nil {
		return x.IP
	}
	return nil
}

func (x *NetAddress) GetPort() uint32 {
	if x != nil {
		return x.Port
	}
	return 0
}

var File_peer_service_proto protoreflect.FileDescriptor

var file_peer_service_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x65, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6f, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x50, 0x65, 0x65, 0x72, 0x73,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x73,
	0x75, 0x62, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x0c, 0x73, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x44, 0x12,
	0x34, 0x0a, 0x15, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x41, 0x6c, 0x6c, 0x53, 0x75, 0x62,
	0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x15,
	0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x41, 0x6c, 0x6c, 0x53, 0x75, 0x62, 0x6e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x73, 0x22, 0x41, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x50, 0x65, 0x65, 0x72,
	0x73, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a,
	0x09, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0b, 0x2e, 0x4e, 0x65, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x09, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x22, 0x4e, 0x0a, 0x0a, 0x4e, 0x65, 0x74, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x50, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x02, 0x49, 0x50, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x32, 0x4c, 0x0a, 0x0b, 0x50, 0x65, 0x65, 0x72,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3d, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x50, 0x65,
	0x65, 0x72, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x14, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x65, 0x65,
	0x72, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e,
	0x47, 0x65, 0x74, 0x50, 0x65, 0x65, 0x72, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x28, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x61, 0x72, 0x6c, 0x73, 0x65, 0x6e, 0x2d, 0x6e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x6b, 0x61, 0x72, 0x6c, 0x73, 0x65, 0x6e, 0x64, 0x2f, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_peer_service_proto_rawDescOnce sync.Once
	file_peer_service_proto_rawDescData = file_peer_service_proto_rawDesc
)

func file_peer_service_proto_rawDescGZIP() []byte {
	file_peer_service_proto_rawDescOnce.Do(func() {
		file_peer_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_peer_service_proto_rawDescData)
	})
	return file_peer_service_proto_rawDescData
}

var file_peer_service_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_peer_service_proto_goTypes = []interface{}{
	(*GetPeersListRequest)(nil),  // 0: GetPeersListRequest
	(*GetPeersListResponse)(nil), // 1: GetPeersListResponse
	(*NetAddress)(nil),           // 2: NetAddress
}
var file_peer_service_proto_depIdxs = []int32{
	2, // 0: GetPeersListResponse.addresses:type_name -> NetAddress
	0, // 1: PeerService.GetPeersList:input_type -> GetPeersListRequest
	1, // 2: PeerService.GetPeersList:output_type -> GetPeersListResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_peer_service_proto_init() }
func file_peer_service_proto_init() {
	if File_peer_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_peer_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPeersListRequest); i {
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
		file_peer_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPeersListResponse); i {
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
		file_peer_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NetAddress); i {
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
			RawDescriptor: file_peer_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_peer_service_proto_goTypes,
		DependencyIndexes: file_peer_service_proto_depIdxs,
		MessageInfos:      file_peer_service_proto_msgTypes,
	}.Build()
	File_peer_service_proto = out.File
	file_peer_service_proto_rawDesc = nil
	file_peer_service_proto_goTypes = nil
	file_peer_service_proto_depIdxs = nil
}
