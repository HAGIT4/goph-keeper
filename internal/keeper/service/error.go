package service

type ErrorInternal struct{}

func (err *ErrorInternal) Error() string {
	return "Internal error"
}

type ErrorUnauthenticated struct{}

func (err *ErrorUnauthenticated) Error() string {
	return "User unathenticated"
}

type ErrorUserNotRegistered struct{}

func (err *ErrorUserNotRegistered) Error() string {
	return "User not registered"
}

type ErrorTokenExpired struct{}

func (err *ErrorTokenExpired) Error() string {
	return "Token expired"
}
