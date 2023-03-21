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
	ID       string   `Json:"id"`
	Isnb     string   `Json:"isdn"`
	Title    string   `Json:"Title"`
	Director Director `Json:"director"`
}

type Director struct {
	FName string `Json:"fName"`
	LName string `Json:"lName"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}

	}
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
		}
	}
	json.NewEncoder(w).Encode(movies)
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
	//set json content type
	w.Header().Set("Content-Type", "application/json")
	//get parameters
	params := mux.Vars(r)
	//loop over the movies
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			//adding a new movie
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

	movies = append(movies, Movie{ID: "1", Isnb: "4231", Title: "RRR", Director: Director{FName: "RAJA", LName: "Mouli"}})
	movies = append(movies, Movie{ID: "2", Isnb: "4232", Title: "BHB", Director: Director{FName: "RAJA", LName: "Mouli"}})
	movies = append(movies, Movie{ID: "3", Isnb: "4233", Title: "KGF", Director: Director{FName: "PRASHANTH", LName: "Neel"}})
	movies = append(movies, Movie{ID: "4 ", Isnb: "4234", Title: "Pushpa", Director: Director{FName: "Sukumar", LName: "K"}})
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Println("Sever is starting at port 5003")
	log.Fatal(http.ListenAndServe(":5003", r))
	fmt.Println(movies)
}
