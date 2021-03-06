// Code generated by go-bindata.
// sources:
// templates/bosh_director.tf
// templates/cf_dns.tf
// templates/cf_lb.tf
// templates/concourse_lb.tf
// templates/jumpbox.tf
// templates/vars.tf
// DO NOT EDIT!

package gcp

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

var _templatesBosh_directorTf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x96\xdf\x8e\xa2\x30\x14\xc6\xef\x79\x8a\x93\x66\x2f\x66\x12\x64\x15\x45\xd9\x0b\x9f\xc4\x98\xa6\x60\x65\xd8\xa9\x94\x94\xa2\xb3\x99\xf0\xee\x9b\xb6\xfc\x17\x19\x9c\x19\x33\x7a\x01\xd2\xef\x9c\xf3\xf5\xd7\x53\xe9\x99\x88\x98\x04\x8c\x02\xca\xf2\x20\xa1\x12\x87\xf1\x41\x20\x78\xb7\x00\xe4\xbf\x94\x02\x00\x6c\x01\x65\x52\xc4\x49\x84\x2c\x80\x03\x3d\x92\x9c\x49\xf5\x70\x31\x77\xf4\xf7\xf7\x62\x8d\xac\xc2\xb2\x04\xcd\x78\x2e\x42\x0a\x28\xe2\x3c\x62\x14\x87\xfc\x94\xe6\x92\xe2\x84\xca\x0b\x17\xaf\x08\x50\x10\xb0\x59\xfd\x4b\xd5\x48\xc8\x49\xd7\xe8\x7f\xb6\x80\x7e\xbd\x9f\x89\x70\x68\x72\xc6\xf1\xa1\xa8\xa3\x2c\x00\x92\x4b\x8e\x43\x41\x89\xa4\xd8\x98\x56\x23\x19\x6c\xe1\x48\x58\x46\x47\xad\x34\xfa\xd2\x8d\x79\x30\x64\xe6\xca\x42\x29\xb5\x00\xe2\x54\x63\xc2\x82\x24\x11\x6d\x84\x2d\x82\x85\x92\x95\x95\xda\xf9\x86\xc9\x38\x2d\x2e\x4e\x46\xd9\x11\xb3\x38\x79\x2d\xc6\xa9\x1e\x63\x41\x2f\x84\x31\x04\x88\xbe\x49\x2a\x12\xc2\xba\xd3\xb8\x9a\x40\x2d\x6b\x79\x9b\xea\x4a\x25\x2d\x90\x65\x01\x18\x3b\x66\xee\x0a\xfa\x0e\x55\x7d\x30\x47\x7b\x25\x20\x8c\xf1\x8b\x76\x02\x90\x72\x21\x33\x63\x66\x87\x5c\x17\xd9\x80\xd6\xfe\xda\x57\x57\xd7\xf3\x3c\x0f\xed\x8d\x4c\x70\xc9\x43\xce\x94\x1d\x19\xa6\xca\x60\xa1\x52\x49\x22\x22\x2a\xb1\x24\x91\xa9\xd4\x9d\x4f\xc0\xb3\x97\x19\x4f\x69\x82\xf6\x53\x49\x35\x21\xe3\xa8\x1a\xdd\x77\xb0\x9a\xe0\x7f\x3a\x37\x7f\xb5\x5a\xea\xab\xbf\x5a\x7d\x23\xc7\x43\x2c\x68\x28\xb9\xb8\x93\x65\x1d\x36\x81\x67\xad\x7d\x34\xd3\xd6\x5c\xfa\x5c\x3f\x05\x28\x4e\xca\x8d\x33\x99\x4d\x15\x31\x93\x7c\x2a\xa2\xc1\x90\x07\x92\x6a\x4d\x6a\xac\xf9\x56\xae\x69\x3f\xd7\x73\xbd\xb9\xb9\xd9\x6c\x36\x3f\xd1\x6f\x7f\xf3\x53\x1a\xf0\x37\xc5\x47\x3f\x18\xa5\xd9\x13\x3f\x90\x63\x59\x69\xd2\x1e\x5e\x2e\xfd\x3f\x5f\x42\x57\x2f\x9a\x3d\xbe\xc3\xee\x6e\xd4\x89\xcd\xf9\x43\x0d\xd9\x62\x15\x87\xa7\x06\xd6\x94\xad\x7d\x4b\x93\x1f\x3e\xb7\xfd\x19\x0f\x09\xcb\x74\xc2\x6a\x40\xbf\xf2\x0d\x0f\x75\x67\xce\x01\x4f\xbd\x23\x81\x0d\xbe\x0d\xf3\x67\xf3\x56\xe7\xb9\x4c\x73\x09\xa8\x73\x18\x3a\x13\x96\xd3\x7b\xb1\xb6\x92\xb5\x0f\x37\xe3\xf9\x1a\xa5\xd3\x1c\x82\x06\x32\x56\xfd\x84\xd5\x50\x2f\xa9\x6e\xb8\x0e\xa8\x4e\x68\x87\xcd\x95\x1f\x0d\xd1\xe9\x68\x6e\x84\x47\x97\xab\x60\xa5\x7e\xe1\x99\x7c\x1a\xc8\x62\xc3\xa2\x87\xb8\xdc\x9f\x18\xd7\xba\x38\xbd\x33\xa5\xf7\x7c\x03\xcb\x17\x72\xae\x6f\xd9\x54\x4d\xd8\xcd\xb5\xd3\xbd\x7b\xbd\x8a\xd5\x36\x76\xea\x83\x44\xb9\x84\x76\x15\x30\xf4\x3f\xa5\x06\xf7\xc3\xb3\x19\x28\xfd\x51\xd5\x2a\xb8\xac\xbc\x1f\x5c\x43\x49\xa2\xa1\x0e\xba\x9d\xba\x8a\x6c\x5a\xf2\x7f\x00\x00\x00\xff\xff\x08\x23\x4c\x5b\xa8\x0c\x00\x00")

func templatesBosh_directorTfBytes() ([]byte, error) {
	return bindataRead(
		_templatesBosh_directorTf,
		"templates/bosh_director.tf",
	)
}

func templatesBosh_directorTf() (*asset, error) {
	bytes, err := templatesBosh_directorTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/bosh_director.tf", size: 3240, mode: os.FileMode(480), modTime: time.Unix(1521677188, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesCf_dnsTf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x96\xcf\x8a\xdb\x30\x10\x87\xef\x7a\x8a\x41\xf4\x54\xb0\x59\xe8\x79\x0f\x85\x9e\x7b\xe9\xb1\x2c\x46\x91\x26\x8e\x8a\xac\x11\x33\xb2\xb3\xe9\xe2\x77\x2f\x72\xe2\x6c\xd2\xad\xe9\xfa\x10\xc8\x61\x73\xc9\x1f\xfd\x34\xf3\xcd\xa7\x60\x34\x18\xf6\x66\x13\x10\xb4\x1c\x24\x63\xd7\x38\xea\x8c\x8f\x1a\x5e\x14\x40\x3e\x24\x84\x47\xd0\x92\xd9\xc7\x56\xab\x51\x29\x46\xa1\x9e\x2d\x82\x6e\x89\xda\x80\x8d\x8b\xd2\x74\x26\x9a\x16\x5d\xf3\x9b\x22\x6a\xd0\x18\x87\xe9\xe7\xe3\xd7\x52\x28\x9a\x0e\xe1\xf4\x7a\x04\xfd\xe9\x65\x30\x5c\x97\x98\x77\x63\x35\xc5\x14\x40\xd9\x32\x07\xcf\xa1\x2b\xaa\xb1\x9e\x72\x28\x96\x7d\xca\x9e\x62\xc9\x7d\xfb\xfe\x03\x4a\x09\xd8\x12\x43\xde\x21\x5c\x55\x07\x8c\x83\x67\x8a\x1d\xc6\x3c\x0d\x40\x7d\x4e\x7d\xfe\x6b\xdc\x09\x57\x90\x07\x64\x39\x12\x0f\x26\xf4\x78\xc4\x58\x18\xb4\xbe\x1c\xb3\x2e\xe0\x73\x85\x71\xd9\x14\xa3\x25\x76\x8d\x60\xd6\xa0\xf7\x3e\x38\x6b\xd8\x55\x2e\xca\x1b\x4f\x8f\xa0\x3f\xd7\xef\x6c\x3e\x9b\x1b\x8f\x7a\x12\x46\x27\xcd\x64\xe7\xe7\xdc\xdc\x52\x97\xfa\x8c\x4d\x1b\x68\x63\x42\x63\x9c\x63\x14\xa9\xed\xb6\x3a\x7d\xd4\x4f\xf3\x81\x9f\xfb\x7f\x2d\xe5\x72\x0e\xaf\x27\xf7\xe5\xe1\x41\x29\x80\x4b\x92\x95\x8e\x46\x5d\x0a\x30\x3b\x93\x8d\x4c\x80\xe7\xcd\xff\x45\xac\x4f\xef\xa3\x7e\x7a\x9f\xe0\x0d\xc9\x6e\x49\x6e\x59\xbb\x81\xdf\x19\xf5\x57\xdf\xa5\x0d\x3d\x57\x3e\xdd\x8f\xd8\xb7\x6c\xab\x8d\xda\x6d\x25\xb2\xab\x12\xd3\xf3\xe1\x5f\x56\xe5\xa6\x52\xaf\xba\xdf\x9d\xd6\x4b\xba\xd5\x62\xb3\x4d\x4b\xff\xd4\x6c\xd3\x6d\x9d\x96\xde\x4c\x7d\x46\xbe\x4b\xa9\xaf\x78\xab\xad\x3a\x4a\x29\x20\x2f\x99\x3d\x2d\xdf\xd6\xee\xfe\x8e\x1e\xad\x57\x58\xab\x6d\x06\x6a\x5b\xc6\xd6\x64\x5a\x34\x7a\x11\xf9\xb0\xba\xf2\x16\xb0\x97\xe5\x8b\xc0\x5e\x3e\x74\xaa\x51\xfd\x09\x00\x00\xff\xff\xab\x21\x00\xec\xae\x0a\x00\x00")

func templatesCf_dnsTfBytes() ([]byte, error) {
	return bindataRead(
		_templatesCf_dnsTf,
		"templates/cf_dns.tf",
	)
}

func templatesCf_dnsTf() (*asset, error) {
	bytes, err := templatesCf_dnsTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/cf_dns.tf", size: 2734, mode: os.FileMode(480), modTime: time.Unix(1520537908, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesCf_lbTf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x58\xcf\x8f\xe3\x26\x14\xbe\xe7\xaf\x40\xa8\xc7\x75\x26\xc9\x4c\xb7\xe9\x61\x4f\x55\xaf\xdb\x1e\x7a\x5b\x8d\x10\xc6\xcf\x09\x0a\x63\x28\xe0\x64\xa3\xd5\xfc\xef\x15\x60\x27\x76\x6c\x63\x3b\x33\xab\x6a\x76\x0f\xf1\x18\xde\xf7\xe0\x7b\x3f\xf8\xf0\x91\x6a\x4e\x53\x01\x08\x1b\x23\x08\x03\x6d\x79\xce\x19\xb5\x80\xd1\x8f\x05\x42\xf6\xac\x00\x7d\x41\xd8\x58\xcd\x8b\x1d\x5e\xbc\x2e\x16\x83\x16\x44\x69\x7e\x74\xbf\x07\x38\x0f\x5a\xcb\xd2\xaa\xd2\x22\xac\x65\x69\x41\x93\x94\xb2\x03\x14\x19\x31\xa0\x8f\x9c\x55\x4e\x8f\x54\x94\xde\xee\x97\x1f\x3b\x29\x77\x02\x08\x93\x2f\xaa\xb4\x70\x3b\x7d\x19\x50\x12\x91\x26\xd5\x48\x52\x8f\x14\xf4\x05\x5e\xfb\x3c\x8a\x94\x70\x35\xe6\x67\x27\x64\x4a\x05\xa1\x59\xa6\xc1\x98\x25\xcb\x93\xfa\xb1\xfa\x6d\x43\x1b\xb3\x27\x4a\xcb\xef\xe7\x69\xe8\x0d\x58\x63\xf6\x89\xb7\xec\x07\xb6\x4c\x91\x39\xeb\x6e\x20\x5b\xa6\x92\x60\xda\x0f\x7d\x32\xb3\x21\x4f\x37\xdb\xd7\x60\x64\xa9\x19\x20\x7c\x63\x93\x73\x0d\x27\x2a\x04\x46\xb8\x7e\x4c\x58\x1e\x3c\xb9\xc0\xa0\xf0\xcf\xbb\x3b\x52\xbd\x84\xe2\x48\x78\xf6\x9a\xb0\x3c\x91\x0a\x0a\xbc\x40\x28\x03\x05\x45\x66\x88\x2c\xd0\x17\xf4\xed\xd6\x41\x01\xf6\x24\xf5\x61\x99\xa6\x22\xa9\x9e\xf1\xb3\x03\x0f\xcf\x17\xf0\x71\xb3\x3a\x51\x16\x08\x51\x21\xe4\xc9\xaf\x11\x21\xa5\xa5\x95\x4c\x0a\x07\x63\x99\xc2\xe1\xa5\xd4\xd6\x04\xec\x6f\x78\xbb\xc2\x9f\x10\x7e\x7a\x7a\xf4\x8e\x5f\x1d\x40\x60\x83\x68\x5a\xec\xc0\xf8\x49\xab\xa5\xff\xff\xb0\xc2\xcf\x6e\x82\xa5\x7a\x07\x96\x58\xba\x0b\xc3\x6f\xce\xef\xe7\x68\x18\xda\x59\x8c\x11\xbe\xe6\x71\x23\x16\x3d\x51\x88\x47\xb7\x82\xcd\xa5\x3e\x51\x9d\xf1\x62\x47\x74\x29\x20\xc0\xef\xad\x55\xc9\x75\x24\x09\x23\x13\xe2\xee\x0c\x1d\xcb\x5c\xd5\xeb\xbd\xbb\x30\x6b\x9e\xd1\x50\x1a\x54\x61\x70\x2e\x43\xd9\x2e\xeb\x95\x8b\xb4\xaa\x46\x03\x22\x27\x82\x17\x07\x8f\xe7\x02\x1f\xc2\xea\xf0\xb6\xab\xb7\xf1\x63\xee\x26\xc8\xfc\x0f\x0c\x99\x36\x45\x66\x1a\x47\xae\x2e\xa2\x24\x75\x62\xd0\xc8\x9f\xda\x43\x87\x97\x2e\x31\x7e\x7e\x98\xec\x9b\x86\x61\x9a\x2b\xcb\x7d\xd7\xc0\x1a\xa8\x10\x67\x44\x91\x90\x34\x43\x29\x15\xb4\x60\xa0\x51\x5a\x5a\x24\xb8\xb1\x90\x21\x6a\x10\x2d\x90\x03\x41\x17\x90\x52\x0b\xf2\x42\xd5\x20\x37\xd5\x78\x8b\x90\x52\x8b\xc4\xbd\x6b\x52\x32\x71\xf7\xe6\x76\xfb\x26\xb2\xff\x61\x12\x4c\x3f\x0b\xb5\xc1\x1c\x2a\x4c\x3f\x17\x6f\x26\x04\xa1\x1b\xc1\x30\xd0\x04\x6f\x66\x39\x5c\xf7\x67\x13\x2b\xde\xf7\x3a\x4a\x06\x57\x10\x57\x42\x89\xd2\x90\xf3\xef\x1d\x2e\x7b\xb2\xa8\x34\xa0\x1d\x23\x47\x9e\x41\xe6\xb6\x80\x2a\x9d\x83\x0e\x70\x46\x0f\xfe\x4d\xc3\x1b\x52\x94\x6b\x5f\x10\x57\x35\x74\x75\x13\x91\x4c\xde\x77\x13\x68\xc8\x28\x9c\x56\x82\xe7\xc0\xce\x4c\x40\x75\x62\x31\x0d\x0e\x28\x85\x5c\x6a\x20\x19\x18\xab\xa5\x73\x6c\x75\x09\xfe\x80\x8a\x31\x56\x85\xf0\x26\x09\xab\x20\xc6\xcf\x8a\xaa\x73\x7b\xde\x72\x5a\x0a\x5b\x1f\x5e\x6f\x14\x70\x53\x4b\x69\x0f\x54\xd8\x3d\x61\x7b\x60\x87\xb0\x7e\x55\xa6\x82\xb3\x24\x0c\x24\xd5\x40\x74\x0b\xc1\xc2\x6f\xc2\x37\xa4\x26\x66\x2d\x08\xa4\xb6\x8d\x2a\xd8\xae\xb6\x2b\xff\x5e\xc3\xbf\x25\x18\x4b\x14\xb5\x7b\x87\xfd\x10\x6c\xf1\x28\xe5\x1d\x47\x53\x16\x3f\xd8\x02\xdc\x99\x3d\xb4\xc8\xc1\x25\x4e\x94\x70\x2e\xc6\xb1\xe5\xf4\x26\x45\xd3\xe0\x63\xc8\xb9\x20\xe8\xb6\xab\x98\x9e\x5b\x3f\xae\x96\x9b\xf5\xda\x6b\xba\xcd\xc6\xcd\x7f\xfc\x75\xb9\xfe\x3d\xbc\x58\x7f\xf6\xa6\x4d\x91\x87\xde\x51\xe6\x75\x2f\x1b\x95\x27\x25\xa5\x18\x53\xf1\x8d\xa9\xed\x6b\xc7\xf5\x9e\x34\x98\x0a\x2d\xfd\x78\xb1\x1c\x29\xa9\xeb\xbc\x19\x69\xd6\x07\x3e\x9c\x63\x97\xd9\x1f\xe7\xd2\xb0\xd9\x6c\x36\xd7\xfc\x1a\xbd\x0e\x8c\x44\x2d\x7e\x0a\xb6\xb2\xe3\xce\xd0\xb9\x22\x00\x63\xb8\x2c\x08\xcd\x73\x5e\x70\xeb\xcf\xb2\xaf\x7f\x7d\xfd\x73\x24\xae\x7d\xe2\x77\x38\xbc\x63\xeb\x68\x09\xd6\x79\x09\x3e\xa8\x52\x1d\x8c\x8f\x47\xd0\xd4\xcd\xe0\xfd\xf3\xc7\xdf\x37\x4a\xfb\xfd\xee\xf1\xf7\x17\x6d\xe3\x46\x3f\xa1\x6a\xdb\x95\x75\xb5\x9d\x54\x5a\x8d\xe9\x1f\xa1\xac\xd6\xab\xcd\x53\xf2\xb8\xf9\xed\xf3\xf6\xfe\xe2\xea\xb0\x1b\xaf\xae\x56\x53\xec\x65\x77\x8c\xd7\x3b\xc4\x41\x24\x8a\x53\xe2\xd8\x91\x07\xf7\x8a\x83\x4e\x6b\xb9\x8b\x80\x68\x73\x71\x52\xac\xb1\x7f\x1f\x43\x1f\xf8\x6e\x20\x3b\x64\xf5\x86\xf3\xd3\x02\xa1\x78\x48\x7b\x7b\x56\x8c\xf2\x71\xc6\x67\x76\xad\xc6\xa2\xa3\x6d\xab\x91\xef\xef\xd1\xbc\x26\x7c\x2a\xbc\xbf\x6b\x9d\xcc\x6c\x8d\x71\x1a\xf9\x3c\xe5\x26\xcc\xcb\xcf\x49\x88\xb3\xf3\x71\x62\x2a\xf6\x48\xfa\x49\x2d\xa6\x37\x1f\x4f\xa6\xfa\x12\x34\x29\x1b\x2f\xb3\xe7\xe7\xe2\xc9\xc4\x73\xd0\x7f\xe1\x79\x87\xe4\x9b\xfe\x51\x39\x42\xc7\x2c\x36\x7e\x02\x19\xdb\xd5\xcf\xe0\xe2\xbf\x00\x00\x00\xff\xff\x26\x2a\xa4\x82\x9c\x19\x00\x00")

func templatesCf_lbTfBytes() ([]byte, error) {
	return bindataRead(
		_templatesCf_lbTf,
		"templates/cf_lb.tf",
	)
}

func templatesCf_lbTf() (*asset, error) {
	bytes, err := templatesCf_lbTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/cf_lb.tf", size: 6556, mode: os.FileMode(480), modTime: time.Unix(1520537908, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesConcourse_lbTf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x93\xc1\x6e\xdb\x30\x0c\x86\xef\x7a\x0a\x82\xd8\x71\x32\x0a\xaf\x87\x5e\x76\x1a\x76\xed\x76\xd8\x6d\x28\x04\xc5\xa6\x1d\xa1\xaa\x28\x48\x72\x8c\xa1\xf0\xbb\x0f\xb2\x1d\xc7\x5d\xb3\x24\x40\x10\xa0\x27\x13\x34\xf5\x93\xfa\x7e\x8a\xbb\xe4\xbb\x04\x58\xb1\xab\xb8\x0b\x91\x54\xd2\xa1\xa5\xa4\x3c\xb3\x45\x78\x15\x00\x3b\x6d\x3b\x82\xaf\x80\x9f\x5e\x5b\xe6\xd6\x92\xaa\xf8\xc5\x77\xe9\x4d\x69\x31\xc5\x72\x8c\x9d\x7e\xa1\x01\xc5\x20\xc4\x7b\x79\xbb\x51\xc6\x9f\x13\xd6\x75\x1d\x28\xc6\x62\x39\x26\xf7\x99\xf9\x3b\xa9\x07\x8a\xdc\x85\x8a\x00\xff\x39\xdf\x98\x40\xbd\xb6\x16\x01\xf7\xa1\x5c\xb4\xa6\xe6\x79\x46\x00\x98\xda\xef\x74\x28\xc8\xed\x94\xa9\x87\x43\x9d\x64\x4f\x0e\x73\x29\xa5\x9e\xc3\xf3\xd1\x49\xe7\x7f\xc5\x66\x63\xe5\x3e\x9e\xaf\x2f\x00\xb4\xb5\xdc\x8f\xed\x00\x7c\xe0\xc4\x15\xdb\x2c\x93\x2a\x8f\x53\x92\x43\x8a\xd3\x18\xbf\xf1\xe1\x0e\x3f\x03\xde\xdf\x7f\xc9\x9f\xb2\x2c\x4b\x7c\x12\x00\x43\x16\x9a\x49\x27\xdd\xc6\xb1\xf4\x70\x99\xa7\x93\x20\x66\x5c\xb8\x72\x40\x2e\xb9\x05\xc3\xff\x19\x9c\xc6\xfc\x66\x55\x70\xb5\x01\x17\x6a\x0b\x80\x48\x31\x1a\x76\x4a\x37\x8d\x71\x26\xfd\xc9\xf5\x8f\x3f\x1e\xbf\x9f\xf1\x97\x43\xaf\x43\x6d\x5c\xab\x42\x67\x09\x01\x63\xdc\xca\x43\x56\x4e\xd9\xb5\xcf\x67\xbc\x8e\x71\x8b\x0b\xe7\x55\xf5\x85\x1b\x1f\xc9\x36\xca\x1a\xf7\x3c\x64\x95\xec\xaa\x0a\xda\xb5\x34\xaa\x8c\x56\x0a\x00\xe3\xd5\x7a\x09\x7e\x7d\xfb\x39\x67\x67\x47\x8e\xb7\xbc\xfa\x2d\xbc\x63\xb5\x4d\xc9\xc7\xab\x68\x8d\x0a\x37\xe3\x95\x5f\xc0\x07\xc3\x75\x35\xad\x9b\xc1\x7a\xb8\xbb\x35\xab\xbf\x01\x00\x00\xff\xff\x5f\xcd\x56\x41\x23\x06\x00\x00")

func templatesConcourse_lbTfBytes() ([]byte, error) {
	return bindataRead(
		_templatesConcourse_lbTf,
		"templates/concourse_lb.tf",
	)
}

func templatesConcourse_lbTf() (*asset, error) {
	bytes, err := templatesConcourse_lbTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/concourse_lb.tf", size: 1571, mode: os.FileMode(480), modTime: time.Unix(1520375763, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesJumpboxTf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x8f\x41\x0a\x83\x30\x10\x45\xf7\x39\x45\x08\xdd\xaa\x20\x64\x23\xf4\x2c\x21\x35\x83\xb5\x44\x27\x4c\x66\x44\x10\xef\x5e\x4a\x6d\xad\xd0\x4d\xbb\x1d\xe6\xbd\xff\x3f\x41\x46\xa1\x16\xb4\xe9\x10\xbb\x08\xae\xc5\x21\x09\x83\xf3\x21\x10\xe4\x6c\xb4\xb9\xc9\x90\x2e\x38\x17\x7d\x32\x7a\x51\x5a\x8f\x7e\x00\x7d\xd6\xe6\xb4\x4c\x9e\x4a\x18\x27\xd7\x87\xb5\xf8\xf8\x52\xab\x52\x28\x9c\x84\xdf\xb0\x13\x8a\x4f\x7a\xf2\x51\x36\xfc\x7b\x62\xb9\x9b\xca\xed\xb4\x36\x75\x7d\xb0\xc2\xcc\x40\xa3\x8f\xee\xd5\xe9\x2f\xeb\x41\x19\x7a\x82\x96\x91\xf6\xe1\x07\xef\x95\x39\xe5\xa6\xaa\x7e\x6b\x6d\xad\xb5\x8f\x94\x7b\x00\x00\x00\xff\xff\x10\xf1\x6b\x1b\x66\x01\x00\x00")

func templatesJumpboxTfBytes() ([]byte, error) {
	return bindataRead(
		_templatesJumpboxTf,
		"templates/jumpbox.tf",
	)
}

func templatesJumpboxTf() (*asset, error) {
	bytes, err := templatesJumpboxTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/jumpbox.tf", size: 358, mode: os.FileMode(480), modTime: time.Unix(1520375763, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesVarsTf = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\xd0\x41\xaa\x03\x21\x0c\x06\xe0\xbd\xa7\x08\x61\x16\xef\x6d\x7a\x83\x9e\xa5\xd8\x31\x95\x14\x31\x92\x11\xa1\x15\xef\x5e\x9c\x19\x98\xe9\xa2\xa0\x4b\xf3\x85\x3f\xfc\xc5\x2a\xdb\x7b\x20\xc0\xa4\xf2\xa4\x39\xdf\xd8\x21\x54\x03\x90\x5f\x89\xe0\x0a\xb8\x64\xe5\xe8\xd1\x34\x63\x0e\xac\xe4\x59\xe2\x00\x7c\x4b\xa4\x01\x46\xb1\x8c\x05\xcf\x4a\x8e\x62\x66\x1b\x96\x9f\x3a\xa9\x14\x76\xa4\x80\x5e\xc4\x87\x3d\xff\xb4\xd9\xfd\x54\x1f\x1c\xe8\x0f\xa7\x5a\xac\x5e\x4e\xc3\x86\xff\x0d\x0d\xc0\xde\x07\xf4\xb7\xfa\xee\x8e\x92\x56\xb3\xd5\x00\xdf\x66\xfb\x6c\xfd\x94\x4f\x00\x00\x00\xff\xff\xeb\xf1\x4c\x79\x5e\x01\x00\x00")

func templatesVarsTfBytes() ([]byte, error) {
	return bindataRead(
		_templatesVarsTf,
		"templates/vars.tf",
	)
}

func templatesVarsTf() (*asset, error) {
	bytes, err := templatesVarsTfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/vars.tf", size: 350, mode: os.FileMode(480), modTime: time.Unix(1520375763, 0)}
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
	"templates/bosh_director.tf": templatesBosh_directorTf,
	"templates/cf_dns.tf": templatesCf_dnsTf,
	"templates/cf_lb.tf": templatesCf_lbTf,
	"templates/concourse_lb.tf": templatesConcourse_lbTf,
	"templates/jumpbox.tf": templatesJumpboxTf,
	"templates/vars.tf": templatesVarsTf,
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
	"templates": &bintree{nil, map[string]*bintree{
		"bosh_director.tf": &bintree{templatesBosh_directorTf, map[string]*bintree{}},
		"cf_dns.tf": &bintree{templatesCf_dnsTf, map[string]*bintree{}},
		"cf_lb.tf": &bintree{templatesCf_lbTf, map[string]*bintree{}},
		"concourse_lb.tf": &bintree{templatesConcourse_lbTf, map[string]*bintree{}},
		"jumpbox.tf": &bintree{templatesJumpboxTf, map[string]*bintree{}},
		"vars.tf": &bintree{templatesVarsTf, map[string]*bintree{}},
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

