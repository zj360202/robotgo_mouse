package main

import (
	"fmt"
	//"robotgo_mouse"
	"github.com/zj360202/robotgo_mouse"
)

func main() {
	fmt.Println("main go")
	////////////////////////////////////////////////////////////
	// 初始化新增接口，不能少，不然部分包含smooth的接口与MoveClick接口无法使用
	//w, h通过其他项目获取  比如：github.com/kbinani/screenshot
	robotgo_mouse.InitScreenSize(1920, 1080)
	////////////////////////////////////////////////////////////

	// 开始使用其他接口
	robotgo_mouse.MoveMouse(200, 200)
	x, y := robotgo_mouse.GetMousePos()
	fmt.Println(x, y)
}
