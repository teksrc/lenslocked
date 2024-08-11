package controllers

import (
	"fmt"
	"net/http"
)

type Users struct{
	Templates struct{
		New Template
	}
}

// could go alternative Templates UserTemplates but not going to.
// type UserTemplates struct{
// 	New views.Template
// }

func (u Users) New(w http.ResponseWriter, r *http.Request) {
	// Need a view to render. Later I can process some data and use it to render, like query params from email etc. 
	u.Templates.New.Execute(w, nil)
}

func (u Users) Create(w http.ResponseWriter, r *http.Request) {
	// https://pkg.go.dev/net/http#Request Form & PostForm -> ParseForm (idempotent)
	// err := r.ParseForm()
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// 	return
	// }
	// fmt.Fprint(w, "Email: ", r.PostForm.Get("email"))
	// fmt.Fprint(w, "Password: ", r.PostForm.Get("password"))

	// Swap to FormValue, https://pkg.go.dev/net/http#Request.FormValue

	fmt.Fprintf(w, "<p>Email: %s</p>", r.FormValue("email"))
	fmt.Fprintf(w, "<p>Password: %s</p>", r.FormValue("password"))
}
