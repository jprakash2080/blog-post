package model

// Post represents a blog post
type Post struct {
	ID      int    `json:"postid"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

type User struct {
    Username string `json:"username"`
    Password string `json:"password"`
}

// Implement database interactions and CRUD operations here
