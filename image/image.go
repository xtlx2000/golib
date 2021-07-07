package image

import (
	"image"
	"os"

	"image/png"

	"github.com/corona10/goimagehash"

	"github.com/xtlx2000/golib/log"
)

func GetImgSize(imgFilename string) (int, int, error) {
	file, err := os.Open(imgFilename)
	defer file.Close()
	if err != nil {
		log.Errorf("OpenImageFile error: %v", err)
		return -1, -1, err
	}
	cfg, _, err := image.DecodeConfig(file)
	if err != nil {
		log.Errorf("DecodeImageConfig error: %v", err)
		return -1, -1, err
	}
	return cfg.Width, cfg.Height, nil
}

func PNGDistance(png1, png2 string) (int, error) {
	file1, err1 := os.Open(png1)
	if err1 != nil {
		log.Errorf("Open error: %v", err1)
		return -1, err1
	}
	file2, err2 := os.Open(png2)
	if err2 != nil {
		log.Errorf("Open error: %v", err2)
		return -1, err2
	}
	defer file1.Close()
	defer file2.Close()

	img1, err1 := png.Decode(file1)
	img2, err2 := png.Decode(file2)

	hash1, _ := goimagehash.AverageHash(img1)
	hash2, _ := goimagehash.AverageHash(img2)
	distance, err := hash1.Distance(hash2)
	if err != nil {
		log.Errorf("distance error: %v", err)
		return -1, err
	}
	return distance, nil
}
