// Code generated by protoc-gen-go. DO NOT EDIT.
// source: message.proto

/*
Package message is a generated protocol buffer package.

It is generated from these files:
	message.proto

It has these top-level messages:
	Signed
	Packet
	Message
	Address
	Registration
	VerificationKey
	EncryptionKey
	DecryptionKey
	Hash
	Signature
	Transaction
	Blame
	Invalid
	Packets
*/
package message

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

type Phase int32

const (
	Phase_NONE                        Phase = 0
	Phase_ANNOUNCEMENT                Phase = 1
	Phase_SHUFFLE                     Phase = 2
	Phase_BROADCAST                   Phase = 3
	Phase_EQUIVOCATION_CHECK          Phase = 4
	Phase_SIGNING                     Phase = 5
	Phase_VERIFICATION_AND_SUBMISSION Phase = 6
	Phase_BLAME                       Phase = 7
)

var Phase_name = map[int32]string{
	0: "NONE",
	1: "ANNOUNCEMENT",
	2: "SHUFFLE",
	3: "BROADCAST",
	4: "EQUIVOCATION_CHECK",
	5: "SIGNING",
	6: "VERIFICATION_AND_SUBMISSION",
	7: "BLAME",
}
var Phase_value = map[string]int32{
	"NONE":                        0,
	"ANNOUNCEMENT":                1,
	"SHUFFLE":                     2,
	"BROADCAST":                   3,
	"EQUIVOCATION_CHECK":          4,
	"SIGNING":                     5,
	"VERIFICATION_AND_SUBMISSION": 6,
	"BLAME": 7,
}

func (x Phase) String() string {
	return proto.EnumName(Phase_name, int32(x))
}
func (Phase) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Reason int32

const (
	Reason_INSUFFICIENTFUNDS             Reason = 0
	Reason_DOUBLESPEND                   Reason = 1
	Reason_EQUIVOCATIONFAILURE           Reason = 2
	Reason_SHUFFLEFAILURE                Reason = 3
	Reason_SHUFFLEANDEQUIVOCATIONFAILURE Reason = 4
	Reason_INVALIDSIGNATURE              Reason = 5
	Reason_MISSINGOUTPUT                 Reason = 6
	Reason_LIAR                          Reason = 7
	Reason_INVALIDFORMAT                 Reason = 8
)

var Reason_name = map[int32]string{
	0: "INSUFFICIENTFUNDS",
	1: "DOUBLESPEND",
	2: "EQUIVOCATIONFAILURE",
	3: "SHUFFLEFAILURE",
	4: "SHUFFLEANDEQUIVOCATIONFAILURE",
	5: "INVALIDSIGNATURE",
	6: "MISSINGOUTPUT",
	7: "LIAR",
	8: "INVALIDFORMAT",
}
var Reason_value = map[string]int32{
	"INSUFFICIENTFUNDS":             0,
	"DOUBLESPEND":                   1,
	"EQUIVOCATIONFAILURE":           2,
	"SHUFFLEFAILURE":                3,
	"SHUFFLEANDEQUIVOCATIONFAILURE": 4,
	"INVALIDSIGNATURE":              5,
	"MISSINGOUTPUT":                 6,
	"LIAR":                          7,
	"INVALIDFORMAT":                 8,
}

func (x Reason) String() string {
	return proto.EnumName(Reason_name, int32(x))
}
func (Reason) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type Signed struct {
	Packet    *Packet    `protobuf:"bytes,1,opt,name=packet" json:"packet,omitempty"`
	Signature *Signature `protobuf:"bytes,2,opt,name=signature" json:"signature,omitempty"`
}

func (m *Signed) Reset()                    { *m = Signed{} }
func (m *Signed) String() string            { return proto.CompactTextString(m) }
func (*Signed) ProtoMessage()               {}
func (*Signed) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Signed) GetPacket() *Packet {
	if m != nil {
		return m.Packet
	}
	return nil
}

func (m *Signed) GetSignature() *Signature {
	if m != nil {
		return m.Signature
	}
	return nil
}

type Packet struct {
	Session      []byte           `protobuf:"bytes,1,opt,name=session,proto3" json:"session,omitempty"`
	Number       uint32           `protobuf:"varint,2,opt,name=number" json:"number,omitempty"`
	FromKey      *VerificationKey `protobuf:"bytes,3,opt,name=from_key,json=fromKey" json:"from_key,omitempty"`
	ToKey        *VerificationKey `protobuf:"bytes,4,opt,name=to_key,json=toKey" json:"to_key,omitempty"`
	Phase        Phase            `protobuf:"varint,5,opt,name=phase,enum=Phase" json:"phase,omitempty"`
	Message      *Message         `protobuf:"bytes,6,opt,name=message" json:"message,omitempty"`
	Registration *Registration    `protobuf:"bytes,7,opt,name=registration" json:"registration,omitempty"`
}

func (m *Packet) Reset()                    { *m = Packet{} }
func (m *Packet) String() string            { return proto.CompactTextString(m) }
func (*Packet) ProtoMessage()               {}
func (*Packet) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Packet) GetSession() []byte {
	if m != nil {
		return m.Session
	}
	return nil
}

func (m *Packet) GetNumber() uint32 {
	if m != nil {
		return m.Number
	}
	return 0
}

func (m *Packet) GetFromKey() *VerificationKey {
	if m != nil {
		return m.FromKey
	}
	return nil
}

func (m *Packet) GetToKey() *VerificationKey {
	if m != nil {
		return m.ToKey
	}
	return nil
}

func (m *Packet) GetPhase() Phase {
	if m != nil {
		return m.Phase
	}
	return Phase_NONE
}

func (m *Packet) GetMessage() *Message {
	if m != nil {
		return m.Message
	}
	return nil
}

func (m *Packet) GetRegistration() *Registration {
	if m != nil {
		return m.Registration
	}
	return nil
}

type Message struct {
	Address   *Address       `protobuf:"bytes,1,opt,name=address" json:"address,omitempty"`
	Key       *EncryptionKey `protobuf:"bytes,2,opt,name=key" json:"key,omitempty"`
	Hash      *Hash          `protobuf:"bytes,3,opt,name=hash" json:"hash,omitempty"`
	Signature *Signature     `protobuf:"bytes,4,opt,name=signature" json:"signature,omitempty"`
	Str       string         `protobuf:"bytes,5,opt,name=str" json:"str,omitempty"`
	Blame     *Blame         `protobuf:"bytes,6,opt,name=blame" json:"blame,omitempty"`
	Next      *Message       `protobuf:"bytes,7,opt,name=next" json:"next,omitempty"`
}

func (m *Message) Reset()                    { *m = Message{} }
func (m *Message) String() string            { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()               {}
func (*Message) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Message) GetAddress() *Address {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *Message) GetKey() *EncryptionKey {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *Message) GetHash() *Hash {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *Message) GetSignature() *Signature {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *Message) GetStr() string {
	if m != nil {
		return m.Str
	}
	return ""
}

func (m *Message) GetBlame() *Blame {
	if m != nil {
		return m.Blame
	}
	return nil
}

func (m *Message) GetNext() *Message {
	if m != nil {
		return m.Next
	}
	return nil
}

type Address struct {
	Address string `protobuf:"bytes,1,opt,name=address" json:"address,omitempty"`
}

func (m *Address) Reset()                    { *m = Address{} }
func (m *Address) String() string            { return proto.CompactTextString(m) }
func (*Address) ProtoMessage()               {}
func (*Address) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Address) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type Registration struct {
	Amount uint64 `protobuf:"varint,1,opt,name=amount" json:"amount,omitempty"`
}

func (m *Registration) Reset()                    { *m = Registration{} }
func (m *Registration) String() string            { return proto.CompactTextString(m) }
func (*Registration) ProtoMessage()               {}
func (*Registration) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Registration) GetAmount() uint64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

type VerificationKey struct {
	Key string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
}

func (m *VerificationKey) Reset()                    { *m = VerificationKey{} }
func (m *VerificationKey) String() string            { return proto.CompactTextString(m) }
func (*VerificationKey) ProtoMessage()               {}
func (*VerificationKey) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *VerificationKey) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type EncryptionKey struct {
	Key string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
}

func (m *EncryptionKey) Reset()                    { *m = EncryptionKey{} }
func (m *EncryptionKey) String() string            { return proto.CompactTextString(m) }
func (*EncryptionKey) ProtoMessage()               {}
func (*EncryptionKey) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *EncryptionKey) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type DecryptionKey struct {
	Key    string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Public string `protobuf:"bytes,2,opt,name=public" json:"public,omitempty"`
}

func (m *DecryptionKey) Reset()                    { *m = DecryptionKey{} }
func (m *DecryptionKey) String() string            { return proto.CompactTextString(m) }
func (*DecryptionKey) ProtoMessage()               {}
func (*DecryptionKey) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *DecryptionKey) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *DecryptionKey) GetPublic() string {
	if m != nil {
		return m.Public
	}
	return ""
}

type Hash struct {
	Hash []byte `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
}

func (m *Hash) Reset()                    { *m = Hash{} }
func (m *Hash) String() string            { return proto.CompactTextString(m) }
func (*Hash) ProtoMessage()               {}
func (*Hash) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *Hash) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

type Signature struct {
	Signature []byte `protobuf:"bytes,1,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (m *Signature) Reset()                    { *m = Signature{} }
func (m *Signature) String() string            { return proto.CompactTextString(m) }
func (*Signature) ProtoMessage()               {}
func (*Signature) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *Signature) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

type Transaction struct {
	Transaction []byte `protobuf:"bytes,1,opt,name=transaction,proto3" json:"transaction,omitempty"`
}

func (m *Transaction) Reset()                    { *m = Transaction{} }
func (m *Transaction) String() string            { return proto.CompactTextString(m) }
func (*Transaction) ProtoMessage()               {}
func (*Transaction) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *Transaction) GetTransaction() []byte {
	if m != nil {
		return m.Transaction
	}
	return nil
}

type Blame struct {
	Reason      Reason           `protobuf:"varint,1,opt,name=reason,enum=Reason" json:"reason,omitempty"`
	Accused     *VerificationKey `protobuf:"bytes,2,opt,name=accused" json:"accused,omitempty"`
	Key         *DecryptionKey   `protobuf:"bytes,3,opt,name=key" json:"key,omitempty"`
	Transaction *Transaction     `protobuf:"bytes,4,opt,name=transaction" json:"transaction,omitempty"`
	Invalid     *Invalid         `protobuf:"bytes,5,opt,name=invalid" json:"invalid,omitempty"`
	Packets     *Packets         `protobuf:"bytes,6,opt,name=packets" json:"packets,omitempty"`
}

func (m *Blame) Reset()                    { *m = Blame{} }
func (m *Blame) String() string            { return proto.CompactTextString(m) }
func (*Blame) ProtoMessage()               {}
func (*Blame) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *Blame) GetReason() Reason {
	if m != nil {
		return m.Reason
	}
	return Reason_INSUFFICIENTFUNDS
}

func (m *Blame) GetAccused() *VerificationKey {
	if m != nil {
		return m.Accused
	}
	return nil
}

func (m *Blame) GetKey() *DecryptionKey {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *Blame) GetTransaction() *Transaction {
	if m != nil {
		return m.Transaction
	}
	return nil
}

func (m *Blame) GetInvalid() *Invalid {
	if m != nil {
		return m.Invalid
	}
	return nil
}

func (m *Blame) GetPackets() *Packets {
	if m != nil {
		return m.Packets
	}
	return nil
}

type Invalid struct {
	Invalid []byte `protobuf:"bytes,1,opt,name=invalid,proto3" json:"invalid,omitempty"`
}

func (m *Invalid) Reset()                    { *m = Invalid{} }
func (m *Invalid) String() string            { return proto.CompactTextString(m) }
func (*Invalid) ProtoMessage()               {}
func (*Invalid) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *Invalid) GetInvalid() []byte {
	if m != nil {
		return m.Invalid
	}
	return nil
}

type Packets struct {
	Packet []*Signed `protobuf:"bytes,1,rep,name=packet" json:"packet,omitempty"`
}

func (m *Packets) Reset()                    { *m = Packets{} }
func (m *Packets) String() string            { return proto.CompactTextString(m) }
func (*Packets) ProtoMessage()               {}
func (*Packets) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *Packets) GetPacket() []*Signed {
	if m != nil {
		return m.Packet
	}
	return nil
}

func init() {
	proto.RegisterType((*Signed)(nil), "Signed")
	proto.RegisterType((*Packet)(nil), "Packet")
	proto.RegisterType((*Message)(nil), "Message")
	proto.RegisterType((*Address)(nil), "Address")
	proto.RegisterType((*Registration)(nil), "Registration")
	proto.RegisterType((*VerificationKey)(nil), "VerificationKey")
	proto.RegisterType((*EncryptionKey)(nil), "EncryptionKey")
	proto.RegisterType((*DecryptionKey)(nil), "DecryptionKey")
	proto.RegisterType((*Hash)(nil), "Hash")
	proto.RegisterType((*Signature)(nil), "Signature")
	proto.RegisterType((*Transaction)(nil), "Transaction")
	proto.RegisterType((*Blame)(nil), "Blame")
	proto.RegisterType((*Invalid)(nil), "Invalid")
	proto.RegisterType((*Packets)(nil), "Packets")
	proto.RegisterEnum("Phase", Phase_name, Phase_value)
	proto.RegisterEnum("Reason", Reason_name, Reason_value)
}

func init() { proto.RegisterFile("message.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 821 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x55, 0xdd, 0x6e, 0xe3, 0x44,
	0x14, 0x5e, 0x37, 0xfe, 0x49, 0x4e, 0x92, 0xee, 0xec, 0x00, 0x4b, 0x80, 0xa2, 0xed, 0x7a, 0x25,
	0x28, 0x45, 0x32, 0xa2, 0x5c, 0x71, 0xe9, 0x24, 0x4e, 0x3b, 0x6a, 0x32, 0x2e, 0x63, 0xbb, 0xb7,
	0x95, 0x9b, 0xcc, 0xb6, 0xd6, 0x36, 0x4e, 0xe4, 0x71, 0x10, 0x79, 0x00, 0xee, 0x79, 0x2a, 0x1e,
	0x84, 0x07, 0xe0, 0x19, 0xd0, 0xfc, 0xb8, 0x4d, 0x50, 0x77, 0xef, 0xfc, 0x7d, 0xe7, 0x9b, 0x93,
	0x73, 0xbe, 0x39, 0x73, 0x02, 0xfd, 0x25, 0x17, 0x22, 0xbf, 0xe3, 0xc1, 0xba, 0x5a, 0xd5, 0x2b,
	0x3f, 0x01, 0x37, 0x29, 0xee, 0x4a, 0xbe, 0xc0, 0x6f, 0xc0, 0x5d, 0xe7, 0xf3, 0x0f, 0xbc, 0x1e,
	0x58, 0xc7, 0xd6, 0x49, 0xf7, 0xcc, 0x0b, 0xae, 0x14, 0x64, 0x86, 0xc6, 0x27, 0xd0, 0x11, 0xc5,
	0x5d, 0x99, 0xd7, 0x9b, 0x8a, 0x0f, 0x0e, 0x94, 0x06, 0x82, 0xa4, 0x61, 0xd8, 0x53, 0xd0, 0xff,
	0xf3, 0x00, 0x5c, 0x7d, 0x18, 0x0f, 0xc0, 0x13, 0x5c, 0x88, 0x62, 0x55, 0xaa, 0xb4, 0x3d, 0xd6,
	0x40, 0xfc, 0x1a, 0xdc, 0x72, 0xb3, 0xbc, 0xe5, 0x95, 0xca, 0xd5, 0x67, 0x06, 0xe1, 0x1f, 0xa1,
	0xfd, 0xbe, 0x5a, 0x2d, 0x6f, 0x3e, 0xf0, 0xed, 0xa0, 0xa5, 0x7e, 0x05, 0x05, 0xd7, 0xbc, 0x2a,
	0xde, 0x17, 0xf3, 0xbc, 0x2e, 0x56, 0xe5, 0x25, 0xdf, 0x32, 0x4f, 0x2a, 0x2e, 0xf9, 0x16, 0x7f,
	0x0f, 0x6e, 0xbd, 0x52, 0x52, 0xfb, 0x23, 0x52, 0xa7, 0x5e, 0x49, 0xe1, 0x11, 0x38, 0xeb, 0xfb,
	0x5c, 0xf0, 0x81, 0x73, 0x6c, 0x9d, 0x1c, 0x9e, 0xb9, 0xc1, 0x95, 0x44, 0x4c, 0x93, 0xd8, 0x07,
	0xcf, 0xd8, 0x32, 0x70, 0x55, 0x9e, 0x76, 0x30, 0xd3, 0x98, 0x35, 0x01, 0xfc, 0x33, 0xf4, 0x2a,
	0x7e, 0x57, 0x88, 0xba, 0x52, 0xb9, 0x07, 0x9e, 0x12, 0xf6, 0x03, 0xb6, 0x43, 0xb2, 0x3d, 0x89,
	0xff, 0x8f, 0x05, 0x9e, 0xc9, 0x23, 0x7f, 0x22, 0x5f, 0x2c, 0x2a, 0x2e, 0x84, 0xf1, 0xb7, 0x1d,
	0x84, 0x1a, 0xb3, 0x26, 0x80, 0x8f, 0xa1, 0x25, 0x5b, 0xd1, 0xde, 0x1e, 0x06, 0x51, 0x39, 0xaf,
	0xb6, 0xeb, 0xa6, 0x11, 0x19, 0xc2, 0x5f, 0x81, 0x7d, 0x9f, 0x8b, 0x7b, 0x63, 0x8c, 0x13, 0x5c,
	0xe4, 0xe2, 0x9e, 0x29, 0x6a, 0xff, 0x7a, 0xec, 0x4f, 0x5c, 0x0f, 0x46, 0xd0, 0x12, 0x75, 0xa5,
	0x9c, 0xe8, 0x30, 0xf9, 0x29, 0xdd, 0xb9, 0x7d, 0xc8, 0x97, 0x4d, 0xf7, 0x6e, 0x30, 0x94, 0x88,
	0x69, 0x12, 0x1f, 0x81, 0x5d, 0xf2, 0x3f, 0x6a, 0xd3, 0xf1, 0x93, 0x35, 0x8a, 0xf5, 0xdf, 0x81,
	0x67, 0x1a, 0x91, 0x97, 0xbd, 0xdb, 0x63, 0xe7, 0xb1, 0x33, 0xff, 0x3b, 0xe8, 0xed, 0xfa, 0x24,
	0x2f, 0x3f, 0x5f, 0xae, 0x36, 0xa5, 0x1e, 0x36, 0x9b, 0x19, 0xe4, 0xbf, 0x83, 0x97, 0xff, 0xbb,
	0x40, 0x59, 0xad, 0x34, 0x45, 0x27, 0x94, 0x9f, 0xfe, 0x5b, 0xe8, 0xef, 0x59, 0xf3, 0x8c, 0xe4,
	0x57, 0xe8, 0x8f, 0xf9, 0x27, 0x25, 0xb2, 0x84, 0xf5, 0xe6, 0xf6, 0xa1, 0x98, 0x2b, 0xbf, 0x3b,
	0xcc, 0x20, 0xff, 0x6b, 0xb0, 0xa5, 0xab, 0x18, 0x1b, 0xab, 0xf5, 0xd8, 0xaa, 0x6f, 0xff, 0x07,
	0xe8, 0x3c, 0x3a, 0x8a, 0x8f, 0x76, 0x0d, 0xd7, 0xaa, 0x9d, 0x37, 0xf0, 0x13, 0x74, 0xd3, 0x2a,
	0x2f, 0x45, 0x3e, 0x57, 0x0d, 0x1f, 0x43, 0xb7, 0x7e, 0x82, 0x46, 0xbe, 0x4b, 0xf9, 0xff, 0x5a,
	0xe0, 0x28, 0xdb, 0xe5, 0x4b, 0xac, 0x78, 0x2e, 0x8c, 0xec, 0xf0, 0xcc, 0x0b, 0x98, 0x82, 0xcc,
	0xd0, 0xf8, 0x14, 0xbc, 0x7c, 0x3e, 0xdf, 0x08, 0xbe, 0x30, 0xb3, 0xf2, 0xcc, 0x0b, 0x31, 0x82,
	0x66, 0xa6, 0x5a, 0x66, 0xa6, 0xf6, 0x5c, 0xd1, 0x46, 0x04, 0xfb, 0xa5, 0xe9, 0xd1, 0xe9, 0x05,
	0x3b, 0xd5, 0xef, 0x15, 0x2a, 0x27, 0xb9, 0x28, 0x7f, 0xcf, 0x1f, 0x8a, 0x85, 0x1a, 0x21, 0x39,
	0x11, 0x44, 0x63, 0xd6, 0x04, 0xa4, 0x46, 0x6f, 0x0d, 0xf1, 0xf8, 0xa0, 0xf4, 0x42, 0x10, 0xac,
	0x09, 0xc8, 0xc1, 0x31, 0xe7, 0xe4, 0xe0, 0x34, 0x29, 0xcd, 0x96, 0x30, 0xd0, 0x3f, 0x05, 0xcf,
	0x1c, 0xdc, 0x5b, 0x50, 0x2d, 0xb5, 0xa0, 0xf4, 0xe6, 0x6a, 0x16, 0xd4, 0xe9, 0x5f, 0x16, 0x38,
	0xea, 0x59, 0xe3, 0x36, 0xd8, 0x34, 0xa6, 0x11, 0x7a, 0x81, 0x11, 0xf4, 0x42, 0x4a, 0xe3, 0x8c,
	0x8e, 0xa2, 0x59, 0x44, 0x53, 0x64, 0xe1, 0x2e, 0x78, 0xc9, 0x45, 0x36, 0x99, 0x4c, 0x23, 0x74,
	0x80, 0xfb, 0xd0, 0x19, 0xb2, 0x38, 0x1c, 0x8f, 0xc2, 0x24, 0x45, 0x2d, 0xfc, 0x1a, 0x70, 0xf4,
	0x5b, 0x46, 0xae, 0xe3, 0x51, 0x98, 0x92, 0x98, 0xde, 0x8c, 0x2e, 0xa2, 0xd1, 0x25, 0xb2, 0xd5,
	0x19, 0x72, 0x4e, 0x09, 0x3d, 0x47, 0x0e, 0x7e, 0x03, 0xdf, 0x5c, 0x47, 0x8c, 0x4c, 0x88, 0x11,
	0x85, 0x74, 0x7c, 0x93, 0x64, 0xc3, 0x19, 0x49, 0x12, 0x12, 0x53, 0xe4, 0xe2, 0x0e, 0x38, 0xc3,
	0x69, 0x38, 0x8b, 0x90, 0x77, 0xfa, 0xb7, 0x05, 0xae, 0xbe, 0x3c, 0xfc, 0x05, 0xbc, 0x22, 0x34,
	0xc9, 0x26, 0x13, 0x32, 0x22, 0x11, 0x4d, 0x27, 0x19, 0x1d, 0x27, 0xe8, 0x05, 0x7e, 0x09, 0xdd,
	0x71, 0x9c, 0x0d, 0xa7, 0x51, 0x72, 0x15, 0xd1, 0x31, 0xb2, 0xf0, 0x97, 0xf0, 0xd9, 0x6e, 0x0d,
	0x93, 0x90, 0x4c, 0x33, 0x26, 0x6b, 0xc5, 0x70, 0x68, 0x0a, 0x6f, 0xb8, 0x16, 0x7e, 0x0b, 0xdf,
	0x1a, 0x2e, 0xa4, 0xe3, 0xe7, 0x8e, 0xd9, 0xf8, 0x73, 0x40, 0x84, 0x5e, 0x87, 0x53, 0x32, 0x96,
	0x2d, 0x84, 0xa9, 0x64, 0x1d, 0xfc, 0x0a, 0xfa, 0xaa, 0x60, 0x7a, 0x1e, 0x67, 0xe9, 0x55, 0x96,
	0x22, 0x57, 0x9a, 0x36, 0x25, 0x21, 0x43, 0x9e, 0x0c, 0x9a, 0x23, 0x93, 0x98, 0xcd, 0xc2, 0x14,
	0xb5, 0x6f, 0x5d, 0xf5, 0x77, 0xf1, 0xcb, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x32, 0x0c, 0x0d,
	0x44, 0x3f, 0x06, 0x00, 0x00,
}
