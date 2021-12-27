package main

import (
	"embed"

	"github.com/FourLineCode/go-dino/game"
)

//go:embed assets/*
var f embed.FS

func main() {
	game.Run(f)
}
