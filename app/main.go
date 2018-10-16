package main

import (
	"flag"
	"github.com/Nicolas-Constanty/Dnai.ML.Server/api"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

func main() {

	password := flag.String("password", "", "Defining server password")
	path := "routes.conf"

	flag.Parse()

	if (*password == "") {
		return
	}

	r := mux.NewRouter()
	api.GenerateRoute(r, path)
	api.Pass, _ = bcrypt.GenerateFromPassword([]byte(*password), 8)
	http.Handle("/", r)
	// Routes consist of a path and a handler function.
	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8000", r))
}


