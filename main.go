package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math"
	"math/big"
)

var client *ethclient.Client

func InitNetwork(networkRPC string) {
	var err error
	client, err = ethclient.Dial(networkRPC)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected")
}

func WalletBalance(address string) (*big.Int, *big.Float) {
	account := common.HexToAddress(address)
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}

	floatBalance := new(big.Float)
	floatBalance.SetString(balance.String())
	ethValue := new(big.Float).Quo(floatBalance, big.NewFloat(math.Pow10(18)))

	return balance, ethValue
}

func WalletBalanceByBlock(address string, block int64) (*big.Int, *big.Float) {
	account := common.HexToAddress(address)
	blockNumber := big.NewInt(block)
	balance, err := client.BalanceAt(context.Background(), account, blockNumber)
	if err != nil {
		log.Fatal(err)
	}

	floatBalance := new(big.Float)
	floatBalance.SetString(balance.String())
	ethValue := new(big.Float).Quo(floatBalance, big.NewFloat(math.Pow10(18)))

	return balance, ethValue
}

func main() {
	avalancheFuji := "https://api.avax-test.network/ext/bc/C/rpc"
	InitNetwork(avalancheFuji)

	walletAddress := "0xCd34A18e1553Ff494E8C342453bea8f4a0feBE8d"
	fmt.Println(WalletBalance(walletAddress))

	fmt.Println(WalletBalanceByBlock(walletAddress, 10956952))

}
