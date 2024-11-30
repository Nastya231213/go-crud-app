package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"golang.org/x/exp/rand"
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

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}
func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.Id == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}
func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content Type", "application/json")

	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.Id = strconv.Itoa(rand.Intn(100000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movies)

}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.Id == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.Id = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
		}

	}

}
func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content Type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies {
		if item.Id == params["id"] {
			json.NewEncoder(w).Encode(item)
		}
	}
}

func main() {
	r := mux.NewRouter()
	movies = append(movies, Movie{Id: "1", Director: &Director{Firstname: "Maria", Lastname: "Dou"}, Isbn: "43323", Title: "Movie"})
	movies = append(movies, Movie{Id: "2", Director: &Director{Firstname: "Darya", Lastname: "Dou"}, Isbn: "43323", Title: "Movie"})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movie/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting server at post 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}
