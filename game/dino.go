package game

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type DinoState int

const (
	StateRunning    DinoState = iota
	StateJumping    DinoState = iota
	StateStationary DinoState = iota
)

type Dino struct {
	X                int
	Y                int
	InitialHeight    int
	Counter          int
	Tick             int
	JumpSpeed        int
	JumpHeight       int
	JumpMaxHeight    int
	JumpDirection    int
	State            DinoState
	SpriteRun0       *ebiten.Image
	SpriteCurrent    *ebiten.Image
	SpriteRun1       *ebiten.Image
	SpriteLose       *ebiten.Image
	SpriteStationary *ebiten.Image
	Options          *ebiten.DrawImageOptions
}

func (d *Dino) Load() error {
	var err error
	d.SpriteRun0, _, err = ebitenutil.NewImageFromFile("assets/dino-run-0.png")
	if err != nil {
		return err
	}
	d.SpriteRun1, _, err = ebitenutil.NewImageFromFile("assets/dino-run-1.png")
	if err != nil {
		return err
	}
	d.SpriteLose, _, err = ebitenutil.NewImageFromFile("assets/dino-lose.png")
	if err != nil {
		return err
	}
	d.SpriteStationary, _, err = ebitenutil.NewImageFromFile("assets/dino-stationary.png")
	if err != nil {
		return err
	}

	d.SpriteCurrent = d.SpriteStationary
	d.State = StateStationary

	d.X = 50
	d.InitialHeight = WINDOW_HEIGHT/2 - 20
	d.Y = d.InitialHeight
	d.Counter = 0
	d.Tick = 10

	d.JumpMaxHeight = 15
	d.JumpSpeed = d.JumpMaxHeight
	d.JumpHeight = 0
	d.JumpDirection = -1

	return nil
}

func (d *Dino) Draw(screen *ebiten.Image) {
	d.Options = &ebiten.DrawImageOptions{}
	d.Options.GeoM.Translate(float64(d.X), float64(d.Y))

	screen.DrawImage(d.SpriteCurrent, d.Options)
}

func (d *Dino) Update(keys []ebiten.Key) {
	d.Counter++

	spacePressed := false
	for _, key := range keys {
		if key == ebiten.KeySpace {
			spacePressed = true
		}
	}
	if spacePressed || d.State == StateJumping {
		d.State = StateJumping
	} else {
		d.State = StateRunning
	}

	if d.Counter == 2000 {
		d.Counter = 0
		d.Tick = int(math.Max(float64(d.Tick-1), 5))
	}

	if d.Counter%d.Tick == 0 && d.State == StateRunning {
		if d.SpriteCurrent == d.SpriteRun0 {
			d.SpriteCurrent = d.SpriteRun1
		} else {
			d.SpriteCurrent = d.SpriteRun0
		}
	} else if d.State == StateJumping {
		d.SpriteCurrent = d.SpriteStationary
		d.Y += (d.JumpSpeed) * d.JumpDirection
		d.JumpHeight += (d.JumpSpeed) * d.JumpDirection

		if d.JumpDirection == -1 {
			d.JumpSpeed = int(math.Max(float64(d.JumpSpeed-1), 1))
		} else if d.JumpDirection == 1 {
			d.JumpSpeed = int(math.Min(float64(d.JumpSpeed+1), float64(d.JumpMaxHeight)))
		}

		if d.JumpHeight < -120 {
			d.JumpDirection = 1
		}
		if d.JumpHeight > 0 {
			d.State = StateRunning
			d.JumpSpeed = d.JumpMaxHeight
			d.JumpHeight = 0
			d.JumpDirection = -1
			d.Y = d.InitialHeight
		}
	}
}
