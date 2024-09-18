package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/gorilla/websocket"
)

var (
	ctx      = context.Background()
	mu       sync.Mutex
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	rooms     = make(map[string]map[*websocket.Conn]string) 
	broadcast = make(chan Message)                          
)

type Message struct {
	Username string `json:"username"`
	Content  string `json:"content"`
	RoomID   string `json:"roomId"`
	System   bool   `json:"system,omitempty"`
}

func main() {
	// Serve static files
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			http.ServeFile(w, r, "./public/join.html")
			return
		}

		// Serve static files from the 'public' directory
		filePath := "./public" + r.URL.Path
		if _, err := os.Stat(filePath); err == nil {
			http.ServeFile(w, r, filePath)
		} else {
			http.NotFound(w, r)
		}
	})

	http.HandleFunc("/api/rooms", getActiveRooms)
	http.HandleFunc("/ws", handleConnections)

	go handleMessages()

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func getActiveRooms(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	mu.Lock()
	var activeRooms []string
	for roomID, clients := range rooms {
		if len(clients) > 0 {
			activeRooms = append(activeRooms, roomID)
		}
	}
	mu.Unlock()

	json.NewEncoder(w).Encode(map[string][]string{
		"rooms": activeRooms,
	})
}

func handleMessages() {
	for {
		msg := <-broadcast
		roomID := msg.RoomID

		mu.Lock()
		clientsInRoom := rooms[roomID]
		mu.Unlock()

		for client := range clientsInRoom {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("Error sending message to client: %v", err)
				client.Close()
				mu.Lock()
				delete(clientsInRoom, client)
				mu.Unlock()
			}
		}
	}
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Error upgrading to websocket: %v", err)
		return
	}
	defer ws.Close()

	roomID := r.URL.Query().Get("roomId")
	if roomID == "" {
		roomID = "defaultRoom"
	}

	username := r.URL.Query().Get("username")
	if username == "" {
		username = "Anonymous"
	}

	mu.Lock()
	if rooms[roomID] == nil {
		rooms[roomID] = make(map[*websocket.Conn]string)
	}
	rooms[roomID][ws] = username
	mu.Unlock()

	log.Printf("Client '%s' connected to room '%s'", username, roomID)

	systemMessage := Message{
		Username: "System",
		Content:  fmt.Sprintf("%s has joined the room.", username),
		RoomID:   roomID,
		System:   true,
	}
	broadcast <- systemMessage

	defer func() {
		mu.Lock()
		delete(rooms[roomID], ws)
		mu.Unlock()

		systemMessage := Message{
			Username: "System",
			Content:  fmt.Sprintf("%s has left the room.", username),
			RoomID:   roomID,
			System:   true,
		}
		broadcast <- systemMessage
	}()

	for {
		var msg Message
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("Error reading message: %v", err)
			break
		}
		msg.RoomID = roomID
		broadcast <- msg
	}
}
