package httpServer

import (
	"errors"
	"fmt"
	"io"
	"learn-go-lang/basics"
	"net/http"
	"os"
)

// function of type `http.HandlerFunc` takes in 2 arguments - http.ResponseWriter and *http.Request
// this type of function is similar to express middleware

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("[REQUEST]: / \n")
	io.WriteString(w, "This is my website!\n")
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		{
			fmt.Printf("[REQUEST]: GET /hello \n")
			// ctx := r.Context()
			name := r.URL.Query().Get("name")
			io.WriteString(w, "Hello, "+name+"\n")
		}
	case "POST":
		{
			fmt.Printf("[REQUEST]: POST /hello \n")
			body := r.Body
			basics.Log("body", body)
			io.WriteString(w, "Hello, POST Hello!\n")
		}
	default:
		{
			fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
		}
	}
}

func BasicServer() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/hello", helloHandler)

	err := http.ListenAndServe(":3333", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}

func Server() {
	// mux is like express router
	mux := http.NewServeMux()
	mux.HandleFunc("/", rootHandler)
	mux.HandleFunc("/hello", helloHandler)

	err := http.ListenAndServe(":3000", mux)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
