package handlers

import (
	"database/sql"

	"github.com/gorilla/mux"
)

func RouterHandlers(router *mux.Router, apiKey string, db *sql.DB) {

	//router.HandleFunc("/api/movie/{movie_id}", getMovieByID(apiKey)).Methods("GET")
	router.HandleFunc("/api/movie/{movie_id}", GetMovieDetails(db)).Methods("GET")
	router.HandleFunc("/api/movie/most-viewed/{n_movies}", GetMostViewedMovies(db)).Methods("GET")

	router.HandleFunc("/api/users", getUsers(db)).Methods("GET")
	router.HandleFunc("/api/users", createUser(db)).Methods("POST")
	router.HandleFunc("/api/user/{user_id}", getUserByID(db)).Methods("GET")
	
	router.HandleFunc("/api/comments", AddComment(db)).Methods("POST")
	router.HandleFunc("/api/comments", GetComments(db)).Methods("GET")
	router.HandleFunc("/api/comments/{comment_id}", DeleteComment(db)).Methods("DELETE")
	router.HandleFunc("/api/comments/{comment_id}", EditComment(db)).Methods("PUT")
	router.HandleFunc("/api/comments/{comment_id}", GetCommentByID(db)).Methods("GET")

}