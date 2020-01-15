package main
import (
	"net/http"
	"log"
	"fmt"
	"html"
	"os"
)

type server struct{}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"message": "hello world"}`))
}

func handlerFunction (w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func main(){
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	s = &server{}
	http.Handle("/", s)

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
