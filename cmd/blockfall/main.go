package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	"log"

	"github.com/avalonbits/blockfall/embeded"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var (
	debug = flag.Bool("debug", false, "If true, enables debug info.")
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
	winX = 1280
	winY = 720
)

type tileInfo struct {
	color int
}

type Game struct {
	tiles    *embeded.Tiles
	tileSize int
	tileMap  [10][20]tileInfo
	tX       int
	tY       int
}

func (g *Game) Update() error {
	for i := range len(g.tileMap) {
		for j := range len(g.tileMap[i]) {
			g.tileMap[i][j].color = min(0, ((i+j)%4)+3)
		}
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	//	maxX, maxY := screen.Size()

	screen.Fill(color.RGBA{
		R: 0x20,
		G: 0x20,
		B: 0x20,
	})
	playArea := screen.SubImage(image.Rectangle{
		Min: image.Point{
			X: g.tX,
			Y: g.tY,
		},
		Max: image.Point{
			X: winX - g.tX,
			Y: winY - g.tY,
		},
	})
	playArea.(*ebiten.Image).Fill(color.Black)

	for i := range len(g.tileMap) {
		for j := range len(g.tileMap[i]) {
			color := g.tileMap[i][j].color
			if color == 0 {
				continue
			}

			tile, err := g.tiles.Tile(embeded.Style1, color)
			if err != nil {
				panic(err)
			}

			w, h := tile.Bounds().Dx(), tile.Bounds().Dy()
			scaleX, scaleY := float64(g.tileSize)/float64(w), float64(g.tileSize)/float64(h)
			opts := &ebiten.DrawImageOptions{}
			opts.GeoM.Scale(scaleX, scaleY)
			opts.GeoM.Translate(float64(g.tX+i*g.tileSize), float64(g.tY+j*g.tileSize))
			screen.DrawImage(tile, opts)
		}
	}

	if *debug {
		ebitenutil.DebugPrintAt(screen, fmt.Sprintf("TPS: %0.2f", ebiten.ActualTPS()), 16, 1200)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return winX, winY
}

func main() {
	flag.Parse()

	const tileScale = float64(min(winY, winX)) / 720.0
	const tileSize = int(32 * tileScale)
	g := &Game{
		tiles:    embeded.GetTiles(),
		tileSize: tileSize,
		tX:       (winX/2 - 10*tileSize/2),
		tY:       (winY/2 - 20*tileSize/2),
	}

	ebiten.SetWindowSize(winX, winY)
	ebiten.SetWindowTitle("Block Fall")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
