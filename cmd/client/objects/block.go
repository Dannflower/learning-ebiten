package objects

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Block struct {
	x, y float64
	img  *ebiten.Image
}

func NewBlock() *Block {

	Block := &Block{}
	Block.img = ebiten.NewImage(20, 20)

	return Block
}

func (b *Block) GetPosition() (x, y float64) {
	return b.x, b.y
}

func (b *Block) SetPosition(x, y float64) {
	b.x = x
	b.y = y
}

func (b *Block) Draw(screen *ebiten.Image) {

	//b.img.Clear()
	b.img.Fill(color.White)
	ops := &ebiten.DrawImageOptions{}
	ops.GeoM.Translate(b.x, b.y)

	screen.DrawImage(b.img, ops)
}
