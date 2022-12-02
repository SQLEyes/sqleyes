package main

import (
	"github.com/sqleyes/engine"
	_ "github.com/sqleyes/plugin-mysql"
	_ "github.com/sqleyes/plugin-postgresql"
	_ "plugin-pgproxy"
)

func main() {
	engine.Run("./config.yaml")
}
