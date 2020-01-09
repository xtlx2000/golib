package main

import (
	"github.com/xtlx2000/golib/rotateLog"
)

func main() {

	log.SteupLog("name.log", 5*1024, 5)
	log.Errorf("wocao %v", nil)
	log.Errorf("wocao1 %v", "string")
	log.Errorf("wocao2 %v", 5)
	log.Errorf("wocao3 %v", 6.6)

	for i := 0; i < 10000; i++ {
		log.Errorf("wocao3 %v", 6.6)
	}
}
