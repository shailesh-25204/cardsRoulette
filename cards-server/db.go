package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/go-redis/redis"
)

type User struct {
	Username string `json:"username" binding:"required"`
	Points   int    `json:"points" binding:"required"`
	Rank     int    `json:"rank"`
}

type Database struct {
	Client *redis.Client
}

type Leaderboard struct {
	Type    string  `json:"type"`
	Count   int     `json:"count"`
	Players []*User `json:"players"`
}

var (
	ErrNil         = errors.New("no matching records found")
	LeaderboardKey = "leaderboard"
)

func (db *Database) dbAddUser(username string) error {
	fmt.Println("string ", username)
	exists, err := db.Client.Exists(username).Result()
	if err != nil {
		fmt.Println("Error checking if key exists:", err)
		return err
	}
	fmt.Println("exists = ", exists)
	if exists == 0 {
		member := redis.Z{
			Score:  float64(0),
			Member: username,
		}
		rank := db.Client.ZAdd(LeaderboardKey, member)
		// pipe := db.Client.TxPipeline()
		// pipe.ZAdd(LeaderboardKey, member)
		// rank := pipe.ZRank(LeaderboardKey, username)

		// _, err := pipe.Exec()
		// if err != nil {
		// 	return err
		// }
		fmt.Println(rank.Val(), err)
	}
	return nil
}

func NewDatabase() (*Database, error) {
	RedisAddr := os.Getenv("REDIS_ENDPOINT")
	RedisPort := os.Getenv("REDIS_PORT")
	RedisPass := os.Getenv("REDIS_PASS")

	fmt.Println("Rcloud ADDRESS = = = => ", fmt.Sprintf("%s:%s", RedisAddr, RedisPort))
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", RedisAddr, RedisPort),
		Password: RedisPass,
		DB:       0,
	})

	//populate database
	pipe := client.TxPipeline()

	for i := 1; i <= 10; i++ {
		member := redis.Z{
			Score:  float64(i),
			Member: fmt.Sprintf("user %d", i),
		}
		pipe.ZAdd(LeaderboardKey, member)
	}
	_, err := pipe.Exec()
	if err != nil {
		return nil, err
	}

	fmt.Println("DATABASE POPULATED ... ... ... ... ... ...")

	if err := client.Ping().Err(); err != nil {
		return nil, err
	}
	return &Database{
		Client: client,
	}, nil
}
