package chart

import "github.com/t-mw/go-chart/drawing"

var viridisColors = [256]drawing.Color{
	drawing.Color{R: 0x44, G: 0x1, B: 0x54, A: 0xff},
	drawing.Color{R: 0x44, G: 0x2, B: 0x55, A: 0xff},
	drawing.Color{R: 0x45, G: 0x3, B: 0x57, A: 0xff},
	drawing.Color{R: 0x45, G: 0x5, B: 0x58, A: 0xff},
	drawing.Color{R: 0x45, G: 0x6, B: 0x5a, A: 0xff},
	drawing.Color{R: 0x46, G: 0x8, B: 0x5b, A: 0xff},
	drawing.Color{R: 0x46, G: 0x9, B: 0x5d, A: 0xff},
	drawing.Color{R: 0x46, G: 0xb, B: 0x5e, A: 0xff},
	drawing.Color{R: 0x46, G: 0xc, B: 0x60, A: 0xff},
	drawing.Color{R: 0x47, G: 0xe, B: 0x61, A: 0xff},
	drawing.Color{R: 0x47, G: 0xf, B: 0x62, A: 0xff},
	drawing.Color{R: 0x47, G: 0x11, B: 0x64, A: 0xff},
	drawing.Color{R: 0x47, G: 0x12, B: 0x65, A: 0xff},
	drawing.Color{R: 0x47, G: 0x14, B: 0x66, A: 0xff},
	drawing.Color{R: 0x48, G: 0x15, B: 0x68, A: 0xff},
	drawing.Color{R: 0x48, G: 0x16, B: 0x69, A: 0xff},
	drawing.Color{R: 0x48, G: 0x18, B: 0x6a, A: 0xff},
	drawing.Color{R: 0x48, G: 0x19, B: 0x6c, A: 0xff},
	drawing.Color{R: 0x48, G: 0x1a, B: 0x6d, A: 0xff},
	drawing.Color{R: 0x48, G: 0x1c, B: 0x6e, A: 0xff},
	drawing.Color{R: 0x48, G: 0x1d, B: 0x6f, A: 0xff},
	drawing.Color{R: 0x48, G: 0x1e, B: 0x70, A: 0xff},
	drawing.Color{R: 0x48, G: 0x20, B: 0x71, A: 0xff},
	drawing.Color{R: 0x48, G: 0x21, B: 0x73, A: 0xff},
	drawing.Color{R: 0x48, G: 0x22, B: 0x74, A: 0xff},
	drawing.Color{R: 0x48, G: 0x24, B: 0x75, A: 0xff},
	drawing.Color{R: 0x48, G: 0x25, B: 0x76, A: 0xff},
	drawing.Color{R: 0x48, G: 0x26, B: 0x77, A: 0xff},
	drawing.Color{R: 0x48, G: 0x27, B: 0x78, A: 0xff},
	drawing.Color{R: 0x47, G: 0x29, B: 0x79, A: 0xff},
	drawing.Color{R: 0x47, G: 0x2a, B: 0x79, A: 0xff},
	drawing.Color{R: 0x47, G: 0x2b, B: 0x7a, A: 0xff},
	drawing.Color{R: 0x47, G: 0x2c, B: 0x7b, A: 0xff},
	drawing.Color{R: 0x47, G: 0x2e, B: 0x7c, A: 0xff},
	drawing.Color{R: 0x46, G: 0x2f, B: 0x7d, A: 0xff},
	drawing.Color{R: 0x46, G: 0x30, B: 0x7e, A: 0xff},
	drawing.Color{R: 0x46, G: 0x31, B: 0x7e, A: 0xff},
	drawing.Color{R: 0x46, G: 0x33, B: 0x7f, A: 0xff},
	drawing.Color{R: 0x45, G: 0x34, B: 0x80, A: 0xff},
	drawing.Color{R: 0x45, G: 0x35, B: 0x81, A: 0xff},
	drawing.Color{R: 0x45, G: 0x36, B: 0x81, A: 0xff},
	drawing.Color{R: 0x44, G: 0x38, B: 0x82, A: 0xff},
	drawing.Color{R: 0x44, G: 0x39, B: 0x83, A: 0xff},
	drawing.Color{R: 0x44, G: 0x3a, B: 0x83, A: 0xff},
	drawing.Color{R: 0x43, G: 0x3b, B: 0x84, A: 0xff},
	drawing.Color{R: 0x43, G: 0x3c, B: 0x84, A: 0xff},
	drawing.Color{R: 0x43, G: 0x3e, B: 0x85, A: 0xff},
	drawing.Color{R: 0x42, G: 0x3f, B: 0x85, A: 0xff},
	drawing.Color{R: 0x42, G: 0x40, B: 0x86, A: 0xff},
	drawing.Color{R: 0x41, G: 0x41, B: 0x86, A: 0xff},
	drawing.Color{R: 0x41, G: 0x42, B: 0x87, A: 0xff},
	drawing.Color{R: 0x41, G: 0x43, B: 0x87, A: 0xff},
	drawing.Color{R: 0x40, G: 0x45, B: 0x88, A: 0xff},
	drawing.Color{R: 0x40, G: 0x46, B: 0x88, A: 0xff},
	drawing.Color{R: 0x3f, G: 0x47, B: 0x88, A: 0xff},
	drawing.Color{R: 0x3f, G: 0x48, B: 0x89, A: 0xff},
	drawing.Color{R: 0x3e, G: 0x49, B: 0x89, A: 0xff},
	drawing.Color{R: 0x3e, G: 0x4a, B: 0x89, A: 0xff},
	drawing.Color{R: 0x3d, G: 0x4b, B: 0x8a, A: 0xff},
	drawing.Color{R: 0x3d, G: 0x4d, B: 0x8a, A: 0xff},
	drawing.Color{R: 0x3c, G: 0x4e, B: 0x8a, A: 0xff},
	drawing.Color{R: 0x3c, G: 0x4f, B: 0x8a, A: 0xff},
	drawing.Color{R: 0x3b, G: 0x50, B: 0x8b, A: 0xff},
	drawing.Color{R: 0x3b, G: 0x51, B: 0x8b, A: 0xff},
	drawing.Color{R: 0x3a, G: 0x52, B: 0x8b, A: 0xff},
	drawing.Color{R: 0x3a, G: 0x53, B: 0x8b, A: 0xff},
	drawing.Color{R: 0x39, G: 0x54, B: 0x8c, A: 0xff},
	drawing.Color{R: 0x39, G: 0x55, B: 0x8c, A: 0xff},
	drawing.Color{R: 0x38, G: 0x56, B: 0x8c, A: 0xff},
	drawing.Color{R: 0x38, G: 0x57, B: 0x8c, A: 0xff},
	drawing.Color{R: 0x37, G: 0x58, B: 0x8c, A: 0xff},
	drawing.Color{R: 0x37, G: 0x59, B: 0x8c, A: 0xff},
	drawing.Color{R: 0x36, G: 0x5b, B: 0x8d, A: 0xff},
	drawing.Color{R: 0x36, G: 0x5c, B: 0x8d, A: 0xff},
	drawing.Color{R: 0x35, G: 0x5d, B: 0x8d, A: 0xff},
	drawing.Color{R: 0x35, G: 0x5e, B: 0x8d, A: 0xff},
	drawing.Color{R: 0x34, G: 0x5f, B: 0x8d, A: 0xff},
	drawing.Color{R: 0x34, G: 0x60, B: 0x8d, A: 0xff},
	drawing.Color{R: 0x33, G: 0x61, B: 0x8d, A: 0xff},
	drawing.Color{R: 0x33, G: 0x62, B: 0x8d, A: 0xff},
	drawing.Color{R: 0x33, G: 0x63, B: 0x8d, A: 0xff},
	drawing.Color{R: 0x32, G: 0x64, B: 0x8e, A: 0xff},
	drawing.Color{R: 0x32, G: 0x65, B: 0x8e, A: 0xff},
	drawing.Color{R: 0x31, G: 0x66, B: 0x8e, A: 0xff},
	drawing.Color{R: 0x31, G: 0x67, B: 0x8e, A: 0xff},
	drawing.Color{R: 0x30, G: 0x68, B: 0x8e, A: 0xff},
	drawing.Color{R: 0x30, G: 0x69, B: 0x8e, A: 0xff},
	drawing.Color{R: 0x2f, G: 0x6a, B: 0x8e, A: 0xff},
	drawing.Color{R: 0x2f, G: 0x6b, B: 0x8e, A: 0xff},
	drawing.Color{R: 0x2f, G: 0x6c, B: 0x8e, A: 0xff},
	drawing.Color{R: 0x2e, G: 0x6d, B: 0x8e, A: 0xff},
	drawing.Color{R: 0x2e, G: 0x6e, B: 0x8e, A: 0xff},
	drawing.Color{R: 0x2d, G: 0x6f, B: 0x8e, A: 0xff},
	drawing.Color{R: 0x2d, G: 0x70, B: 0x8e, A: 0xff},
	drawing.Color{R: 0x2d, G: 0x70, B: 0x8e, A: 0xff},
	drawing.Color{R: 0x2c, G: 0x71, B: 0x8e, A: 0xff},
	drawing.Color{R: 0x2c, G: 0x72, B: 0x8e, A: 0xff},
	drawing.Color{R: 0x2b, G: 0x73, B: 0x8e, A: 0xff},
	drawing.Color{R: 0x2b, G: 0x74, B: 0x8e, A: 0xff},
	drawing.Color{R: 0x2b, G: 0x75, B: 0x8e, A: 0xff},
	drawing.Color{R: 0x2a, G: 0x76, B: 0x8e, A: 0xff},
	drawing.Color{R: 0x2a, G: 0x77, B: 0x8e, A: 0xff},
	drawing.Color{R: 0x29, G: 0x78, B: 0x8e, A: 0xff},
	drawing.Color{R: 0x29, G: 0x79, B: 0x8e, A: 0xff},
	drawing.Color{R: 0x29, G: 0x7a, B: 0x8e, A: 0xff},
	drawing.Color{R: 0x28, G: 0x7b, B: 0x8e, A: 0xff},
	drawing.Color{R: 0x28, G: 0x7c, B: 0x8e, A: 0xff},
	drawing.Color{R: 0x28, G: 0x7d, B: 0x8e, A: 0xff},
	drawing.Color{R: 0x27, G: 0x7e, B: 0x8e, A: 0xff},
	drawing.Color{R: 0x27, G: 0x7f, B: 0x8e, A: 0xff},
	drawing.Color{R: 0x26, G: 0x80, B: 0x8e, A: 0xff},
	drawing.Color{R: 0x26, G: 0x81, B: 0x8e, A: 0xff},
	drawing.Color{R: 0x26, G: 0x82, B: 0x8e, A: 0xff},
	drawing.Color{R: 0x25, G: 0x83, B: 0x8e, A: 0xff},
	drawing.Color{R: 0x25, G: 0x83, B: 0x8e, A: 0xff},
	drawing.Color{R: 0x25, G: 0x84, B: 0x8e, A: 0xff},
	drawing.Color{R: 0x24, G: 0x85, B: 0x8e, A: 0xff},
	drawing.Color{R: 0x24, G: 0x86, B: 0x8e, A: 0xff},
	drawing.Color{R: 0x23, G: 0x87, B: 0x8e, A: 0xff},
	drawing.Color{R: 0x23, G: 0x88, B: 0x8e, A: 0xff},
	drawing.Color{R: 0x23, G: 0x89, B: 0x8e, A: 0xff},
	drawing.Color{R: 0x22, G: 0x8a, B: 0x8d, A: 0xff},
	drawing.Color{R: 0x22, G: 0x8b, B: 0x8d, A: 0xff},
	drawing.Color{R: 0x22, G: 0x8c, B: 0x8d, A: 0xff},
	drawing.Color{R: 0x21, G: 0x8d, B: 0x8d, A: 0xff},
	drawing.Color{R: 0x21, G: 0x8e, B: 0x8d, A: 0xff},
	drawing.Color{R: 0x21, G: 0x8f, B: 0x8d, A: 0xff},
	drawing.Color{R: 0x20, G: 0x90, B: 0x8d, A: 0xff},
	drawing.Color{R: 0x20, G: 0x91, B: 0x8c, A: 0xff},
	drawing.Color{R: 0x20, G: 0x92, B: 0x8c, A: 0xff},
	drawing.Color{R: 0x20, G: 0x93, B: 0x8c, A: 0xff},
	drawing.Color{R: 0x1f, G: 0x93, B: 0x8c, A: 0xff},
	drawing.Color{R: 0x1f, G: 0x94, B: 0x8c, A: 0xff},
	drawing.Color{R: 0x1f, G: 0x95, B: 0x8b, A: 0xff},
	drawing.Color{R: 0x1f, G: 0x96, B: 0x8b, A: 0xff},
	drawing.Color{R: 0x1f, G: 0x97, B: 0x8b, A: 0xff},
	drawing.Color{R: 0x1e, G: 0x98, B: 0x8b, A: 0xff},
	drawing.Color{R: 0x1e, G: 0x99, B: 0x8a, A: 0xff},
	drawing.Color{R: 0x1e, G: 0x9a, B: 0x8a, A: 0xff},
	drawing.Color{R: 0x1e, G: 0x9b, B: 0x8a, A: 0xff},
	drawing.Color{R: 0x1e, G: 0x9c, B: 0x89, A: 0xff},
	drawing.Color{R: 0x1e, G: 0x9d, B: 0x89, A: 0xff},
	drawing.Color{R: 0x1e, G: 0x9e, B: 0x89, A: 0xff},
	drawing.Color{R: 0x1e, G: 0x9f, B: 0x88, A: 0xff},
	drawing.Color{R: 0x1e, G: 0xa0, B: 0x88, A: 0xff},
	drawing.Color{R: 0x1f, G: 0xa1, B: 0x88, A: 0xff},
	drawing.Color{R: 0x1f, G: 0xa2, B: 0x87, A: 0xff},
	drawing.Color{R: 0x1f, G: 0xa3, B: 0x87, A: 0xff},
	drawing.Color{R: 0x1f, G: 0xa3, B: 0x86, A: 0xff},
	drawing.Color{R: 0x20, G: 0xa4, B: 0x86, A: 0xff},
	drawing.Color{R: 0x20, G: 0xa5, B: 0x86, A: 0xff},
	drawing.Color{R: 0x21, G: 0xa6, B: 0x85, A: 0xff},
	drawing.Color{R: 0x21, G: 0xa7, B: 0x85, A: 0xff},
	drawing.Color{R: 0x22, G: 0xa8, B: 0x84, A: 0xff},
	drawing.Color{R: 0x23, G: 0xa9, B: 0x83, A: 0xff},
	drawing.Color{R: 0x23, G: 0xaa, B: 0x83, A: 0xff},
	drawing.Color{R: 0x24, G: 0xab, B: 0x82, A: 0xff},
	drawing.Color{R: 0x25, G: 0xac, B: 0x82, A: 0xff},
	drawing.Color{R: 0x26, G: 0xad, B: 0x81, A: 0xff},
	drawing.Color{R: 0x27, G: 0xae, B: 0x81, A: 0xff},
	drawing.Color{R: 0x28, G: 0xaf, B: 0x80, A: 0xff},
	drawing.Color{R: 0x29, G: 0xaf, B: 0x7f, A: 0xff},
	drawing.Color{R: 0x2a, G: 0xb0, B: 0x7f, A: 0xff},
	drawing.Color{R: 0x2b, G: 0xb1, B: 0x7e, A: 0xff},
	drawing.Color{R: 0x2c, G: 0xb2, B: 0x7d, A: 0xff},
	drawing.Color{R: 0x2e, G: 0xb3, B: 0x7c, A: 0xff},
	drawing.Color{R: 0x2f, G: 0xb4, B: 0x7c, A: 0xff},
	drawing.Color{R: 0x30, G: 0xb5, B: 0x7b, A: 0xff},
	drawing.Color{R: 0x32, G: 0xb6, B: 0x7a, A: 0xff},
	drawing.Color{R: 0x33, G: 0xb7, B: 0x79, A: 0xff},
	drawing.Color{R: 0x35, G: 0xb7, B: 0x79, A: 0xff},
	drawing.Color{R: 0x36, G: 0xb8, B: 0x78, A: 0xff},
	drawing.Color{R: 0x38, G: 0xb9, B: 0x77, A: 0xff},
	drawing.Color{R: 0x39, G: 0xba, B: 0x76, A: 0xff},
	drawing.Color{R: 0x3b, G: 0xbb, B: 0x75, A: 0xff},
	drawing.Color{R: 0x3d, G: 0xbc, B: 0x74, A: 0xff},
	drawing.Color{R: 0x3e, G: 0xbd, B: 0x73, A: 0xff},
	drawing.Color{R: 0x40, G: 0xbe, B: 0x72, A: 0xff},
	drawing.Color{R: 0x42, G: 0xbe, B: 0x71, A: 0xff},
	drawing.Color{R: 0x44, G: 0xbf, B: 0x70, A: 0xff},
	drawing.Color{R: 0x46, G: 0xc0, B: 0x6f, A: 0xff},
	drawing.Color{R: 0x48, G: 0xc1, B: 0x6e, A: 0xff},
	drawing.Color{R: 0x49, G: 0xc2, B: 0x6d, A: 0xff},
	drawing.Color{R: 0x4b, G: 0xc2, B: 0x6c, A: 0xff},
	drawing.Color{R: 0x4d, G: 0xc3, B: 0x6b, A: 0xff},
	drawing.Color{R: 0x4f, G: 0xc4, B: 0x6a, A: 0xff},
	drawing.Color{R: 0x51, G: 0xc5, B: 0x69, A: 0xff},
	drawing.Color{R: 0x53, G: 0xc6, B: 0x68, A: 0xff},
	drawing.Color{R: 0x55, G: 0xc6, B: 0x66, A: 0xff},
	drawing.Color{R: 0x58, G: 0xc7, B: 0x65, A: 0xff},
	drawing.Color{R: 0x5a, G: 0xc8, B: 0x64, A: 0xff},
	drawing.Color{R: 0x5c, G: 0xc9, B: 0x63, A: 0xff},
	drawing.Color{R: 0x5e, G: 0xc9, B: 0x62, A: 0xff},
	drawing.Color{R: 0x60, G: 0xca, B: 0x60, A: 0xff},
	drawing.Color{R: 0x62, G: 0xcb, B: 0x5f, A: 0xff},
	drawing.Color{R: 0x65, G: 0xcc, B: 0x5e, A: 0xff},
	drawing.Color{R: 0x67, G: 0xcc, B: 0x5c, A: 0xff},
	drawing.Color{R: 0x69, G: 0xcd, B: 0x5b, A: 0xff},
	drawing.Color{R: 0x6c, G: 0xce, B: 0x5a, A: 0xff},
	drawing.Color{R: 0x6e, G: 0xce, B: 0x58, A: 0xff},
	drawing.Color{R: 0x70, G: 0xcf, B: 0x57, A: 0xff},
	drawing.Color{R: 0x73, G: 0xd0, B: 0x55, A: 0xff},
	drawing.Color{R: 0x75, G: 0xd0, B: 0x54, A: 0xff},
	drawing.Color{R: 0x77, G: 0xd1, B: 0x52, A: 0xff},
	drawing.Color{R: 0x7a, G: 0xd2, B: 0x51, A: 0xff},
	drawing.Color{R: 0x7c, G: 0xd2, B: 0x4f, A: 0xff},
	drawing.Color{R: 0x7f, G: 0xd3, B: 0x4e, A: 0xff},
	drawing.Color{R: 0x81, G: 0xd4, B: 0x4c, A: 0xff},
	drawing.Color{R: 0x84, G: 0xd4, B: 0x4b, A: 0xff},
	drawing.Color{R: 0x86, G: 0xd5, B: 0x49, A: 0xff},
	drawing.Color{R: 0x89, G: 0xd5, B: 0x48, A: 0xff},
	drawing.Color{R: 0x8b, G: 0xd6, B: 0x46, A: 0xff},
	drawing.Color{R: 0x8e, G: 0xd7, B: 0x44, A: 0xff},
	drawing.Color{R: 0x90, G: 0xd7, B: 0x43, A: 0xff},
	drawing.Color{R: 0x93, G: 0xd8, B: 0x41, A: 0xff},
	drawing.Color{R: 0x95, G: 0xd8, B: 0x3f, A: 0xff},
	drawing.Color{R: 0x98, G: 0xd9, B: 0x3e, A: 0xff},
	drawing.Color{R: 0x9b, G: 0xd9, B: 0x3c, A: 0xff},
	drawing.Color{R: 0x9d, G: 0xda, B: 0x3a, A: 0xff},
	drawing.Color{R: 0xa0, G: 0xda, B: 0x39, A: 0xff},
	drawing.Color{R: 0xa3, G: 0xdb, B: 0x37, A: 0xff},
	drawing.Color{R: 0xa5, G: 0xdb, B: 0x35, A: 0xff},
	drawing.Color{R: 0xa8, G: 0xdc, B: 0x33, A: 0xff},
	drawing.Color{R: 0xab, G: 0xdc, B: 0x32, A: 0xff},
	drawing.Color{R: 0xad, G: 0xdd, B: 0x30, A: 0xff},
	drawing.Color{R: 0xb0, G: 0xdd, B: 0x2e, A: 0xff},
	drawing.Color{R: 0xb3, G: 0xdd, B: 0x2d, A: 0xff},
	drawing.Color{R: 0xb5, G: 0xde, B: 0x2b, A: 0xff},
	drawing.Color{R: 0xb8, G: 0xde, B: 0x29, A: 0xff},
	drawing.Color{R: 0xbb, G: 0xdf, B: 0x27, A: 0xff},
	drawing.Color{R: 0xbd, G: 0xdf, B: 0x26, A: 0xff},
	drawing.Color{R: 0xc0, G: 0xdf, B: 0x24, A: 0xff},
	drawing.Color{R: 0xc3, G: 0xe0, B: 0x23, A: 0xff},
	drawing.Color{R: 0xc5, G: 0xe0, B: 0x21, A: 0xff},
	drawing.Color{R: 0xc8, G: 0xe1, B: 0x20, A: 0xff},
	drawing.Color{R: 0xcb, G: 0xe1, B: 0x1e, A: 0xff},
	drawing.Color{R: 0xcd, G: 0xe1, B: 0x1d, A: 0xff},
	drawing.Color{R: 0xd0, G: 0xe2, B: 0x1c, A: 0xff},
	drawing.Color{R: 0xd3, G: 0xe2, B: 0x1b, A: 0xff},
	drawing.Color{R: 0xd5, G: 0xe2, B: 0x1a, A: 0xff},
	drawing.Color{R: 0xd8, G: 0xe3, B: 0x19, A: 0xff},
	drawing.Color{R: 0xdb, G: 0xe3, B: 0x18, A: 0xff},
	drawing.Color{R: 0xdd, G: 0xe3, B: 0x18, A: 0xff},
	drawing.Color{R: 0xe0, G: 0xe4, B: 0x18, A: 0xff},
	drawing.Color{R: 0xe2, G: 0xe4, B: 0x18, A: 0xff},
	drawing.Color{R: 0xe5, G: 0xe4, B: 0x18, A: 0xff},
	drawing.Color{R: 0xe8, G: 0xe5, B: 0x19, A: 0xff},
	drawing.Color{R: 0xea, G: 0xe5, B: 0x19, A: 0xff},
	drawing.Color{R: 0xed, G: 0xe5, B: 0x1a, A: 0xff},
	drawing.Color{R: 0xef, G: 0xe6, B: 0x1b, A: 0xff},
	drawing.Color{R: 0xf2, G: 0xe6, B: 0x1c, A: 0xff},
	drawing.Color{R: 0xf4, G: 0xe6, B: 0x1e, A: 0xff},
	drawing.Color{R: 0xf7, G: 0xe6, B: 0x1f, A: 0xff},
	drawing.Color{R: 0xf9, G: 0xe7, B: 0x21, A: 0xff},
	drawing.Color{R: 0xfb, G: 0xe7, B: 0x23, A: 0xff},
	drawing.Color{R: 0xfe, G: 0xe7, B: 0x24, A: 0xff},
}

// Viridis creates a color map provider.
func Viridis(v, vmin, vmax float64) drawing.Color {
	normalized := (v - vmin) / (vmax - vmin)
	index := uint8(normalized * 255)
	return viridisColors[index]
}
