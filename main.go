package main

import(
  "fmt"
  "log"
  "encoding/json"
  "math/random"
  "net/http"
  "srtconv"
  "github.com/gorilla/mux"
  )

type Movie struct{
  Id string 'json:"id"'
  Isbn string 'json:"isbn"'
  Title string 'json:"title"'
  Director *Director 'json:"director"'
}

type Director struct {
  Firstname string 'json:"firstname"'
  Lastname string 'json:"lastname"'
}

var movies []movie

func main(){
  r := mux.NewRouter()
  r.HandleFunc("/movie", getMovies).Methods("GET")
  r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
  r.HandleFunc("/movies", createMovie).Methods("POST")
  r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
  r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")
}
