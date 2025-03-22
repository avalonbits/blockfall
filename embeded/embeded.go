package embeded

import (
	"embed"
	"image"
	"io/fs"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

//go:embed tiles/*
var tiles embed.FS

func GetTiles() *Tiles {
	staticFS, err := fs.Sub(tiles, "tiles")
	if err != nil {
		panic(err)
	}
	return &Tiles{
		FS:    staticFS,
		cache: map[int]*ebiten.Image{},
	}
}

type Tiles struct {
	fs.FS
	cache map[int]*ebiten.Image
}

type Style string

const (
	Style1 = Style("style1")
	Style2 = Style("style2")
	Style3 = Style("style3")
)

func (t *Tiles) Tile(style Style, piece int) (*ebiten.Image, error) {
	if img, ok := t.cache[piece]; ok {
		return img, nil
	}

	img, _, err := ebitenutil.NewImageFromFileSystem(t.FS, string(style)+".png")
	if err != nil {
		return nil, err
	}
	tile := img.SubImage(image.Rectangle{
		Min: image.Point{
			X: 0,
			Y: int(piece-1) * 256,
		},
		Max: image.Point{
			X: 256,
			Y: int(piece) * 256,
		},
	})

	t.cache[piece] = tile.(*ebiten.Image)
	return t.cache[piece], nil
}
