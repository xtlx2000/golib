/* usage:
func main() {
	p := memory.NewObjectPool(func() interface{} {
		return nil
	})

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
)

type ObjectPool struct {
	pool map[string]interface{}
	lock sync.RWMutex

	New func() interface{}
}

func NewObjectPool(New func() interface{}) *ObjectPool {
	return &ObjectPool{
		pool: make(map[string]interface{}),
		New:  New,
	}
}

func (p *ObjectPool) Put(key string, value interface{}) {
	p.lock.Lock()
	p.pool[key] = value
	p.lock.Unlock()
}

func (p *ObjectPool) Get(key string) interface{} {
	p.lock.RLock()
	value, isExist := p.pool[key]
	p.lock.RUnlock()
	if !isExist {
		return p.New()
	}
	return value
}
