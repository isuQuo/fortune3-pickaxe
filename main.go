package main

import (
	"fmt"
	"fortune3-pickaxe/controllers"
	"fortune3-pickaxe/models"
	"fortune3-pickaxe/templates"
	"fortune3-pickaxe/views"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/gorilla/csrf"
)

func errorHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Page not found", http.StatusNotFound)
}

func main() {
	cfg := models.DefaultPostgresConfig()
	db, err := models.Open(cfg)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	r := chi.NewRouter()

	r.Get("/", controllers.StaticHandler(views.Must(views.ParseFS(
		templates.FS, "home.gohtml", "tailwind.gohtml",
	))))

	goalsC := controllers.Goals{}
	goalsC.Templates.New = views.Must(views.ParseFS(
		templates.FS, "goal.gohtml", "tailwind.gohtml",
	))

	usersC := controllers.Users{
		UserService: &models.UserService{
			DB: db,
		},
	}
	usersC.Templates.New = views.Must(views.ParseFS(
		templates.FS, "signup.gohtml", "tailwind.gohtml"))
	usersC.Templates.SignIn = views.Must(views.ParseFS(
		templates.FS, "signin.gohtml", "tailwind.gohtml"))

	r.Get("/signup", usersC.New)
	r.Post("/users", usersC.Create)
	r.Get("/signin", usersC.SignIn)
	r.Post("/signin", usersC.Authenticate)
	r.Get("/users/me", usersC.CurrentUser)

	r.Get("/goal", goalsC.New)
	r.Post("/create-goal", goalsC.Create)
	r.NotFound(errorHandler)

	fmt.Println("Starting the server on :3000...")
	CSRF := csrf.Protect([]byte("gFvi45R4fy5xNBlnEeZtQbfAVCYEIAUX"), csrf.Secure(false))
	http.ListenAndServe("127.0.0.1:3000", CSRF(r))
}
