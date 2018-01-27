package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello")

	for i := range []int{1, 2, 3} {
		u := unko{Name: fmt.Sprintf("Name %d", i), Age: i}
		fmt.Println(u)
	}
}

type Task struct {
	Name string
	Age  int
}
