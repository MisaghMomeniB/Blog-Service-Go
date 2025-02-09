package handler

import (
    "encoding/json"
    "fmt"
    "net/http"
    "go-web-service/pkg/model"
    "time"
)

// In-memory store for posts (a simple example)
var posts = []model.Post{
    {ID: 1, Title: "First Post", Content: "This is the first blog post.", CreatedAt: time.Now()},
}

// PostsHandler handles fetching all blog posts
func PostsHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(posts)
}

// CreatePostHandler handles creating a new blog post
func CreatePostHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        return
    }

	var newPost model.Post
    err := json.NewDecoder(r.Body).Decode(&newPost)
    if err != nil {
        http.Error(w, "Unable to parse request body", http.StatusBadRequest)
        return
    }

	    // Set the ID and created time
		newPost.ID = len(posts) + 1
		newPost.CreatedAt = time.Now()
	
		// Save the new post
		posts = append(posts, newPost)
	
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(newPost)
	}