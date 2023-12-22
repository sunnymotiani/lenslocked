package models

import "database/sql"

type Session struct {
	ID     int
	UserID int
	//TOKEN IS ONLY SET WHEN CREATING N NEW SESSION, When look up a session
	//THIS will be left empty as we only store the hash of session token
	//in our database and we connot reverse it into a raw token.
	Token     string
	TokenHash string
}

type SessionService struct {
	DB *sql.DB
}

func (ss *SessionService) Create(userID int) (*Session, error) {
	//TODO CREATE SESSION TOKEN
	//TODO : IMPLIMWNT SessionService.Create
	return nil, nil
}

func (ss *SessionService) User(token string) (*User, error) {
	//TODO : Impliment SessionService.User
	return nil, nil
}
