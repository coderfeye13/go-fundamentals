package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Pet represents a single pet from the Petstore API
type Pet struct {
	ID        int64    `json:"id"`
	Name      string   `json:"name"`
	Status    string   `json:"status"`
	PhotoURLs []string `json:"photoUrls"`
}

// getPetsByStatus fetches all pets with the given status
func getPetsByStatus(status string) ([]Pet, error) {
	url := fmt.Sprintf("https://petstore.swagger.io/v2/pet/findByStatus?status=%s", status)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("GET request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("reading body failed: %w", err)
	}

	var pets []Pet
	err = json.Unmarshal(body, &pets)
	if err != nil {
		return nil, fmt.Errorf("JSON parse failed: %w", err)
	}

	return pets, nil
}

// getPetByID fetches a single pet by ID
func getPetByID(id int64) (*Pet, error) {
	url := fmt.Sprintf("https://petstore.swagger.io/v2/pet/%d", id)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("GET request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("reading body failed: %w", err)
	}

	var pet Pet
	err = json.Unmarshal(body, &pet)
	if err != nil {
		return nil, fmt.Errorf("JSON parse failed: %w", err)
	}

	return &pet, nil
}

// createPet sends a POST request to create a new pet
func createPet(pet Pet) (*Pet, error) {
	// Marshal: Go struct → JSON bytes
	data, err := json.Marshal(pet)
	if err != nil {
		return nil, fmt.Errorf("marshal failed: %w", err)
	}

	resp, err := http.Post(
		"https://petstore.swagger.io/v2/pet",
		"application/json",
		bytes.NewBuffer(data),
	)
	if err != nil {
		return nil, fmt.Errorf("POST request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("reading body failed: %w", err)
	}

	var created Pet
	err = json.Unmarshal(body, &created)
	if err != nil {
		return nil, fmt.Errorf("JSON parse failed: %w", err)
	}

	return &created, nil
}

// deletePet sends a DELETE request for the given pet ID
func deletePet(id int64) error {
	url := fmt.Sprintf("https://petstore.swagger.io/v2/pet/%d", id)

	// http.Delete doesn't exist — we need a custom request
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return fmt.Errorf("creating DELETE request failed: %w", err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("DELETE request failed: %w", err)
	}
	defer resp.Body.Close()

	fmt.Printf("DELETE status: %d\n", resp.StatusCode)
	return nil
}

func main() {
	// --- GET all available pets ---
	fmt.Println("=== GET: Available Pets ===")
	pets, err := getPetsByStatus("available")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Found %d pets\n", len(pets))

	// --- GET single pet ---
	fmt.Println("\n=== GET: Single Pet ===")
	pet, err := getPetByID(pets[0].ID)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Pet: ID=%d Name=%s Status=%s\n", pet.ID, pet.Name, pet.Status)

	// --- POST: Create new pet ---
	fmt.Println("\n=== POST: Create Pet ===")
	newPet := Pet{
		ID:        12345678,
		Name:      "GoLand_Dog",
		Status:    "available",
		PhotoURLs: []string{"https://example.com/goland_dog.jpg"},
	}
	created, err := createPet(newPet)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Created: ID=%d Name=%s\n", created.ID, created.Name)

	// --- DELETE: Remove the pet we just created ---
	fmt.Println("\n=== DELETE: Remove Pet ===")
	err = deletePet(created.ID)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Pet deleted successfully!")
}
