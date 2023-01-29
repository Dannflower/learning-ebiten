package main

import (
	"log"
	"strconv"

	"github.com/dannflower/learning-ebiten/cmd/client/objects"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	tickCount int
	block     *objects.Block
	incX      bool
	incY      bool
}

func (g *Game) Update() error {

	width, height := g.Layout(ebiten.WindowSize())
	g.tickCount++

	x, y := g.block.GetPosition()

	if x == 0 {
		g.incX = true
	}

	if x == float64(width) {
		g.incX = false
	}

	if y == 0 {
		g.incY = true
	}

	if y == float64(height) {
		g.incY = false
	}

	if g.incX {
		x++
	} else {
		x--
	}

	if g.incY {
		y++
	} else {
		y--
	}

	g.block.SetPosition(x, y)

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	ebitenutil.DebugPrint(screen, strconv.FormatFloat(ebiten.ActualFPS(), 'f', 0, 64))

	g.block.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {

	game := &Game{}
	game.block = objects.NewBlock()

	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
