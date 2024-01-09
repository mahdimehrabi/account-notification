package infrastructures

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Env has environment stored
type Env struct {
	ServerPort       string
	RedisAddr        string
	RedisDB          int
	RedisPassword    string
	NotificationAddr string
}

// NewEnv creates a new environment
func NewEnv() *Env {
	env := Env{}
	return &env
}

// LoadEnv loads environment
func (env *Env) LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	env.ServerPort = os.Getenv("ServerPort")
	env.RedisAddr = os.Getenv("RedisAddr")
	rdb := os.Getenv("RedisDB")
	rdbi, err := strconv.Atoi(rdb)
	if err != nil {
		log.Fatal(err)
	}
	env.RedisDB = rdbi
	env.RedisPassword = os.Getenv("RedisPassword")
	env.NotificationAddr = os.Getenv("NotificationAddr")
}
