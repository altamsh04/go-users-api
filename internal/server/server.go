package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {
    r := mux.NewRouter()

    // Routes
    r.HandleFunc("/", handleHome).Methods("GET")
    r.HandleFunc("/v1/users", fetchAllUsers).Methods("GET")
    r.HandleFunc("/v1/users", addNewUser).Methods("POST")
    r.HandleFunc("/v1/users/{userID}", deleteUser).Methods("DELETE")

    fmt.Println("Server listening at port 8000")
    if err := http.ListenAndServe(":8000", r); err != nil {
        fmt.Println("Failed to start server:", err)
    }
}
