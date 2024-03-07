package main

import (
	"encoding/json"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type Hub struct {
	redisCL *Database
	clients map[uuid.UUID]*websocket.Conn
}

func newHub(database *Database) *Hub {
	return &Hub{
		redisCL: database,
		clients: map[uuid.UUID]*websocket.Conn{},
	}
}

func (hub *Hub) GetLeaderboard() (*Leaderboard, error) {
	scores := hub.redisCL.Client.ZRangeWithScores(LeaderboardKey, -10, -1)
	if scores == nil {
		return nil, ErrNil
	}
	count := len(scores.Val())
	users := make([]*User, count)
	for idx, member := range scores.Val() {
		users[idx] = &User{
			Username: member.Member.(string),
			Points:   int(member.Score),
			Rank:     5 - idx,
		}
	}
	leaderboard := &Leaderboard{
		Type:    "leaderboard",
		Count:   count,
		Players: users,
	}
	return leaderboard, nil
}

func (h *Hub) run() {
	ticker := time.NewTicker(3 * time.Second)
	defer func() {
		ticker.Stop()
	}()
	for {
		select {
		case <-ticker.C:
			// fmt.Println("HUB SIZE = ", len(h.clients))
			leaderboard, err := h.GetLeaderboard()
			if err != nil {
				log.Printf("Could Not Fetch Leaderboard")
			}
			jsonData, err := json.Marshal(&leaderboard)
			if err != nil {
				log.Printf("Could Not Fetch Leaderboard")
			}
			for _, conn := range h.clients {
				// conn.write
				conn.WriteMessage(websocket.TextMessage, jsonData)

			}
		}
	}
}
