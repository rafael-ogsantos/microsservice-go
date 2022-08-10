package main

import (
	"auth/data"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtkey = []byte("my_secret_key")

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

func (app *Config) SignIn(w http.ResponseWriter, r *http.Request) {
	var creds Credentials
	err := app.readJSON(w, r, &creds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := app.Models.User.GetByEmail(creds.Email)

	checkPassword, err := app.Models.User.PasswordMatches(user.Password, creds.Password)
	if !checkPassword {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	// if !ok || expectedPassword != creds.Password {
	// 	w.WriteHeader(http.StatusUnauthorized)
	// 	return
	// }

	expirationTime := time.Now().Add(5 * time.Minute)

	claims := &Claims{
		Email: creds.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtkey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})

	type jsonResponse struct {
		Error   bool   `json:"error"`
		Message string `json:"message"`
		// Token   string `json:"token"`
	}

	// c, err := r.Cookie("token")
	// if err != nil {
	// 	log.Println(err)
	// 	if err == http.ErrNoCookie {
	// 		w.WriteHeader(http.StatusUnauthorized)
	// 		return
	// 	}
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	return
	// }

	var resp jsonResponse
	resp.Error = false
	resp.Message = "Logged Successfully"
	// resp.Token = c.Value

	app.writeJSON(w, http.StatusAccepted, resp)
}

func (app *Config) SignUp(w http.ResponseWriter, r *http.Request) {
	var requestPayload data.User

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	userID, err := app.Models.User.Insert(requestPayload)
	if err != nil {
		app.errorJSON(w, errors.New("failed to register a user"), http.StatusBadRequest)
		return
	}

	fmt.Println(userID)

	payload := jsonResponse{
		Error: false,

		Message: "User with",
	}

	app.writeJSON(w, http.StatusAccepted, payload)
}

// func Welcome(w http.ResponseWriter, r *http.Request) {
// 	c, err := r.Cookie("token")
// 	if err != nil {
// 		if err == http.ErrNoCookie {
// 			w.WriteHeader(http.StatusUnauthorized)
// 			return
// 		}
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}

// 	tknStr := c.Value

// 	claims := &Claims{}

// 	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(t *jwt.Token) (interface{}, error) {
// 		return jwtkey, nil
// 	})

// 	if err != nil {
// 		if err == jwt.ErrSignatureInvalid {
// 			w.WriteHeader(http.StatusUnauthorized)
// 			return
// 		}
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}
// 	if !tkn.Valid {
// 		w.WriteHeader(http.StatusUnauthorized)
// 		return
// 	}
// 	w.Write([]byte(fmt.Sprintf("Welcome %s!", claims.Username)))
// }
