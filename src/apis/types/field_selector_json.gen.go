// Code generated by protoc-gen-jsonshim. DO NOT EDIT.
package types

import (
	bytes "bytes"
	jsonpb "github.com/golang/protobuf/jsonpb"
)

// MarshalJSON is a custom marshaler for FieldSelector
func (this *FieldSelector) MarshalJSON() ([]byte, error) {
	str, err := FieldSelectorMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for FieldSelector
func (this *FieldSelector) UnmarshalJSON(b []byte) error {
	return FieldSelectorUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for Condition
func (this *Condition) MarshalJSON() ([]byte, error) {
	str, err := FieldSelectorMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for Condition
func (this *Condition) UnmarshalJSON(b []byte) error {
	return FieldSelectorUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

var (
	FieldSelectorMarshaler   = &jsonpb.Marshaler{}
	FieldSelectorUnmarshaler = &jsonpb.Unmarshaler{AllowUnknownFields: true}
)
