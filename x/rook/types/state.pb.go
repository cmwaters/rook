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
	Landscape_L_UNKNOWN Landscape = 0
	Landscape_PLAINS    Landscape = 1
	Landscape_FOREST    Landscape = 2
	Landscape_MOUNTAINS Landscape = 3
	Landscape_LAKE      Landscape = 4
)

var Landscape_name = map[int32]string{
	0: "L_UNKNOWN",
	1: "PLAINS",
	2: "FOREST",
	3: "MOUNTAINS",
	4: "LAKE",
}

var Landscape_value = map[string]int32{
	"L_UNKNOWN": 0,
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
	Settlement_S_UNKNOWN  Settlement = 0
	Settlement_TOWN       Settlement = 1
	Settlement_CITY       Settlement = 2
	Settlement_CAPITAL    Settlement = 3
	Settlement_LUMBERMILL Settlement = 4
	Settlement_QUARRY     Settlement = 5
	Settlement_FARM       Settlement = 6
	Settlement_ROOK       Settlement = 7
)

var Settlement_name = map[int32]string{
	0: "S_UNKNOWN",
	1: "TOWN",
	2: "CITY",
	3: "CAPITAL",
	4: "LUMBERMILL",
	5: "QUARRY",
	6: "FARM",
	7: "ROOK",
}

var Settlement_value = map[string]int32{
	"S_UNKNOWN":  0,
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
	Landscape  Landscape `protobuf:"varint,1,opt,name=landscape,proto3,enum=rook.rook.Landscape" json:"landscape,omitempty"`
	Population uint32    `protobuf:"varint,2,opt,name=population,proto3" json:"population,omitempty"`
	Faction    *Faction  `protobuf:"bytes,3,opt,name=faction,proto3" json:"faction,omitempty"`
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
	return Landscape_L_UNKNOWN
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

type Faction struct {
	Moniker     string                `protobuf:"bytes,1,opt,name=moniker,proto3" json:"moniker,omitempty"`
	Food        uint32                `protobuf:"varint,2,opt,name=food,proto3" json:"food,omitempty"`
	Wood        uint32                `protobuf:"varint,3,opt,name=wood,proto3" json:"wood,omitempty"`
	Stone       uint32                `protobuf:"varint,4,opt,name=stone,proto3" json:"stone,omitempty"`
	Population  map[uint32]uint32     `protobuf:"bytes,5,rep,name=population,proto3" json:"population,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	Settlements map[uint32]Settlement `protobuf:"bytes,6,rep,name=settlements,proto3" json:"settlements,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3,enum=rook.rook.Settlement"`
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

func (m *Faction) GetFood() uint32 {
	if m != nil {
		return m.Food
	}
	return 0
}

func (m *Faction) GetWood() uint32 {
	if m != nil {
		return m.Wood
	}
	return 0
}

func (m *Faction) GetStone() uint32 {
	if m != nil {
		return m.Stone
	}
	return 0
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
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x54, 0x4f, 0x6b, 0xdb, 0x4e,
	0x10, 0xf5, 0x5a, 0xb2, 0x1d, 0x8f, 0x7f, 0x4e, 0x96, 0x25, 0xbf, 0x22, 0x72, 0x10, 0xc6, 0x25,
	0x10, 0xd2, 0xa2, 0x80, 0xda, 0x43, 0x29, 0x94, 0xa2, 0x04, 0xa7, 0x18, 0xcb, 0x7f, 0xba, 0x92,
	0x29, 0xe9, 0xa5, 0x28, 0xce, 0xa6, 0x35, 0x91, 0xb5, 0x42, 0x5a, 0x37, 0xcd, 0xb5, 0xa7, 0x1e,
	0xfb, 0xb1, 0x7a, 0xcc, 0xa5, 0xd0, 0x63, 0x49, 0xbe, 0x48, 0xd9, 0x95, 0xff, 0x6c, 0x53, 0x5f,
	0xec, 0xb7, 0x33, 0xef, 0xbd, 0x19, 0xcd, 0x2c, 0x0b, 0x38, 0xe3, 0xfc, 0xea, 0x28, 0x17, 0x91,
	0x60, 0x4e, 0x9a, 0x71, 0xc1, 0x49, 0x5d, 0x46, 0x1c, 0xf9, 0xd3, 0xfe, 0x89, 0xe0, 0xbf, 0x51,
	0x94, 0x89, 0x69, 0x14, 0x07, 0x92, 0x41, 0x5c, 0x30, 0x66, 0x51, 0x6a, 0xa1, 0x96, 0x71, 0xd0,
	0x70, 0x5b, 0xce, 0x8a, 0xe9, 0xe8, 0x2c, 0xa7, 0x1f, 0xa5, 0x9d, 0x44, 0x64, 0x37, 0x54, 0x92,
	0xc9, 0x73, 0xa8, 0x67, 0x2c, 0xe7, 0xf3, 0x6c, 0xc2, 0x72, 0xab, 0xdc, 0x42, 0x07, 0x0d, 0xf7,
	0x91, 0xa6, 0xa4, 0x8b, 0x5c, 0xc0, 0x04, 0x5d, 0x13, 0x09, 0x01, 0x33, 0x17, 0x2c, 0xb5, 0x8c,
	0x16, 0x3a, 0x68, 0x52, 0x85, 0xf7, 0xde, 0xc0, 0xd6, 0xd2, 0x9a, 0x60, 0x30, 0xae, 0xd8, 0x8d,
	0x85, 0x54, 0x5a, 0x42, 0xb2, 0x0f, 0x95, 0xcf, 0x51, 0x3c, 0x67, 0x8b, 0x1a, 0x3b, 0x5a, 0x8d,
	0x70, 0x1a, 0x33, 0x5a, 0x64, 0x5f, 0x96, 0x5f, 0xa0, 0xf6, 0x37, 0x04, 0xa6, 0x8c, 0x11, 0x17,
	0xea, 0x71, 0x94, 0x5c, 0xe4, 0x93, 0x28, 0x65, 0xca, 0x6b, 0xdb, 0xdd, 0xd5, 0x74, 0xfe, 0x32,
	0x47, 0xd7, 0x34, 0x62, 0x03, 0xa4, 0x3c, 0x9d, 0xc7, 0x91, 0x98, 0xf2, 0x44, 0x15, 0x6b, 0x52,
	0x2d, 0x42, 0x9e, 0x42, 0xed, 0x32, 0x9a, 0xa8, 0xa4, 0xa1, 0x3a, 0x21, 0x9a, 0xe3, 0x69, 0x91,
	0xa1, 0x4b, 0x4a, 0xfb, 0xab, 0x01, 0xb5, 0x45, 0x90, 0x58, 0x50, 0x9b, 0xf1, 0x64, 0x7a, 0xc5,
	0x32, 0xd5, 0x4b, 0x9d, 0x2e, 0x8f, 0x72, 0x1a, 0x97, 0x9c, 0x5f, 0x2c, 0xaa, 0x29, 0x2c, 0x63,
	0xd7, 0x32, 0xb6, 0x98, 0x90, 0xc4, 0x64, 0x17, 0x2a, 0xb9, 0xe0, 0x09, 0xb3, 0x4c, 0x15, 0x2c,
	0x0e, 0xe4, 0xf8, 0xaf, 0x8e, 0x2b, 0x6a, 0x79, 0xed, 0x7f, 0x9b, 0x72, 0x46, 0x2b, 0x52, 0xb1,
	0x3e, 0xfd, 0xab, 0x3a, 0xd0, 0xc8, 0x99, 0x10, 0x31, 0x9b, 0xb1, 0x44, 0xe4, 0x56, 0x55, 0x99,
	0x3c, 0xde, 0x60, 0x12, 0xac, 0x59, 0x85, 0x8b, 0xae, 0xdb, 0x7b, 0x05, 0x3b, 0x0f, 0xaa, 0x6c,
	0xd8, 0xe4, 0xae, 0xbe, 0xc9, 0xa6, 0xb6, 0xb8, 0xbd, 0x31, 0xe0, 0x87, 0xfe, 0x1b, 0xf4, 0x4f,
	0x74, 0xfd, 0xb6, 0xfb, 0xbf, 0xd6, 0xe5, 0x5a, 0xad, 0xdf, 0x87, 0x1e, 0x34, 0xb4, 0x6b, 0xb8,
	0x9a, 0x36, 0xd2, 0xa6, 0xbd, 0x9a, 0x6c, 0x59, 0x9f, 0xec, 0x86, 0x1d, 0x1c, 0xf6, 0xa0, 0xbe,
	0xba, 0x37, 0xa4, 0x09, 0x75, 0xff, 0xc3, 0x78, 0xd0, 0x1b, 0x0c, 0xdf, 0x0d, 0x70, 0x89, 0x00,
	0x54, 0x47, 0xbe, 0xd7, 0x1d, 0x04, 0x18, 0x49, 0x7c, 0x3a, 0xa4, 0x9d, 0x20, 0xc4, 0x65, 0x49,
	0xeb, 0x0f, 0xc7, 0x83, 0x50, 0xa5, 0x0c, 0xb2, 0x05, 0xa6, 0xef, 0xf5, 0x3a, 0xd8, 0x3c, 0x8c,
	0x01, 0xd6, 0x2d, 0x4b, 0x5a, 0xa0, 0xb9, 0x6d, 0x81, 0x19, 0x4a, 0x84, 0x24, 0x3a, 0xe9, 0x86,
	0x67, 0xb8, 0x4c, 0x1a, 0x50, 0x3b, 0xf1, 0x46, 0xdd, 0xd0, 0xf3, 0xb1, 0x41, 0xb6, 0x01, 0xfc,
	0x71, 0xff, 0xb8, 0x43, 0xfb, 0x5d, 0xdf, 0xc7, 0xa6, 0x2c, 0xf9, 0x76, 0xec, 0x51, 0x7a, 0x86,
	0x2b, 0x52, 0x72, 0xea, 0xd1, 0x3e, 0xae, 0x4a, 0x44, 0x87, 0xc3, 0x1e, 0xae, 0x1d, 0xbf, 0xfe,
	0x71, 0x67, 0xa3, 0xdb, 0x3b, 0x1b, 0xfd, 0xbe, 0xb3, 0xd1, 0xf7, 0x7b, 0xbb, 0x74, 0x7b, 0x6f,
	0x97, 0x7e, 0xdd, 0xdb, 0xa5, 0xf7, 0xfb, 0x1f, 0xa7, 0xe2, 0xd3, 0xfc, 0xdc, 0x99, 0xf0, 0xd9,
	0xd1, 0x64, 0x76, 0x1d, 0x09, 0x96, 0xe5, 0x47, 0xea, 0xe9, 0xf8, 0x52, 0xfc, 0x89, 0x9b, 0x94,
	0xe5, 0xe7, 0x55, 0xf5, 0x84, 0x3c, 0xfb, 0x13, 0x00, 0x00, 0xff, 0xff, 0x27, 0x52, 0x34, 0x01,
	0x56, 0x04, 0x00, 0x00,
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
			dAtA[i] = 0x32
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
			dAtA[i] = 0x2a
		}
	}
	if m.Stone != 0 {
		i = encodeVarintState(dAtA, i, uint64(m.Stone))
		i--
		dAtA[i] = 0x20
	}
	if m.Wood != 0 {
		i = encodeVarintState(dAtA, i, uint64(m.Wood))
		i--
		dAtA[i] = 0x18
	}
	if m.Food != 0 {
		i = encodeVarintState(dAtA, i, uint64(m.Food))
		i--
		dAtA[i] = 0x10
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
	if m.Food != 0 {
		n += 1 + sovState(uint64(m.Food))
	}
	if m.Wood != 0 {
		n += 1 + sovState(uint64(m.Wood))
	}
	if m.Stone != 0 {
		n += 1 + sovState(uint64(m.Stone))
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
		case 4:
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
		case 5:
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
		case 6:
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
