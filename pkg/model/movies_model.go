package model

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"proyecto_final/internal/config"
)

// Estructura para los datos de la película
type Movie struct {
	ID 			int		`json:"id"`
	Title       string `json:"title"`
	Overview    string `json:"overview"`
	ReleaseDate string `json:"release_date"`
}

type Movie_DB struct {
	IDMovie			string `json:"movie_id"`
	Title			string `json:"title"`
	Visualizationzs	string `json:"visualizations"`
} 

func GetMovieByID(movieID string) (Movie, error){
	//lectura del archivo config.yaml - configuraciones
	cfg, err := config.LoadConfig()

	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	var movie Movie

	// Construye la URL de la solicitud
	url := fmt.Sprintf("%s/movie/%s?api_key=%s",cfg.TmdbUrl, movieID, cfg.ApiKey)

	// Haz la solicitud HTTP GET
	resp, err := http.Get(url)
	if err != nil {
		return movie, err
	}
	defer resp.Body.Close()

	// Verifica que la respuesta sea 200 OK
	if resp.StatusCode != http.StatusOK {
		return movie, fmt.Errorf("error: status code %d", resp.StatusCode)
	}

	// Decodifica la respuesta JSON en la estructura Movie
	if err := json.NewDecoder(resp.Body).Decode(&movie); err != nil {
		return movie, err
	}

	return movie, err

}

func GetMostViewedMoviesFromDB(db *sql.DB, n int) ([]Movie_DB, error) {
	rows, err := db.Query("SELECT movie_id, title, visualizations FROM movies ORDER BY visualizations DESC LIMIT ?", n)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var mostViewedMovies []Movie_DB
	for rows.Next() {
		var movie Movie_DB
		err := rows.Scan(&movie.IDMovie, &movie.Title, &movie.Visualizationzs)
		if err != nil {
			return nil, err
		}
		mostViewedMovies = append(mostViewedMovies, movie)
	}
	return mostViewedMovies, nil

}


