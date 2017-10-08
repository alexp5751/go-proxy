package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/alexp5751/go-proxy/handler"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	router := mux.NewRouter()
	router.Handle("/", rootHandler{})
	router.Handle("/weather/{latitude},{longitude}", handler.Weather{})

	corsRouter := cors.Default().Handler(router)
	logrus.Infof("Listening on port :%s", port)
	log.Fatal(http.ListenAndServe(":"+port, corsRouter))
}

type rootHandler struct {
}

func (h rootHandler) ServeHTTP(writer http.ResponseWriter, response *http.Request) {
	help := "Candidate Education Widget (Hai)  -   /candidate-education\n"
	help += "Candidate Folders Widget (Ray)    -   /candidate-folders\n"
	help += "Bar Chart Widget (Bipol)          -   /bar-chart\n"
	help += "Header (Thomas)                   -   /header\n"
	help += "Search (Thomas)                   -   /search\n"
	help += "Events (Randy)                    -   /events\n"
	raw := []byte(help)
	writer.Write(raw)
}
