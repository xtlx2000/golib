/* Usage:
func main() {
	p := memory.NewObjectPool()
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

	------  output  --------
	2019/04/10 06:05:07 [INFO] main.go:19:  1
	2019/04/10 06:05:07 [INFO] main.go:20:  2
	2019/04/10 06:05:07 [INFO] main.go:21:  3
	2019/04/10 06:05:07 [INFO] main.go:22:  4
	2019/04/10 06:05:07 [INFO] main.go:23:  <nil>
	2019/04/10 06:05:07 [INFO] main.go:24:  <nil>
}
*/
package memory

import (
	"sync"
	"time"
)

type object struct {
	o        interface{}
	lastTime time.Time
}

func newObject(o interface{}) *object {
	return &object{
		o:        o,
		lastTime: time.Now(),
	}
}

type ObjectPool struct {
	pool map[string]*object
	lock sync.RWMutex

	New func() interface{}
}

func NewObjectPool() *ObjectPool {
	return &ObjectPool{
		pool: make(map[string]*object),
	}
}

func (p *ObjectPool) Start() error {
	// TODO
	return nil
}

func (p *ObjectPool) Stop() error {
	// TODO
	return nil
}

func (p *ObjectPool) Put(key string, value interface{}) {
	p.lock.Lock()
	defer p.lock.Unlock()
	p.pool[key] = newObject(value)
}

func (p *ObjectPool) Get(key string) interface{} {
	p.lock.RLock()
	defer p.lock.RUnlock()
	value, isExist := p.pool[key]
	if !isExist {
		return p.New()
	}
	value.lastTime = time.Now()
	return value.o
}
