package exercises

type Person struct {
	// Üç alan: FirstName, LastName string | Age int
}

func MakePerson(firstName, lastName string, age int) Person {
	// Person döndür (pointer değil)
}

func MakePersonPointer(firstName, lastName string, age int) *Person {
	// *Person döndür
}

func main() {
	// her iki fonksiyonu çağır
	// sonuçları bir değişkene ata
}
