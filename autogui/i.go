package autogui

import (
	"time"

	"github.com/go-vgo/robotgo"
)

// keyborad
func KeyPress(key string) {
	robotgo.KeyTap(key)
}

func Delay(ms int) {
	time.Sleep(time.Millisecond * time.Duration(ms))
}

func KeyDown(key string) {
	robotgo.KeyToggle(key, "down")
}

func KeyUp(key string) {
	robotgo.KeyToggle(key, "up")
}

func DoubleKeyClick(key1, key2 string) {
	KeyDown(key1)
	KeyDown(key2)
	Delay(500)
	KeyUp(key2)
	KeyUp(key1)
}

// mouse
func MoveTo(pos PositionInfo) {
	robotgo.MoveMouse(pos.X, pos.Y)
}

func LeftClick() {
	robotgo.MouseClick("left", true)
}

func RightClick() {
	Delay(500)
	robotgo.MouseClick("right", true)
}

func GetMousePos() PositionInfo {
	x, y := robotgo.GetMousePos()
	return NewPositionInfo(x, y)
}
