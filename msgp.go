package main

//go:generate msgp

type InnerStructA struct {
	Bool    bool     `msgp:"bool"`
	Float64 float64  `msgp:"f64"`
	Int8    int8     `msgp:"i8"`
	Uint64s []uint64 `msgp:"u64s"`
}

type StructA struct {
	Str          string            `msgp:"str"`
	Slice        []byte            `msgp:"slice"`
	MapStrUint16 map[string]uint16 `msgp:"map_str_u"`
	Inner        InnerStructA      `msgp:"inner_struct"`
	Bt           uint8             `msgp:"bt"`
	Int          int32             `msgp:"int"`
	Int64        int64             `msgp:"i64"`
	Strs         []string          `msgp:"strs"`
	SliceI64     []int64           `msgp:"slice_i64"`
	MapStrUint64 map[string]uint64 `msgp:"mp"`
	Recursive    []StructA         `msgp:"recursive"`
	Float64      float64           `msgp:"f64"`
	MapStrSlice  map[string][]int8 `msgp:"map2"`
	F32          float32           `msgp:"float32"`
}

var InnerStructA_inst = InnerStructA{
	Bool:    true,
	Float64: -234.52,
	Int8:    23,
	Uint64s: []uint64{898998, 3, 922337203685477577, 5235, 0, 3424, 555, 325, 65748987656, 98989283, 18446744073709551615, 9223372036854775807},
}

var Struct1A_inst = StructA{
	Str:   "hello+lk\n\nsdj flasjfdlkaj\nsdlf sd skd akdska\tjs\nldfk\"jal kas_ ad\"fsfka\tjs\nldfk\"jal kas_ ad\"fska\tjs\nldfk\"jal kas_ adk",
	Slice: []byte{'a', 'b', 'c', '7', 0x03, 0x05},
	MapStrUint16: map[string]uint16{
		"abc":                  23,
		"lasdhflahsdlfkhlasdf": 4,
		"c":                    1356,
		"another":              23422,
		"abcd":                 12525,
	},
	Inner: InnerStructA{false, 213.513, 23, []uint64{7869, 0, 23424, 211, 970372, 7324743636, 32, 11, 52355, 113, 3255}},
	Bt:    87,
	Int:   82,
	Int64: -27823325,
	Strs: []string{"", "ab", "abcdefg hijk", "j\nsdlfka\tjs\nldfk\"jal kas_ ad\"fsfka", "ab", "abcdefg hijk", "j\nsdlfka\tjs\nldfk\"jal kas_ ad\"fsfka",
		"ab", "abcdefg hijk", "j\nsdlfka\tjs\nldfk\"jal kas_ ad\"fsfka"},
	SliceI64:     []int64{898, 9223372036854775807, -234324, 0, 235, 11, 0, 92233720368547758},
	MapStrUint64: map[string]uint64{"jkjkjk kkjk": 2342, "kjkj akjsdfla lajsdf lakjd laskjdf alsdjkf": 9923848234},
	Recursive: []StructA{
		{Str: "hello there!", Inner: InnerStructA_inst, Int64: 89324723752735, F32: 44.0},
		{Str: "and hello there!", Int64: -24723752735, F32: 4.9},
	},
	Float64: 3.14159265359,
	MapStrSlice: map[string][]int8{
		"":                                     nil,
		"uiouoiyioutuuiouoiyioutuuiouoiyioutu": {0x32, 0x12, 0x78, 0x00, 0x76},
		"RTQYWYEUIIUQIE": {0x32, 0x12, 0x78, 0x00, 0x76, 0x32, 0x12, 0x78, 0x00, 0x76, 0x32, 0x12, 0x78, 0x76,
			0x32, 0x12, 0x78, 0x00, 0x76, 0x32, 0x12, 0x78, 0x00, 0x76, 0x32, 0x12, 0x78, 0x76, 0x32, 0x12, 0x78, 0x00, 0x76, 0x32,
			0x12, 0x78, 0x00, 0x76, 0x32, 0x12, 0x78, 0x76, 0x32, 0x12, 0x78, 0x00, 0x76, 0x32, 0x12, 0x78, 0x00, 0x76, 0x32, 0x12},
	},
	F32: 536.4,
}
