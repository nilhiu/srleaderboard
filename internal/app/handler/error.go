package handler

import (
	"errors"
	"net/http"
)

var (
	HTTPErrNotFound            = errors.New("Not Found")
	HTTPErrInternalServerError = errors.New("Internal Server Error")
)

func WriteErrorHeader(w http.ResponseWriter, err error) {
	if errors.Is(err, HTTPErrNotFound) {
		w.WriteHeader(http.StatusNotFound)
	} else if errors.Is(err, HTTPErrInternalServerError) {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
