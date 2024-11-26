package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)
type Movie struct {
	Id       string    `json:"id"`
	Director *Director `json:"director"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

func main() {
	r := mux.NewRouter()
	movies=append(movies,Mvoie(ID: "1", Isbn:"43323",title:"Movie ",Director:&Director{Firstname: "John",Lastname: "Doe"}))

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("movies", createMovies).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movie/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting server at post 8000\n")
	log.Fata(http.ListenAndServe(":8000", r))
}
