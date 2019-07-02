// Code generated for package templates by go-bindata DO NOT EDIT. (@generated)
// sources:
// templates/fs.go
// templates/module.go.tpl
// templates/resolver.go.tpl
// templates/schema.graphql
package templates

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

var _templatesFsGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func templatesFsGoBytes() ([]byte, error) {
	return bindataRead(
		_templatesFsGo,
		"templates/fs.go",
	)
}

func templatesFsGo() (*asset, error) {
	bytes, err := templatesFsGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/fs.go", size: 0, mode: os.FileMode(438), modTime: time.Unix(1562060847, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesModuleGoTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x8e\xc1\x6e\xea\x30\x10\x45\xd7\x99\xaf\x18\x65\x65\xa3\x27\x67\xcd\x5b\x16\x58\xb0\x28\x54\x2d\x3f\x60\xec\xc1\xb8\x75\xec\x30\xb1\xa1\x08\xe5\xdf\xab\x14\x22\x55\x95\xba\xb1\x3c\x73\x74\xef\x99\xa6\xc1\x45\xb2\x84\x8e\x22\xb1\xce\x64\x71\x7f\xc5\x43\xd0\xad\x8f\x2e\xa9\x96\x1a\xc7\xba\x3b\x9e\xc2\x3f\x5c\x6e\x71\xb3\xdd\xe1\x6a\xb9\xde\x29\x80\x4e\x9b\x0f\xed\x08\x1f\x18\xc0\xb7\x5d\xe2\x8c\x02\xaa\xfa\x67\xdc\x8e\x9f\x1a\xaa\xda\xf9\x7c\x2c\x7b\x65\x52\xdb\xcc\xe7\x96\x7a\xef\x62\xdf\xb8\x53\x70\x14\x27\x47\x0d\x12\x20\x5f\x3b\xc2\xe7\x64\x4b\x20\xec\x33\x17\x93\x6f\x03\xc0\xa1\x44\x83\x62\x76\xdf\x4b\x5c\xa4\x78\xf0\xae\x30\x09\x1f\xdf\xc9\xe4\xc4\x38\xfb\x36\xa9\xf5\x63\x96\x78\x83\x6a\x82\xea\xc9\x47\x2b\x22\x5d\xc4\xc3\xa4\x56\x9f\x64\x4a\xd6\xfb\x40\x6f\xe6\x48\xad\x96\x52\xed\xd2\x0b\xa7\xb3\xb7\xc4\x62\xb4\x09\x4e\x29\xe3\x6c\x7c\x5f\xa9\x4f\xe1\x4c\x2c\xf1\xaf\xf8\x28\xab\x98\x72\xe1\x88\x1b\xba\xfc\xc6\xe2\x7e\xef\x6d\x2a\xea\xff\xe3\xd8\x3b\x48\xa8\x06\x09\xc3\x57\x00\x00\x00\xff\xff\x5d\x33\xce\x9f\x84\x01\x00\x00")

func templatesModuleGoTplBytes() ([]byte, error) {
	return bindataRead(
		_templatesModuleGoTpl,
		"templates/module.go.tpl",
	)
}

func templatesModuleGoTpl() (*asset, error) {
	bytes, err := templatesModuleGoTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/module.go.tpl", size: 388, mode: os.FileMode(438), modTime: time.Unix(1562055432, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesResolverGoTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x8f\x41\x4e\x03\x31\x0c\x45\xf7\x3e\xc5\x67\x37\xb0\xe9\x09\x38\x02\x0b\x7a\x01\x14\x82\x67\xc6\x10\xec\xd4\x71\x5a\x55\x55\xef\x8e\xa8\xda\x42\x2b\xcd\xca\x96\xbf\xde\xb3\x7e\x4d\xf9\x2b\x4d\x8c\xc9\x53\x9d\x37\x85\x68\xb5\x42\xcc\xd2\x30\x4a\x61\x48\x43\x42\x8b\xe4\x21\x3a\xa1\x9a\x68\x60\x34\x47\xcc\x8c\xea\xf6\xc9\x39\xd0\x2a\x67\x19\x25\xc3\xb9\x59\xd9\xb2\xb7\x5f\x87\x04\x76\x52\x0a\xd4\x02\xef\x0c\xe7\x89\x95\x3d\x05\x7f\x3c\x10\xc5\xbe\x32\xdc\x2c\xd6\x67\x04\x68\xe1\x3d\xc7\xe1\x48\xb4\x4d\x8e\x37\x5c\x92\xb5\x59\xe0\x19\xca\xbb\xe1\x3f\xf0\x48\x34\x76\xcd\x18\x9e\x6e\xae\x78\xed\xec\xfb\xe1\x3c\xaf\xf6\x03\x39\x47\x77\x85\x4a\xa1\xe3\x02\xfa\xd2\x23\x85\x98\x0e\x7f\xeb\xa2\xe0\x54\x60\x73\xf3\xe3\x5a\xe0\x94\x7d\xdf\x1b\x2e\xf1\x4f\x00\x00\x00\xff\xff\xd6\xae\x57\x56\x70\x01\x00\x00")

func templatesResolverGoTplBytes() ([]byte, error) {
	return bindataRead(
		_templatesResolverGoTpl,
		"templates/resolver.go.tpl",
	)
}

func templatesResolverGoTpl() (*asset, error) {
	bytes, err := templatesResolverGoTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/resolver.go.tpl", size: 368, mode: os.FileMode(438), modTime: time.Unix(1562055432, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesSchemaGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\xa9\x2c\x48\x55\x08\x2c\x4d\x2d\xaa\x54\xa8\x56\x48\xcb\x49\xcc\xcd\xcc\x4b\xcf\xb7\x52\x08\x2e\x29\xca\xcc\x4b\x57\xa8\xe5\x02\xcb\xfb\x96\x96\x24\x96\x64\xe6\xe7\x61\x57\x52\x9c\x9c\x98\x93\x58\xa4\x10\x92\x99\x9b\x0a\x63\xfb\x26\x16\x00\x02\x00\x00\xff\xff\x60\x5b\x70\x13\x59\x00\x00\x00")

func templatesSchemaGraphqlBytes() ([]byte, error) {
	return bindataRead(
		_templatesSchemaGraphql,
		"templates/schema.graphql",
	)
}

func templatesSchemaGraphql() (*asset, error) {
	bytes, err := templatesSchemaGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/schema.graphql", size: 89, mode: os.FileMode(438), modTime: time.Unix(1562055432, 0)}
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
	"templates/fs.go":           templatesFsGo,
	"templates/module.go.tpl":   templatesModuleGoTpl,
	"templates/resolver.go.tpl": templatesResolverGoTpl,
	"templates/schema.graphql":  templatesSchemaGraphql,
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
		"fs.go":           &bintree{templatesFsGo, map[string]*bintree{}},
		"module.go.tpl":   &bintree{templatesModuleGoTpl, map[string]*bintree{}},
		"resolver.go.tpl": &bintree{templatesResolverGoTpl, map[string]*bintree{}},
		"schema.graphql":  &bintree{templatesSchemaGraphql, map[string]*bintree{}},
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
