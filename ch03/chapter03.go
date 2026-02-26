package main

import (
	"fmt"
	"maps"
	"slices"
)

func main() {
	fmt.Println("===== ARRAYS =====")
	arraysDemo()

	fmt.Println("\n===== SLICES BASICS & SLICING =====")
	slicesBasicsDemo()

	fmt.Println("\n===== SLICES: SHARED MEMORY, APPEND, LEN/CAP =====")
	slicesAppendLenCapDemo()

	fmt.Println("\n===== SLICES: copy() =====")
	copyDemo()

	fmt.Println("\n===== MAPS BASICS (nil vs empty), LEN, READ/WRITE =====")
	mapsBasicsDemo()

	fmt.Println("\n===== MAPS: COMMA-OK IDIOM =====")
	commaOkDemo()

	fmt.Println("\n===== MAPS: delete() and clear() =====")
	deleteClearDemo()

	fmt.Println("\n===== COMPARING SLICES & MAPS (Go 1.21+) =====")
	compareSlicesMapsDemo()

	fmt.Println("\n===== MAP AS SET (bool vs struct{}) =====")
	setDemo()
}

func arraysDemo() {
	// Arrays: size is part of the type ([3]int != [4]int).
	// TR: Array value type'tır; assignment deep copy yapar.
	// EN: Arrays are value types; assignment copies the full array.

	x := [...]int{1, 2, 3} // inferred length => [3]int
	y := [3]int{10, 20, 30}
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("x == y:", x == y) // comparable

	a := [3]int{1, 2, 3}
	b := a     // deep copy
	b[0] = 100 // changes b only
	fmt.Println("a after b changed:", a)
	fmt.Println("b:", b)

	// Uncomment to see compile-time type mismatch:
	// var p [3]int
	// var q [4]int
	// p = q // cannot use q (type [4]int) as type [3]int
}

func slicesBasicsDemo() {
	// Slices: []T, length not part of type. Zero value is nil.
	// TR: Slice comparable değildir; sadece nil ile karşılaştırılabilir.
	// EN: Slices are not comparable (except to nil).

	var nilSlice []int
	fmt.Println("nilSlice == nil:", nilSlice == nil, "len:", len(nilSlice), "cap:", cap(nilSlice))

	x := []string{"a", "b", "c", "d"}
	y := x[:2]  // [a b]
	z := x[1:]  // [b c d]
	d := x[1:3] // [b c]
	e := x[:]   // full slice
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)
	fmt.Println("d:", d)
	fmt.Println("e:", e)

	// Slices share underlying array
	d[0] = "X" // changes x[1]
	fmt.Println("after d[0] = X -> x:", x, "d:", d)
}

func slicesAppendLenCapDemo() {
	// make([]T, length, capacity)
	// TR: length görünen kısım, capacity büyüyebileceğin sınır.
	// EN: length is visible size, capacity is growth limit before reallocation.

	a := make([]int, 3, 10)
	fmt.Println("a:", a, "len:", len(a), "cap:", cap(a))

	a[0], a[1], a[2] = 1, 2, 3
	fmt.Println("a filled:", a, "len:", len(a), "cap:", cap(a))

	// append may or may not allocate a new underlying array depending on cap
	b := append(a, 4) // cap is enough => shares underlying array with a
	fmt.Println("b := append(a,4) -> b:", b, "len:", len(b), "cap:", cap(b))

	// You can "see" appended element by extending a within its capacity
	fmt.Println("a[:4]:", a[:4]) // [1 2 3 4]

	// Show shared memory effect
	b[0] = 100
	fmt.Println("after b[0]=100:")
	fmt.Println("a (len 3):", a)
	fmt.Println("a[:4]:", a[:4])
	fmt.Println("b:", b)

	// Now show a case where append likely allocates a new array (cap full)
	c := []int{1, 2, 3} // typically len=3 cap=3
	d := append(c, 4)   // likely reallocates
	d[0] = 999
	fmt.Println("c:", c, "len:", len(c), "cap:", cap(c))
	fmt.Println("d:", d, "len:", len(d), "cap:", cap(d))
	fmt.Println("note: if append reallocated, c[0] stays 1; otherwise it could change (rare).")
}

func copyDemo() {
	// copy(dst, src) copies min(len(dst), len(src)) elements and returns count.
	// TR: capacity değil length önemli.
	// EN: length matters, not capacity.

	x := []int{1, 2, 3, 4}
	y := make([]int, 4)
	n := copy(y, x)
	fmt.Println("copy full -> y:", y, "copied:", n)

	// Copy subset
	y2 := make([]int, 2)
	n2 := copy(y2, x)
	fmt.Println("copy subset -> y2:", y2, "copied:", n2)

	// Copy from middle
	y3 := make([]int, 2)
	copy(y3, x[2:]) // [3 4]
	fmt.Println("copy from middle -> y3:", y3)

	// Overlapping copy (safe)
	x2 := []int{1, 2, 3, 4}
	n3 := copy(x2[:3], x2[1:]) // shift left by 1
	fmt.Println("overlap copy -> x2:", x2, "copied:", n3)

	// copy with arrays via slicing
	arr := [4]int{5, 6, 7, 8}
	y4 := make([]int, 2)
	copy(y4, arr[:])
	fmt.Println("copy array->slice -> y4:", y4)

	x3 := []int{1, 2, 3, 4}
	copy(arr[:], x3)
	fmt.Println("copy slice->array -> arr:", arr)
}

func mapsBasicsDemo() {
	// map[key]value; zero value is nil.
	var nilMap map[string]int
	fmt.Println("nilMap == nil:", nilMap == nil, "len:", len(nilMap))
	fmt.Println("read missing from nilMap:", nilMap["missing"]) // returns 0
	// Writing to nilMap would panic:
	// nilMap["x"] = 1 // panic

	emptyMap := map[string]int{} // not nil; safe to write
	fmt.Println("emptyMap == nil:", emptyMap == nil, "len:", len(emptyMap))
	emptyMap["hello"] = 5
	emptyMap["world"] = 0
	fmt.Println("emptyMap after writes:", emptyMap, "len:", len(emptyMap))

	// Non-empty literal with slice values
	teams := map[string][]string{
		"Orcas":   {"Fred", "Ralph", "Bijou"},
		"Lions":   {"Sarah", "Peter", "Billie"},
		"Kittens": {"Waldo", "Raul", "Ze"},
	}
	fmt.Println("teams:", teams)

	// make with initial capacity hint
	ages := make(map[int][]string, 10)
	fmt.Println("ages len initially:", len(ages))
	ages[20] = []string{"Furkan"}
	fmt.Println("ages:", ages)
}

func commaOkDemo() {
	m := map[string]int{
		"hello": 5,
		"world": 0,
	}

	v, ok := m["hello"]
	fmt.Println(`m["hello"] ->`, v, ok) // 5 true

	v, ok = m["world"]
	fmt.Println(`m["world"] ->`, v, ok) // 0 true (key exists, value is zero)

	v, ok = m["goodbye"]
	fmt.Println(`m["goodbye"] ->`, v, ok) // 0 false (key missing)
}

func deleteClearDemo() {
	m := map[string]int{
		"hello": 5,
		"world": 10,
	}
	fmt.Println("before delete:", m, "len:", len(m))
	delete(m, "hello") // safe even if key missing
	fmt.Println("after delete hello:", m, "len:", len(m))

	// delete on nil map is safe (no-op)
	var nilMap map[string]int
	delete(nilMap, "x") // no panic
	fmt.Println("delete on nil map: ok")

	// clear empties the map (Go 1.21+)
	m2 := map[string]int{"hello": 5, "world": 10}
	fmt.Println("before clear:", m2, "len:", len(m2))
	clear(m2)
	fmt.Println("after clear:", m2, "len:", len(m2))
}

func compareSlicesMapsDemo() {
	// slices.Equal requires element type to be comparable
	x := []int{1, 2, 3, 4, 5}
	y := []int{1, 2, 3, 4, 5}
	z := []int{1, 2, 3, 4, 5, 6}
	fmt.Println("slices.Equal(x,y):", slices.Equal(x, y)) // true
	fmt.Println("slices.Equal(x,z):", slices.Equal(x, z)) // false

	// maps.Equal compares key/value pairs (order doesn't matter)
	m := map[string]int{"hello": 5, "world": 10}
	n := map[string]int{"world": 10, "hello": 5}
	fmt.Println("maps.Equal(m,n):", maps.Equal(m, n)) // true
}

func setDemo() {
	// Map as set with bool values
	intSet := map[int]bool{}
	vals := []int{5, 10, 2, 5, 8, 7, 3, 9, 1, 2, 10}
	for _, v := range vals {
		intSet[v] = true
	}
	fmt.Println("len(vals):", len(vals), "len(intSet):", len(intSet)) // duplicates removed
	fmt.Println("intSet[5]:", intSet[5])
	fmt.Println("intSet[500]:", intSet[500])
	if intSet[100] {
		fmt.Println("100 is in the set")
	} else {
		fmt.Println("100 is NOT in the set")
	}

	// Map as set with struct{} (zero-byte value), requires comma-ok to check membership clearly
	intSet2 := map[int]struct{}{}
	for _, v := range vals {
		intSet2[v] = struct{}{}
	}
	if _, ok := intSet2[5]; ok {
		fmt.Println("5 is in the struct{} set")
	}
	if _, ok := intSet2[500]; !ok {
		fmt.Println("500 is NOT in the struct{} set")
	}
}
