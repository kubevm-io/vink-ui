// Code generated by protoc-gen-deepcopy. DO NOT EDIT.
package v1alpha1

import (
	proto "github.com/golang/protobuf/proto"
)

// DeepCopyInto supports using CustomResourceDefinitionBak within kubernetes types, where deepcopy-gen is used.
func (in *CustomResourceDefinitionBak) DeepCopyInto(out *CustomResourceDefinitionBak) {
	p := proto.Clone(in).(*CustomResourceDefinitionBak)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomResourceDefinitionBak. Required by controller-gen.
func (in *CustomResourceDefinitionBak) DeepCopy() *CustomResourceDefinitionBak {
	if in == nil {
		return nil
	}
	out := new(CustomResourceDefinitionBak)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new CustomResourceDefinitionBak. Required by controller-gen.
func (in *CustomResourceDefinitionBak) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}
