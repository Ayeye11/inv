package myhttp

import (
	"fmt"
	"log"
	"net/http"
)

func SendError(w http.ResponseWriter, err error) error {
	errHTTP, ok := err.(*ErrorHTTP)
	if !ok {
		log.Println(fmt.Errorf("fail to parse error:\n%v", err))
		return SendJSON(w, http.StatusInternalServerError, map[string]string{"error": serverErrorMessages[500]})
	}

	if errHTTP.status < 500 {
		return SendJSON(w, errHTTP.status, map[string]string{"error": errHTTP.data})
	}

	log.Println(errHTTP.data)
	return SendJSON(w, errHTTP.status, map[string]string{"error": errHTTP.CheckServerErrors().Error()})
}

func SendMessage(w http.ResponseWriter, status int, message string) error {
	return SendJSON(w, status, map[string]string{"message": message})
}
