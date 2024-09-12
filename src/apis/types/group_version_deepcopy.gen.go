// Code generated by protoc-gen-deepcopy. DO NOT EDIT.
package types

import (
	proto "github.com/golang/protobuf/proto"
)

// DeepCopyInto supports using GroupVersionResource within kubernetes types, where deepcopy-gen is used.
func (in *GroupVersionResource) DeepCopyInto(out *GroupVersionResource) {
	p := proto.Clone(in).(*GroupVersionResource)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GroupVersionResource. Required by controller-gen.
func (in *GroupVersionResource) DeepCopy() *GroupVersionResource {
	if in == nil {
		return nil
	}
	out := new(GroupVersionResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new GroupVersionResource. Required by controller-gen.
func (in *GroupVersionResource) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using GroupVersionKind within kubernetes types, where deepcopy-gen is used.
func (in *GroupVersionKind) DeepCopyInto(out *GroupVersionKind) {
	p := proto.Clone(in).(*GroupVersionKind)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GroupVersionKind. Required by controller-gen.
func (in *GroupVersionKind) DeepCopy() *GroupVersionKind {
	if in == nil {
		return nil
	}
	out := new(GroupVersionKind)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new GroupVersionKind. Required by controller-gen.
func (in *GroupVersionKind) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using GroupVersionResourceIdentifier within kubernetes types, where deepcopy-gen is used.
func (in *GroupVersionResourceIdentifier) DeepCopyInto(out *GroupVersionResourceIdentifier) {
	p := proto.Clone(in).(*GroupVersionResourceIdentifier)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GroupVersionResourceIdentifier. Required by controller-gen.
func (in *GroupVersionResourceIdentifier) DeepCopy() *GroupVersionResourceIdentifier {
	if in == nil {
		return nil
	}
	out := new(GroupVersionResourceIdentifier)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new GroupVersionResourceIdentifier. Required by controller-gen.
func (in *GroupVersionResourceIdentifier) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}