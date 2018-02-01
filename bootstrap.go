package tipsGo

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/shchy/tipsGo/models"
)

func main() {
	fmt.Println("hello")
	fmt.Println(models.NewTask("namae", 1))

	r := mux.NewRouter()
	// r.HandleFunc("/", HomeHandler)
	// r.HandleFunc("/products", ProductsHandler)
	// r.HandleFunc("/articles", ArticlesHandler)
	http.Handle("/", r)
}
