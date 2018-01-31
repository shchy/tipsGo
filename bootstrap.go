package tips

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("hello")
	fmt.Println(TaskImpl{name: "namae", value: 1})

	//

	r := mux.NewRouter()
	// r.HandleFunc("/", HomeHandler)
	// r.HandleFunc("/products", ProductsHandler)
	// r.HandleFunc("/articles", ArticlesHandler)
	http.Handle("/", r)
}
