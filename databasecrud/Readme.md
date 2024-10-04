# Project Documentation: Go SQLite CRUD Application with Context Support

This project is a Go application that demonstrates how to perform basic CRUD (Create, Read, Update, Delete) operations on a SQLite database using the `database/sql` package and the `modernc.org/sqlite` driver. It showcases the use of `context.Context` to manage database query timeouts and cancellations.

---

## **Features**

1. **Database Connection**: Establishes a connection to a SQLite database file (`names.db`) using the `modernc.org/sqlite` driver.
2. Database ops [here](./databaseops.md)

3. **Listing SQL Drivers**: Outputs all registered SQL drivers available in the Go environment.

4. **CRUD Operations**:
   - **Create**: Inserts new records into the `names` table.
   - **Read**: Queries and retrieves all records from the `names` table.
   - **Update**: Modifies existing records in the `names` table.
   - **Delete**: Removes records from the `names` table.

5. **Context Usage**: Utilizes `context.Context` with a timeout for database operations, allowing for better control over query execution times.

6. **Error Handling**: Implements basic error handling and logging to track and display errors during database interactions.

---

## **Getting Started**

- **Prerequisites**:
  - Go installed (version 1.13 or higher).
  - SQLite installed (optional, for managing the database directly).
  - Set up your `GOPATH` and environment variables.

- **Dependencies**:
  - Install the SQLite driver for Go:
    ```sh
    go get modernc.org/sqlite
    ```

- **Database Setup**:
  - Create a SQLite database file named `names.db`.
  - Create a `names` table with the following SQL command:
    ```sql
    CREATE TABLE names (id INTEGER PRIMARY KEY, name TEXT);
    ```

- **Running the Application**:
  - Execute the program using:
    ```sh
    go run main.go
    ```
  - The application will perform the following actions:
    - List available SQL drivers.
    - Open the database connection.
    - Insert a new user (`Id: 6, Name: "John"`).
    - Update the user's name to `"Johnny"`.
    - Delete the user.
    - Query and print all users in the database.

---

## **Project Structure**

- **main.go**: The main Go file containing all the code for database operations, including functions for listing drivers, opening the database, and performing CRUD operations with context support.

---

## **Technologies Used**

- **Go**: Programming language used to build the application.

- **SQLite**: Lightweight, file-based SQL database for data storage.

- **modernc.org/sqlite**: Pure Go SQLite driver for the `database/sql` package.

---

