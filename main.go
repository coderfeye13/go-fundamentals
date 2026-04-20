package main

import "fmt"

type Pet struct {
	ID     int
	Name   string
	Status string
}

func filterPets(pets []Pet, status string) []Pet {
	result := []Pet{}
	for _, pet := range pets {
		if pet.Status == status {
			result = append(result, pet)
		}
	}
	return result
}
func countByStatus(pets []Pet) map[string]int {
	// your code here
	resultC := map[string]int{} //empty map

	for _, pet := range pets {
		resultC[pet.Status]++
	}
	return resultC
}

func main() {
	pets := []Pet{
		{1, "Karamel", "available"},
		{2, "Doggie", "sold"},
		{3, "Fluffy", "available"},
		{4, "Rex", "pending"},
		{5, "Buddy", "sold"},
		{6, "Max", "available"},
	}

	result := filterPets(pets, "available")
	fmt.Println(result)
	// expected: [{1 Karamel available} {3 Fluffy available}]

	counts := countByStatus(pets)
	fmt.Println(counts)
	// expected: map[available:3 pending:1 sold:2]
}
