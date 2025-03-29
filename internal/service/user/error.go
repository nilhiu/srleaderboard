package user

import "errors"

var (
	ErrRegisterUserExists     = errors.New("user already exists")
	ErrRegisterInsertFailed   = errors.New("failed to insert user into the database")
	ErrLoginIncorrectPassword = errors.New("incorrect password")
)
