package handlers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Admin is the handler for the admin welcome screen
func Admin(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	t, err := loadTemplate("./layouts/admin/admin.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	err = t.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
