package models

import (
	"github.com/DMohanty99/eventApp/db"
)

type User struct {
	Id       int64
	Email    string `binding:"required" json:"email"`
	Password string `binding:"required" json:"password"`
}

func (u *User) Save() error {
	query := "INSERT INTO users(email,password) VALUES (?,?) "
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close() // defer statement used to close the statement

	result, err := stmt.Exec(u.Email, u.Password)

	if err != nil {
		return err
	}

	userID, err := result.LastInsertId()

	u.Id = userID
	if err != nil {
		return err
	}

	return err
}
