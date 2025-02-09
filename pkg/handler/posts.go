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