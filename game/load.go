package game

import (
	_ "image/png"
)

var (
	EntityGround *Ground
)

func LoadEntities() error {
	var err error
	EntityGround = &Ground{}
	if err = EntityGround.Load(); err != nil {
		return err
	}

	return nil
}
