package database

import (
	"time"
)

func (db *Database) GetManyUsers() ([]User, error) {
	var users []User
	query := "SELECT id, username, password, created_at FROM users"

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var u User
		err := rows.Scan(&u.ID, &u.Username, &u.Password, &u.CreatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	err = rows.Close()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (db *Database) GetSingleUser(id int) (User, error) {
	var user User
	query := "SELECT id, username, password, created_at FROM users WHERE id = ?"

	err := db.QueryRow(query, id).Scan(&user.ID, &user.Username, &user.Password, &user.CreatedAt)
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func (db *Database) SignupUser(credentials Credentials) (string, error) {
	query := "INSERT INTO users (username, password, created_at) VALUES (?, ?, ?)"

	hash, err := HashPassword(credentials.Password)
	if err != nil {
		return "", err
	}

	_, err = db.Exec(query, credentials.Username, hash, time.Now())
	if err != nil {
		return "", err
	}

	tokenString, err := CreateJWTString(credentials.Username)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (db *Database) LoginUser(credentials Credentials) (string, error) {
	var user User
	query := "SELECT username, password from users WHERE username = ?"

	err := db.QueryRow(query, credentials.Username).Scan(&user.Username, &user.Password)
	if err != nil {
		return "", err
	}

	err = CheckPasswordHash(credentials.Password, user.Password)
	if err != nil {
		return "", err
	}

	tokenString, err := CreateJWTString(user.Username)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
