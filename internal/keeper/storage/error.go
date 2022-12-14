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

type ErrorTextDataNotCreated struct {
	e error
}

func (err *ErrorTextDataNotCreated) Error() string {
	return fmt.Sprintf("storage: textdate not created: %s", err.e)
}

func NewErrorTextDataNotCreated(e error) (err error) {
	return &ErrorLoginPassNotCreated{
		e: e,
	}
}

type ErrorTextDataNotFoundByID struct {
	id     uuid.UUID
	userID uuid.UUID
	e      error
}

func (err *ErrorTextDataNotFoundByID) Error() string {
	return fmt.Sprintf("storage: textdata with ID %s and user ID %s not found: %s", err.id, err.userID, err.e)
}

func NewErrorTextDataNotFoundByID(id uuid.UUID, userID uuid.UUID, e error) (err error) {
	return &ErrorLoginPassNotFoundByID{
		id:     id,
		userID: userID,
		e:      e,
	}
}

type ErrorTextDataNotFoundByKeyword struct {
	keyword string
	userID  uuid.UUID
	e       error
}

func (err *ErrorTextDataNotFoundByKeyword) Error() string {
	return fmt.Sprintf(`storage: textdata with keyword %s and user ID %s not found: %s`,
		err.keyword, err.userID, err.e)
}

func NewErrorTextDataNotFoundByKeyword(keyword string, userID uuid.UUID, e error) (err error) {
	return &ErrorTextDataNotFoundByKeyword{
		keyword: keyword,
		userID:  userID,
		e:       e,
	}
}

type ErrorTextDataNotUpdated struct {
	id     uuid.UUID
	userID uuid.UUID
	e      error
}

func (err *ErrorTextDataNotUpdated) Error() string {
	return fmt.Sprintf("storage: textdata with ID %s and user ID %s is not updated: %s",
		err.id, err.userID, err.e)
}

func NewErrorTextDataNotUpdated(id uuid.UUID, userID uuid.UUID, e error) (err error) {
	return &ErrorLoginNotUpdatedByID{
		id:     id,
		userID: userID,
		e:      e,
	}
}

type ErrorTextDataNotDeleted struct {
	id     uuid.UUID
	userID uuid.UUID
	e      error
}

func (err *ErrorTextDataNotDeleted) Error() string {
	return fmt.Sprintf("storage: textdata with ID %s and user ID %s is not deleted: %s",
		err.id, err.userID, err.e)
}

func NewErrorTextDataNotDeleted(id uuid.UUID, userID uuid.UUID, e error) (err error) {
	return &ErrorTextDataNotDeleted{
		id:     id,
		userID: userID,
		e:      e,
	}
}

type ErrorTextDataNotFoundForUser struct {
	userID uuid.UUID
	e      error
}

func (err *ErrorTextDataNotFoundForUser) Error() string {
	return fmt.Sprintf("storage: textdata for user with ID %s not found: %s",
		err.userID, err.e,
	)
}

func NewErrorTextDataNotFoundForUser(userID uuid.UUID, e error) (err error) {
	return &ErrorTextDataNotFoundForUser{
		userID: userID,
		e:      e,
	}
}

type ErrorCardDataNotCreated struct {
	e error
}

func (err *ErrorCardDataNotCreated) Error() string {
	return fmt.Sprintf("storage: carddata not created: %s", err.e)
}

func NewErrorCardDataNotCreated(e error) (err error) {
	return &ErrorCardDataNotCreated{
		e: e,
	}
}

type ErrorCardDataNotFoundByKeyword struct {
	keyword string
	userID  uuid.UUID
	e       error
}

func (err *ErrorCardDataNotFoundByKeyword) Error() string {
	return fmt.Sprintf("storage: carddata with keyword %s not found: %s", err.keyword, err.e)
}

func NewErrorCardDataNotFoundByKeyword(keyword string, userID uuid.UUID, e error) (err error) {
	return &ErrorCardDataNotFoundByKeyword{
		keyword: keyword,
		userID:  userID,
		e:       e,
	}
}

type ErrorNoCardDataFoundForUser struct {
	userID uuid.UUID
}

func (err *ErrorNoCardDataFoundForUser) Error() string {
	return fmt.Sprintf("storage: no carddata found for user with ID %s", err.userID)
}

func NewErrorNoCardDataFoundForUser(userID uuid.UUID) (err error) {
	return &ErrorNoCardDataFoundForUser{
		userID: userID,
	}
}

type ErrorBinaryNotCreated struct {
	e error
}

func (err *ErrorBinaryNotCreated) Error() string {
	return fmt.Sprintf("storage: binary not created: %s", err.e)
}

func NewErrorBinaryNotCreated(e error) (err error) {
	return &ErrorBinaryNotCreated{
		e: e,
	}
}

type ErrorBinaryNotFoundByKeyword struct {
	keyword string
	userID  uuid.UUID
	e       error
}

func (err *ErrorBinaryNotFoundByKeyword) Error() string {
	return fmt.Sprintf("storage: binary with keyword %s not found: %s", err.keyword, err.e)
}

func NewErrorBinaryNotFoundByKeyword(keyword string, userID uuid.UUID, e error) (err error) {
	return &ErrorBinaryNotFoundByKeyword{
		keyword: keyword,
		userID:  userID,
		e:       e,
	}
}

type ErrorNoBinaryFoundForUser struct {
	userID uuid.UUID
}

func (err *ErrorNoBinaryFoundForUser) Error() string {
	return fmt.Sprintf("storage: no binary found for user with ID %s", err.userID)
}

func NewErrorNoBinaryFoundForUser(userID uuid.UUID) (err error) {
	return &ErrorNoBinaryFoundForUser{
		userID: userID,
	}
}
