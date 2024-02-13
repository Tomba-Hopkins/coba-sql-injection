package main

import (
	"coba-sqli/tambahan"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", tambahan.HandlerLogin)
	mux.HandleFunc("/register", tambahan.HandlerRegister)
	mux.HandleFunc("/dashboard", tambahan.HandlerDashboard)
	fileserver := http.FileServer(http.Dir("assets"))
	mux.Handle("/static/", http.StripPrefix("/static", fileserver))

	log.Println("Web started at port 8080")

	err := http.ListenAndServe("localhost:8080", mux)
	if err != nil {
		panic(err)
	}
}