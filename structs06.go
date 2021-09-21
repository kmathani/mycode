package main

import "fmt"

type astro struct {
	name string
	age int
	mission string
	isneeded bool
}

type nasaMission struct {
	people  []astro 
	number int
	message string
}


func main() {
	ast1 := astro{"Megan McArthur",35, "ISS" , false}
	ast2 := astro{"Barry Wilmore", 43, "Boening Crew Flight Test", true}
	ast3 := astro{"Rahim Benzma",35, "Space-X Crew-3", true}

	fmt.Println(ast1);
	fmt.Println(ast2);
	fmt.Println(ast3);

	p := []astro{ast1, ast2, ast3}

	//display the slice
	fmt.Println(p)

	fmt.Println(p[1].mission)

	s := nasaMission{p,3,"success"}

	fmt.Println("Simple Print" , s);
	fmt.Printf("%+v", s);

}

