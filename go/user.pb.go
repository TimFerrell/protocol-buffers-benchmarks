// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

/*
Package usertest is a generated protocol buffer package.

It is generated from these files:
	user.proto

It has these top-level messages:
	User
	Users
*/
package usertest

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type User struct {
	Id      int32  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Name    string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Email   string `protobuf:"bytes,3,opt,name=email" json:"email,omitempty"`
	Twitter string `protobuf:"bytes,4,opt,name=twitter" json:"twitter,omitempty"`
	Street  string `protobuf:"bytes,5,opt,name=street" json:"street,omitempty"`
	City    string `protobuf:"bytes,6,opt,name=city" json:"city,omitempty"`
	State   string `protobuf:"bytes,7,opt,name=state" json:"state,omitempty"`
	Zip     string `protobuf:"bytes,8,opt,name=zip" json:"zip,omitempty"`
	Phone   string `protobuf:"bytes,9,opt,name=phone" json:"phone,omitempty"`
}

func (m *User) Reset()                    { *m = User{} }
func (m *User) String() string            { return proto.CompactTextString(m) }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *User) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetTwitter() string {
	if m != nil {
		return m.Twitter
	}
	return ""
}

func (m *User) GetStreet() string {
	if m != nil {
		return m.Street
	}
	return ""
}

func (m *User) GetCity() string {
	if m != nil {
		return m.City
	}
	return ""
}

func (m *User) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *User) GetZip() string {
	if m != nil {
		return m.Zip
	}
	return ""
}

func (m *User) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

type Users struct {
	Users []*User `protobuf:"bytes,1,rep,name=users" json:"users,omitempty"`
}

func (m *Users) Reset()                    { *m = Users{} }
func (m *Users) String() string            { return proto.CompactTextString(m) }
func (*Users) ProtoMessage()               {}
func (*Users) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Users) GetUsers() []*User {
	if m != nil {
		return m.Users
	}
	return nil
}

func init() {
	proto.RegisterType((*User)(nil), "usertest.User")
	proto.RegisterType((*Users)(nil), "usertest.Users")
}

func init() { proto.RegisterFile("user.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 204 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x34, 0x90, 0xcf, 0x4a, 0xc6, 0x40,
	0x0c, 0xc4, 0xd9, 0xb6, 0xdb, 0xef, 0x6b, 0x84, 0x22, 0x41, 0x24, 0xc7, 0x52, 0x3c, 0xf4, 0x62,
	0x0f, 0xfa, 0x26, 0x05, 0x1f, 0xa0, 0xda, 0x80, 0x0b, 0xf6, 0x0f, 0xbb, 0x11, 0xd1, 0xe7, 0xf3,
	0xc1, 0x24, 0x59, 0x7b, 0x9b, 0xdf, 0x4c, 0x18, 0x98, 0x00, 0x7c, 0x26, 0x8e, 0xe3, 0x11, 0x77,
	0xd9, 0xf1, 0xaa, 0x5a, 0x38, 0x49, 0xff, 0xeb, 0xa0, 0x7a, 0x49, 0x1c, 0xb1, 0x85, 0x22, 0x2c,
	0xe4, 0x3a, 0x37, 0xf8, 0xa9, 0x08, 0x0b, 0x22, 0x54, 0xdb, 0xbc, 0x32, 0x15, 0x9d, 0x1b, 0x9a,
	0xc9, 0x34, 0xde, 0x81, 0xe7, 0x75, 0x0e, 0x1f, 0x54, 0x9a, 0x99, 0x01, 0x09, 0x2e, 0xf2, 0x15,
	0x44, 0x38, 0x52, 0x65, 0xfe, 0x89, 0x78, 0x0f, 0x75, 0x92, 0xc8, 0x2c, 0xe4, 0x2d, 0xf8, 0x27,
	0xed, 0x7e, 0x0b, 0xf2, 0x4d, 0x75, 0xee, 0x56, 0xad, 0xdd, 0x49, 0x66, 0x61, 0xba, 0xe4, 0x6e,
	0x03, 0xbc, 0x85, 0xf2, 0x27, 0x1c, 0x74, 0x35, 0x4f, 0xa5, 0xde, 0x1d, 0xef, 0xfb, 0xc6, 0xd4,
	0xe4, 0x3b, 0x83, 0xfe, 0x11, 0xbc, 0xae, 0x48, 0xf8, 0x00, 0x5e, 0xb7, 0x25, 0x72, 0x5d, 0x39,
	0xdc, 0x3c, 0xb5, 0xe3, 0xb9, 0x74, 0xd4, 0x7c, 0xca, 0xe1, 0x6b, 0x6d, 0x6f, 0x78, 0xfe, 0x0b,
	0x00, 0x00, 0xff, 0xff, 0x84, 0x9e, 0x6e, 0x83, 0x14, 0x01, 0x00, 0x00,
}
