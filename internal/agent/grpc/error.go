package grcp

type ErrorNoAuthToken struct{}

func (err *ErrorNoAuthToken) Error() string {
	return "No auth token found"
}
