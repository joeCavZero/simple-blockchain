package main

import (
	"project/api"
	"project/blockchain"
)

func main() {
	newBlockchain := blockchain.NewBlockchain()

	api := api.NewApi(newBlockchain)
	api.Run(":8000")

}
