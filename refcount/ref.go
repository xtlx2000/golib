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

func NewPool(maxCount int) Pool {
	return newRefPool(maxCount)
}
