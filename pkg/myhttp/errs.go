package myhttp

import (
	"fmt"
	"net/http"
)

type ErrorHTTP struct {
	Status int
	Data   string
}

func NewErrorHTTP(status int, data string) *ErrorHTTP {
	return &ErrorHTTP{status, data}
}

func (e *ErrorHTTP) Error() string {
	return e.Data
}

var serverErrorMessages = map[int]string{
	http.StatusInternalServerError: "internal server error. Please try again later",
	http.StatusBadGateway:          "bad gateway. The server received an invalid response from an upstream server",
	http.StatusServiceUnavailable:  "service unavailable. The server is currently overloaded or under maintenance. Please try again later",
	http.StatusGatewayTimeout:      "gateway timeout. The server took too long to respond",
}

func (e *ErrorHTTP) CheckServerErrors() error {
	if val, exists := serverErrorMessages[e.Status]; exists {
		return fmt.Errorf("%s", val)
	}

	return fmt.Errorf("%s", "something went wrong... Please try again later")
}

func ParseError(err error) *ErrorHTTP {
	if errHTTP, ok := err.(*ErrorHTTP); ok {
		return errHTTP
	}

	msg := fmt.Sprintf("fail to parse error:\n%v\n", err)
	return NewErrorHTTP(http.StatusInternalServerError, msg)
}
