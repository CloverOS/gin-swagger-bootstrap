// Code generaTed by fileb0x at "2022-08-24 20:50:50.963947 +0800 CST m=+0.122204641" from config file "b0x.yaml" DO NOT EDIT.
// modified(2022-08-24 19:59:39.989 +0800 CST)
// original path: swagger-bootstrap/group.json

package bootstrap

import (
	"os"
)

// FileGroupJSON is "group.json"
var FileGroupJSON = []byte("\x5b\x0a\x20\x20\x7b\x0a\x20\x20\x20\x20\x22\x6e\x61\x6d\x65\x22\x3a\x20\x22\x73\x77\x61\x67\x67\x65\x72\x22\x2c\x0a\x20\x20\x20\x20\x22\x75\x72\x6c\x22\x3a\x20\x22\x2e\x2e\x2f\x73\x77\x61\x67\x67\x65\x72\x2f\x73\x77\x61\x67\x67\x65\x72\x2e\x6a\x73\x6f\x6e\x22\x2c\x0a\x20\x20\x20\x20\x22\x73\x77\x61\x67\x67\x65\x72\x56\x65\x72\x73\x69\x6f\x6e\x22\x3a\x20\x22\x32\x2e\x30\x22\x0a\x20\x20\x7d\x0a\x5d\x0a")

func init() {

	f, err := FS.OpenFile(CTX, "group.json", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		panic(err)
	}

	_, err = f.Write(FileGroupJSON)
	if err != nil {
		panic(err)
	}

	err = f.Close()
	if err != nil {
		panic(err)
	}
}