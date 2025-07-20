package cli

import (
	"os"
	"strconv"

	"github.com/joeCavZero/simple-blockchain/src/logkit"
)

type Cli struct {
	Mode string
	Port uint16
}

func NewCli() *Cli {
	clilk := logkit.NewLogkit("CLI")

	mode := os.Getenv("MODE")
	if mode != "prod" && mode != "test" {
		clilk.Error("Mode must be 'prod' or 'test'.")
		os.Exit(1)
	}
	port_str := os.Getenv("PORT")
	if port_str == "" {
		port_str = "8000"
	}
	port_int, err := strconv.Atoi(port_str)
	if err != nil {
		clilk.Error("Port must be a number.")
		os.Exit(1)
	}
	port := uint16(port_int)
	return &Cli{
		Mode: mode,
		Port: port,
	}
}
