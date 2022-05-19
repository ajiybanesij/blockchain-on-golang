package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
)

import (
	"log"
)

var client *ethclient.Client

// Settin up the client
func InitNetwork(networkRPC string) {
	var err error
	client, err = ethclient.Dial(networkRPC)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected")
}

func main() {
	avalancheFuji := "https://api.avax-test.network/ext/bc/C/rpc"
	InitNetwork(avalancheFuji)
}
