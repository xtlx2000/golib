package refcount

import (
	"errors"
	"sync"
)

type pool struct {
	refMap   []*Item
	currIdx  int
	maxCount int
	mutex    sync.RWMutex
}

func newRefPool(maxCount int) *pool {
	return &pool{
		refMap:   make([]*Item, maxCount, maxCount),
		currIdx:  0,
		maxCount: maxCount,
	}
}

func (this *pool) New() (*Item, error) {
	this.mutex.Lock()
	defer this.mutex.Unlock()

	it := &Item{
		Id:       0,
		Value:    nil,
		refCount: 1, // pool hold 1
	}

	this.currIdx = (this.currIdx + 1) % this.maxCount
	for {
		if this.refMap[this.currIdx] != nil && this.refMap[this.currIdx].refCount > 1 {
			this.currIdx = (this.currIdx + 1) % this.maxCount
		} else {
			break
		}
	}
	it.Id = this.currIdx
	this.refMap[it.Id] = it
	it.refCount++ // New()'s caller hold 1

	return it, nil
}
func (this *pool) Get(id int) (*Item, error) {
	this.mutex.Lock()
	defer this.mutex.Unlock()

	if id >= this.maxCount {
		return nil, errors.New("exceed maxCount")
	}
	it := this.refMap[id]
	if it == nil || it.Id != id {
		return nil, errors.New("no item")
	}
	it.refCount++ // Get()'s caller hold 1
	return it, nil
}
func (this *pool) Release(it *Item) {
	this.mutex.Lock()
	defer this.mutex.Unlock()

	if it.refCount <= 1 {
		return
	}
	it.refCount--
}
