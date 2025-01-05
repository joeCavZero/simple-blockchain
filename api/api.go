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

	router.HandleFunc("/api/blocks", a.createBlock).Methods("POST")

	router.HandleFunc("/api/blocks/{index}", a.getBlock).Methods("GET")

	router.HandleFunc("/api/validate", a.validate).Methods("GET")

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

func (a *Api) createBlock(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var body_data map[string]string
	err := json.NewDecoder(r.Body).Decode(&body_data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	data, ok := body_data["data"]
	if !ok {
		http.Error(w, "Data field is required", http.StatusBadRequest)
		return
	}

	block := a.blockchain.AddBlock(data)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(block)
}

func (a *Api) validate(w http.ResponseWriter, r *http.Request) {
	err := a.blockchain.ValidateChain()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusOK)
}
