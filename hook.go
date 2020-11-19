// Copyright 2016 The go-vgo Project Developers. See the COPYRIGHT
// file at the top-level directory of this distribution and at
// https://github.com/go-vgo/robotgo/blob/master/LICENSE
//
// Licensed under the Apache License, Version 2.0 <LICENSE-APACHE or
// http://www.apache.org/licenses/LICENSE-2.0> or the MIT license
// <LICENSE-MIT or http://opensource.org/licenses/MIT>, at your
// option. This file may not be copied, modified, or distributed
// except according to those terms.

package robotgo_mouse

import (
	"strconv"

	hook "github.com/robotn/gohook"
)

/*
 ___________    ____  _______ .__   __. .___________.
|   ____\   \  /   / |   ____||  \ |  | |           |
|  |__   \   \/   /  |  |__   |   \|  | `---|  |----`
|   __|   \      /   |   __|  |  . `  |     |  |
|  |____   \    /    |  |____ |  |\   |     |  |
|_______|   \__/     |_______||__| \__|     |__|
*/

// EventStart start global event hook
// return event channel
func EventStart() chan hook.Event {
	return hook.Start()
}

// EventEnd removes global event hook
func EventEnd() {
	hook.End()
}

// Start start global event hook
// return event channel
func Start() chan hook.Event {
	return hook.Start()
}

// End removes global event hook
func End() {
	hook.End()
}

// StopEvent stop event listener
func StopEvent() {
	hook.StopEvent()
}

// EventProcess return go hook process
func EventProcess(events chan hook.Event) chan bool {
	return hook.Process(events)
}

// EventHook register gohook event
func EventHook(when uint8, keysPressed []string, callback func(hook.Event)) {
	hook.Register(when, keysPressed, callback)
}

// AddEvent add event listener,
//
// parameters for the string type,
// the keyboard corresponding key parameters,
//
// mouse arguments: mleft, center, mright, wheelDown, wheelUp,
// wheelLeft, wheelRight.
//
// Use "robotgo.AddEvents()" or "gohook" add asynchronous event listener
func AddEvent(key string) bool {
	var (
		// cs   *C.char
		mArr = []string{"mleft", "center", "mright", "wheelDown",
			"wheelUp", "wheelLeft", "wheelRight"}
		mouseBool bool
	)

	for i := 0; i < len(mArr); i++ {
		if key == mArr[i] {
			mouseBool = true
		}
	}

	if len(key) > 1 && !mouseBool {
		key = strconv.Itoa(int(Keycode[key]))
	}

	geve := hook.AddEvent(key)
	// defer C.free(unsafe.Pointer(cs))
	return geve == 0
}

// AddEvents add global event hook
//
// robotgo.AddEvents("q")
// robotgo.AddEvents("q", "ctrl")
// robotgo.AddEvents("q", "ctrl", "shift")
func AddEvents(key string, arr ...string) bool {
	s := hook.Start()
	// defer hook.End()

	ct := false
	k := 0
	for {
		e := <-s

		l := len(arr)
		if l > 0 {
			for i := 0; i < l; i++ {
				ukey := Keycode[arr[i]]

				if e.Kind == hook.KeyHold && e.Keycode == ukey {
					k++
				}

				if k == l {
					ct = true
				}

				if e.Kind == hook.KeyUp && e.Keycode == ukey {
					if k > 0 {
						k--
					}
					// time.Sleep(10 * time.Microsecond)
					ct = false
				}
			}
		} else {
			ct = true
		}

		if ct && e.Kind == hook.KeyUp && e.Keycode == Keycode[key] {
			hook.End()
			// k = 0
			break
		}
	}

	return true
}

// AddMouse add mouse event hook
//
// mouse arguments: left, center, right, wheelDown, wheelUp,
// wheelLeft, wheelRight.
//
// robotgo.AddMouse("left")
// robotgo.AddMouse("left", 100, 100)
func AddMouse(btn string, x ...int16) bool {
	s := hook.Start()
	ukey := MouseMap[btn]

	ct := false
	for {
		e := <-s

		if len(x) > 1 {
			if e.Kind == hook.MouseMove && e.X == x[0] && e.Y == x[1] {
				ct = true
			}
		} else {
			ct = true
		}

		if ct && e.Kind == hook.MouseDown && e.Button == ukey {
			hook.End()
			break
		}
	}

	return true
}

// AddMousePos add listen mouse event pos hook
func AddMousePos(x, y int16) bool {
	s := hook.Start()

	for {
		e := <-s
		if e.Kind == hook.MouseMove && e.X == x && e.Y == y {
			hook.End()
			break
		}
	}

	return true
}
