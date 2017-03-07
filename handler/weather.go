package handler

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type Weather struct {
}

func (w Weather) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	resp, err := http.Get(fmt.Sprintf("https://api.darksky.net/forecast/%s/%s,%s,%s", os.Getenv("DARK_SKY_API_KEY"), vars["latitude"], vars["longitude"], vars["time"]))
	if err != nil {
		writer.Write([]byte(err.Error()))
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		writer.Write([]byte(err.Error()))
	}
	writer.Write(body)
}
