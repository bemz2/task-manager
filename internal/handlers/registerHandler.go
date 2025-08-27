package handlers

import (
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"task-manager/internal/db"
)

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func RegisterHandler(conn *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req RegisterRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil || req.Username == "" || req.Password == "" {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}
		hash := sha256.Sum256([]byte(req.Password))
		passwordHash := hex.EncodeToString(hash[:])

		err = db.CreateUser(conn, req.Username, passwordHash)

		if err != nil {
			http.Error(w, "User creation failed", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("User created"))
	}
}
