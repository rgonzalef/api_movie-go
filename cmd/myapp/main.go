package main

import (
	"fmt"
	"log"
	"net/http"
	"proyecto_final/internal/config"
	"proyecto_final/internal/database"
	"proyecto_final/internal/handlers"

	"github.com/gorilla/mux"
)


func main() {
	//lectura del archivo config.yaml - configuraciones
	cfg, err := config.LoadConfig()

	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	//conectarse a la DB
	db, err := database.Connect(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Error connecting to Database: %v", err)
	}
	defer db.Close()

	//creamos tablas users y commets
	if err := database.CreateTableUsers(db); err != nil {
		log.Fatalf("Error creating table: %v", err)
	}

	if err := database.CreateTableMovies(db); err != nil {
		log.Fatalf("Error creating table: %v", err)
	}

	if err := database.CreateTableComments(db); err != nil {
		log.Fatalf("Error creating table: %v", err)
	}

	// Crea un nuevo router Gorilla Mux
	router := mux.NewRouter()
	handlers.RouterHandlers(router, cfg.ApiKey, db )

	

	// Iniciamos el servidor en el puerto 8080
	fmt.Printf("Server running on http://localhost%s", cfg.ServerAddress )
	if err := http.ListenAndServe(cfg.ServerAddress, router); err != nil {
		log.Fatalf("Error iniciando el servidor: %v", err)
		//os.Exit(1)
	}
}
