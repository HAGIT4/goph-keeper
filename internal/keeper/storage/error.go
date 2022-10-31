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

type ErrorLoginPassNotCreated struct {
	e error
}

func (err *ErrorLoginPassNotCreated) Error() string {
	return fmt.Sprintf("storage: loginpass not created: %s", err.e)
}

func NewErrorLoginPassNotCreated(e error) (err error) {
	return &ErrorLoginPassNotCreated{
		e: e,
	}
}

type ErrorLoginPassNotFoundByID struct {
	id     uuid.UUID
	userID uuid.UUID
	e      error
}

func (err *ErrorLoginPassNotFoundByID) Error() string {
	return fmt.Sprintf("storage: loginpass with ID %s and user ID %s not found: %s", err.id, err.userID, err.e)
}

func NewErrorLoginPassNotFoundByID(id uuid.UUID, userID uuid.UUID, e error) (err error) {
	return &ErrorLoginPassNotFoundByID{
		id:     id,
		userID: userID,
		e:      e,
	}
}

type ErrorLoginPassNotFoundByKeyword struct {
	keyword string
	userID  uuid.UUID
	e       error
}

func (err *ErrorLoginPassNotFoundByKeyword) Error() string {
	return fmt.Sprintf("storage: loginpass with keyword %s not found: %s", err.keyword, err.e)
}

func NewErrorLoginPassNotFoundByKeyword(keyword string, userID uuid.UUID, e error) (err error) {
	return &ErrorLoginPassNotFoundByKeyword{
		keyword: keyword,
		userID:  userID,
		e:       e,
	}
}

type ErrorLoginNotUpdatedByID struct {
	id     uuid.UUID
	userID uuid.UUID
	login  string
	e      error
}

func (err *ErrorLoginNotUpdatedByID) Error() string {
	return fmt.Sprintf("storage: loginpass with ID %s and user ID %s is not updated with login %s: %s",
		err.id, err.userID, err.login, err.e)
}

func NewErrorLoginNotUpdatedByID(id uuid.UUID, userID uuid.UUID, login string, e error) (err error) {
	return &ErrorLoginNotUpdatedByID{
		id:     id,
		userID: userID,
		login:  login,
		e:      e,
	}
}

type ErrorNoLoginPassFoundForUser struct {
	userID uuid.UUID
}

func (err *ErrorNoLoginPassFoundForUser) Error() string {
	return fmt.Sprintf("storage: no loginpass found for user with ID %s", err.userID)
}

func NewErrorNoLoginPassFoundForUser(userID uuid.UUID) (err error) {
	return &ErrorNoLoginPassFoundForUser{
		userID: userID,
	}
}
