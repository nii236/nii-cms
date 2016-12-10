package handlers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Landing is the handler for the admin welcome screen
func Landing(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	t, err := loadTemplate("./layouts/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	err = t.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
