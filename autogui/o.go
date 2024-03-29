package autogui

import (
	"github.com/go-vgo/robotgo"
	"github.com/xtlx2000/golib/common"
	//"github.com/xtlx2000/golib/image"
	//"github.com/xtlx2000/golib/log"
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

func FindPosInScreen(bmpFilename string, tol float64) (PositionInfo, error) {
	/*
		// get screen bmp
		screenBMP := robotgo.CaptureScreen()
		defer robotgo.FreeBitmap(screenBMP)
		// get target bmp
		targetBMP := robotgo.OpenBitmap(bmpFilename)
		defer robotgo.FreeBitmap(targetBMP)

		width, height, err := image.GetImgSize(bmpFilename)
		if err != nil {
			log.Errorf("GetImgSize error: %v", err)
			return NewPositionInfo(-1, -1), common.NOT_FOUND
		}
		// find
		fx, fy := robotgo.FindBitmap(targetBMP, screenBMP, tol)
		if fx == -1 || fy == -1 {
			return NewPositionInfo(-1, -1), common.NOT_FOUND
		}

		fx = fx + (width / 2)
		fy = fy + (height / 2)

		return NewPositionInfo(fx, fy), nil
	*/
	return NewPositionInfo(-1, -1), common.NOT_FOUND
}

func CaptureScreen(sWidth, sHeight, width, height int, imgFilename string) {
	/*
		screenBMP := robotgo.CaptureScreen(sWidth, sHeight, width, height)
		defer robotgo.FreeBitmap(screenBMP)
		robotgo.SaveBitmap(screenBMP, imgFilename)
	*/
	img := robotgo.CaptureImg(sWidth, sHeight, width, height)
	robotgo.Save(img, imgFilename)
}
