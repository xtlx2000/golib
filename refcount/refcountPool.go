package refcount

type Item struct {
	Id       int
	Value    interface{}
	refCount int
}

type Pool interface {
	New() (*Item, error)
	Get(id int) (*Item, error)
	Release(*Item)
}

func NewRotateRefCountPool(maxCount int) Pool {
	return newRotateRefCountPool(maxCount)
}

func NewIncrementRefCountPool() Pool {
	return newIncrementRefCountPool()
}
