// Code generated by vfsgen; DO NOT EDIT.

// +build !dev

package api

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

// Api statically implements the virtual filesystem provided to vfsgen.
var Api = func() http.FileSystem {
	fs := vfsgen۰FS{
		"/": &vfsgen۰DirInfo{
			name:    "/",
			modTime: time.Time{},
		},
		"/api.proto": &vfsgen۰CompressedFileInfo{
			name:             "api.proto",
			modTime:          time.Time{},
			uncompressedSize: 10486,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x5a\xdf\x6f\xdb\x38\xf2\x7f\xf7\x5f\x31\xf0\xcb\x37\xfd\xa2\xb5\xdb\x74\x77\xaf\x48\x2e\x77\x97\x73\xb2\xad\xd1\xd6\x09\x22\xb7\x8b\x7d\x32\x68\x69\x2c\xf3\x22\x91\x5a\x92\x8a\x6b\x2c\xfa\xbf\x1f\xf8\x4b\x26\x65\xc9\x4e\xdb\x2c\x70\x7e\xd8\x8d\xc4\x99\xe1\x7c\x66\x86\xf3\x83\xea\x78\x0c\x13\x5e\x6d\x05\xcd\xd7\x0a\x4e\x5f\xbe\x7a\x03\x09\x29\x65\xcd\x72\x48\xae\x12\x98\x14\xbc\xce\x60\x46\x14\x7d\x40\x98\xf0\xb2\xaa\x15\x65\x39\xcc\x91\x94\x40\x6a\xb5\xe6\x42\x8e\x06\xe3\xf1\x60\x3c\x86\x0f\x34\x45\x26\x31\x83\x9a\x65\x28\x40\xad\x11\x2e\x2b\x92\xae\xd1\xaf\x3c\x87\xcf\x28\x24\xe5\x0c\x4e\x47\x2f\xe1\x44\x13\x0c\xdd\xd2\xf0\xd9\xb9\x16\xb1\xe5\x35\x94\x64\x0b\x8c\x2b\xa8\x25\x82\x5a\x53\x09\x2b\x5a\x20\xe0\x97\x14\x2b\x05\x94\x41\xca\xcb\xaa\xa0\x84\xa5\x08\x1b\xaa\xd6\x66\x1f\x27\x45\x6b\x02\xbf\x3b\x19\x7c\xa9\x08\x65\x40\x20\xe5\xd5\x16\xf8\x2a\x24\x04\xa2\x9c\xd2\xfa\xb7\x56\xaa\x3a\x1b\x8f\x37\x9b\xcd\x88\x18\x85\x47\x5c\xe4\xe3\xc2\x92\xca\xf1\x87\xe9\xe4\x7a\x96\x5c\xbf\x38\x1d\xbd\x74\x4c\x9f\x58\x81\x52\x82\xc0\x3f\x6a\x2a\x30\x83\xe5\x16\x48\x55\x15\x34\x25\xcb\x02\xa1\x20\x1b\xe0\x02\x48\x2e\x10\x33\x50\x5c\x2b\xbd\x11\x54\xdb\xed\x39\x48\xbe\x52\x1b\x22\x50\x8b\xc9\xa8\x54\x82\x2e\x6b\x15\xd9\xcc\xab\x48\x65\x44\xc0\x19\x10\x06\xc3\xcb\x04\xa6\xc9\x10\xfe\x7d\x99\x4c\x93\xe7\x5a\xc8\x6f\xd3\xf9\xbb\x9b\x4f\x73\xf8\xed\xf2\xee\xee\x72\x36\x9f\x5e\x27\x70\x73\x07\x93\x9b\xd9\xd5\x74\x3e\xbd\x99\x25\x70\xf3\x2b\x5c\xce\x7e\x87\xf7\xd3\xd9\xd5\x73\x40\xaa\xd6\x28\x00\xbf\x54\x42\x23\xe0\x02\xa8\xb6\x26\x66\xc6\x74\x09\x62\xa4\xc2\x8a\x5b\x95\x64\x85\x29\x5d\xd1\x14\x0a\xc2\xf2\x9a\xe4\x08\x39\x7f\x40\xc1\x74\x24\x54\x28\x4a\x2a\xb5\x57\x25\x10\x96\x69\x31\x05\x2d\xa9\x22\xca\xbc\xda\xc3\x35\x1a\x98\x9d\x5c\x88\x4d\x66\x93\x39\xfc\x5d\xda\xa7\x51\xaa\x83\x8d\x99\x58\xfb\x57\x5e\x12\x5a\x8c\x52\x5e\xfe\x63\x30\x90\x5b\xa6\xc8\x17\xb8\x80\x61\x25\xb8\xe2\xaf\x87\xe7\x83\x41\x45\xd2\x7b\xad\x49\x5a\x12\x29\xd7\xe7\x83\x01\x2d\x2b\x2e\x14\x0c\x73\xce\xf3\x02\xc7\xa4\xa2\x63\xc2\x18\x77\x8a\x8c\x0c\xe7\xf0\xbc\x21\x33\xcf\xe9\x8b\x1c\xd9\x0b\xb9\x21\x79\x8e\x62\xcc\x2b\x43\xda\xc9\x36\x18\xd8\x55\x38\xc9\x45\x95\x8e\x72\xa2\x70\x43\xb6\x76\x39\x5d\xe4\xc8\x16\x4e\xca\xc8\x49\x19\xf1\x0a\x19\xa9\xe8\xc3\xa9\x5f\x79\x06\x17\xf0\xe7\x00\x80\xb2\x15\x3f\x33\x7f\x01\x28\xaa\x0a\x3c\x83\xe1\xa4\xa8\xa5\x42\x01\x1f\x09\x23\x39\x0a\xb8\xbc\x9d\x42\x92\xbc\x83\x4a\xf0\x07\x9a\xa1\x18\x9e\x1b\xf2\x07\x7b\x7e\xce\x60\xf8\xf0\x72\xf4\x6a\xf4\xd2\xbd\x4e\x39\x53\x24\x55\x5e\xa8\xfe\x31\x52\x6a\xb9\xa1\x9d\x1d\xb1\xfe\xd5\xa2\x38\x83\xa1\x8e\x7b\x79\x36\x1e\xe7\x54\xad\xeb\xa5\xb6\xf5\xd8\x79\xe2\x45\xca\x52\x35\x4e\x4b\xf2\x42\xca\x75\xc0\x87\xda\x29\x67\x30\x3c\xe8\x30\x47\xff\x55\xff\xcf\xfc\x07\xbf\x28\x14\x8c\x14\x8b\x8c\xa7\xd2\x2b\xf9\x3d\x2a\x64\x28\x53\x41\x8d\x7d\xcf\x60\xf8\x91\x0b\x04\xb2\xe4\xb5\x82\x47\x99\xef\xeb\x00\x40\xa6\x6b\x2c\x51\x9e\xc1\xbb\xf9\xfc\x36\x39\x6f\xbf\xd1\x2f\x52\xce\x64\x6d\xde\x0c\xdd\xa1\xd6\xfb\x8d\xff\x23\x39\x33\x62\x2a\xc1\xb3\x3a\xed\x5b\xff\x7a\x3e\x18\x48\x14\x0f\x34\xc5\x46\x2b\x0b\x58\x9f\x55\x5a\x14\x56\x27\x93\x05\x09\xa4\x96\xc2\xac\x8b\x2a\x85\x89\x40\xa2\xd0\xf3\x9d\x44\x8f\x1f\x65\xfe\x0c\x04\xaa\x5a\x30\xd9\x5a\xba\xc3\xaa\xd8\x3e\x0b\xbc\xdf\xc4\xaa\x39\x0b\x23\x52\xd1\x91\xb6\xb4\x8f\xc0\xdd\xaf\xe2\x52\xc1\x19\x0c\xcd\x71\x79\x78\x35\x76\x0a\x0d\x23\xa2\x25\xcf\xb6\x9a\xe8\xff\x77\xaf\xbf\x3a\x1f\x47\xc8\x04\x2a\x41\xf1\xc1\xe6\x10\xa9\x88\xaa\xa5\xce\xbb\x0d\x4c\x9d\x1f\x80\x2a\x09\xf7\xf5\x12\x53\xce\x56\x34\x37\x29\x26\xe5\x8c\x61\xaa\xe8\x03\x55\xdb\xc6\x14\x6f\x51\x35\x76\xd8\xfd\x1d\x1b\x61\xf7\xfe\xfb\x2d\x90\xe3\x61\x03\x74\x22\xcd\xb0\x40\x85\x1d\x0e\xbc\x32\x0b\x8d\xe2\xd1\x63\xac\x7b\xb4\xf4\xfd\xea\x3b\x4d\xbe\x19\x41\xe3\x2b\x02\x05\x95\x4a\xfb\xc9\x31\xca\x0e\x17\x7c\xd0\x24\x27\xf1\x73\x9f\x2b\xf4\xda\x53\xbb\x63\xac\x75\x3c\x82\x48\x73\xfa\x40\x63\x3c\x43\xe9\x43\x50\x87\x18\xd9\x1d\x3b\xcc\xf6\xbc\xb6\x53\x7e\xa6\x19\x13\xcb\x77\xd2\xf9\xba\x0f\x76\x40\xf2\xe4\xe8\x0d\x1c\x8b\xe6\xb8\x5b\x6b\xc1\x7c\x9d\x30\xa5\x46\x94\x26\x41\xb9\x4c\x49\x2a\x0a\x3a\x3f\xc5\xe8\x5d\x5f\x36\x0d\xc8\x4f\x76\xaf\xf7\x20\xbb\xf7\x4f\x86\xd3\xa9\x7b\x04\x1b\xc9\x32\xe3\x58\xa8\x38\x2f\x74\x5f\x75\xd8\xa9\x97\x59\xa6\x7d\x72\xab\x89\x4f\x82\x87\x18\x4d\xb0\xf0\xe4\x59\x74\xac\x15\xfd\xbe\x54\xda\x24\x98\x1d\xe0\x95\xe0\xe5\x11\xc8\x36\xa7\xec\x50\xc7\xcf\x5d\xf9\xe7\xc7\xb1\xf7\x26\xa0\x16\xfa\x4e\x98\x32\x25\x85\x2d\x17\xac\x2e\x97\x28\x74\x1a\x2a\x49\xba\xa6\x0c\x25\x98\xe6\x3d\xc0\x7f\xf4\x18\x27\x5a\xda\x0e\x7d\xf4\x18\x83\x8f\x96\x7e\xc0\xef\xf5\x13\xbb\xdd\x1d\xdf\xba\xca\x05\xc9\xd0\x29\xe2\x33\x58\x4e\x1f\x90\x75\xe5\xae\x4f\x96\xdc\x25\xa2\xf6\x21\xee\x5d\xdd\x3b\xd6\xbd\x94\x4f\x9e\xd0\x1c\xc0\x63\x07\x5e\x29\x2c\x2b\xa5\x8f\xba\xb7\xc8\x7e\xc5\x8d\x95\x86\x93\xf8\x39\xc6\x18\xaf\x3d\xb9\xdf\xf7\x50\x1d\x73\xfd\xd7\xc1\x00\x59\x5d\xfa\x46\xd1\xd5\x9d\xa6\x5d\x9c\x71\x05\x12\x95\x79\x4c\xe6\x97\xf3\x4f\xc9\xe2\xd3\x2c\xb9\xbd\x9e\x4c\x7f\x9d\x5e\x5f\xc1\x05\xbc\x3c\xf7\xa4\xf3\x35\xc2\xed\xdd\xcd\xe7\x69\x32\xbd\x99\x4d\x67\x6f\x4d\xed\x43\xa0\x2c\xd3\xfd\x29\x4a\x73\xc6\x7c\x75\xa4\x12\x96\xa8\x67\xb7\xd4\x34\x91\xd9\xc8\x48\x89\xd8\x2f\xe0\x55\x24\xfb\xee\xd3\xec\xa8\xd8\x35\xd1\x72\x75\x88\x5a\xb1\xb6\xdd\x93\xb0\xaa\x8b\x62\x0b\xb5\xd4\xc3\xb1\xdd\xca\x4b\xbb\x80\xd3\x78\x97\xeb\xc9\xcd\x6c\x32\xfd\xd0\xbd\x13\x51\x20\x79\x89\xb0\xe1\xe2\x5e\xcb\x25\xba\x65\xc4\x62\xeb\xc0\x64\x9c\xa1\x9e\x92\x03\x95\x9e\x83\xac\xd3\x35\x10\xe9\xe2\x47\x93\xe9\xe5\x92\x18\x85\xb9\x6d\x14\x9a\x99\xdc\x29\x17\x28\x71\x01\xaf\x23\x05\x93\xf9\xcd\xed\xed\xa3\xcd\x6b\x53\x63\xe6\xfc\xe7\x38\x2f\xe0\xa7\x48\xe4\xf5\xdd\xdd\xcd\xdd\x41\x79\x25\xd1\x10\xa1\x66\xd6\x84\x86\xd9\x72\x5d\xc0\xcf\x91\xac\xab\xeb\xb7\x77\x97\x57\xd7\x57\x07\xc5\xb9\x5b\x0b\x09\xb5\xd4\x2d\x79\x6a\x62\x5e\x71\x10\x28\x95\x9e\xa8\xb4\xbb\x60\x55\x33\xb3\x40\x0a\xdf\x93\x37\xb2\x2f\xe0\x97\x73\x1d\xb9\x25\x4a\xa9\x07\xf0\xf6\x90\x12\xc4\x2f\x29\xd1\x5f\xbc\xf8\xdd\x15\xd7\x58\x82\x2c\x6e\x88\xa5\x12\xda\x60\x7a\x68\xdd\x0b\x3d\xdf\xcb\xf0\x15\xbc\xaf\x97\x28\x18\x6a\x44\x3a\x25\xea\x40\xf0\xcd\xde\x08\x26\x9c\x29\xc1\x0b\xa8\x0a\xc2\x1a\x2e\x09\x44\x20\x64\xa8\x50\x94\x94\xd9\xab\x1a\xad\xce\x47\x5b\x64\x92\x0a\xd3\x51\xa8\xc1\xfd\x1b\xb9\xf0\x1b\x86\xd1\xf9\xd1\x17\xa5\xcd\x9a\xa6\x6b\x73\x07\x25\xa8\xc4\x08\x5a\x1a\x2a\x60\x18\x9d\x4a\xb7\xfa\x45\xb0\xa3\xa7\x5c\x18\xca\x85\x6d\x56\xc3\x50\x7b\xc4\x6e\x36\xff\x61\x65\x4f\x5a\x28\xdc\x5a\xc5\x48\x5d\xe8\x4a\x24\x6d\xc8\xf5\x79\xcc\x64\xc1\x60\x44\x5d\xa3\xb9\x21\x32\x67\x43\x45\xf8\x36\x44\x46\xe5\xd7\x98\x92\xda\x6b\x30\x94\x36\x4d\x2d\x75\xe5\xe5\xf7\x7b\x4e\xcc\x50\x11\x5a\xc8\x76\x34\x38\x56\x1d\x7b\x15\x67\xd2\x59\xcd\xd5\x1f\x85\x65\x43\x68\x7c\x11\x40\x88\x26\xc2\xc7\x44\x5c\xc1\xf9\x3d\x66\x50\x57\xdd\xf1\xd6\x29\xba\x65\x9a\x69\xeb\x98\xdb\x54\x23\xb7\x52\x61\xb9\x0f\x3e\x84\x72\x65\xd0\x1f\x04\xd4\x9e\x14\x43\x8f\x10\xa5\x73\x4a\xb0\xf7\xff\x49\xab\xba\xe2\x90\xa1\x54\x82\x6f\x8f\xa2\xda\x1f\x37\x77\x3b\x4c\x78\x5d\x64\x11\xb6\x25\x7a\xc1\xee\x80\x76\xf9\x35\x69\x26\x7c\xcd\x1a\x46\x81\x53\xc4\xcd\x5f\xfd\xbe\x73\x63\x24\xfc\xd9\xbf\xfc\x43\x3e\x70\x4c\x1f\x3a\x07\x5c\x7f\x76\x3a\xc2\x6d\x5f\xe7\x90\xe8\x50\xb4\x1d\xcc\x66\x97\x59\x46\x6d\x52\xed\x18\xcc\xe2\x3b\x93\x1e\x91\x96\x60\xe1\xb5\x6a\xd7\xcf\x7e\xfe\xb8\xc3\x68\x1c\xf3\xba\x0b\x64\x10\xad\xff\x9b\x50\xc3\x13\x11\x5c\x25\x29\xee\x6f\x92\xf4\x9f\x3d\x62\x03\xfa\x76\x71\xff\x66\xeb\xc5\x59\x75\x57\x9c\x3e\x90\x25\x16\x3b\xdb\x69\xd9\xcc\xd9\x8f\x40\xa1\x17\x8f\x17\x3d\x52\xd4\x7d\x0c\x76\xcd\x47\xa8\x57\xde\x5d\xd1\x5b\x3b\xdb\x81\x41\xa2\x0d\xfb\xa8\x30\xfa\xf1\x6a\xe7\xf5\x9e\x22\x15\xe9\x6f\x94\x90\xcd\x07\x81\x1e\x91\xd1\xb9\x6a\xdb\xc3\x89\x88\x90\x6e\x2b\x8c\x46\x3e\xc5\x83\xeb\xd1\x13\xa9\x08\xcb\x88\xc8\x74\x31\xca\xab\xfa\x59\x68\x04\xca\xf4\x6a\x8a\x46\x44\xfb\x1c\xec\x0f\x93\x66\x95\x32\xf5\xfa\x14\x52\x5e\x33\xd5\xc4\xfe\x71\xf3\xed\x19\xac\xd7\x48\x2c\x38\x24\x8e\xab\xe9\xdf\x0f\x39\xbb\x65\xdc\x36\xeb\x71\x8b\x9e\xfe\x05\x16\x7d\xfd\xed\x16\xfd\xc9\x5b\xf4\x2d\xaa\xb0\x6d\x33\xb7\xf1\xf6\xce\x29\x48\xf2\xbb\xcb\x25\x9b\xff\xc7\x63\xb0\xc9\x5e\xdb\xc1\x73\xfb\xaa\xb2\xcf\xd7\x2e\x0c\x2b\xe0\x15\x0a\xeb\x3f\xdd\xa9\xdc\xbc\xef\xa9\xc9\x5e\x54\xc7\x9d\xd7\x6e\xca\x73\xb0\x15\xc9\xfd\x48\x91\x53\xdd\xa6\x54\x5c\x52\xc5\xc5\xb6\x21\x74\xc6\xcb\xa9\x0a\xfa\xc6\x57\xe7\x6d\x41\x6b\x22\xd7\x3e\x34\xb4\xa4\x94\x97\x25\x55\x5d\x52\xec\xca\xce\xa9\xfd\x7d\x99\x12\x88\x06\x6a\x5a\x20\x61\xb0\x59\x23\x83\x65\x4d\x8b\x4e\xb1\x9a\x78\x61\xc7\x82\xc6\xb7\x4e\xf4\x95\x7e\xc9\x57\x86\x37\x6b\xf3\x9a\x97\x8b\xcc\xf2\xfd\x14\xf1\x7d\xde\x79\x38\xe7\x7a\x9a\xc8\x6c\x0a\x2e\x2b\xea\xa6\x94\x50\x07\x1e\xd8\xe7\xe7\x48\xce\xc4\x72\x08\x23\xa2\xcd\x97\xfa\x45\x33\x76\x04\x5c\xb7\x05\x51\xda\x73\x40\x95\x35\x82\x25\xcc\x4c\xf8\x8c\x41\xd4\xcc\x7c\x9e\xe4\xac\x2d\xb1\xf2\x8c\x17\xf0\x37\x3f\x88\x0f\x5a\x90\x82\xa0\x30\x4b\x1d\xb1\xe2\xd0\x2c\xc2\xf2\xd6\xd1\xe9\x1c\xba\x7d\x39\xd8\xe3\xd9\x81\x77\x83\x66\x80\x49\x39\x93\x34\x43\xa3\xbf\xc6\xe7\x6e\x1a\x1e\xd3\xcb\x1e\xbe\xd4\x09\x9a\x40\xc2\xda\x2d\xa0\xdb\xa5\xbf\x03\x34\x6a\x47\x93\x56\xc5\xa5\xa4\xcb\x02\xc1\xfe\x73\x02\xc6\x37\x71\x0a\xf3\x25\xcc\xf3\xb4\x2d\xb6\x77\x7b\xf3\x17\xd9\xa8\x03\x80\x11\xb2\xc1\xf0\xaa\x80\xff\x33\xaa\xbb\xe1\x5c\xd8\xab\x73\x7b\x9e\x22\xd2\x4e\x49\x04\x64\x9d\xa6\x28\xe5\xaa\x2e\xfa\x47\xa6\x40\x6c\x7c\x65\x7d\xc4\x0e\x3c\xbe\x1d\xd7\x79\x3f\x54\xdd\xd1\xcd\x3a\xf1\x7f\x0c\x6a\x85\x96\xd2\xd1\x3c\x3d\x7a\xd0\x3c\xed\x83\x70\x7c\xcc\xdc\xdd\xf4\x7e\xf3\xa0\xb9\x37\xe7\x7c\x8b\xe1\xdc\x05\xf6\xce\x76\x8f\x36\x9c\x93\xb5\x53\x5c\xc7\x97\xdc\xc9\xec\x0c\xfd\xc6\x5c\x0b\x4b\xdd\x3f\x1d\xc5\x1f\x9d\xba\x71\x84\x2d\xc7\x64\x87\xe9\x8f\x1a\xc5\xf6\x20\x8e\xa6\x50\xf7\x7c\xbe\xf2\x1b\xf8\xc9\x5c\x4b\x7d\x8b\xca\x1b\x56\x33\x73\x11\x5c\x51\xb9\x10\xb2\xbd\xf1\x61\x30\xad\x50\x68\x77\x4e\x1d\x4d\xfb\x9e\xf9\x27\xa6\xe3\x68\x7d\x2c\x08\x19\xe3\xc6\xe4\xf4\x7c\xd0\xdd\xe8\x13\x2f\x20\xea\x0c\x7c\x90\x87\xd7\xb1\xe0\xaf\x64\x33\x84\xfb\x37\x0d\xd0\x56\x6d\xb9\x7f\x23\x35\x45\xe2\x27\x84\xb8\x11\x08\xae\xaf\x5c\x41\xe9\xe0\xff\xdc\xbe\x7d\x02\x7f\x27\x44\x12\xbb\xbd\x9d\x76\x17\x74\xaf\x56\x96\x84\xc8\xc4\x2c\x4e\xb3\xbd\x6a\xbd\xe3\x5f\x73\xa9\xb4\xc1\xbb\xd8\xdf\xb9\xb5\xbd\x22\x6d\xd8\xed\x6d\x69\x27\x72\xcd\x1c\x41\xff\x65\x9f\x7d\x7a\xab\x33\x82\x40\xd9\xc9\x3d\xbd\xd5\x8b\x5d\x55\xf9\x2d\x2a\xd9\x7c\xa5\xb6\xb7\x65\x74\xef\x5f\x4f\xb4\x33\x94\xd5\xa3\x89\x8f\xf6\xac\xdb\xf1\xf9\xeb\x29\x92\x76\xfb\x9b\xd3\xf1\x53\x1b\x64\x22\xfb\x35\x2c\xf8\xe6\xf5\xe8\x4c\x14\x25\x7e\x2b\x27\xb6\x4a\xa4\x97\xc9\xde\x07\xd2\xf6\x3e\x71\x37\x8a\x28\xef\xc5\x1b\xf7\x15\xdc\xd9\xfe\x67\xbe\x88\x6f\xef\xdc\xf6\xa9\xf5\xa3\x1e\xfb\x6f\x00\x00\x00\xff\xff\xa8\x0c\x79\x3b\xf6\x28\x00\x00"),
		},
	}
	fs["/"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/api.proto"].(os.FileInfo),
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
			gr:                        gr,
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
