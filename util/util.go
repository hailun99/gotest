package util

import (
	"fmt"
)

func aa() {
	fmt.Println("aa")
}

type A struct {
	Id   int64
	Name string
}

func (a A) print() {

}
