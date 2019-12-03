package main

import (
	"fmt"
)

type inter interface {
	Start()
}

type interStr struct {
}

func newInter() inter {
	return &interStr{}
}

func (i *interStr) Start() {

}

func main() {

	i := newInter()
	fmt.Printf("%v", i)
}
