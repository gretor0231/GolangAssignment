package main

import (
	"net/http"
)

func (app *Config) Broker(w http.ResponseWriter, r *http.Request) {
	payload := jsonResponse{
		Error:   false,
		Message: "Hit the broker",
	}

	_ = app.writeJSON(w, http.StatusOK, payload)
}

func (app *Config) Add(w http.ResponseWriter, r *http.Request) {
	payload := jsonResponse{
		Error:   false,
		Message: "Add a person to the list",
	}

	_ = app.writeJSON(w, http.StatusOK, payload)
}

func (app *Config) Remove(w http.ResponseWriter, r *http.Request) {
	payload := jsonResponse{
		Error:   false,
		Message: "Remove a person to the list",
	}

	_ = app.writeJSON(w, http.StatusOK, payload)
}

func (app *Config) Query(w http.ResponseWriter, r *http.Request) {
	payload := jsonResponse{
		Error:   false,
		Message: "Query a person to the list",
	}

	_ = app.writeJSON(w, http.StatusOK, payload)
}
