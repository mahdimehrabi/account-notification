package main

import (
	"bytes"
	"crypto/rand"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	rand2 "math/rand"
	"net/http"
	"time"
)

type CreateDTO struct {
	AccountID int64   `json:"accountID" validate:"required"`
	Balance   float64 `json:"balance" validate:"required"`
}

type TransactionDTO struct {
	FromID   int64   `json:"fromID" validate:"required"`
	ToID     int64   `json:"toID" validate:"required"`
	Amount   float64 `json:"amount" validate:"required"`
	CreateAt int64   `json:"-"`
}

func generateRandomBytes(n int) ([]byte, error) {
	randomBytes := make([]byte, 50+n)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return nil, err
	}
	return randomBytes, nil
}

var r *rand2.Rand
var accounts = make([]CreateDTO, 0)

func main() {
	s := rand2.NewSource(time.Now().UnixNano())
	r = rand2.New(s)
	accountCount := 1000
	createAccounts(accountCount)
	fmt.Println(accountCount, "accounts created successfully")

	// Seed the random number generator with the current time in nanoseconds
	start := time.Now()

	// Define the URL to send the POST request to
	url := "http://localhost:8000/api/accounts/send/"
	//
	// Get the number of requests to send from command line input
	numRequests := flag.Int("n", 1, "Number of requests to send")
	flag.Parse()

	// Counters for successful and failed requests
	successfulRequests := 0
	failedRequests := 0

	for i := 0; i < *numRequests; i++ {
		// Generate random data for the Req
		tr := TransactionDTO{
			FromID: accounts[r.Intn(len(accounts)-1)].AccountID,
			ToID:   accounts[r.Intn(len(accounts)-1)].AccountID,
			Amount: 1,
		}

		// Convert the Req to a JSON payload
		payload, err := json.Marshal(tr)
		if err != nil {
			fmt.Println("Error marshaling JSON:", err)
			failedRequests++
			continue
		}

		// Send the POST request
		resp, err := http.Post(url, "application/json", bytes.NewBuffer(payload))
		if err != nil {
			fmt.Println("Error sending POST request:", err)
			failedRequests++
			continue
		}

		// Read and parse the response
		if resp.StatusCode == http.StatusOK {
			fmt.Printf("send transaction %d successfully\n", i)
			successfulRequests++
		} else {
			bf := bytes.NewBuffer(make([]byte, 0))
			_, err = io.Copy(bf, resp.Body)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Printf("Received a non-OK response: %d for item:%d ,resp:%s\n", resp.StatusCode, i, bf.String())
			failedRequests++
		}
		resp.Body.Close()

		time.Sleep(420 * time.Nanosecond)
	}
	end := time.Now()
	fmt.Printf("Successful requests: %d\n", successfulRequests)
	fmt.Printf("Failed requests: %d\n", failedRequests)
	fmt.Printf("Benchmark: %s \n", end.Sub(start))
}

func createAccounts(numRequests int) {
	url := "http://localhost:8000/api/accounts/"
	for i := 0; i < numRequests; i++ {
		// Generate random data for the Req
		account := CreateDTO{
			AccountID: int64(r.Intn(99999999999)),
			Balance:   1000.0,
		}

		req, err := json.Marshal(account)
		if err != nil {
			log.Fatal(err)
		}

		// Send the POST request
		resp, err := http.Post(url, "application/json", bytes.NewBuffer(req))
		if err != nil {
			log.Fatal(err)
		}
		resp.Body.Close()

		// Read and parse the response
		if resp.StatusCode == http.StatusOK {
			fmt.Printf("Created account %d successfully\n", i)
		} else {
			fmt.Printf("Received a non-OK response: %d for item:%d\n", resp.StatusCode, i)
		}
		accounts = append(accounts, account)
		time.Sleep(420 * time.Nanosecond)
	}
}
