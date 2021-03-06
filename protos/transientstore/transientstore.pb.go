// Code generated by protoc-gen-go. DO NOT EDIT.
// source: transientstore/transientstore.proto

/*
Package transientstore is a generated protocol buffer package.

It is generated from these files:
	transientstore/transientstore.proto

It has these top-level messages:
	TxPvtReadWriteSetWithConfigInfo
*/
package transientstore

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import rwset "github.com/hyperledger/fabric/protos/ledger/rwset"
import common2 "github.com/hyperledger/fabric/protos/common"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// TxPvtReadWriteSetWithConfigInfo encapsulates the transaction's private
// read-write set and additional information about the configurations such as
// the latest collection config when the transaction is simulated
type TxPvtReadWriteSetWithConfigInfo struct {
	PvtRwset          *rwset.TxPvtReadWriteSet                    `protobuf:"bytes,1,opt,name=pvt_rwset,json=pvtRwset" json:"pvt_rwset,omitempty"`
	CollectionConfigs map[string]*common2.CollectionConfigPackage `protobuf:"bytes,2,rep,name=collection_configs,json=collectionConfigs" json:"collection_configs,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *TxPvtReadWriteSetWithConfigInfo) Reset()                    { *m = TxPvtReadWriteSetWithConfigInfo{} }
func (m *TxPvtReadWriteSetWithConfigInfo) String() string            { return proto.CompactTextString(m) }
func (*TxPvtReadWriteSetWithConfigInfo) ProtoMessage()               {}
func (*TxPvtReadWriteSetWithConfigInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *TxPvtReadWriteSetWithConfigInfo) GetPvtRwset() *rwset.TxPvtReadWriteSet {
	if m != nil {
		return m.PvtRwset
	}
	return nil
}

func (m *TxPvtReadWriteSetWithConfigInfo) GetCollectionConfigs() map[string]*common2.CollectionConfigPackage {
	if m != nil {
		return m.CollectionConfigs
	}
	return nil
}

func init() {
	proto.RegisterType((*TxPvtReadWriteSetWithConfigInfo)(nil), "transientstore.TxPvtReadWriteSetWithConfigInfo")
}

func init() { proto.RegisterFile("transientstore/transientstore.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 299 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x51, 0xc1, 0x4b, 0xfb, 0x30,
	0x18, 0x65, 0x1d, 0xbf, 0x1f, 0x2e, 0x03, 0xd1, 0x1c, 0xb4, 0xec, 0xb2, 0xa1, 0x97, 0x1d, 0x24,
	0x81, 0x8d, 0x81, 0x78, 0x74, 0x28, 0x78, 0x1b, 0x51, 0x18, 0x78, 0x19, 0x59, 0xf6, 0xad, 0x0d,
	0xeb, 0x92, 0x92, 0x7c, 0xad, 0xf6, 0x1f, 0xf3, 0xef, 0x93, 0x36, 0xa0, 0xae, 0x13, 0xbc, 0x94,
	0xe6, 0x7d, 0xdf, 0x7b, 0x2f, 0xef, 0x85, 0x5c, 0xa3, 0x93, 0xc6, 0x6b, 0x30, 0xe8, 0xd1, 0x3a,
	0xe0, 0x87, 0x47, 0x96, 0x3b, 0x8b, 0x96, 0x9e, 0x1e, 0xa2, 0x83, 0x38, 0x83, 0x4d, 0x02, 0x8e,
	0xbb, 0x37, 0x0f, 0x18, 0xbe, 0x61, 0x73, 0x70, 0xa9, 0xec, 0x7e, 0x6f, 0x0d, 0x57, 0x36, 0xcb,
	0x40, 0xa1, 0xb6, 0x26, 0x0c, 0xae, 0x3e, 0x22, 0x32, 0x7c, 0x79, 0x5f, 0x94, 0x28, 0x40, 0x6e,
	0x96, 0x4e, 0x23, 0x3c, 0x03, 0x2e, 0x35, 0xa6, 0x73, 0x6b, 0xb6, 0x3a, 0x79, 0x32, 0x5b, 0x4b,
	0x67, 0xa4, 0x97, 0x97, 0xb8, 0x6a, 0xf4, 0xe2, 0xce, 0xa8, 0x33, 0xee, 0x4f, 0x62, 0x16, 0xd4,
	0x8f, 0xa8, 0xe2, 0x24, 0x2f, 0x51, 0xd4, 0x33, 0x5a, 0x10, 0xfa, 0x6d, 0xb7, 0x52, 0x8d, 0x9e,
	0x8f, 0xa3, 0x51, 0x77, 0xdc, 0x9f, 0x3c, 0xb2, 0x56, 0xa0, 0x3f, 0xee, 0xc0, 0xe6, 0x5f, 0x4a,
	0x01, 0xf4, 0x0f, 0x06, 0x5d, 0x25, 0xce, 0x55, 0x1b, 0x1f, 0x00, 0xb9, 0xf8, 0x7d, 0x99, 0x9e,
	0x91, 0xee, 0x0e, 0xaa, 0x26, 0x41, 0x4f, 0xd4, 0xbf, 0x74, 0x46, 0xfe, 0x95, 0x32, 0x2b, 0x20,
	0x8e, 0x9a, 0x54, 0x43, 0x16, 0x6a, 0x3a, 0x72, 0x5b, 0x48, 0xb5, 0x93, 0x09, 0x88, 0xb0, 0x7d,
	0x17, 0xdd, 0x76, 0xee, 0x15, 0xb9, 0xb1, 0x2e, 0x61, 0x69, 0x95, 0x83, 0x0b, 0xb5, 0xb3, 0xad,
	0x5c, 0x3b, 0xad, 0x42, 0xb1, 0xbe, 0x15, 0xf0, 0x75, 0x9a, 0x68, 0x4c, 0x8b, 0x75, 0xed, 0xc0,
	0x7f, 0x90, 0x78, 0x20, 0xf1, 0x40, 0x6a, 0x3d, 0xf3, 0xfa, 0x7f, 0x03, 0x4f, 0x3f, 0x03, 0x00,
	0x00, 0xff, 0xff, 0x43, 0x79, 0x87, 0x0d, 0x0e, 0x02, 0x00, 0x00,
}
