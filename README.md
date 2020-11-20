# robotgo_mouse
robotgo_mouse只包含鼠标模拟操作的robotgo简化版(没有任何新功能开发，完全属于robotgo的子集合)

以下是我个人理解与测试，错误的地方请原谅与告知:
  1. keyboard感觉好像少几个按键( "-", "=", "[", "]", "\", ";", "'", ",", ".", "/")；
  2. 截屏功能时间开销又太大；(测试截屏到保存一张png开销在500ms)

所以就单独提取了mouse一块的功能

从完整引用robotgo项目生产exe 10MB到现在3.5MB

```Go
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
```

### 引用
[robotgo](https://github.com/go-vgo/robotgo) 原项目

## License

Robotgo is primarily distributed under the terms of both the MIT license and the Apache License (Version 2.0), with portions covered by various BSD-like licenses.

See [LICENSE-APACHE](http://www.apache.org/licenses/LICENSE-2.0), [LICENSE-MIT](https://github.com/go-vgo/robotgo/blob/master/LICENSE).
