package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello")

	for i := range []int{1, 2, 3} {
		u := Task{Name: fmt.Sprintf("Name %d", i), Age: i}
		fmt.Println(u)
	}
}

// Task 最小単位の作業
type Task struct {
	Name string
	Age  int
}
