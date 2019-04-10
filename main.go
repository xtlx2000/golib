package main

import (
	"strconv"
	"time"

	"github.com/xtlx2000/golib/log"
	"github.com/xtlx2000/golib/memory"
)

func main() {
	p := memory.NewTimeObjectPool(10)
	p.Start()

	for i := 0; i < 10; i++ {
		p.Put(strconv.Itoa(i), i)
	}
	time.Sleep(9 * time.Second)
	for i := 0; i < 11; i++ {
		value := p.Get(strconv.Itoa(i))
		log.Infof("get key=%v, value=%v", strconv.Itoa(i), value)
	}

	time.Sleep(20 * time.Second)
	for i := 0; i < 11; i++ {
		value := p.Get(strconv.Itoa(i))
		log.Infof("get key=%v, value=%v", strconv.Itoa(i), value)
	}
}
