package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"proyecto_final/pkg/model"
	"proyecto_final/pkg/services"
	"strconv"

	"github.com/gorilla/mux"
)



func GetMovieDetails(db *sql.DB) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		movieID := vars["movie_id"]
	
		movieDetails, err := model.GetMovieByID(movieID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		flag, err := services.CheckIfMovieInDB(db, movieID)
		if err != nil {
		// Maneja el error apropiadamente
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
		}

		if flag {
			// La película existe en la base de datos
			services.IncrementViewCount(db, movieID)
			//fmt.Println("Se encontró el ID de la película en la base de datos.")
		} else {
			// La película no existe, agrega la información
			services.AddMovieInfo(db, movieID, movieDetails.Title)
			//fmt.Println("No se encontró el ID de la película en la base de datos, se agregó un nuevo registro.")
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(movieDetails)
	}
} 



func GetMostViewedMovies(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){
		params := mux.Vars(r)
		n, _ := strconv.Atoi(params["n_movies"])
		movies, err := model.GetMostViewedMoviesFromDB(db, n)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(movies)


	}
}







