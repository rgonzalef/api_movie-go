package model

import "database/sql"

type Comment struct {
	ID         int    `json:"id"`
	UserID     int 	`json:"user_id"`
	MovieID    int    `json:"movie_id"`
	Content    string `json:"comment_text"`
	CreateTime string `json:"timestamp"`
}

func AddCommentToDB(db *sql.DB, comment *Comment) error {
	query := "INSERT INTO comments (user_id, movie_id, comment_text, timestamp) VALUES (?, ?, ?, NOW())"
	result, err := db.Exec(query, comment.UserID, comment.MovieID, comment.Content)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	comment.ID = int(id)

	return nil

}

func GetCommentsFromDB(db *sql.DB) ([]Comment, error) {
	rows, err := db.Query(`SELECT id, user_id, movie_id, comment_text, timestamp FROM comments`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []Comment
	for rows.Next(){
		var comment Comment
		if err := rows.Scan(&comment.ID, &comment.UserID, &comment.MovieID, &comment.Content, &comment.CreateTime); err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}

	return comments, nil
}

func GetCommentsByIDFromDB(db *sql.DB, comment_id int) (*Comment, error) {
	var comment Comment
	query := "SELECT id, user_id, movie_id, comment_text, timestamp FROM comments WHERE id = ?"
	err := db.QueryRow(query, comment_id).Scan(&comment.ID, &comment.UserID, &comment.MovieID, &comment.Content, &comment.CreateTime)
	if err != nil {
		return nil, err
	}
	
	return &comment, nil
}

func DeleteCommentFromDB(db *sql.DB, comment_id int) error {
	query := "DELETE FROM comments WHERE id = ?"
	_, err := db.Exec(query, comment_id)
	if err != nil {
		return err
	}

	return nil

}

func EditCommentFromDB(db *sql.DB, comment_id int, new_comment string) error {
	query := "UPDATE comments SET comment_text = ? WHERE id = ?"
	_, err := db.Exec(query,new_comment, comment_id)
	if err != nil {
		return err
	}

	return nil

}
