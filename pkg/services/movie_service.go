package services

import (
	"database/sql"
	"fmt"
)

func CheckIfMovieInDB(db *sql.DB, movieID string) (bool, error) {
	var exists bool

	query := "SELECT EXISTS(SELECT 1 FROM movies WHERE movie_id = ?)"

	// Ejecuta la consulta para verificar si existe la película
	err := db.QueryRow(query, movieID).Scan(&exists)
	if err != nil {
		if err == sql.ErrNoRows {
			// No se encontraron filas, pero no es un error
			return false, nil
		}
		// Si ocurre otro error, lo devolvemos
		return false, err
	}

	return exists, nil
}

//agrega los campos id_movie y title a la tabla movies

func AddMovieInfo(db *sql.DB, movieID string, title string) error {
	query := "INSERT INTO movies (movie_id, title, visualizations) VALUES (?, ?, 1)"
	_, err := db.Exec(query, movieID, title )
	if err != nil {
		return fmt.Errorf("error al almacenar la información: %v", err)
		
	}
	return nil
}

// IncrementViewCount incrementa el contador de visualizaciones de una película en la base de datos.
func IncrementViewCount(db *sql.DB, movieID string) error {
	
	// Incrementar el contador de visualizaciones en la base de datos
	query := "UPDATE movies SET visualizations = visualizations + 1 WHERE movie_id = ?"
	_, err := db.Exec(query, movieID)
	if err != nil {
		return fmt.Errorf("error al incrementar el contador de visualizaciones: %v", err)
		
	}
	return nil
}


