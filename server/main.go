package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type DeckResponse struct {
	Success bool `json:"success"`
	Deck    Deck `json:"deck"`
}

type ShuffleRequest struct {
	Deck Deck `json:"deck"`
}

type DealRequest struct {
	Deck           Deck `json:"deck"`
	Players        int  `json:"players"`
	CardsPerPlayer int  `json:"cards_per_player"`
}

func main() {
	http.HandleFunc("/create/deck", APICreate_Deck)
	http.HandleFunc("/shuffle", APIShuffle_Deck)
	http.HandleFunc("/deal", APIDeal)
	fmt.Println("Server started")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func APICreate_Deck(w http.ResponseWriter, r *http.Request) {
	deck := Create_Deck()
	resp := DeckResponse{
		Success: true,
		Deck:    deck,
	}
	jresp, _ := json.Marshal(resp)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jresp)

}

func APIShuffle_Deck(w http.ResponseWriter, r *http.Request) {
	var req ShuffleRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		fmt.Println("Error Decoding Request")
	}

	deck := req.Deck
	deck.Shuffle()
	resp := DeckResponse{
		Success: true,
		Deck:    deck,
	}

	jresp, _ := json.Marshal(resp)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jresp)
}

func APIDeal(w http.ResponseWriter, r *http.Request) {
	var req DealRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		fmt.Println("Error Decoding Request")
	}

	deck := req.Deck
	players := req.Players
	cardsper := req.CardsPerPlayer
	deck.Deal(players, cardsper)

	resp := DeckResponse{
		Success: true,
		Deck:    deck,
	}

	jresp, _ := json.Marshal(resp)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jresp)
}
