package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/TharinduBalasooriya/newmovieapi/db"
	"github.com/TharinduBalasooriya/newmovieapi/models"
	"github.com/gorilla/mux"
)

func AddMovie(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	var movie models.Movie

	err := decoder.Decode(&movie)
	db.MovieList = append(db.MovieList, movie)

	if err != nil {
		log.Println(err)
	}

	w.WriteHeader(http.StatusOK)

}

func GetAllMovies(w http.ResponseWriter , r *http.Request){
	allMovies := db.MovieList
	json.NewEncoder(w).Encode(allMovies)
}

func GetMovieById(w http.ResponseWriter , r *http.Request){

	id := mux.Vars(r)["id"]

	for _,movie := range db.MovieList{

		if id == movie.Id{

			json.NewEncoder(w).Encode(movie)
			return
		}
	}

	message := models.Mesaage{Message: "Movie Not found"}
	json.NewEncoder(w).Encode(message)

}


func DeleteMOvieById(w http.ResponseWriter , r *http.Request){

	id := mux.Vars(r)["id"]

	for index,movie := range db.MovieList{

		if id == movie.Id{
			db.MovieList = append(db.MovieList[:index], db.MovieList[index+1:]...)
		}
	}
	json.NewEncoder(w).Encode(db.MovieList)
}
