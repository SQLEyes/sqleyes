package main

import (
	"engine"
	_ "plugin-postgresql"
)

func main() {
	engine.Run("./config.yaml")
}
