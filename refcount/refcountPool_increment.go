package refcount

import (
	"errors"
	"sync"
)

type incrementRefCountPool struct {
	refMap  map[int]*Item
	currIdx int
	mutex   sync.RWMutex
}

func newIncrementRefCountPool() *incrementRefCountPool {
	return &incrementRefCountPool{
		refMap:  make(map[int]*Item),
		currIdx: 0,
	}
}

func (this *incrementRefCountPool) New() (*Item, error) {
	this.mutex.Lock()
	defer this.mutex.Unlock()

	this.currIdx++
	it := &Item{
		Id:       this.currIdx,
		Value:    nil,
		refCount: 1, // incrementRefCountPool hold 1
	}

	this.refMap[it.Id] = it
	it.refCount++ // New()'s caller hold 1

	return it, nil
}
func (this *incrementRefCountPool) Get(id int) (*Item, error) {
	this.mutex.Lock()
	defer this.mutex.Unlock()

	it, isExist := this.refMap[id]
	if !isExist {
		return nil, errors.New("item not found")
	}

	it.refCount++ // Get()'s caller hold 1
	return it, nil
}
func (this *incrementRefCountPool) Release(it *Item) {
	this.mutex.Lock()
	defer this.mutex.Unlock()

	if it.refCount <= 1 {
		return
	}
	it.refCount--
}
