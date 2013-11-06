package generator

import (
	"bytes"
	"compress/gzip"
	"io"
)

// decoder_tmpl returns raw, uncompressed file data.
func decoder_tmpl() []byte {
	gz, err := gzip.NewReader(bytes.NewBuffer([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0xac, 0x56,
		0xdf, 0x6f, 0xdb, 0x36, 0x10, 0x7e, 0x96, 0xfe, 0x8a, 0xab, 0x80, 0x34,
		0x52, 0x1a, 0xc8, 0xc5, 0x56, 0xf4, 0x21, 0x83, 0x1f, 0xda, 0x2c, 0x1d,
		0x3a, 0xa4, 0x0e, 0x10, 0x37, 0xcf, 0x83, 0x2c, 0x9d, 0x12, 0x26, 0x36,
		0xa9, 0x91, 0xb4, 0x9b, 0x42, 0xf3, 0xff, 0xbe, 0x3b, 0xd2, 0xb2, 0x64,
		0xc5, 0xf2, 0xd0, 0x7a, 0x2f, 0x12, 0x45, 0xde, 0x7d, 0xdf, 0x77, 0x3f,
		0x48, 0xaa, 0xae, 0x0b, 0x2c, 0x85, 0x44, 0x88, 0x0a, 0xcc, 0x55, 0x81,
		0x95, 0x16, 0x0b, 0x61, 0xc5, 0x0a, 0xa3, 0xf5, 0x3a, 0x0c, 0xea, 0x5a,
		0x94, 0x20, 0x8c, 0xfd, 0x5e, 0x21, 0xa4, 0x10, 0x19, 0xab, 0x85, 0xbc,
		0x77, 0x2b, 0x01, 0x2d, 0xa0, 0xd6, 0x70, 0x31, 0x06, 0x93, 0xde, 0x62,
		0x56, 0x4c, 0xdd, 0x5a, 0xbc, 0x4a, 0x7e, 0x73, 0xf3, 0xaf, 0xc6, 0x20,
		0xc5, 0x1c, 0x6a, 0xb2, 0x0c, 0x34, 0xda, 0xa5, 0x96, 0x3c, 0x4d, 0x5f,
		0x0e, 0x15, 0x65, 0xb1, 0x07, 0x5e, 0x48, 0xbb, 0x1f, 0xfb, 0xb3, 0xb4,
		0xc7, 0x02, 0xbf, 0x7f, 0x37, 0x08, 0xfd, 0xfe, 0xdd, 0x51, 0xe0, 0xcb,
		0x41, 0xd9, 0x77, 0xe2, 0x48, 0xdd, 0xcb, 0x03, 0xc2, 0xef, 0xc4, 0xd1,
		0xca, 0xcb, 0xb9, 0xca, 0xec, 0xaf, 0xbf, 0xec, 0xc7, 0xff, 0xe4, 0x17,
		0x8f, 0x27, 0x18, 0x0a, 0xe0, 0x93, 0x5f, 0x3c, 0x8a, 0x60, 0xa6, 0xd4,
		0x7c, 0x3f, 0xfa, 0x47, 0x5a, 0xf9, 0x21, 0xe8, 0xe6, 0x1d, 0xf6, 0x77,
		0x84, 0xb1, 0x99, 0x76, 0x1c, 0x2d, 0xc5, 0x04, 0xbf, 0xd5, 0xb5, 0x59,
		0xce, 0xbc, 0x8c, 0xf5, 0xfa, 0xcf, 0xe9, 0xcd, 0x64, 0x9a, 0x67, 0xf2,
		0x77, 0xe7, 0xa0, 0x63, 0x93, 0xa4, 0x7e, 0xb8, 0x4f, 0x41, 0x57, 0x40,
		0x87, 0x36, 0xac, 0xb2, 0xfc, 0x29, 0xbb, 0x47, 0xa8, 0xeb, 0x74, 0x92,
		0x2d, 0xd0, 0x3d, 0x78, 0x41, 0x2c, 0x2a, 0xa5, 0x2d, 0xc4, 0x61, 0x10,
		0x91, 0x8f, 0xd2, 0x26, 0xa2, 0x51, 0xb9, 0xb0, 0xfc, 0x12, 0x8a, 0x9f,
		0xf7, 0xc2, 0x3e, 0x2c, 0x67, 0x69, 0xae, 0x16, 0xa3, 0x19, 0xca, 0xd9,
		0xa3, 0x7a, 0x90, 0x46, 0xc9, 0xd1, 0x02, 0xef, 0xb3, 0x47, 0x1e, 0x18,
		0x92, 0x26, 0x51, 0x47, 0x61, 0xc2, 0xd1, 0xe9, 0x4c, 0x12, 0x0b, 0x4b,
		0x37, 0xac, 0x3d, 0x74, 0x41, 0xec, 0x92, 0x72, 0x3c, 0x9b, 0x58, 0x80,
		0xf6, 0xfc, 0x32, 0xb7, 0xac, 0xdc, 0xc0, 0x06, 0x28, 0x9d, 0xfa, 0x77,
		0x48, 0xea, 0xca, 0xa5, 0xcc, 0x7d, 0x3e, 0x06, 0x00, 0x62, 0x0d, 0x42,
		0xb9, 0x92, 0xa0, 0x4e, 0xe0, 0x6c, 0x98, 0x88, 0x18, 0x36, 0xa9, 0x79,
		0x3d, 0x68, 0x54, 0x9b, 0x8b, 0xad, 0x08, 0x22, 0xdd, 0xe8, 0x88, 0x75,
		0xb2, 0x3e, 0xac, 0x65, 0xa7, 0x38, 0xfd, 0x30, 0xfe, 0x1f, 0x55, 0xad,
		0x82, 0x18, 0x0f, 0x00, 0x26, 0xb0, 0xe9, 0x8c, 0xca, 0x6a, 0x38, 0xeb,
		0xd9, 0x25, 0xe0, 0x2a, 0xec, 0x93, 0x4d, 0x7d, 0x86, 0xa9, 0x71, 0x6d,
		0x67, 0xd5, 0xd3, 0x39, 0xfc, 0x75, 0xde, 0x76, 0x38, 0x2b, 0x8f, 0xff,
		0xa3, 0xb5, 0x00, 0xe7, 0x06, 0xc1, 0x7b, 0xb3, 0x51, 0x13, 0xf5, 0xd7,
		0xeb, 0x8f, 0xb7, 0x1f, 0x2e, 0xaf, 0x7a, 0x0e, 0xd4, 0x57, 0x9c, 0xd1,
		0x38, 0xba, 0x7a, 0xae, 0x30, 0xb7, 0x58, 0xc0, 0x69, 0x7d, 0x1a, 0x25,
		0xdc, 0xa2, 0x61, 0x30, 0x1a, 0xc1, 0xa5, 0xc6, 0xcc, 0x52, 0xdb, 0x3c,
		0x20, 0xa8, 0xd9, 0x23, 0x59, 0x30, 0xb2, 0xb0, 0x50, 0x28, 0x34, 0xf2,
		0xd4, 0x02, 0x3e, 0xd3, 0xae, 0x4c, 0x9d, 0xda, 0x33, 0x0e, 0x6d, 0xdc,
		0xaa, 0xf2, 0xdf, 0xfd, 0xfc, 0xd5, 0x6b, 0xc6, 0x0e, 0x56, 0x1c, 0x10,
		0x5b, 0x78, 0x9a, 0x6b, 0xa5, 0x2a, 0x50, 0x2b, 0x4a, 0xfc, 0x13, 0x7e,
		0x1f, 0xad, 0xb2, 0xf9, 0x12, 0xa1, 0xca, 0x04, 0x89, 0x23, 0x68, 0x59,
		0xe0, 0x33, 0x9b, 0xbf, 0x0d, 0x83, 0xd2, 0x67, 0x89, 0x5d, 0xb8, 0xb5,
		0x40, 0x48, 0x76, 0x20, 0xa3, 0x60, 0x95, 0x39, 0x5f, 0xf0, 0x77, 0x15,
		0x4d, 0xb8, 0xe4, 0xd1, 0x83, 0xc0, 0x5e, 0x64, 0xb0, 0x3d, 0x38, 0x86,
		0x8f, 0x88, 0x9d, 0x3c, 0x8e, 0x3b, 0x79, 0xbc, 0x6d, 0xf3, 0xd8, 0xb8,
		0x10, 0xc4, 0x41, 0x97, 0xcb, 0x9b, 0x2f, 0x5f, 0x3e, 0x78, 0x0f, 0x4e,
		0x9f, 0x0b, 0x88, 0xd6, 0xdf, 0xfa, 0xa9, 0x06, 0x85, 0x36, 0x77, 0x7a,
		0xc5, 0x25, 0x29, 0xe3, 0xe8, 0x4e, 0x62, 0x53, 0x10, 0xda, 0xdf, 0x8b,
		0x8c, 0x4b, 0xe2, 0xce, 0xad, 0x60, 0xdb, 0x18, 0xdd, 0xd8, 0x0e, 0x34,
		0x47, 0x2f, 0x2e, 0x87, 0xc1, 0xc5, 0x0d, 0xf6, 0xb4, 0xc8, 0xf4, 0xeb,
		0xed, 0xe7, 0xc9, 0x1f, 0x3b, 0xa1, 0x0d, 0x88, 0x3a, 0xa1, 0xde, 0x3f,
		0x31, 0x44, 0xd6, 0x69, 0x1b, 0x50, 0x7a, 0x93, 0xfe, 0xe8, 0xbc, 0xc5,
		0x54, 0x4f, 0x28, 0xb9, 0xf4, 0x31, 0x91, 0x25, 0xe7, 0x1b, 0x83, 0xd8,
		0x8b, 0x4f, 0x92, 0x36, 0x6d, 0x8e, 0x94, 0x2b, 0x38, 0xee, 0xd9, 0x34,
		0x7a, 0x3b, 0x35, 0xe7, 0x76, 0xcc, 0xd5, 0x5c, 0xc9, 0x34, 0xdc, 0x9f,
		0x8f, 0x43, 0xbb, 0xe5, 0x50, 0x9d, 0x5f, 0xed, 0x14, 0xed, 0xfa, 0x66,
		0xf2, 0x53, 0xb9, 0x70, 0xd2, 0x7e, 0x2c, 0x07, 0x1c, 0xa0, 0xa3, 0x32,
		0xdf, 0x84, 0xcd, 0x1f, 0x5c, 0x2b, 0xbb, 0xef, 0xe6, 0xe0, 0x2e, 0x05,
		0xce, 0x0b, 0x7f, 0x72, 0x07, 0x6e, 0x9a, 0x14, 0x93, 0x91, 0x24, 0xd4,
		0xed, 0x64, 0x9e, 0x71, 0x1e, 0xeb, 0xed, 0x34, 0xfc, 0x03, 0xf4, 0x5f,
		0x27, 0x6d, 0x09, 0xd1, 0xc9, 0xdf, 0x74, 0x97, 0x5d, 0x38, 0xab, 0xe6,
		0x32, 0xdd, 0xfe, 0xf2, 0x35, 0xd7, 0x99, 0x5f, 0xf5, 0x5b, 0xf3, 0xf5,
		0x2a, 0xad, 0x6b, 0xc7, 0xb9, 0xc3, 0xc0, 0xce, 0x16, 0x17, 0xd5, 0x9c,
		0x0f, 0x85, 0x17, 0xbf, 0x8e, 0xad, 0xd9, 0xf6, 0x7a, 0xeb, 0x12, 0x36,
		0xb7, 0xf7, 0x59, 0xf4, 0xd3, 0x5c, 0xee, 0x52, 0x7e, 0x49, 0xd3, 0x1b,
		0x77, 0x86, 0x6d, 0xb7, 0xf3, 0xa6, 0x7b, 0xf3, 0xc6, 0x1f, 0x6c, 0x9d,
		0x6d, 0xcb, 0x57, 0x70, 0x63, 0xff, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff,
		0xf1, 0xa3, 0x14, 0x74, 0x15, 0x0b, 0x00, 0x00,
	}))

	if err != nil {
		panic("Decompression failed: " + err.Error())
	}

	var b bytes.Buffer
	io.Copy(&b, gz)
	gz.Close()

	return b.Bytes()
}