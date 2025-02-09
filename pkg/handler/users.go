package handler

import (
	"encoding/json"
    "net/http"
    "go-web-service/pkg/model" // Import the model package
)

func UsersHandler(w http.ResponseWriter, r *http.Request) {
    // Sample user data
    users := []model.User{
        {Name: "Ali", Email: "ali@example.com"},
        {Name: "Sara", Email: "sara@example.com"},
    }

    // Set content-type to JSON
    w.Header().Set("Content-Type", "application/json")

    // Encode users slice to JSON and write to response
    json.NewEncoder(w).Encode(users)
}
