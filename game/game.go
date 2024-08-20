package game

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/joaoAraki27/go-project/assets"
)

type Game struct {
	Player           *Player
	lasers           []*Laser
	meteors          []*Meteor
	MeteorSpawnTimer *Timer
	score            int
}

func NewGame() *Game {
	g := &Game{
		MeteorSpawnTimer: NewTimer(24),
	}
	player := Newplayer(g)
	g.Player = player
	return g
}

// Responsavel por atualizar a logica do jogo
// Roda em 60 FPS
// A lib garante 60 X por segundo - tem um forloop na lib que atualiza 60 fps
func (g *Game) Update() error {
	g.Player.Update()

	for _, l := range g.lasers {
		l.Update()
	}

	g.MeteorSpawnTimer.Update()
	if g.MeteorSpawnTimer.isReady() {
		g.MeteorSpawnTimer.Reset()
		m := NewMeteor()
		g.meteors = append(g.meteors, m)
	}

	for _, m := range g.meteors {
		m.Update()
	}

	for _, m := range g.meteors {
		if m.Collider().Intersects(g.Player.Collider()) {
			fmt.Println("voce perdeu")
			g.Reset()
		}
	}

	for i, m := range g.meteors {
		for j, l := range g.lasers {
			if m.Collider().Intersects(l.Collider()) {
				g.meteors = append(g.meteors[:i], g.meteors[i+1:]...)
				g.lasers = append(g.lasers[:j], g.lasers[j+1:]...)
				g.score += 1
			}
		}
	}

	return nil
}

// Responsavel por desenhar objetos na tela
// A lib garante 60 X por segundo - tem um forloop na lib que atualiza 60 fps
func (g *Game) Draw(screen *ebiten.Image) {
	g.Player.Draw(screen)

	for _, l := range g.lasers {
		l.Draw(screen)
	}

	for _, m := range g.meteors {
		m.Draw(screen)
	}

	text.Draw(screen, fmt.Sprintf("Pontos : %d", g.score), assets.FontUi, 20, 100, color.White)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func (g *Game) AddLasers(laser *Laser) {
	g.lasers = append(g.lasers, laser)
}

func (g *Game) Reset() {
	g.Player = Newplayer(g)
	g.meteors = nil
	g.lasers = nil
	g.MeteorSpawnTimer.Reset()
	g.score = 0
}
