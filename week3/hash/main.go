package main

import "fmt"

func main() {
	type Person struct {
		Name  string
		Likes []string
	}

	var people [3]*Person
	likes := make(map[string][]*Person)

	people[0] = &Person{
		Name:  "Rod",
		Likes: []string{"cheese", "bacon"},
	}

	people[1] = &Person{
		Name:  "Iara",
		Likes: []string{"cheese", "aveia"},
	}

	people[2] = &Person{
		Name:  "Fred",
		Likes: []string{"bacon", "aveia"},
	}

	for _, p := range people {
		for _, l := range p.Likes {
			likes[l] = append(likes[l], p)
		}
	}

	for _, p := range likes["cheese"] {
		fmt.Println(p.Name, "likes cheese.")
	}

	fmt.Println(len(likes["bacon"]), "people like bacon.")
	fmt.Printf("%#v\n", likes)

	fmt.Printf("Len of likes: %d\n", len(likes))

	v, ok := likes["cuzcuz"]
	fmt.Println(v, ok)
	v, ok = likes["cheese"]
	fmt.Println(v, ok)

}
