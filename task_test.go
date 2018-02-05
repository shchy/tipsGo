package tipsGo

import (
	"fmt"
	"testing"

	"github.com/shchy/tipsGo/models"
)

func TestA(t *testing.T) {
	fmt.Println("hello")
	u := models.User{}
	fmt.Println(u.Name)
}
