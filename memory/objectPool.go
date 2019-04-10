/* Usage:
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

type Object struct {
	o        interface{}
	lastTime time.Time
}

func NewObject(o interface{}) *Object {
	return &Object{
		o:        o,
		lastTime: time.Now(),
	}
}

type ObjectPool struct {
	pool map[string]*Object
	lock sync.RWMutex

	New func() interface{}
}

func NewObjectPool(New func() interface{}) *ObjectPool {
	return &ObjectPool{
		pool: make(map[string]*Object),
		New:  New,
	}
}

func (p *ObjectPool) Start() {
	// TODO
}

func (p *ObjectPool) Stop() {
	// TODO
}

func (p *ObjectPool) Put(key string, value interface{}) {
	p.lock.Lock()
	defer p.lock.Unlock()
	p.pool[key] = NewObject(value)
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
