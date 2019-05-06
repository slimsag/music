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
			modTime:          mustUnmarshalTextTime("2019-05-06T04:55:27.346416353Z"),
			uncompressedSize: 3859,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x57\xb1\x92\xdc\xc6\x11\xcd\xf1\x15\x5d\x60\xa0\xdb\xaa\x5d\xac\xc4\x93\x13\x92\x47\xd7\xc9\x24\x2d\x5a\x12\x03\x17\x03\xbb\x5c\x2e\x55\x03\x68\x00\xcd\x1b\xcc\x80\x33\x83\xc5\xe1\x54\x0c\x9c\x38\x71\x60\x67\x0a\x94\x39\x72\xd9\x56\xe2\xcc\xb1\xfd\x25\xaa\xb2\xf4\x1b\xae\xee\x01\xf6\xf6\x8e\x52\xc2\xe3\xce\xf6\xf4\x74\xf7\x7b\xfd\xba\xf7\x01\x3c\xe3\x1b\x86\xb3\x5f\x74\x6c\x29\x10\x44\x8f\x36\x1c\xc8\x07\x82\xc6\x8c\x91\x36\x59\xf6\x84\x1b\x8f\x3d\xc1\xc4\x75\xec\x2e\xf2\xf3\x9f\x79\xea\x73\xe8\x88\xdb\x2e\x5e\xe4\x0f\x3f\xd4\x8f\xc1\x57\x17\x79\x17\xe3\x10\x1e\xed\xf7\xd3\x34\x15\xb3\x1b\xe3\x58\x52\x51\xb9\x7e\x4f\x7d\x49\xf5\xfe\xf9\xaf\x7f\x3b\xf7\x6f\x5e\xfc\x86\xf8\xe5\xcf\x43\x44\x1f\x2f\x3e\x3e\xcf\x41\x7d\x97\xce\xd7\xe4\x2f\xf2\x0f\x73\x40\x63\xdc\x74\x91\x63\x55\x91\x21\xef\x7a\x8a\xe4\x1f\x03\x8e\xd1\x0d\x06\xe7\xc7\x40\xb6\xf2\xf3\x10\xa9\xde\xf5\x54\x33\x3e\x86\x76\xf6\x2e\x54\x6e\xa0\xc7\x30\x70\x15\x47\x4f\x3b\xb6\xbb\xe5\xbf\x8b\xbf\x66\x34\x26\x54\x9e\xc8\x3e\x7d\xb2\x4f\xe9\x3c\xcd\xb2\x4b\xa8\x25\x77\x0e\x80\xf0\x53\xf9\x6f\x61\x0c\x23\x1a\x33\x43\x8f\x35\x81\x1b\x23\xb8\x06\x3e\xc1\xbe\x74\x0e\xce\xca\x31\x42\x70\x3d\x41\xef\xbc\xfc\x53\x93\xb7\xe0\x2c\x05\x40\x4f\x80\x26\xb8\xe5\x5a\x03\x83\xc1\x10\xb9\xda\x42\x4f\x11\xcd\x16\xd0\xd6\x70\xd9\xb4\x68\xe1\x0d\xd6\xb4\x29\xb2\xec\x65\x94\x50\x46\xcb\x6f\x47\x82\xca\xf5\x03\x7a\xaa\x21\xba\xd5\xef\x44\x21\xca\xdf\xca\xd9\x8a\x7c\x4c\xf1\x05\x60\x0b\xb1\x23\x68\xb0\x8a\x10\x3b\x8c\xc0\x11\x3a\x94\x9c\x0e\xe4\x67\x88\x1d\x5b\x18\x70\x20\x0f\x3d\xf5\xa5\x47\x2b\xce\x0f\xe4\xd9\xb6\x80\xd0\x39\x43\x60\x09\xbd\x3a\xe9\xdd\x18\xbb\x81\xa9\x22\x98\x3a\xae\x3a\xa8\x70\x0c\xf2\x46\x94\x38\x06\xef\xea\xb1\x22\x40\x28\xc7\x9b\x1b\xb9\x6f\x1d\x07\xd2\x0a\x4c\xf2\xb2\xe1\x2b\xf9\xd6\x53\xa5\x78\x42\xcf\xd7\x54\xc3\xc4\xb1\x03\x84\x2b\xbc\x71\xae\x80\xd7\xf2\xcc\x31\x10\xb4\x50\x12\x44\xa1\x12\xd9\x94\xad\xa1\x10\x5c\xca\x49\xaf\x40\x70\xa3\xad\xb7\xe0\x3c\x18\xe7\xc2\x6a\xc6\xb6\xf2\x84\x81\x80\x63\x91\x65\xe2\x75\x72\xbe\x86\x5c\x20\xcd\x81\x63\x20\xd3\x28\xb4\x16\x9e\xdb\xd6\x70\xe8\x12\xb6\x86\x23\x79\x8c\xec\x2c\x34\x2e\x65\xbd\x62\xaf\x0e\xe4\x30\xd7\xd2\xe6\x70\xf6\xc3\x3f\xbf\xf9\xdf\xb7\x7f\xd9\x14\x90\xb0\x51\x40\xaf\xac\x9b\x2c\x60\x90\xa7\x92\xc9\x66\x29\x56\xc2\x9b\xd0\x86\xd5\x83\x46\x9d\x77\x64\xdb\x9d\x1a\x7f\xff\xb7\xbf\xab\xbd\x18\x49\xfd\xf2\xfb\x74\xcb\xd3\x05\xe7\xf9\xc6\xd9\x88\x66\x39\x2d\xe0\x73\x0c\xd1\xcc\x5b\xc8\x6f\xba\x31\xf9\xfa\xe1\x1f\xff\xbe\xeb\xab\x4c\xa4\x5c\x6f\x5c\x1a\x23\xbc\x8b\xdd\x31\xb5\x03\x7a\x46\x1b\x13\x39\x2b\x67\x5b\xe3\x7a\xa9\x05\xd5\xdb\x04\xd2\x52\xbc\x92\xc4\x5f\xe2\x43\x88\x42\xc5\x3e\x55\xab\x2f\xb2\xec\xc1\x03\x78\x3d\x0f\x14\xc4\xb7\x98\x07\x2d\xbe\xb0\xdd\x13\xc4\x49\x58\xc2\x3d\xfa\x19\x2a\x8c\xd4\x3a\x3f\xaf\x86\x70\x86\x29\xa2\xa5\x58\xfe\x4e\xb3\xed\xe1\x36\x69\x33\x43\x47\xa6\xde\x3c\xca\xb2\x1d\x5c\x42\xfe\x76\xc9\xf8\xfb\x6f\xfe\xa5\x19\x4b\x89\x82\xb3\xed\x31\xd5\xd7\x1d\xcd\xea\xd0\x53\x3b\x1a\xf4\x10\xf8\x86\x6a\x6d\x31\x83\xbe\xa5\xf5\x20\x75\xcc\x76\x21\x72\x6a\x80\x89\x6b\x02\x8f\xb6\xd5\x2e\x55\xae\x05\x28\x29\x4e\x44\x36\x75\x50\x4d\x34\x24\xa6\x07\x68\xbc\xeb\xa1\xc4\x10\xe0\x97\x29\x7b\x7d\x44\xcd\x3a\x6e\x3b\x18\x38\x56\xdd\x1d\xe3\xc5\xae\x48\xa9\x94\x28\x54\xf8\xcf\xb7\x92\xcd\x5f\xff\x78\xcc\x46\x8e\xdf\xcf\x26\xf4\x52\xb1\x33\xf4\x12\x14\x9c\x7f\x58\xf5\x9b\xbb\xcf\xed\xf4\x39\xaa\x0b\xc8\xcb\xff\xfe\xc1\xb6\xc9\xeb\x26\x29\x5a\x9c\x07\x4d\x29\x7a\xac\x59\xf8\x8e\xe6\xc8\x74\x37\x90\x47\xe8\xc7\xc0\x15\x4c\x8a\x5e\xec\x38\xa4\x08\xc0\x79\x6e\xd9\x0a\x2d\xf4\x35\x16\xa8\x23\x59\x18\x03\xd5\x70\x86\x01\xa8\x68\x0b\xc8\x85\x1f\x86\xb0\x5e\x9c\x05\xb6\x2d\xf9\x7c\x93\x28\xf2\x85\xb8\x46\x03\xd6\x45\xed\xb5\x2c\x7b\x45\xe8\xcd\x2c\x72\x9c\xd8\x90\x1e\x9f\xdd\x08\x0d\x5b\x11\x09\xa3\xd6\x22\x07\x6c\x01\xed\xac\x84\x53\xbe\x2c\xba\xb7\xfa\x2a\xe0\xa5\x0d\x91\xb0\xde\x8a\x2e\xe9\xc5\x74\x69\xcd\xee\x77\x81\xfb\x41\x64\x6d\xb9\xf0\xfb\xb3\x75\x30\x91\x2d\x26\xbe\xe2\x41\x26\x47\xe1\x7c\xbb\x97\x4f\xfb\x57\x63\x5f\x92\xa7\xfa\xcb\x3e\x45\xfd\xe5\x7a\xf1\xc1\xf3\x6b\x14\x4f\x41\xe5\xb9\xd1\x68\xd9\x46\xb2\xaa\x40\x32\x8f\x7e\xb4\xba\xea\x66\xab\xd6\x6b\x74\x13\x29\xf7\x1b\xe0\xe4\xc5\x10\xfa\xdb\x80\xef\xc5\x5b\xc0\xa7\x6e\xa2\x03\xf9\xed\x6a\x7e\xef\xd1\x34\x10\x64\x86\x70\x4f\xaa\x31\x57\xbb\xc1\x0d\x4b\x4d\x17\xf1\x4c\x60\x2a\xae\x27\xc5\xb5\xd4\x70\x54\x9f\xfd\x58\x75\xaa\x9c\x62\x40\x65\xe0\x48\xa0\x62\x8f\xc7\xca\x27\x17\x8d\x22\xcb\xb6\x0d\x70\x26\x64\x8d\x58\xa2\xcc\xd6\x90\x6f\x24\xa2\x9a\xac\x8b\x4b\xd2\x27\x88\x0b\x4a\x30\x78\xac\x22\x57\xb4\x85\x97\x1f\x1c\x08\x1a\x25\xb2\x86\x14\x9d\x0a\xff\xaa\x31\x6e\x88\xdc\xa3\x81\x49\x4a\x2a\x3e\x43\xe5\xb9\xa4\x53\xaa\xac\x6a\xdd\x38\x99\xe8\xd2\xbc\xa2\xff\xce\x86\x47\x59\xf6\x51\x01\x97\x12\x78\x49\x2d\x5b\x2b\x85\x93\x0c\x6b\x67\x3f\x88\x60\x69\x99\xa2\xd4\x8b\xc2\xa4\x47\x25\x64\x95\xb0\xd9\x8d\x7e\x91\xa7\x65\xe2\x75\xaa\x05\x25\x0a\xdb\x9d\x5d\xd4\xea\x8a\xe6\x13\x53\x0e\x1b\xb8\x7c\xf5\x6c\x41\x31\x74\x44\x71\x09\x12\xa3\xba\x0f\xb2\x2f\x45\xee\xa9\xc8\x1e\x16\xba\x61\x85\xe3\xf4\x4c\x4f\xeb\xa4\xee\xc9\x93\x99\x21\x74\xdc\x44\x21\x70\xe3\xe9\xed\x48\xb6\x9a\x6f\x9f\x3f\x22\x59\xc0\xcb\x06\x10\x54\xf7\xf4\xfc\x04\x36\x69\x5d\xe7\x65\x8b\xd1\xf8\x52\xf6\x32\x5b\x95\x2d\x1c\xe1\xcd\x18\x4e\x02\x73\xa9\xbf\x5c\xec\x68\x49\x48\xd2\xd3\x5e\x5f\xfa\x69\xdd\x79\x54\x0e\x85\x00\x24\xb2\xc4\x61\x99\x6e\xef\x17\x57\x95\x7c\x70\x61\x81\x34\xf1\xd0\xc3\x30\xfa\xaa\xc3\xe5\x94\x6c\x64\x4f\xfa\xd6\x51\x6f\x93\x86\x4e\x1d\x1b\x4a\xd5\x64\xdb\x16\xd9\x79\x5a\x13\x8e\xc4\x4b\x84\x94\xe7\x65\xcd\xaa\x8c\xae\x2b\x4e\xf0\xaf\xa5\xe2\x08\x35\x87\x88\xb6\x92\x77\xf0\x27\xdb\x6a\x81\xd2\x93\x61\xc1\xde\x42\x18\xcb\x68\x08\x64\xd5\xb4\x11\x6a\x27\x93\xb1\x74\x07\xd2\x52\x94\x64\xdc\x94\xb8\xa2\xf2\x10\x4e\xb8\x1e\x06\xaa\xb8\xe1\x2a\x61\x59\x64\x1f\xdf\x8f\x77\x0d\x95\xae\x07\x4f\x21\xf0\x61\x89\xec\x7e\x44\xa8\xab\x95\x20\xb5\x58\xa6\x42\x5d\x4b\xb7\x0b\xfe\x63\xdc\xb9\x66\x27\x15\x4b\xac\xc1\x04\x97\x5c\xf0\x84\xda\xbf\x9f\x8a\x7a\xab\xe0\xa7\x95\x91\x65\x79\xaa\x3a\xf4\xf1\x51\x96\x7d\xf5\x55\x21\xe4\x7b\x71\xec\xe0\x17\xce\x7f\x46\x33\xe4\xcf\xf2\x77\xef\xb2\xec\x57\x42\x0b\x5d\xda\xd6\xc1\x79\xa2\xb5\xa7\x61\xda\xfa\x0e\xcb\xd7\x2f\xb6\xef\x25\x54\x3b\x0a\xaa\x34\x1d\x6a\x1d\xe7\x05\xad\x1e\x87\x41\x0a\x53\xb3\xa7\x2a\x1a\x6d\xf2\x94\xcc\x42\x2c\x96\xa1\xab\xdb\xa6\x54\x84\x03\xb4\x64\xc9\xa3\x11\x0a\x29\xaf\x24\x06\x1c\x06\xc5\x4e\x2e\x1b\x03\x6c\x43\xf4\x63\x4f\xb2\xd2\x70\x1a\x09\xc2\xa9\x23\xd9\xd5\x7d\x96\x1d\xc5\x34\xde\xc1\x68\x0f\xab\x90\x89\x5f\xef\x50\x96\x12\x85\x3f\x6d\xcf\x29\x54\xa1\x8b\xa6\x16\x3a\x1e\x1e\xc1\x77\x5f\xff\x69\x21\x41\x58\x97\xe7\x3b\x72\x3b\x5a\xdd\xad\x65\xa3\x92\x80\xbf\xfb\xfa\xcf\xb0\x1c\x14\xf0\x85\x30\x82\xfb\xc1\xf9\x88\x56\x97\x39\xed\x15\xae\x6b\x43\xf0\xe4\x7a\xd7\x7f\xf4\xf4\xc9\x5e\xff\x6c\xf5\xe3\xc3\xf4\xf1\xe1\xd3\xe4\x4b\x8e\xce\xd3\xd1\xf9\xd3\xc4\xd6\xdb\x6a\xca\x78\x91\xae\x14\x8f\xe5\x22\x8f\x86\x0e\x64\xb4\x43\x53\x8f\x62\xd5\x31\x1d\x6e\xf5\x2f\x91\xd6\x59\x93\x7e\x24\xb4\x0b\x15\x05\xbd\xca\xd9\x03\xcd\x02\x83\x12\x21\x72\xcf\xb6\xdd\xd7\xa3\x5f\x60\xc7\x55\xc0\xe9\x7a\xa0\x4a\x36\x85\x96\x0f\x64\x17\xe8\x70\xa5\xc5\x71\x54\xe4\x27\xdb\xe3\x2d\x19\xa1\x9c\x45\x0d\x8e\xa3\x75\x42\x1b\xd3\x2f\x00\x91\xd5\xd4\xad\xa9\xd4\x2a\xbd\xfa\x43\x75\x51\x51\x59\xe2\xb8\xba\x4a\xbf\x12\x68\xbe\xab\xe4\xda\xba\x3f\xc6\xfe\x4f\xe6\xcf\x68\x16\xe2\xaf\x1b\xca\xd1\x46\x3f\x7d\xce\x21\xbe\x7b\x97\xfd\x3f\x00\x00\xff\xff\xca\xef\xbf\x07\x13\x0f\x00\x00"),
		},
		"/pages/index.md": &vfsgen۰CompressedFileInfo{
			name:             "index.md",
			modTime:          mustUnmarshalTextTime("2019-05-06T04:58:09.409866327Z"),
			uncompressedSize: 2452,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x96\xcb\x6f\x14\x47\x10\xc6\xef\xf3\x57\x54\xd6\x87\xf5\x4a\xf6\xae\x0d\xe1\x40\xcc\x82\x20\x98\x47\x04\x11\xc2\x09\x02\x21\x0e\x35\x33\x35\xd3\xe5\xed\xee\x9a\xf4\x63\xc7\xb3\x88\xfc\xed\x51\xf7\x8c\xbd\x0e\x8a\x72\x42\xf2\xcd\x8f\xae\xea\x9a\xef\xf7\x7d\xdd\x7d\x00\x5e\xb3\xf1\xd8\xce\x3d\x98\xe8\xb9\x02\x47\x5e\xa2\xab\xc8\x17\xc5\x2b\x72\x04\x83\xc4\xb9\xd6\xd0\xb0\xad\xa7\x15\x68\x6b\x68\xc9\x92\x43\x0d\x6c\x1b\x71\x06\x03\x8b\x05\x2c\x25\x06\xf8\x55\xb1\x25\x4f\xc0\xd6\x07\x17\x0d\xd9\xe0\x8f\xc0\x88\x0f\x7a\x80\xa0\x08\x9e\xf3\x8e\xa1\xd1\x31\xd0\xb2\x28\x5e\x37\xa9\x3f\xd4\x62\xe7\x01\x36\x56\x7a\x50\xd2\x43\x10\x70\x84\xf5\x4d\x2b\xcd\x1b\x02\x43\x47\xa9\xde\x13\xa0\x23\xe8\x9c\x94\x58\x4e\x2d\x4b\xf2\x61\x3f\xf7\xed\x89\xd3\x16\x73\x03\xa8\x7b\x1c\x3c\x28\xec\xba\x21\x75\xaf\x14\x86\xff\x1b\xd7\x0b\xf0\x38\x59\x8f\x36\xa4\x8a\xcb\xe8\x03\x78\x25\x12\xc0\x10\x20\x18\xf2\x1e\x5b\x82\xc3\xaf\x5f\x97\x6f\x87\x73\x83\xac\xbf\x7d\x5b\xfc\x54\x14\x07\x07\x70\x21\x4d\xe8\xd3\x90\x35\x6d\x49\x4b\x97\x9a\xc2\x0a\x2e\xa5\x04\x69\x1a\x72\xfe\xe6\xbb\xd3\x22\x95\x44\x6e\xc4\x81\x19\xa0\x17\xb7\x01\xf4\x80\xe0\xbf\xeb\x41\x2e\x7d\x3d\x7b\xe8\xa9\xf4\x1c\x08\xd8\x27\xc9\x52\xdd\x20\x71\x09\x17\x44\xa9\xc1\xe7\x97\x1c\x5e\xc5\xf2\xcb\xa1\x0a\xa1\xf3\xbf\xac\x56\x2d\x07\x15\xcb\x65\x25\x66\x35\x81\x5e\x80\x38\xf8\xfc\x86\xed\x86\xea\xd7\x76\xbf\xb2\xef\xfb\xa5\xce\x7f\x65\x9b\xd7\xb3\x5d\xf9\x40\x9d\x22\x7b\xdc\xc6\x40\x1b\xb4\x3e\x1c\x9f\x9e\xde\x7f\xf8\x00\x4f\x4f\x4f\x56\x8b\xc4\xa0\x61\x4d\x3e\x6b\x47\x98\xc4\x3e\x38\x18\xf9\x1e\x5e\xcb\xfa\x0c\x4d\x29\x02\x2f\x12\xef\x45\x51\x3c\xe2\xc6\xa1\x21\xe8\xb9\x0e\x6a\x3d\xbb\xff\xc0\x91\x99\x81\x22\x6e\x55\x58\xcf\xee\x9d\xe4\x5f\xbd\xab\xd6\xb3\xdb\x63\x0d\x12\x43\x2c\x29\x4f\x45\xa6\xa4\x7a\x75\xfe\xfe\xd3\x60\x2e\x5f\x7c\x24\x7e\xfd\xc4\x07\x74\x61\xfd\xf3\xfd\x19\xe4\xde\xa5\xb8\x9a\xdc\x7a\x76\x32\x03\xd4\x5a\xfa\xf5\x0c\xab\x8a\x34\x39\x31\x14\xc8\x9d\x01\xc6\x20\x9d\xc6\xe1\x0c\xc8\x56\x6e\xe8\x02\xd5\xc7\x86\x6a\xc6\x33\x68\x07\x27\xbe\x92\x8e\xce\xa0\xe3\x2a\x44\x47\xc7\x6c\x8f\xa7\x1f\xa7\x7e\x4d\xd4\xda\x57\x8e\xc8\x3e\x7e\xb4\x1a\x3f\xe7\x71\x51\xfc\x91\xd8\x70\x22\x57\x72\x18\x0d\x2b\x41\x91\xbb\x16\x20\x1b\xde\x03\x5d\x55\xd4\x05\xe0\x00\x2a\x63\xde\x92\x4b\x16\x66\x0b\x1d\x76\xe4\xc0\x90\x29\x1d\x5a\x82\x4a\xb6\xe4\xd8\xb6\x20\x96\x40\x89\xa6\x23\xe8\x15\x57\x0a\x2a\x8c\x3e\x49\x9e\x5d\xd9\x39\xa9\x63\x95\xfc\x58\xc6\xdd\x2e\x2d\xb7\xc2\x9e\xe0\xd0\x4b\xb4\xb5\x07\x2f\x86\xfa\x6c\x30\xb6\x25\x85\x9e\xc8\x02\x82\xa3\x2a\x6b\x94\x83\x8c\xb0\xc1\x9d\xc8\x11\xd4\xd4\x91\xad\xc7\x2d\xc7\x14\x62\xb4\x21\xdb\xb4\xd5\x91\x72\xd0\x6e\xc6\x13\xbb\x58\x16\xc5\xe7\x0f\x4c\xfd\x08\xfc\x26\x7d\x5f\x0e\x9f\x74\xeb\x9a\x77\xbc\x62\x5b\xd3\xd5\x22\x7b\xe2\xdc\xa9\xb8\xf7\xc4\x96\x45\xb3\xfd\x81\x6e\x28\xab\xe7\xdb\x4b\xf5\xf0\xaa\xfc\x28\x93\x1b\xee\xd6\x0c\xcf\xd0\x73\x85\x5a\x0f\x80\x37\xe7\xcb\x96\x9c\x4f\x87\xa4\x34\x59\xc9\x0f\x59\x83\xa3\x5b\x86\x10\xab\x87\xec\x8a\x44\xdc\x87\x0c\x3f\x01\x62\x0f\x5d\xf4\x8a\xea\xdb\xde\x42\x68\x23\x07\x74\xff\x36\xc1\x5f\x31\x9d\x0c\xa5\x13\xac\xc1\xa1\x6d\x29\xed\x36\x3a\x62\xfc\x57\xae\x4d\xdb\xab\x68\xd0\xc2\x56\xb8\xa2\x31\xb5\x2f\xe3\x4e\x91\x6d\xf7\x90\x76\x9c\xec\xfb\x03\x21\x5d\x5d\xe8\xa7\x6f\x1f\xd4\x7f\xfa\xdf\xe4\x4e\xe1\xbc\x19\xf5\x33\x7c\x05\xfb\x48\x74\x8c\x56\xa6\x3c\x4c\xca\xae\x17\xf0\x49\x22\x74\x3a\x56\x1b\xf8\x5e\xf7\x23\x28\xe3\x4d\x8c\xff\xbe\x77\x3a\x01\xf3\xf9\xe6\xb8\x0e\xb7\x89\x95\x02\x8d\xae\x25\x37\xe1\x08\x0a\xed\x7e\x87\x0a\x2d\x74\xe4\xd2\x05\x3a\x42\x78\xc7\x1d\xee\x09\xcc\xc7\x65\xf3\x1f\xc8\xe0\xed\xbb\x37\xe1\xe4\xc5\xd3\xfa\x3d\x9f\x4f\x41\x39\x7d\x78\xa7\x30\x2e\xc4\x85\x64\xd2\xef\xc4\x9d\x52\x91\x9c\x1a\xa4\x83\xc6\x51\xf0\xf9\xb2\x9c\xde\x11\xd1\x53\x9d\xef\xbe\xca\x11\x86\x9c\x14\x98\xf7\xb8\x1d\xe6\x93\xdd\x13\xc9\x54\x7d\x8d\x25\xd7\x26\x1c\x41\x62\x9b\x8e\xe5\x4e\x0d\xd7\x19\x1d\x91\x4d\x29\x4a\xdc\x14\x8f\x4b\x38\xa4\x82\xdb\xc8\x96\xf0\xbb\xa4\x97\x46\xae\x9b\xe7\x07\x83\xa1\x7a\xfe\x5f\xde\x70\x98\x8f\x7f\x87\x1d\xd7\x7a\x18\x6d\x44\x35\xf4\x1c\x14\x84\x5e\xd2\x25\xdc\xf0\x96\xd2\x2b\xa5\x25\xe7\x21\x3d\x48\x20\xb0\xa1\x65\xf1\x4f\x00\x00\x00\xff\xff\xdf\xf1\xc3\xc7\x94\x09\x00\x00"),
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
