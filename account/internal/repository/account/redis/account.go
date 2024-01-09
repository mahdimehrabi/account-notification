package redis

import (
	"encoding/json"
	"github.com/go-redis/redis"
	"github.com/mahdimehrabi/account-notification/account/internal/entity"
	"github.com/mahdimehrabi/account-notification/account/internal/repository/account"
	"github.com/rs/zerolog/log"
	"strconv"
)

type Redis struct {
	c *redis.Client
}

// I know I must implemented store accounts seperately but bare with me
func (r Redis) List() ([]entity.Account, error) {
	// Get all keys from Redis
	keys, err := r.c.Keys("*").Result()
	if err != nil {
		return nil, err
	}

	// Iterate through keys and retrieve corresponding accounts
	var accounts []entity.Account
	for _, key := range keys {
		str, err := r.c.Get(key).Result()
		if err != nil {
			return nil, err
		}

		var account entity.Account
		if err := json.Unmarshal([]byte(str), &account); err != nil {
			return nil, err
		}

		accounts = append(accounts, account)
	}

	return accounts, nil
}

func (r Redis) Create(account *entity.Account) error {
	accJSON, err := json.Marshal(account)
	if err != nil {
		return err
	}
	//I could use HSEt but we keep it simple
	if err := r.c.Set(strconv.Itoa(int(account.AccountID)), accJSON, 0).Err(); err != nil {
		return err
	}
	return nil
}

func (r Redis) Send(transaction *entity.Transaction) error {
	from, to, amount := transaction.From, transaction.To, transaction.Amount

	if err := r.get(from); err != nil {
		return err
	}
	if err := r.get(to); err != nil {
		return err
	}

	if from.Balance < amount {
		return account.ErrInsufficientAmount
	}
	pipeLine := r.c.TxPipeline()

	from.Balance -= amount
	fromBt, err := json.Marshal(from)
	if err != nil {
		return err
	}
	to.Balance += amount
	toBt, err := json.Marshal(to)
	if err != nil {
		return err
	}
	pipeLine.Set(strconv.FormatInt(from.AccountID, 10), fromBt, 0)
	pipeLine.Set(strconv.FormatInt(to.AccountID, 10), toBt, 0)

	if _, err := pipeLine.Exec(); err != nil {
		return err
	}

	return nil
}

func (r Redis) get(ent *entity.Account) error {
	str, err := r.c.Get(strconv.FormatInt(ent.AccountID, 10)).Result()
	if err != nil {
		return account.ErrNotFound
	}

	if err := json.Unmarshal([]byte(str), ent); err != nil {
		log.Error().Msg(err.Error())
		return err
	}
	return nil
}

func (r Redis) SendContemplate(transaction *entity.Transaction) error {
	from, to, amount := transaction.From, transaction.To, transaction.Amount

	if err := r.get(from); err != nil {
		return err
	}
	if err := r.get(to); err != nil {
		return err
	}

	if from.Balance < amount {
		return account.ErrInsufficientAmount
	}
	pipeLine := r.c.TxPipeline()

	from.Balance += amount
	fromBt, err := json.Marshal(from)
	if err != nil {
		return err
	}
	to.Balance -= amount
	toBt, err := json.Marshal(to)
	if err != nil {
		return err
	}
	pipeLine.Set(strconv.FormatInt(from.AccountID, 10), fromBt, 0)
	pipeLine.Set(strconv.FormatInt(to.AccountID, 10), toBt, 0)

	if _, err := pipeLine.Exec(); err != nil {
		return err
	}

	return nil
}

func NewRedis(redisAddr, redisPassword string, redisDB int) *Redis {
	c := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		DB:       redisDB,
		Password: redisPassword,
	})

	return &Redis{
		c: c,
	}
}
