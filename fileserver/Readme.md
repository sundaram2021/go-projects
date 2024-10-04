# Project Documentation: Go Application for JSON File CRUD Operations

This project is a Go application that performs basic CRUD (Create, Read, Update, Delete) operations on a JSON file named `person.json`. It demonstrates how to work with JSON data and file I/O in Go.

---

## **Features**

1. **Data Structure**:
   - Defines a `Person` struct with `Name` (string) and `Age` (uint) fields.

2. **Create Data**:
   - Function `CreateData(a []Person)`:
     - Creates a new JSON file `person.json`.
     - Writes an array of `Person` structs to the file in JSON format.

3. **Read Data**:
   - Function `ReadData()`:
     - Opens and reads the `person.json` file.
     - Outputs the JSON content to the console.

4. **Update Data**:
   - Function `UpdateData(a []Person)`:
     - Opens `person.json` in write mode.
     - Overwrites the existing content with new JSON data from an array of `Person` structs.

5. **Delete Data**:
   - Function `DeletData()`:
     - Deletes the `person.json` file from the filesystem.

6. **Main Execution Flow**:
   - Creates initial data with two `Person` entries and writes to `person.json`.
   - Reads and displays the file content.
   - Updates the data with a new set of `Person` entries.
   - Reads and displays the updated file content.
   - Deletes the `person.json` file.

---

## **How to Run**

1. **Prerequisites**:
   - Install Go (version 1.13 or higher).

2. **Setup**:
   - Save the provided code in a file named `main.go`.

3. **Running the Application**:
   - Open a terminal and navigate to the directory containing `main.go`.
   - Execute the command:
     ```sh
     go run main.go
     ```
   - The application will perform the following steps:
     - Create `person.json` with initial data.
     - Read and display the content of `person.json`.
     - Update `person.json` with new data.
     - Read and display the updated content.
     - Delete `person.json`.

4. **Expected Output**:
   - Messages indicating the success or failure of each operation.
   - The content of `person.json` before and after the update.

---

## **Technologies Used**

- **JSON Encoding/Decoding**: Utilizes Go's `encoding/json` package to marshal and unmarshal data.
- **File I/O Operations**: Uses Go's `os` and `io` packages for file handling.
- **Standard Library**: No external dependencies required.

---