package controllers

import (
	"encoding/json"
	//"log"
	"fmt"
	"time"
	//"error"
	"net/http"
	"database/sql"
	"github.com/gorilla/mux"
	"blog/models"
	"blog/utils"
	"blog/const"
	database "blog/utils"
	"github.com/dgrijalva/jwt-go"
)

// PostController handles blog post related operations
type PostController struct{}

func AuthenticateUser(username, password string) (bool, error) {
    // Establish a database connection
    db := database.GetDB()
    defer db.Close()

    // Prepare a query to fetch the user with the provided username and verify the MD5-hashed password
    query := "SELECT 1 FROM User WHERE Username = ? AND Password = MD5(?)"

    fmt.Println("Query:", query,"=",username, "=",password)

    var result int
    err := db.QueryRow(query, username, password).Scan(&result)
    fmt.Println("Error:", err)

    if err != nil {
        if err == sql.ErrNoRows {
            // User not found
            return false, nil
        }
        return false, err
    }
    // If a result is found, the authentication is successful
    return true, nil
}

func GenerateJWTToken(user model.User) (string, error) {
    // If authentication is successful, generate a JWT token
    token := jwt.New(jwt.SigningMethodHS256)
    claims := token.Claims.(jwt.MapClaims)
    claims["username"] = user.Username

    // Set expiration time based on JWTTokensExpiry
    claims["exp"] = time.Now().Add(constants.JWTTokensExpiry).Unix()

    tokenString, err := token.SignedString(constants.JWTSecret)

    if err != nil {
        return "", err
    }

    return tokenString, nil
}

func Login(w http.ResponseWriter, r *http.Request) {
    // Parse user credentials from the request
    var user model.User
    err := json.NewDecoder(r.Body).Decode(&user)
    if err != nil {
        http.Error(w, "Failed to parse user credentials", http.StatusBadRequest)
        return
    }

    // Authenticate the user and validate credentials
    authenticated, err := AuthenticateUser(user.Username, user.Password)
    if err != nil {
        http.Error(w, "Authentication Invalid!", http.StatusUnauthorized)
        return
    }

    if !authenticated {
        http.Error(w, "Authentication failed!", http.StatusUnauthorized)
        return
    }

    // If authentication is successful, generate a JWT token
    token, err := GenerateJWTToken(user)
    if err != nil {
        http.Error(w, "Failed to generate token", http.StatusUnauthorized)
        return
    }

    // Return the token in the response
    response := map[string]string{"token": token}
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}

// ListPosts returns a list of all blog posts
func (pc PostController) ListPosts(w http.ResponseWriter, r *http.Request) {

    // Establish a database connection
    db := utils.GetDB()
    defer db.Close()

    // Execute a SQL query to retrieve all posts
    rows, err := db.Query("SELECT postID,title,content,author FROM BlogPosts")
    if err != nil {
        http.Error(w, "Failed to retrieve posts", http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    // Create a slice to store the retrieved posts
    var posts []model.Post

    // Iterate through the rows and scan data into the Post struct
    for rows.Next() {
        var post model.Post
        if err := rows.Scan(&post.ID, &post.Title, &post.Content, &post.Author); err != nil {
					//fmt.Println("err:fetch:", err)
            http.Error(w, "Failed to scan posts", http.StatusInternalServerError)
            return
        }
        posts = append(posts, post)
    }
    // Check for errors during row iteration
    if err := rows.Err(); err != nil {
        http.Error(w, "Failed to retrieve posts", http.StatusInternalServerError)
        return
    }

    // Convert posts to JSON and send it in the response
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(posts)
}

// GetPost returns a specific blog post by ID
func (pc PostController) GetSinglePost(w http.ResponseWriter, r *http.Request) {
    // Get the post ID from the request parameters
    vars := mux.Vars(r)
    postID, ok := vars["id"]
    if !ok {
        http.Error(w, "Invalid post ID", http.StatusBadRequest)
        return
    }
    // Establish a database connection
    db := utils.GetDB()
    defer db.Close()

    // Execute a SQL query to retrieve the post by ID
    query := "SELECT postid,title,content,author FROM BlogPosts WHERE postid = ?"
    row := db.QueryRow(query, postID)

    var post model.Post
    err := row.Scan(&post.ID, &post.Title, &post.Content, &post.Author)

    if err != nil {
        if err == sql.ErrNoRows {
            http.Error(w, "Post not found", http.StatusNotFound)
            return
        }
        http.Error(w, "Failed to retrieve the post", http.StatusInternalServerError)
        return
    }

    // Convert the post to JSON and send it in the response
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(post)
}

//creates a new blog post
func (pc PostController) CreatePost(w http.ResponseWriter, r *http.Request) {
    // Establish a database connection
    db := utils.GetDB()
    defer db.Close()

    // Parse the request body to create a new post
    var newPost model.Post
    if err := json.NewDecoder(r.Body).Decode(&newPost); err != nil {
        http.Error(w, "Invalid request data", http.StatusBadRequest)
        return
    }
    // Validate the new post data
    if newPost.Title == "" || newPost.Content == "" || newPost.Author == "" {
        http.Error(w, "Title, content, and author are required fields", http.StatusBadRequest)
        return
    }

    // database insertion
    query := "INSERT INTO BlogPosts (title, content, author, CreatedAt, UpdatedAt) VALUES (?, ?, ?, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)"
    result, err := db.Exec(query, newPost.Title, newPost.Content, newPost.Author)
		fmt.Println("error insert:", err)
    if err != nil {
        http.Error(w, "Failed to create the post", http.StatusInternalServerError)
        return
    }

    // Get the ID of the newly created post
    newPostID, _ := result.LastInsertId()
    newPost.ID = int(newPostID)

    // Return the newly created post in the response
    w.WriteHeader(http.StatusCreated)
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(newPost)
}

// updates an existing blog post
func (pc PostController) UpdatePost(w http.ResponseWriter, r *http.Request) {
    // Get the post ID from the request parameters
    vars := mux.Vars(r)
    postID, ok := vars["id"]
    if !ok {
        http.Error(w, "Invalid post ID", http.StatusBadRequest)
        return
    }
    // Establish a database connection
    db := database.GetDB()
    defer db.Close()

    // Parse the request body to update the post
    var updatedPost model.Post
    if err := json.NewDecoder(r.Body).Decode(&updatedPost); err != nil {
        http.Error(w, "Invalid request data", http.StatusBadRequest)
        return
    }

    // Validate the updated post data (e.g., required fields)
    if updatedPost.Title == "" || updatedPost.Content == "" || updatedPost.Author == "" {
        http.Error(w, "Title, content, and author are required fields", http.StatusBadRequest)
        return
    }

    // Implement database update logic
    query := "UPDATE blogposts SET title = ?, content = ?, author = ?, UpdatedAt = CURRENT_TIMESTAMP WHERE postid = ?"
    _, err := db.Exec(query, updatedPost.Title, updatedPost.Content, updatedPost.Author, postID)
		fmt.Println("error:update", err)
    if err != nil {
        http.Error(w, "Failed to update the post", http.StatusInternalServerError)
        return
    }

    // Return a success response
    w.WriteHeader(http.StatusOK) // StatusOK
}

// DeletePost deletes a blog post by ID
func (pc PostController) DeletePost(w http.ResponseWriter, r *http.Request) {
    // Get the post ID from the request parameters
    vars := mux.Vars(r)
    postID, ok := vars["id"]
    if !ok {
        http.Error(w, "Invalid post ID", http.StatusBadRequest)
        return
    }
    // Establish a database connection
    db := database.GetDB()
    defer db.Close()

    // Execute a SQL query to delete the post by ID
    query := "DELETE FROM blogposts WHERE postid = ?"
    _, err := db.Exec(query, postID)
    if err != nil {
        http.Error(w, "Failed to delete the post", http.StatusInternalServerError)
        return
    }

    // Return a success response
    w.WriteHeader(http.StatusOK) // StatusOK
}
