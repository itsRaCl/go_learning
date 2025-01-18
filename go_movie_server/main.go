package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"first_name"`
	Lastname  string `json:"last_name"`
}

var movies []Movie

func getMovies(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(res).Encode(movies); err != nil {
		log.Fatal(err)
	}
}

func deleteMovie(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)

	for idx, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:idx], movies[idx+1:]...)
			break
		}
	}
}

func getMovie(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)

	for _, item := range movies {
		if item.ID == params["id"] {
			if err := json.NewEncoder(res).Encode(item); err != nil {
				log.Fatal(err)
			}
			break
		}
	}
}

func createMovie(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	var newMovie Movie

	_ = json.NewDecoder(req.Body).Decode(&newMovie)

	newMovie.ID = strconv.Itoa(rand.Intn(100000000))

	movies = append(movies, newMovie)

	fmt.Fprintf(res, "Movie Created Successfully!\n")
}

func updateMovie(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	var newMovie Movie
	params := mux.Vars(req)

	_ = json.NewDecoder(req.Body).Decode(&newMovie)

	for idx, item := range movies {
		if item.ID == params["id"] {
			newMovie.ID = item.ID
			movies = append(movies[:idx], movies[idx+1:]...)
			movies = append(movies, newMovie)
		}
	}

	_ = json.NewEncoder(res).Encode(newMovie)
}
func main() {
	r := mux.NewRouter()

	movies = append(movies, Movie{ID: "1", Isbn: "438227", Title: "The Perks of Being a Wallflower", Director: &Director{Firstname: "John", Lastname: "Doe"}})
	movies = append(movies, Movie{ID: "2", Isbn: "438228", Title: "Iron Man", Director: &Director{Firstname: "John", Lastname: "Faveau"}})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting server at port :8080\n")
	log.Fatal(http.ListenAndServe(":8080", r))
}
