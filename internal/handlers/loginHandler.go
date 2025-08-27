package handlers

import (
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"task-manager/internal/db"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func LoginHandler(conn *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req LoginRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil || req.Username == "" || req.Password == "" {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}
		storedHash, err := db.GetPasswordHash(conn, req.Username)
		if err != nil {
			http.Error(w, "User not found", http.StatusUnauthorized)
			return
		}
		hash := sha256.Sum256([]byte(req.Password))
		passwordHash := hex.EncodeToString(hash[:])
		if passwordHash != storedHash {
			http.Error(w, "Invalid username or password", http.StatusUnauthorized)
			return
		}

		w.Write([]byte("Login successful"))
	}
}
