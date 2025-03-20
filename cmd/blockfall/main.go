package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	winY = 1280
	winX = 720
)

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return winX, winY
}

func main() {
	ebiten.SetWindowSize(winX, winY)
	ebiten.SetWindowTitle("Hello, World! Ebitengine")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
