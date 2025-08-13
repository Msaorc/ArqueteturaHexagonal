package server

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Msaorc/ArqueteturaHexagonal/adapters/web/handler"
	"github.com/Msaorc/ArqueteturaHexagonal/application"
	"github.com/codegangsta/negroni"
	"github.com/gorila/mux"
)

type WebServer struct {
	Service application.ProductServiceInterface
}

func MakeNewWebserver() *WebServer {
	return &WebServer{}
}

func (w WebServer) Server() {

	r := mux.NewRouter()
	n := negroni.New(
		negroni.NewLogger(),
	)

	handler.MakeProductHandlers(r, n, w.Service)
	http.Handle("/", r)

	server := &http.Server{
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
		Addr:              ":8080",
		Handler:           http.DefaultServeMux,
		ErrorLog:          log.New(os.Stderr, "log: ", log.Lshortfile),
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
