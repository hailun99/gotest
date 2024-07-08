package main

import (
	"github.com/hailun99/gotest/util"
	"github.com/hailun99/gotest/tset"
	"fmt"

)

func init(){
	fmt.Println(3456)
}

func init2(){
	fmt.Println(6789)
}



func main(){
	fmt.Println("Hello World!")
	r := util.Add(1,2)
	fmt.Println(r)

	tset.Abc()
}


