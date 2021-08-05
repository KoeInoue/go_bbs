// Code generated for package main by go-bindata DO NOT EDIT. (@generated)
// sources:
// assets/build/asset-manifest.json
// assets/build/css/2.4608ec59.chunk.css
// assets/build/css/2.4608ec59.chunk.css.map
// assets/build/css/2.8840a649.chunk.css
// assets/build/css/2.8840a649.chunk.css.map
// assets/build/css/main.5be27138.chunk.css
// assets/build/css/main.5be27138.chunk.css.map
// assets/build/css/main.c19d1c86.chunk.css
// assets/build/css/main.c19d1c86.chunk.css.map
// assets/build/favicon.ico
// assets/build/index.html
// assets/build/js/2.30a5f63c.chunk.js
// assets/build/js/2.30a5f63c.chunk.js.map
// assets/build/js/main.093370dd.chunk.js
// assets/build/js/main.093370dd.chunk.js.map
// assets/build/js/main.90e62041.chunk.js
// assets/build/js/main.90e62041.chunk.js.map
// assets/build/js/runtime-main.0d813b15.js
// assets/build/js/runtime-main.0d813b15.js.map
// assets/build/js/runtime-main.4a474bcd.js
// assets/build/js/runtime-main.4a474bcd.js.map
// assets/build/logo192.png
// assets/build/logo512.png
// assets/build/manifest.json
// assets/build/media/fa-brands-400.099a9556.woff
// assets/build/media/fa-brands-400.30cc681d.eot
// assets/build/media/fa-brands-400.3b89dd10.ttf
// assets/build/media/fa-brands-400.ba7ed552.svg
// assets/build/media/fa-brands-400.f7307680.woff2
// assets/build/media/fa-regular-400.0bb42845.svg
// assets/build/media/fa-regular-400.1f77739c.ttf
// assets/build/media/fa-regular-400.7124eb50.woff
// assets/build/media/fa-regular-400.7630483d.eot
// assets/build/media/fa-regular-400.f0f82301.woff2
// assets/build/media/fa-solid-900.1042e8ca.eot
// assets/build/media/fa-solid-900.376c1f97.svg
// assets/build/media/fa-solid-900.605ed792.ttf
// assets/build/media/fa-solid-900.9fe5a17c.woff
// assets/build/media/fa-solid-900.e8a427e1.woff2
// assets/build/media/login.db9aa21e.png
// assets/build/precache-manifest.63f194368b5e14fc8d9700498338ad87.js
// assets/build/precache-manifest.ae6b8c83066240b96bd573fcc5f7e0da.js
// assets/build/precache-manifest.c5fc2df4798feaf42383f744c6942a0a.js
// assets/build/robots.txt
// assets/build/service-worker.js
// assets/css/main.css
// assets/css/main.css.map
// assets/favicon.ico
// assets/img/books.png
// assets/img/data.svg
// assets/img/facebook.svg
// assets/img/googleplay.png
// assets/img/ios.png
// assets/img/lesson.png
// assets/img/mail-send.svg
// assets/img/mentioned.svg
// assets/img/notification.svg
// assets/img/progress.svg
// assets/img/quiz.svg
// assets/img/reading.svg
// assets/img/security.svg
// assets/img/social-login.svg
// assets/img/storage.svg
// assets/img/study-data.png
// assets/img/text.svg
// assets/img/twitter.svg
// assets/img/youtube.svg
// assets/index.html
// assets/js/production.js
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// bindataRead reads the given file from disk. It returns an error on failure.
func bindataRead(path, name string) ([]byte, error) {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset %s at %s: %v", name, path, err)
	}
	return buf, err
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

// assetsBuildAssetManifestJson reads file data from disk. It returns an error on failure.
func assetsBuildAssetManifestJson() (*asset, error) {
	path := "/go/src/app/assets/build/asset-manifest.json"
	name := "assets/build/asset-manifest.json"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assetsBuildCss24608ec59ChunkCss reads file data from disk. It returns an error on failure.
func assetsBuildCss24608ec59ChunkCss() (*asset, error) {
	path := "/go/src/app/assets/build/css/2.4608ec59.chunk.css"
	name := "assets/build/css/2.4608ec59.chunk.css"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assetsBuildCss24608ec59ChunkCssMap reads file data from disk. It returns an error on failure.
func assetsBuildCss24608ec59ChunkCssMap() (*asset, error) {
	path := "/go/src/app/assets/build/css/2.4608ec59.chunk.css.map"
	name := "assets/build/css/2.4608ec59.chunk.css.map"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assetsBuildCss28840a649ChunkCss reads file data from disk. It returns an error on failure.
func assetsBuildCss28840a649ChunkCss() (*asset, error) {
	path := "/go/src/app/assets/build/css/2.8840a649.chunk.css"
	name := "assets/build/css/2.8840a649.chunk.css"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assetsBuildCss28840a649ChunkCssMap reads file data from disk. It returns an error on failure.
func assetsBuildCss28840a649ChunkCssMap() (*asset, error) {
	path := "/go/src/app/assets/build/css/2.8840a649.chunk.css.map"
	name := "assets/build/css/2.8840a649.chunk.css.map"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assetsBuildCssMain5be27138ChunkCss reads file data from disk. It returns an error on failure.
func assetsBuildCssMain5be27138ChunkCss() (*asset, error) {
	path := "/go/src/app/assets/build/css/main.5be27138.chunk.css"
	name := "assets/build/css/main.5be27138.chunk.css"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assetsBuildCssMain5be27138ChunkCssMap reads file data from disk. It returns an error on failure.
func assetsBuildCssMain5be27138ChunkCssMap() (*asset, error) {
	path := "/go/src/app/assets/build/css/main.5be27138.chunk.css.map"
	name := "assets/build/css/main.5be27138.chunk.css.map"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assetsBuildCssMainC19d1c86ChunkCss reads file data from disk. It returns an error on failure.
func assetsBuildCssMainC19d1c86ChunkCss() (*asset, error) {
	path := "/go/src/app/assets/build/css/main.c19d1c86.chunk.css"
	name := "assets/build/css/main.c19d1c86.chunk.css"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assetsBuildCssMainC19d1c86ChunkCssMap reads file data from disk. It returns an error on failure.
func assetsBuildCssMainC19d1c86ChunkCssMap() (*asset, error) {
	path := "/go/src/app/assets/build/css/main.c19d1c86.chunk.css.map"
	name := "assets/build/css/main.c19d1c86.chunk.css.map"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assetsBuildFaviconIco reads file data from disk. It returns an error on failure.
func assetsBuildFaviconIco() (*asset, error) {
	path := "/go/src/app/assets/build/favicon.ico"
	name := "assets/build/favicon.ico"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assetsBuildIndexHtml reads file data from disk. It returns an error on failure.
func assetsBuildIndexHtml() (*asset, error) {
	path := "/go/src/app/assets/build/index.html"
	name := "assets/build/index.html"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assetsBuildJs230a5f63cChunkJs reads file data from disk. It returns an error on failure.
func assetsBuildJs230a5f63cChunkJs() (*asset, error) {
	path := "/go/src/app/assets/build/js/2.30a5f63c.chunk.js"
	name := "assets/build/js/2.30a5f63c.chunk.js"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assetsBuildJs230a5f63cChunkJsMap reads file data from disk. It returns an error on failure.
func assetsBuildJs230a5f63cChunkJsMap() (*asset, error) {
	path := "/go/src/app/assets/build/js/2.30a5f63c.chunk.js.map"
	name := "assets/build/js/2.30a5f63c.chunk.js.map"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assetsBuildJsMain093370ddChunkJs reads file data from disk. It returns an error on failure.
func assetsBuildJsMain093370ddChunkJs() (*asset, error) {
	path := "/go/src/app/assets/build/js/main.093370dd.chunk.js"
	name := "assets/build/js/main.093370dd.chunk.js"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assetsBuildJsMain093370ddChunkJsMap reads file data from disk. It returns an error on failure.
func assetsBuildJsMain093370ddChunkJsMap() (*asset, error) {
	path := "/go/src/app/assets/build/js/main.093370dd.chunk.js.map"
	name := "assets/build/js/main.093370dd.chunk.js.map"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assetsBuildJsMain90e62041ChunkJs reads file data from disk. It returns an error on failure.
func assetsBuildJsMain90e62041ChunkJs() (*asset, error) {
	path := "/go/src/app/assets/build/js/main.90e62041.chunk.js"
	name := "assets/build/js/main.90e62041.chunk.js"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assetsBuildJsMain90e62041ChunkJsMap reads file data from disk. It returns an error on failure.
func assetsBuildJsMain90e62041ChunkJsMap() (*asset, error) {
	path := "/go/src/app/assets/build/js/main.90e62041.chunk.js.map"
	name := "assets/build/js/main.90e62041.chunk.js.map"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assetsBuildJsRuntimeMain0d813b15Js reads file data from disk. It returns an error on failure.
func assetsBuildJsRuntimeMain0d813b15Js() (*asset, error) {
	path := "/go/src/app/assets/build/js/runtime-main.0d813b15.js"
	name := "assets/build/js/runtime-main.0d813b15.js"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assetsBuildJsRuntimeMain0d813b15JsMap reads file data from disk. It returns an error on failure.
func assetsBuildJsRuntimeMain0d813b15JsMap() (*asset, error) {
	path := "/go/src/app/assets/build/js/runtime-main.0d813b15.js.map"
	name := "assets/build/js/runtime-main.0d813b15.js.map"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assetsBuildJsRuntimeMain4a474bcdJs reads file data from disk. It returns an error on failure.
func assetsBuildJsRuntimeMain4a474bcdJs() (*asset, error) {
	path := "/go/src/app/assets/build/js/runtime-main.4a474bcd.js"
	name := "assets/build/js/runtime-main.4a474bcd.js"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assetsBuildJsRuntimeMain4a474bcdJsMap reads file data from disk. It returns an error on failure.
func assetsBuildJsRuntimeMain4a474bcdJsMap() (*asset, error) {
	path := "/go/src/app/assets/build/js/runtime-main.4a474bcd.js.map"
	name := "assets/build/js/runtime-main.4a474bcd.js.map"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assetsBuildLogo192Png reads file data from disk. It returns an error on failure.
func assetsBuildLogo192Png() (*asset, error) {
	path := "/go/src/app/assets/build/logo192.png"
	name := "assets/build/logo192.png"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assetsBuildLogo512Png reads file data from disk. It returns an error on failure.
func assetsBuildLogo512Png() (*asset, error) {
	path := "/go/src/app/assets/build/logo512.png"
	name := "assets/build/logo512.png"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assetsBuildManifestJson reads file data from disk. It returns an error on failure.
func assetsBuildManifestJson() (*asset, error) {
	path := "/go/src/app/assets/build/manifest.json"
	name := "assets/build/manifest.json"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assetsBuildMediaFaBrands400099a9556Woff reads file data from disk. It returns an error on failure.
func assetsBuildMediaFaBrands400099a9556Woff() (*asset, error) {
	path := "/go/src/app/assets/build/media/fa-brands-400.099a9556.woff"
	name := "assets/build/media/fa-brands-400.099a9556.woff"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assetsBuildMediaFaBrands40030cc681dEot reads file data from disk. It returns an error on failure.
func assetsBuildMediaFaBrands40030cc681dEot() (*asset, error) {
	path := "/go/src/app/assets/build/media/fa-brands-400.30cc681d.eot"
	name := "assets/build/media/fa-brands-400.30cc681d.eot"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assetsBuildMediaFaBrands4003b89dd10Ttf reads file data from disk. It returns an error on failure.
func assetsBuildMediaFaBrands4003b89dd10Ttf() (*asset, error) {
	path := "/go/src/app/assets/build/media/fa-brands-400.3b89dd10.ttf"
	name := "assets/build/media/fa-brands-400.3b89dd10.ttf"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assetsBuildMediaFaBrands400Ba7ed552Svg reads file data from disk. It returns an error on failure.
func assetsBuildMediaFaBrands400Ba7ed552Svg() (*asset, error) {
	path := "/go/src/app/assets/build/media/fa-brands-400.ba7ed552.svg"
	name := "assets/build/media/fa-brands-400.ba7ed552.svg"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assetsBuildMediaFaBrands400F7307680Woff2 reads file data from disk. It returns an error on failure.
func assetsBuildMediaFaBrands400F7307680Woff2() (*asset, error) {
	path := "/go/src/app/assets/build/media/fa-brands-400.f7307680.woff2"
	name := "assets/build/media/fa-brands-400.f7307680.woff2"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assetsBuildMediaFaRegular4000bb42845Svg reads file data from disk. It returns an error on failure.
func assetsBuildMediaFaRegular4000bb42845Svg() (*asset, error) {
	path := "/go/src/app/assets/build/media/fa-regular-400.0bb42845.svg"
	name := "assets/build/media/fa-regular-400.0bb42845.svg"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assetsBuildMediaFaRegular4001f77739cTtf reads file data from disk. It returns an error on failure.
func assetsBuildMediaFaRegular4001f77739cTtf() (*asset, error) {
	path := "/go/src/app/assets/build/media/fa-regular-400.1f77739c.ttf"
	name := "assets/build/media/fa-regular-400.1f77739c.ttf"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assetsBuildMediaFaRegular4007124eb50Woff reads file data from disk. It returns an error on failure.
func assetsBuildMediaFaRegular4007124eb50Woff() (*asset, error) {
	path := "/go/src/app/assets/build/media/fa-regular-400.7124eb50.woff"
	name := "assets/build/media/fa-regular-400.7124eb50.woff"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assetsBuildMediaFaRegular4007630483dEot reads file data from disk. It returns an error on failure.
func assetsBuildMediaFaRegular4007630483dEot() (*asset, error) {
	path := "/go/src/app/assets/build/media/fa-regular-400.7630483d.eot"
	name := "assets/build/media/fa-regular-400.7630483d.eot"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assetsBuildMediaFaRegular400F0f82301Woff2 reads file data from disk. It returns an error on failure.
func assetsBuildMediaFaRegular400F0f82301Woff2() (*asset, error) {
	path := "/go/src/app/assets/build/media/fa-regular-400.f0f82301.woff2"
	name := "assets/build/media/fa-regular-400.f0f82301.woff2"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assetsBuildMediaFaSolid9001042e8caEot reads file data from disk. It returns an error on failure.
func assetsBuildMediaFaSolid9001042e8caEot() (*asset, error) {
	path := "/go/src/app/assets/build/media/fa-solid-900.1042e8ca.eot"
	name := "assets/build/media/fa-solid-900.1042e8ca.eot"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assetsBuildMediaFaSolid900376c1f97Svg reads file data from disk. It returns an error on failure.
func assetsBuildMediaFaSolid900376c1f97Svg() (*asset, error) {
	path := "/go/src/app/assets/build/media/fa-solid-900.376c1f97.svg"
	name := "assets/build/media/fa-solid-900.376c1f97.svg"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assetsBuildMediaFaSolid900605ed792Ttf reads file data from disk. It returns an error on failure.
func assetsBuildMediaFaSolid900605ed792Ttf() (*asset, error) {
	path := "/go/src/app/assets/build/media/fa-solid-900.605ed792.ttf"
	name := "assets/build/media/fa-solid-900.605ed792.ttf"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assetsBuildMediaFaSolid9009fe5a17cWoff reads file data from disk. It returns an error on failure.
func assetsBuildMediaFaSolid9009fe5a17cWoff() (*asset, error) {
	path := "/go/src/app/assets/build/media/fa-solid-900.9fe5a17c.woff"
	name := "assets/build/media/fa-solid-900.9fe5a17c.woff"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assetsBuildMediaFaSolid900E8a427e1Woff2 reads file data from disk. It returns an error on failure.
func assetsBuildMediaFaSolid900E8a427e1Woff2() (*asset, error) {
	path := "/go/src/app/assets/build/media/fa-solid-900.e8a427e1.woff2"
	name := "assets/build/media/fa-solid-900.e8a427e1.woff2"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assetsBuildMediaLoginDb9aa21ePng reads file data from disk. It returns an error on failure.
func assetsBuildMediaLoginDb9aa21ePng() (*asset, error) {
	path := "/go/src/app/assets/build/media/login.db9aa21e.png"
	name := "assets/build/media/login.db9aa21e.png"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assetsBuildPrecacheManifest63f194368b5e14fc8d9700498338ad87Js reads file data from disk. It returns an error on failure.
func assetsBuildPrecacheManifest63f194368b5e14fc8d9700498338ad87Js() (*asset, error) {
	path := "/go/src/app/assets/build/precache-manifest.63f194368b5e14fc8d9700498338ad87.js"
	name := "assets/build/precache-manifest.63f194368b5e14fc8d9700498338ad87.js"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assetsBuildPrecacheManifestAe6b8c83066240b96bd573fcc5f7e0daJs reads file data from disk. It returns an error on failure.
func assetsBuildPrecacheManifestAe6b8c83066240b96bd573fcc5f7e0daJs() (*asset, error) {
	path := "/go/src/app/assets/build/precache-manifest.ae6b8c83066240b96bd573fcc5f7e0da.js"
	name := "assets/build/precache-manifest.ae6b8c83066240b96bd573fcc5f7e0da.js"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assetsBuildPrecacheManifestC5fc2df4798feaf42383f744c6942a0aJs reads file data from disk. It returns an error on failure.
func assetsBuildPrecacheManifestC5fc2df4798feaf42383f744c6942a0aJs() (*asset, error) {
	path := "/go/src/app/assets/build/precache-manifest.c5fc2df4798feaf42383f744c6942a0a.js"
	name := "assets/build/precache-manifest.c5fc2df4798feaf42383f744c6942a0a.js"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assetsBuildRobotsTxt reads file data from disk. It returns an error on failure.
func assetsBuildRobotsTxt() (*asset, error) {
	path := "/go/src/app/assets/build/robots.txt"
	name := "assets/build/robots.txt"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assetsBuildServiceWorkerJs reads file data from disk. It returns an error on failure.
func assetsBuildServiceWorkerJs() (*asset, error) {
	path := "/go/src/app/assets/build/service-worker.js"
	name := "assets/build/service-worker.js"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assetsCssMainCss reads file data from disk. It returns an error on failure.
func assetsCssMainCss() (*asset, error) {
	path := "/go/src/app/assets/css/main.css"
	name := "assets/css/main.css"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assetsCssMainCssMap reads file data from disk. It returns an error on failure.
func assetsCssMainCssMap() (*asset, error) {
	path := "/go/src/app/assets/css/main.css.map"
	name := "assets/css/main.css.map"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assetsFaviconIco reads file data from disk. It returns an error on failure.
func assetsFaviconIco() (*asset, error) {
	path := "/go/src/app/assets/favicon.ico"
	name := "assets/favicon.ico"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assetsImgBooksPng reads file data from disk. It returns an error on failure.
func assetsImgBooksPng() (*asset, error) {
	path := "/go/src/app/assets/img/books.png"
	name := "assets/img/books.png"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assetsImgDataSvg reads file data from disk. It returns an error on failure.
func assetsImgDataSvg() (*asset, error) {
	path := "/go/src/app/assets/img/data.svg"
	name := "assets/img/data.svg"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assetsImgFacebookSvg reads file data from disk. It returns an error on failure.
func assetsImgFacebookSvg() (*asset, error) {
	path := "/go/src/app/assets/img/facebook.svg"
	name := "assets/img/facebook.svg"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assetsImgGoogleplayPng reads file data from disk. It returns an error on failure.
func assetsImgGoogleplayPng() (*asset, error) {
	path := "/go/src/app/assets/img/googleplay.png"
	name := "assets/img/googleplay.png"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assetsImgIosPng reads file data from disk. It returns an error on failure.
func assetsImgIosPng() (*asset, error) {
	path := "/go/src/app/assets/img/ios.png"
	name := "assets/img/ios.png"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assetsImgLessonPng reads file data from disk. It returns an error on failure.
func assetsImgLessonPng() (*asset, error) {
	path := "/go/src/app/assets/img/lesson.png"
	name := "assets/img/lesson.png"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assetsImgMailSendSvg reads file data from disk. It returns an error on failure.
func assetsImgMailSendSvg() (*asset, error) {
	path := "/go/src/app/assets/img/mail-send.svg"
	name := "assets/img/mail-send.svg"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assetsImgMentionedSvg reads file data from disk. It returns an error on failure.
func assetsImgMentionedSvg() (*asset, error) {
	path := "/go/src/app/assets/img/mentioned.svg"
	name := "assets/img/mentioned.svg"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assetsImgNotificationSvg reads file data from disk. It returns an error on failure.
func assetsImgNotificationSvg() (*asset, error) {
	path := "/go/src/app/assets/img/notification.svg"
	name := "assets/img/notification.svg"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assetsImgProgressSvg reads file data from disk. It returns an error on failure.
func assetsImgProgressSvg() (*asset, error) {
	path := "/go/src/app/assets/img/progress.svg"
	name := "assets/img/progress.svg"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assetsImgQuizSvg reads file data from disk. It returns an error on failure.
func assetsImgQuizSvg() (*asset, error) {
	path := "/go/src/app/assets/img/quiz.svg"
	name := "assets/img/quiz.svg"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assetsImgReadingSvg reads file data from disk. It returns an error on failure.
func assetsImgReadingSvg() (*asset, error) {
	path := "/go/src/app/assets/img/reading.svg"
	name := "assets/img/reading.svg"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assetsImgSecuritySvg reads file data from disk. It returns an error on failure.
func assetsImgSecuritySvg() (*asset, error) {
	path := "/go/src/app/assets/img/security.svg"
	name := "assets/img/security.svg"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assetsImgSocialLoginSvg reads file data from disk. It returns an error on failure.
func assetsImgSocialLoginSvg() (*asset, error) {
	path := "/go/src/app/assets/img/social-login.svg"
	name := "assets/img/social-login.svg"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assetsImgStorageSvg reads file data from disk. It returns an error on failure.
func assetsImgStorageSvg() (*asset, error) {
	path := "/go/src/app/assets/img/storage.svg"
	name := "assets/img/storage.svg"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assetsImgStudyDataPng reads file data from disk. It returns an error on failure.
func assetsImgStudyDataPng() (*asset, error) {
	path := "/go/src/app/assets/img/study-data.png"
	name := "assets/img/study-data.png"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assetsImgTextSvg reads file data from disk. It returns an error on failure.
func assetsImgTextSvg() (*asset, error) {
	path := "/go/src/app/assets/img/text.svg"
	name := "assets/img/text.svg"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assetsImgTwitterSvg reads file data from disk. It returns an error on failure.
func assetsImgTwitterSvg() (*asset, error) {
	path := "/go/src/app/assets/img/twitter.svg"
	name := "assets/img/twitter.svg"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assetsImgYoutubeSvg reads file data from disk. It returns an error on failure.
func assetsImgYoutubeSvg() (*asset, error) {
	path := "/go/src/app/assets/img/youtube.svg"
	name := "assets/img/youtube.svg"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assetsIndexHtml reads file data from disk. It returns an error on failure.
func assetsIndexHtml() (*asset, error) {
	path := "/go/src/app/assets/index.html"
	name := "assets/index.html"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
}

// assetsJsProductionJs reads file data from disk. It returns an error on failure.
func assetsJsProductionJs() (*asset, error) {
	path := "/go/src/app/assets/js/production.js"
	name := "assets/js/production.js"
	bytes, err := bindataRead(path, name)
	if err != nil {
		return nil, err
	}

	fi, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset info %s at %s: %v", name, path, err)
	}

	a := &asset{bytes: bytes, info: fi}
	return a, err
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
	"assets/build/asset-manifest.json":                                   assetsBuildAssetManifestJson,
	"assets/build/css/2.4608ec59.chunk.css":                              assetsBuildCss24608ec59ChunkCss,
	"assets/build/css/2.4608ec59.chunk.css.map":                          assetsBuildCss24608ec59ChunkCssMap,
	"assets/build/css/2.8840a649.chunk.css":                              assetsBuildCss28840a649ChunkCss,
	"assets/build/css/2.8840a649.chunk.css.map":                          assetsBuildCss28840a649ChunkCssMap,
	"assets/build/css/main.5be27138.chunk.css":                           assetsBuildCssMain5be27138ChunkCss,
	"assets/build/css/main.5be27138.chunk.css.map":                       assetsBuildCssMain5be27138ChunkCssMap,
	"assets/build/css/main.c19d1c86.chunk.css":                           assetsBuildCssMainC19d1c86ChunkCss,
	"assets/build/css/main.c19d1c86.chunk.css.map":                       assetsBuildCssMainC19d1c86ChunkCssMap,
	"assets/build/favicon.ico":                                           assetsBuildFaviconIco,
	"assets/build/index.html":                                            assetsBuildIndexHtml,
	"assets/build/js/2.30a5f63c.chunk.js":                                assetsBuildJs230a5f63cChunkJs,
	"assets/build/js/2.30a5f63c.chunk.js.map":                            assetsBuildJs230a5f63cChunkJsMap,
	"assets/build/js/main.093370dd.chunk.js":                             assetsBuildJsMain093370ddChunkJs,
	"assets/build/js/main.093370dd.chunk.js.map":                         assetsBuildJsMain093370ddChunkJsMap,
	"assets/build/js/main.90e62041.chunk.js":                             assetsBuildJsMain90e62041ChunkJs,
	"assets/build/js/main.90e62041.chunk.js.map":                         assetsBuildJsMain90e62041ChunkJsMap,
	"assets/build/js/runtime-main.0d813b15.js":                           assetsBuildJsRuntimeMain0d813b15Js,
	"assets/build/js/runtime-main.0d813b15.js.map":                       assetsBuildJsRuntimeMain0d813b15JsMap,
	"assets/build/js/runtime-main.4a474bcd.js":                           assetsBuildJsRuntimeMain4a474bcdJs,
	"assets/build/js/runtime-main.4a474bcd.js.map":                       assetsBuildJsRuntimeMain4a474bcdJsMap,
	"assets/build/logo192.png":                                           assetsBuildLogo192Png,
	"assets/build/logo512.png":                                           assetsBuildLogo512Png,
	"assets/build/manifest.json":                                         assetsBuildManifestJson,
	"assets/build/media/fa-brands-400.099a9556.woff":                     assetsBuildMediaFaBrands400099a9556Woff,
	"assets/build/media/fa-brands-400.30cc681d.eot":                      assetsBuildMediaFaBrands40030cc681dEot,
	"assets/build/media/fa-brands-400.3b89dd10.ttf":                      assetsBuildMediaFaBrands4003b89dd10Ttf,
	"assets/build/media/fa-brands-400.ba7ed552.svg":                      assetsBuildMediaFaBrands400Ba7ed552Svg,
	"assets/build/media/fa-brands-400.f7307680.woff2":                    assetsBuildMediaFaBrands400F7307680Woff2,
	"assets/build/media/fa-regular-400.0bb42845.svg":                     assetsBuildMediaFaRegular4000bb42845Svg,
	"assets/build/media/fa-regular-400.1f77739c.ttf":                     assetsBuildMediaFaRegular4001f77739cTtf,
	"assets/build/media/fa-regular-400.7124eb50.woff":                    assetsBuildMediaFaRegular4007124eb50Woff,
	"assets/build/media/fa-regular-400.7630483d.eot":                     assetsBuildMediaFaRegular4007630483dEot,
	"assets/build/media/fa-regular-400.f0f82301.woff2":                   assetsBuildMediaFaRegular400F0f82301Woff2,
	"assets/build/media/fa-solid-900.1042e8ca.eot":                       assetsBuildMediaFaSolid9001042e8caEot,
	"assets/build/media/fa-solid-900.376c1f97.svg":                       assetsBuildMediaFaSolid900376c1f97Svg,
	"assets/build/media/fa-solid-900.605ed792.ttf":                       assetsBuildMediaFaSolid900605ed792Ttf,
	"assets/build/media/fa-solid-900.9fe5a17c.woff":                      assetsBuildMediaFaSolid9009fe5a17cWoff,
	"assets/build/media/fa-solid-900.e8a427e1.woff2":                     assetsBuildMediaFaSolid900E8a427e1Woff2,
	"assets/build/media/login.db9aa21e.png":                              assetsBuildMediaLoginDb9aa21ePng,
	"assets/build/precache-manifest.63f194368b5e14fc8d9700498338ad87.js": assetsBuildPrecacheManifest63f194368b5e14fc8d9700498338ad87Js,
	"assets/build/precache-manifest.ae6b8c83066240b96bd573fcc5f7e0da.js": assetsBuildPrecacheManifestAe6b8c83066240b96bd573fcc5f7e0daJs,
	"assets/build/precache-manifest.c5fc2df4798feaf42383f744c6942a0a.js": assetsBuildPrecacheManifestC5fc2df4798feaf42383f744c6942a0aJs,
	"assets/build/robots.txt":                                            assetsBuildRobotsTxt,
	"assets/build/service-worker.js":                                     assetsBuildServiceWorkerJs,
	"assets/css/main.css":                                                assetsCssMainCss,
	"assets/css/main.css.map":                                            assetsCssMainCssMap,
	"assets/favicon.ico":                                                 assetsFaviconIco,
	"assets/img/books.png":                                               assetsImgBooksPng,
	"assets/img/data.svg":                                                assetsImgDataSvg,
	"assets/img/facebook.svg":                                            assetsImgFacebookSvg,
	"assets/img/googleplay.png":                                          assetsImgGoogleplayPng,
	"assets/img/ios.png":                                                 assetsImgIosPng,
	"assets/img/lesson.png":                                              assetsImgLessonPng,
	"assets/img/mail-send.svg":                                           assetsImgMailSendSvg,
	"assets/img/mentioned.svg":                                           assetsImgMentionedSvg,
	"assets/img/notification.svg":                                        assetsImgNotificationSvg,
	"assets/img/progress.svg":                                            assetsImgProgressSvg,
	"assets/img/quiz.svg":                                                assetsImgQuizSvg,
	"assets/img/reading.svg":                                             assetsImgReadingSvg,
	"assets/img/security.svg":                                            assetsImgSecuritySvg,
	"assets/img/social-login.svg":                                        assetsImgSocialLoginSvg,
	"assets/img/storage.svg":                                             assetsImgStorageSvg,
	"assets/img/study-data.png":                                          assetsImgStudyDataPng,
	"assets/img/text.svg":                                                assetsImgTextSvg,
	"assets/img/twitter.svg":                                             assetsImgTwitterSvg,
	"assets/img/youtube.svg":                                             assetsImgYoutubeSvg,
	"assets/index.html":                                                  assetsIndexHtml,
	"assets/js/production.js":                                            assetsJsProductionJs,
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
	"assets": &bintree{nil, map[string]*bintree{
		"build": &bintree{nil, map[string]*bintree{
			"asset-manifest.json": &bintree{assetsBuildAssetManifestJson, map[string]*bintree{}},
			"css": &bintree{nil, map[string]*bintree{
				"2.4608ec59.chunk.css":        &bintree{assetsBuildCss24608ec59ChunkCss, map[string]*bintree{}},
				"2.4608ec59.chunk.css.map":    &bintree{assetsBuildCss24608ec59ChunkCssMap, map[string]*bintree{}},
				"2.8840a649.chunk.css":        &bintree{assetsBuildCss28840a649ChunkCss, map[string]*bintree{}},
				"2.8840a649.chunk.css.map":    &bintree{assetsBuildCss28840a649ChunkCssMap, map[string]*bintree{}},
				"main.5be27138.chunk.css":     &bintree{assetsBuildCssMain5be27138ChunkCss, map[string]*bintree{}},
				"main.5be27138.chunk.css.map": &bintree{assetsBuildCssMain5be27138ChunkCssMap, map[string]*bintree{}},
				"main.c19d1c86.chunk.css":     &bintree{assetsBuildCssMainC19d1c86ChunkCss, map[string]*bintree{}},
				"main.c19d1c86.chunk.css.map": &bintree{assetsBuildCssMainC19d1c86ChunkCssMap, map[string]*bintree{}},
			}},
			"favicon.ico": &bintree{assetsBuildFaviconIco, map[string]*bintree{}},
			"index.html":  &bintree{assetsBuildIndexHtml, map[string]*bintree{}},
			"js": &bintree{nil, map[string]*bintree{
				"2.30a5f63c.chunk.js":          &bintree{assetsBuildJs230a5f63cChunkJs, map[string]*bintree{}},
				"2.30a5f63c.chunk.js.map":      &bintree{assetsBuildJs230a5f63cChunkJsMap, map[string]*bintree{}},
				"main.093370dd.chunk.js":       &bintree{assetsBuildJsMain093370ddChunkJs, map[string]*bintree{}},
				"main.093370dd.chunk.js.map":   &bintree{assetsBuildJsMain093370ddChunkJsMap, map[string]*bintree{}},
				"main.90e62041.chunk.js":       &bintree{assetsBuildJsMain90e62041ChunkJs, map[string]*bintree{}},
				"main.90e62041.chunk.js.map":   &bintree{assetsBuildJsMain90e62041ChunkJsMap, map[string]*bintree{}},
				"runtime-main.0d813b15.js":     &bintree{assetsBuildJsRuntimeMain0d813b15Js, map[string]*bintree{}},
				"runtime-main.0d813b15.js.map": &bintree{assetsBuildJsRuntimeMain0d813b15JsMap, map[string]*bintree{}},
				"runtime-main.4a474bcd.js":     &bintree{assetsBuildJsRuntimeMain4a474bcdJs, map[string]*bintree{}},
				"runtime-main.4a474bcd.js.map": &bintree{assetsBuildJsRuntimeMain4a474bcdJsMap, map[string]*bintree{}},
			}},
			"logo192.png":   &bintree{assetsBuildLogo192Png, map[string]*bintree{}},
			"logo512.png":   &bintree{assetsBuildLogo512Png, map[string]*bintree{}},
			"manifest.json": &bintree{assetsBuildManifestJson, map[string]*bintree{}},
			"media": &bintree{nil, map[string]*bintree{
				"fa-brands-400.099a9556.woff":   &bintree{assetsBuildMediaFaBrands400099a9556Woff, map[string]*bintree{}},
				"fa-brands-400.30cc681d.eot":    &bintree{assetsBuildMediaFaBrands40030cc681dEot, map[string]*bintree{}},
				"fa-brands-400.3b89dd10.ttf":    &bintree{assetsBuildMediaFaBrands4003b89dd10Ttf, map[string]*bintree{}},
				"fa-brands-400.ba7ed552.svg":    &bintree{assetsBuildMediaFaBrands400Ba7ed552Svg, map[string]*bintree{}},
				"fa-brands-400.f7307680.woff2":  &bintree{assetsBuildMediaFaBrands400F7307680Woff2, map[string]*bintree{}},
				"fa-regular-400.0bb42845.svg":   &bintree{assetsBuildMediaFaRegular4000bb42845Svg, map[string]*bintree{}},
				"fa-regular-400.1f77739c.ttf":   &bintree{assetsBuildMediaFaRegular4001f77739cTtf, map[string]*bintree{}},
				"fa-regular-400.7124eb50.woff":  &bintree{assetsBuildMediaFaRegular4007124eb50Woff, map[string]*bintree{}},
				"fa-regular-400.7630483d.eot":   &bintree{assetsBuildMediaFaRegular4007630483dEot, map[string]*bintree{}},
				"fa-regular-400.f0f82301.woff2": &bintree{assetsBuildMediaFaRegular400F0f82301Woff2, map[string]*bintree{}},
				"fa-solid-900.1042e8ca.eot":     &bintree{assetsBuildMediaFaSolid9001042e8caEot, map[string]*bintree{}},
				"fa-solid-900.376c1f97.svg":     &bintree{assetsBuildMediaFaSolid900376c1f97Svg, map[string]*bintree{}},
				"fa-solid-900.605ed792.ttf":     &bintree{assetsBuildMediaFaSolid900605ed792Ttf, map[string]*bintree{}},
				"fa-solid-900.9fe5a17c.woff":    &bintree{assetsBuildMediaFaSolid9009fe5a17cWoff, map[string]*bintree{}},
				"fa-solid-900.e8a427e1.woff2":   &bintree{assetsBuildMediaFaSolid900E8a427e1Woff2, map[string]*bintree{}},
				"login.db9aa21e.png":            &bintree{assetsBuildMediaLoginDb9aa21ePng, map[string]*bintree{}},
			}},
			"precache-manifest.63f194368b5e14fc8d9700498338ad87.js": &bintree{assetsBuildPrecacheManifest63f194368b5e14fc8d9700498338ad87Js, map[string]*bintree{}},
			"precache-manifest.ae6b8c83066240b96bd573fcc5f7e0da.js": &bintree{assetsBuildPrecacheManifestAe6b8c83066240b96bd573fcc5f7e0daJs, map[string]*bintree{}},
			"precache-manifest.c5fc2df4798feaf42383f744c6942a0a.js": &bintree{assetsBuildPrecacheManifestC5fc2df4798feaf42383f744c6942a0aJs, map[string]*bintree{}},
			"robots.txt":                                            &bintree{assetsBuildRobotsTxt, map[string]*bintree{}},
			"service-worker.js":                                     &bintree{assetsBuildServiceWorkerJs, map[string]*bintree{}},
		}},
		"css": &bintree{nil, map[string]*bintree{
			"main.css":     &bintree{assetsCssMainCss, map[string]*bintree{}},
			"main.css.map": &bintree{assetsCssMainCssMap, map[string]*bintree{}},
		}},
		"favicon.ico": &bintree{assetsFaviconIco, map[string]*bintree{}},
		"img": &bintree{nil, map[string]*bintree{
			"books.png":        &bintree{assetsImgBooksPng, map[string]*bintree{}},
			"data.svg":         &bintree{assetsImgDataSvg, map[string]*bintree{}},
			"facebook.svg":     &bintree{assetsImgFacebookSvg, map[string]*bintree{}},
			"googleplay.png":   &bintree{assetsImgGoogleplayPng, map[string]*bintree{}},
			"ios.png":          &bintree{assetsImgIosPng, map[string]*bintree{}},
			"lesson.png":       &bintree{assetsImgLessonPng, map[string]*bintree{}},
			"mail-send.svg":    &bintree{assetsImgMailSendSvg, map[string]*bintree{}},
			"mentioned.svg":    &bintree{assetsImgMentionedSvg, map[string]*bintree{}},
			"notification.svg": &bintree{assetsImgNotificationSvg, map[string]*bintree{}},
			"progress.svg":     &bintree{assetsImgProgressSvg, map[string]*bintree{}},
			"quiz.svg":         &bintree{assetsImgQuizSvg, map[string]*bintree{}},
			"reading.svg":      &bintree{assetsImgReadingSvg, map[string]*bintree{}},
			"security.svg":     &bintree{assetsImgSecuritySvg, map[string]*bintree{}},
			"social-login.svg": &bintree{assetsImgSocialLoginSvg, map[string]*bintree{}},
			"storage.svg":      &bintree{assetsImgStorageSvg, map[string]*bintree{}},
			"study-data.png":   &bintree{assetsImgStudyDataPng, map[string]*bintree{}},
			"text.svg":         &bintree{assetsImgTextSvg, map[string]*bintree{}},
			"twitter.svg":      &bintree{assetsImgTwitterSvg, map[string]*bintree{}},
			"youtube.svg":      &bintree{assetsImgYoutubeSvg, map[string]*bintree{}},
		}},
		"index.html": &bintree{assetsIndexHtml, map[string]*bintree{}},
		"js": &bintree{nil, map[string]*bintree{
			"production.js": &bintree{assetsJsProductionJs, map[string]*bintree{}},
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
