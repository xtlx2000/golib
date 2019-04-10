package main

import (
	"github.com/xtlx2000/golib/memory"

	"github.com/xtlx2000/golib/log"
)

func main() {
	p := memory.NewObjectPool(func() interface{} {
		return nil
	})
	p.Start()

	p.Put("1", 1)
	p.Put("2", 2)
	p.Put("3", 3)
	p.Put("4", 4)

	log.Infof("%v", p.Get("1"))
	log.Infof("%v", p.Get("2"))
	log.Infof("%v", p.Get("3"))
	log.Infof("%v", p.Get("4"))
	log.Infof("%v", p.Get("5"))
	log.Infof("%v", p.Get("6"))
}
