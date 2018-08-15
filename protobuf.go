package main

var InnerStructB_inst = InnerStructB{
	Bool:    true,
	Float64: -234.52,
	Int8:    23,
	Uint64S: []uint64{898998, 3, 922337203685477577, 5235, 0, 3424, 555, 325, 65748987656, 98989283, 18446744073709551615, 9223372036854775807},
}

var StructB_inst = StructB{
	Str:   "hello+lk\n\nsdj flasjfdlkaj\nsdlf sd skd akdska\tjs\nldfk\"jal kas_ ad\"fsfka\tjs\nldfk\"jal kas_ ad\"fska\tjs\nldfk\"jal kas_ adk",
	Slice: []byte{'a', 'b', 'c', '7', 0x03, 0x05},
	MapStrUint16: map[string]uint32{
		"abc":                  23,
		"lasdhflahsdlfkhlasdf": 4,
		"c":                    1356,
		"another":              23422,
		"abcd":                 12525,
	},
	Inner: &InnerStructB{Bool: false, Float64: 213.513, Int8: 23, Uint64S: []uint64{7869, 0, 23424, 211, 970372, 7324743636, 32, 11, 52355, 113, 3255}},
	Bt:    87,
	Int:   82,
	Int64: -27823325,
	Strs: []string{"", "ab", "abcdefg hijk", "j\nsdlfka\tjs\nldfk\"jal kas_ ad\"fsfka", "ab", "abcdefg hijk", "j\nsdlfka\tjs\nldfk\"jal kas_ ad\"fsfka",
		"ab", "abcdefg hijk", "j\nsdlfka\tjs\nldfk\"jal kas_ ad\"fsfka"},
	SliceI64:     []int64{898, 9223372036854775807, -234324, 0, 235, 11, 0, 92233720368547758},
	MapStrUint64: map[string]uint64{"jkjkjk kkjk": 2342, "kjkj akjsdfla lajsdf lakjd laskjdf alsdjkf": 9923848234},
	Recursive: []*StructB{
		{Str: "hello there!", Inner: &InnerStructB_inst, Int64: 89324723752735, F32: 44.0},
		{Str: "and hello there!", Int64: -24723752735, F32: 4.9},
	},
	Float64: 3.14159265359,
	MapStrSlice: map[string]*Int32S{
		"":                                     nil,
		"uiouoiyioutuuiouoiyioutuuiouoiyioutu": {Ints: []int32{0x32, 0x12, 0x78, 0x00, 0x76}},
		"RTQYWYEUIIUQIE": {Ints: []int32{0x32, 0x12, 0x78, 0x00, 0x76, 0x32, 0x12, 0x78, 0x00, 0x76, 0x32, 0x12, 0x78, 0x76,
			0x32, 0x12, 0x78, 0x00, 0x76, 0x32, 0x12, 0x78, 0x00, 0x76, 0x32, 0x12, 0x78, 0x76, 0x32, 0x12, 0x78, 0x00, 0x76, 0x32,
			0x12, 0x78, 0x00, 0x76, 0x32, 0x12, 0x78, 0x76, 0x32, 0x12, 0x78, 0x00, 0x76, 0x32, 0x12, 0x78, 0x00, 0x76, 0x32, 0x12}},
	},
	F32: 536.4,
}
