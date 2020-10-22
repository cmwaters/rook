// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: rook/config.proto

package types

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type GameConfig struct {
	InitialTeams uint32                 `protobuf:"varint,1,opt,name=initialTeams,proto3" json:"initialTeams,omitempty"`
	Map          *MapConfig             `protobuf:"bytes,2,opt,name=map,proto3" json:"map,omitempty"`
	Production   *ProductionRatesConfig `protobuf:"bytes,3,opt,name=production,proto3" json:"production,omitempty"`
	Construction *SettlementCostsConfig `protobuf:"bytes,4,opt,name=construction,proto3" json:"construction,omitempty"`
}

func (m *GameConfig) Reset()         { *m = GameConfig{} }
func (m *GameConfig) String() string { return proto.CompactTextString(m) }
func (*GameConfig) ProtoMessage()    {}
func (*GameConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_3cd723d25700bbf6, []int{0}
}
func (m *GameConfig) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GameConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GameConfig.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GameConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GameConfig.Merge(m, src)
}
func (m *GameConfig) XXX_Size() int {
	return m.Size()
}
func (m *GameConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_GameConfig.DiscardUnknown(m)
}

var xxx_messageInfo_GameConfig proto.InternalMessageInfo

func (m *GameConfig) GetInitialTeams() uint32 {
	if m != nil {
		return m.InitialTeams
	}
	return 0
}

func (m *GameConfig) GetMap() *MapConfig {
	if m != nil {
		return m.Map
	}
	return nil
}

func (m *GameConfig) GetProduction() *ProductionRatesConfig {
	if m != nil {
		return m.Production
	}
	return nil
}

func (m *GameConfig) GetConstruction() *SettlementCostsConfig {
	if m != nil {
		return m.Construction
	}
	return nil
}

type MapConfig struct {
	Width  uint32 `protobuf:"varint,1,opt,name=width,proto3" json:"width,omitempty"`
	Height uint32 `protobuf:"varint,2,opt,name=height,proto3" json:"height,omitempty"`
	Seed   int64  `protobuf:"varint,3,opt,name=seed,proto3" json:"seed,omitempty"`
}

func (m *MapConfig) Reset()         { *m = MapConfig{} }
func (m *MapConfig) String() string { return proto.CompactTextString(m) }
func (*MapConfig) ProtoMessage()    {}
func (*MapConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_3cd723d25700bbf6, []int{1}
}
func (m *MapConfig) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MapConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MapConfig.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MapConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MapConfig.Merge(m, src)
}
func (m *MapConfig) XXX_Size() int {
	return m.Size()
}
func (m *MapConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_MapConfig.DiscardUnknown(m)
}

var xxx_messageInfo_MapConfig proto.InternalMessageInfo

func (m *MapConfig) GetWidth() uint32 {
	if m != nil {
		return m.Width
	}
	return 0
}

func (m *MapConfig) GetHeight() uint32 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *MapConfig) GetSeed() int64 {
	if m != nil {
		return m.Seed
	}
	return 0
}

type ProductionRatesConfig struct {
	Town       uint32 `protobuf:"varint,1,opt,name=town,proto3" json:"town,omitempty"`
	City       uint32 `protobuf:"varint,2,opt,name=city,proto3" json:"city,omitempty"`
	Lumbermill uint32 `protobuf:"varint,3,opt,name=lumbermill,proto3" json:"lumbermill,omitempty"`
	Quarry     uint32 `protobuf:"varint,4,opt,name=quarry,proto3" json:"quarry,omitempty"`
	Farm       uint32 `protobuf:"varint,5,opt,name=farm,proto3" json:"farm,omitempty"`
}

func (m *ProductionRatesConfig) Reset()         { *m = ProductionRatesConfig{} }
func (m *ProductionRatesConfig) String() string { return proto.CompactTextString(m) }
func (*ProductionRatesConfig) ProtoMessage()    {}
func (*ProductionRatesConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_3cd723d25700bbf6, []int{2}
}
func (m *ProductionRatesConfig) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ProductionRatesConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ProductionRatesConfig.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ProductionRatesConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProductionRatesConfig.Merge(m, src)
}
func (m *ProductionRatesConfig) XXX_Size() int {
	return m.Size()
}
func (m *ProductionRatesConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_ProductionRatesConfig.DiscardUnknown(m)
}

var xxx_messageInfo_ProductionRatesConfig proto.InternalMessageInfo

func (m *ProductionRatesConfig) GetTown() uint32 {
	if m != nil {
		return m.Town
	}
	return 0
}

func (m *ProductionRatesConfig) GetCity() uint32 {
	if m != nil {
		return m.City
	}
	return 0
}

func (m *ProductionRatesConfig) GetLumbermill() uint32 {
	if m != nil {
		return m.Lumbermill
	}
	return 0
}

func (m *ProductionRatesConfig) GetQuarry() uint32 {
	if m != nil {
		return m.Quarry
	}
	return 0
}

func (m *ProductionRatesConfig) GetFarm() uint32 {
	if m != nil {
		return m.Farm
	}
	return 0
}

type SettlementCostsConfig struct {
	Town       *ResourceSet `protobuf:"bytes,1,opt,name=town,proto3" json:"town,omitempty"`
	City       *ResourceSet `protobuf:"bytes,2,opt,name=city,proto3" json:"city,omitempty"`
	Capital    *ResourceSet `protobuf:"bytes,3,opt,name=capital,proto3" json:"capital,omitempty"`
	Farm       *ResourceSet `protobuf:"bytes,4,opt,name=farm,proto3" json:"farm,omitempty"`
	Lumbermill *ResourceSet `protobuf:"bytes,5,opt,name=lumbermill,proto3" json:"lumbermill,omitempty"`
	Quarry     *ResourceSet `protobuf:"bytes,6,opt,name=quarry,proto3" json:"quarry,omitempty"`
	Rook       *ResourceSet `protobuf:"bytes,7,opt,name=rook,proto3" json:"rook,omitempty"`
}

func (m *SettlementCostsConfig) Reset()         { *m = SettlementCostsConfig{} }
func (m *SettlementCostsConfig) String() string { return proto.CompactTextString(m) }
func (*SettlementCostsConfig) ProtoMessage()    {}
func (*SettlementCostsConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_3cd723d25700bbf6, []int{3}
}
func (m *SettlementCostsConfig) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SettlementCostsConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SettlementCostsConfig.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SettlementCostsConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SettlementCostsConfig.Merge(m, src)
}
func (m *SettlementCostsConfig) XXX_Size() int {
	return m.Size()
}
func (m *SettlementCostsConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_SettlementCostsConfig.DiscardUnknown(m)
}

var xxx_messageInfo_SettlementCostsConfig proto.InternalMessageInfo

func (m *SettlementCostsConfig) GetTown() *ResourceSet {
	if m != nil {
		return m.Town
	}
	return nil
}

func (m *SettlementCostsConfig) GetCity() *ResourceSet {
	if m != nil {
		return m.City
	}
	return nil
}

func (m *SettlementCostsConfig) GetCapital() *ResourceSet {
	if m != nil {
		return m.Capital
	}
	return nil
}

func (m *SettlementCostsConfig) GetFarm() *ResourceSet {
	if m != nil {
		return m.Farm
	}
	return nil
}

func (m *SettlementCostsConfig) GetLumbermill() *ResourceSet {
	if m != nil {
		return m.Lumbermill
	}
	return nil
}

func (m *SettlementCostsConfig) GetQuarry() *ResourceSet {
	if m != nil {
		return m.Quarry
	}
	return nil
}

func (m *SettlementCostsConfig) GetRook() *ResourceSet {
	if m != nil {
		return m.Rook
	}
	return nil
}

func init() {
	proto.RegisterType((*GameConfig)(nil), "rook.rook.GameConfig")
	proto.RegisterType((*MapConfig)(nil), "rook.rook.MapConfig")
	proto.RegisterType((*ProductionRatesConfig)(nil), "rook.rook.ProductionRatesConfig")
	proto.RegisterType((*SettlementCostsConfig)(nil), "rook.rook.SettlementCostsConfig")
}

func init() { proto.RegisterFile("rook/config.proto", fileDescriptor_3cd723d25700bbf6) }

var fileDescriptor_3cd723d25700bbf6 = []byte{
	// 447 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x93, 0xcf, 0x6a, 0xdb, 0x40,
	0x10, 0xc6, 0xad, 0xf8, 0x4f, 0xc8, 0x24, 0x86, 0x76, 0x49, 0x82, 0xe8, 0x41, 0x18, 0x41, 0x4b,
	0xf0, 0x41, 0x2e, 0x29, 0xf4, 0xda, 0xd2, 0x14, 0x7a, 0x0a, 0x14, 0xa5, 0xa7, 0xde, 0xd6, 0xf2,
	0xc4, 0x5e, 0xaa, 0xd5, 0xaa, 0xbb, 0x23, 0x5c, 0x3f, 0x43, 0x2f, 0x7d, 0xac, 0x1e, 0x73, 0xec,
	0x31, 0xb5, 0x5f, 0xa4, 0xec, 0x6a, 0x6d, 0x64, 0xea, 0xe8, 0x62, 0xcf, 0xce, 0x7c, 0xf3, 0xe3,
	0xdb, 0x4f, 0x2c, 0x3c, 0xd7, 0x4a, 0x7d, 0x9b, 0x64, 0xaa, 0xb8, 0x17, 0xf3, 0xa4, 0xd4, 0x8a,
	0x14, 0x3b, 0xb1, 0xad, 0xc4, 0xfe, 0xbc, 0x78, 0xe6, 0xa6, 0x86, 0x38, 0x61, 0x3d, 0x8c, 0x1f,
	0x03, 0x80, 0x4f, 0x5c, 0xe2, 0x8d, 0xdb, 0x60, 0x31, 0x9c, 0x89, 0x42, 0x90, 0xe0, 0xf9, 0x17,
	0xe4, 0xd2, 0x84, 0xc1, 0x28, 0xb8, 0x1a, 0xa6, 0x7b, 0x3d, 0xf6, 0x0a, 0xba, 0x92, 0x97, 0xe1,
	0xd1, 0x28, 0xb8, 0x3a, 0xbd, 0x3e, 0x4f, 0x76, 0xf4, 0xe4, 0x96, 0x97, 0x35, 0x26, 0xb5, 0x02,
	0xf6, 0x1e, 0xa0, 0xd4, 0x6a, 0x56, 0x65, 0x24, 0x54, 0x11, 0x76, 0x9d, 0x7c, 0xd4, 0x90, 0x7f,
	0xde, 0x0d, 0x53, 0x4e, 0x68, 0xfc, 0x6a, 0x63, 0x87, 0x7d, 0x84, 0xb3, 0x4c, 0x15, 0x86, 0xb4,
	0x67, 0xf4, 0xfe, 0x63, 0xdc, 0x21, 0x51, 0x8e, 0x12, 0x0b, 0xba, 0x51, 0x86, 0xb6, 0x8c, 0xbd,
	0xad, 0xf8, 0x16, 0x4e, 0x76, 0xce, 0xd8, 0x39, 0xf4, 0x97, 0x62, 0x46, 0x0b, 0x7f, 0xb3, 0xfa,
	0xc0, 0x2e, 0x61, 0xb0, 0x40, 0x31, 0x5f, 0x90, 0xbb, 0xd5, 0x30, 0xf5, 0x27, 0xc6, 0xa0, 0x67,
	0x10, 0x67, 0xce, 0x7c, 0x37, 0x75, 0x75, 0xfc, 0x33, 0x80, 0x8b, 0x83, 0xd6, 0xad, 0x9a, 0xd4,
	0xb2, 0xf0, 0x68, 0x57, 0xdb, 0x5e, 0x26, 0x68, 0xe5, 0xb9, 0xae, 0x66, 0x11, 0x40, 0x5e, 0xc9,
	0x29, 0x6a, 0x29, 0xf2, 0xdc, 0xb1, 0x87, 0x69, 0xa3, 0x63, 0xdd, 0x7c, 0xaf, 0xb8, 0xd6, 0x2b,
	0x77, 0xe1, 0x61, 0xea, 0x4f, 0x96, 0x75, 0xcf, 0xb5, 0x0c, 0xfb, 0x35, 0xcb, 0xd6, 0xf1, 0xdf,
	0x23, 0xb8, 0x38, 0x18, 0x02, 0x1b, 0x37, 0xdc, 0x9c, 0x5e, 0x5f, 0x36, 0x42, 0x4b, 0xd1, 0xa8,
	0x4a, 0x67, 0x78, 0x87, 0xe4, 0x5d, 0x8e, 0x1b, 0x2e, 0x5b, 0xb4, 0xce, 0xfd, 0x6b, 0x38, 0xce,
	0x78, 0x29, 0x88, 0xe7, 0xfe, 0x9b, 0x3e, 0x25, 0xdf, 0xca, 0x2c, 0xdd, 0xf9, 0xee, 0xb5, 0xd3,
	0xad, 0x86, 0xbd, 0xdd, 0xcb, 0xa6, 0xdf, 0xba, 0xd1, 0xcc, 0x2c, 0xd9, 0x65, 0x36, 0x68, 0xdd,
	0xd9, 0x66, 0x39, 0x86, 0x9e, 0x9d, 0x85, 0xc7, 0xed, 0x9e, 0x6c, 0xe7, 0xc3, 0xbb, 0xdf, 0xeb,
	0x28, 0x78, 0x58, 0x47, 0xc1, 0xe3, 0x3a, 0x0a, 0x7e, 0x6d, 0xa2, 0xce, 0xc3, 0x26, 0xea, 0xfc,
	0xd9, 0x44, 0x9d, 0xaf, 0x2f, 0xe7, 0x82, 0x16, 0xd5, 0x34, 0xc9, 0x94, 0x9c, 0x64, 0x72, 0xc9,
	0x09, 0xb5, 0x99, 0xb8, 0x37, 0xf6, 0xa3, 0xfe, 0xa3, 0x55, 0x89, 0x66, 0x3a, 0x70, 0x6f, 0xed,
	0xcd, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd4, 0x82, 0x98, 0x4c, 0x9d, 0x03, 0x00, 0x00,
}

func (m *GameConfig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GameConfig) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GameConfig) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Construction != nil {
		{
			size, err := m.Construction.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintConfig(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if m.Production != nil {
		{
			size, err := m.Production.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintConfig(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.Map != nil {
		{
			size, err := m.Map.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintConfig(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.InitialTeams != 0 {
		i = encodeVarintConfig(dAtA, i, uint64(m.InitialTeams))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *MapConfig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MapConfig) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MapConfig) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Seed != 0 {
		i = encodeVarintConfig(dAtA, i, uint64(m.Seed))
		i--
		dAtA[i] = 0x18
	}
	if m.Height != 0 {
		i = encodeVarintConfig(dAtA, i, uint64(m.Height))
		i--
		dAtA[i] = 0x10
	}
	if m.Width != 0 {
		i = encodeVarintConfig(dAtA, i, uint64(m.Width))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *ProductionRatesConfig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ProductionRatesConfig) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ProductionRatesConfig) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Farm != 0 {
		i = encodeVarintConfig(dAtA, i, uint64(m.Farm))
		i--
		dAtA[i] = 0x28
	}
	if m.Quarry != 0 {
		i = encodeVarintConfig(dAtA, i, uint64(m.Quarry))
		i--
		dAtA[i] = 0x20
	}
	if m.Lumbermill != 0 {
		i = encodeVarintConfig(dAtA, i, uint64(m.Lumbermill))
		i--
		dAtA[i] = 0x18
	}
	if m.City != 0 {
		i = encodeVarintConfig(dAtA, i, uint64(m.City))
		i--
		dAtA[i] = 0x10
	}
	if m.Town != 0 {
		i = encodeVarintConfig(dAtA, i, uint64(m.Town))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *SettlementCostsConfig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SettlementCostsConfig) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SettlementCostsConfig) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Rook != nil {
		{
			size, err := m.Rook.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintConfig(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x3a
	}
	if m.Quarry != nil {
		{
			size, err := m.Quarry.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintConfig(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x32
	}
	if m.Lumbermill != nil {
		{
			size, err := m.Lumbermill.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintConfig(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	if m.Farm != nil {
		{
			size, err := m.Farm.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintConfig(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if m.Capital != nil {
		{
			size, err := m.Capital.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintConfig(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.City != nil {
		{
			size, err := m.City.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintConfig(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.Town != nil {
		{
			size, err := m.Town.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintConfig(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintConfig(dAtA []byte, offset int, v uint64) int {
	offset -= sovConfig(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GameConfig) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.InitialTeams != 0 {
		n += 1 + sovConfig(uint64(m.InitialTeams))
	}
	if m.Map != nil {
		l = m.Map.Size()
		n += 1 + l + sovConfig(uint64(l))
	}
	if m.Production != nil {
		l = m.Production.Size()
		n += 1 + l + sovConfig(uint64(l))
	}
	if m.Construction != nil {
		l = m.Construction.Size()
		n += 1 + l + sovConfig(uint64(l))
	}
	return n
}

func (m *MapConfig) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Width != 0 {
		n += 1 + sovConfig(uint64(m.Width))
	}
	if m.Height != 0 {
		n += 1 + sovConfig(uint64(m.Height))
	}
	if m.Seed != 0 {
		n += 1 + sovConfig(uint64(m.Seed))
	}
	return n
}

func (m *ProductionRatesConfig) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Town != 0 {
		n += 1 + sovConfig(uint64(m.Town))
	}
	if m.City != 0 {
		n += 1 + sovConfig(uint64(m.City))
	}
	if m.Lumbermill != 0 {
		n += 1 + sovConfig(uint64(m.Lumbermill))
	}
	if m.Quarry != 0 {
		n += 1 + sovConfig(uint64(m.Quarry))
	}
	if m.Farm != 0 {
		n += 1 + sovConfig(uint64(m.Farm))
	}
	return n
}

func (m *SettlementCostsConfig) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Town != nil {
		l = m.Town.Size()
		n += 1 + l + sovConfig(uint64(l))
	}
	if m.City != nil {
		l = m.City.Size()
		n += 1 + l + sovConfig(uint64(l))
	}
	if m.Capital != nil {
		l = m.Capital.Size()
		n += 1 + l + sovConfig(uint64(l))
	}
	if m.Farm != nil {
		l = m.Farm.Size()
		n += 1 + l + sovConfig(uint64(l))
	}
	if m.Lumbermill != nil {
		l = m.Lumbermill.Size()
		n += 1 + l + sovConfig(uint64(l))
	}
	if m.Quarry != nil {
		l = m.Quarry.Size()
		n += 1 + l + sovConfig(uint64(l))
	}
	if m.Rook != nil {
		l = m.Rook.Size()
		n += 1 + l + sovConfig(uint64(l))
	}
	return n
}

func sovConfig(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozConfig(x uint64) (n int) {
	return sovConfig(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GameConfig) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConfig
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GameConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GameConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field InitialTeams", wireType)
			}
			m.InitialTeams = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.InitialTeams |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Map", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Map == nil {
				m.Map = &MapConfig{}
			}
			if err := m.Map.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Production", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Production == nil {
				m.Production = &ProductionRatesConfig{}
			}
			if err := m.Production.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Construction", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Construction == nil {
				m.Construction = &SettlementCostsConfig{}
			}
			if err := m.Construction.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipConfig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthConfig
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthConfig
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MapConfig) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConfig
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MapConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MapConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Width", wireType)
			}
			m.Width = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Width |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Height |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Seed", wireType)
			}
			m.Seed = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Seed |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipConfig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthConfig
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthConfig
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ProductionRatesConfig) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConfig
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ProductionRatesConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ProductionRatesConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Town", wireType)
			}
			m.Town = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Town |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field City", wireType)
			}
			m.City = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.City |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Lumbermill", wireType)
			}
			m.Lumbermill = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Lumbermill |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Quarry", wireType)
			}
			m.Quarry = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Quarry |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Farm", wireType)
			}
			m.Farm = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Farm |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipConfig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthConfig
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthConfig
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *SettlementCostsConfig) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConfig
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SettlementCostsConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SettlementCostsConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Town", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Town == nil {
				m.Town = &ResourceSet{}
			}
			if err := m.Town.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field City", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.City == nil {
				m.City = &ResourceSet{}
			}
			if err := m.City.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Capital", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Capital == nil {
				m.Capital = &ResourceSet{}
			}
			if err := m.Capital.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Farm", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Farm == nil {
				m.Farm = &ResourceSet{}
			}
			if err := m.Farm.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Lumbermill", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Lumbermill == nil {
				m.Lumbermill = &ResourceSet{}
			}
			if err := m.Lumbermill.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Quarry", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Quarry == nil {
				m.Quarry = &ResourceSet{}
			}
			if err := m.Quarry.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Rook", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Rook == nil {
				m.Rook = &ResourceSet{}
			}
			if err := m.Rook.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipConfig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthConfig
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthConfig
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipConfig(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowConfig
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthConfig
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupConfig
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthConfig
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthConfig        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowConfig          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupConfig = fmt.Errorf("proto: unexpected end of group")
)
