package controllers

import "net/http"

// would we use generics here instead of an interface as a parameter?
func StaticHandler(tpl Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, r, nil)
	}
}
