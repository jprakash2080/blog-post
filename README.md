# Sample API Authentication with `curl`

This guide demonstrates how to authenticate with an API to obtain an access token using `curl`. In this example, we'll use the `/login` endpoint to obtain a JWT token for further API requests.

## Authentication

To authenticate and obtain a token, send a POST request to the `/login` endpoint with your username and password.

Replace `<your-username>` and `<your-password>` with your actual credentials.

```bash
curl --location 'http://localhost:8080/login' \
--header 'Content-Type: application/json' \
--data '{
    "username": "<your-username>",
    "password": "<your-password>"
}'

```

# Create a New Blog Post with JWT Token

This guide demonstrates how to use `curl` to create a new blog post through an API with authentication. In this example, we'll send a POST request to the `/auth/posts` endpoint to create a new blog post. The request includes an access token in the `Authorization` header for authentication.

## Create a New Blog Post

To create a new blog post, send a POST request to the `/auth/posts` endpoint. Be sure to include your JWT token in the `Authorization` header, along with the title, content, and author of the post.

Replace `<your-jwt-token>`, `<your-title>`, `<your-content>`, and `<your-author>` with your actual token and post data.

```bash
curl --location 'http://localhost:8080/auth/posts' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer <your-jwt-token>' \
--data '{
  "title": "<your-title>",
  "content": "<your-content>",
  "author": "<your-author>"
}'

```
# Update a Blog Post with JWT Token

This guide demonstrates how to use `curl` to update an existing blog post through an API with authentication. In this example, we'll send a PUT request to the `/auth/posts/{id}` endpoint to update a blog post. The request includes an access token in the `Authorization` header for authentication.

## Update a Blog Post

To update a blog post, send a PUT request to the `/auth/posts/{id}` endpoint. Be sure to include your JWT token in the `Authorization` header, along with the updated title, content, and author.

Replace `<your-jwt-token>`, `<post-id>`, `<updated-title>`, `<updated-content>`, and `<updated-author>` with your actual token, post ID, and updated post data.

```bash
curl --location --request PUT 'http://localhost:8080/auth/posts/<post-id>' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer <your-jwt-token>' \
--data '{
  "title": "<updated-title>",
  "content": "<updated-content>",
  "author": "<updated-author>"
}'

```
# Retrieve a List of Blog Posts with JWT Token

This guide demonstrates how to use `curl` to retrieve a list of all blog posts through an API with authentication. In this example, we'll send a GET request to the `/auth/fetchposts` endpoint to fetch all blog posts. The request includes an access token in the `Authorization` header for authentication.

## Retrieve a List of Blog Posts

To retrieve a list of all blog posts, send a GET request to the `/auth/fetchposts` endpoint. Be sure to include your JWT token in the `Authorization` header.

Replace `<your-jwt-token>` with your actual token.

```bash
curl --location 'http://localhost:8080/auth/fetchposts' \
--header 'Authorization: Bearer <your-jwt-token>'

```

# Retrieve a Specific Blog Post by ID with JWT Token

This guide demonstrates how to use `curl` to retrieve a specific blog post by its ID through an API with authentication. In this example, we'll send a GET request to the `/auth/getpost/{id}` endpoint to fetch a blog post by its unique ID. The request includes an access token in the `Authorization` header for authentication.

## Retrieve a Specific Blog Post by ID

To retrieve a specific blog post, send a GET request to the `/auth/getpost/{id}` endpoint, where `{id}` should be replaced with the ID of the post you want to fetch. Ensure you include your JWT token in the `Authorization` header.

Replace `<your-jwt-token>` with your actual token and `<post-id>` with the ID of the blog post you want to retrieve.

```bash
curl --location 'http://localhost:8080/auth/getpost/<post-id>' \
--header 'Authorization: Bearer <your-jwt-token>'

```

# Retrieve a Specific Blog Post by ID with JWT Token

This guide demonstrates how to use `curl` to retrieve a specific blog post by its ID through an API with authentication. In this example, we'll send a GET request to the `/auth/getpost/{id}` endpoint to fetch a blog post by its unique ID. The request includes an access token in the `Authorization` header for authentication.


## Retrieve a Specific Blog Post by ID

To retrieve a specific blog post, send a GET request to the `/auth/getpost/{id}` endpoint, where `{id}` should be replaced with the ID of the post you want to fetch. Ensure you include your JWT token in the `Authorization` header.

Replace `<your-jwt-token>` with your actual token and `<post-id>` with the ID of the blog post you want to retrieve.

```bash
curl --location 'http://localhost:8080/auth/getpost/<post-id>' \
--header 'Authorization: Bearer <your-jwt-token>'

```
