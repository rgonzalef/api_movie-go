package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql" // Importa el driver de MySQL
)

func Connect(databaseUrl string) (*sql.DB, error) {
	db, err := sql.Open("mysql", databaseUrl)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, err
}

func CreateTableUsers(db *sql.DB) error {
	query := `
		CREATE TABLE IF NOT EXISTS users(
		    id INT AUTO_INCREMENT, 
		    name VARCHAR(100) NOT NULL,
			lastname VARCHAR(100) NOT NULL,
			email VARCHAR(100) NOT NULL UNIQUE,
			password VARCHAR(255) NOT NULL,
		    PRIMARY KEY(id)
		)
       `
	_, err := db.Exec(query)
	if err != nil {
		return err
	}

	log.Println("Table users created or already exists")
	return nil
	
	   
}

func CreateTableMovies(db *sql.DB) error {
	query := `
		CREATE TABLE IF NOT EXISTS movies (
            movie_id VARCHAR(100) NOT NULL,
			title VARCHAR(255),
            visualizations INT DEFAULT 0
             
)
       `
	_, err := db.Exec(query)
	if err != nil {
		return err
	}

	log.Println("Table movies created or already exists")
	return nil
	
	   
}


func CreateTableComments(db *sql.DB) error {
	query := `
		CREATE TABLE IF NOT EXISTS comments(
            id INT AUTO_INCREMENT PRIMARY KEY,
            user_id INT NOT NULL ,
            movie_id VARCHAR(100),
            comment_text TEXT,
            timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
            FOREIGN KEY (user_id) REFERENCES users(id)
            
)
       `
	_, err := db.Exec(query)
	if err != nil {
		return err
	}

	log.Println("Table commets created or already exists")
	return nil
	
	   
}

