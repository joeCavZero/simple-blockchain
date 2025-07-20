package main

import (
	"os"

	"github.com/joeCavZero/simple-blockchain/src/api"
	"github.com/joeCavZero/simple-blockchain/src/blockchain"
	"github.com/joeCavZero/simple-blockchain/src/cli"
	"github.com/joeCavZero/simple-blockchain/src/logkit"
)

func main() {
	mainlk := logkit.NewLogkit("Main")

	cli := cli.NewCli()

	switch cli.Mode {
	case "prod":
		newBlockchain := blockchain.NewBlockchain()

		api := api.NewApi(newBlockchain)
		api.Run(cli.Port)
	case "test":
		mainlk.Info("Test mod is not implemented yet")
		os.Exit(0)
	default:
		mainlk.Error("Invalid mode")
		os.Exit(1)
	}

}
