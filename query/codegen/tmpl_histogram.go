package codegen

import (
	"bytes"
	"compress/gzip"
	"io"
)

// histogram_tmpl returns raw, uncompressed file data.
func histogram_tmpl() []byte {
	gz, err := gzip.NewReader(bytes.NewBuffer([]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x00, 0xff, 0x74, 0x54,
		0xc1, 0x92, 0x9b, 0x30, 0x0c, 0xbd, 0xf3, 0x15, 0x9a, 0xed, 0xa1, 0x64,
		0x16, 0xb2, 0xc9, 0x07, 0x70, 0xea, 0x69, 0x0f, 0xfd, 0x82, 0x9d, 0x4c,
		0xc6, 0x80, 0x08, 0xee, 0x82, 0x3d, 0xb5, 0xcd, 0x6e, 0x5a, 0x26, 0xfd,
		0xf6, 0xca, 0xc6, 0x60, 0x70, 0x36, 0x97, 0xa0, 0x58, 0xd2, 0x93, 0xf4,
		0xf4, 0xec, 0x71, 0xac, 0xb1, 0xe1, 0x02, 0xe1, 0xa9, 0xe5, 0xda, 0xc8,
		0x8b, 0x62, 0xfd, 0xd3, 0xed, 0x96, 0xe4, 0x39, 0xfc, 0x68, 0xb1, 0x7a,
		0xd7, 0xc0, 0x1b, 0x60, 0xf0, 0xc1, 0xba, 0x01, 0x81, 0x6b, 0x32, 0x97,
		0xb0, 0x7d, 0xd2, 0x0c, 0xa2, 0x32, 0x5c, 0x0a, 0xd0, 0xef, 0x7f, 0xce,
		0x5c, 0x9f, 0x17, 0x57, 0xba, 0x58, 0xbb, 0x04, 0x00, 0x14, 0x9a, 0x41,
		0x09, 0x08, 0xa7, 0xf0, 0xaf, 0x00, 0xc1, 0x3b, 0x60, 0xa2, 0x5e, 0xe1,
		0x9d, 0x03, 0xc0, 0xf9, 0x0c, 0x45, 0x01, 0x46, 0x0d, 0xb8, 0x4b, 0x50,
		0xd4, 0x89, 0xeb, 0x47, 0x21, 0x33, 0x68, 0x5b, 0x10, 0xf8, 0x19, 0xd2,
		0x40, 0x96, 0xbf, 0xb0, 0x32, 0x51, 0x37, 0x01, 0x89, 0x82, 0x53, 0xdb,
		0x85, 0x6f, 0x62, 0xdc, 0x54, 0x29, 0x6c, 0x89, 0x0c, 0x7a, 0x2e, 0x8a,
		0x03, 0x7d, 0xd8, 0xd5, 0x7e, 0x2a, 0x39, 0x08, 0x63, 0x8d, 0x4f, 0x5e,
		0x9b, 0xd6, 0x1a, 0x25, 0x17, 0xba, 0x18, 0x6f, 0xd9, 0xc4, 0x83, 0x35,
		0x6f, 0xa1, 0x2d, 0x29, 0x3e, 0x50, 0x99, 0x0d, 0x35, 0x94, 0x69, 0x5a,
		0x1f, 0x0c, 0x5c, 0x18, 0x49, 0x4e, 0x8d, 0x06, 0x64, 0xe3, 0xa0, 0xdc,
		0xdc, 0x55, 0x87, 0x4c, 0x69, 0x30, 0x2d, 0xfa, 0xc0, 0x87, 0x13, 0xd0,
		0x7a, 0x58, 0xc7, 0xff, 0xe2, 0x96, 0x55, 0xc3, 0xca, 0x0e, 0xf7, 0x5a,
		0x2a, 0x13, 0xce, 0xf7, 0x13, 0x92, 0x75, 0x87, 0x33, 0x1a, 0x2e, 0x5b,
		0xff, 0x65, 0x57, 0x28, 0x20, 0x4e, 0x79, 0x3b, 0x9e, 0xb2, 0xfb, 0xc3,
		0x6f, 0xf1, 0xc9, 0x29, 0xd9, 0x40, 0x3b, 0xa6, 0x08, 0xad, 0x67, 0xa6,
		0xdd, 0x57, 0xc8, 0xbb, 0xd4, 0x59, 0xfa, 0x37, 0x35, 0x75, 0x97, 0xbb,
		0xb3, 0x6d, 0x91, 0x9a, 0xee, 0xd2, 0x0b, 0x38, 0x58, 0x1a, 0xc4, 0x17,
		0xc0, 0x47, 0x70, 0x34, 0xaf, 0x6b, 0xba, 0xa5, 0x90, 0x2b, 0xdd, 0x8e,
		0x94, 0x6f, 0x27, 0xde, 0xc1, 0x4b, 0x0c, 0x17, 0x57, 0xf7, 0x40, 0x5f,
		0x54, 0x9f, 0x4b, 0x2c, 0xd5, 0x1b, 0xa9, 0x80, 0x93, 0x10, 0x22, 0xc4,
		0xfc, 0x08, 0xb5, 0xb4, 0x02, 0x5f, 0xe5, 0xda, 0xfd, 0xbe, 0xf1, 0x13,
		0x65, 0x1f, 0xc8, 0x73, 0xd7, 0xbd, 0xd7, 0x84, 0x93, 0xff, 0x22, 0xa1,
		0x57, 0xa1, 0xbd, 0x82, 0xfc, 0x3d, 0x73, 0x8a, 0x11, 0x80, 0x57, 0xca,
		0xe3, 0xe2, 0xf2, 0xf0, 0xce, 0x05, 0x8d, 0x70, 0x07, 0x11, 0x38, 0xf1,
		0x52, 0x8d, 0x29, 0xb7, 0xe3, 0xda, 0x8b, 0x47, 0xf3, 0x3c, 0x5e, 0x83,
		0xbf, 0x29, 0xb6, 0x3b, 0xca, 0x16, 0x35, 0x5e, 0xe7, 0x15, 0x37, 0x9d,
		0x94, 0x2a, 0x4d, 0x1d, 0x74, 0xbe, 0xe5, 0xfb, 0x25, 0xa2, 0x6f, 0x17,
		0xe7, 0x52, 0xd0, 0xa4, 0x0e, 0xda, 0x56, 0xea, 0x3c, 0x19, 0x1c, 0x76,
		0x59, 0xdc, 0x47, 0x7e, 0xdc, 0xaa, 0x77, 0xe2, 0xd3, 0x86, 0x9f, 0x36,
		0xb2, 0x5d, 0x9f, 0x3f, 0xc3, 0x71, 0xe1, 0xf2, 0x27, 0xaa, 0x0b, 0x11,
		0x2c, 0xe9, 0x45, 0x0b, 0x53, 0x7b, 0x42, 0x25, 0x8d, 0xa7, 0xf6, 0xf0,
		0x6a, 0xbe, 0x13, 0xd7, 0x5a, 0x0f, 0x3d, 0xd6, 0x34, 0x31, 0x33, 0x50,
		0x92, 0x27, 0x84, 0x6b, 0xd0, 0x2d, 0x53, 0x68, 0xd1, 0xec, 0xed, 0xd4,
		0xac, 0xa7, 0x1f, 0x7a, 0x28, 0x2a, 0xa2, 0x05, 0xdd, 0xdd, 0xa5, 0xe2,
		0xd3, 0x2b, 0xf1, 0x70, 0x1f, 0xbd, 0x6d, 0x23, 0x65, 0xf4, 0x74, 0xf8,
		0x15, 0xb0, 0x99, 0x7a, 0x4b, 0xb1, 0xd3, 0x0c, 0x9d, 0x40, 0x69, 0x35,
		0xd2, 0x69, 0xa4, 0x80, 0x72, 0x7e, 0x14, 0x97, 0x80, 0x59, 0x75, 0x2c,
		0x56, 0x1b, 0xe5, 0xae, 0x74, 0x16, 0xec, 0x67, 0x28, 0x67, 0xdb, 0xc5,
		0x4d, 0x1b, 0x9c, 0x7e, 0xfd, 0x52, 0xd9, 0x44, 0xd4, 0x38, 0xd2, 0x87,
		0x1e, 0xb1, 0xe4, 0x7f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x1f, 0x32, 0x13,
		0xf9, 0xff, 0x05, 0x00, 0x00,
	}))

	if err != nil {
		panic("Decompression failed: " + err.Error())
	}

	var b bytes.Buffer
	io.Copy(&b, gz)
	gz.Close()

	return b.Bytes()
}