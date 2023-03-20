package WebSocketManager

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"github.com/google/uuid"
)

type WsClientData struct {
	uid  string
	conn *websocket.Conn
}
type ConnectedWsClient struct {
	mutex sync.Mutex
	conn  *websocket.Conn
}
type WsClientMessage struct {
	uid   string
	event string
	data  any
}

var (
	connectedClients = map[string]*ConnectedWsClient{}
	registerClient   = make(chan WsClientData, 25)
	unregisterClient = make(chan string, 25)
	clientMessages   = make(chan WsClientMessage, 100)
)

func runClientHandler() {
	for {
		select {
		case webSocketClient := <-registerClient:
			connectedClients[webSocketClient.uid] = &ConnectedWsClient{conn: webSocketClient.conn}
			fmt.Println("New Connection:", webSocketClient.uid)

		case clientId := <-unregisterClient:
			delete(connectedClients, clientId)
			fmt.Println("Disconnected:", clientId)

		case clientMessage := <-clientMessages:
			go func() {
				handleWsClientMessage(clientMessage)
			}()
		}
	}
}

func handleWsClientMessage(wsClientMessage WsClientMessage) {
	var clientData, ok = connectedClients[wsClientMessage.uid]
	if ok {
		clientData.mutex.Lock()
		defer func() {
			var clientData, ok = connectedClients[wsClientMessage.uid]
			if ok {
				clientData.mutex.Unlock()
			}
		}()

		switch wsClientMessage.event {
		}
	}
}

func ConfigureWebsocket(server *fiber.App) {
	server.Use("/ws", func(ctx *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(ctx) {
			return ctx.Next()
		}
		return ctx.SendStatus(fiber.StatusUpgradeRequired)
	})

	go runClientHandler()

	server.Get("/ws", websocket.New(func(conn *websocket.Conn) {
		var uid = uuid.New().String() + "-" + uuid.New().String()
		registerClient <- WsClientData{uid, conn}

		defer func(conn *websocket.Conn) {
			unregisterClient <- uid
			var err = conn.Close()
			if err != nil {
				fmt.Println("Error when trying to close the connection for client:", uid)
			}
		}(conn)

		var isCloseMessage = false
		conn.SetCloseHandler(func(code int, text string) error {
			isCloseMessage = true
			return nil
		})

		var (
			msg  []byte
			err  error
			temp map[string]any
		)
		for {
			if _, msg, err = conn.ReadMessage(); err != nil {
				if !isCloseMessage {
					fmt.Println("Error when reading the message:", err)
				}
				break
			}

			temp = make(map[string]any)
			var err = json.Unmarshal(msg, &temp)
			if err != nil {
				fmt.Println("Error when parsing the websocket message")
				continue
			}
			clientMessages <- WsClientMessage{uid, temp["event"].(string), temp["data"]}
		}
	}))
}
