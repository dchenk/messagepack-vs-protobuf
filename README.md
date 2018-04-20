# messagepack-vs-protobuf
Comparing the MessagePack and Protocol Buffers formats for size and ease of decoding/encoding.

#### Encoded Size for Unsigned Integers
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

Protocol Buffers turns out to encode many but not all unsigned integers more compactly, that is, in the ranges: 256-16,383; 65,536-268,435,455; and 4,294,967,296-72,057,594,037,927,935).
