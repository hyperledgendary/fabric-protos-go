// Copyright the Hyperledger Fabric contributors. All rights reserved.
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: peer/chaincode_event.proto

package peer

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

//ChaincodeEvent is used for events and registrations that are specific to chaincode
//string type - "chaincode"
type ChaincodeEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChaincodeId string `protobuf:"bytes,1,opt,name=chaincode_id,json=chaincodeId,proto3" json:"chaincode_id,omitempty"`
	TxId        string `protobuf:"bytes,2,opt,name=tx_id,json=txId,proto3" json:"tx_id,omitempty"`
	EventName   string `protobuf:"bytes,3,opt,name=event_name,json=eventName,proto3" json:"event_name,omitempty"`
	Payload     []byte `protobuf:"bytes,4,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *ChaincodeEvent) Reset() {
	*x = ChaincodeEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_peer_chaincode_event_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChaincodeEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChaincodeEvent) ProtoMessage() {}

func (x *ChaincodeEvent) ProtoReflect() protoreflect.Message {
	mi := &file_peer_chaincode_event_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChaincodeEvent.ProtoReflect.Descriptor instead.
func (*ChaincodeEvent) Descriptor() ([]byte, []int) {
	return file_peer_chaincode_event_proto_rawDescGZIP(), []int{0}
}

func (x *ChaincodeEvent) GetChaincodeId() string {
	if x != nil {
		return x.ChaincodeId
	}
	return ""
}

func (x *ChaincodeEvent) GetTxId() string {
	if x != nil {
		return x.TxId
	}
	return ""
}

func (x *ChaincodeEvent) GetEventName() string {
	if x != nil {
		return x.EventName
	}
	return ""
}

func (x *ChaincodeEvent) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

var File_peer_chaincode_event_proto protoreflect.FileDescriptor

var file_peer_chaincode_event_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x70, 0x65, 0x65, 0x72, 0x2f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x63, 0x6f, 0x64, 0x65,
	0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x22, 0x81, 0x01, 0x0a, 0x0e, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x63, 0x6f,
	0x64, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x68, 0x61, 0x69, 0x6e,
	0x63, 0x6f, 0x64, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63,
	0x68, 0x61, 0x69, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x12, 0x13, 0x0a, 0x05, 0x74, 0x78,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x78, 0x49, 0x64, 0x12,
	0x1d, 0x0a, 0x0a, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x42, 0x69, 0x0a, 0x22, 0x6f, 0x72, 0x67, 0x2e,
	0x68, 0x79, 0x70, 0x65, 0x72, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x66, 0x61, 0x62, 0x72,
	0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x70, 0x65, 0x65, 0x72, 0x42, 0x15,
	0x43, 0x68, 0x61, 0x69, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x50, 0x61,
	0x63, 0x6b, 0x61, 0x67, 0x65, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x68, 0x79, 0x70, 0x65, 0x72, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2f, 0x66, 0x61,
	0x62, 0x72, 0x69, 0x63, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2d, 0x67, 0x6f, 0x2f, 0x70,
	0x65, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_peer_chaincode_event_proto_rawDescOnce sync.Once
	file_peer_chaincode_event_proto_rawDescData = file_peer_chaincode_event_proto_rawDesc
)

func file_peer_chaincode_event_proto_rawDescGZIP() []byte {
	file_peer_chaincode_event_proto_rawDescOnce.Do(func() {
		file_peer_chaincode_event_proto_rawDescData = protoimpl.X.CompressGZIP(file_peer_chaincode_event_proto_rawDescData)
	})
	return file_peer_chaincode_event_proto_rawDescData
}

var file_peer_chaincode_event_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_peer_chaincode_event_proto_goTypes = []interface{}{
	(*ChaincodeEvent)(nil), // 0: protos.ChaincodeEvent
}
var file_peer_chaincode_event_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_peer_chaincode_event_proto_init() }
func file_peer_chaincode_event_proto_init() {
	if File_peer_chaincode_event_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_peer_chaincode_event_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChaincodeEvent); i {
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
			RawDescriptor: file_peer_chaincode_event_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_peer_chaincode_event_proto_goTypes,
		DependencyIndexes: file_peer_chaincode_event_proto_depIdxs,
		MessageInfos:      file_peer_chaincode_event_proto_msgTypes,
	}.Build()
	File_peer_chaincode_event_proto = out.File
	file_peer_chaincode_event_proto_rawDesc = nil
	file_peer_chaincode_event_proto_goTypes = nil
	file_peer_chaincode_event_proto_depIdxs = nil
}
