//go:generate rsrc -ico main.ico -manifest goversioninfo.exe.manifest -o main.syso
package main

import (
	"engine"
	_ "plugin-postgresql"
)

func main() {
	engine.Run()
}
