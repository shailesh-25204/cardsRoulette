package main

import (
	// "fmt"
	"math/rand"
)

// const cardTypes = [4]string{"skip", "reset", "bomb", "defuse"}
func createNewDeck() *game {
	var cardTypes = [4]string{"skip", "reset", "bomb", "defuse"}

	var d []string

	for i := 0; i < 5; i++ {
		card := cardTypes[rand.Intn(4)]
		d = append(d, card)
	}

	return &game{
		deck:    d,
		result:  false,
		defuses: 0,
	}
}

type game struct {
	deck    []string
	result  bool
	defuses int
}

func (c *Client) gameEngine() {
	// fmt.Println("deck: ", c.game.deck)
	ind := rand.Intn(len(c.game.deck))
	// fmt.Println("rand intn  -> ", ind)
	// fmt.Println("card: ", c.game.deck[ind])
	card := c.game.deck[ind]

	defer func() {
		c.gameCh <- card
	}()

	switch card {
	case "bomb":
		if c.game.defuses <= 0 {
			c.game.deck = nil
			c.game.result = false
			return
		} else {
			c.game.defuses--
		}
	case "reset":
		def := c.game.defuses
		c.game = createNewDeck()
		c.game.defuses = def
		return
	case "defuse":
		c.game.defuses++
	}

	c.game.deck = append(c.game.deck[:ind], c.game.deck[ind+1:]...)
	if len(c.game.deck) == 0 {
		c.game.result = true
		c.hub.redisCL.Client.ZIncrBy("leaderboard", 1, c.username)
	}
}

/*
	gameCh channel strings:
	newgame -> conn.write c.game, c.game.win = false
	reveal -> gameEngine (choose random card and evaluate game state)
	card types skip reset bomb defuse -> write


*/
