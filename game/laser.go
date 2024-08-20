package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/joaoAraki27/go-project/assets"
)

type Laser struct {
	image    *ebiten.Image
	position Vector
}

func NewLaser(position Vector) *Laser {
	image := assets.LaserSprite

	bound := image.Bounds()
	halfW := float64(bound.Dx()) / 2 //Metade da largura da imagem do laser
	halfH := float64(bound.Dy()) / 2

	position.x -= halfW
	position.y -= halfH

	return &Laser{
		image:    image,
		position: position,
	}
}

func (l *Laser) Update() {
	speed := 10.0

	l.position.y += -speed
}

func (l *Laser) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	//Posicao X e Y que a imagem sera desenhada a tela
	op.GeoM.Translate(l.position.x, l.position.y)
	//Desenha imagem na tela
	screen.DrawImage(l.image, op)
}

func (l *Laser) Collider() Rect {
	bounds := l.image.Bounds()

	return NewRect(
		l.position.x,
		l.position.y,
		float64(bounds.Dx()),
		float64(bounds.Dy()),
	)
}
