// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: geobuf.proto

/*
	Package geobuf is a generated protocol buffer package.

	It is generated from these files:
		geobuf.proto

	It has these top-level messages:
		Value
		Feature
		FeatureCollection
*/
package geobuf

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import binary "encoding/binary"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// GeomType is described in section 4.3.4 of the specification
type GeomType int32

const (
	GeomType_UNKNOWN    GeomType = 0
	GeomType_POINT      GeomType = 1
	GeomType_LINESTRING GeomType = 2
	GeomType_POLYGON    GeomType = 3
)

var GeomType_name = map[int32]string{
	0: "UNKNOWN",
	1: "POINT",
	2: "LINESTRING",
	3: "POLYGON",
}
var GeomType_value = map[string]int32{
	"UNKNOWN":    0,
	"POINT":      1,
	"LINESTRING": 2,
	"POLYGON":    3,
}

func (x GeomType) String() string {
	return proto.EnumName(GeomType_name, int32(x))
}
func (GeomType) EnumDescriptor() ([]byte, []int) { return fileDescriptorGeobuf, []int{0} }

// Variant type encoding
// The use of values is described in section 4.1 of the specification
type Value struct {
	// Exactly one of these values must be present in a valid message
	StringValue string  `protobuf:"bytes,1,opt,name=string_value,json=stringValue,proto3" json:"string_value,omitempty"`
	FloatValue  float32 `protobuf:"fixed32,2,opt,name=float_value,json=floatValue,proto3" json:"float_value,omitempty"`
	DoubleValue float64 `protobuf:"fixed64,3,opt,name=double_value,json=doubleValue,proto3" json:"double_value,omitempty"`
	IntValue    int64   `protobuf:"varint,4,opt,name=int_value,json=intValue,proto3" json:"int_value,omitempty"`
	UintValue   uint64  `protobuf:"varint,5,opt,name=uint_value,json=uintValue,proto3" json:"uint_value,omitempty"`
	SintValue   int64   `protobuf:"zigzag64,6,opt,name=sint_value,json=sintValue,proto3" json:"sint_value,omitempty"`
	BoolValue   bool    `protobuf:"varint,7,opt,name=bool_value,json=boolValue,proto3" json:"bool_value,omitempty"`
}

func (m *Value) Reset()                    { *m = Value{} }
func (m *Value) String() string            { return proto.CompactTextString(m) }
func (*Value) ProtoMessage()               {}
func (*Value) Descriptor() ([]byte, []int) { return fileDescriptorGeobuf, []int{0} }

func (m *Value) GetStringValue() string {
	if m != nil {
		return m.StringValue
	}
	return ""
}

func (m *Value) GetFloatValue() float32 {
	if m != nil {
		return m.FloatValue
	}
	return 0
}

func (m *Value) GetDoubleValue() float64 {
	if m != nil {
		return m.DoubleValue
	}
	return 0
}

func (m *Value) GetIntValue() int64 {
	if m != nil {
		return m.IntValue
	}
	return 0
}

func (m *Value) GetUintValue() uint64 {
	if m != nil {
		return m.UintValue
	}
	return 0
}

func (m *Value) GetSintValue() int64 {
	if m != nil {
		return m.SintValue
	}
	return 0
}

func (m *Value) GetBoolValue() bool {
	if m != nil {
		return m.BoolValue
	}
	return false
}

type Feature struct {
	Id          uint64            `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Properties  map[string]*Value `protobuf:"bytes,2,rep,name=Properties" json:"Properties,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value"`
	Type        GeomType          `protobuf:"varint,3,opt,name=type,proto3,enum=GeomType" json:"type,omitempty"`
	Geometry    []uint64          `protobuf:"varint,4,rep,packed,name=Geometry" json:"Geometry,omitempty"`
	BoundingBox []int64           `protobuf:"varint,5,rep,packed,name=BoundingBox" json:"BoundingBox,omitempty"`
}

func (m *Feature) Reset()                    { *m = Feature{} }
func (m *Feature) String() string            { return proto.CompactTextString(m) }
func (*Feature) ProtoMessage()               {}
func (*Feature) Descriptor() ([]byte, []int) { return fileDescriptorGeobuf, []int{1} }

func (m *Feature) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Feature) GetProperties() map[string]*Value {
	if m != nil {
		return m.Properties
	}
	return nil
}

func (m *Feature) GetType() GeomType {
	if m != nil {
		return m.Type
	}
	return GeomType_UNKNOWN
}

func (m *Feature) GetGeometry() []uint64 {
	if m != nil {
		return m.Geometry
	}
	return nil
}

func (m *Feature) GetBoundingBox() []int64 {
	if m != nil {
		return m.BoundingBox
	}
	return nil
}

type FeatureCollection struct {
	Features []*Feature `protobuf:"bytes,1,rep,name=Features" json:"Features,omitempty"`
}

func (m *FeatureCollection) Reset()                    { *m = FeatureCollection{} }
func (m *FeatureCollection) String() string            { return proto.CompactTextString(m) }
func (*FeatureCollection) ProtoMessage()               {}
func (*FeatureCollection) Descriptor() ([]byte, []int) { return fileDescriptorGeobuf, []int{2} }

func (m *FeatureCollection) GetFeatures() []*Feature {
	if m != nil {
		return m.Features
	}
	return nil
}

func init() {
	proto.RegisterType((*Value)(nil), "Value")
	proto.RegisterType((*Feature)(nil), "Feature")
	proto.RegisterType((*FeatureCollection)(nil), "FeatureCollection")
	proto.RegisterEnum("GeomType", GeomType_name, GeomType_value)
}
func (m *Value) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Value) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.StringValue) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintGeobuf(dAtA, i, uint64(len(m.StringValue)))
		i += copy(dAtA[i:], m.StringValue)
	}
	if m.FloatValue != 0 {
		dAtA[i] = 0x15
		i++
		binary.LittleEndian.PutUint32(dAtA[i:], uint32(math.Float32bits(float32(m.FloatValue))))
		i += 4
	}
	if m.DoubleValue != 0 {
		dAtA[i] = 0x19
		i++
		binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.DoubleValue))))
		i += 8
	}
	if m.IntValue != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintGeobuf(dAtA, i, uint64(m.IntValue))
	}
	if m.UintValue != 0 {
		dAtA[i] = 0x28
		i++
		i = encodeVarintGeobuf(dAtA, i, uint64(m.UintValue))
	}
	if m.SintValue != 0 {
		dAtA[i] = 0x30
		i++
		i = encodeVarintGeobuf(dAtA, i, uint64((uint64(m.SintValue)<<1)^uint64((m.SintValue>>63))))
	}
	if m.BoolValue {
		dAtA[i] = 0x38
		i++
		if m.BoolValue {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	return i, nil
}

func (m *Feature) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Feature) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Id != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintGeobuf(dAtA, i, uint64(m.Id))
	}
	if len(m.Properties) > 0 {
		for k, _ := range m.Properties {
			dAtA[i] = 0x12
			i++
			v := m.Properties[k]
			msgSize := 0
			if v != nil {
				msgSize = v.Size()
				msgSize += 1 + sovGeobuf(uint64(msgSize))
			}
			mapSize := 1 + len(k) + sovGeobuf(uint64(len(k))) + msgSize
			i = encodeVarintGeobuf(dAtA, i, uint64(mapSize))
			dAtA[i] = 0xa
			i++
			i = encodeVarintGeobuf(dAtA, i, uint64(len(k)))
			i += copy(dAtA[i:], k)
			if v != nil {
				dAtA[i] = 0x12
				i++
				i = encodeVarintGeobuf(dAtA, i, uint64(v.Size()))
				n1, err := v.MarshalTo(dAtA[i:])
				if err != nil {
					return 0, err
				}
				i += n1
			}
		}
	}
	if m.Type != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintGeobuf(dAtA, i, uint64(m.Type))
	}
	if len(m.Geometry) > 0 {
		dAtA3 := make([]byte, len(m.Geometry)*10)
		var j2 int
		for _, num := range m.Geometry {
			for num >= 1<<7 {
				dAtA3[j2] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j2++
			}
			dAtA3[j2] = uint8(num)
			j2++
		}
		dAtA[i] = 0x22
		i++
		i = encodeVarintGeobuf(dAtA, i, uint64(j2))
		i += copy(dAtA[i:], dAtA3[:j2])
	}
	if len(m.BoundingBox) > 0 {
		dAtA5 := make([]byte, len(m.BoundingBox)*10)
		var j4 int
		for _, num1 := range m.BoundingBox {
			num := uint64(num1)
			for num >= 1<<7 {
				dAtA5[j4] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j4++
			}
			dAtA5[j4] = uint8(num)
			j4++
		}
		dAtA[i] = 0x2a
		i++
		i = encodeVarintGeobuf(dAtA, i, uint64(j4))
		i += copy(dAtA[i:], dAtA5[:j4])
	}
	return i, nil
}

func (m *FeatureCollection) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FeatureCollection) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Features) > 0 {
		for _, msg := range m.Features {
			dAtA[i] = 0xa
			i++
			i = encodeVarintGeobuf(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func encodeVarintGeobuf(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Value) Size() (n int) {
	var l int
	_ = l
	l = len(m.StringValue)
	if l > 0 {
		n += 1 + l + sovGeobuf(uint64(l))
	}
	if m.FloatValue != 0 {
		n += 5
	}
	if m.DoubleValue != 0 {
		n += 9
	}
	if m.IntValue != 0 {
		n += 1 + sovGeobuf(uint64(m.IntValue))
	}
	if m.UintValue != 0 {
		n += 1 + sovGeobuf(uint64(m.UintValue))
	}
	if m.SintValue != 0 {
		n += 1 + sozGeobuf(uint64(m.SintValue))
	}
	if m.BoolValue {
		n += 2
	}
	return n
}

func (m *Feature) Size() (n int) {
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovGeobuf(uint64(m.Id))
	}
	if len(m.Properties) > 0 {
		for k, v := range m.Properties {
			_ = k
			_ = v
			l = 0
			if v != nil {
				l = v.Size()
				l += 1 + sovGeobuf(uint64(l))
			}
			mapEntrySize := 1 + len(k) + sovGeobuf(uint64(len(k))) + l
			n += mapEntrySize + 1 + sovGeobuf(uint64(mapEntrySize))
		}
	}
	if m.Type != 0 {
		n += 1 + sovGeobuf(uint64(m.Type))
	}
	if len(m.Geometry) > 0 {
		l = 0
		for _, e := range m.Geometry {
			l += sovGeobuf(uint64(e))
		}
		n += 1 + sovGeobuf(uint64(l)) + l
	}
	if len(m.BoundingBox) > 0 {
		l = 0
		for _, e := range m.BoundingBox {
			l += sovGeobuf(uint64(e))
		}
		n += 1 + sovGeobuf(uint64(l)) + l
	}
	return n
}

func (m *FeatureCollection) Size() (n int) {
	var l int
	_ = l
	if len(m.Features) > 0 {
		for _, e := range m.Features {
			l = e.Size()
			n += 1 + l + sovGeobuf(uint64(l))
		}
	}
	return n
}

func sovGeobuf(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozGeobuf(x uint64) (n int) {
	return sovGeobuf(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Value) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGeobuf
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Value: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Value: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StringValue", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGeobuf
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGeobuf
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StringValue = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 5 {
				return fmt.Errorf("proto: wrong wireType = %d for field FloatValue", wireType)
			}
			var v uint32
			if (iNdEx + 4) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint32(binary.LittleEndian.Uint32(dAtA[iNdEx:]))
			iNdEx += 4
			m.FloatValue = float32(math.Float32frombits(v))
		case 3:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field DoubleValue", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.DoubleValue = float64(math.Float64frombits(v))
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IntValue", wireType)
			}
			m.IntValue = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGeobuf
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.IntValue |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UintValue", wireType)
			}
			m.UintValue = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGeobuf
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UintValue |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SintValue", wireType)
			}
			var v uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGeobuf
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			v = (v >> 1) ^ uint64((int64(v&1)<<63)>>63)
			m.SintValue = int64(v)
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BoolValue", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGeobuf
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.BoolValue = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipGeobuf(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGeobuf
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
func (m *Feature) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGeobuf
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Feature: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Feature: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGeobuf
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Properties", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGeobuf
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGeobuf
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Properties == nil {
				m.Properties = make(map[string]*Value)
			}
			var mapkey string
			var mapvalue *Value
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowGeobuf
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					wire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				fieldNum := int32(wire >> 3)
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowGeobuf
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= (uint64(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthGeobuf
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var mapmsglen int
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowGeobuf
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapmsglen |= (int(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					if mapmsglen < 0 {
						return ErrInvalidLengthGeobuf
					}
					postmsgIndex := iNdEx + mapmsglen
					if mapmsglen < 0 {
						return ErrInvalidLengthGeobuf
					}
					if postmsgIndex > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = &Value{}
					if err := mapvalue.Unmarshal(dAtA[iNdEx:postmsgIndex]); err != nil {
						return err
					}
					iNdEx = postmsgIndex
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipGeobuf(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if skippy < 0 {
						return ErrInvalidLengthGeobuf
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Properties[mapkey] = mapvalue
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGeobuf
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= (GeomType(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType == 0 {
				var v uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowGeobuf
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.Geometry = append(m.Geometry, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowGeobuf
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= (int(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthGeobuf
				}
				postIndex := iNdEx + packedLen
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				for iNdEx < postIndex {
					var v uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowGeobuf
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= (uint64(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.Geometry = append(m.Geometry, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Geometry", wireType)
			}
		case 5:
			if wireType == 0 {
				var v int64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowGeobuf
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= (int64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.BoundingBox = append(m.BoundingBox, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowGeobuf
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= (int(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthGeobuf
				}
				postIndex := iNdEx + packedLen
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				for iNdEx < postIndex {
					var v int64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowGeobuf
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= (int64(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.BoundingBox = append(m.BoundingBox, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field BoundingBox", wireType)
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGeobuf(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGeobuf
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
func (m *FeatureCollection) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGeobuf
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: FeatureCollection: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FeatureCollection: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Features", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGeobuf
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGeobuf
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Features = append(m.Features, &Feature{})
			if err := m.Features[len(m.Features)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGeobuf(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGeobuf
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
func skipGeobuf(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGeobuf
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
					return 0, ErrIntOverflowGeobuf
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGeobuf
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthGeobuf
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowGeobuf
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipGeobuf(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthGeobuf = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGeobuf   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("geobuf.proto", fileDescriptorGeobuf) }

var fileDescriptorGeobuf = []byte{
	// 431 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x92, 0xcd, 0x8e, 0xd3, 0x30,
	0x14, 0x85, 0xe7, 0xe6, 0xa7, 0x4d, 0x6e, 0x46, 0x25, 0x78, 0x15, 0x01, 0x13, 0x4c, 0x35, 0x0b,
	0x8b, 0x45, 0x16, 0x65, 0x33, 0xb0, 0x41, 0x2a, 0x2a, 0x55, 0xc4, 0x28, 0xad, 0x4c, 0x01, 0xb1,
	0x42, 0x2d, 0xf5, 0x54, 0x11, 0x21, 0xae, 0x12, 0x07, 0x91, 0xb7, 0xe0, 0xb1, 0x58, 0xf2, 0x08,
	0xa8, 0x6c, 0x78, 0x0b, 0x50, 0x12, 0x4f, 0x1b, 0xcd, 0x2e, 0xf9, 0xce, 0xf1, 0x95, 0x3f, 0x5d,
	0xe3, 0xf9, 0x4e, 0xc8, 0x4d, 0x75, 0x13, 0xed, 0x0b, 0xa9, 0xe4, 0xf8, 0x2f, 0xa0, 0xfd, 0x7e,
	0x9d, 0x55, 0x82, 0x3c, 0xc1, 0xf3, 0x52, 0x15, 0x69, 0xbe, 0xfb, 0xf4, 0xad, 0xf9, 0x0f, 0x80,
	0x02, 0x73, 0xb9, 0xd7, 0xb1, 0xae, 0xf2, 0x18, 0xbd, 0x9b, 0x4c, 0xae, 0x95, 0x6e, 0x18, 0x14,
	0x98, 0xc1, 0xb1, 0x45, 0xc7, 0x19, 0x5b, 0x59, 0x6d, 0x32, 0xa1, 0x1b, 0x26, 0x05, 0x06, 0xdc,
	0xeb, 0x58, 0x57, 0x79, 0x88, 0x6e, 0x9a, 0xdf, 0x4e, 0xb0, 0x28, 0x30, 0x93, 0x3b, 0x69, 0xae,
	0xcf, 0x5f, 0x20, 0x56, 0xa7, 0xd4, 0xa6, 0xc0, 0x2c, 0xee, 0x56, 0xfd, 0xb8, 0x3c, 0xc5, 0x03,
	0x0a, 0x8c, 0x70, 0xb7, 0xec, 0xc7, 0x1b, 0x29, 0x33, 0x1d, 0x0f, 0x29, 0x30, 0x87, 0xbb, 0x0d,
	0x69, 0xe3, 0xf1, 0x3f, 0xc0, 0xe1, 0x6b, 0xb1, 0x56, 0x55, 0x21, 0xc8, 0x08, 0x8d, 0x78, 0xdb,
	0x2a, 0x5a, 0xdc, 0x88, 0xb7, 0xe4, 0x0a, 0x71, 0x59, 0xc8, 0xbd, 0x28, 0x54, 0x2a, 0xca, 0xc0,
	0xa0, 0x26, 0xf3, 0x26, 0x41, 0xa4, 0xdb, 0xd1, 0x29, 0x9a, 0xe5, 0xaa, 0xa8, 0x79, 0xaf, 0x4b,
	0x2e, 0xd0, 0x52, 0xf5, 0xbe, 0x53, 0x1d, 0x4d, 0xdc, 0x68, 0x2e, 0xe4, 0xd7, 0x55, 0xbd, 0x17,
	0xbc, 0xc5, 0x24, 0x44, 0xa7, 0x21, 0x42, 0x15, 0x75, 0x60, 0x51, 0x93, 0x59, 0x53, 0xc3, 0x07,
	0x7e, 0x64, 0xe4, 0x12, 0xbd, 0xa9, 0xac, 0xf2, 0x6d, 0x9a, 0xef, 0xa6, 0xf2, 0x7b, 0x60, 0x53,
	0x93, 0x99, 0x6d, 0xa5, 0x8f, 0x1f, 0xcc, 0xf0, 0xde, 0x9d, 0x3b, 0x10, 0x1f, 0xcd, 0x2f, 0xa2,
	0xd6, 0x5b, 0x6a, 0x3e, 0xc9, 0x23, 0xb4, 0x4f, 0x7b, 0xf1, 0x26, 0x83, 0xa8, 0xd5, 0xe6, 0x1d,
	0x7c, 0x61, 0x5c, 0xc1, 0xf8, 0x39, 0xde, 0xd7, 0x4a, 0xaf, 0x64, 0x96, 0x89, 0xcf, 0x2a, 0x95,
	0x39, 0xb9, 0x44, 0x47, 0xc3, 0x32, 0x80, 0x56, 0xdc, 0xb9, 0x15, 0xe7, 0xc7, 0xe4, 0xe9, 0xcb,
	0xce, 0xa3, 0x31, 0x23, 0x1e, 0x0e, 0xdf, 0x25, 0x6f, 0x92, 0xc5, 0x87, 0xc4, 0x3f, 0x23, 0x2e,
	0xda, 0xcb, 0x45, 0x9c, 0xac, 0x7c, 0x20, 0x23, 0xc4, 0xeb, 0x38, 0x99, 0xbd, 0x5d, 0xf1, 0x38,
	0x99, 0xfb, 0x46, 0xd3, 0x5b, 0x2e, 0xae, 0x3f, 0xce, 0x17, 0x89, 0x6f, 0x4e, 0xfd, 0x9f, 0x87,
	0x10, 0x7e, 0x1d, 0x42, 0xf8, 0x7d, 0x08, 0xe1, 0xc7, 0x9f, 0xf0, 0x6c, 0x33, 0x68, 0x5f, 0xe0,
	0xb3, 0xff, 0x01, 0x00, 0x00, 0xff, 0xff, 0x62, 0x77, 0x0c, 0x95, 0x91, 0x02, 0x00, 0x00,
}
