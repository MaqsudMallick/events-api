package models

import (
	"errors"
	"module-name/db"
	"module-name/utils"
)

type User struct {
	ID       int64
	Username string `binding:"required"`
	Password string `binding:"required"`
}

func (u *User) SaveUser() error {
	query := `INSERT INTO users (username, password) VALUES ($1, $2) RETURNING id`
	hash, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}
	err = db.DB.QueryRow(query, u.Username, hash).Scan(&u.ID)
	if err != nil {
		return errors.New("User already exists")
	}
	return nil
}

func (u *User) IsUser() bool {
	query := `SELECT id, password FROM users WHERE username = $1`
	var password string
	err := db.DB.QueryRow(query, u.Username).Scan(&u.ID, &password)
	if err != nil {
		return false
	}
	return utils.CompareHashAndPassword(password, u.Password)
}
