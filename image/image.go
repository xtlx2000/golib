package image

import (
	"image"
	"os"

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
