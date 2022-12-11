package stm32crc

//region #### CRC32STM32

var crc32table [256]uint32

func Init() {
	var poly uint32 = 0x04C11DB7
	var i uint32
	var j uint32
	for i = 0; i < 256; i++ {
		c := i << 24
		for j = 0; j < 8; j++ {
			if (c & 0x80000000) > 0 {
				c = (c << 1) ^ poly
			} else {
				c = c << 1
			}
		}
		crc32table[i] = c & 0xffffffff
	}
}

func Crc32(bytesArr []byte) uint32 {

	length := len(bytesArr)
	var crc uint32 = 0xffffffff

	k := 0
	for length >= 4 {
		var v uint32
		v = ((uint32(bytesArr[k]) << 24) & 0xFF000000) | ((uint32(bytesArr[k+1]) << 16) & 0xFF0000) |
			((uint32(bytesArr[k+2]) << 8) & 0xFF00) | (uint32(bytesArr[k+3]) & 0xFF)

		crc = ((crc << 8) & 0xffffffff) ^ crc32table[0xFF&((crc>>24)^v)]
		crc = ((crc << 8) & 0xffffffff) ^ crc32table[0xFF&((crc>>24)^(v>>8))]
		crc = ((crc << 8) & 0xffffffff) ^ crc32table[0xFF&((crc>>24)^(v>>16))]
		crc = ((crc << 8) & 0xffffffff) ^ crc32table[0xFF&((crc>>24)^(v>>24))]

		k += 4
		length -= 4
	}

	if length > 0 {
		var v uint32 = 0

		for i := 0; i < length; i++ {
			v |= uint32(bytesArr[k+i])<<24 - uint32(i)*8
		}

		if length == 1 {
			v &= 0xFF000000
		} else if length == 2 {
			v &= 0xFFFF0000
		} else if length == 3 {
			v &= 0xFFFFFF00
		}

		crc = ((crc << 8) & 0xffffffff) ^ crc32table[0xFF&((crc>>24)^(v))]
		crc = ((crc << 8) & 0xffffffff) ^ crc32table[0xFF&((crc>>24)^(v>>8))]
		crc = ((crc << 8) & 0xffffffff) ^ crc32table[0xFF&((crc>>24)^(v>>16))]
		crc = ((crc << 8) & 0xffffffff) ^ crc32table[0xFF&((crc>>24)^(v>>24))]
	}

	return crc

}

//endregion
