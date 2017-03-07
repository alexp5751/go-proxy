package main

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.Handle("/", RootHandler{})
	logrus.Info("Listening on port :4400")
	log.Fatal(http.ListenAndServe(":4400", router))
}

type RootHandler struct {
}

func (h RootHandler) ServeHTTP(writer http.ResponseWriter, response *http.Request) {
	raw, err := ioutil.ReadFile("./data.json")
	if err != nil {
		logrus.Errorf("The file could not be read: %s", err.Error())
		writer.Write([]byte("a server error occurred :("))
		return
	}
	writer.Write(raw)
}
