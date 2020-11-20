// Copyright 2016 The go-vgo Project Developers. See the COPYRIGHT
// file at the top-level directory of this distribution and at
// https://github.com/go-vgo/robotgo/blob/master/LICENSE
//
// Licensed under the Apache License, Version 2.0 <LICENSE-APACHE or
// http://www.apache.org/licenses/LICENSE-2.0> or the MIT license
// <LICENSE-MIT or http://opensource.org/licenses/MIT>, at your
// option. This file may not be copied, modified, or distributed
// except according to those terms.

// +build darwin windows

package robotgo_mouse

//import (
//	"fmt"
//	"testing"
//)
//
//func TestMoveMouse(t *testing.T) {
//	MoveMouse(20, 20)
//	MilliSleep(10)
//	x, y := GetMousePos()
//
//	fmt.Println(x, y)
//
//	//tt.Equal(t, 20, x)
//	//tt.Equal(t, 20, y)
//}

//
//func TestMoveMouseSmooth(t *testing.T) {
//	b := MoveMouseSmooth(100, 100)
//	MilliSleep(10)
//	x, y := GetMousePos()
//
//	tt.True(t, b)
//	tt.Equal(t, 100, x)
//	tt.Equal(t, 100, y)
//}
//
//func TestDragMouse(t *testing.T) {
//	DragMouse(500, 500)
//	MilliSleep(10)
//	x, y := GetMousePos()
//
//	tt.Equal(t, 500, x)
//	tt.Equal(t, 500, y)
//}
//
//func TestScrollMouse(t *testing.T) {
//	ScrollMouse(120, "up")
//	MilliSleep(100)
//
//	Scroll(210, 210)
//}
//
//func TestMoveRelative(t *testing.T) {
//	Move(200, 200)
//	MilliSleep(10)
//
//	MoveRelative(10, -10)
//	MilliSleep(10)
//
//	x, y := GetMousePos()
//	tt.Equal(t, 210, x)
//	tt.Equal(t, 190, y)
//}
//
//func TestMoveSmoothRelative(t *testing.T) {
//	Move(200, 200)
//	MilliSleep(10)
//
//	MoveSmoothRelative(10, -10)
//	MilliSleep(10)
//
//	x, y := GetMousePos()
//	tt.Equal(t, 210, x)
//	tt.Equal(t, 190, y)
//}
