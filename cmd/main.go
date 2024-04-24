package main

import (
	"fmt"
	"hello"
	"os"
)

func main() {
	if len(os.Args) > 1 {
	fmt.Println(hello.Say(os.Args[1]))
	} else {
		fmt.Println("Hello World!")
	}

	//declaration

	a := 2
	b := 2.1

	fmt.Printf("a : %T %v \n",a,a)
	fmt.Printf("b : %T %v \n",b,b)
	}


