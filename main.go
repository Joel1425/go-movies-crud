package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	ISBN     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}
type Director struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) { //to get all movies
	w.Header().Set("Content-Type", "application/json") //setting json content type
	json.NewEncoder(w).Encode(movies)                  // return movies
}

func deleteMovie(w http.ResponseWriter, r *http.Request) { // to delete a movie
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...) // in place of that movie all the rest movies are appended hence its deleted
			break
		}
	}
	json.NewEncoder(w).Encode(movies) // return what is left
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies { // use _ if youre not using anything
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item) // return in json format
			return
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(100000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") //setting json content type
	params := mux.Vars(r)                              // params
	// looping over movies
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}

}
func main() {
	r := mux.NewRouter()
	movies = append(movies, Movie{ID: "101", ISBN: "MOV101", Title: "Dil to Pagal Hai", Director: &Director{FirstName: "Virat", LastName: "Kohli"}}) //& is used for the reference
	movies = append(movies, Movie{ID: "102", ISBN: "MOV102", Title: "Malamaal Weekly", Director: &Director{FirstName: "Rohit", LastName: "Sharma"}})
	movies = append(movies, Movie{ID: "103", ISBN: "MOV103", Title: "Hera Pheri", Director: &Director{FirstName: "Bhuvi", LastName: "Kumar"}})
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting Server at Port: 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}
