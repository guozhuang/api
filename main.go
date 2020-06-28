package main

import (
	"api/cmd"
)

func main() {
	app := cmd.Init()
	app.Run()
}
