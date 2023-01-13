package main

import (
	"fmt"
	"fortune3-pickaxe/controllers"
	"fortune3-pickaxe/templates"
	"fortune3-pickaxe/views"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func errorHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Page not found", http.StatusNotFound)
}

func main() {
	r := chi.NewRouter()

	r.Get("/", controllers.StaticHandler(views.Must(views.ParseFS(
		templates.FS, "home.gohtml", "tailwind.gohtml",
	))))

	goalsC := controllers.Goals{}
	goalsC.Templates.New = views.Must(views.ParseFS(
		templates.FS, "goal.gohtml", "tailwind.gohtml",
	))

	r.Get("/goal", goalsC.New)
	r.Post("/create-goal", goalsC.Create)
	r.NotFound(errorHandler)

	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe("127.0.0.1:3000", r)
}
