package handlers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// SignUpCallback is the handler for the Register Callback
func SignUpCallback(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
