package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/joaoAraki27/go-project/assets"
)

type Player struct {
	image             *ebiten.Image
	position          Vector
	game              *Game
	laserLoadingTimer *Timer
}

func Newplayer(game *Game) *Player {
	image := assets.PlayerSprite

	bounds := image.Bounds()
	halfW := float64(bounds.Dx()) / 2

	position := Vector{
		x: screenWidth/2 - halfW,
		y: 500,
	}
	return &Player{
		image:             image,
		game:              game,
		position:          position,
		laserLoadingTimer: NewTimer(12),
	}
}

func (p *Player) Update() {
	speed := 10.0

	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		p.position.x -= speed
	} else if ebiten.IsKeyPressed(ebiten.KeyRight) {
		p.position.x += speed
	} else if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		p.position.y -= speed
	} else if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		p.position.y += speed
	}

	p.laserLoadingTimer.Update()
	if ebiten.IsKeyPressed(ebiten.KeySpace) && p.laserLoadingTimer.isReady() {

		p.laserLoadingTimer.Reset()

		bound := p.image.Bounds()
		halfW := float64(bound.Dx()) / 2 //Metade da largura da imagem do laser // Half the width of the laser image
		halfH := float64(bound.Dy()) / 2

		spawnPos := Vector{
			p.position.x + halfW,
			p.position.y - halfH/2,
		}

		laser := NewLaser(spawnPos)
		p.game.AddLasers(laser)
	}
}

func (p *Player) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	//The x and y position where the image will be drawn on the screen
	op.GeoM.Translate(p.position.x, p.position.y)
	//Draw an image on the screen
	screen.DrawImage(p.image, op)
}

func (p *Player) Collider() Rect {
	bounds := p.image.Bounds()

	return NewRect(
		p.position.x,
		p.position.y,
		float64(bounds.Dx()),
		float64(bounds.Dy()),
	)
}
