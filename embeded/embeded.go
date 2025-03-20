package embeded

import (
	"embed"
	"io/fs"
)

//go:embed tiles/*
var tiles embed.FS

func Tiles() fs.FS {
	staticFS, err := fs.Sub(tiles, "tiles")
	if err != nil {
		panic(err)
	}
	return staticFS
}
