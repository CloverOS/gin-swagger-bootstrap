// Code generated by fileb0x at "2022-08-24 20:50:50.897295 +0800 CST m=+0.055551102" from config file "b0x.yaml" DO NOT EDIT.
// modification hash(2e51b7103fbfb00f6415e63d30f5a435.29b50aea162ce1dda98bfb15f6263dde)

package file

import (
	"bytes"

	"context"
	"io"
	"net/http"
	"os"
	"path"

	"golang.org/x/net/webdav"
)

var (
	// CTX is a context for webdav vfs
	CTX = context.Background()

	// FS is a virtual memory file system
	FS = webdav.NewMemFS()

	// Handler is used to server files through a http handler
	Handler *webdav.Handler

	// HTTP is the http file system
	HTTP http.FileSystem = new(HTTPFS)
)

// HTTPFS implements http.FileSystem
type HTTPFS struct {
	// Prefix allows to limit the path of all requests. F.e. a prefix "css" would allow only calls to /css/*
	Prefix string
}

func init() {
	err := CTX.Err()
	if err != nil {
		panic(err)
	}

	err = FS.Mkdir(CTX, "layer3.0.3/", 0777)
	if err != nil && err != os.ErrExist {
		panic(err)
	}

	err = FS.Mkdir(CTX, "layer3.0.3/skin/", 0777)
	if err != nil && err != os.ErrExist {
		panic(err)
	}

	err = FS.Mkdir(CTX, "layer3.0.3/skin/default/", 0777)
	if err != nil && err != os.ErrExist {
		panic(err)
	}

	err = FS.Mkdir(CTX, "layer3.0.3/mobile/", 0777)
	if err != nil && err != os.ErrExist {
		panic(err)
	}

	err = FS.Mkdir(CTX, "layer3.0.3/mobile/need/", 0777)
	if err != nil && err != os.ErrExist {
		panic(err)
	}

	err = FS.Mkdir(CTX, "ace-editor/", 0777)
	if err != nil && err != os.ErrExist {
		panic(err)
	}

	err = FS.Mkdir(CTX, "ace-editor/snippets/", 0777)
	if err != nil && err != os.ErrExist {
		panic(err)
	}

	err = FS.Mkdir(CTX, "bootstrap/", 0777)
	if err != nil && err != os.ErrExist {
		panic(err)
	}

	err = FS.Mkdir(CTX, "bootstrap/css/", 0777)
	if err != nil && err != os.ErrExist {
		panic(err)
	}

	err = FS.Mkdir(CTX, "bootstrap/js/", 0777)
	if err != nil && err != os.ErrExist {
		panic(err)
	}

	err = FS.Mkdir(CTX, "bootstrap/fonts/", 0777)
	if err != nil && err != os.ErrExist {
		panic(err)
	}

	err = FS.Mkdir(CTX, "highlight/", 0777)
	if err != nil && err != os.ErrExist {
		panic(err)
	}

	err = FS.Mkdir(CTX, "highlight/styles/", 0777)
	if err != nil && err != os.ErrExist {
		panic(err)
	}

	err = FS.Mkdir(CTX, "images/", 0777)
	if err != nil && err != os.ErrExist {
		panic(err)
	}

	err = FS.Mkdir(CTX, "layui/", 0777)
	if err != nil && err != os.ErrExist {
		panic(err)
	}

	err = FS.Mkdir(CTX, "layui/css/", 0777)
	if err != nil && err != os.ErrExist {
		panic(err)
	}

	err = FS.Mkdir(CTX, "layui/css/modules/", 0777)
	if err != nil && err != os.ErrExist {
		panic(err)
	}

	err = FS.Mkdir(CTX, "layui/css/modules/layer/", 0777)
	if err != nil && err != os.ErrExist {
		panic(err)
	}

	err = FS.Mkdir(CTX, "layui/css/modules/layer/default/", 0777)
	if err != nil && err != os.ErrExist {
		panic(err)
	}

	err = FS.Mkdir(CTX, "layui/lay/", 0777)
	if err != nil && err != os.ErrExist {
		panic(err)
	}

	err = FS.Mkdir(CTX, "layui/lay/modules/", 0777)
	if err != nil && err != os.ErrExist {
		panic(err)
	}

	err = FS.Mkdir(CTX, "layui/font/", 0777)
	if err != nil && err != os.ErrExist {
		panic(err)
	}

	err = FS.Mkdir(CTX, "ext/", 0777)
	if err != nil && err != os.ErrExist {
		panic(err)
	}

	err = FS.Mkdir(CTX, "ext/treetable-lay/", 0777)
	if err != nil && err != os.ErrExist {
		panic(err)
	}

	err = FS.Mkdir(CTX, "jquery/", 0777)
	if err != nil && err != os.ErrExist {
		panic(err)
	}

	err = FS.Mkdir(CTX, "jquery/clipboard/", 0777)
	if err != nil && err != os.ErrExist {
		panic(err)
	}

	err = FS.Mkdir(CTX, "iconfont/", 0777)
	if err != nil && err != os.ErrExist {
		panic(err)
	}

	err = FS.Mkdir(CTX, "cdao/", 0777)
	if err != nil && err != os.ErrExist {
		panic(err)
	}

	err = FS.Mkdir(CTX, "ace/", 0777)
	if err != nil && err != os.ErrExist {
		panic(err)
	}

	err = FS.Mkdir(CTX, "bootstrap-tabx/", 0777)
	if err != nil && err != os.ErrExist {
		panic(err)
	}

	err = FS.Mkdir(CTX, "bootstrap-tabx/css/", 0777)
	if err != nil && err != os.ErrExist {
		panic(err)
	}

	err = FS.Mkdir(CTX, "bootstrap-tabx/js/", 0777)
	if err != nil && err != os.ErrExist {
		panic(err)
	}

	err = FS.Mkdir(CTX, "bootstrap-tabx/img/", 0777)
	if err != nil && err != os.ErrExist {
		panic(err)
	}

	Handler = &webdav.Handler{
		FileSystem: FS,
		LockSystem: webdav.NewMemLS(),
	}

}

// Open a file
func (hfs *HTTPFS) Open(path string) (http.File, error) {
	path = hfs.Prefix + path

	f, err := FS.OpenFile(CTX, path, os.O_RDONLY, 0644)
	if err != nil {
		return nil, err
	}

	return f, nil
}

// ReadFile is adapTed from ioutil
func ReadFile(path string) ([]byte, error) {
	f, err := FS.OpenFile(CTX, path, os.O_RDONLY, 0644)
	if err != nil {
		return nil, err
	}

	buf := bytes.NewBuffer(make([]byte, 0, bytes.MinRead))

	// If the buffer overflows, we will get bytes.ErrTooLarge.
	// Return that as an error. Any other panic remains.
	defer func() {
		e := recover()
		if e == nil {
			return
		}
		if panicErr, ok := e.(error); ok && panicErr == bytes.ErrTooLarge {
			err = panicErr
		} else {
			panic(e)
		}
	}()
	_, err = buf.ReadFrom(f)
	return buf.Bytes(), err
}

// WriteFile is adapTed from ioutil
func WriteFile(filename string, data []byte, perm os.FileMode) error {
	f, err := FS.OpenFile(CTX, filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, perm)
	if err != nil {
		return err
	}
	n, err := f.Write(data)
	if err == nil && n < len(data) {
		err = io.ErrShortWrite
	}
	if err1 := f.Close(); err == nil {
		err = err1
	}
	return err
}

// WalkDirs looks for files in the given dir and returns a list of files in it
// usage for all files in the b0x: WalkDirs("", false)
func WalkDirs(name string, includeDirsInList bool, files ...string) ([]string, error) {
	f, err := FS.OpenFile(CTX, name, os.O_RDONLY, 0)
	if err != nil {
		return nil, err
	}

	fileInfos, err := f.Readdir(0)
	if err != nil {
		return nil, err
	}

	err = f.Close()
	if err != nil {
		return nil, err
	}

	for _, info := range fileInfos {
		filename := path.Join(name, info.Name())

		if includeDirsInList || !info.IsDir() {
			files = append(files, filename)
		}

		if info.IsDir() {
			files, err = WalkDirs(filename, includeDirsInList, files...)
			if err != nil {
				return nil, err
			}
		}
	}

	return files, nil
}
