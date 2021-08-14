package routes

import (
	"encoding/json"
	"net/http"

	"github.com/TharinduBalasooriya/newmovieapi/controller"
	"github.com/TharinduBalasooriya/newmovieapi/models"
	"github.com/gorilla/mux"
)


func MovieRoutes() *mux.Router{

	var router =  mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/",func(rw http.ResponseWriter, r *http.Request) {
		
		message := models.Mesaage{
			Message: "Welcome to the movie API!!!",
		}
		json.NewEncoder(rw).Encode(message)
	})


	router.HandleFunc("/api/movies",controller.GetAllMovies).Methods(http.MethodGet)
	router.HandleFunc("/api/movies/{id}",controller.GetMovieById).Methods(http.MethodGet)
	router.HandleFunc("/api/movies/{id}",controller.DeleteMOvieById).Methods(http.MethodDelete)
	router.HandleFunc("/api/movies",controller.AddMovie).Methods(http.MethodPost)

	return router
}