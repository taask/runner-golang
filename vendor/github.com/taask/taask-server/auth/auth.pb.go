// Code generated by protoc-gen-go. DO NOT EDIT.
// source: auth.proto

package auth

import (
	fmt "fmt"
	simplcrypto "github.com/cohix/simplcrypto"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Attempt struct {
	MemberUUID           string                          `protobuf:"bytes,1,opt,name=MemberUUID,proto3" json:"MemberUUID,omitempty"`
	GroupUUID            string                          `protobuf:"bytes,2,opt,name=GroupUUID,proto3" json:"GroupUUID,omitempty"`
	PubKey               *simplcrypto.SerializablePubKey `protobuf:"bytes,3,opt,name=PubKey,proto3" json:"PubKey,omitempty"`
	AuthHashSig          *simplcrypto.Signature          `protobuf:"bytes,4,opt,name=AuthHashSig,proto3" json:"AuthHashSig,omitempty"`
	Timestamp            int64                           `protobuf:"varint,5,opt,name=Timestamp,proto3" json:"Timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *Attempt) Reset()         { *m = Attempt{} }
func (m *Attempt) String() string { return proto.CompactTextString(m) }
func (*Attempt) ProtoMessage()    {}
func (*Attempt) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{0}
}

func (m *Attempt) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Attempt.Unmarshal(m, b)
}
func (m *Attempt) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Attempt.Marshal(b, m, deterministic)
}
func (m *Attempt) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Attempt.Merge(m, src)
}
func (m *Attempt) XXX_Size() int {
	return xxx_messageInfo_Attempt.Size(m)
}
func (m *Attempt) XXX_DiscardUnknown() {
	xxx_messageInfo_Attempt.DiscardUnknown(m)
}

var xxx_messageInfo_Attempt proto.InternalMessageInfo

func (m *Attempt) GetMemberUUID() string {
	if m != nil {
		return m.MemberUUID
	}
	return ""
}

func (m *Attempt) GetGroupUUID() string {
	if m != nil {
		return m.GroupUUID
	}
	return ""
}

func (m *Attempt) GetPubKey() *simplcrypto.SerializablePubKey {
	if m != nil {
		return m.PubKey
	}
	return nil
}

func (m *Attempt) GetAuthHashSig() *simplcrypto.Signature {
	if m != nil {
		return m.AuthHashSig
	}
	return nil
}

func (m *Attempt) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

type Session struct {
	MemberUUID           string                 `protobuf:"bytes,1,opt,name=MemberUUID,proto3" json:"MemberUUID,omitempty"`
	GroupUUID            string                 `protobuf:"bytes,2,opt,name=GroupUUID,proto3" json:"GroupUUID,omitempty"`
	SessionChallengeSig  *simplcrypto.Signature `protobuf:"bytes,3,opt,name=SessionChallengeSig,proto3" json:"SessionChallengeSig,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *Session) Reset()         { *m = Session{} }
func (m *Session) String() string { return proto.CompactTextString(m) }
func (*Session) ProtoMessage()    {}
func (*Session) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{1}
}

func (m *Session) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Session.Unmarshal(m, b)
}
func (m *Session) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Session.Marshal(b, m, deterministic)
}
func (m *Session) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Session.Merge(m, src)
}
func (m *Session) XXX_Size() int {
	return xxx_messageInfo_Session.Size(m)
}
func (m *Session) XXX_DiscardUnknown() {
	xxx_messageInfo_Session.DiscardUnknown(m)
}

var xxx_messageInfo_Session proto.InternalMessageInfo

func (m *Session) GetMemberUUID() string {
	if m != nil {
		return m.MemberUUID
	}
	return ""
}

func (m *Session) GetGroupUUID() string {
	if m != nil {
		return m.GroupUUID
	}
	return ""
}

func (m *Session) GetSessionChallengeSig() *simplcrypto.Signature {
	if m != nil {
		return m.SessionChallengeSig
	}
	return nil
}

type MemberGroup struct {
	UUID                 string   `protobuf:"bytes,1,opt,name=UUID,proto3" json:"UUID,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	JoinCode             string   `protobuf:"bytes,3,opt,name=JoinCode,proto3" json:"JoinCode,omitempty"`
	AuthHash             []byte   `protobuf:"bytes,4,opt,name=AuthHash,proto3" json:"AuthHash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MemberGroup) Reset()         { *m = MemberGroup{} }
func (m *MemberGroup) String() string { return proto.CompactTextString(m) }
func (*MemberGroup) ProtoMessage()    {}
func (*MemberGroup) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{2}
}

func (m *MemberGroup) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MemberGroup.Unmarshal(m, b)
}
func (m *MemberGroup) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MemberGroup.Marshal(b, m, deterministic)
}
func (m *MemberGroup) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MemberGroup.Merge(m, src)
}
func (m *MemberGroup) XXX_Size() int {
	return xxx_messageInfo_MemberGroup.Size(m)
}
func (m *MemberGroup) XXX_DiscardUnknown() {
	xxx_messageInfo_MemberGroup.DiscardUnknown(m)
}

var xxx_messageInfo_MemberGroup proto.InternalMessageInfo

func (m *MemberGroup) GetUUID() string {
	if m != nil {
		return m.UUID
	}
	return ""
}

func (m *MemberGroup) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *MemberGroup) GetJoinCode() string {
	if m != nil {
		return m.JoinCode
	}
	return ""
}

func (m *MemberGroup) GetAuthHash() []byte {
	if m != nil {
		return m.AuthHash
	}
	return nil
}

func init() {
	proto.RegisterType((*Attempt)(nil), "taask.server.auth.Attempt")
	proto.RegisterType((*Session)(nil), "taask.server.auth.Session")
	proto.RegisterType((*MemberGroup)(nil), "taask.server.auth.MemberGroup")
}

func init() { proto.RegisterFile("auth.proto", fileDescriptor_8bbd6f3875b0e874) }

var fileDescriptor_8bbd6f3875b0e874 = []byte{
	// 348 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x52, 0x41, 0x4b, 0xc3, 0x30,
	0x18, 0xa5, 0x6e, 0x6e, 0x2e, 0xf3, 0x62, 0x3c, 0x58, 0x86, 0x48, 0x29, 0x1e, 0x76, 0x31, 0x15,
	0xbd, 0xe8, 0x71, 0x4e, 0xd0, 0x29, 0x8a, 0x74, 0xee, 0xe2, 0x2d, 0x9d, 0x1f, 0x6d, 0x58, 0xd3,
	0x84, 0x24, 0x15, 0xe7, 0xcf, 0xf0, 0xe7, 0xf9, 0x6b, 0xa4, 0x49, 0xdd, 0x8a, 0x88, 0x08, 0x5e,
	0x4a, 0xf2, 0xde, 0xf7, 0xbe, 0xf7, 0x1a, 0x1e, 0x42, 0xb4, 0x34, 0x19, 0x91, 0x4a, 0x18, 0x81,
	0x77, 0x0c, 0xa5, 0x7a, 0x41, 0x34, 0xa8, 0x17, 0x50, 0xa4, 0x22, 0x06, 0xc7, 0x29, 0x33, 0x59,
	0x99, 0x90, 0xb9, 0xe0, 0xd1, 0x5c, 0x64, 0xec, 0x35, 0xd2, 0x8c, 0xcb, 0x7c, 0xae, 0x96, 0xd2,
	0x88, 0xc8, 0xca, 0x22, 0x0e, 0x5a, 0xd3, 0x14, 0xdc, 0x92, 0x3f, 0x29, 0x16, 0xb0, 0x94, 0x94,
	0x29, 0xa7, 0x08, 0x3f, 0x3c, 0xd4, 0x1d, 0x19, 0x03, 0x5c, 0x1a, 0x7c, 0x80, 0xd0, 0x1d, 0xf0,
	0x04, 0xd4, 0x6c, 0x36, 0xb9, 0xf4, 0xbd, 0xc0, 0x1b, 0xf6, 0xe2, 0x06, 0x82, 0xf7, 0x51, 0xef,
	0x4a, 0x89, 0x52, 0x5a, 0x7a, 0xc3, 0xd2, 0x6b, 0x00, 0x9f, 0xa1, 0xce, 0x43, 0x99, 0xdc, 0xc2,
	0xd2, 0x6f, 0x05, 0xde, 0xb0, 0x7f, 0x12, 0x10, 0xeb, 0x4d, 0x9c, 0x39, 0x99, 0x82, 0x62, 0x34,
	0x67, 0x6f, 0x34, 0xc9, 0xc1, 0xcd, 0xc5, 0xf5, 0x3c, 0x3e, 0x47, 0xfd, 0x51, 0x69, 0xb2, 0x6b,
	0xaa, 0xb3, 0x29, 0x4b, 0xfd, 0xb6, 0x95, 0xef, 0x7d, 0x93, 0xb3, 0xb4, 0xa0, 0xa6, 0x54, 0x10,
	0x37, 0x67, 0xab, 0x48, 0x8f, 0x8c, 0x83, 0x36, 0x94, 0x4b, 0x7f, 0x33, 0xf0, 0x86, 0xad, 0x78,
	0x0d, 0x84, 0xef, 0x1e, 0xea, 0x4e, 0x41, 0x6b, 0x26, 0x8a, 0x7f, 0xfe, 0xdc, 0x04, 0xed, 0xd6,
	0x8b, 0xc6, 0x19, 0xcd, 0x73, 0x28, 0x52, 0xa8, 0xa2, 0xb6, 0x7e, 0x8f, 0xfa, 0x93, 0x26, 0xe4,
	0xa8, 0xef, 0x6c, 0xed, 0x76, 0x8c, 0x51, 0xbb, 0x91, 0xc8, 0x9e, 0x2b, 0xec, 0x9e, 0x72, 0xa8,
	0x63, 0xd8, 0x33, 0x1e, 0xa0, 0xad, 0x1b, 0xc1, 0x8a, 0xb1, 0x78, 0x06, 0x6b, 0xdb, 0x8b, 0x57,
	0xf7, 0x8a, 0xfb, 0x7a, 0x14, 0xfb, 0x7a, 0xdb, 0xf1, 0xea, 0x7e, 0x71, 0xf8, 0x14, 0x36, 0x4a,
	0x61, 0x4b, 0xe6, 0xbe, 0x47, 0xae, 0x6a, 0x51, 0x55, 0xb5, 0xa4, 0x63, 0xdb, 0x70, 0xfa, 0x19,
	0x00, 0x00, 0xff, 0xff, 0x0f, 0xe5, 0xa0, 0xc4, 0x92, 0x02, 0x00, 0x00,
}
