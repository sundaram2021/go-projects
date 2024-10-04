# Project Documentation: JSON to Excel Conversion Tool

This repository contains a Go application that converts JSON data to an Excel file. It provides a convenient way to transform structured JSON data into a format that is easily readable and manageable in spreadsheet software.

---

## Features

1. **JSON Creation**:
   - Generates a JSON file from a predefined data structure.
   - Writes the JSON data to `data.json` in the current directory.

2. **Excel Conversion**:
   - Reads the JSON data from `data.json`.
   - Converts the JSON data to an Excel file named `data.xlsx`, creating a well-structured spreadsheet.

3. **Data Compatibility**:
   - Handles both array of objects and single object JSON formats.
   - Automatically populates the Excel sheet with appropriate headers and rows based on the JSON structure.

---

## How to Use

### Prerequisites

- Go installed on your machine (version 1.16 or higher recommended).
- Access to command-line interface/terminal.

### Setup and Operation

1. **Clone the Repository**:
   ```sh
   git clone https://github.com/yourusername/json-to-excel.git
   cd json-to-excel
   ```

2. **Running the Application**:
   - Execute the application with the following command:
     ```sh
     go run main.go
     ```
   - This command performs the following actions:
     - Creates a `data.json` file with example JSON data.
     - Converts the JSON data into an Excel file (`data.xlsx`).

### Understanding the Code

- The `CreateJson` function:
  - Takes interface{} data and marshals it into JSON.
  - Writes this JSON data to a file named `data.json`.
  
- The `JsonToExcel` function:
  - Opens the `data.json` file and reads its content.
  - Unmarshals the JSON data to dynamically determine its structure.
  - Based on the JSON structure (array of objects or a single object), it populates an Excel file.
  - Saves the populated data into `data.xlsx`.

---

## Technologies Used

- **Go Programming Language**: Core language used for application development.
- **GitHub**: Used for version control and hosting the project repository.
- **Tealeg/xlsx**: A Go package for reading and writing XLSX files.

---

## Dependencies

- **xlsx Package**:
  - To install the `xlsx` package, run:
    ```sh
    go get github.com/tealeg/xlsx
    ```

---

## Example JSON Data Format

- Array of Objects:
  ```json
  [
    {"name": "Sundarm", "age": 23, "city": "City A"},
    {"name": "Aman", "age": 26, "city": "City B"},
    {"name": "Dheeraj", "age": 25, "city": "City C"},
    {"name": "Avnish", "age": 21, "city": "City D"}
  ]
  ```

---
