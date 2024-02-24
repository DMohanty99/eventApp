package models

import (
	"errors"

	"github.com/DMohanty99/eventApp/db"
	"github.com/DMohanty99/eventApp/utils"
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

	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}

	result, err := stmt.Exec(u.Email, hashedPassword)

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

func (u *User) ValidateCred() error {

	query := "SELECT id,password FROM users where email = ?"
	row := db.DB.QueryRow(query, u.Email)

	var retrievedPassword string
	err := row.Scan(&u.Id, &retrievedPassword)

	if err != nil {
		return errors.New("Invalid Credentials")
	}
	ok := utils.CheckPasswordHash(u.Password, retrievedPassword)

	if ok != true {
		return errors.New("Invalid Credentials")
	}
	return nil

}
