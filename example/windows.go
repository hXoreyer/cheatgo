package main

import (
	"github.com/hxoreyer/cheatgo"
	"github.com/lxn/win"
)

func main() {
	window, err := cheatgo.NewWindow("windows", 1024, 768)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()
	window.Rectangle(cheatgo.RECT{
		Left:   100,
		Top:    100,
		Right:  256,
		Bottom: 321}, 1, cheatgo.NewRGB("#ffa07a"), win.PS_SOLID, cheatgo.NewRGB("#cfcfcf"))

	window.RunLoop(nil)
}
