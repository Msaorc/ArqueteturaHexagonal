package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Msaorc/ArqueteturaHexagonal/application"
)

func MakeProductHandlers(r *mux.Router, negroni *negroni.Negroni, service application.ProductServiceInterface) {
	r.Handler("/product/{id}", n.with(
		negroni.Wrap(getProduct(service)),
	)).Methods("GET", "OPTIONS")
}

func getProduct(service application.ProductInterface) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		vars := mux.Vars(r)
		id := vars["id"]
		product, err := service.GetId(id)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		err = json.NewEncoder(w).Encode(product)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	})
}
