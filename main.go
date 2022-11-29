package main

import (
	"github.com/sqleyes/engine"
	_ "github.com/sqleyes/plugin-postgresql"
)

func main() {
	engine.Run("./config.yaml")
}
