package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"src/utils"

	"github.com/gorilla/mux"
)

//Message structure
type Message struct {
	Country    string `json:"country"`
	ClientType string `json:"client-type"`
}

//Archivo de los seguros de caucion
type Archivo struct {
	Name string `json:"name"`
	File string `json:"file"`
}

//SeguroDeCaucion para sancor
type SeguroDeCaucion struct {
	BupID             int       `json:"bupId"`
	ProductorID       int       `json:"productorId"`
	Name              string    `json:"name"`
	TipoDoc           string    `json:"tipoDocumento"`
	NumeroDoc         int       `json:"numeroDocumento"`
	CodigoArea        int       `json:"codigoArea"`
	Phone             int       `json:"phone"`
	Email             string    `json:"email"`
	ProductoID        int       `json:"proheroku ductId"`
	CoberturaID       int       `json:"coberageId"`
	SumaAsegurada     int       `json:"sumaAsegurada"`
	Objeto            string    `json:"objeto"`
	TipoCertificacion string    `json:"tipoCertificacion"`
	Observaciones     string    `json:"observaciones"`
	Archivos          []Archivo `json:"archivos"`
}

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
	if r.Method == "POST" {
		w.WriteHeader(http.StatusOK)
		b, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		// Unmarshal
		var msg Message
		err = json.Unmarshal(b, &msg)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		clientType := msg.ClientType
		country := msg.Country

		bodyString := writeBodyForReg(country, clientType)

		w.Write([]byte(bodyString))
	} else {
		w.WriteHeader(http.StatusAccepted)
	}
}

func regs01Params(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	clientID := r.Header.Get("client-id")
	clientPass := r.Header.Get("client-pass")
	if clientID != "fede" || clientPass != "cloudgaia1" {
		w.WriteHeader(http.StatusUnauthorized)
	} else if r.Method == "GET" {
		params := mux.Vars(r)
		country := params["country"]
		clientType := params["clientType"]

		bodyString := writeBodyForReg(country, clientType)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(bodyString))
	}
}

func writeBodyForReg(country string, clientType string) string {
	var bodyArray []string

	bodyArray = append(bodyArray, `"description 1-> `+country+"-"+clientType+`"`)
	bodyArray = append(bodyArray, `"description 2"`)
	bodyArray = append(bodyArray, `"description 3"`)
	bodyArray = append(bodyArray, `"description 4"`)

	bodyString := `{"message" : [`
	bodyString += strings.Join(bodyArray, ",")
	bodyString += "]}"

	return bodyString
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
			body = body[:len(body)-1]
			body += "]"

			w.Write([]byte(body))
		}
	}
}

func regs02(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "POST" {
		w.Write([]byte(`{"message": "hola!"}`))
	}
}

func regs02Params(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "GET" {
		params := mux.Vars(r)
		docType := params["doc-type"]
		docNumber := params["doc-number"]
		body := `[{"doc-number": "123456789","doc-type": "dni", "name":"Fede", "lastname":"Longhi"}, {"doc-number": "987654321","doc-type": "passport", "name":"Fede", "lastname":"Longhi"},`
		body += `{"doc-number": ` + `"` + docNumber + `"` + `,"doc-type": ` + `"` + docType + `"` + `,"name": "Fede", "lastName":"Longhi"}]`

		w.Write([]byte(body))
	}
}

func orderServiceHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "GET" {
		// params := mux.Vars(r)
		// clientID := params["clientId"]
		role := r.Header.Get("role")
		var body []byte
		if role == "1" {
			body = utils.ReadFile("./resources/ordenes_1.JSON")
		}
		if role == "2" {
			body = utils.ReadFile("./resources/ordenes_2.JSON")
		}
		if role == "3" {
			body = utils.ReadFile("./resources/ordenes_3.JSON")
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(body))
	}
}

func segurosDeCaucionForClientHandler(w http.ResponseWriter, r *http.Request) {
	var body []byte
	if r.Method == "GET" {
		w.Header().Set("Content-Type", "application/json")
		body = utils.ReadFile("./resources/seguros-caucion.JSON")
		w.WriteHeader(http.StatusOK)
		w.Write(body)
	}
}

func segurosDeCaucionHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var body []byte
	if r.Method == "GET" {
		body = utils.ReadFile("./resources/seguros-caucion-all.JSON")
		w.Write(body)
	} else if r.Method == "POST" {
		b, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		var seguroCaucion SeguroDeCaucion
		err = json.Unmarshal(b, &seguroCaucion)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(b)
	}
}

func segurosDeCaucionFilteredHandler(w http.ResponseWriter, r *http.Request) {
	var body []byte
	if r.Method == "GET" {
		body = utils.ReadFile("./resources/seguros-caucion-filtered.JSON")
		w.WriteHeader(http.StatusOK)
		w.Write(body)
	}
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		//log.Fatal("$PORT must be set")
		port = "8080"
	}

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", homeLink)

	//**** Sancor endpoints ****//
	router.HandleFunc("/client/{clientId}/ordenesDeServicio", orderServiceHandler)
	router.HandleFunc("/client/{clientId}/segurosDeCaucion", segurosDeCaucionForClientHandler)
	router.HandleFunc("/segurosDeCaucion", segurosDeCaucionHandler)
	router.HandleFunc("/segurosDeCaucion/filtered", segurosDeCaucionFilteredHandler)
	//*********************************//

	//*********************************//
	router.HandleFunc("/regs-01/country/{country}/client-type/{clientType}", regs01Params)
	router.HandleFunc("/regs-01", regs01)
	router.HandleFunc("/regs-02/doc-type/{doc-type}/doc-number/{doc-number}", regs02Params)
	router.HandleFunc("/regs-02", regs02)
	router.HandleFunc("/loop/{number}", loopHandler)

	log.Fatal(http.ListenAndServe(":"+port, router))
}
