// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: rook/state.proto

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

type Landscape int32

const (
	Landscape_UNKNOWN   Landscape = 0
	Landscape_PLAINS    Landscape = 1
	Landscape_FOREST    Landscape = 2
	Landscape_MOUNTAINS Landscape = 3
	Landscape_LAKE      Landscape = 4
)

var Landscape_name = map[int32]string{
	0: "UNKNOWN",
	1: "PLAINS",
	2: "FOREST",
	3: "MOUNTAINS",
	4: "LAKE",
}

var Landscape_value = map[string]int32{
	"UNKNOWN":   0,
	"PLAINS":    1,
	"FOREST":    2,
	"MOUNTAINS": 3,
	"LAKE":      4,
}

func (x Landscape) String() string {
	return proto.EnumName(Landscape_name, int32(x))
}

func (Landscape) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_5fdb64879681e3ce, []int{0}
}

type Settlement int32

const (
	Settlement_NONE       Settlement = 0
	Settlement_TOWN       Settlement = 1
	Settlement_CITY       Settlement = 2
	Settlement_CAPITAL    Settlement = 3
	Settlement_LUMBERMILL Settlement = 4
	Settlement_QUARRY     Settlement = 5
	Settlement_FARM       Settlement = 6
	Settlement_ROOK       Settlement = 7
)

var Settlement_name = map[int32]string{
	0: "NONE",
	1: "TOWN",
	2: "CITY",
	3: "CAPITAL",
	4: "LUMBERMILL",
	5: "QUARRY",
	6: "FARM",
	7: "ROOK",
}

var Settlement_value = map[string]int32{
	"NONE":       0,
	"TOWN":       1,
	"CITY":       2,
	"CAPITAL":    3,
	"LUMBERMILL": 4,
	"QUARRY":     5,
	"FARM":       6,
	"ROOK":       7,
}

func (x Settlement) String() string {
	return proto.EnumName(Settlement_name, int32(x))
}

func (Settlement) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_5fdb64879681e3ce, []int{1}
}

// This is what is visible to each player
type PartialState struct {
	Map       map[uint32]*Tile `protobuf:"bytes,1,rep,name=map,proto3" json:"map,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Resources *ResourceSet     `protobuf:"bytes,2,opt,name=resources,proto3" json:"resources,omitempty"`
	Step      uint32           `protobuf:"varint,3,opt,name=step,proto3" json:"step,omitempty"`
}

func (m *PartialState) Reset()         { *m = PartialState{} }
func (m *PartialState) String() string { return proto.CompactTextString(m) }
func (*PartialState) ProtoMessage()    {}
func (*PartialState) Descriptor() ([]byte, []int) {
	return fileDescriptor_5fdb64879681e3ce, []int{0}
}
func (m *PartialState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PartialState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PartialState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PartialState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PartialState.Merge(m, src)
}
func (m *PartialState) XXX_Size() int {
	return m.Size()
}
func (m *PartialState) XXX_DiscardUnknown() {
	xxx_messageInfo_PartialState.DiscardUnknown(m)
}

var xxx_messageInfo_PartialState proto.InternalMessageInfo

func (m *PartialState) GetMap() map[uint32]*Tile {
	if m != nil {
		return m.Map
	}
	return nil
}

func (m *PartialState) GetResources() *ResourceSet {
	if m != nil {
		return m.Resources
	}
	return nil
}

func (m *PartialState) GetStep() uint32 {
	if m != nil {
		return m.Step
	}
	return 0
}

type Tile struct {
	Landscape  Landscape  `protobuf:"varint,1,opt,name=landscape,proto3,enum=rook.rook.Landscape" json:"landscape,omitempty"`
	Population uint32     `protobuf:"varint,2,opt,name=population,proto3" json:"population,omitempty"`
	Faction    *Faction   `protobuf:"bytes,3,opt,name=faction,proto3" json:"faction,omitempty"`
	Settlement Settlement `protobuf:"varint,4,opt,name=settlement,proto3,enum=rook.rook.Settlement" json:"settlement,omitempty"`
}

func (m *Tile) Reset()         { *m = Tile{} }
func (m *Tile) String() string { return proto.CompactTextString(m) }
func (*Tile) ProtoMessage()    {}
func (*Tile) Descriptor() ([]byte, []int) {
	return fileDescriptor_5fdb64879681e3ce, []int{1}
}
func (m *Tile) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Tile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Tile.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Tile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Tile.Merge(m, src)
}
func (m *Tile) XXX_Size() int {
	return m.Size()
}
func (m *Tile) XXX_DiscardUnknown() {
	xxx_messageInfo_Tile.DiscardUnknown(m)
}

var xxx_messageInfo_Tile proto.InternalMessageInfo

func (m *Tile) GetLandscape() Landscape {
	if m != nil {
		return m.Landscape
	}
	return Landscape_UNKNOWN
}

func (m *Tile) GetPopulation() uint32 {
	if m != nil {
		return m.Population
	}
	return 0
}

func (m *Tile) GetFaction() *Faction {
	if m != nil {
		return m.Faction
	}
	return nil
}

func (m *Tile) GetSettlement() Settlement {
	if m != nil {
		return m.Settlement
	}
	return Settlement_NONE
}

type Faction struct {
	Moniker     string                `protobuf:"bytes,1,opt,name=moniker,proto3" json:"moniker,omitempty"`
	Resources   *ResourceSet          `protobuf:"bytes,2,opt,name=resources,proto3" json:"resources,omitempty"`
	Population  map[uint32]uint32     `protobuf:"bytes,3,rep,name=population,proto3" json:"population,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	Settlements map[uint32]Settlement `protobuf:"bytes,4,rep,name=settlements,proto3" json:"settlements,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3,enum=rook.rook.Settlement"`
}

func (m *Faction) Reset()         { *m = Faction{} }
func (m *Faction) String() string { return proto.CompactTextString(m) }
func (*Faction) ProtoMessage()    {}
func (*Faction) Descriptor() ([]byte, []int) {
	return fileDescriptor_5fdb64879681e3ce, []int{2}
}
func (m *Faction) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Faction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Faction.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Faction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Faction.Merge(m, src)
}
func (m *Faction) XXX_Size() int {
	return m.Size()
}
func (m *Faction) XXX_DiscardUnknown() {
	xxx_messageInfo_Faction.DiscardUnknown(m)
}

var xxx_messageInfo_Faction proto.InternalMessageInfo

func (m *Faction) GetMoniker() string {
	if m != nil {
		return m.Moniker
	}
	return ""
}

func (m *Faction) GetResources() *ResourceSet {
	if m != nil {
		return m.Resources
	}
	return nil
}

func (m *Faction) GetPopulation() map[uint32]uint32 {
	if m != nil {
		return m.Population
	}
	return nil
}

func (m *Faction) GetSettlements() map[uint32]Settlement {
	if m != nil {
		return m.Settlements
	}
	return nil
}

type ResourceSet struct {
	Food  uint32 `protobuf:"varint,1,opt,name=food,proto3" json:"food,omitempty"`
	Stone uint32 `protobuf:"varint,2,opt,name=stone,proto3" json:"stone,omitempty"`
	Wood  uint32 `protobuf:"varint,3,opt,name=wood,proto3" json:"wood,omitempty"`
}

func (m *ResourceSet) Reset()         { *m = ResourceSet{} }
func (m *ResourceSet) String() string { return proto.CompactTextString(m) }
func (*ResourceSet) ProtoMessage()    {}
func (*ResourceSet) Descriptor() ([]byte, []int) {
	return fileDescriptor_5fdb64879681e3ce, []int{3}
}
func (m *ResourceSet) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ResourceSet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ResourceSet.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ResourceSet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResourceSet.Merge(m, src)
}
func (m *ResourceSet) XXX_Size() int {
	return m.Size()
}
func (m *ResourceSet) XXX_DiscardUnknown() {
	xxx_messageInfo_ResourceSet.DiscardUnknown(m)
}

var xxx_messageInfo_ResourceSet proto.InternalMessageInfo

func (m *ResourceSet) GetFood() uint32 {
	if m != nil {
		return m.Food
	}
	return 0
}

func (m *ResourceSet) GetStone() uint32 {
	if m != nil {
		return m.Stone
	}
	return 0
}

func (m *ResourceSet) GetWood() uint32 {
	if m != nil {
		return m.Wood
	}
	return 0
}

func init() {
	proto.RegisterEnum("rook.rook.Landscape", Landscape_name, Landscape_value)
	proto.RegisterEnum("rook.rook.Settlement", Settlement_name, Settlement_value)
	proto.RegisterType((*PartialState)(nil), "rook.rook.PartialState")
	proto.RegisterMapType((map[uint32]*Tile)(nil), "rook.rook.PartialState.MapEntry")
	proto.RegisterType((*Tile)(nil), "rook.rook.Tile")
	proto.RegisterType((*Faction)(nil), "rook.rook.Faction")
	proto.RegisterMapType((map[uint32]uint32)(nil), "rook.rook.Faction.PopulationEntry")
	proto.RegisterMapType((map[uint32]Settlement)(nil), "rook.rook.Faction.SettlementsEntry")
	proto.RegisterType((*ResourceSet)(nil), "rook.rook.ResourceSet")
}

func init() { proto.RegisterFile("rook/state.proto", fileDescriptor_5fdb64879681e3ce) }

var fileDescriptor_5fdb64879681e3ce = []byte{
	// 596 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0xcd, 0xc6, 0x6e, 0xd3, 0x4c, 0x48, 0xbb, 0x5a, 0x15, 0x64, 0xf5, 0x60, 0x45, 0x41, 0x95,
	0xaa, 0x82, 0x5c, 0x29, 0x80, 0x84, 0x90, 0x10, 0x72, 0x2b, 0x17, 0x45, 0x71, 0xec, 0xb0, 0x71,
	0x84, 0xca, 0xcd, 0x4d, 0xb7, 0x25, 0xaa, 0xe3, 0xb5, 0xec, 0x0d, 0xa5, 0xff, 0x82, 0x7f, 0xc4,
	0x95, 0x63, 0x2f, 0x48, 0x1c, 0x51, 0xfb, 0x2f, 0x38, 0xa1, 0xdd, 0x7c, 0x2d, 0x25, 0x5c, 0xb8,
	0x24, 0xcf, 0xfb, 0xde, 0xbc, 0x99, 0x9d, 0x19, 0x2d, 0xe0, 0x9c, 0xf3, 0xcb, 0x83, 0x42, 0xc4,
	0x82, 0x39, 0x59, 0xce, 0x05, 0x27, 0x55, 0x79, 0xe2, 0xc8, 0x9f, 0xe6, 0x77, 0x04, 0x0f, 0x7a,
	0x71, 0x2e, 0x46, 0x71, 0xd2, 0x97, 0x0a, 0xd2, 0x02, 0x63, 0x1c, 0x67, 0x16, 0x6a, 0x18, 0x7b,
	0xb5, 0x56, 0xc3, 0x59, 0x28, 0x1d, 0x5d, 0xe5, 0x74, 0xe3, 0xcc, 0x4b, 0x45, 0x7e, 0x4d, 0xa5,
	0x98, 0x3c, 0x87, 0x6a, 0xce, 0x0a, 0x3e, 0xc9, 0x87, 0xac, 0xb0, 0xca, 0x0d, 0xb4, 0x57, 0x6b,
	0x3d, 0xd2, 0x22, 0xe9, 0x8c, 0xeb, 0x33, 0x41, 0x97, 0x42, 0x42, 0xc0, 0x2c, 0x04, 0xcb, 0x2c,
	0xa3, 0x81, 0xf6, 0xea, 0x54, 0xe1, 0x9d, 0xb7, 0xb0, 0x31, 0xb7, 0x26, 0x18, 0x8c, 0x4b, 0x76,
	0x6d, 0x21, 0x45, 0x4b, 0x48, 0x76, 0x61, 0xed, 0x53, 0x9c, 0x4c, 0xd8, 0x2c, 0xc7, 0x96, 0x96,
	0x23, 0x1a, 0x25, 0x8c, 0x4e, 0xd9, 0x57, 0xe5, 0x97, 0xa8, 0xf9, 0x15, 0x81, 0x29, 0xcf, 0x48,
	0x0b, 0xaa, 0x49, 0x9c, 0x9e, 0x15, 0xc3, 0x38, 0x63, 0xca, 0x6b, 0xb3, 0xb5, 0xad, 0xc5, 0xf9,
	0x73, 0x8e, 0x2e, 0x65, 0xc4, 0x06, 0xc8, 0x78, 0x36, 0x49, 0x62, 0x31, 0xe2, 0xa9, 0x4a, 0x56,
	0xa7, 0xda, 0x09, 0x79, 0x0a, 0x95, 0xf3, 0x78, 0xa8, 0x48, 0x43, 0x55, 0x42, 0x34, 0xc7, 0xe3,
	0x29, 0x43, 0xe7, 0x12, 0xf2, 0x02, 0xa0, 0x60, 0x42, 0x24, 0x6c, 0xcc, 0x52, 0x61, 0x99, 0xaa,
	0x84, 0x87, 0x5a, 0x40, 0x7f, 0x41, 0x52, 0x4d, 0xd8, 0xfc, 0x55, 0x86, 0xca, 0xcc, 0x8b, 0x58,
	0x50, 0x19, 0xf3, 0x74, 0x74, 0xc9, 0x72, 0x75, 0x85, 0x2a, 0x9d, 0x7f, 0xfe, 0x67, 0xeb, 0x0f,
	0xff, 0xb8, 0xa0, 0xa1, 0x66, 0xdd, 0xfc, 0xfb, 0x0e, 0x4e, 0x6f, 0x21, 0x9a, 0x4e, 0x5b, 0x6f,
	0x82, 0x07, 0xb5, 0x65, 0xb5, 0x85, 0x65, 0x2a, 0x93, 0xc7, 0x2b, 0x4c, 0x96, 0xf7, 0x2b, 0xa6,
	0x2e, 0x7a, 0xdc, 0xce, 0x6b, 0xd8, 0xba, 0x97, 0x65, 0xc5, 0xe0, 0xb7, 0xf5, 0xc1, 0xd7, 0xb5,
	0x39, 0xef, 0x0c, 0x00, 0xdf, 0xf7, 0x5f, 0x11, 0xff, 0x44, 0x8f, 0xff, 0x67, 0xf7, 0xb5, 0xf5,
	0xe9, 0x40, 0x4d, 0x6b, 0x9d, 0x5c, 0xd5, 0x73, 0xce, 0xcf, 0x66, 0x96, 0x0a, 0xcb, 0x9a, 0x0a,
	0xc1, 0xd3, 0x45, 0x4d, 0xea, 0x43, 0x2a, 0xaf, 0xa4, 0x72, 0xb6, 0xd4, 0x12, 0xef, 0xb7, 0xa1,
	0xba, 0x58, 0x33, 0x52, 0x83, 0xca, 0x20, 0xe8, 0x04, 0xe1, 0xfb, 0x00, 0x97, 0x08, 0xc0, 0x7a,
	0xcf, 0x77, 0xdb, 0x41, 0x1f, 0x23, 0x89, 0x8f, 0x43, 0xea, 0xf5, 0x23, 0x5c, 0x26, 0x75, 0xa8,
	0x76, 0xc3, 0x41, 0x10, 0x29, 0xca, 0x20, 0x1b, 0x60, 0xfa, 0x6e, 0xc7, 0xc3, 0xe6, 0xfe, 0x05,
	0xc0, 0xb2, 0x60, 0x79, 0x1e, 0x84, 0x81, 0x87, 0x4b, 0x12, 0x45, 0xd2, 0x12, 0x49, 0x74, 0xd4,
	0x8e, 0x4e, 0x70, 0x59, 0x66, 0x3a, 0x72, 0x7b, 0xed, 0xc8, 0xf5, 0xb1, 0x41, 0x36, 0x01, 0xfc,
	0x41, 0xf7, 0xd0, 0xa3, 0xdd, 0xb6, 0xef, 0x63, 0x53, 0x66, 0x7b, 0x37, 0x70, 0x29, 0x3d, 0xc1,
	0x6b, 0x32, 0xe4, 0xd8, 0xa5, 0x5d, 0xbc, 0x2e, 0x11, 0x0d, 0xc3, 0x0e, 0xae, 0x1c, 0xbe, 0xf9,
	0x76, 0x6b, 0xa3, 0x9b, 0x5b, 0x1b, 0xfd, 0xbc, 0xb5, 0xd1, 0x97, 0x3b, 0xbb, 0x74, 0x73, 0x67,
	0x97, 0x7e, 0xdc, 0xd9, 0xa5, 0x0f, 0xbb, 0x17, 0x23, 0xf1, 0x71, 0x72, 0xea, 0x0c, 0xf9, 0xf8,
	0x60, 0x38, 0xbe, 0x8a, 0x05, 0xcb, 0x8b, 0x03, 0xf5, 0xc4, 0x7c, 0x9e, 0xfe, 0x89, 0xeb, 0x8c,
	0x15, 0xa7, 0xeb, 0xea, 0xa9, 0x79, 0xf6, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x8a, 0x4f, 0xd0, 0xbc,
	0x7e, 0x04, 0x00, 0x00,
}

func (m *PartialState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PartialState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PartialState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Step != 0 {
		i = encodeVarintState(dAtA, i, uint64(m.Step))
		i--
		dAtA[i] = 0x18
	}
	if m.Resources != nil {
		{
			size, err := m.Resources.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintState(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Map) > 0 {
		for k := range m.Map {
			v := m.Map[k]
			baseI := i
			if v != nil {
				{
					size, err := v.MarshalToSizedBuffer(dAtA[:i])
					if err != nil {
						return 0, err
					}
					i -= size
					i = encodeVarintState(dAtA, i, uint64(size))
				}
				i--
				dAtA[i] = 0x12
			}
			i = encodeVarintState(dAtA, i, uint64(k))
			i--
			dAtA[i] = 0x8
			i = encodeVarintState(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *Tile) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Tile) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Tile) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Settlement != 0 {
		i = encodeVarintState(dAtA, i, uint64(m.Settlement))
		i--
		dAtA[i] = 0x20
	}
	if m.Faction != nil {
		{
			size, err := m.Faction.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintState(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.Population != 0 {
		i = encodeVarintState(dAtA, i, uint64(m.Population))
		i--
		dAtA[i] = 0x10
	}
	if m.Landscape != 0 {
		i = encodeVarintState(dAtA, i, uint64(m.Landscape))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Faction) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Faction) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Faction) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Settlements) > 0 {
		for k := range m.Settlements {
			v := m.Settlements[k]
			baseI := i
			i = encodeVarintState(dAtA, i, uint64(v))
			i--
			dAtA[i] = 0x10
			i = encodeVarintState(dAtA, i, uint64(k))
			i--
			dAtA[i] = 0x8
			i = encodeVarintState(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.Population) > 0 {
		for k := range m.Population {
			v := m.Population[k]
			baseI := i
			i = encodeVarintState(dAtA, i, uint64(v))
			i--
			dAtA[i] = 0x10
			i = encodeVarintState(dAtA, i, uint64(k))
			i--
			dAtA[i] = 0x8
			i = encodeVarintState(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.Resources != nil {
		{
			size, err := m.Resources.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintState(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Moniker) > 0 {
		i -= len(m.Moniker)
		copy(dAtA[i:], m.Moniker)
		i = encodeVarintState(dAtA, i, uint64(len(m.Moniker)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ResourceSet) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ResourceSet) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ResourceSet) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Wood != 0 {
		i = encodeVarintState(dAtA, i, uint64(m.Wood))
		i--
		dAtA[i] = 0x18
	}
	if m.Stone != 0 {
		i = encodeVarintState(dAtA, i, uint64(m.Stone))
		i--
		dAtA[i] = 0x10
	}
	if m.Food != 0 {
		i = encodeVarintState(dAtA, i, uint64(m.Food))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintState(dAtA []byte, offset int, v uint64) int {
	offset -= sovState(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *PartialState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Map) > 0 {
		for k, v := range m.Map {
			_ = k
			_ = v
			l = 0
			if v != nil {
				l = v.Size()
				l += 1 + sovState(uint64(l))
			}
			mapEntrySize := 1 + sovState(uint64(k)) + l
			n += mapEntrySize + 1 + sovState(uint64(mapEntrySize))
		}
	}
	if m.Resources != nil {
		l = m.Resources.Size()
		n += 1 + l + sovState(uint64(l))
	}
	if m.Step != 0 {
		n += 1 + sovState(uint64(m.Step))
	}
	return n
}

func (m *Tile) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Landscape != 0 {
		n += 1 + sovState(uint64(m.Landscape))
	}
	if m.Population != 0 {
		n += 1 + sovState(uint64(m.Population))
	}
	if m.Faction != nil {
		l = m.Faction.Size()
		n += 1 + l + sovState(uint64(l))
	}
	if m.Settlement != 0 {
		n += 1 + sovState(uint64(m.Settlement))
	}
	return n
}

func (m *Faction) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Moniker)
	if l > 0 {
		n += 1 + l + sovState(uint64(l))
	}
	if m.Resources != nil {
		l = m.Resources.Size()
		n += 1 + l + sovState(uint64(l))
	}
	if len(m.Population) > 0 {
		for k, v := range m.Population {
			_ = k
			_ = v
			mapEntrySize := 1 + sovState(uint64(k)) + 1 + sovState(uint64(v))
			n += mapEntrySize + 1 + sovState(uint64(mapEntrySize))
		}
	}
	if len(m.Settlements) > 0 {
		for k, v := range m.Settlements {
			_ = k
			_ = v
			mapEntrySize := 1 + sovState(uint64(k)) + 1 + sovState(uint64(v))
			n += mapEntrySize + 1 + sovState(uint64(mapEntrySize))
		}
	}
	return n
}

func (m *ResourceSet) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Food != 0 {
		n += 1 + sovState(uint64(m.Food))
	}
	if m.Stone != 0 {
		n += 1 + sovState(uint64(m.Stone))
	}
	if m.Wood != 0 {
		n += 1 + sovState(uint64(m.Wood))
	}
	return n
}

func sovState(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozState(x uint64) (n int) {
	return sovState(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PartialState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowState
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
			return fmt.Errorf("proto: PartialState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PartialState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Map", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
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
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Map == nil {
				m.Map = make(map[uint32]*Tile)
			}
			var mapkey uint32
			var mapvalue *Tile
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowState
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
				if fieldNum == 1 {
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowState
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapkey |= uint32(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
				} else if fieldNum == 2 {
					var mapmsglen int
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowState
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapmsglen |= int(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					if mapmsglen < 0 {
						return ErrInvalidLengthState
					}
					postmsgIndex := iNdEx + mapmsglen
					if postmsgIndex < 0 {
						return ErrInvalidLengthState
					}
					if postmsgIndex > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = &Tile{}
					if err := mapvalue.Unmarshal(dAtA[iNdEx:postmsgIndex]); err != nil {
						return err
					}
					iNdEx = postmsgIndex
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipState(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if skippy < 0 {
						return ErrInvalidLengthState
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Map[mapkey] = mapvalue
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Resources", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
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
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Resources == nil {
				m.Resources = &ResourceSet{}
			}
			if err := m.Resources.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Step", wireType)
			}
			m.Step = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Step |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipState(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthState
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthState
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
func (m *Tile) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowState
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
			return fmt.Errorf("proto: Tile: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Tile: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Landscape", wireType)
			}
			m.Landscape = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Landscape |= Landscape(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Population", wireType)
			}
			m.Population = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Population |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Faction", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
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
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Faction == nil {
				m.Faction = &Faction{}
			}
			if err := m.Faction.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Settlement", wireType)
			}
			m.Settlement = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Settlement |= Settlement(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipState(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthState
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthState
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
func (m *Faction) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowState
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
			return fmt.Errorf("proto: Faction: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Faction: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Moniker", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Moniker = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Resources", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
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
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Resources == nil {
				m.Resources = &ResourceSet{}
			}
			if err := m.Resources.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Population", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
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
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Population == nil {
				m.Population = make(map[uint32]uint32)
			}
			var mapkey uint32
			var mapvalue uint32
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowState
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
				if fieldNum == 1 {
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowState
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapkey |= uint32(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
				} else if fieldNum == 2 {
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowState
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapvalue |= uint32(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipState(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if skippy < 0 {
						return ErrInvalidLengthState
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Population[mapkey] = mapvalue
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Settlements", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
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
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Settlements == nil {
				m.Settlements = make(map[uint32]Settlement)
			}
			var mapkey uint32
			var mapvalue Settlement
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowState
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
				if fieldNum == 1 {
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowState
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapkey |= uint32(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
				} else if fieldNum == 2 {
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowState
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapvalue |= Settlement(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipState(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if skippy < 0 {
						return ErrInvalidLengthState
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Settlements[mapkey] = mapvalue
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipState(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthState
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthState
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
func (m *ResourceSet) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowState
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
			return fmt.Errorf("proto: ResourceSet: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ResourceSet: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Food", wireType)
			}
			m.Food = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Food |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Stone", wireType)
			}
			m.Stone = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Stone |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Wood", wireType)
			}
			m.Wood = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Wood |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipState(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthState
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthState
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
func skipState(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowState
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
					return 0, ErrIntOverflowState
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
					return 0, ErrIntOverflowState
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
				return 0, ErrInvalidLengthState
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupState
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthState
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthState        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowState          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupState = fmt.Errorf("proto: unexpected end of group")
)