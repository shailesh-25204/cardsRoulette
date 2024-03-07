package main

import (
	"bytes"
	"encoding/json"

	// "fmt"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

const (
	writeWait = 60 * time.Second

	pongWait = 60 * time.Second

	pingPeriod = (pongWait * 9) / 10

	maxMessageSize = 512
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

var allowedOrigins = map[string]bool{
	"http://localhost:5173": true,
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		origin := r.Header.Get("Origin")
		return allowedOrigins[origin]
	},
}

type Client struct {
	id       uuid.UUID
	hub      *Hub
	conn     *websocket.Conn
	game     *game
	gameCh   chan string
	rwCh     chan string
	username string
}

type Request struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

type Response struct {
	Type     string `json:"type"`
	Len      int    `json:"len"`
	NextCard string `json:"nextCard"`
	Result   bool   `json:"result"`
}

func (c *Client) readPump() {
	defer func() {
		delete(c.hub.clients, c.id)
		c.conn.Close()
	}()
	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
		// fmt.Println("will unmarshal now", string(message))
		reqStr := string(message)
		if len(reqStr) == 2 && reqStr == "{}" {
			log.Printf("Empty Json Object")
			return
		}
		var req Request
		err2 := json.Unmarshal(message, &req)
		if false {
			log.Printf("error decoding sakura response: %v", err2)
			if e, ok := err2.(*json.SyntaxError); ok {
				log.Printf("syntax error at byte offset %d", e.Offset)
			}
			log.Printf("sakura response: %q", message)
			return
		}
		// fmt.Println("The req is :", req)

		//Handle New User
		if c.username == "" && req.Type == "newUser" {
			c.username = req.Data
			c.rwCh <- req.Type
			//Add user to datebase
			// fmt.Println("dbAddUser Call")
			_ = c.hub.redisCL.dbAddUser(c.username)
		}
		// fmt.Println("c.username -> ", c.username)
		if req.Type == "move" {
			if req.Data == "newGame" {
				c.game.result = false
				c.game.defuses = 0
				c.game = createNewDeck()
				c.rwCh <- "newGame"
			}
			if req.Data == "reveal" && len(c.game.deck) > 0 {
				c.gameEngine()
			}
		}

		// if input == "newgame" {
		// 	c.gameCh <- input
		// 	c.game.result = false
		// 	c.game.defuses = 0
		// } else if input == "reveal" {
		// 	if (len(c.game.deck)) > 0 {
		// 		c.gameEngine()
		// 	}
		// }
		// c.gameCh <- input
	}
}

func (c *Client) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		delete(c.hub.clients, c.id)
		ticker.Stop()
		c.conn.Close()
	}()
	for {
		// fmt.Println("Write Pump Running...")
		select {
		case rwMessage := <-c.rwCh:
			// fmt.Println("recieved this  ", rwMessage)
			var jsonData []byte
			var err error
			switch rwMessage {
			case "newUser":
				res := Request{
					Type: "regSuccess",
					Data: c.username,
				}
				jsonData, err = json.Marshal(&res)
			case "newGame":
				res := Response{
					Type:     "gameState",
					Len:      len(c.game.deck),
					NextCard: "",
					Result:   c.game.result,
				}
				jsonData, err = json.Marshal(res)
			}
			if err != nil {
				return
			}

			c.conn.WriteMessage(websocket.TextMessage, jsonData)

		case message := <-c.gameCh:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			res := Response{
				Type:     "gameState",
				Len:      len(c.game.deck),
				NextCard: message,
				Result:   c.game.result,
			}

			jsonData, err := json.Marshal(res)
			if err != nil {
				return
			}

			c.conn.WriteMessage(websocket.TextMessage, jsonData)
			// // if !ok {
			// // 	c.conn.WriteMessage(websocket.CloseMessage, []byte{})
			// // 	return
			// // }

			// w, err := c.conn.NextWriter(websocket.TextMessage)
			// if err != nil {
			// 	return
			// }
			// w.Write([]byte(message))

			// n := len(c.gameCh)
			// for i := 0; i < n; i++ {
			// 	w.Write(newline)
			// 	w.Write([]byte(<-c.gameCh))
			// }
			// fmt.Println("WritePump deck ", c.game.deck)
			// if len(c.game.deck) == 0 {
			// 	w.Write(newline)
			// 	if c.game.result {
			// 		w.Write([]byte("WIN"))
			// 	} else {
			// 		w.Write([]byte("LOSS"))
			// 	}
			// }
			// if err := w.Close(); err != nil {
			// 	return
			// }
		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

func serveWs(hub *Hub, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	newGame := createNewDeck()
	id := uuid.New()
	client := &Client{id: id, hub: hub, conn: conn, game: newGame, gameCh: make(chan string), username: "", rwCh: make(chan string)}
	client.hub.clients[id] = conn

	// Client is added in database after getting username

	go client.writePump()
	go client.readPump()
	// go client.gameEngine()

}

/*
	------>
	newGame
	reveal
	newUser
	<------
	newUser : username
	card : card
*/
