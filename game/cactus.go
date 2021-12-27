package game

import (
	"bytes"
	_ "image/png"
	"math"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Cactus struct {
	X1       int
	X2       int
	Y        int
	MinDist  int
	Speed    int
	Counter  int
	Sprite   *ebiten.Image
	Options1 *ebiten.DrawImageOptions
	Options2 *ebiten.DrawImageOptions
}

func (c *Cactus) Load() error {
	var err error
	imageByte, _ := files.ReadFile("assets/cactus.png")
	reader := bytes.NewReader(imageByte)
	c.Sprite, _, err = ebitenutil.NewImageFromReader(reader)
	if err != nil {
		return err
	}

	c.MinDist = getRandomMinDist()
	c.X1 = CACTUS_START_X
	c.X2 = c.X1 + c.MinDist
	c.Y = WINDOW_HEIGHT / 2

	c.Speed = PLATFORM_MIN_SPEED
	c.Counter = 0

	return nil
}

func (c *Cactus) Draw(screen *ebiten.Image) {
	c.Options1 = &ebiten.DrawImageOptions{}
	c.Options1.GeoM.Translate(float64(c.X1), float64(c.Y))
	c.Options2 = &ebiten.DrawImageOptions{}
	c.Options2.GeoM.Translate(float64(c.X2), float64(c.Y))

	screen.DrawImage(c.Sprite, c.Options1)
	screen.DrawImage(c.Sprite, c.Options2)
}

func (c *Cactus) Update(g *Game) {
	if g.State == GameStateLost {
		return
	}

	c.Counter++
	if c.Counter == SPEED_INCREASE_INTERVAL {
		c.Counter = 0
		c.Speed = int(math.Min(float64(c.Speed+1), PLATFORM_MAX_SPEED))
	}

	c.X1 -= c.Speed
	c.X2 -= c.Speed

	dinoX := EntityDino.X + EntityDino.SpriteCurrent.Bounds().Max.X
	dinoY := EntityDino.Y + EntityDino.SpriteCurrent.Bounds().Max.Y
	cactusX1 := c.X1 + c.Sprite.Bounds().Max.X
	cactusX2 := c.X2 + c.Sprite.Bounds().Max.X
	if dinoX > c.X1+20 && EntityDino.X < cactusX1 {
		if dinoY >= c.Y+10 {
			g.State = GameStateLost
		}
	} else if dinoX > c.X2+20 && EntityDino.X < cactusX2 {
		if dinoY >= c.Y+10 {
			g.State = GameStateLost
		}
	}

	if c.X1 <= 0-c.Sprite.Bounds().Max.X {
		c.X1 = CACTUS_START_X
		g.Score++
	}
	if c.X2 <= 0-c.Sprite.Bounds().Max.X {
		c.X2 = c.X1 + c.MinDist
		c.MinDist = getRandomMinDist()
		g.Score++
	}
}

func getRandomMinDist() int {
	return rand.Intn(CACTUS_MAX_DIST-CACTUS_MIN_DIST) + CACTUS_MIN_DIST
}
