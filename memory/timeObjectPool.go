/* Usage:
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
*/
package memory

import (
	"time"

	"github.com/emirpasic/gods/maps/linkedhashmap"
	"github.com/xtlx2000/golib/log"
)

type TimeObjectPool struct {
	hashMap *linkedhashmap.Map // map[string]*object
	// for getter goroutine
	getChan chan *objectGetter
	// for putter goroutine
	putChan chan *objectPutter
	// for patroltask goroutine
	patrolInterval int
	patrolNum      int
	validDuration  float64
	iter           linkedhashmap.Iterator
}

func NewTimeObjectPool(patrolInterval, patrolNum int, validDuration float64) *TimeObjectPool {
	hashMap := linkedhashmap.New()
	return &TimeObjectPool{
		hashMap:        hashMap,
		getChan:        make(chan *objectGetter),
		putChan:        make(chan *objectPutter),
		patrolInterval: patrolInterval,
		patrolNum:      patrolNum,
		validDuration:  validDuration,
		iter:           hashMap.Iterator(),
	}
}

func (p *TimeObjectPool) Start() error {
	go p.loop()
	return nil
}

func (p *TimeObjectPool) Stop() error {
	// TODO
	return nil
}

type objectPutter struct {
	key        string
	value      interface{}
	resultChan chan error
}

func (p *TimeObjectPool) Put(key string, value interface{}) error {
	putter := &objectPutter{
		key:        key,
		value:      value,
		resultChan: make(chan error),
	}
	p.putChan <- putter
	return <-putter.resultChan
}

type objectGetter struct {
	key        string
	resultChan chan interface{}
}

func (p *TimeObjectPool) Get(key string) interface{} {
	getter := &objectGetter{
		key:        key,
		resultChan: make(chan interface{}),
	}
	p.getChan <- getter
	return <-getter.resultChan
}

func (p *TimeObjectPool) loop() {
	for {
		select {
		// handle getter
		case getter := <-p.getChan:
			value, isExist := p.hashMap.Get(getter.key)
			var result interface{}
			if !isExist {
				result = nil
			} else {
				result = value.(*object).o
				value.(*object).lastTime = time.Now()
			}
			getter.resultChan <- result
		// handle putter
		case putter := <-p.putChan:
			p.hashMap.Put(putter.key, newObject(putter.value))
			putter.resultChan <- nil
		case <-time.After(time.Second * time.Duration(p.patrolInterval)):
			p.patrolTask()
		}
	}
}

func (p *TimeObjectPool) patrolTask() {
	log.Debugf("patrol task")
	now := time.Now()
	// iterate 100 items every task
	for i := 0; i < p.patrolNum; i++ {
		if !p.iter.Next() {
			p.iter.Begin()
			return
		}
		key, value := p.iter.Key(), p.iter.Value()

		if now.Sub(value.(*object).lastTime).Seconds() >= p.validDuration {
			log.Debugf("key=%v time exceed, removed.", key)
			p.hashMap.Remove(key)
		}

	}
}
