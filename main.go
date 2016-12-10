package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/nii236/nii-cms/handlers"
)

func main() {
	r := httprouter.New()
	r.GET("/", handlers.Landing)
	r.GET("/admin", handlers.Admin)
	r.GET("/admin/sign_up", handlers.SignUp)
	r.GET("/admin/sign_in", handlers.SignIn)
	r.POST("/admin/sign_up_callback", handlers.SignUpCallback)
	r.POST("/admin/sign_in_callback", handlers.SignInCallback)
	http.ListenAndServe(":8080", r)
}
