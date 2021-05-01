package autogui

import (
	"github.com/go-vgo/robotgo"
	"github.com/xtlx2000/WowAutoFisher/eve/common"
	"github.com/xtlx2000/golib/image"
	"github.com/xtlx2000/golib/log"
)

type PositionInfo struct {
	X int
	Y int
}

func NewPositionInfo(x, y int) PositionInfo {
	return PositionInfo{
		X: x,
		Y: y,
	}
}

func FindPosInScreen(bmpFilename string) (PositionInfo, error) {
	// get screen bmp
	screenBMP := robotgo.CaptureScreen()
	robotgo.SaveBitmap(screenBMP, "img/screen.png")
	defer robotgo.FreeBitmap(screenBMP)
	// get target bmp
	targetBMP := robotgo.OpenBitmap(bmpFilename)
	defer robotgo.FreeBitmap(targetBMP)

	width, height, err := image.GetImgSize(bmpFilename)
	if err != nil {
		log.Errorf("GetImgSize error: %v", err)
		return NewPositionInfo(-1, -1), err
	}
	// find
	fx, fy := robotgo.FindBitmap(targetBMP, screenBMP, 0.2)
	if fx == -1 || fy == -1 {
		return NewPositionInfo(-1, -1), common.POS_NOT_FOUND
	}

	fx = fx + (width / 2)
	fy = fy + (height / 2)

	return NewPositionInfo(fx, fy), nil
}
