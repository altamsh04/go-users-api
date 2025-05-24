package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/altamsh04/go-users-api/internal/database"
	"github.com/gorilla/mux"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	Mobile   string `json:"mobile"`
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<center><h1>Welcome To Users API Backend!</h1></center>")
}

func fetchAllUsers(w http.ResponseWriter, r *http.Request) {
	rows, err := database.DB.Query("SELECT id, username, fullname, email, mobile FROM users")
	if err != nil {
		http.Error(w, "Database query failed", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	users := []User{}

	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Username, &user.Fullname, &user.Email, &user.Mobile)
		if err != nil {
			http.Error(w, "Error reading data", http.StatusInternalServerError)
			return
		}
		users = append(users, user)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func addNewUser(w http.ResponseWriter, r *http.Request) {
	var user User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if user.Username == "" || user.Fullname == "" || user.Email == "" || user.Mobile == "" {
		http.Error(w, "Please provide username, fullname, email and mobile number", http.StatusBadRequest)
		return
	}

	query := `INSERT INTO users (username, fullname, email, mobile) VALUES (?, ?, ?, ?)`
	_, err = database.DB.Exec(query, user.Username, user.Fullname, user.Email, user.Mobile)
	if err != nil {
		http.Error(w, "Failed to insert user", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "User added successfully",
	})

}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["userID"]
	if userID == "" {
		http.Error(w, "userID is required", http.StatusBadRequest)
		return
	}

	query := `DELETE FROM users WHERE id = ?`
	result, err := database.DB.Exec(query, userID)
	if err != nil {
		http.Error(w, "Failed to delete user from database.", http.StatusInternalServerError)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		http.Error(w, "Failed to confirm deletion.", http.StatusInternalServerError)
		return
	}
	if rowsAffected == 0 {
		http.Error(w, "No user found with that userid.", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "User deleted successfully",
	})
}
