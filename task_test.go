package tips

import (
	"fmt"
	"testing"
)

func TestA(t *testing.T) {
	fmt.Println("hello")
	fmt.Println(TaskImpl{name: "namae", value: 1})
	u := User{}
	fmt.Println(u.GetName())
}
