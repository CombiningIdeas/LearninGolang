package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

type ByName []Person

func (a ByName) Len() int           { return len(a) }
func (a ByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByName) Less(i, j int) bool { return a[i].Name < a[j].Name }

func main() {
	people := []Person{
		{Name: "Alice", Age: 25},
		{Name: "Charlie", Age: 30},
		{Name: "Bob", Age: 20},
	}

	sort.Sort(ByName(people))

	for _, person := range people {
		fmt.Println(person.Name, person.Age)
	}
}
