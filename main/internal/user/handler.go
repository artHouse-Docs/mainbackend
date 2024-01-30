package user

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = user.Register(ctx)
	if err != nil {
		switch err.Error() {
		case "user email already exists":
			w.WriteHeader(http.StatusConflict)
			return
		case "database unavailable":
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	refreshCookie := http.Cookie{
		Name:     "refreshToken",
		Value:    user.RefreshToken,
		Path:     "/",
		MaxAge:   0,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	}

	accessToken := http.Cookie{
		Name:     "accessToken",
		Value:    user.AccessToken,
		Path:     "/",
		MaxAge:   0,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	}

	http.SetCookie(w, &refreshCookie)
	http.SetCookie(w, &accessToken)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(user)
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var user User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	result, err := user.Login(ctx)
	if err != nil {
		switch err.Error() {
		case "user not found":
			w.WriteHeader(http.StatusNotFound)
			return
		case "wrong password":
			w.WriteHeader(http.StatusUnauthorized)
			return
		default:
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Println(err.Error())
			return
		}
	}

	if result {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(user)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
