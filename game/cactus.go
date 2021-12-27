package game

import (
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
	c.Sprite, _, err = ebitenutil.NewImageFromFile("assets/cactus.png")
	if err != nil {
		return err
	}

	c.MinDist = getRandomMinDist()
	c.X1 = 900
	c.X2 = c.X1 + c.MinDist
	c.Y = WINDOW_HEIGHT / 2

	c.Speed = 5
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
	c.Counter++
	if c.Counter == 1000 {
		c.Counter = 0
		c.Speed = int(math.Min(float64(c.Speed+1), 15))
	}

	c.X1 -= c.Speed
	c.X2 -= c.Speed

	if c.X1 <= 0-c.Sprite.Bounds().Max.X {
		c.X1 = 900
		g.Score++
	}
	if c.X2 <= 0-c.Sprite.Bounds().Max.X {
		c.X2 = c.X1 + c.MinDist
		c.MinDist = getRandomMinDist()
		g.Score++
	}
}

func getRandomMinDist() int {
	return rand.Intn(650-450) + 450
}
