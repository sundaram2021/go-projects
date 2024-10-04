# Project Documentation: Gorilla Mux REST API

This Go application utilizes the Gorilla Mux router to implement a simple REST API for user management, including features for user registration and login. It demonstrates the creation, reading, and management of user data, leveraging the robust routing capabilities of Gorilla Mux.

---

## Features

1. **Root Endpoint**:
   - **GET /**: Displays a welcome message along with the requested URI.

2. **Numeric Request Handler**:
   - **GET /{num}**: Responds with the number specified in the URL.

3. **Query Parameter Handler**:
   - **GET /query**: Extracts and displays a query parameter named `query`.

4. **User Registration**:
   - **POST /register**: Allows new users to register by providing their ID, name, and password. Registers the user and adds them to the registered users list.

5. **User Login**:
   - **POST /login**: Handles user login by verifying provided credentials against registered users. Logs the user in and adds them to the logged-in users list.

---

## How to Use

### Prerequisites

- **Go Environment**: Ensure Go is installed on your machine (version 1.13 or higher is recommended).
- **Gorilla Mux Package**: Required for routing, which can be installed using:
  ```bash
  go get -u github.com/gorilla/mux
  ```

### Setup

1. **Clone the Repository**:
   Assuming this code is hosted on a GitHub repository, you can clone it using:
   ```bash
   git clone repo_url
   ```

2. **Navigate to the Project Directory**:
   ```bash
   cd server
   ```

3. **Build and Run the Application**:
   Compile and run the application using:
   ```bash
   go run main.go
   ```

### Testing the API

- Use curl or Postman to interact with the API:
  - **Root Endpoint**:
    ```bash
    curl http://localhost:8080/
    ```
  - **Numeric Endpoint**:
    ```bash
    curl http://localhost:8080/123
    ```
  - **Query Parameter**:
    ```bash
    curl "http://localhost:8080/query?query=test"
    ```
  - **Register User**:
    ```bash
    curl -X POST http://localhost:8080/register -H "Content-Type: application/json" -d '{"id":1, "name":"John Doe", "password":"123456"}'
    ```
  - **Login User**:
    ```bash
    curl -X POST http://localhost:8080/login -H "Content-Type: application/json" -d '{"name":"John Doe", "password":"123456"}'
    ```

---

## Technologies Used

- **Go**: Core programming language.
- **Gorilla Mux**: HTTP routing and URL matcher for building Go web servers.

---

## Dependencies

- **Gorilla Mux**:
  Responsible for handling HTTP requests and routing in a more flexible way than the standard net/http package in Go.

---

This documentation provides a comprehensive guide to setting up and using the REST API developed with Gorilla Mux, including detailed instructions for interacting with each endpoint. Whether for educational purposes or practical applications, this project offers a solid foundation for further development and customization.