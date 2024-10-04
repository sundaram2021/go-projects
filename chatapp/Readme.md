# Project Documentation: WebSocket Chat Application with Multiple Rooms

This project is a real-time chat application built using Go, WebSockets, and HTML/CSS. It allows users to join chat rooms via unique room IDs and communicate with others in the same room. Below are the key components and features of the application:

---

## **Features**

1. **Multiple Chat Rooms**: Users can create or join chat rooms using unique room IDs. Each room operates independently, ensuring messages are only shared among users in the same room.

2. **User Interface Improvements**: The `join.html` and `interface.html` pages have been enhanced with modern CSS to provide a more user-friendly experience.

3. **User Notifications**: When a new user joins or leaves a room, the application sends a system message to all users in that room to notify them of the change.

4. **Active Rooms Display**: The home page lists all active rooms, showing users which rooms are currently in use. However, users must enter the room ID manually to join; direct links are not provided to prevent unauthorized access.

5. **WebSocket Communication**: Real-time messaging is implemented using WebSockets, allowing for instantaneous communication between clients and the server.

6. **Server-Side Room Management**: The server maintains a mapping of rooms and connected clients, ensuring that messages are routed correctly and that resources are managed efficiently.

7. **Client-Side Storage**: Usernames and room IDs are stored in `localStorage` to persist user data across sessions.


---

## **Getting Started**

- **Running the Server**: Use `go run main.go` to start the server. Ensure that all HTML, CSS, and JavaScript files are correctly placed in the `public` directory.

- **Accessing the Application**: Navigate to `http://localhost:8080/` in your web browser to access the chat application's home page.

- **Joining a Room**: Enter your username and a room ID on the home page to join or create a chat room.

---

## **Project Structure**

- **main.go**: The main server file containing the Go code for handling WebSocket connections, message broadcasting, and room management.

- **public/**: Directory containing all static files served to the client, including HTML pages and CSS styles.

  - **join.html**: The home page where users enter their username and room ID.

  - **interface.html**: The chat interface displayed after joining a room.

  - **static/styles.css**: CSS file containing styles for the application's UI.

---

## **Technologies Used**

- **Go**: Backend language used for server-side logic and WebSocket handling.

- **HTML/CSS**: Frontend technologies for structuring and styling the user interface.

- **JavaScript**: Client-side scripting for WebSocket communication and dynamic content updates.

---

