package main

import (
	"fmt"
	"image"
	"image/color"
	"log"
	"math/rand/v2"

	"github.com/avalonbits/blockfall/embeded"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type shape [][]byte
type Tetromino struct {
	shape   []shape
	current int
}

func (t *Tetromino) Rotate() {
	t.current++
}

func (t Tetromino) Shape() [][]byte {
	which := t.current % len(t.shape)
	return [][]byte(t.shape[which])
}
func (t Tetromino) DefaultShape() [][]byte {
	return [][]byte(t.shape[0])
}

var pieces = []*Tetromino{
	{
		shape: []shape{
			{
				{1, 0, 0, 0},
				{1, 0, 0, 0},
				{1, 0, 0, 0},
				{1, 0, 0, 0},
			},
		},
	},
	{
		shape: []shape{
			{
				{0, 2, 0, 0},
				{0, 2, 0, 0},
				{2, 2, 0, 0},
				{0, 0, 0, 0},
			},
		},
	},
	{
		shape: []shape{
			{
				{3, 0, 0, 0},
				{3, 0, 0, 0},
				{3, 3, 0, 0},
				{0, 0, 0, 0},
			},
		},
	},
	{
		shape: []shape{
			{
				{0, 0, 0, 0},
				{4, 4, 0, 0},
				{4, 4, 0, 0},
				{0, 0, 0, 0},
			},
		},
	},
	{
		shape: []shape{
			{
				{0, 0, 0, 0},
				{0, 5, 5, 0},
				{5, 5, 0, 0},
				{0, 0, 0, 0},
			},
		},
	},
	{
		shape: []shape{
			{
				{6, 6, 6, 0},
				{0, 6, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
			},
		},
	},
	{
		shape: []shape{
			{
				{0, 0, 0, 0},
				{7, 7, 0, 0},
				{0, 7, 7, 0},
				{0, 0, 0, 0},
			},
		},
	},
}

const (
	NoPiece = iota
	Piece_I
	Piece_J
	Piece_L
	Piece_O
	Piece_S
	Piece_T
	Piece_Z
)

const (
	winY = 1280
	winX = 720
)

type Game struct {
	tiles   *embeded.Tiles
	tileMap [10][20]int
	style   int
	count   int
}

func (g *Game) Update() error {

	for i := range len(g.tileMap) {
		for j := range len(g.tileMap[i]) {
			if rand.IntN(2) == 0 {
				g.tileMap[i][j] = 7
			} else {
				g.tileMap[i][j] = 5
			}
		}
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{
		R: 0x20,
		G: 0x20,
		B: 0x20,
	})
	playArea := screen.SubImage(image.Rectangle{
		Min: image.Point{
			X: 8,
			Y: 8,
		},
		Max: image.Point{
			X: winX - 8,
			Y: winY - 8,
		},
	})
	playArea.(*ebiten.Image).Fill(color.Black)

	for i := range len(g.tileMap) {
		for j := range len(g.tileMap[i]) {
			tile, err := g.tiles.Tile(embeded.Style1, g.tileMap[i][j])
			if err != nil {
				panic(err)
			}
			w, h := tile.Bounds().Dx(), tile.Bounds().Dy()
			scale := 0.1875
			wf, hf := float64(w)*scale, float64(h)*scale
			opts := &ebiten.DrawImageOptions{}
			opts.GeoM.Scale(scale, scale)
			opts.GeoM.Translate(64+float64(i)*wf, 64+float64(j)*hf)
			screen.DrawImage(tile, opts)
		}
	}

	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("TPS: %0.2f", ebiten.ActualTPS()), 16, 1200)

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return winX, winY
}

func main() {
	g := &Game{
		tiles: embeded.GetTiles(),
		style: 1,
	}

	ebiten.SetWindowSize(winX, winY)
	ebiten.SetWindowTitle("Hello, World! Ebitengine")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
