package model

import "database/sql"

type User struct {
	ID    int    	`json:"id"`
	Name  string 	`json:"name"`
	LastName string `json:"lastname"`
	Email string 	`json:"email"`
	Password string `json:"password"`
}

func GetUsers(db *sql.DB) ([]User, error) {
	rows, err := db.Query(`SELECT id, name, lastname, email FROM users`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next(){
		var user User
		if err := rows.Scan(&user.ID, &user.Name, &user.LastName, &user.Email); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func CreateUser(db *sql.DB, user *User) error {
	query := "INSERT INTO users (name, lastname, email, password) VALUES (?, ?, ?, ?)"
	result, err := db.Exec(query, user.Name, user.LastName, user.Email, user.Password)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	user.ID = int(id)

	return nil
}

func GetUserByID(db *sql.DB, id int) (*User, error) {
	var user User
	query := "SELECT id, name, lastname, email FROM users WHERE id = ?"
	err := db.QueryRow(query, id).Scan(&user.ID, &user.Name, &user.LastName, &user.Email)
	if err != nil {
		return nil, err
	}

	return &user, nil
}