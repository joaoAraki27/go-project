package game

import (
	"math/rand" // Import the math/rand package

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/joaoAraki27/go-project/assets"
)

type Meteor struct {
	image    *ebiten.Image
	speed    float64
	position Vector
}

func NewMeteor() *Meteor {
	image := assets.MeteorSprites[rand.Intn(len(assets.MeteorSprites))]

	speed := 5 + rand.Float64()*13

	position := Vector{
		x: rand.Float64() * screenWidth,
		y: -100,
	}

	return &Meteor{
		image:    image,
		speed:    speed,
		position: position,
	}
}

func (m *Meteor) Update() {
	m.position.y += m.speed
}

func (m *Meteor) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	op.GeoM.Translate(m.position.x, m.position.y)

	screen.DrawImage(m.image, op)
}

func (m *Meteor) Collider() Rect {
	bounds := m.image.Bounds()

	return NewRect(
		m.position.x,
		m.position.y,
		float64(bounds.Dx()),
		float64(bounds.Dy()),
	)
}
