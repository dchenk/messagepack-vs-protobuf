# messagepack-vs-protobuf
Comparing the MessagePack and Protocol Buffers (PB) formats for size and ease of decoding/encoding.

#### Encoded Size for Unsigned Integers (and "sint" integers)
| Value (hex) | Value (dec) | MessagePack (bytes) | Proto. Buffers (bytes) |
| :---: | :---: | :---: | :---: |
| 00-7F | 0 - 127 | 1 | 1 |
| 80-FF | 128 - 255 | 2 | 2 |
| 100-3FFF | 256 - 16,383 | 3 | 2 |
| 4000-FFFF | 16,384 - 65,535 | 3 | 3 |
| 10000-1FFFFF | 65,536 - 2,097,151 | 5 | 3 |
| 200000-FFFFFFF | 2,097,152 - 268,435,455 | 5 | 4|
| ... | 268,435,456 - 2<sup>32</sup>-1 | 5 | 5 |
| ... | 2<sup>32</sup> - 2<sup>35</sup>-1 | 9 | 5 |
| ... | 2<sup>35</sup> - 2<sup>42</sup>-1 | 9 | 6 |
| ... | 2<sup>42</sup> - 2<sup>49</sup>-1 | 9 | 7 |
| ... | 2<sup>49</sup> - 2<sup>56</sup>-1 | 9 | 8 |
| ... | 2<sup>56</sup> - 2<sup>63</sup>-1 | 9 | 9 |
| ... | 2<sup>63</sup> - 2<sup>64</sup>-1 | 9 | 10 |
* This table describes only non-negative integer values in MessagePack and only the types uint32, uint64, sint32, and sint64 in PB. The int32 and int64 types in PB always encode using 10 bytes, since the very last bit in these numbers holds the sign.

For just these types of integers, PB turns out to encode many but not all integer values more compactly, that is, values in the ranges: 256-16,383; 65,536-268,435,455; and 4,294,967,296-72,057,594,037,927,935.

Actual messages are not, of course, just variable-length integers: in MessagePack you need a "map key", and in PB you need the field number. Given a PB encoded message, you can't simply construct a meaningful JSON representation because all object keys would just be the field numbers (as strings); yet MessagePack can inter-operate with JSON using string-type keys to elements in a map/"object".

#### Fixed-length Encoded Numbers
In both the MessagePack and Protocol Buffers formats, you can have 32-bit numbers (e.g., float32) that always take up 5 bytes with the leading byte indicating the type, and 64-bit numbers (e.g., float64), which take a total of 9 bytes.
