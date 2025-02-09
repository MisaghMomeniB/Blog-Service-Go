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

// DeletePostHandler handles deleting a blog post by ID
func DeletePostHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodDelete {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        return
    }

    id := r.URL.Query().Get("id")
    if id == "" {
        http.Error(w, "Missing post ID", http.StatusBadRequest)
        return
    }

    var postID int
    fmt.Sscanf(id, "%d", &postID)

    var index = -1
    for i, post := range posts {
        if post.ID == postID {
            index = i
            break
        }
    }

    if index == -1 {
        http.Error(w, "Post not found", http.StatusNotFound)
        return
    }