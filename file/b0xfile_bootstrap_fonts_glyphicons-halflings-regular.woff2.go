// Code generaTed by fileb0x at "2022-08-24 20:50:50.957808 +0800 CST m=+0.116065737" from config file "b0x.yaml" DO NOT EDIT.
// modified(2021-06-28 19:10:38 +0800 CST)
// original path: swagger-bootstrap/bootstrap/fonts/glyphicons-halflings-regular.woff2

package file

import (
	"os"
)

// FileBootstrapFontsGlyphiconsHalflingsRegularWoff2 is "bootstrap/fonts/glyphicons-halflings-regular.woff2"
var FileBootstrapFontsGlyphiconsHalflingsRegularWoff2 = []byte("")

func init() {

	f, err := FS.OpenFile(CTX, "bootstrap/fonts/glyphicons-halflings-regular.woff2", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		panic(err)
	}

	_, err = f.Write(FileBootstrapFontsGlyphiconsHalflingsRegularWoff2)
	if err != nil {
		panic(err)
	}

	err = f.Close()
	if err != nil {
		panic(err)
	}
}
