package main

import (
	"fmt"
	"html/template"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
)

type GetIndexHandler struct {
	once  sync.Once
	templ *template.Template
}

func (handler *GetIndexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	handler.once.Do(func() {
		handler.templ = template.Must(template.ParseFiles("templates/index.html"))
	})
	w.Header().Set("Content-Type", "text/html; charset=ISO-8859-1")
	w.WriteHeader(http.StatusOK)
	handler.templ.Execute(w, nil)
}

func GetNameHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	response := fmt.Sprintf("Hello %s!", vars["name"])
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, response)
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, "Oops! You're trying to access an invalid resource. Try learning to spell correctly.")
}

func main() {
	handle := mux.NewRouter()
	handle.NotFoundHandler = http.HandlerFunc(NotFoundHandler)
	handle.Handle("/", &GetIndexHandler{}).Methods("GET")
	handle.HandleFunc("/hello/{name}", GetNameHandler).Methods("GET")

	server := http.Server{
		Addr:    ":8080",
		Handler: handle,
	}

	server.ListenAndServe()
}
