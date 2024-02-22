package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"os"
	"time"
)

const path = "./entries.json"

// raffleEntry is the struct we unmarshal raffle entries into
type raffleEntry struct {
	ID   string `json:"id"`
	NAME string `json:"name"`
}

// importData reads the raffle entries from file and creates the entries slice.
func importData() []raffleEntry {
	file, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	var entries []raffleEntry

	err = json.Unmarshal(file, &entries)
	if err != nil {
		log.Fatal(err)
	}

	return entries
}

// getWinner returns a random winner from a slice of raffle entries.
func getWinner(entries []raffleEntry) raffleEntry {
	source := rand.NewSource(time.Now().Unix())
	r := rand.New(source)
	winner := r.Intn(len(entries))

	return entries[winner]
}

func main() {
	entries := importData()
	log.Println("And... the raffle winning entry is...")
	winner := getWinner(entries)
	time.Sleep(500 * time.Millisecond)
	log.Println(winner)
}
