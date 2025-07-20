# SIMPLE BLOCKCHAIN

A simple blockchain implementation in Go, designed for learning purposes.

## Features
- Basic proof-of-work mechanism
- Block and blockchain structures
- Mining simulation
- Easy to understand and extend

## Requirements
- Go 1.24.0 or later

## Getting Started

### Running Locally without Docker

1. Clone the repository:
```bash
git clone https://github.com/joeCavZero/simple-blockchain.git
cd simple-blockchain
```

2. Install the dependencies:
```bash
go mod tidy
```

2. Run:
```bash
go run main.go
```

### Running with Docker
1. Clone the repository:
```bash
git clone https://github.com/joeCavZero/simple-blockchain.git
cd simple-blockchain
```
2. Build the Docker image:
```bash
sh scripts/docker-build.sh
```
3. Run the Docker container:
```bash
sh scripts/docker-run.sh
```

# API

## GET
- ```/api/blocks``` -
 Retorn all blocks from the blockchain.
 
- ```/api/blocks/{index}``` - Returns a specific block by its index.

- ```/api/validate``` - Validates the integrity of the blockchain.

## POST

- ```/api/mine``` - Mines a new block with the data provided in the request body.
```json
{
    "data": "Daniel da Silva Cavalcante bought the LATAM Airlines"
}
```

- ```/api/difficulty``` - Sets the mining difficulty.
```json
{
    "difficulty": 6
}
```