package main

import (
	"fmt"
	"github.com/maxxxlounge/pizza"
)

var (
	p pizza.Pizza
	s []pizza.Slice
)

func main (){
	p,err := pizza.New("../data/example.in")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(p)
}

