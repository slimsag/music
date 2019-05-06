// Code generated by vfsgen; DO NOT EDIT.

// +build !dev

package website

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	pathpkg "path"
	"time"
)

// Assets contains project assets.
var Assets = func() http.FileSystem {
	mustUnmarshalTextTime := func(text string) time.Time {
		var t time.Time
		err := t.UnmarshalText([]byte(text))
		if err != nil {
			panic(err)
		}
		return t
	}

	fs := vfsgen۰FS{
		"/": &vfsgen۰DirInfo{
			name:    "/",
			modTime: mustUnmarshalTextTime("2019-04-22T00:26:06.494328833Z"),
		},
		"/pages": &vfsgen۰DirInfo{
			name:    "pages",
			modTime: mustUnmarshalTextTime("2019-04-22T07:44:10.078887074Z"),
		},
		"/pages/dizi": &vfsgen۰DirInfo{
			name:    "dizi",
			modTime: mustUnmarshalTextTime("2019-05-06T04:56:01.540191921Z"),
		},
		"/pages/dizi/index.md": &vfsgen۰CompressedFileInfo{
			name:             "index.md",
			modTime:          mustUnmarshalTextTime("2019-05-06T05:10:53.856078062Z"),
			uncompressedSize: 3855,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x57\xb1\x92\xdc\xc6\x11\xcd\xf1\x15\x5d\x60\xa0\xdb\xaa\x5d\xac\xc8\x93\x1c\x90\x3c\xba\x4e\x26\x69\xd1\x92\x18\xb8\x18\xd8\xe5\x72\xa9\x1a\x40\x03\x68\xde\x60\x06\x9c\x19\x2c\x0e\xa7\x62\xe0\xc4\x89\x03\x3b\x53\xa0\xcc\x91\xcb\xb6\x12\x67\x8e\xed\x2f\x51\x95\xa5\xdf\x70\x75\x0f\xb0\xb7\x47\x4a\x09\x8f\x3b\xdb\xd3\xd3\xdd\xef\xf5\xeb\xde\x7b\xf0\x94\x6f\x18\xce\x7e\xd1\xb1\xa5\x40\x10\x3d\xda\x70\x20\x1f\x08\x1a\x33\x46\xda\x64\xd9\x63\x6e\x3c\xf6\x04\x13\xd7\xb1\xbb\xc8\x3f\xfe\xd9\x87\x39\x74\xc4\x6d\x17\x2f\xf2\xf3\xfb\x1f\xe7\x10\x7c\x75\x91\x77\x31\x0e\xe1\xe1\x7e\x3f\x4d\x53\x31\xbb\x31\x8e\x25\x15\x95\xeb\xf7\xd4\x97\x54\xef\x9f\xfd\xfa\xb7\x73\xff\xfa\xf9\x6f\x88\x5f\xfc\x3c\x44\xf4\xf1\xe2\xa3\xf3\x1c\xd4\x6f\xe9\x7c\x4d\xfe\x22\xff\x30\x07\x34\xc6\x4d\x17\x39\x56\x15\x19\xf2\xae\xa7\x48\xfe\x11\xe0\x18\xdd\x60\x70\x7e\x04\x64\x2b\x3f\x0f\x91\xea\x5d\x4f\x35\xe3\x23\x68\x67\xef\x42\xe5\x06\x7a\x04\x03\x57\x71\xf4\xb4\x63\xbb\x5b\xfe\xbb\xf8\x6b\x46\x63\x42\xe5\x89\xec\x93\xc7\xfb\x94\xca\x93\x2c\xbb\x84\x5a\xf2\xe6\x00\x08\x3f\x95\xfb\x16\xc6\x30\xa2\x31\x33\xf4\x58\x13\xb8\x31\x82\x6b\xe0\x13\xec\x4b\xe7\xe0\xac\x1c\x23\x04\xd7\x13\xf4\xce\xcb\x3f\x35\x79\x0b\xce\x52\x00\xf4\x04\x68\x82\x5b\xae\x35\x30\x18\x0c\x91\xab\x2d\xf4\x14\xd1\x6c\x01\x6d\x0d\x97\x4d\x8b\x16\x5e\x63\x4d\x9b\x22\xcb\x5e\x44\x09\x65\xb4\xfc\x66\x24\xa8\x5c\x3f\xa0\xa7\x1a\xa2\x5b\xfd\x4e\x14\xa2\xfc\xad\x9c\xad\xc8\xc7\x14\x5f\x00\xb6\x10\x3b\x82\x06\xab\x08\xb1\xc3\x08\x1c\xa1\x43\xc9\xe9\x40\x7e\x86\xd8\xb1\x85\x01\x07\xf2\xd0\x53\x5f\x7a\xb4\xe2\xfc\x40\x9e\x6d\x0b\x08\x9d\x33\x04\x96\xd0\xab\x93\xde\x8d\xb1\x1b\x98\x2a\x82\xa9\xe3\xaa\x83\x0a\xc7\x20\x6f\x44\x89\x63\xf0\xae\x1e\x2b\x02\x84\x72\xbc\xb9\x91\xfb\xd6\x71\x20\xad\xc0\x24\x2f\x1b\xbe\x92\x6f\x3d\x55\x8a\x27\xf4\x7c\x4d\x35\x4c\x1c\x3b\x40\xb8\xc2\x1b\xe7\x0a\x78\x25\xcf\x1c\x03\x41\x0b\x25\x41\x14\x22\x91\x4d\xd9\x1a\x0a\xc1\xa5\x9c\xf4\x0a\x04\x37\xda\x7a\x0b\xce\x83\x71\x2e\xac\x66\x6c\x2b\x4f\x18\x08\x38\x16\x59\x26\x5e\x27\xe7\x6b\xc8\x05\xd2\x1c\x38\x06\x32\x8d\x42\x6b\xe1\x99\x6d\x0d\x87\x2e\x61\x6b\x38\x92\xc7\xc8\xce\x42\xe3\x52\xd6\x2b\xf6\xea\x40\x0e\x73\x2d\x6d\x0e\x67\x3f\xfc\xf3\x9b\xff\x7d\xfb\x97\x4d\x01\x09\x1b\x05\xf4\xca\xba\xc9\x02\x06\x79\x2a\x99\x6c\x96\x62\x25\xbc\x09\x6d\x58\x3d\x68\xd4\x79\x47\xb6\xdd\xa9\xf1\xf7\x7f\xfb\xbb\xda\x8b\x91\xd4\x2f\x7f\x97\x6e\x79\xba\xe0\x3c\xdf\x38\x1b\xd1\x2c\xa7\x05\x7c\x8e\x21\x9a\x79\x0b\xf9\x4d\x37\x26\x5f\x3f\xfc\xe3\xdf\x77\x7d\x95\x89\x94\xeb\x8d\x4b\x63\x84\x77\xb1\x3b\xa6\x76\x40\xcf\x68\x63\x22\x67\xe5\x6c\x6b\x5c\x2f\xb5\xa0\x7a\x9b\x40\x5a\x8a\x57\x92\xf8\x4b\x7c\x08\x51\xa8\xd8\xa7\x6a\xf5\x45\x96\xdd\xbb\x07\xaf\xe6\x81\x82\xf8\x16\xf3\xa0\xc5\x17\xb6\x7b\x82\x38\x09\x4b\xb8\x47\x3f\x43\x85\x91\x5a\xe7\xe7\xd5\x10\xce\x30\x45\xb4\x14\xcb\xdf\x69\xb6\x3d\xdc\x26\x6d\x66\xe8\xc8\xd4\x9b\x87\x59\xb6\x83\x4b\xc8\xdf\x2c\x19\x7f\xff\xcd\xbf\x34\x63\x29\x51\x70\xb6\x3d\xa6\xfa\xaa\xa3\x59\x1d\x7a\x6a\x47\x83\x1e\x02\xdf\x50\xad\x2d\x66\xd0\xb7\xb4\x1e\xa4\x8e\xd9\x2e\x44\x4e\x0d\x30\x71\x4d\xe0\xd1\xb6\xda\xa5\xca\xb5\x00\x25\xc5\x89\xc8\xa6\x0e\xaa\x89\x86\xc4\xf4\x00\x8d\x77\x3d\x94\x18\x02\xfc\x32\x65\xaf\x8f\xa8\x59\xc7\x6d\x07\x03\xc7\xaa\xbb\x63\xbc\xd8\x15\x29\x95\x12\x85\x0a\xff\xf9\x56\xb2\xf9\xeb\x1f\x8f\xd9\xc8\xf1\xfb\xd9\x84\x5e\x2a\x76\x86\x5e\x82\x82\xf3\x0f\xab\x7e\x73\xf7\xb9\x9d\x3e\x47\x75\x01\x79\xf9\xdf\x3f\xd8\x36\x79\xdd\x24\x45\x8b\xf3\xa0\x29\x45\x8f\x35\x0b\xdf\xd1\x1c\x99\xee\x06\xf2\x08\xfd\x18\xb8\x82\x49\xd1\x8b\x1d\x87\x14\x01\x38\xcf\x2d\x5b\xa1\x85\xbe\xc6\x02\x75\x24\x0b\x63\xa0\x1a\xce\x30\x00\x15\x6d\x01\xb9\xf0\xc3\x10\xd6\x8b\xb3\xc0\xb6\x25\x9f\x6f\x12\x45\xbe\x10\xd7\x68\xc0\xba\xa8\xbd\x96\x65\x2f\x09\xbd\x99\x45\x8e\x13\x1b\xd2\xe3\xb3\x1b\xa1\x61\x2b\x22\x61\xd4\x5a\xe4\x80\x2d\xa0\x9d\x95\x70\xca\x97\x45\xf7\x56\x5f\x05\xbc\xb0\x21\x12\xd6\x5b\xd1\x25\xbd\x98\x2e\xad\xd9\xfd\x2e\x70\x3f\x88\xac\x2d\x17\x7e\x7f\xb6\x0e\x26\xb2\xc5\xc4\x57\x3c\xc8\xe4\x28\x9c\x6f\xf7\xf2\x69\xff\x72\xec\x4b\xf2\x54\x7f\xd9\xa7\xa8\xbf\x5c\x2f\xde\x7b\x76\x8d\xe2\x29\xa8\x3c\x37\x1a\x2d\xdb\x48\x56\x15\x48\xe6\xd1\x8f\x56\x57\xdd\x6c\xd5\x7a\x8d\x6e\x22\xe5\x7e\x03\x9c\xbc\x18\x42\x7f\x1b\xf0\x3b\xf1\x16\xf0\xa9\x9b\xe8\x40\x7e\xbb\x9a\xbf\xf3\x68\x1a\x08\x32\x43\xb8\x27\xd5\x98\xab\xdd\xe0\x86\xa5\xa6\x8b\x78\x26\x30\x15\xd7\x93\xe2\x5a\x6a\x38\xaa\xcf\x7e\xac\x3a\x55\x4e\x31\xa0\x32\x70\x24\x50\xb1\xc7\x63\xe5\x93\x8b\x46\x91\x65\xdb\x06\x38\x13\xb2\x46\x2c\x51\x66\x6b\xc8\x37\x12\x51\x4d\xd6\xc5\x25\xe9\x13\xc4\x05\x25\x18\x3c\x56\x91\x2b\xda\xc2\x8b\x0f\x0e\x04\x8d\x12\x59\x43\x8a\x4e\x85\x7f\xd5\x18\x37\x44\xee\xd1\xc0\x24\x25\x15\x9f\xa1\xf2\x5c\xd2\x29\x55\x56\xb5\x6e\x9c\x4c\x74\x69\x5e\xd1\x7f\x67\xc3\xc3\x2c\xbb\x5f\xc0\xa5\x04\x5e\x52\xcb\xd6\x4a\xe1\x24\xc3\xda\xd9\x0f\x22\x58\x5a\xa6\x28\xf5\xa2\x30\xe9\x51\x09\x59\x25\x6c\x76\xa3\x5f\xe4\x69\x99\x78\x9d\x6a\x41\x89\xc2\x76\x67\x17\xb5\xba\xa2\xf9\xc4\x94\xc3\x06\x2e\x5f\x3e\x5d\x50\x0c\x1d\x51\x5c\x82\xc4\xa8\xee\x83\xec\x4a\x91\x7b\x2a\xb2\x07\x85\x6e\x57\xe1\x38\x3d\xd3\xd3\x3a\xa9\x7b\xf2\x64\x66\x08\x1d\x37\x51\x08\xdc\x78\x7a\x33\x92\xad\xe6\xdb\xe7\x8f\x48\x16\xf0\xa2\x01\x04\xd5\x3d\x3d\x3f\x81\x4d\x5a\xd7\x79\xd9\x62\x34\xbe\x94\xbd\xcc\x56\x65\x0b\x47\x78\x3d\x86\x93\xc0\x5c\xea\x2f\x17\x3b\x5a\x12\x92\xf4\xb4\xd7\x97\x7e\x5a\x77\x1e\x95\x43\x21\x00\x89\x2c\x71\x58\xa6\xdb\xfb\xc5\x55\x25\x1f\x5c\x58\x20\x4d\x3c\xf4\x30\x8c\xbe\xea\x70\x39\x25\x1b\xd9\x93\xbe\x75\xd4\xdb\xa4\xa1\x53\xc7\x86\x52\x35\xd9\xb6\x45\x76\x9e\xd6\x84\x23\xf1\x12\x21\xe5\x79\x59\xb3\x2a\xa3\xeb\x8a\x13\xfc\x6b\xa9\x38\x42\xcd\x21\xa2\xad\xe4\x1d\xfc\xc9\xb6\x5a\xa0\xf4\x64\x58\xb0\xb7\x10\xc6\x32\x1a\x02\x59\x35\x6d\x84\xda\xc9\x64\x2c\xdd\x81\xb4\x14\x25\x19\x37\x25\xae\xa8\x3c\x84\x13\xae\x87\x81\x2a\x6e\xb8\x4a\x58\x16\xd9\x47\xef\xc6\xbb\x86\x4a\xd7\x83\xa7\x10\xf8\xb0\x44\xf6\x6e\x44\xa8\xab\x95\x20\xb5\x58\xa6\x42\x5d\x4b\xb7\x0b\xfe\x63\xdc\xb9\x66\x27\x15\x4b\xac\xc1\x04\x97\x5c\xf0\x84\xda\xbf\x9f\x8a\x7a\xab\xe0\xa7\x95\x91\x65\x79\xaa\x3a\xf4\xf1\x61\x96\x7d\xf5\x55\x21\xe4\x7b\x7e\xec\xe0\xe7\xce\x7f\x46\x33\xe4\x4f\xf3\xb7\x6f\xb3\xec\x57\x42\x0b\x5d\xda\xd6\xc1\x79\xa2\xb5\xa7\x61\xda\xfa\x0e\xcb\xd7\x2f\xb6\xef\x25\x54\x3b\x0a\xaa\x34\x1d\x6a\x1d\xe7\x05\xad\x1e\x87\x41\x0a\x53\xb3\xa7\x2a\x1a\x6d\xf2\x94\xcc\x42\x2c\x96\xa1\xab\xdb\xa6\x54\x84\x03\xb4\x64\xc9\xa3\x11\x0a\x29\xaf\x24\x06\x1c\x06\xc5\x4e\x2e\x1b\x03\x6c\x43\xf4\x63\x4f\xb2\xd2\x70\x1a\x09\xc2\xa9\x23\xd9\xd5\x7d\x96\x1d\xc5\x34\xde\xc1\x68\x0f\xab\x90\x89\x5f\xef\x50\x96\x12\x85\x3f\x6d\xcf\x29\x54\xa1\x8b\xa6\x16\x3a\x1e\x1e\xc2\x77\x5f\xff\x69\x21\x41\x58\x97\xe7\x3b\x72\x3b\x5a\xdd\xad\x65\xa3\x92\x80\xbf\xfb\xfa\xcf\xb0\x1c\x14\xf0\x85\x30\x82\xfb\xc1\xf9\x88\x56\x97\x39\xed\x15\xae\x6b\x43\xf0\xf8\x7a\xd7\xdf\x7f\xf2\x78\xaf\x7f\xb6\xfa\xf1\x41\xfa\xf8\xe0\x49\xf2\x25\x47\xe7\xe9\xe8\xfc\x49\x62\xeb\x6d\x35\x65\xbc\x48\x57\x8a\xc7\x72\x91\x47\x43\x07\x32\xda\xa1\xa9\x47\xb1\xea\x98\x0e\xb7\xfa\x97\x48\xeb\xac\x49\x3f\x12\xda\x85\x8a\x82\x5e\xe5\xec\x81\x66\x81\x41\x89\x10\xb9\x67\xdb\xee\xeb\xd1\x2f\xb0\xe3\x2a\xe0\x74\x3d\x50\x25\x9b\x42\xcb\x07\xb2\x0b\x74\xb8\xd2\xe2\x38\x2a\xf2\x93\xed\xf1\x96\x8c\x50\xce\xa2\x06\xc7\xd1\x3a\xa1\x8d\xe9\x17\x80\xc8\x6a\xea\xd6\x54\x6a\x95\x5e\xfd\x91\xba\xa8\xa8\x2c\x71\x5c\x5d\xa5\x5f\x09\x34\xdf\x55\x72\x6d\xdd\x1f\x63\xff\x27\xf3\x67\x34\x0b\xf1\xd7\x0d\xe5\x68\xa3\x9f\x3e\xe7\x10\xdf\xbe\xcd\xfe\x1f\x00\x00\xff\xff\xc5\x60\x01\x73\x0f\x0f\x00\x00"),
		},
		"/pages/index.md": &vfsgen۰CompressedFileInfo{
			name:             "index.md",
			modTime:          mustUnmarshalTextTime("2019-05-06T05:10:48.10050764Z"),
			uncompressedSize: 2436,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x96\xdf\x6f\xdb\x36\x10\xc7\xdf\xf5\x57\xdc\x9c\x07\x27\x40\x62\xc7\xeb\x32\xa0\x4b\xdd\xa2\x5d\xd3\x1f\x43\x3b\x14\xcd\x56\xb4\x28\xfa\x70\x92\x4e\xe2\xc5\x24\x4f\xe3\x0f\x2b\x72\xd1\xfd\xed\x03\x29\x25\xce\x8a\x61\x4f\x05\xfa\x66\x58\x77\xc7\xe3\x7d\xbe\x5f\x92\x07\xe0\x35\x1b\x8f\xed\xdc\x83\x89\x9e\x2b\x70\xe4\x25\xba\x8a\x7c\x51\xbc\x20\x47\x30\x48\x9c\x6b\x0d\x0d\xdb\x7a\x8a\x40\x5b\x43\x4b\x96\x1c\x6a\x60\xdb\x88\x33\x18\x58\x2c\x60\x29\x31\xc0\xaf\x8a\x2d\x79\x02\xb6\x3e\xb8\x68\xc8\x06\x7f\x0c\x46\x7c\xd0\x03\x04\x45\xf0\x94\x77\x0c\x8d\x8e\x81\x16\x45\xf1\xb2\x49\xf5\xa1\x16\x3b\x0f\xb0\xb1\xd2\x83\x92\x1e\x82\x80\x23\xac\x6f\x4b\x69\xde\x10\x18\x3a\x4e\xf9\x9e\x00\x1d\x41\xe7\xa4\xc4\x72\x2a\x59\x92\x0f\xfb\xbe\xef\x76\x9c\x96\x98\x1b\x40\xdd\xe3\xe0\x41\x61\xd7\x0d\xa9\x7a\xa5\x30\xfc\x5f\xbb\x5e\x80\xc7\xce\x7a\xb4\x21\x65\x5c\x45\x1f\xc0\x2b\x91\x00\x86\x00\xc1\x90\xf7\xd8\x12\x1c\x7e\xfe\xbc\x78\x3d\x5c\x18\x64\xfd\xe5\xcb\xd1\x0f\x45\x71\x70\x00\x97\xd2\x84\x3e\x35\x59\xd3\x96\xb4\x74\xa9\x28\x2c\xe1\x4a\x4a\x90\xa6\x21\xe7\x6f\xf7\x9d\x82\x54\x1a\x72\x23\x0e\xcc\x00\xbd\xb8\x0d\xa0\x07\x04\xff\x55\x0d\x72\x69\xf7\xec\xa1\xa7\xd2\x73\x20\x60\x9f\x46\x96\xf2\x06\x89\x0b\xb8\x24\x4a\x05\x3e\x3e\xe7\xf0\x22\x96\x9f\x0e\x55\x08\x9d\xff\x65\xb9\x6c\x39\xa8\x58\x2e\x2a\x31\xcb\x09\xf4\x11\x88\x83\x8f\xaf\xd8\x6e\xa8\x7e\x69\xf7\x91\x7d\xdf\x2f\x74\xfe\x97\x6d\x8e\x67\xbb\xf4\x81\x3a\x45\xf6\xa4\x8d\x81\x36\x68\x7d\x38\x59\xad\xee\xdd\x3f\xc3\xd5\xea\x74\x79\x94\x18\x34\xac\xc9\xe7\xd9\x11\xa6\x61\x1f\x1c\x8c\x7c\x0f\x6f\xc6\xfa\x04\x4d\x29\x02\xcf\x12\xef\xa3\xa2\x78\xc0\x8d\x43\x43\xd0\x73\x1d\xd4\x7a\x76\xf6\xf3\xe9\x0c\x14\x71\xab\xc2\x7a\x76\x6f\x75\x36\x03\xef\xaa\xf5\xec\x6e\x4b\x83\xc4\x10\x4b\xca\x1d\x91\x29\xa9\x5e\x5e\xbc\xfd\x30\x98\xab\x67\xef\x89\x5f\x3e\xf2\x01\x5d\x58\xff\x74\x6f\x06\xb9\x6e\x29\xae\x26\xb7\x9e\x9d\xce\x00\xb5\x96\x7e\x3d\xc3\xaa\x22\x4d\x4e\x0c\x05\x72\xe7\x80\x31\x48\xa7\x71\x38\x07\xb2\x95\x1b\xba\x40\xf5\x89\xa1\x9a\xf1\x1c\xda\xc1\x89\xaf\xa4\xa3\x73\xe8\xb8\x0a\xd1\xd1\x09\xdb\x93\xe9\xe7\x54\xaf\x89\x5a\xfb\xca\x11\xd9\x87\x0f\x96\xe3\x56\x1e\x16\xc5\x1f\x89\x0b\x27\x6a\x25\x87\x51\xac\x12\x14\xb9\x9b\xcd\x67\xb1\x7b\xa0\xeb\x8a\xba\x00\x1c\x40\x65\xc4\x5b\x72\x49\xbe\x6c\xa1\xc3\x8e\x1c\x18\x32\xa5\x43\x4b\x50\xc9\x96\x1c\xdb\x16\xc4\x12\x28\xd1\x74\x0c\xbd\xe2\x4a\x41\x85\xd1\xa7\x71\x67\x45\x76\x4e\xea\x58\x25\x2d\x96\x71\xb7\x4b\xe1\x56\xd8\x13\x1c\x7a\x89\xb6\xf6\xe0\xc5\x50\x9f\xc5\xc5\xb6\xa4\xd0\x13\x59\x40\x70\x54\xe5\x19\x65\x13\x23\x6c\x70\x27\x72\x0c\x35\x75\x64\xeb\x71\xc9\xd1\x81\x18\x6d\xc8\x12\x6d\x75\xa4\x6c\xb2\xdb\xf6\xc4\x1e\x2d\x8a\xe2\xe3\x3b\xa6\x7e\x84\x7d\xeb\xbc\x4f\x87\x8f\xba\x75\xcd\x3b\x5e\xb2\xad\xe9\xfa\x28\xeb\xe1\xc2\xa9\xb8\xd7\xc3\x96\x45\xb3\xfd\x46\x4a\x28\xab\xa7\xdb\x2b\x75\xff\xba\x7c\x2f\x93\x12\xbe\xaf\x10\x9e\xa0\xe7\x0a\xb5\x1e\x00\x6f\xcf\x95\x2d\x39\x9f\x0e\x47\x69\xf2\x14\xdf\xe5\xfd\x1f\xdf\x11\x83\x58\x3d\x64\x45\x24\xda\x3e\x64\xf0\x09\x0e\x7b\xe8\xa2\x57\x54\xdf\xd5\x15\x42\x1b\x39\xa0\xfb\xb7\x00\xfe\x8a\xe9\x44\x28\x9d\x60\x0d\x0e\x6d\x4b\x69\xb5\x51\x0d\xe3\xa7\x9c\x9b\x96\x57\xd1\xa0\x85\xad\x70\x45\xa3\x5b\x9f\xc7\x9d\x22\xdb\xee\x01\xed\x38\x49\xf7\x1b\x01\xba\xbe\xd4\x8f\x5f\x9f\xd5\x7f\xfa\xdf\xe4\xbb\x82\x79\x35\xce\xce\xf0\x35\xec\xad\xd0\x31\x5a\x99\x7c\x30\x4d\x75\x7d\x04\x1f\x24\x42\xa7\x63\xb5\x81\xaf\x67\x7e\x0c\x65\xbc\xb5\xef\xdf\x3f\xae\x26\x58\x3e\xdf\x16\x37\xa6\x36\xb1\x52\xa0\xd1\xb5\xe4\x26\x14\x41\xa1\xdd\xaf\x50\xa1\x85\x8e\x5c\xba\x34\x47\x00\x6f\xb8\xc3\xfd\xf4\xe7\x63\xd8\xfc\x1b\xcd\xff\xf5\x9b\x57\xe1\xf4\xd9\xe3\xfa\x2d\x5f\x4c\x06\x59\xdd\xff\xae\x20\x2e\xc5\x85\x24\xce\xaf\x06\x3b\xb9\x21\x29\x34\x48\x07\x8d\xa3\xe0\xf3\xe5\x38\xbd\x1b\xa2\xa7\x3a\xdf\x75\x95\x23\x0c\xd9\x21\x30\xef\x71\x3b\xcc\x27\x99\x27\x8a\x29\xfb\x06\x49\xce\x4d\x28\x82\xc4\x36\x1d\xc5\x9d\x1a\x6e\xbc\x39\xe2\x9a\xdc\x93\x98\x29\x1e\x43\x38\xa4\x84\xbb\xb8\x16\xf0\xbb\xa4\x97\x45\xce\x9b\xe7\x07\x82\xa1\x7a\xfe\x5f\xba\x70\x98\x8f\x7c\x87\x1d\xd7\x7a\x18\x25\x44\x35\xf4\x1c\x14\x84\x5e\xd2\xa5\xdb\xf0\x96\xd2\xab\xa4\x25\xe7\x21\x3d\x40\x20\xb0\xa1\x45\xf1\x4f\x00\x00\x00\xff\xff\x82\x2e\xae\x72\x84\x09\x00\x00"),
		},
		"/songs": &vfsgen۰DirInfo{
			name:    "songs",
			modTime: mustUnmarshalTextTime("2019-04-21T22:36:45.045873893Z"),
		},
		"/songs/bangtan_boys-save_me.txt": &vfsgen۰CompressedFileInfo{
			name:             "bangtan_boys-save_me.txt",
			modTime:          mustUnmarshalTextTime("2019-04-21T22:36:45.04522815Z"),
			uncompressedSize: 2029,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x54\x4b\x4e\xc6\x20\x10\xde\x73\x0a\x92\xff\x0a\x8c\x8b\xee\xc0\xe1\x77\xe5\xca\x13\xb0\x20\x6a\x34\x6d\xa2\x68\xd2\xdb\x1b\x28\x23\x0c\xad\xd5\x1a\x6b\xba\xf8\x93\xe1\x91\x8f\x79\x3f\x08\x8f\xe1\xd9\x77\xf2\xce\xbd\x7b\x79\xeb\x85\x7b\x0b\x0f\xc3\x4b\x27\x8d\xeb\xef\x83\xeb\xa5\x19\xc6\x57\xf1\xe4\xc7\x4e\xa2\x10\x06\x32\x69\x10\xe7\x13\x88\xb4\xd1\xdd\x4e\xa7\x05\x81\x89\xe8\x71\xa3\x14\x02\x61\xf1\x34\x2a\x72\x72\x6c\x06\xec\xcb\x84\x04\x17\x3f\xe3\x51\x13\xb1\xb6\xd8\x06\xd6\x68\x42\xfd\x98\xdb\x92\x93\x3c\x71\x55\xc6\x75\xca\x38\x5e\x4d\x2a\x6e\x5a\xef\xb1\x36\xfc\x3b\xb9\x5c\xc8\x7a\x31\xb7\x8c\xca\xfb\x3a\x1b\xfe\xbf\x9e\x52\xcb\x1c\x78\x8c\x5d\x2f\xb5\xef\xb2\xb2\xcf\x34\x71\xb9\x35\x65\x3b\x8d\xce\x37\x31\x35\x9d\xf3\x55\x3b\x5e\x9f\x80\xad\x1c\x26\x42\x21\x5b\xd9\x9f\x34\x55\x66\x16\x41\x26\x97\x64\x8d\xca\xb3\xc4\xa2\x68\xef\xb4\x34\xcc\x2a\xd0\xc0\xd4\xbe\x7f\xaf\xf4\x32\x4a\x07\x1e\xa5\xf3\xea\x37\x6c\x54\xf2\x6b\xf7\x97\x32\x23\xfc\xaf\x47\x6a\xc9\x38\x48\x87\x2b\xeb\x47\x00\x00\x00\xff\xff\x0a\x15\x36\xdf\xed\x07\x00\x00"),
		},
		"/songs/girls_generation-gee.txt": &vfsgen۰CompressedFileInfo{
			name:             "girls_generation-gee.txt",
			modTime:          mustUnmarshalTextTime("2019-04-21T22:36:45.045620288Z"),
			uncompressedSize: 2556,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xe4\x55\xcd\x6e\x83\x30\x0c\xbe\xfb\x29\x2c\x71\xe7\x54\x38\x70\x2b\x81\xe4\x39\x7a\x88\x34\xb4\xaa\x95\x18\x3b\xec\xed\x27\x27\x31\x75\xc2\xcf\xfa\xc3\xb4\x9f\x1e\x3e\x1c\x6c\x93\x7c\x38\xfe\x92\xa1\x1b\x8e\xb6\x42\x63\x2d\x1c\xde\x87\x97\x73\x5f\xa1\xe9\xfa\xe3\x1b\x1a\x7b\xb2\xfd\x61\xe8\xce\x27\x78\xb5\x1f\x15\xb6\x00\x26\x2b\x40\x67\x05\xb4\x05\x82\x12\xd6\x0d\xa0\x2d\x11\x74\x56\x3a\xdb\x08\x5b\x53\x90\x1e\xfc\xa9\x1c\xb3\x5d\xff\x9a\x32\x2e\xf0\x44\x28\xc0\x84\xa4\xbd\xd5\x4f\x18\xff\x66\xf4\x30\xd4\x02\x6e\x4b\xdc\x94\x2b\x07\xaf\xa7\xa0\x42\x1d\xe5\xfe\x49\x78\xdf\x72\x1e\xbf\x13\x1a\xa6\x30\x07\x17\x65\xd0\xd2\xf5\x2e\x74\xc8\xbe\x40\x87\x91\x2c\xcf\xdf\x84\xb5\x53\x4e\x4c\x9e\xf9\x4d\x12\x24\xea\x10\xd7\x61\x49\x0f\x10\x63\x8c\x26\x33\x72\xe2\xb4\x8c\xf7\xa0\x16\x6d\xbd\x86\x09\x93\x96\xdb\x5b\xdf\xb9\x32\x15\xcf\xd7\xab\x9c\xb4\xcd\x52\x2c\xff\x32\xe8\x1f\xfb\x99\xb2\x7f\x97\xf6\xfe\x9f\xf4\xae\x94\x1d\xfe\xac\xee\x1e\x15\x1e\xce\x29\x0f\x97\xa5\xf7\x90\xbc\xfe\x92\xbe\x72\x2a\x3b\x2b\x2d\xa9\xda\x0e\xe3\xbb\x33\xbd\xdf\xe4\x26\xb9\x22\x0a\x5f\x7a\x7b\xd2\x00\xf9\x44\xf3\x2e\xc5\x2f\x2a\x8a\x47\x39\xd1\x4f\xea\x4b\x63\x08\x66\x2b\x27\xe7\x06\x7b\xf9\x3b\xb6\x32\x27\x47\xfe\xa4\x8d\xf9\x19\x00\x00\xff\xff\x30\x72\x88\xbe\xfc\x09\x00\x00"),
		},
		"/songs/spinning_song_tsumugi_uta.txt": &vfsgen۰CompressedFileInfo{
			name:             "spinning_song_tsumugi_uta.txt",
			modTime:          mustUnmarshalTextTime("2019-04-21T22:36:45.045768862Z"),
			uncompressedSize: 1137,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x52\xb1\x4e\x03\x31\x0c\xdd\xfd\x15\x96\x2a\x21\x98\xba\x34\x37\xdc\x96\xf4\xae\x0b\x9d\x28\x7c\x40\x24\xa2\x23\x82\xba\x08\xee\x06\x76\x3e\x81\x1f\xe0\x1f\xf8\x25\xc4\x6f\x20\x9f\xe3\xa4\x3d\x2a\x21\x46\xa4\x0e\x51\xee\x9e\x9f\x9f\x9f\xed\xf4\xb1\x7f\x08\x35\x6e\x1e\x23\x51\xa4\x0e\x37\x3b\xea\x70\x8e\xd7\xcf\xc3\x76\xe8\x22\xde\xf4\x1e\xe7\xf8\xf5\xf1\xfe\xf9\xf6\x0a\x7e\xe8\xef\x76\x4f\x35\x5e\x45\xc2\x33\x5c\x07\xc2\xf3\x4b\xdf\xf9\x6d\xa4\x30\x62\x9e\x6e\x31\x03\xeb\x40\x17\x70\x1f\x5e\x6a\x6c\x00\xdc\x02\x1a\x03\xd0\x1a\xb0\x06\x56\x33\x39\x05\xe4\x3f\xbb\x60\x60\x82\x19\x68\x2a\x58\xce\x2a\xbd\xdc\x88\x69\xac\x35\x4c\x3e\xc8\x3a\xc9\xff\x90\xe7\xc3\x18\x23\x29\xb5\x94\x52\xb6\x53\xf6\x61\x01\x89\x64\x53\x2c\xef\x84\x5c\xa5\x58\x16\x11\x05\x11\x99\x94\xb3\xb9\x2a\xab\x4b\x91\x36\xb7\xf9\xab\x3d\x57\xa6\x72\xcc\x84\xf8\xb0\x13\x9b\x65\x26\x4e\x3b\xd9\xef\xd7\x16\x66\x31\xd9\x1a\x58\xca\xb0\x8f\x99\xfc\xcf\x4f\x20\x4d\x55\x5b\x1e\xef\x0a\x56\x39\x43\x17\xb0\xb7\x26\xfe\x6c\x12\xe7\xef\x0a\x69\x86\x65\x37\x4e\x17\xf5\x1d\x00\x00\xff\xff\xa9\x10\x36\xdd\x71\x04\x00\x00"),
		},
		"/songs/tengen_toppa-sorairo_days.txt": &vfsgen۰CompressedFileInfo{
			name:             "tengen_toppa-sorairo_days.txt",
			modTime:          mustUnmarshalTextTime("2019-04-21T22:36:45.045931603Z"),
			uncompressedSize: 962,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x52\x3b\x0e\x83\x30\x0c\xdd\x7d\x8a\x48\x5c\x01\x18\xd8\x08\xa6\x39\x40\xb9\x40\x86\xa8\x45\xad\x08\xa2\x74\xe0\xf6\x95\x1d\x03\x41\x55\x3f\x03\x55\xa5\x67\x48\xec\x67\x3f\xdb\xca\xd8\x8e\x57\x57\xa8\xa3\x1f\x6c\x3b\x78\x85\x76\xba\x81\xbd\x8f\x67\x3f\x14\xaa\x71\xdd\xc9\x75\xaa\xf1\x7d\x6f\xe1\xe2\xa6\x42\x21\x80\x4e\x09\x55\x92\x2d\xa6\x53\x88\x5d\x5b\xc2\x67\x3e\x32\xf6\xaf\x5b\x32\xe0\x90\x64\x2f\x8d\x84\x6b\x06\x5f\xc3\x91\x03\xeb\x95\x2c\xb4\x58\x86\xda\x92\x84\xef\x92\x70\xc9\x0b\x5f\x0e\xff\x46\x63\x91\x99\x07\xa2\xa1\xcb\xb9\x1e\xf2\x5a\x84\x4a\x0c\x13\x7e\x42\x32\x1b\x27\x71\xbe\xca\x8e\x5a\x33\x0c\xbe\x6e\xea\xc5\xab\x37\x21\x48\x47\xcc\xa1\x4a\xf2\x75\x14\x09\xc5\x88\x44\xea\x78\x8d\xf3\x7c\xbb\xe8\x68\xf1\x13\x09\x73\xe0\x3f\x99\xe6\x54\xa6\xaf\x02\xf2\x42\x9f\x9c\xd2\xde\x9f\x9f\xee\x23\x00\x00\xff\xff\xcc\x50\x1f\xc4\xc2\x03\x00\x00"),
		},
	}
	fs["/"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/pages"].(os.FileInfo),
		fs["/songs"].(os.FileInfo),
	}
	fs["/pages"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/pages/dizi"].(os.FileInfo),
		fs["/pages/index.md"].(os.FileInfo),
	}
	fs["/pages/dizi"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/pages/dizi/index.md"].(os.FileInfo),
	}
	fs["/songs"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/songs/bangtan_boys-save_me.txt"].(os.FileInfo),
		fs["/songs/girls_generation-gee.txt"].(os.FileInfo),
		fs["/songs/spinning_song_tsumugi_uta.txt"].(os.FileInfo),
		fs["/songs/tengen_toppa-sorairo_days.txt"].(os.FileInfo),
	}

	return fs
}()

type vfsgen۰FS map[string]interface{}

func (fs vfsgen۰FS) Open(path string) (http.File, error) {
	path = pathpkg.Clean("/" + path)
	f, ok := fs[path]
	if !ok {
		return nil, &os.PathError{Op: "open", Path: path, Err: os.ErrNotExist}
	}

	switch f := f.(type) {
	case *vfsgen۰CompressedFileInfo:
		gr, err := gzip.NewReader(bytes.NewReader(f.compressedContent))
		if err != nil {
			// This should never happen because we generate the gzip bytes such that they are always valid.
			panic("unexpected error reading own gzip compressed bytes: " + err.Error())
		}
		return &vfsgen۰CompressedFile{
			vfsgen۰CompressedFileInfo: f,
			gr: gr,
		}, nil
	case *vfsgen۰DirInfo:
		return &vfsgen۰Dir{
			vfsgen۰DirInfo: f,
		}, nil
	default:
		// This should never happen because we generate only the above types.
		panic(fmt.Sprintf("unexpected type %T", f))
	}
}

// vfsgen۰CompressedFileInfo is a static definition of a gzip compressed file.
type vfsgen۰CompressedFileInfo struct {
	name              string
	modTime           time.Time
	compressedContent []byte
	uncompressedSize  int64
}

func (f *vfsgen۰CompressedFileInfo) Readdir(count int) ([]os.FileInfo, error) {
	return nil, fmt.Errorf("cannot Readdir from file %s", f.name)
}
func (f *vfsgen۰CompressedFileInfo) Stat() (os.FileInfo, error) { return f, nil }

func (f *vfsgen۰CompressedFileInfo) GzipBytes() []byte {
	return f.compressedContent
}

func (f *vfsgen۰CompressedFileInfo) Name() string       { return f.name }
func (f *vfsgen۰CompressedFileInfo) Size() int64        { return f.uncompressedSize }
func (f *vfsgen۰CompressedFileInfo) Mode() os.FileMode  { return 0444 }
func (f *vfsgen۰CompressedFileInfo) ModTime() time.Time { return f.modTime }
func (f *vfsgen۰CompressedFileInfo) IsDir() bool        { return false }
func (f *vfsgen۰CompressedFileInfo) Sys() interface{}   { return nil }

// vfsgen۰CompressedFile is an opened compressedFile instance.
type vfsgen۰CompressedFile struct {
	*vfsgen۰CompressedFileInfo
	gr      *gzip.Reader
	grPos   int64 // Actual gr uncompressed position.
	seekPos int64 // Seek uncompressed position.
}

func (f *vfsgen۰CompressedFile) Read(p []byte) (n int, err error) {
	if f.grPos > f.seekPos {
		// Rewind to beginning.
		err = f.gr.Reset(bytes.NewReader(f.compressedContent))
		if err != nil {
			return 0, err
		}
		f.grPos = 0
	}
	if f.grPos < f.seekPos {
		// Fast-forward.
		_, err = io.CopyN(ioutil.Discard, f.gr, f.seekPos-f.grPos)
		if err != nil {
			return 0, err
		}
		f.grPos = f.seekPos
	}
	n, err = f.gr.Read(p)
	f.grPos += int64(n)
	f.seekPos = f.grPos
	return n, err
}
func (f *vfsgen۰CompressedFile) Seek(offset int64, whence int) (int64, error) {
	switch whence {
	case io.SeekStart:
		f.seekPos = 0 + offset
	case io.SeekCurrent:
		f.seekPos += offset
	case io.SeekEnd:
		f.seekPos = f.uncompressedSize + offset
	default:
		panic(fmt.Errorf("invalid whence value: %v", whence))
	}
	return f.seekPos, nil
}
func (f *vfsgen۰CompressedFile) Close() error {
	return f.gr.Close()
}

// vfsgen۰DirInfo is a static definition of a directory.
type vfsgen۰DirInfo struct {
	name    string
	modTime time.Time
	entries []os.FileInfo
}

func (d *vfsgen۰DirInfo) Read([]byte) (int, error) {
	return 0, fmt.Errorf("cannot Read from directory %s", d.name)
}
func (d *vfsgen۰DirInfo) Close() error               { return nil }
func (d *vfsgen۰DirInfo) Stat() (os.FileInfo, error) { return d, nil }

func (d *vfsgen۰DirInfo) Name() string       { return d.name }
func (d *vfsgen۰DirInfo) Size() int64        { return 0 }
func (d *vfsgen۰DirInfo) Mode() os.FileMode  { return 0755 | os.ModeDir }
func (d *vfsgen۰DirInfo) ModTime() time.Time { return d.modTime }
func (d *vfsgen۰DirInfo) IsDir() bool        { return true }
func (d *vfsgen۰DirInfo) Sys() interface{}   { return nil }

// vfsgen۰Dir is an opened dir instance.
type vfsgen۰Dir struct {
	*vfsgen۰DirInfo
	pos int // Position within entries for Seek and Readdir.
}

func (d *vfsgen۰Dir) Seek(offset int64, whence int) (int64, error) {
	if offset == 0 && whence == io.SeekStart {
		d.pos = 0
		return 0, nil
	}
	return 0, fmt.Errorf("unsupported Seek in directory %s", d.name)
}

func (d *vfsgen۰Dir) Readdir(count int) ([]os.FileInfo, error) {
	if d.pos >= len(d.entries) && count > 0 {
		return nil, io.EOF
	}
	if count <= 0 || count > len(d.entries)-d.pos {
		count = len(d.entries) - d.pos
	}
	e := d.entries[d.pos : d.pos+count]
	d.pos += count
	return e, nil
}
