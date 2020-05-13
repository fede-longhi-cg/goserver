package main

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

func homeLink(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "GET":
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "get called"}`))
	case "POST":
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"message": "post called"}`))
	case "PUT":
		w.WriteHeader(http.StatusAccepted)
		w.Write([]byte(`{"message": "put called"}`))
	case "DELETE":
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "delete called"}`))
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "not found"}`))
	}
}

func regs01(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "GET" {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": ["description1","description2","description3"]}`))
	}
}

func regs01Params(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "GET" {
		params := mux.Vars(r)
		country := params["country"]
		clientType := params["clientType"]

		var bodyArray []string

		bodyArray = append(bodyArray, "description 1-> "+country+"-"+clientType)
		bodyArray = append(bodyArray, "description 2")
		bodyArray = append(bodyArray, "description 3")
		bodyArray = append(bodyArray, "description 4")

		bodyString := "["
		bodyString += strings.Join(bodyArray, ",")
		bodyString += "]"
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(bodyString))
	}
}

func loopHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "GET" {
		vars := mux.Vars(r)
		number, err := strconv.Atoi(vars["number"])
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`should be a number`))
		} else {
			w.WriteHeader(http.StatusOK)
			body := "["
			for i := 0; i < number; i++ {
				body += strconv.Itoa(i) + ","
			}

			w.Write([]byte(body))
		}
	}
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", homeLink)
	router.HandleFunc("/regs-01/country/{country}/client-type/{clientType}", regs01Params)
	router.HandleFunc("/regs-01/", regs01)
	router.HandleFunc("/loop/{number}", loopHandler)

	log.Fatal(http.ListenAndServe(":"+port, router))
}
