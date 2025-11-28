package rmproto

func CRC8(data []byte) byte {
	var crc byte = 0x00
	for _, b := range data {
		crc ^= b
		for i := 0; i < 8; i++ {
			if (crc & 0x80) != 0 {
				crc = (crc << 1) ^ 0x31 // Polynomial x^8 + x^5 + x^4 + 1
			} else {
				crc <<= 1
			}
		}
	}
	return crc
}

func CRC16(data []byte) uint16 {
	var crc uint16 = 0xFFFF
	for _, b := range data {
		crc ^= uint16(b)
		for i := 0; i < 8; i++ {
			if (crc & 0x0001) != 0 {
				crc = (crc >> 1) ^ 0xA001 // Polynomial x^16 + x^15 + x^2 + 1
			} else {
				crc >>= 1
			}
		}
	}
	return crc
}
