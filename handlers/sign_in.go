package handlers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// SignIn is the handler for the Sign In page
func SignIn(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	t, err := loadTemplate("./layouts/admin/sign_in/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	err = t.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
