package game

import (
	_ "image/png"
)

var (
	EntityGround *Ground
	EntityDino   *Dino
	EntityCactus *Cactus
)

func LoadEntities() error {
	var err error
	EntityGround = &Ground{}
	if err = EntityGround.Load(); err != nil {
		return err
	}

	EntityDino = &Dino{}
	if err = EntityDino.Load(); err != nil {
		return err
	}

	EntityCactus = &Cactus{}
	if err = EntityCactus.Load(); err != nil {
		return err
	}

	return nil
}
