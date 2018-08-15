package main

// THIS FILE WAS PRODUCED BY THE MSGP CODE GENERATION TOOL (github.com/dchenk/msgp).
// DO NOT EDIT.

import (
	"github.com/dchenk/msgp/msgp"
)

// DecodeMsg implements msgp.Decoder
func (z *InnerStructA) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch string(field) {
		case "bool":
			z.Bool, err = dc.ReadBool()
			if err != nil {
				return
			}
		case "f64":
			z.Float64, err = dc.ReadFloat64()
			if err != nil {
				return
			}
		case "i8":
			z.Int8, err = dc.ReadInt8()
			if err != nil {
				return
			}
		case "u64s":
			var zb0002 uint32
			zb0002, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Uint64s) >= int(zb0002) {
				z.Uint64s = (z.Uint64s)[:zb0002]
			} else {
				z.Uint64s = make([]uint64, zb0002)
			}
			for za0001 := range z.Uint64s {
				z.Uint64s[za0001], err = dc.ReadUint64()
				if err != nil {
					return
				}
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encoder
func (z *InnerStructA) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 4
	// write "bool"
	err = en.Append(0x84, 0xa4, 0x62, 0x6f, 0x6f, 0x6c)
	if err != nil {
		return
	}
	err = en.WriteBool(z.Bool)
	if err != nil {
		return
	}
	// write "f64"
	err = en.Append(0xa3, 0x66, 0x36, 0x34)
	if err != nil {
		return
	}
	err = en.WriteFloat64(z.Float64)
	if err != nil {
		return
	}
	// write "i8"
	err = en.Append(0xa2, 0x69, 0x38)
	if err != nil {
		return
	}
	err = en.WriteInt8(z.Int8)
	if err != nil {
		return
	}
	// write "u64s"
	err = en.Append(0xa4, 0x75, 0x36, 0x34, 0x73)
	if err != nil {
		return
	}
	err = en.WriteArrayHeader(uint32(len(z.Uint64s)))
	if err != nil {
		return
	}
	for za0001 := range z.Uint64s {
		err = en.WriteUint64(z.Uint64s[za0001])
		if err != nil {
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *InnerStructA) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 4
	// string "bool"
	o = append(o, 0x84, 0xa4, 0x62, 0x6f, 0x6f, 0x6c)
	o = msgp.AppendBool(o, z.Bool)
	// string "f64"
	o = append(o, 0xa3, 0x66, 0x36, 0x34)
	o = msgp.AppendFloat64(o, z.Float64)
	// string "i8"
	o = append(o, 0xa2, 0x69, 0x38)
	o = msgp.AppendInt8(o, z.Int8)
	// string "u64s"
	o = append(o, 0xa4, 0x75, 0x36, 0x34, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Uint64s)))
	for za0001 := range z.Uint64s {
		o = msgp.AppendUint64(o, z.Uint64s[za0001])
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *InnerStructA) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch string(field) {
		case "bool":
			z.Bool, bts, err = msgp.ReadBoolBytes(bts)
			if err != nil {
				return
			}
		case "f64":
			z.Float64, bts, err = msgp.ReadFloat64Bytes(bts)
			if err != nil {
				return
			}
		case "i8":
			z.Int8, bts, err = msgp.ReadInt8Bytes(bts)
			if err != nil {
				return
			}
		case "u64s":
			var zb0002 uint32
			zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Uint64s) >= int(zb0002) {
				z.Uint64s = (z.Uint64s)[:zb0002]
			} else {
				z.Uint64s = make([]uint64, zb0002)
			}
			for za0001 := range z.Uint64s {
				z.Uint64s[za0001], bts, err = msgp.ReadUint64Bytes(bts)
				if err != nil {
					return
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *InnerStructA) Msgsize() (s int) {
	s = 1 + 5 + msgp.BoolSize + 4 + msgp.Float64Size + 3 + msgp.Int8Size + 5 + msgp.ArrayHeaderSize + (len(z.Uint64s) * (msgp.Uint64Size))
	return
}

// DecodeMsg implements msgp.Decoder
func (z *StructA) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch string(field) {
		case "str":
			z.Str, err = dc.ReadString()
			if err != nil {
				return
			}
		case "slice":
			z.Slice, err = dc.ReadBytes(z.Slice)
			if err != nil {
				return
			}
		case "map_str_u":
			var zb0002 uint32
			zb0002, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.MapStrUint16 == nil && zb0002 > 0 {
				z.MapStrUint16 = make(map[string]uint16, zb0002)
			} else if len(z.MapStrUint16) > 0 {
				for key := range z.MapStrUint16 {
					delete(z.MapStrUint16, key)
				}
			}
			for zb0002 > 0 {
				zb0002--
				var za0001 string
				var za0002 uint16
				za0001, err = dc.ReadString()
				if err != nil {
					return
				}
				za0002, err = dc.ReadUint16()
				if err != nil {
					return
				}
				z.MapStrUint16[za0001] = za0002
			}
		case "inner_struct":
			err = z.Inner.DecodeMsg(dc)
			if err != nil {
				return
			}
		case "bt":
			z.Bt, err = dc.ReadUint8()
			if err != nil {
				return
			}
		case "int":
			z.Int, err = dc.ReadInt32()
			if err != nil {
				return
			}
		case "i64":
			z.Int64, err = dc.ReadInt64()
			if err != nil {
				return
			}
		case "strs":
			var zb0003 uint32
			zb0003, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Strs) >= int(zb0003) {
				z.Strs = (z.Strs)[:zb0003]
			} else {
				z.Strs = make([]string, zb0003)
			}
			for za0003 := range z.Strs {
				z.Strs[za0003], err = dc.ReadString()
				if err != nil {
					return
				}
			}
		case "slice_i64":
			var zb0004 uint32
			zb0004, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.SliceI64) >= int(zb0004) {
				z.SliceI64 = (z.SliceI64)[:zb0004]
			} else {
				z.SliceI64 = make([]int64, zb0004)
			}
			for za0004 := range z.SliceI64 {
				z.SliceI64[za0004], err = dc.ReadInt64()
				if err != nil {
					return
				}
			}
		case "mp":
			var zb0005 uint32
			zb0005, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.MapStrUint64 == nil && zb0005 > 0 {
				z.MapStrUint64 = make(map[string]uint64, zb0005)
			} else if len(z.MapStrUint64) > 0 {
				for key := range z.MapStrUint64 {
					delete(z.MapStrUint64, key)
				}
			}
			for zb0005 > 0 {
				zb0005--
				var za0005 string
				var za0006 uint64
				za0005, err = dc.ReadString()
				if err != nil {
					return
				}
				za0006, err = dc.ReadUint64()
				if err != nil {
					return
				}
				z.MapStrUint64[za0005] = za0006
			}
		case "recursive":
			var zb0006 uint32
			zb0006, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Recursive) >= int(zb0006) {
				z.Recursive = (z.Recursive)[:zb0006]
			} else {
				z.Recursive = make([]StructA, zb0006)
			}
			for za0007 := range z.Recursive {
				err = z.Recursive[za0007].DecodeMsg(dc)
				if err != nil {
					return
				}
			}
		case "f64":
			z.Float64, err = dc.ReadFloat64()
			if err != nil {
				return
			}
		case "map2":
			var zb0007 uint32
			zb0007, err = dc.ReadMapHeader()
			if err != nil {
				return
			}
			if z.MapStrSlice == nil && zb0007 > 0 {
				z.MapStrSlice = make(map[string][]int8, zb0007)
			} else if len(z.MapStrSlice) > 0 {
				for key := range z.MapStrSlice {
					delete(z.MapStrSlice, key)
				}
			}
			for zb0007 > 0 {
				zb0007--
				var za0008 string
				var za0009 []int8
				za0008, err = dc.ReadString()
				if err != nil {
					return
				}
				var zb0008 uint32
				zb0008, err = dc.ReadArrayHeader()
				if err != nil {
					return
				}
				if cap(za0009) >= int(zb0008) {
					za0009 = (za0009)[:zb0008]
				} else {
					za0009 = make([]int8, zb0008)
				}
				for za0010 := range za0009 {
					za0009[za0010], err = dc.ReadInt8()
					if err != nil {
						return
					}
				}
				z.MapStrSlice[za0008] = za0009
			}
		case "float32":
			z.F32, err = dc.ReadFloat32()
			if err != nil {
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encoder
func (z *StructA) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 14
	// write "str"
	err = en.Append(0x8e, 0xa3, 0x73, 0x74, 0x72)
	if err != nil {
		return
	}
	err = en.WriteString(z.Str)
	if err != nil {
		return
	}
	// write "slice"
	err = en.Append(0xa5, 0x73, 0x6c, 0x69, 0x63, 0x65)
	if err != nil {
		return
	}
	err = en.WriteBytes(z.Slice)
	if err != nil {
		return
	}
	// write "map_str_u"
	err = en.Append(0xa9, 0x6d, 0x61, 0x70, 0x5f, 0x73, 0x74, 0x72, 0x5f, 0x75)
	if err != nil {
		return
	}
	err = en.WriteMapHeader(uint32(len(z.MapStrUint16)))
	if err != nil {
		return
	}
	for za0001, za0002 := range z.MapStrUint16 {
		err = en.WriteString(za0001)
		if err != nil {
			return
		}
		err = en.WriteUint16(za0002)
		if err != nil {
			return
		}
	}
	// write "inner_struct"
	err = en.Append(0xac, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x5f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74)
	if err != nil {
		return
	}
	err = z.Inner.EncodeMsg(en)
	if err != nil {
		return
	}
	// write "bt"
	err = en.Append(0xa2, 0x62, 0x74)
	if err != nil {
		return
	}
	err = en.WriteUint8(z.Bt)
	if err != nil {
		return
	}
	// write "int"
	err = en.Append(0xa3, 0x69, 0x6e, 0x74)
	if err != nil {
		return
	}
	err = en.WriteInt32(z.Int)
	if err != nil {
		return
	}
	// write "i64"
	err = en.Append(0xa3, 0x69, 0x36, 0x34)
	if err != nil {
		return
	}
	err = en.WriteInt64(z.Int64)
	if err != nil {
		return
	}
	// write "strs"
	err = en.Append(0xa4, 0x73, 0x74, 0x72, 0x73)
	if err != nil {
		return
	}
	err = en.WriteArrayHeader(uint32(len(z.Strs)))
	if err != nil {
		return
	}
	for za0003 := range z.Strs {
		err = en.WriteString(z.Strs[za0003])
		if err != nil {
			return
		}
	}
	// write "slice_i64"
	err = en.Append(0xa9, 0x73, 0x6c, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x36, 0x34)
	if err != nil {
		return
	}
	err = en.WriteArrayHeader(uint32(len(z.SliceI64)))
	if err != nil {
		return
	}
	for za0004 := range z.SliceI64 {
		err = en.WriteInt64(z.SliceI64[za0004])
		if err != nil {
			return
		}
	}
	// write "mp"
	err = en.Append(0xa2, 0x6d, 0x70)
	if err != nil {
		return
	}
	err = en.WriteMapHeader(uint32(len(z.MapStrUint64)))
	if err != nil {
		return
	}
	for za0005, za0006 := range z.MapStrUint64 {
		err = en.WriteString(za0005)
		if err != nil {
			return
		}
		err = en.WriteUint64(za0006)
		if err != nil {
			return
		}
	}
	// write "recursive"
	err = en.Append(0xa9, 0x72, 0x65, 0x63, 0x75, 0x72, 0x73, 0x69, 0x76, 0x65)
	if err != nil {
		return
	}
	err = en.WriteArrayHeader(uint32(len(z.Recursive)))
	if err != nil {
		return
	}
	for za0007 := range z.Recursive {
		err = z.Recursive[za0007].EncodeMsg(en)
		if err != nil {
			return
		}
	}
	// write "f64"
	err = en.Append(0xa3, 0x66, 0x36, 0x34)
	if err != nil {
		return
	}
	err = en.WriteFloat64(z.Float64)
	if err != nil {
		return
	}
	// write "map2"
	err = en.Append(0xa4, 0x6d, 0x61, 0x70, 0x32)
	if err != nil {
		return
	}
	err = en.WriteMapHeader(uint32(len(z.MapStrSlice)))
	if err != nil {
		return
	}
	for za0008, za0009 := range z.MapStrSlice {
		err = en.WriteString(za0008)
		if err != nil {
			return
		}
		err = en.WriteArrayHeader(uint32(len(za0009)))
		if err != nil {
			return
		}
		for za0010 := range za0009 {
			err = en.WriteInt8(za0009[za0010])
			if err != nil {
				return
			}
		}
	}
	// write "float32"
	err = en.Append(0xa7, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x33, 0x32)
	if err != nil {
		return
	}
	err = en.WriteFloat32(z.F32)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *StructA) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 14
	// string "str"
	o = append(o, 0x8e, 0xa3, 0x73, 0x74, 0x72)
	o = msgp.AppendString(o, z.Str)
	// string "slice"
	o = append(o, 0xa5, 0x73, 0x6c, 0x69, 0x63, 0x65)
	o = msgp.AppendBytes(o, z.Slice)
	// string "map_str_u"
	o = append(o, 0xa9, 0x6d, 0x61, 0x70, 0x5f, 0x73, 0x74, 0x72, 0x5f, 0x75)
	o = msgp.AppendMapHeader(o, uint32(len(z.MapStrUint16)))
	for za0001, za0002 := range z.MapStrUint16 {
		o = msgp.AppendString(o, za0001)
		o = msgp.AppendUint16(o, za0002)
	}
	// string "inner_struct"
	o = append(o, 0xac, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x5f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74)
	o, err = z.Inner.MarshalMsg(o)
	if err != nil {
		return
	}
	// string "bt"
	o = append(o, 0xa2, 0x62, 0x74)
	o = msgp.AppendUint8(o, z.Bt)
	// string "int"
	o = append(o, 0xa3, 0x69, 0x6e, 0x74)
	o = msgp.AppendInt32(o, z.Int)
	// string "i64"
	o = append(o, 0xa3, 0x69, 0x36, 0x34)
	o = msgp.AppendInt64(o, z.Int64)
	// string "strs"
	o = append(o, 0xa4, 0x73, 0x74, 0x72, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Strs)))
	for za0003 := range z.Strs {
		o = msgp.AppendString(o, z.Strs[za0003])
	}
	// string "slice_i64"
	o = append(o, 0xa9, 0x73, 0x6c, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x36, 0x34)
	o = msgp.AppendArrayHeader(o, uint32(len(z.SliceI64)))
	for za0004 := range z.SliceI64 {
		o = msgp.AppendInt64(o, z.SliceI64[za0004])
	}
	// string "mp"
	o = append(o, 0xa2, 0x6d, 0x70)
	o = msgp.AppendMapHeader(o, uint32(len(z.MapStrUint64)))
	for za0005, za0006 := range z.MapStrUint64 {
		o = msgp.AppendString(o, za0005)
		o = msgp.AppendUint64(o, za0006)
	}
	// string "recursive"
	o = append(o, 0xa9, 0x72, 0x65, 0x63, 0x75, 0x72, 0x73, 0x69, 0x76, 0x65)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Recursive)))
	for za0007 := range z.Recursive {
		o, err = z.Recursive[za0007].MarshalMsg(o)
		if err != nil {
			return
		}
	}
	// string "f64"
	o = append(o, 0xa3, 0x66, 0x36, 0x34)
	o = msgp.AppendFloat64(o, z.Float64)
	// string "map2"
	o = append(o, 0xa4, 0x6d, 0x61, 0x70, 0x32)
	o = msgp.AppendMapHeader(o, uint32(len(z.MapStrSlice)))
	for za0008, za0009 := range z.MapStrSlice {
		o = msgp.AppendString(o, za0008)
		o = msgp.AppendArrayHeader(o, uint32(len(za0009)))
		for za0010 := range za0009 {
			o = msgp.AppendInt8(o, za0009[za0010])
		}
	}
	// string "float32"
	o = append(o, 0xa7, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x33, 0x32)
	o = msgp.AppendFloat32(o, z.F32)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *StructA) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch string(field) {
		case "str":
			z.Str, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "slice":
			z.Slice, bts, err = msgp.ReadBytesBytes(bts, z.Slice)
			if err != nil {
				return
			}
		case "map_str_u":
			var zb0002 uint32
			zb0002, bts, err = msgp.ReadMapHeaderBytes(bts)
			if err != nil {
				return
			}
			if z.MapStrUint16 == nil && zb0002 > 0 {
				z.MapStrUint16 = make(map[string]uint16, zb0002)
			} else if len(z.MapStrUint16) > 0 {
				for key := range z.MapStrUint16 {
					delete(z.MapStrUint16, key)
				}
			}
			for zb0002 > 0 {
				var za0001 string
				var za0002 uint16
				zb0002--
				za0001, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
				za0002, bts, err = msgp.ReadUint16Bytes(bts)
				if err != nil {
					return
				}
				z.MapStrUint16[za0001] = za0002
			}
		case "inner_struct":
			bts, err = z.Inner.UnmarshalMsg(bts)
			if err != nil {
				return
			}
		case "bt":
			z.Bt, bts, err = msgp.ReadUint8Bytes(bts)
			if err != nil {
				return
			}
		case "int":
			z.Int, bts, err = msgp.ReadInt32Bytes(bts)
			if err != nil {
				return
			}
		case "i64":
			z.Int64, bts, err = msgp.ReadInt64Bytes(bts)
			if err != nil {
				return
			}
		case "strs":
			var zb0003 uint32
			zb0003, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Strs) >= int(zb0003) {
				z.Strs = (z.Strs)[:zb0003]
			} else {
				z.Strs = make([]string, zb0003)
			}
			for za0003 := range z.Strs {
				z.Strs[za0003], bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
			}
		case "slice_i64":
			var zb0004 uint32
			zb0004, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.SliceI64) >= int(zb0004) {
				z.SliceI64 = (z.SliceI64)[:zb0004]
			} else {
				z.SliceI64 = make([]int64, zb0004)
			}
			for za0004 := range z.SliceI64 {
				z.SliceI64[za0004], bts, err = msgp.ReadInt64Bytes(bts)
				if err != nil {
					return
				}
			}
		case "mp":
			var zb0005 uint32
			zb0005, bts, err = msgp.ReadMapHeaderBytes(bts)
			if err != nil {
				return
			}
			if z.MapStrUint64 == nil && zb0005 > 0 {
				z.MapStrUint64 = make(map[string]uint64, zb0005)
			} else if len(z.MapStrUint64) > 0 {
				for key := range z.MapStrUint64 {
					delete(z.MapStrUint64, key)
				}
			}
			for zb0005 > 0 {
				var za0005 string
				var za0006 uint64
				zb0005--
				za0005, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
				za0006, bts, err = msgp.ReadUint64Bytes(bts)
				if err != nil {
					return
				}
				z.MapStrUint64[za0005] = za0006
			}
		case "recursive":
			var zb0006 uint32
			zb0006, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Recursive) >= int(zb0006) {
				z.Recursive = (z.Recursive)[:zb0006]
			} else {
				z.Recursive = make([]StructA, zb0006)
			}
			for za0007 := range z.Recursive {
				bts, err = z.Recursive[za0007].UnmarshalMsg(bts)
				if err != nil {
					return
				}
			}
		case "f64":
			z.Float64, bts, err = msgp.ReadFloat64Bytes(bts)
			if err != nil {
				return
			}
		case "map2":
			var zb0007 uint32
			zb0007, bts, err = msgp.ReadMapHeaderBytes(bts)
			if err != nil {
				return
			}
			if z.MapStrSlice == nil && zb0007 > 0 {
				z.MapStrSlice = make(map[string][]int8, zb0007)
			} else if len(z.MapStrSlice) > 0 {
				for key := range z.MapStrSlice {
					delete(z.MapStrSlice, key)
				}
			}
			for zb0007 > 0 {
				var za0008 string
				var za0009 []int8
				zb0007--
				za0008, bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
				var zb0008 uint32
				zb0008, bts, err = msgp.ReadArrayHeaderBytes(bts)
				if err != nil {
					return
				}
				if cap(za0009) >= int(zb0008) {
					za0009 = (za0009)[:zb0008]
				} else {
					za0009 = make([]int8, zb0008)
				}
				for za0010 := range za0009 {
					za0009[za0010], bts, err = msgp.ReadInt8Bytes(bts)
					if err != nil {
						return
					}
				}
				z.MapStrSlice[za0008] = za0009
			}
		case "float32":
			z.F32, bts, err = msgp.ReadFloat32Bytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *StructA) Msgsize() (s int) {
	s = 1 + 4 + msgp.StringPrefixSize + len(z.Str) + 6 + msgp.BytesPrefixSize + len(z.Slice) + 10 + msgp.MapHeaderSize
	if z.MapStrUint16 != nil {
		for za0001, za0002 := range z.MapStrUint16 {
			_ = za0002
			s += msgp.StringPrefixSize + len(za0001) + msgp.Uint16Size
		}
	}
	s += 13 + z.Inner.Msgsize() + 3 + msgp.Uint8Size + 4 + msgp.Int32Size + 4 + msgp.Int64Size + 5 + msgp.ArrayHeaderSize
	for za0003 := range z.Strs {
		s += msgp.StringPrefixSize + len(z.Strs[za0003])
	}
	s += 10 + msgp.ArrayHeaderSize + (len(z.SliceI64) * (msgp.Int64Size)) + 3 + msgp.MapHeaderSize
	if z.MapStrUint64 != nil {
		for za0005, za0006 := range z.MapStrUint64 {
			_ = za0006
			s += msgp.StringPrefixSize + len(za0005) + msgp.Uint64Size
		}
	}
	s += 10 + msgp.ArrayHeaderSize
	for za0007 := range z.Recursive {
		s += z.Recursive[za0007].Msgsize()
	}
	s += 4 + msgp.Float64Size + 5 + msgp.MapHeaderSize
	if z.MapStrSlice != nil {
		for za0008, za0009 := range z.MapStrSlice {
			_ = za0009
			s += msgp.StringPrefixSize + len(za0008) + msgp.ArrayHeaderSize + (len(za0009) * (msgp.Int8Size))
		}
	}
	s += 8 + msgp.Float32Size
	return
}
