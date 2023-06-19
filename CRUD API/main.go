package main

import (
	"fmt"
	"log"
	"net/http"
)

type Movie struct {
	ID       string    `json:id`
	Isbn     string    `json:isbn`
	Title    string    `json:title`
	Director *Director `json:director`
}

type Director struct {
	FirstName string `json:firstname`
	LastName  string `json:lastname`
}

var movies []Movie

func main() {
	r := mux.newRouter()

	movies = append(movies, Movie{ID: "1", Isbn: "438227", Title: "Movie One", Director: &Director{FirstName: "Lisa", LastName: "Foo"}})
	movies = append(movies, Movie{ID: "2", Isbn: "438229", Title: "Movie Two", Director: &Director{FirstName: "Julie", LastName: "Voo"}})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovies).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovies).Methods("DELETE")

	fmt.Println("Starting server at port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))

}
