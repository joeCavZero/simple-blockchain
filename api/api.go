package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"project/blockchain"
	"strconv"

	"github.com/gorilla/mux"
)

type Api struct {
	blockchain *blockchain.Blockchain
}

func NewApi(bc *blockchain.Blockchain) *Api {
	return &Api{
		blockchain: bc,
	}
}

func (a *Api) Run(addr string) {
	router := mux.NewRouter()

	router.HandleFunc("/api/blocks", a.getBlocks).Methods("GET")

	router.HandleFunc("/api/blocks/{index}", a.getBlock).Methods("GET")

	router.HandleFunc("/api/validate", a.validate).Methods("GET")

	router.HandleFunc("/api/mine", a.mine).Methods("POST")

	router.HandleFunc("/api/difficulty", a.setDifficulty).Methods("POST")

	fmt.Printf("Listening on port %s\n", addr)
	err := http.ListenAndServe(addr, router)
	if err != nil {
		panic(err)
	}
}

func (a *Api) getBlocks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	blocks := a.blockchain.GetChain()

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(blocks)
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

	block, err := a.blockchain.GetBlock(uint64(index))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(block)
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

	mining_result := a.blockchain.Mine(newBlock)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(mining_result)
}

func (a *Api) validate(w http.ResponseWriter, r *http.Request) {
	err := a.blockchain.ValidateChain()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusOK)
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
}
