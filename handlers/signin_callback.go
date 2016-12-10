package handlers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// SignInCallback is the handler for the Sign In callback
func SignInCallback(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
