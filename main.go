package main

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
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

func CreateWallet() (common.Address, *ecdsa.PrivateKey) {
	generatedPrivateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}

	//privateKeyBytes := crypto.FromECDSA(generatedPrivateKey)
	//privateKey := hexutil.Encode(privateKeyBytes)[2:]
	//fmt.Println(privateKey)

	publicKey := generatedPrivateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	address := crypto.PubkeyToAddress(*publicKeyECDSA)
	return address, generatedPrivateKey
}

func ImportWallet(privateKey string) (common.Address, *ecdsa.PrivateKey) {
	importedPrivateKey, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		log.Fatal(err)
	}

	publicKey := importedPrivateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	address := crypto.PubkeyToAddress(*publicKeyECDSA)
	return address, importedPrivateKey
}

func main() {
	avalancheFuji := "https://api.avax-test.network/ext/bc/C/rpc"
	InitNetwork(avalancheFuji)

	newPublicKey, newPrivateKey := CreateWallet()
	fmt.Println("New Public Key: ", newPublicKey)
	fmt.Println("New Private Key: ", newPrivateKey)

	importedPublicKey, importedPrivateKey := ImportWallet("4b82b08ab9a64de944fef3d4184e36771dcd31ad8d8c584ce4b353765e692195")
	fmt.Println("Imported Public Key: ", importedPublicKey)
	fmt.Println("Imported Private Key: ", importedPrivateKey)
}
