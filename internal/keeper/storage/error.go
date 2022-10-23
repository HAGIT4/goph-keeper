package storage

import (
	"fmt"

	"github.com/google/uuid"
)

type ErrorUsernameNotFound struct {
	Username string
}

func (err *ErrorUsernameNotFound) Error() string {
	return fmt.Sprintf("storage: user with username %s not found", err.Username)
}

type ErrorUserIDNotFound struct {
	ID uuid.UUID
}

func (err *ErrorUserIDNotFound) Error() string {
	return fmt.Sprintf("storage: user with ID %s not found", err.ID)
}

type ErrorUserNotCreated struct{}

func (err *ErrorUserNotCreated) Error() string {
	return "storage: user not created"
}
