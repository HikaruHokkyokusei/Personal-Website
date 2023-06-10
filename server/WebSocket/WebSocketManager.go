package WebSocketManager

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
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

type GenericWsMessage struct {
	Event string          `json:"event"`
	Data  json.RawMessage `json:"data"`
}

type WsClientMessage struct {
	cliId string
	msg   GenericWsMessage
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

func sendMessage(uid string, message GenericWsMessage) {
	if payload, err := json.Marshal(message); err == nil {
		if err := connectedClients[uid].conn.WriteMessage(websocket.TextMessage, payload); err != nil {
			panic(err)
		}
	} else {
		panic(err)
	}
}

func handleWsClientMessage(wsClientMessage WsClientMessage) {
	var clientData, ok = connectedClients[wsClientMessage.cliId]
	if ok {
		clientData.mutex.Lock()
		defer func() {
			var clientData, ok = connectedClients[wsClientMessage.cliId]
			if ok {
				clientData.mutex.Unlock()
			}
		}()

		switch wsClientMessage.msg.Event {
		}
	}
}

func ConfigureWebsocket(app fiber.Router) {
	app.Use("/", func(ctx *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(ctx) {
			ctx.Locals("allowed", true)
			return ctx.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	go runClientHandler()

	app.Get("/", websocket.New(func(conn *websocket.Conn) {
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
			temp GenericWsMessage
		)
		for {
			if _, msg, err = conn.ReadMessage(); err == nil {
				temp = GenericWsMessage{}
				if err := json.Unmarshal(msg, &temp); err == nil {
					clientMessages <- WsClientMessage{uid, temp}
				} else {
					fmt.Println("Error when parsing the websocket message")
				}
			} else {
				if !isCloseMessage {
					fmt.Println("Error when reading the message:", err)
				}
				break
			}
		}
	}))
}
