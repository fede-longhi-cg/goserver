package main
import (
	"net/http"
	"log"
	"fmt"
	"html"
)

func handlerFunction (w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func main(){

	http.HandleFunc("/", handlerFunction)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
