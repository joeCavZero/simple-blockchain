package main

import (
	"github.com/joeCavZero/simple-blockchain/api"
	"github.com/joeCavZero/simple-blockchain/blockchain"
)

func main() {
	newBlockchain := blockchain.NewBlockchain()

	api := api.NewApi(newBlockchain)
	api.Run(":8000")

}
