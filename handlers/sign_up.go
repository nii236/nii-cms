package handlers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// SignUp is the handler for the registration page
func SignUp(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	t, err := loadTemplate("./layouts/admin/sign_up/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	err = t.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
