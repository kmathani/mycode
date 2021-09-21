package main

import (
 "fmt"
)

type User struct {
 Name string
 Pets []string
}

func (u User) newPet() {
 u.Pets = append(u.Pets, "Lucy")                    // Simple function should ensure "Fido" is added to User
 fmt.Println(u)                                     // This user is **NOT** the same user as the one in main()!!
}


func  main() {

	u := User{Name: "Anna", Pets: []string{"Baily"}}
	u.newPet()
	fmt.Println(u)
}

