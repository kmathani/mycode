package main

import "fmt"

type astro struct {
	name string
	age int
	mission string
	isneeded bool
}

func main() {
	ast1 := astro{"Megan McArthur",35, "ISS" , false}
	ast2 := astro{"Barry Wilmore", 43, "Boening Crew Flight Test", true}
	ast3 := astro{"Rahim Benzma",35, "Space-X Crew-3", true}

	fmt.Println(ast1);
	fmt.Println(ast2);
	fmt.Println(ast3);
}

