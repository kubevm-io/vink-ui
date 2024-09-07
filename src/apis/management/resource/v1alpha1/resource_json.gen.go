// Code generated by protoc-gen-jsonshim. DO NOT EDIT.
package v1alpha1

import (
	bytes "bytes"
	jsonpb "github.com/golang/protobuf/jsonpb"
)

// MarshalJSON is a custom marshaler for CreateRequest
func (this *CreateRequest) MarshalJSON() ([]byte, error) {
	str, err := ResourceMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for CreateRequest
func (this *CreateRequest) UnmarshalJSON(b []byte) error {
	return ResourceUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for UpdateRequest
func (this *UpdateRequest) MarshalJSON() ([]byte, error) {
	str, err := ResourceMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for UpdateRequest
func (this *UpdateRequest) UnmarshalJSON(b []byte) error {
	return ResourceUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for DeleteRequest
func (this *DeleteRequest) MarshalJSON() ([]byte, error) {
	str, err := ResourceMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for DeleteRequest
func (this *DeleteRequest) UnmarshalJSON(b []byte) error {
	return ResourceUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

var (
	ResourceMarshaler   = &jsonpb.Marshaler{}
	ResourceUnmarshaler = &jsonpb.Unmarshaler{AllowUnknownFields: true}
)
