package main

type Item struct {
	Id int
}

type item struct {
	Item
	num int
}

func newitem() *item {
	return &item{
		Item{
			Id: 1,
		},
		num: 2,
	}
}

func main() {
	var i *item
	i = newitem()
}
