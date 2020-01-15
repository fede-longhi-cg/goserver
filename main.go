package main
import (
	"net/http"
	"log"
	"fmt"
	"html"
	"os"
)

func handlerFunction (w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func main(){
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	http.HandleFunc("/", handlerFunction)

	log.Fatal(http.ListenAndServe(port, nil))
}
