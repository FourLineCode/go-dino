package game

import (
	"bytes"
	_ "image/png"
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
	imageByte1, _ := files.ReadFile("assets/dino-run-0.png")
	reader1 := bytes.NewReader(imageByte1)
	d.SpriteRun0, _, err = ebitenutil.NewImageFromReader(reader1)
	if err != nil {
		return err
	}
	imageByte2, _ := files.ReadFile("assets/dino-run-1.png")
	reader2 := bytes.NewReader(imageByte2)
	d.SpriteRun1, _, err = ebitenutil.NewImageFromReader(reader2)
	if err != nil {
		return err
	}
	imageByte3, _ := files.ReadFile("assets/dino-lose.png")
	reader3 := bytes.NewReader(imageByte3)
	d.SpriteLose, _, err = ebitenutil.NewImageFromReader(reader3)
	if err != nil {
		return err
	}
	imageByte4, _ := files.ReadFile("assets/dino-stationary.png")
	reader4 := bytes.NewReader(imageByte4)
	d.SpriteStationary, _, err = ebitenutil.NewImageFromReader(reader4)
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

func (d *Dino) Update(g *Game) {
	if g.State == GameStateLost {
		d.SpriteCurrent = d.SpriteLose
		return
	}

	d.Counter++

	spacePressed := false
	for _, key := range g.keys {
		if key == ebiten.KeySpace {
			spacePressed = true
		}
	}
	if spacePressed || d.State == StateJumping {
		d.State = StateJumping
	} else {
		d.State = StateRunning
	}

	if d.Counter == SPEED_INCREASE_INTERVAL*2 {
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
