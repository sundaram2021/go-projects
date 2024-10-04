# Project Documentation: Basic Arithmetic REST API

This repository hosts a simple REST API implemented in Go that provides basic arithmetic operations such as addition, subtraction, multiplication, and division. It also includes a feature to extract a name from the URL path and return it in a JSON response. The API uses the built-in `net/http` package for handling HTTP requests and responses.

---

## Features

1. **Arithmetic Operations**:
   - **Add**: Adds two numbers provided as query parameters.
   - **Subtract**: Subtracts the second number from the first, provided as query parameters.
   - **Multiply**: Multiplies two numbers provided as query parameters.
   - **Divide**: Divides the first number by the second, provided as query parameters, including error handling for division by zero.

2. **Name Extraction**:
   - Extracts a name from the URL path and returns it in a JSON response.

3. **Error Handling**:
   - Provides clear error messages for invalid input parameters and division by zero scenarios.

---

## How to Use

### Prerequisites

- Go installed on your machine (version 1.13 or higher is recommended).
- Basic knowledge of REST APIs and how to interact with them, either via tools like curl, Postman, or programmatically.

### Running the API

1. **Clone the Repository**:
   ```bash
   git clone reppo_url
   cd std-server
   ```

2. **Build and Run the Application**:
   ```bash
   go run main.go
   ```
   This command starts the server on port 8080.

### Testing the API

- Use curl, Postman, or any other HTTP client to make requests to the API:

  - **Addition**:
    ```bash
    curl "http://localhost:8080/add?x=10&y=5"
    ```

  - **Subtraction**:
    ```bash
    curl "http://localhost:8080/subtract?x=10&y=5"
    ```

  - **Multiplication**:
    ```bash
    curl "http://localhost:8080/multiply?x=10&y=5"
    ```

  - **Division**:
    ```bash
    curl "http://localhost:8080/divide?x=10&y=5"
    ```

  - **Name Extraction**:
    ```bash
    curl "http://localhost:8080/John"
    ```

---

## Technologies Used

- **Go (Golang)**: Core programming language used to develop the REST API.
- **Net/HTTP Package**: Used to set up the HTTP server and handle requests and responses.

---

## API Endpoints

- **/add**: Takes two query parameters `x` and `y`, adds them, and returns the result.
- **/subtract**: Takes two query parameters `x` and `y`, subtracts `y` from `x`, and returns the result.
- **/multiply**: Takes two query parameters `x` and `y`, multiplies them, and returns the result.
- **/divide**: Takes two query parameters `x` and `y`, divides `x` by `y`, and handles division by zero.
- **/:name**: Takes a name as part of the URL path and returns it in a JSON response.

---

This documentation provides a comprehensive guide on how to set up, run, and interact with the Basic Arithmetic REST API, detailing the endpoints available and how to use them effectively. Whether for educational purposes, API testing, or as a base for more complex calculator applications, this project serves as a practical example of REST API development in Go.