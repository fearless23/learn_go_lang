package chi

import (
	"fmt"
	"learn-go-lang/basics"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}

func getHelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("[REQUEST]: GET /hello \n")
	name := r.URL.Query().Get("name")
	id := chi.URLParam(r, "id")
	w.Write([]byte("Hello, " + name + "--" + id + "\n"))
}

func postHelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("[REQUEST]: POST /hello \n")
	_body := r.Body
	basics.Log("body", _body)
	w.Write([]byte("Hello, POST Hello!\n"))
}

func Server() {
	router := chi.NewRouter()
	router.Use(middleware.Heartbeat("/healthcheck"))
	router.Use(middleware.Logger)
	router.Use(middleware.AllowContentType("application/json", "text/xml"))

	router.Get("/", rootHandler)
	router.Get("/hello", getHelloHandler)
	router.Get("/hello/{id}", getHelloHandler)
	router.Post("/hello", postHelloHandler)

	userRouter := UserRouter()
	router.Mount("/user", userRouter)

	http.ListenAndServe(":3000", router)
}
