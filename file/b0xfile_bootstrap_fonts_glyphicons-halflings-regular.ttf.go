// Code generaTed by fileb0x at "2022-08-24 20:50:51.019364 +0800 CST m=+0.177623102" from config file "b0x.yaml" DO NOT EDIT.
// modified(2021-06-28 19:10:38 +0800 CST)
// original path: swagger-bootstrap/bootstrap/fonts/glyphicons-halflings-regular.ttf

package file

import (
	"os"
)

// FileBootstrapFontsGlyphiconsHalflingsRegularTtf is "bootstrap/fonts/glyphicons-halflings-regular.ttf"
var FileBootstrapFontsGlyphiconsHalflingsRegularTtf = []byte("")

func init() {

	f, err := FS.OpenFile(CTX, "bootstrap/fonts/glyphicons-halflings-regular.ttf", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		panic(err)
	}

	_, err = f.Write(FileBootstrapFontsGlyphiconsHalflingsRegularTtf)
	if err != nil {
		panic(err)
	}

	err = f.Close()
	if err != nil {
		panic(err)
	}
}
