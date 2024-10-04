# Project Documentation: RESTful API with JWT Authentication using Go and Gin

This project is a Go application that implements a RESTful API for managing recipes, featuring user authentication with JSON Web Tokens (JWT). It utilizes the Gin web framework for handling HTTP requests and middleware, and includes session management with secure cookies.

---

## Features

1. **User Authentication**:
   - **Sign-In**: Users can sign in to obtain a JWT for authenticated requests (`POST /signin`).
   - **Token Refresh**: Users can refresh their JWT before it expires (`POST /refresh`).

2. **Recipes Management** (Protected Endpoints):
   - **Create Recipe**: Add a new recipe to the collection (`POST /recipes`).
   - **Update Recipe**: Modify an existing recipe by ID (`PUT /recipes/:id`).
   - **Delete Recipe**: Remove a recipe by ID (`DELETE /recipes/:id`).
   - **Search Recipes**: Find recipes by tag (`GET /recipes/search`).

3. **Public Endpoints**:
   - **List All Recipes**: Retrieve all available recipes (`GET /recipes`).

4. **Session Management**:
   - Uses secure cookie-based sessions to store user session data.

5. **JWT Middleware**:
   - Protects routes by validating JWTs provided in the `Authorization` header.

6. **Data Persistence**:
   - Recipes are loaded from and saved to a `recipes.json` file.

---

## How to Run

### Prerequisites

- **Go**: Install Go (version 1.13 or higher).
- **Dependencies**: Install required Go packages.

### Installation

1. **Clone the Repository**:

   ```sh
   git clone repo_url
   cd project_directory
   ```

2. **Install Dependencies**:

   Ensure you have the following Go packages installed:

   ```sh
   go get github.com/gin-gonic/gin
   go get github.com/gin-contrib/sessions
   go get github.com/gin-contrib/sessions/cookie
   go get github.com/dgrijalva/jwt-go
   go get github.com/joho/godotenv
   go get github.com/rs/xid
   ```

3. **Environment Variables**:

   - Create a `.env` file in the project root directory.
   - Add the following environment variable:

     ```env
     JWT_SECRET=your_jwt_secret_key
     ```

4. **Prepare the `recipes.json` File**:

   - Ensure there is a `recipes.json` file in the same directory as `main.go`.
   - The file should contain an array of recipes. Example content:

     ```json
     [
       {
         "id": "1",
         "name": "Spaghetti Bolognese",
         "tags": ["Italian", "Pasta"],
         "ingredients": ["Spaghetti", "Tomato Sauce", "Ground Beef"],
         "instructions": ["Boil pasta", "Cook sauce", "Mix together"],
         "publishedAt": "2021-01-01T00:00:00Z"
       }
     ]
     ```

### Running the Application

1. **Start the Server**:

   ```sh
   go run main.go
   ```

   The server will start and listen on the default port (usually `:8080`).

2. **Testing the API**:

   - Use an API client like **Postman** or **cURL** to interact with the endpoints.

### API Endpoints

1. **Public Endpoints**:

   - **List All Recipes**:

     ```http
     GET /recipes
     ```

2. **Authentication**:

   - **Sign In**:

     ```http
     POST /signin
     ```

     - **Request Body**:

       ```json
       {
         "username": "your_username",
         "password": "your_password"
       }
       ```

     - **Response**: Returns a JWT token and expiration time.

   - **Refresh Token**:

     ```http
     POST /refresh
     ```

     - **Headers**:

       ```
       Authorization: your_jwt_token
       ```

3. **Protected Endpoints** (Require JWT):

   - **Create Recipe**:

     ```http
     POST /recipes
     ```

     - **Headers**:

       ```
       Authorization: your_jwt_token
       ```

     - **Request Body**:

       ```json
       {
         "name": "New Recipe",
         "tags": ["Tag1", "Tag2"],
         "ingredients": ["Ingredient1", "Ingredient2"],
         "instructions": ["Step 1", "Step 2"]
       }
       ```

   - **Update Recipe**:

     ```http
     PUT /recipes/:id
     ```

     - Replace `:id` with the recipe ID.
     - **Headers**:

       ```
       Authorization: your_jwt_token
       ```

     - **Request Body**: Same as the create recipe body.

   - **Delete Recipe**:

     ```http
     DELETE /recipes/:id
     ```

     - Replace `:id` with the recipe ID.
     - **Headers**:

       ```
       Authorization: your_jwt_token
       ```

   - **Search Recipes**:

     ```http
     GET /recipes/search?tag=desired_tag
     ```

     - **Headers**:

       ```
       Authorization: your_jwt_token
       ```

---

## Technologies Used

- **Go (Golang)**: Programming language used for server-side development.

- **Gin Web Framework**: High-performance HTTP web framework for Go.

- **JWT (JSON Web Tokens)**: Used for secure user authentication.

- **Gin Sessions**: Session management middleware for Gin.

- **Environment Variables**: Managed using `github.com/joho/godotenv`.

- **Unique ID Generation**: Using `github.com/rs/xid` for generating unique recipe IDs.

- **JSON Handling**: Standard `encoding/json` package for marshaling and unmarshaling data.

- **File I/O**: Using `io/ioutil` and `os` packages for reading from and writing to files.

---

This documentation provides an overview of the application's key features, instructions on how to set it up and run, and the technologies involved in its development.