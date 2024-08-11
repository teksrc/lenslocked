package controllers

import (
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
