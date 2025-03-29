package handler

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/nilhiu/srleaderboard/internal/app/ajax"
	"github.com/nilhiu/srleaderboard/internal/app/htmx"
	"github.com/nilhiu/srleaderboard/internal/service/user"
)

// AuthHandler provides methods to handle HTTP requests about authetication.
type AuthHandler struct {
	ctx     context.Context
	userSvc user.Service
}

// NewAuthHandler creates a new AuthHandler.
func NewAuthHandler(ctx context.Context, userSvc user.Service) *AuthHandler {
	return &AuthHandler{
		ctx:     ctx,
		userSvc: userSvc,
	}
}

// Login is a handler, which authenticates an user based on if the provided
// username and password ([ajax.LoginRequest]) was correct. Supports both
// AJAX and HTMX requests (`Hx-Request`). Returns a JWT token in
// the response's `Set-Cookie` header field.
func (h *AuthHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req ajax.LoginRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if r.Header.Get("Hx-Request") != "true" {
			h.loginAJAX(req).ServeHTTP(w, r)
		} else {
			h.loginHTMX(req).ServeHTTP(w, r)
		}
	}
}

// Register is a handler, which registers an user based on if the provided
// username, email, and password ([ajax.RegisterRequest]) was correct. Supports
// both AJAX and HTMX requests (`Hx-Request`). Returns a JWT token in
// the response's `Set-Cookie` header field.
func (h *AuthHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req ajax.RegisterRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if r.Header.Get("Hx-Request") != "true" {
			h.registerAJAX(req).ServeHTTP(w, r)
		} else {
			h.registerHTMX(req).ServeHTTP(w, r)
		}
	}
}

// LogOut is a handler, which log an user out. Doesn't require the user to be
// actually authenticated, as it only deletes the JWT cookie. Supports both
// AJAX and HTMX requests (`Hx-Request`).
func (h *AuthHandler) Logout() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Hx-Request") != "true" {
			h.logoutAJAX()
		} else {
			h.logoutHTMX()
		}
	}
}

func (h *AuthHandler) loginAJAX(req ajax.LoginRequest) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokStr, err := h.userSvc.Login(req.Username, req.Password)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		writeJWTCookie(w, tokStr)
		w.WriteHeader(http.StatusOK)
	}
}

func (h *AuthHandler) loginHTMX(req ajax.LoginRequest) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokStr, err := h.userSvc.Login(req.Username, req.Password)
		if err != nil {
			htmx.NewTrigger().Add("auth").
				AlertError("failed to log in, incorrect information").
				Write(w)
			w.WriteHeader(http.StatusUnauthorized)
		} else {
			writeJWTCookie(w, tokStr)

			htmx.NewTrigger().Add("auth").
				AlertOK("logged in successfully as " + req.Username).
				Write(w)
			w.WriteHeader(http.StatusNoContent)
		}
	}
}

func (h *AuthHandler) registerAJAX(req ajax.RegisterRequest) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokStr, err := h.userSvc.Register(req.Username, req.Email, req.Password)
		if err != nil {
			if errors.Is(err, user.ErrRegisterUserExists) {
				w.WriteHeader(http.StatusConflict)
			}
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		writeJWTCookie(w, tokStr)
		w.WriteHeader(http.StatusOK)
	}
}

func (h *AuthHandler) registerHTMX(req ajax.RegisterRequest) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokStr, err := h.userSvc.Register(req.Username, req.Email, req.Password)
		if err != nil {
			if errors.Is(err, user.ErrRegisterUserExists) {
				htmx.NewTrigger().Add("auth").
					AlertError("user with that username already exists").
					Write(w)
				w.WriteHeader(http.StatusConflict)
			}

			w.WriteHeader(http.StatusInternalServerError)
		} else {
			writeJWTCookie(w, tokStr)

			htmx.NewTrigger().Add("auth").
				AlertOK("successfully registered").
				Write(w)
			w.WriteHeader(http.StatusNoContent)
		}
	}
}

func (h *AuthHandler) logoutAJAX() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		deleteJWTCookie(w)
		w.WriteHeader(http.StatusOK)
	}
}

func (h *AuthHandler) logoutHTMX() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		deleteJWTCookie(w)

		htmx.NewTrigger().Add("auth").
			AlertOK("logged out successfully").
			Write(w)
		w.WriteHeader(http.StatusNoContent)
	}
}

func writeJWTCookie(w http.ResponseWriter, token string) {
	http.SetCookie(w, &http.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
		Path:     "/",
	})
}

func deleteJWTCookie(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:     "jwt",
		Value:    "",
		MaxAge:   -1,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		Path:     "/",
	})
}
