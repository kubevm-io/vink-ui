// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: types/group_version.proto

package types

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

type GroupVersionResourceEnum int32

const (
	GroupVersionResourceEnum_UNSPECIFIED              GroupVersionResourceEnum = 0
	GroupVersionResourceEnum_VIRTUAL_MACHINE          GroupVersionResourceEnum = 1
	GroupVersionResourceEnum_VIRTUAL_MACHINE_INSTANCE GroupVersionResourceEnum = 2
	GroupVersionResourceEnum_DATA_VOLUME              GroupVersionResourceEnum = 3
	GroupVersionResourceEnum_NODE                     GroupVersionResourceEnum = 4
	GroupVersionResourceEnum_NAMESPACE                GroupVersionResourceEnum = 5
	GroupVersionResourceEnum_MULTUS                   GroupVersionResourceEnum = 6
	GroupVersionResourceEnum_SUBNET                   GroupVersionResourceEnum = 7
	GroupVersionResourceEnum_VPC                      GroupVersionResourceEnum = 8
	GroupVersionResourceEnum_IPPOOL                   GroupVersionResourceEnum = 9
)

// Enum value maps for GroupVersionResourceEnum.
var (
	GroupVersionResourceEnum_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "VIRTUAL_MACHINE",
		2: "VIRTUAL_MACHINE_INSTANCE",
		3: "DATA_VOLUME",
		4: "NODE",
		5: "NAMESPACE",
		6: "MULTUS",
		7: "SUBNET",
		8: "VPC",
		9: "IPPOOL",
	}
	GroupVersionResourceEnum_value = map[string]int32{
		"UNSPECIFIED":              0,
		"VIRTUAL_MACHINE":          1,
		"VIRTUAL_MACHINE_INSTANCE": 2,
		"DATA_VOLUME":              3,
		"NODE":                     4,
		"NAMESPACE":                5,
		"MULTUS":                   6,
		"SUBNET":                   7,
		"VPC":                      8,
		"IPPOOL":                   9,
	}
)

func (x GroupVersionResourceEnum) Enum() *GroupVersionResourceEnum {
	p := new(GroupVersionResourceEnum)
	*p = x
	return p
}

func (x GroupVersionResourceEnum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GroupVersionResourceEnum) Descriptor() protoreflect.EnumDescriptor {
	return file_types_group_version_proto_enumTypes[0].Descriptor()
}

func (GroupVersionResourceEnum) Type() protoreflect.EnumType {
	return &file_types_group_version_proto_enumTypes[0]
}

func (x GroupVersionResourceEnum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GroupVersionResourceEnum.Descriptor instead.
func (GroupVersionResourceEnum) EnumDescriptor() ([]byte, []int) {
	return file_types_group_version_proto_rawDescGZIP(), []int{0}
}

// GroupVersionResource unambiguously identifies a resource.  It doesn't anonymously include GroupVersion
// to avoid automatic coercion.  It doesn't use a GroupVersion to avoid custom marshalling
type GroupVersionResource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Group    string `protobuf:"bytes,1,opt,name=group,proto3" json:"group,omitempty"`
	Version  string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	Resource string `protobuf:"bytes,3,opt,name=resource,proto3" json:"resource,omitempty"`
}

func (x *GroupVersionResource) Reset() {
	*x = GroupVersionResource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_types_group_version_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupVersionResource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupVersionResource) ProtoMessage() {}

func (x *GroupVersionResource) ProtoReflect() protoreflect.Message {
	mi := &file_types_group_version_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupVersionResource.ProtoReflect.Descriptor instead.
func (*GroupVersionResource) Descriptor() ([]byte, []int) {
	return file_types_group_version_proto_rawDescGZIP(), []int{0}
}

func (x *GroupVersionResource) GetGroup() string {
	if x != nil {
		return x.Group
	}
	return ""
}

func (x *GroupVersionResource) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *GroupVersionResource) GetResource() string {
	if x != nil {
		return x.Resource
	}
	return ""
}

// GroupVersionKind unambiguously identifies a kind.  It doesn't anonymously include GroupVersion
// to avoid automatic coercion.  It doesn't use a GroupVersion to avoid custom marshalling
type GroupVersionKind struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Group   string `protobuf:"bytes,1,opt,name=group,proto3" json:"group,omitempty"`
	Version string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	Kind    string `protobuf:"bytes,3,opt,name=kind,proto3" json:"kind,omitempty"`
}

func (x *GroupVersionKind) Reset() {
	*x = GroupVersionKind{}
	if protoimpl.UnsafeEnabled {
		mi := &file_types_group_version_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupVersionKind) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupVersionKind) ProtoMessage() {}

func (x *GroupVersionKind) ProtoReflect() protoreflect.Message {
	mi := &file_types_group_version_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupVersionKind.ProtoReflect.Descriptor instead.
func (*GroupVersionKind) Descriptor() ([]byte, []int) {
	return file_types_group_version_proto_rawDescGZIP(), []int{1}
}

func (x *GroupVersionKind) GetGroup() string {
	if x != nil {
		return x.Group
	}
	return ""
}

func (x *GroupVersionKind) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *GroupVersionKind) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

type GroupVersionResourceIdentifier struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Option:
	//
	//	*GroupVersionResourceIdentifier_Enum
	//	*GroupVersionResourceIdentifier_Custom
	Option isGroupVersionResourceIdentifier_Option `protobuf_oneof:"option"`
}

func (x *GroupVersionResourceIdentifier) Reset() {
	*x = GroupVersionResourceIdentifier{}
	if protoimpl.UnsafeEnabled {
		mi := &file_types_group_version_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupVersionResourceIdentifier) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupVersionResourceIdentifier) ProtoMessage() {}

func (x *GroupVersionResourceIdentifier) ProtoReflect() protoreflect.Message {
	mi := &file_types_group_version_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupVersionResourceIdentifier.ProtoReflect.Descriptor instead.
func (*GroupVersionResourceIdentifier) Descriptor() ([]byte, []int) {
	return file_types_group_version_proto_rawDescGZIP(), []int{2}
}

func (m *GroupVersionResourceIdentifier) GetOption() isGroupVersionResourceIdentifier_Option {
	if m != nil {
		return m.Option
	}
	return nil
}

func (x *GroupVersionResourceIdentifier) GetEnum() GroupVersionResourceEnum {
	if x, ok := x.GetOption().(*GroupVersionResourceIdentifier_Enum); ok {
		return x.Enum
	}
	return GroupVersionResourceEnum_UNSPECIFIED
}

func (x *GroupVersionResourceIdentifier) GetCustom() *GroupVersionResource {
	if x, ok := x.GetOption().(*GroupVersionResourceIdentifier_Custom); ok {
		return x.Custom
	}
	return nil
}

type isGroupVersionResourceIdentifier_Option interface {
	isGroupVersionResourceIdentifier_Option()
}

type GroupVersionResourceIdentifier_Enum struct {
	Enum GroupVersionResourceEnum `protobuf:"varint,1,opt,name=enum,proto3,enum=vink.kubevm.io.apis.types.GroupVersionResourceEnum,oneof"`
}

type GroupVersionResourceIdentifier_Custom struct {
	Custom *GroupVersionResource `protobuf:"bytes,2,opt,name=custom,proto3,oneof"`
}

func (*GroupVersionResourceIdentifier_Enum) isGroupVersionResourceIdentifier_Option() {}

func (*GroupVersionResourceIdentifier_Custom) isGroupVersionResourceIdentifier_Option() {}

var File_types_group_version_proto protoreflect.FileDescriptor

var file_types_group_version_proto_rawDesc = []byte{
	0x0a, 0x19, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x76, 0x69, 0x6e,
	0x6b, 0x2e, 0x6b, 0x75, 0x62, 0x65, 0x76, 0x6d, 0x2e, 0x69, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x22, 0x62, 0x0a, 0x14, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x56,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1a,
	0x0a, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x22, 0x56, 0x0a, 0x10, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x4b, 0x69, 0x6e, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x12,
	0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6b, 0x69,
	0x6e, 0x64, 0x22, 0xc0, 0x01, 0x0a, 0x1e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x49, 0x0a, 0x04, 0x65, 0x6e, 0x75, 0x6d, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x33, 0x2e, 0x76, 0x69, 0x6e, 0x6b, 0x2e, 0x6b, 0x75, 0x62, 0x65, 0x76,
	0x6d, 0x2e, 0x69, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x48, 0x00, 0x52, 0x04, 0x65, 0x6e, 0x75, 0x6d,
	0x12, 0x49, 0x0a, 0x06, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x2f, 0x2e, 0x76, 0x69, 0x6e, 0x6b, 0x2e, 0x6b, 0x75, 0x62, 0x65, 0x76, 0x6d, 0x2e, 0x69,
	0x6f, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x48, 0x00, 0x52, 0x06, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x42, 0x08, 0x0a, 0x06, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2a, 0xb5, 0x01, 0x0a, 0x18, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x56,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x45, 0x6e,
	0x75, 0x6d, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45,
	0x44, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x56, 0x49, 0x52, 0x54, 0x55, 0x41, 0x4c, 0x5f, 0x4d,
	0x41, 0x43, 0x48, 0x49, 0x4e, 0x45, 0x10, 0x01, 0x12, 0x1c, 0x0a, 0x18, 0x56, 0x49, 0x52, 0x54,
	0x55, 0x41, 0x4c, 0x5f, 0x4d, 0x41, 0x43, 0x48, 0x49, 0x4e, 0x45, 0x5f, 0x49, 0x4e, 0x53, 0x54,
	0x41, 0x4e, 0x43, 0x45, 0x10, 0x02, 0x12, 0x0f, 0x0a, 0x0b, 0x44, 0x41, 0x54, 0x41, 0x5f, 0x56,
	0x4f, 0x4c, 0x55, 0x4d, 0x45, 0x10, 0x03, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x44, 0x45, 0x10,
	0x04, 0x12, 0x0d, 0x0a, 0x09, 0x4e, 0x41, 0x4d, 0x45, 0x53, 0x50, 0x41, 0x43, 0x45, 0x10, 0x05,
	0x12, 0x0a, 0x0a, 0x06, 0x4d, 0x55, 0x4c, 0x54, 0x55, 0x53, 0x10, 0x06, 0x12, 0x0a, 0x0a, 0x06,
	0x53, 0x55, 0x42, 0x4e, 0x45, 0x54, 0x10, 0x07, 0x12, 0x07, 0x0a, 0x03, 0x56, 0x50, 0x43, 0x10,
	0x08, 0x12, 0x0a, 0x0a, 0x06, 0x49, 0x50, 0x50, 0x4f, 0x4f, 0x4c, 0x10, 0x09, 0x42, 0x26, 0x5a,
	0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x75, 0x62, 0x65,
	0x76, 0x6d, 0x2e, 0x69, 0x6f, 0x2f, 0x76, 0x69, 0x6e, 0x6b, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_types_group_version_proto_rawDescOnce sync.Once
	file_types_group_version_proto_rawDescData = file_types_group_version_proto_rawDesc
)

func file_types_group_version_proto_rawDescGZIP() []byte {
	file_types_group_version_proto_rawDescOnce.Do(func() {
		file_types_group_version_proto_rawDescData = protoimpl.X.CompressGZIP(file_types_group_version_proto_rawDescData)
	})
	return file_types_group_version_proto_rawDescData
}

var file_types_group_version_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_types_group_version_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_types_group_version_proto_goTypes = []interface{}{
	(GroupVersionResourceEnum)(0),          // 0: vink.kubevm.io.apis.types.GroupVersionResourceEnum
	(*GroupVersionResource)(nil),           // 1: vink.kubevm.io.apis.types.GroupVersionResource
	(*GroupVersionKind)(nil),               // 2: vink.kubevm.io.apis.types.GroupVersionKind
	(*GroupVersionResourceIdentifier)(nil), // 3: vink.kubevm.io.apis.types.GroupVersionResourceIdentifier
}
var file_types_group_version_proto_depIdxs = []int32{
	0, // 0: vink.kubevm.io.apis.types.GroupVersionResourceIdentifier.enum:type_name -> vink.kubevm.io.apis.types.GroupVersionResourceEnum
	1, // 1: vink.kubevm.io.apis.types.GroupVersionResourceIdentifier.custom:type_name -> vink.kubevm.io.apis.types.GroupVersionResource
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_types_group_version_proto_init() }
func file_types_group_version_proto_init() {
	if File_types_group_version_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_types_group_version_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupVersionResource); i {
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
		file_types_group_version_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupVersionKind); i {
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
		file_types_group_version_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupVersionResourceIdentifier); i {
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
	file_types_group_version_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*GroupVersionResourceIdentifier_Enum)(nil),
		(*GroupVersionResourceIdentifier_Custom)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_types_group_version_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_types_group_version_proto_goTypes,
		DependencyIndexes: file_types_group_version_proto_depIdxs,
		EnumInfos:         file_types_group_version_proto_enumTypes,
		MessageInfos:      file_types_group_version_proto_msgTypes,
	}.Build()
	File_types_group_version_proto = out.File
	file_types_group_version_proto_rawDesc = nil
	file_types_group_version_proto_goTypes = nil
	file_types_group_version_proto_depIdxs = nil
}
