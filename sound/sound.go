package sound

import (
	"io"
	"os"

	"github.com/hajimehoshi/go-mp3"
	"github.com/hajimehoshi/oto"

	"github.com/xtlx2000/golib/log"
)

func PlayMp3(filepath string) error {
	f, err := os.Open(filepath)
	if err != nil {
		log.Errorf("open error: %v", err)
		return err
	}
	defer f.Close()

	d, err := mp3.NewDecoder(f)
	if err != nil {
		log.Errorf("decode error: %v", err)
		return err
	}
	ctx, err := oto.NewContext(d.SampleRate(), 2, 2, 8192)
	if err != nil {
		log.Errorf("oto error: %v", err)
		return err
	}
	defer ctx.Close()

	player := ctx.NewPlayer()
	defer player.Close()

	if _, err = io.Copy(player, d); err != nil {
		return err
	}
	return nil
}
