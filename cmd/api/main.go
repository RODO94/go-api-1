package main

import (
	"api-1/internal/handlers"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
)

func main(){

	log.SetReportCaller(true)
	var r *chi.Mux = chi.NewRouter()
	handlers.Handler(r)

	fmt.Println("Starting Go AP")

	err := http.ListenAndServe("localhost:8000", r)

	if err!= nil {
		log.Error(err)
	}
}