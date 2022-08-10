package main

import (
	"broker/event"
	"encoding/json"
	"net/http"
	"time"
)

type RequestPayload struct {
	Action string      `json:"action"`
	Auth   AuthPayload `json:"auth,omitempty"`
}

type AuthPayload struct {
	Name string      `json:"name"`
	Data AuthService `json:"data"`
}

type AuthService struct {
	Email      string    `json:"email,omitempty"`
	FirstName  string    `json:"first_name,omitempty"`
	LastName   string    `json:"last_name,omitempty"`
	Password   string    `json:"password,omitempty"`
	UserActive int       `json:"user_active,omitempty"`
	CreatedAt  time.Time `json:"created_at,omitempty"`
	UpdatedAt  time.Time `json:"updated_at,omitempty"`
}

func (app *Config) Broker(w http.ResponseWriter, r *http.Request) {
	payload := jsonResponse{
		Error:   false,
		Message: "Hit the broker",
	}

	_ = app.writeJSON(w, http.StatusOK, payload)
}

func (app *Config) HandleSubmission(w http.ResponseWriter, r *http.Request) {
	var requestPayload RequestPayload

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	switch requestPayload.Action {
	case "auth":
		app.authenticate(w, requestPayload.Auth)
	case "signup":
		app.signup(w, requestPayload.Auth)
	}
}

func (app *Config) authenticate(w http.ResponseWriter, a AuthPayload) {
	err := app.pushToQueue(a, "auth.REGISTER")
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	var payload jsonResponse
	payload.Error = false
	payload.Message = "User Authenticated"

	app.writeJSON(w, http.StatusAccepted, payload)
}

func (app *Config) signup(w http.ResponseWriter, a AuthPayload) {
	err := app.pushToQueue(a, "auth.SIGNUP")
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	var payload jsonResponse
	payload.Error = false
	payload.Message = "User Login"

	app.writeJSON(w, http.StatusAccepted, payload)
}

func (app *Config) pushToQueue(a AuthPayload, name string) error {
	emitter, err := event.NewEventEmitter(app.Rabbit)
	if err != nil {
		return err
	}

	j, _ := json.MarshalIndent(&a, "", "\t")
	err = emitter.Push(string(j), name)
	if err != nil {
		return err
	}

	return nil
}
