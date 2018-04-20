# messagepack-vs-protobuf
Comparing the MessagePack and Protocol Buffers formats for size and ease of decoding/encoding.

| Value (hex) | Value (dec) | MessagePack (bytes) | Proto. Buffers (bytes) |
| :---: | :---: | :---: | :---: |
| 00-7F | 0-127 | 1 | 1 |
| 80-FF | 128-255 | 2 | 2 |
| 100-3FFF | 256-16,383 | 3 | 2 |
| 4000-FFFF | 16,384-65,535 | 3 | 3 |

Protocol Buffers turns out to encode many but not all unsigned integers more compactly, that is, in the ranges: 256-16,383; 65,536-268,435,455; and 4,294,967,296-72,057,594,037,927,935).
