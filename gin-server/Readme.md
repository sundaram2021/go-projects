# Project Documentation: RESTful API for Recipes and Orders using Go and Gin

This project is a Go application that implements a RESTful API for managing recipes and customer orders. It uses the Gin web framework to handle HTTP requests and responses.

---

## **Features**

1. **Recipes Endpoint**:
   - **`GET /recipes`**: Retrieves a list of recipes from a `recipes.json` file. Each recipe includes the item name, recipe details, and price.

2. **Orders Management**:
   - **`POST /orders`**: Creates a new order by accepting a list of item names and calculates the total price based on the recipes.
   - **`PUT /orders`**: Updates an existing order. (Currently, it functions the same as creating an order.)
   - **`DELETE /orders/:item`**: Deletes a specific item from the order and recalculates the total price.
   - **`GET /orders`**: Retrieves the current order details. (Note: This function may need further implementation as it returns a nil order.)

3. **Welcome Message**:
   - **`GET /`**: Returns a welcome message indicating that the Recipes API is active.

---

## **How to Run**

1. **Prerequisites**:
   - Install Go (version 1.13 or higher).
   - Install the Gin web framework:
     ```sh
     go get -u github.com/gin-gonic/gin
     ```

2. **Setup**:
   - Ensure you have a `recipes.json` file in the same directory as `main.go`. This file should contain an array of `MenuItem` objects with `item`, `recipe`, and `price` fields. Example:
     ```json
     [
       {
         "item": "Pizza",
         "recipe": "Dough, Tomato Sauce, Cheese",
         "price": 12.99
       },
       {
         "item": "Burger",
         "recipe": "Bun, Patty, Lettuce, Tomato",
         "price": 9.99
       }
     ]
     ```

3. **Running the Application**:
   - Execute the application using:
     ```sh
     go run main.go
     ```
   - The server will start on port `8080` and display:
     ```
     server is running on port 8080...
     ```

4. **Testing the API**:
   - Use `curl`, Postman, or any HTTP client to interact with the API endpoints.
   - **Example Requests**:
     - **Get Recipes**:
       ```sh
       curl http://localhost:8080/recipes
       ```
     - **Create Order**:
       ```sh
       curl -X POST http://localhost:8080/orders -H "Content-Type: application/json" -d '{"orders":["Pizza","Burger"]}'
       ```
     - **Delete Item from Order**:
       ```sh
       curl -X DELETE http://localhost:8080/orders/Pizza
       ```

---

## **Technologies Used**

- **Go Programming Language**: Core language used for application development.
- **Gin Web Framework**: Used for handling HTTP requests and routing.
- **JSON**: Data format for input/output, using Go's `encoding/json` package.
- **File I/O**: Reading from `recipes.json` using Go's `io/ioutil` and `os` packages.

---