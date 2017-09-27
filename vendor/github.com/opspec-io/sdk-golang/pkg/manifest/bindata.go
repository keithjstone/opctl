// Code generated by go-bindata.
// sources:
// github.com/opspec-io/spec/spec/pkg-manifest.schema.json
// DO NOT EDIT!

package manifest

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _githubComOpspecIoSpecSpecPkgManifestSchemaJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5d\x6d\x6f\xdc\x36\x12\xfe\xee\x5f\x41\x6c\x0d\x74\x7d\xd9\x17\xe7\xa5\xb9\xab\x8b\xc2\x30\xd2\x04\x97\x43\x9b\x04\x4d\x93\x03\x6a\x6f\x0a\xae\x34\xbb\xcb\x5a\x12\x55\x92\xb2\x93\xf6\xf2\xdf\x0f\x24\xb5\x7a\xa1\x48\xad\x24\x4b\xf5\x26\xb1\x3f\xb4\xb1\xc8\x19\x71\x1e\x3e\x1a\x0e\x67\x28\xf9\xaf\x03\x84\x46\x87\xdc\xdb\x40\x88\x47\x27\x68\xb4\x11\x22\x3e\x99\xcf\x7f\xe7\x34\x9a\xea\xab\x33\xca\xd6\x73\x9f\xe1\x95\x98\x1e\x3f\x9a\xeb\x6b\x5f\x8d\x26\x52\x4e\x10\x11\x80\x94\x8a\xb1\x77\x89\xd7\xf0\x13\x8e\xc8\x0a\xb8\xd0\xad\x3e\x70\x8f\x91\x58\x10\x1a\xc9\x3e\xdb\x46\xb4\xa2\x0c\x61\x94\x8a\xe8\xae\x31\xa3\x31\x30\x41\x80\x8f\x4e\x90\x1c\x12\x42\xa3\x08\x87\x90\xfd\x56\x55\xf7\x02\x87\x80\xe8\x0a\x89\x0d\x20\x1a\x2b\x35\xaa\x9b\xf8\x10\xab\x21\x71\xc1\x48\xb4\x1e\xa9\xcb\x1f\x75\xab\xa1\xc2\xa5\xf9\x87\xfc\xd7\xb6\x37\x20\x51\x9c\x08\x5e\xd4\x7d\xc8\x60\x25\x7b\x7f\x35\xf7\x61\x45\x22\x22\xb5\xf2\x79\x8c\x19\x0e\xcb\xa2\x34\x11\x9d\x65\x59\x12\xed\x96\xe3\x9e\x31\xd8\x2b\x60\xbc\x1e\x89\xb7\xba\x87\x0d\x05\xc7\x3d\x20\x7c\x0b\x2c\xbd\xcd\x41\x7a\xab\x11\x83\x3f\x12\xc2\xc0\x1f\x9d\xa0\xf3\xc2\xdc\x1e\x20\xb4\x50\xed\xd8\xf7\x95\x3c\x0e\x5e\x15\x79\xb0\xc2\x01\x87\x94\x49\xd9\x2d\x72\x7e\x60\xc6\xf0\x87\x57\x0a\x8c\x82\x05\x19\x25\x0b\xcd\x13\x87\x79\x67\xb2\x0b\x52\x78\x82\x00\x26\xcd\xc4\x91\x75\xae\xe9\xf2\x77\xf0\x44\x7e\xdd\xc2\xd7\x7c\x4c\xa5\x4b\xee\xce\x35\x8c\xcc\x9a\x6d\x5c\xdb\xfe\x7c\x9c\x98\xaa\x56\x38\x09\x44\x9d\x1a\x3d\xbc\x5a\x2d\x84\xbf\x06\x8f\x81\x55\x8d\x01\xdf\x73\xcd\x0a\xa5\x14\x11\x8e\xb8\x16\x9c\xb8\xee\xbe\xa4\x34\x00\x1c\xd5\xdf\xdf\xa3\x11\x17\x0c\x93\x48\x54\xd1\x72\xd2\x4e\x0d\xe1\x49\x41\xb2\x7c\x8b\x03\xc7\xed\x6a\x89\x77\x60\x8a\x67\xa2\x3b\xf9\xaa\x3a\x55\x48\x8f\x72\x82\xa4\xbf\x2f\x4a\x8f\x63\xc5\x08\x27\xab\x8b\x9d\x3a\x53\x35\xc4\xef\x9f\x0b\x08\x4d\x94\xcd\x49\xfe\xcf\xeb\x97\x2f\xd0\x6b\xe5\xf4\xd1\xf9\x56\x06\x5d\xc2\x87\x6b\xca\xfc\xc5\x58\x2e\x17\xfc\x64\x3e\x17\x94\x06\x7c\x46\x40\xac\xd4\x72\xb1\x11\x61\x90\xae\x19\xd7\x8c\xac\x37\x62\x5a\x58\x50\xa6\x57\x38\x20\x3e\x96\x77\x98\x1e\x1f\x7f\xc5\xc1\x53\xff\xfc\x66\x76\xff\xf8\xa8\xc4\x9e\xcc\x26\x12\x09\x58\x03\x2b\x37\x86\x24\x22\x61\x22\x1f\xfe\xe3\x03\xcb\xf4\xca\xf6\xf6\x06\xa6\x32\x43\x19\x78\xbf\x4f\x03\x49\x5b\xeb\x48\xcf\xa6\xdd\xcf\x4c\x7b\x3c\xfb\xd6\xb0\x8c\x46\xf0\x72\x55\xe2\xbe\xfc\x69\xf8\x3c\x4b\x58\xdc\x8f\xf3\xa4\x5e\xa5\x0d\x96\xee\x77\x2b\xbb\x8f\xf2\x6f\x0b\xeb\xb4\xe4\xde\xa1\x35\xfd\x0c\xd1\x81\xa6\xaa\xf2\x98\xb5\x84\xa5\x68\x6c\x12\x91\x3f\x12\x68\x6d\x68\x41\x6c\x28\x23\x1f\x3a\x1e\xb5\xca\x2a\xd4\xce\xbf\x97\xe3\x49\xc2\xdc\x01\x48\xd6\xe8\x0a\x3f\x7e\x20\x0c\x3c\x41\x59\xbf\x21\x88\x4f\xd8\xfe\x05\x20\x95\x08\x5b\xf5\x44\x57\x38\x48\xe0\x3b\x69\x31\x5e\x72\x1a\x24\x02\x50\x8c\xc5\x06\xad\x18\x0d\x11\xa3\x54\x48\x3c\xe2\xcb\x35\xa2\x0c\x31\x08\xb0\x20\x57\x69\x0f\xe9\x30\x59\xcc\x40\x80\xaf\x7b\xcb\x48\xc4\x27\x0c\x91\x08\x5d\x6f\x88\xb7\x49\x03\x56\x19\x97\xc8\xe8\xd8\x19\x94\x34\x31\xac\x7d\x4c\xe4\x67\x53\xdb\x39\x2e\xda\xa3\xa0\x45\x52\xca\x1a\xb2\xac\x48\x00\xee\x07\x20\x6f\x75\x3d\x01\xcf\x48\x00\xbd\x92\x5f\xde\xf2\x8e\xfd\xb7\xcd\x7e\x39\x0b\x9f\x05\xf1\x15\x9d\xac\xcc\x8f\x93\x20\x78\xc2\xc0\x2f\x47\xe9\x0e\xb6\x1a\x28\x49\x39\x88\x04\xc1\x01\x47\x09\x07\x1f\xf9\x89\x9c\x05\x84\x13\xb1\x91\xd7\x3d\xb5\x9c\xa1\x6b\x22\xf4\x3c\x72\x9a\x30\x0f\xd2\xa7\x83\x84\x78\x0d\x92\x11\xc5\xfc\x09\xaa\x7b\x26\x12\x0e\xcc\xc8\xa5\xd8\x46\x05\xef\x63\x06\x5c\x6d\xf7\x3d\x0a\xcc\x23\xcb\x00\x90\xa0\x48\x53\x44\x93\xb5\x49\xd8\x90\xeb\xb1\x47\x0c\x31\xe6\x5c\xae\xf7\xb7\x39\x9c\x0a\x47\xec\xd3\x9f\x21\x67\x1b\xfe\x96\x16\xed\xa3\x86\x28\x09\x97\xc0\x76\xed\xf4\xaa\xbd\xba\x67\x25\x82\x40\xc5\xe2\xcd\xe3\x50\x29\x30\xd0\x1e\xe8\xc1\x03\x47\x60\xa6\xb7\xc6\xa5\x26\x7b\x28\xef\x98\xe9\x2a\x60\x45\x4f\x62\x0f\xd5\xa3\x0f\x2d\x81\x91\x02\x43\x01\xe3\x8a\x58\x6f\x01\x18\x88\x92\xb0\x0d\x2e\xb2\xff\x50\xb0\xb8\x92\x02\xcd\x61\xd9\x4a\x68\x20\x76\x5b\xbf\xa2\x2c\xc4\xe6\x7a\xd7\x74\x47\x9b\x3d\xc0\xb6\x3d\xbd\x0d\xc8\x9f\xb5\xef\xe1\xca\xd7\xeb\x21\xa2\x25\x28\x5f\xef\xd2\x60\x2c\xdf\x95\xf6\x74\xfa\xce\xab\x3b\xe0\xad\x4a\xa3\x65\xd1\x72\x8f\x1b\xe2\xf7\x69\x6a\xa2\x4d\xee\x48\x8a\x0c\xc5\x12\x07\x49\xcc\x29\x37\x12\x44\xad\x8d\xd0\x22\x03\x19\xf1\xa8\x8b\x11\x49\x20\x48\x1c\x40\x3b\x3f\x96\x4b\x0d\x95\xe9\xea\x60\x4a\x44\x2b\xcf\x5c\x9d\x0d\x11\x15\x43\x91\xe9\x9b\x46\x09\x92\x1a\xbf\x5a\x34\x6b\xeb\x37\x1a\x1b\xa6\x04\x86\x32\xcd\xc5\xb1\xbf\x6b\x91\x69\x15\x9e\x5b\xc2\x26\xf7\x76\xb3\xd8\xee\x0a\xbd\x5f\x68\xf7\xda\xe7\x96\x33\x65\xf4\xde\x6d\x3a\xdd\x8b\x5e\x1f\x5b\xbc\x74\x9d\xba\xd5\xaa\x4f\x2d\xe5\xf6\x6b\x23\x59\x9e\x84\xf2\x56\x52\x13\x6d\xd7\x76\xa0\xda\xeb\x26\x95\x9f\x57\x2e\x7a\xee\x5a\xc2\x73\xc1\xa1\x16\x0e\x57\x14\xdc\xb5\x06\xd4\xd1\xd4\xa2\xe0\x50\xa6\xba\x7c\x71\x27\x53\x0b\xc4\x6b\x6c\xe5\x56\x66\x28\x03\xcd\x75\xb4\x7b\xe8\x5e\x75\x86\xf6\xd0\xdd\xe9\x77\x6b\x71\x88\x07\x9f\xea\xc7\x0e\x24\x8c\x67\x16\xd5\xf8\x99\x46\xcb\x70\x4d\x2d\xcb\x01\x18\x16\x02\x58\xc7\xa7\xa4\x22\x3c\x14\x7c\xff\xdc\x57\xf8\x1a\xdc\xac\x61\xc5\x6f\x78\x10\xff\xd5\xad\x44\xdb\x66\x21\xef\xa5\xbe\xbb\x73\xf7\xe9\x43\x0c\x91\x0f\x91\xd7\x12\xec\xa2\xdc\x50\x20\x7f\x12\x75\xf0\xba\x10\xb3\x7d\xc1\xfb\xb3\x4c\x16\x99\xb9\x93\x51\x94\x04\x41\x35\xca\x4d\x1d\x50\xe9\xf2\x62\xb7\xd7\xf8\x22\x12\xb2\xd5\x90\x75\x37\x30\x5f\x44\x42\xb6\x03\x30\x5f\x46\x12\xa1\x03\x30\x9f\x5a\xd2\xa8\xc6\xc4\x1b\xe4\x46\xb4\x56\x77\x6e\xa4\xd8\xee\xca\x8d\xbc\x54\x7d\x7a\xcd\x8d\xa4\xfd\xf6\x36\x37\x62\xf3\xde\x37\xcf\x8d\x68\xad\xb7\x9b\x1b\xa9\x7d\x92\xf6\x2b\x37\x52\x9e\x84\x72\x6e\x84\x53\xef\x12\x6a\x78\x5d\x6c\xdf\xc9\x52\x63\xb6\x5e\x2b\xd9\x5a\xbe\xbb\x78\xad\x6f\x7b\x4b\xbc\x6e\x4f\x48\x3d\xdc\xcf\xe2\x44\x46\x8a\xbc\x9d\x2c\x0a\xbc\x5d\x89\xb4\x6a\xaf\xbb\xba\x7a\x7a\xd9\xf1\x4e\x47\x05\xb0\xbb\x30\xae\x2b\x30\x9f\xe5\x56\xa9\x71\x72\xce\x5e\x57\x37\xdd\x72\x12\x03\xe3\xa0\x0e\xbf\x95\xb0\xd0\xd2\x83\xa0\x61\xa6\x97\xda\x96\xfa\x7d\x2c\x60\x2a\x48\x08\x3b\x8b\xfd\xe5\x7c\xc4\x56\x0c\x69\xdb\xfa\xb5\x69\xf6\xd0\xac\xc2\xda\x26\xad\xc5\xd9\x81\xdc\x4a\xa3\x6d\x51\xb7\x5e\xd5\xa0\x26\x7d\x39\x9b\xaa\x13\x6d\x53\xf9\x84\xed\x02\xef\x0c\x69\x91\xf4\x10\x1c\x83\x15\x30\x88\x3c\x40\x98\x23\xf5\x60\x82\x8f\x96\x1f\xd0\xf9\x9a\x88\x4d\xb2\x9c\x79\x34\x9c\x6b\x81\xb9\x4f\xa4\xb9\xcb\x44\x6a\x9a\x67\x72\x39\xde\x3b\x24\x04\x03\xd8\x36\xdc\x9f\xdd\x7f\x98\xab\xe8\x17\x60\x13\x90\x7e\x70\x86\x10\x13\x4b\xea\xa2\xd6\xef\x48\x91\xa1\x58\xf9\xa0\x57\xd0\xb4\x75\xfd\x20\xb5\xa1\x5c\x18\x07\xfe\x1a\x80\xb5\x95\x1a\x0a\xaf\x87\xbd\xe2\x95\xd9\xd8\x0f\x64\x24\xbe\x7a\xd4\x0e\x2e\x29\x31\x14\x54\x8f\x7a\x85\x4a\xd9\xd6\x1b\x4c\x8f\x5b\xc3\xf4\x78\x28\x98\xbe\xe9\x1b\xa6\xc7\x3d\xc1\x94\x30\xd2\x0e\xa5\x84\x91\xa1\x40\x7a\xdc\x2b\x48\xd2\xb2\x7e\x30\xe2\x10\x5e\x35\x38\x56\x78\x86\x38\x84\x38\x12\xc4\x43\xe9\x8b\xde\xe6\x32\xa9\x15\x49\x8c\x34\x76\x27\xf3\x79\x7e\x69\xde\xab\xf5\xe9\x98\xeb\x01\x38\xb0\xb5\x18\xa7\x0e\x7f\x84\x68\x2d\x36\x2d\x2b\xf9\x5a\x68\xa0\x38\xda\x55\xd9\xdd\x51\xc4\xbf\x6f\xb7\x70\x3b\xd8\x7d\xb2\xd0\x55\x7c\xdd\x75\x4c\x61\x52\x36\x60\x9b\x8d\xb3\x1f\x5f\xf8\xd4\xf2\xb8\x35\x9b\xbf\x2f\x2f\x6f\xdf\x61\x27\x9c\x1e\x1d\xe8\x70\xda\x60\x20\x70\xcc\xea\x78\x8d\xb3\xcb\x37\xb2\x23\x06\x6b\x78\xdf\x4b\x0a\x5f\xdf\xa7\x26\xd5\x59\x68\x6f\x9d\xea\xd4\x2f\xcd\x74\x4a\x75\x6a\xf3\xf7\x36\x85\x3f\x50\xc6\x54\x03\x76\xab\x29\xfc\xda\x87\x6a\xcf\xb2\xb2\xa5\x49\x30\xde\x94\x33\x19\x6d\x20\xfe\xaa\x4b\xbd\xa9\xf6\xd0\xd2\xe8\x7c\xfa\xdb\x0c\x4f\xff\x3c\x9b\xfe\x7a\x3c\xfd\x76\x71\xaf\xe3\x8b\x1c\x35\x9f\x1a\x79\x95\x7f\x82\xc7\x31\xe5\x0d\xb5\x65\xef\x88\xf7\xa0\x2b\x7f\xdd\xb6\x07\x65\xc5\xc3\xd4\x3d\xa8\x2b\xd6\x1f\x7b\x50\x57\x2c\xfb\xf4\xa1\xae\xe0\x5a\x9b\x84\xa4\xdd\x9d\xbc\x79\xa0\xc7\xe6\xe8\xcd\x3e\x2e\xa7\x9e\x3f\x38\x9e\xad\x77\x95\xe5\x7f\xed\x8e\x6a\xdc\x1f\xd2\x29\x20\xdb\x40\x4f\xb3\x57\x23\x1a\x28\xaa\x2b\x97\xb7\x52\x54\x17\xaf\x59\xfd\x16\xf7\xd6\xf6\x75\xd8\x5b\xd7\xac\xb3\x58\xee\xb7\x3c\x1c\x04\x68\xcd\x70\xbc\xc9\x7c\xda\x77\x88\x03\xa0\x6d\xc0\x02\xd1\xec\x9a\x5c\x92\x18\x7c\xa2\x3f\xaf\x26\x7f\x9b\x3f\xc1\x41\xf0\x9b\x12\xbb\xe1\x2c\x7a\x34\x12\x98\x44\xc0\xa4\xc6\xce\xc8\xc7\x37\x91\x96\x5e\x3f\x08\x20\xb8\x89\x0e\x0e\x8c\x60\x53\x83\x75\xae\xca\x06\xdb\x66\xad\xdc\xa3\x73\xc5\x2f\x53\xd3\x26\x20\xf2\x42\xf3\x50\xb6\x8d\x3b\x4f\x68\x18\xe2\xc8\x47\x2c\x89\xe4\xee\x1c\xa3\xec\x5e\xdf\x21\x7a\x05\x8c\x11\x1f\x38\xc2\xd1\x07\xc4\x41\x20\x2c\x54\x9c\xa2\x13\xe1\x01\x5c\x81\x25\xc1\xeb\x8e\xee\x91\x3b\xc2\xb7\x0d\xad\xf5\xbb\xd8\xb5\xd3\x6a\x7b\x23\xbb\x3c\xb9\xe9\x6f\x66\x30\x48\x98\x35\x7a\xaa\x39\x0b\x6c\x33\x66\xfb\xd5\x15\x02\x1c\x91\x48\xa1\x98\xcf\x6a\x45\x78\xd7\xf1\xe8\xb4\xdb\xbb\xf1\xb9\x0e\x38\x16\x27\x47\xa7\x32\xfc\xb8\xb8\x98\x17\x22\x90\x43\xab\x94\x33\x14\xd9\xfe\xd8\x44\x6c\x26\x8d\xaf\x49\x10\xa0\x25\xa0\x25\x4d\x22\x5f\xcd\x0c\x0e\xb3\x4f\x45\x6c\xbf\x14\x50\x4d\xe8\x54\x20\x54\x67\x1c\xad\x9d\x3e\xda\x65\x9b\x8e\xd0\x27\x0c\x31\x58\xe9\xef\x16\x94\x46\xb5\x7b\x50\xf6\x43\xaa\xe9\xb0\x2c\x57\x17\x95\x6b\x66\xaf\x8a\x31\xcd\x42\x66\x8b\xe8\x08\xa2\xab\xb7\xb8\x17\x5e\x3e\x8d\xae\x08\xa3\x51\x08\x91\x40\x57\x98\x11\xbc\x0c\x7a\x65\xe8\xf9\xbb\xef\x6f\x81\x88\x24\x42\xdc\xa3\xb1\x2a\xd6\xa1\xeb\xb9\x26\x66\x84\xc3\x5b\x65\x63\x47\x7f\x96\x2a\x6b\xeb\xd5\xd2\x21\xdf\x36\x53\xe5\xf6\xa0\x0f\x9e\x3e\x23\xfd\xf2\xf2\xce\x73\xd6\x8e\x50\x4e\x5b\x1a\xbc\xa6\x51\xf5\xe7\xe9\x48\x55\x24\xd3\x85\x9e\x35\x91\x17\xd2\x19\x0b\x33\xff\x9a\x35\x55\x3e\xc4\xb0\x3d\x4f\x20\x68\xf6\xa9\x1d\x2b\xac\x1d\x7c\x80\x85\x06\xd6\x8f\x08\x35\xb8\x4d\x2e\xd6\x61\x82\xac\x09\x9c\x12\x58\xc6\xd5\x45\x7f\x73\x6c\xf9\x0a\x11\xb2\xcc\x83\xfa\xb2\x73\xc9\xbb\x20\x0f\x47\xf2\x29\xcd\x8e\x5d\xa8\x02\x96\xfa\xf2\x15\x15\x1b\xbd\x03\xd6\x3d\xf9\xce\xaa\x55\xed\x08\x63\xca\xec\x29\x3a\x73\xef\x2d\xfb\xa5\x4e\x23\xfb\x00\x57\x3e\x5c\x41\xd5\x85\x0d\xe5\x35\x99\x43\x27\xa1\x9b\x39\xce\x73\xe5\x1f\xc7\x53\xfd\xff\xa3\xd3\xb1\xf0\xe2\xff\x25\x7e\x7c\x74\xda\x90\xee\xff\xa6\x5c\x20\x69\xf0\x98\x1f\xc9\x11\x2f\x89\xf2\x80\x76\xc2\xef\x28\xfd\xa1\x72\x4a\xbf\x32\xb8\x2e\x4c\xed\x4c\x33\x9d\x1d\xea\xb4\xd6\x35\xc5\xfe\xc4\x9d\x5e\xcc\x3a\x55\xf6\x77\x5b\x76\xa4\xc7\x71\xb1\xef\x4b\x6f\x81\x42\x1c\xc7\xa0\xd6\x1e\xbc\x6d\xb2\x1d\x86\x42\x4d\x5c\xf9\x90\xa8\x0a\xff\x29\x33\xf7\xbd\x7d\x82\xfa\x6e\xe6\x5e\xed\xdd\x58\x0a\x1f\x18\x43\x31\x83\x15\x79\x5f\x86\x52\x07\x73\x7b\x0a\xe5\xcb\xa4\xc9\x8b\x09\x7f\x37\x94\x34\x11\x9f\x18\x94\xd7\x94\x5d\xfe\x50\xf9\xe2\xa9\xcd\xd0\xff\x52\x76\x29\xad\xf0\x0b\x5f\x5d\x15\x1b\x34\x2e\xe7\x56\x0a\xe7\x23\xd4\xf2\xbf\xfb\x14\xc4\x81\xcb\xd2\x72\x51\xc6\xb9\xee\xa6\x91\x4f\xe1\xda\xa2\x8f\x6a\x8e\xbd\x50\x93\xc7\xeb\x07\xc6\xbd\x5a\xbc\x6b\x14\x3b\xb3\x6c\x69\x53\xe7\xf4\x1a\x8d\xcd\xbc\x5a\xdd\x8b\xde\x75\x39\xb7\xf8\xd2\xac\x59\xee\x52\xb7\x4b\x25\xba\x49\x30\x89\xe4\x78\x3a\x2f\xac\x79\xb9\x39\x61\x64\x9a\x05\x41\x77\x41\xa6\xe5\xee\xd5\x3f\xc1\x91\xb5\x98\xa5\x5e\x69\x2a\x0e\xc8\x9f\xc0\xd1\xf3\x17\xaf\xde\xfc\xf2\xdb\x8b\xb3\x9f\x9e\xea\x70\xee\xed\xd9\x8f\x6f\x9e\xca\xcd\x55\x7a\xe8\xfc\xeb\xbc\xc3\x89\x6e\xfc\x7a\x86\x9e\xaf\xb6\xfd\x38\x92\xfb\xbf\x09\x22\x02\xfd\xf4\xe6\xf5\x2f\xea\xf3\x6d\x9c\x27\x21\xf8\x69\x8f\xef\xd1\xe1\x38\x57\x51\xe3\x54\x6e\x1a\x98\xd4\x56\x3e\xb3\x6e\x1d\x37\xcf\xfd\x6f\x76\x3f\xd1\x5c\x4a\xf5\x6f\xb5\x64\x4d\x35\x24\xcb\xe9\xf5\xf2\xcd\x2f\x19\xdf\x0a\x24\xd3\xf4\x2a\x34\x6a\x92\x95\x7a\xd7\x50\x4d\x75\x90\x4c\x2b\x08\xdc\x51\xcd\xd4\xb8\x5f\x59\x10\x67\xbc\xd0\xf0\x90\x42\xbd\xa7\x95\xab\x8e\xe9\x69\xeb\x88\xdd\x52\xfd\xee\x57\x14\xfe\x9e\xa8\x86\xc6\x37\x08\x67\x4a\x95\x4a\x5b\x50\x53\xea\xd0\x39\xb4\xd9\x6a\x71\x05\x38\x37\x3e\x82\xe7\xb9\x5e\xc0\x6a\x08\x62\x36\xc0\xee\x50\x16\x0a\xb6\xd6\xca\x79\xde\xdc\x19\x46\xad\x63\x6f\x41\x4c\x87\xd7\x1d\xc2\xc2\xd2\x67\x83\xb0\xd0\xec\x3a\x83\x70\x16\xe9\xbf\xc8\x34\x49\xbf\xd1\x37\x49\xdf\x47\x9f\x20\xca\xd2\x9d\xdc\x0c\xe9\x03\x81\x3c\x5b\x42\xf4\x07\xe5\x69\x80\x05\xf8\x08\x0b\xc4\x92\x48\x90\x10\x26\x88\x41\x1c\x60\x4f\xee\x9b\x0e\xc7\x3f\x3f\x7d\x76\x24\x37\x46\x62\x03\xba\x46\x82\xe8\x0a\xfd\xfc\xf4\xd9\x4c\xfe\x27\x5f\x8d\xd2\x82\xcf\xe1\x58\x95\x79\x54\xa6\xfa\x72\x8d\x56\x1c\x1d\x8e\xe7\x72\xcb\x75\x24\x37\x92\x33\xf4\x02\xb8\x90\x7a\x93\x58\x46\xc8\x34\x4a\xcb\xd8\x99\x1e\x9e\xc4\x31\x65\x02\xfc\x09\x22\x33\x98\x49\xe9\x54\xe5\x51\x85\x3e\xe7\xdb\x99\xcf\xbe\xb4\x37\xc9\x29\x95\xb9\x7c\xe3\x7c\x87\xfe\x53\x65\x76\xa2\xaa\x26\x17\xc2\x57\x0f\x66\xc7\xb3\xe3\xea\xf1\xfa\xf1\xf6\x88\x47\xf9\x20\x3d\x8f\xc1\x9b\x6b\x99\xd9\x46\x84\x41\x75\xf8\x66\xf0\x5f\xcc\xa4\xbd\x1b\xa7\x39\xb4\x8b\x8b\x99\xe5\x9f\xe3\xd3\x93\xf1\xc5\x85\xca\xb3\x9d\x4d\x7f\xc5\xd3\x3f\xa7\x8b\x7b\xe3\xd3\x93\x8b\x8b\x59\xe9\xd2\xd1\x3f\x8e\x8e\x4e\xd5\xf5\x7b\x85\xeb\x17\x17\xd3\x8b\x8b\xd9\xe2\xde\xd1\xe9\x61\xe1\x4f\xb6\x1d\x7c\x3c\x38\xf8\x7f\x00\x00\x00\xff\xff\x48\xdf\x7f\x2b\x16\x70\x00\x00")

func githubComOpspecIoSpecSpecPkgManifestSchemaJsonBytes() ([]byte, error) {
	return bindataRead(
		_githubComOpspecIoSpecSpecPkgManifestSchemaJson,
		"github.com/opspec-io/spec/spec/pkg-manifest.schema.json",
	)
}

func githubComOpspecIoSpecSpecPkgManifestSchemaJson() (*asset, error) {
	bytes, err := githubComOpspecIoSpecSpecPkgManifestSchemaJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "github.com/opspec-io/spec/spec/pkg-manifest.schema.json", size: 28694, mode: os.FileMode(420), modTime: time.Unix(1506459897, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"github.com/opspec-io/spec/spec/pkg-manifest.schema.json": githubComOpspecIoSpecSpecPkgManifestSchemaJson,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"github.com": &bintree{nil, map[string]*bintree{
		"opspec-io": &bintree{nil, map[string]*bintree{
			"spec": &bintree{nil, map[string]*bintree{
				"spec": &bintree{nil, map[string]*bintree{
					"pkg-manifest.schema.json": &bintree{githubComOpspecIoSpecSpecPkgManifestSchemaJson, map[string]*bintree{}},
				}},
			}},
		}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
