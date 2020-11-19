package main

import (
	"fmt"
	"robotgo_mouse/robotgo_mouse"
	//"github.com/zj360202/robotgo_mouse/robotgo_mouse"
)

func main() {
	fmt.Println("main go")
	robotgo_mouse.Test()

	robotgo_mouse.MoveMouse(200, 200)
	x, y := robotgo_mouse.GetMousePos()
	fmt.Println(x, y)
}
