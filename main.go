package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Struct for character data
type Character struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Specie string `json:"species"`
	Status string `json:"status"`
}

// Function to fetch character data from API
func getCharacterData(url string) (Character, error) {
	response, err := http.Get(url)
	if err != nil {
		return Character{}, err
	}
	defer response.Body.Close()

	var character Character
	err = json.NewDecoder(response.Body).Decode(&character)
	if err != nil {
		return Character{}, err
	}
	return character, nil
}

// Handler for the /check_endpoint endpoint
func checkEndpoint(w http.ResponseWriter, r *http.Request) {
	// Fetch character data from first API
	character1, err := getCharacterData("https://rickandmortyapi.com/api/character/222")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Fetch character data from second API
	character2, err := getCharacterData("https://rickandmortyapi.com/api/character/222")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Compare character data
	if character1 == character2 {
		// If data is the same, return status code 200 (OK)
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Characters are the same")
	} else {
		// If data is different, return status code 401 (Unauthorized)
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "Characters are different")
	}
}

func main() {
	// Define endpoint handler
	http.HandleFunc("/check_endpoint", checkEndpoint)

	// Start HTTP server
	fmt.Println("Server listening on port 8082...")
	log.Fatal(http.ListenAndServe(":8082", nil))
}
