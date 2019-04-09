package main

import (
	"github.com/xtlx2000/golib/memory"

	"github.com/cxr29/log"
)

func main() {
	p := memory.NewObjectPool(func() interface{} {
		return nil
	})

	p.Put("1", 1)
	p.Put("2", 2)
	p.Put("3", 3)
	p.Put("4", 4)

	log.Infof("%v\n", p.Get("1"))
	log.Infof("%v\n", p.Get("2"))
	log.Infof("%v\n", p.Get("3"))
	log.Infof("%v\n", p.Get("4"))
	log.Infof("%v\n", p.Get("5"))
	log.Infof("%v\n", p.Get("6"))
}
