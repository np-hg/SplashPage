// Code generated for package main by go-bindata DO NOT EDIT. (@generated)
// sources:
// data/index.html
package main

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

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// Mode return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _dataIndexHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x56\x4d\x6f\xe3\x36\x10\xbd\x07\xc8\x7f\x98\x2e\xb0\x70\x82\x86\xb6\x92\xd6\x28\xa2\x6a\x7d\x68\x5a\xa0\x0b\x14\x68\x81\xf4\x92\x23\x25\x8e\xa4\x69\x28\x52\x20\x69\x5b\xae\xe1\xff\x5e\xd0\x62\x1c\x59\x1f\x46\x7a\xd9\x20\xb6\x25\xce\xd3\x7c\x3c\xbe\xe1\x28\xf9\xee\xd7\x3f\x9f\xfe\x7e\xf9\xeb\x37\x28\x5d\x25\x57\xd7\x57\x49\xf8\xf5\x57\xc8\xc5\xea\xfa\x0a\x00\x20\xb1\x6e\x27\x31\xdc\xf8\xbf\xfd\x1e\x28\x87\xf9\xd3\xda\x3a\x5d\x3d\x3d\x3f\xc3\xe1\xf0\x6e\x0c\x80\x09\xeb\x7e\x0f\x28\x2d\x9e\xad\xa5\x5a\xec\x60\x7f\xee\x21\xe5\xd9\x6b\x61\xf4\x5a\x09\x66\xe9\x5f\x8c\xe1\x3e\x8a\x3e\xff\x3c\x89\xc9\xb4\xd4\x26\x06\x53\xa4\x37\x8f\x3f\xde\xc1\xe3\xf2\x0e\x1e\x1f\x6f\x3b\xf8\x83\x2f\xea\xed\x66\x6e\xd1\x6c\x28\x43\x46\x99\x56\x17\x22\x07\xaf\xdb\x92\x1c\xf6\x62\xd7\xda\x92\x23\xad\x62\x30\x28\xb9\xa3\x4d\x1f\x20\xc8\xd6\x92\xef\x62\x20\x25\x49\x21\x4b\xa5\xce\x5e\x7b\x98\x5c\x2b\x17\xca\x7b\x58\xd6\xcd\x98\x35\xe7\x15\xc9\x5d\x0c\x96\x2b\xcb\x2c\x1a\xca\xfb\x24\x68\x23\xd0\x30\xc3\x05\xad\x6d\x0c\x3f\x0c\xdc\x54\xdc\x14\xa4\x3c\x81\x03\x53\x89\x54\x94\x2e\x86\xfb\x91\xe8\xad\xdf\x18\x32\x2e\xb3\x9b\x87\xba\x81\xef\x21\x9a\x3f\x6c\xb6\xb7\x60\xb5\x24\x01\xce\x70\x65\x6b\x6e\x50\xb9\xe9\x6d\xd1\x86\x8e\xb1\x43\x92\xa9\x1e\x44\xe9\x90\x2d\xa9\x8e\x21\xd3\xca\xa1\x72\x1e\x7a\xf7\xa1\xc7\x5a\xfa\x32\xbd\x41\x33\xa8\xa0\xf1\x56\x52\xc5\xa5\x04\x3c\xa8\xe4\x42\x6f\x63\x88\x20\x82\x65\xdd\x1c\x3f\xa6\x48\xf9\x4d\x74\x07\xe1\x7f\xbe\xfc\x98\x98\x56\xbc\x2f\xa7\x77\xa1\xf0\xd4\x6a\xb9\x1e\x28\xc9\xe9\x3a\x86\xe5\x40\xdc\xac\xb2\xec\xc8\x71\xae\x4d\x15\xb7\x74\x4b\xee\xf0\xe5\x86\x2d\xa3\xcf\xb7\x7d\x27\x1f\x46\xb6\x72\x60\x12\x73\x17\xc3\xc3\x59\xd8\x4b\x65\xad\xa8\x2a\xfe\x7f\x69\xd8\x38\xc6\x25\x15\x2a\x86\x0c\x95\x1b\x6c\xd1\x37\xaf\xbd\x61\x5b\x12\xae\x8c\xe1\x7e\xd9\xb6\xc3\x54\xf5\x85\x21\xd1\xaf\xf7\xd4\xd2\xde\xd8\x73\xed\x97\x98\xc3\xaa\xf6\xd1\xfd\xc1\xb1\xae\x94\x6f\xc7\x28\xaa\x9b\xee\xf7\xd8\x63\x2d\x9a\x15\xbc\x1e\x3d\x07\x8e\x18\xa3\xb7\x93\x80\x50\x12\x5f\x3b\xdd\xb3\xf8\xae\xc8\xa5\xd7\x76\x49\x42\xa0\x9a\xdc\x6c\xdf\x76\x9c\x14\x9a\xc9\x9a\x73\x89\x83\x03\x4a\x62\xc3\x04\x19\xcc\x5a\x15\xb4\x75\xf4\x40\xff\xac\xad\xa3\x7c\xc7\x42\x63\x4f\x08\xe1\x28\x12\x46\x0e\x2b\x3b\x81\xa8\x48\xb1\xd3\x79\x35\x9c\x05\x35\x17\x82\x54\xc1\x8e\x8a\xea\x71\x74\x5e\x6a\x29\x0c\x1b\x11\xf3\xa9\xd0\xb1\x83\xfa\xac\x69\x46\x78\x0e\x76\xd3\xa6\x37\x0d\x48\xb5\x73\xba\xba\xb4\x89\x3f\xf5\x44\xd2\x1b\x9d\x4a\x9c\x26\x67\xb2\x78\x1b\xcb\xc9\x22\x0c\xeb\xeb\xab\xc4\xcf\xd2\x30\xa9\xf7\x7b\x3f\xa4\x7f\xe1\x4a\xa1\xf9\x5a\xf1\x02\xbb\x33\x37\x11\xb4\x81\x4c\x72\x6b\xbf\xcc\x3c\x23\x92\xd4\xeb\x6c\x95\x78\x62\x3a\xab\x54\x15\x33\xb0\x26\xfb\xf2\xc9\x4f\xf4\x8e\x2b\x38\x1c\x3e\x2d\x56\xc9\x42\xd0\xe6\x14\xed\x94\x5d\xc8\xaf\x13\xe1\x24\xaf\xd9\x6a\x3c\x05\xaf\xf1\xd9\xaa\xbb\x4f\xc1\xe7\x96\x5c\x09\xf3\xaf\x99\x56\x7f\x90\x7a\xb5\x83\x37\x8d\x00\x33\x5c\x15\x08\xf3\x51\x33\xbc\xbf\xb1\x3c\x97\x7a\x3b\x89\xe9\xe7\xd4\x3d\x00\xbb\x79\x8f\x3e\xc7\xa1\x34\x98\xb7\x3c\xfd\x6e\x30\xf7\x04\x9d\xd1\x49\x55\xd1\x7a\x7a\xe7\xf3\xd9\x64\x6f\x3c\xf2\x0b\xfe\xbb\x2c\x8f\xd4\xd5\xd5\xc4\x07\x4c\x23\xcb\xdd\x00\x6f\xd7\xc9\x22\x28\xe9\x28\x2f\xff\x56\xf8\x5f\x00\x00\x00\xff\xff\x09\x76\xd5\x3d\x2d\x0a\x00\x00")

func dataIndexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_dataIndexHtml,
		"data/index.html",
	)
}

func dataIndexHtml() (*asset, error) {
	bytes, err := dataIndexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/index.html", size: 2605, mode: os.FileMode(438), modTime: time.Unix(1610130443, 0)}
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
	"data/index.html": dataIndexHtml,
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
	"data": &bintree{nil, map[string]*bintree{
		"index.html": &bintree{dataIndexHtml, map[string]*bintree{}},
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
