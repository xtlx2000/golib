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
