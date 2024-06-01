package main

import (
	"encoding/json"
	"fmt"
	"log"
	"sync"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

type Message struct {
	Sender  string `json:"sender"`
	Content string `json:"content"`
}

func main() {
	app := fiber.New()

	app.Use(func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) { // Returns true if the client requested upgrade to the WebSocket protocol
			return c.Next()
		}
		return c.SendStatus(fiber.StatusUpgradeRequired)
	})

	app.Get("/ws/:id", websocket.New(func(c *websocket.Conn) {
		defer func() {
			unregister <- c
			c.Close()
		}()

		id := c.Params("id")

		register <- c

		for {
			messageType, message, err := c.ReadMessage()
			if err != nil {
				if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
					fmt.Println("read error:", err)
				}
				return // Calls the deferred function, i.e. closes the connection on error
			}

			if messageType == websocket.TextMessage {
				var receivedMsg Message
				if err := json.Unmarshal(message, &receivedMsg); err != nil {
					fmt.Println("unmarshal error:", err)
					continue
				}
				receivedMsg.Sender = id // Ensure the sender is set to the correct ID

				msgBytes, _ := json.Marshal(receivedMsg)
				broadcast <- string(msgBytes)
			} else {
				fmt.Println("websocket message received of type", messageType)
			}
		}
	}))

	go runHub()

	log.Fatalln(app.Listen(":8000"))
}

type client struct {
	isClosing bool
	mu        sync.Mutex
}

var clients = make(map[*websocket.Conn]*client)
var register = make(chan *websocket.Conn)
var broadcast = make(chan string)
var unregister = make(chan *websocket.Conn)

func runHub() {
	for {
		select {
		case connection := <-register:
			clients[connection] = &client{}
			fmt.Println("connection registered")

		case message := <-broadcast:
			fmt.Println("message received:", message)
			for connection, c := range clients {
				go func(connection *websocket.Conn, c *client) {
					c.mu.Lock()
					defer c.mu.Unlock()
					if c.isClosing {
						return
					}
					if err := connection.WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
						c.isClosing = true
						fmt.Println("write error:", err)
						connection.WriteMessage(websocket.CloseMessage, []byte{})
						connection.Close()
						unregister <- connection
					}
				}(connection, c)
			}

		case connection := <-unregister:
			delete(clients, connection)
			fmt.Println("connection unregistered")
		}
	}
}
