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

// creo la estructura del json
type Movie struct {
	Id       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

// director
type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

//defino pelis
var movies []Movie

//funciones del CRUD
func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	// no usamos index entonces cuando no
	// usamos un identificador l dejamos en blanco
	for _, item := range movies {
		if item.Id == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var movie Movie
	// blank identifier
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.Id = strconv.Itoa(rand.Intn(10000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	//set json content
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	//loop over the movies
	//delete the movie with id
	//add a new movie the movie that we send in the body of postman
	for index, item := range movies {
		if item.Id == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.Id = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.Id == params["id"] {
			// en go la forma de aliminar de un slice un elemento
			// es correr todos los elementos de la derecha
			// hacia la izquierda un lugar
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

// PPAL
func main() {

	r := mux.NewRouter()

	// precargo unas mubis
	movies = append(movies, Movie{
		Id:    "1",
		Isbn:  "3241234421",
		Title: "DepredaGO",
		Director: &Director{
			Firstname: "Roman",
			Lastname:  "GOlanski"}})
	movies = append(movies, Movie{
		Id:    "2",
		Isbn:  "2143652173",
		Title: "RanGO",
		Director: &Director{
			Firstname: "Martin",
			Lastname:  "sGOrcese"}})

	// defino las rutas
	r.HandleFunc("/movie", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	// aviso que voy a iniciar servicio y escucho&sirvo en el puerto
	fmt.Printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}
