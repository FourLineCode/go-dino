package game

import (
	_ "image/png"
)

var (
	EntityGround *Ground
	EntityDino   *Dino
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

	return nil
}
