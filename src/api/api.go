package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/joeCavZero/simple-blockchain/src/blockchain"

	"github.com/joeCavZero/simple-blockchain/src/logkit"

	"github.com/gorilla/mux"
)

var apilk = logkit.NewLogkit("API")

type Api struct {
	blockchain *blockchain.Blockchain
}

func NewApi(bc *blockchain.Blockchain) *Api {
	return &Api{
		blockchain: bc,
	}
}

func (a *Api) Run(addr uint16) {
	router := mux.NewRouter()

	router.HandleFunc("/api/blocks", a.getBlocks).Methods("GET")

	router.HandleFunc("/api/blocks/{index}", a.getBlock).Methods("GET")

	router.HandleFunc("/api/validate", a.validate).Methods("GET")

	router.HandleFunc("/api/mine", a.mine).Methods("POST")

	router.HandleFunc("/api/difficulty", a.setDifficulty).Methods("POST")

	apilk.Infof("API is running on port %d", addr)

	err := http.ListenAndServe(
		fmt.Sprintf(":%d", addr),
		router,
	)
	if err != nil {
		panic(err)
	}
}

func (a *Api) getBlocks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	blocks := a.blockchain.GetChain()

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(blocks)

	ip := r.RemoteAddr
	apilk.Info("List of blocks sent to the client", ip)
}

func (a *Api) getBlock(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	raw_index, ok := params["index"]
	if !ok {
		http.Error(w, "Index is required", http.StatusBadRequest)
		return
	}

	index, err := strconv.Atoi(raw_index)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	ip := r.RemoteAddr

	block, err := a.blockchain.GetBlock(uint64(index))
	if err != nil {
		apilk.Error("Block", raw_index, "not found", ip)
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(block)

	apilk.Info("Block", raw_index, "sent to the client", ip)
}

func (a *Api) mine(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var body_data map[string]string
	err := json.NewDecoder(r.Body).Decode(&body_data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// data
	data, ok := body_data["data"]
	if !ok {
		http.Error(w, "'data' is required", http.StatusBadRequest)
		return
	}

	newBlock := a.blockchain.CreateBlock(data)

	mining_result := a.blockchain.Mine(&newBlock)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(mining_result)

	ip := r.RemoteAddr
	apilk.Infof(
		"Block %d mined by the client %s",
		mining_result.MinedBlock.Index,
		ip,
	)
}

func (a *Api) validate(w http.ResponseWriter, r *http.Request) {
	err := a.blockchain.ValidateChain()
	ip := r.RemoteAddr
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		apilk.Error("Chain validation failed", ip)
		return
	}

	w.WriteHeader(http.StatusOK)

	apilk.Info("Chain validated by the client", ip)
}

func (a *Api) setDifficulty(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var body_data map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&body_data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// difficulty
	interface_difficulty, ok := body_data["difficulty"]
	if !ok {
		http.Error(w, "'difficulty' is required", http.StatusBadRequest)
		return
	}
	difficulty_float, ok := interface_difficulty.(float64)
	if !ok {
		http.Error(w, "'difficulty' must be a number", http.StatusBadRequest)
		return
	}

	if difficulty_float < 0 {
		http.Error(w, "'difficulty' must be greater than 0", http.StatusBadRequest)
		return
	}

	difficulty := uint64(difficulty_float)
	if !ok {
		http.Error(w, "'difficulty' must be an uint64", http.StatusBadRequest)
		return
	}

	a.blockchain.SetDifficulty(difficulty)
	w.WriteHeader(http.StatusOK)

	ip := r.RemoteAddr
	apilk.Infof(
		"Difficulty set to %d by the client %s",
		difficulty,
		ip,
	)
}
