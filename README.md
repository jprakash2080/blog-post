  ## Database Setup

  To get started with MyBlog, you can set up the required database and tables. Use the following steps:

  1. Navigate to the 'SQL' folder in this repository.

  2. Locate the 'database.sql' file in the 'SQL' folder.

  3. Run the 'database.sql' script in your MySQL database management tool to create the database, tables, and sample data.

  # SQL Schema for MyBlog Database

  This repository contains the SQL schema for the MyBlog database, which includes tables for blog posts and user authentication.

  ## Database Details

  - **Database Name**: myblog
  - **Character Set**: utf8mb4
  - **Collation**: utf8mb4_0900_ai_ci

  ## BlogPosts Table Definition

  The `BlogPosts` table stores information about blog posts, including their titles, content, authors, creation timestamps, and update timestamps.

  ### Columns:

  - `PostID` (int): A unique identifier for each blog post.
  - `Title` (varchar): The title of the blog post (up to 255 characters).
  - `Content` (text): The content of the blog post.
  - `Author` (varchar): The author of the blog post (up to 100 characters).
  - `CreatedAt` (timestamp): The timestamp when the blog post was created.
  - `UpdatedAt` (timestamp): The timestamp of the last update to the blog post.

  ## Usage

  The `BlogPosts` table schema is suitable for managing blog posts in your application. You can create this table by executing the provided SQL command.

  To create the `BlogPosts` table, use the following SQL command:

  ```sql
  CREATE TABLE `BlogPosts` (
    `PostID` int NOT NULL AUTO_INCREMENT,
    `Title` varchar(255) NOT NULL,
    `Content` text NOT NULL,
    `Author` varchar(100) NOT NULL,
    `CreatedAt` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `UpdatedAt` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`PostID`)
  ) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
  ```
  ## User Table Definition

  The `User` table is used for user authentication and stores information about registered users.

  ### Columns:

  - `UserID` (int): A unique identifier for each user.
  - `Username` (varchar): The username of the user (up to 255 characters).
  - `Password` (varchar): The hashed password of the user.

  ## Usage

  The `User` table schema is suitable for managing user information in your application. You can create this table by executing the provided SQL command.

  To create the `User` table, use the following SQL command:

  ```sql
  CREATE TABLE `User` (
    `UserID` int NOT NULL AUTO_INCREMENT,
    `Username` varchar(255) NOT NULL,
    `Password` varchar(255) NOT NULL,
    PRIMARY KEY (`UserID`)
  ) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
  ```
  ## Usage

  You can use the provided SQL schema to create the MyBlog database and its tables. The database and tables are designed for managing blog posts and user authentication in your application.

  To create the `myblog` database, use the following SQL command:

  ```sql
  CREATE DATABASE `myblog` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;

  ```


  # `API API endpoint`
  # API Authentication with `curl`

  This guide demonstrates how to authenticate with an API to obtain an access token using `curl`. In this example, we'll use the `/login` endpoint to obtain a token for further API requests.

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

  # Create a New Blog Post with Auth Token

  This guide demonstrates how to use `curl` to create a new blog post through an API with authentication. In this example, we'll send a POST request to the `/auth/posts` endpoint to create a new blog post. The request includes an access token in the `Authorization` header for authentication.

  ## Create a New Blog Post

  To create a new blog post, send a POST request to the `/auth/posts` endpoint. Be sure to include your Auth token in the `Authorization` header, along with the title, content, and author of the post.

  Replace `<your-auth-token>`, `<your-title>`, `<your-content>`, and `<your-author>` with your actual token and post data.

  ```bash
  curl --location 'http://localhost:8080/auth/posts' \
  --header 'Content-Type: application/json' \
  --header 'Authorization: <your-auth-token>' \
  --data '{
    "title": "<your-title>",
    "content": "<your-content>",
    "author": "<your-author>"
  }'

  ```
  # Update a Blog Post with Auth Token

  This guide demonstrates how to use `curl` to update an existing blog post through an API with authentication. In this example, we'll send a PUT request to the `/auth/posts/{id}` endpoint to update a blog post. The request includes an access token in the `Authorization` header for authentication.

  ## Update a Blog Post

  To update a blog post, send a PUT request to the `/auth/posts/{id}` endpoint. Be sure to include your Auth token in the `Authorization` header, along with the updated title, content, and author.

  Replace `<your-auth-token>`, `<post-id>`, `<updated-title>`, `<updated-content>`, and `<updated-author>` with your actual token, post ID, and updated post data.

  ```bash
  curl --location --request PUT 'http://localhost:8080/auth/posts/<post-id>' \
  --header 'Content-Type: application/json' \
  --header 'Authorization: <your-auth-token>' \
  --data '{
    "title": "<updated-title>",
    "content": "<updated-content>",
    "author": "<updated-author>"
  }'

  ```
  # Retrieve a List of Blog Posts with Auth Token

  This guide demonstrates how to use `curl` to retrieve a list of all blog posts through an API with authentication. In this example, we'll send a GET request to the `/auth/fetchposts` endpoint to fetch all blog posts. The request includes an access token in the `Authorization` header for authentication.

  ## Retrieve a List of Blog Posts

  To retrieve a list of all blog posts, send a GET request to the `/auth/fetchposts` endpoint. Be sure to include your Auth token in the `Authorization` header.

  Replace `<your-Auth-token>` with your actual token.

  ```bash
  curl --location 'http://localhost:8080/auth/fetchposts' \
  --header 'Authorization: <your-Auth-token>'

  ```

  # Retrieve a Specific Blog Post by ID with Auth Token

  This guide demonstrates how to use `curl` to retrieve a specific blog post by its ID through an API with authentication. In this example, we'll send a GET request to the `/auth/getpost/{id}` endpoint to fetch a blog post by its unique ID. The request includes an access token in the `Authorization` header for authentication.

  ## Retrieve a Specific Blog Post by ID

  To retrieve a specific blog post, send a GET request to the `/auth/getpost/{id}` endpoint, where `{id}` should be replaced with the ID of the post you want to fetch. Ensure you include your Auth token in the `Authorization` header.

  Replace `<your-Auth-token>` with your actual token and `<post-id>` with the ID of the blog post you want to retrieve.

  ```bash
  curl --location 'http://localhost:8080/auth/getpost/<post-id>' \
  --header 'Authorization: <your-Auth-token>'

  ```

  # Retrieve a Specific Blog Post by ID with Auth Token

  This guide demonstrates how to use `curl` to retrieve a specific blog post by its ID through an API with authentication. In this example, we'll send a GET request to the `/auth/getpost/{id}` endpoint to fetch a blog post by its unique ID. The request includes an access token in the `Authorization` header for authentication.


  ## Retrieve a Specific Blog Post by ID

  To retrieve a specific blog post, send a GET request to the `/auth/getpost/{id}` endpoint, where `{id}` should be replaced with the ID of the post you want to fetch. Ensure you include your Auth token in the `Authorization` header.

  Replace `<your-Auth-token>` with your actual token and `<post-id>` with the ID of the blog post you want to retrieve.

  ```bash
  curl --location 'http://localhost:8080/auth/getpost/<post-id>' \
  --header 'Authorization: <your-Auth-token>'
  ```

  ## Postman Collection

  For an easy way to interact with the MyBlog API, you can import the Postman collection containing all the available API endpoints.

  ### Import the Collection

  You can import the MyBlog Postman Collection in two ways:

  - **Import from Postman Link:** [MyBlog Postman Collection](https://bold-trinity-204583.postman.co/workspace/MyWorkspace~0ef8d7b8-cea0-4d54-8811-c2907e90808d/collection/203659-c045be72-9607-4e6f-adb1-501db6c0dee7?action=share&creator=203659)

  - **Import from File:** You can also import Postman endpoints from the JSON file located in the 'Postman' folder under 'SQL'.

  ### Included Requests

  The MyBlog Postman Collection includes a variety of requests to interact with the MyBlog API:

  - Create, update, retrieve, and delete blog posts.
  - User authentication and user-related endpoints.

  Import this collection into your Postman app to make it easy to work with the MyBlog API.

  If you have any questions or need further assistance, please refer to the API documentation or contact the development team.


  ## To Do

  Map the `BlogPosts` table with the `User` table to store the author's name and user ID from the `User` table.
