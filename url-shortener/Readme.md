# Project Documentation: URL Shortener Service

This Go application implements a basic URL shortener service. It provides functionality to shorten URLs via a web interface and redirects users to the original URL when accessing the shortened version. The URLs and their mappings are persisted locally in a JSON file.

---

## Features

1. **URL Shortening**:
   - Accepts a long URL via an HTTP request and provides a shortened version.
   - Uses SHA1 hashing to generate a unique short identifier for each URL.

2. **Redirection**:
   - Redirects users to the original URL when they access the shortened URL.

3. **Persistent Storage**:
   - Stores URL mappings in a local file (`urls.json`) to retain data across server restarts.
   - Loads existing URL mappings from the file at startup.

4. **Concurrency Safe**:
   - Uses mutex locks to ensure that URL mappings are safely accessed and modified across multiple requests.

---

## How to Use

### Prerequisites

- Go installed on your machine (Go 1.15 or higher recommended).
- Basic understanding of HTTP protocols and RESTful principles.

### Running the Application

1. **Clone the Repository**:
   Assuming the code is hosted on a GitHub repository, you can clone it using:
   ```bash
   git clone repo_url
   cd url-shortener
   ```

2. **Start the Server**:
   Run the server using:
   ```bash
   go run main.go
   ```
   The server starts on port 8080.

### Interacting with the API

- **Shorten a URL**:
  Use a browser or a tool like curl to request a shortened URL:
  ```bash
  curl "http://localhost:8080/shorten?url=https://example.com"
  ```
  This will return a shortened URL that you can visit in your browser.

- **Visit a Shortened URL**:
  Access the shortened URL in a browser or with curl, and you will be redirected to the original URL.

---

## API Endpoints

- **POST /shorten**:
  - **Purpose**: Generates a shortened URL.
  - **Query Parameters**:
    - **url**: The original URL to be shortened.
  - **Response**: A shortened URL.

- **GET /{shortURL}**:
  - **Purpose**: Redirects to the original URL associated with the shortened URL.
  - **Path Variables**:
    - **shortURL**: The short identifier of the original URL.
  - **Response**: HTTP redirect to the original URL.

---

## Implementation Details

- **Short URL Generation**:
  - Uses SHA1 to hash the original URL and truncates it to create a unique identifier.
  
- **Concurrency Handling**:
  - Uses `sync.Mutex` to manage concurrent read/write access to the URL store.

- **Data Persistence**:
  - Uses `json.Marshal` and `ioutil.WriteFile` to serialize URL mappings to a file.
  - Uses `json.Unmarshal` and `ioutil.ReadFile` to load URL mappings from a file.

---

## Technologies Used

- **Go (Golang)**: The primary programming language used.
- **Standard Library Packages**:
  - `net/http` for HTTP server functionality.
  - `encoding/json` for JSON serialization.
  - `crypto/sha1` for generating hashes.
  - `sync` for handling concurrency.
  - `ioutil` and `os` for file operations.

---

This documentation outlines how to set up, run, and interact with the URL shortener service, making it accessible for developers interested in understanding or enhancing the project. Whether used for educational purposes or as a template for more complex applications, this project serves as a practical demonstration of basic web server operations in Go.