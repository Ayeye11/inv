package myhttp

import (
	"log"
	"net/http"
)

func SendError(w http.ResponseWriter, err error) error {
	errHTTP := ParseError(err)

	if errHTTP.Status < 500 {
		return SendJSON(w, errHTTP.Status, map[string]string{"error": errHTTP.Data})
	}

	log.Println(errHTTP.Data)
	return SendJSON(w, errHTTP.Status, map[string]string{"error": errHTTP.CheckServerErrors().Error()})
}

func SendMessage(w http.ResponseWriter, status int, message string) error {
	return SendJSON(w, status, map[string]string{"message": message})
}
