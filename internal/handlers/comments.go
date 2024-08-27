package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"proyecto_final/pkg/model"
	"strconv"

	"github.com/gorilla/mux"
)

func AddComment(db *sql.DB) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		var comment model.Comment
		if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err := model.AddCommentToDB(db, &comment); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(comment)
	}
}

func GetComments(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		comments, err := model.GetCommentsFromDB(db)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(comments)

	}
}

func GetCommentByID(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["comment_id"])
		if err != nil {
			http.Error(w, "Invalid comment ID", http.StatusBadRequest)
			return
		}

		comments, err := model.GetCommentsByIDFromDB(db, id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(comments)

	}
}

func DeleteComment(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Obtener el ID del comentario desde los parámetros de la URL
		vars := mux.Vars(r)
		commentID, err := strconv.Atoi(vars["comment_id"])
		if err != nil {
			http.Error(w, "Invalid comment ID", http.StatusBadRequest)
			return
		}

		// Obtener el ID del usuario desde los encabezados
		idUser, err := strconv.Atoi(r.Header.Get("id_user"))
		if err != nil {
			http.Error(w, "Invalid user ID", http.StatusBadRequest)
			return
		}

		// Obtener el comentario de la base de datos
		comments, err := model.GetCommentsByIDFromDB(db, commentID)
		if err != nil {
			http.Error(w, "Error fetching comment", http.StatusInternalServerError)
			return
		}

		// Verificar si el usuario es el propietario del comentario
		if comments.UserID == idUser {
			// Eliminar el comentario si el usuario es el propietario
			if err := model.DeleteCommentFromDB(db, commentID); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			// Responder con éxito
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode("Comment deleted successfully")
		} else {
			// Si el usuario no es el propietario, responder con un error de autorización
			http.Error(w, "You are not authorized to delete this comment", http.StatusUnauthorized)
		}
	}
}


func EditComment(db *sql.DB) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		// Obtener el ID del comentario desde los parámetros de la URL
		vars := mux.Vars(r)
		commentID, err := strconv.Atoi(vars["comment_id"]) 
		if err != nil {
			http.Error(w, "Invalid comment ID", http.StatusBadRequest)
			return
		}

		// Obtener el ID del usuario desde los encabezados
		idUser, err := strconv.Atoi(r.Header.Get("id_user"))
		if err != nil {
			http.Error(w, "Invalid user ID", http.StatusBadRequest)
			return
		}

		// Obtener el comentario de la base de datos
		comments, err := model.GetCommentsByIDFromDB(db, commentID)
		if err != nil {
			http.Error(w, "Error fetching comment", http.StatusInternalServerError)
			return
		}

		//Obtener el comentario nuevo del Body
		var comment model.Comment
			if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
		
		// Verificar si el usuario es el propietario del comentario
		if comments.UserID == idUser {
			// Eliminar el comentario si el usuario es el propietario
			if err := model.EditCommentFromDB(db, commentID, comment.Content ); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			// Responder con éxito
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode("Comment updated successfully")
		} else {
			// Si el usuario no es el propietario, responder con un error de autorización
			http.Error(w, "You are not authorized to edit this comment", http.StatusUnauthorized)
		}


	}
}