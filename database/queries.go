package database

func (db *Database) GetManyUsers() ([]User, error) {

	rows, err := db.Query("SELECT id, username, password, created_at FROM users")
	if err != nil {
		return nil, err
	}

	var users []User
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
