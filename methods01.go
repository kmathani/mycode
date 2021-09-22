package main
import "fmt"

type author struct {
	name string
	branch string
	books int
	salary int
}

func (a author) show() {
	fmt.Println("Author Name:", a.name)
	fmt.Println("Branch Name:",a.branch)
	fmt.Println("Published Articles:", a.books)
	fmt.Println("Salary:",a.salary)
}

func (a *author) bookup() {
	a.books ++
}


func main() {
	res := author{
		name: "RZFreez",
		branch: "Pennsylvania",
		books: 14,
		salary: 34000,
	}

	res.show()

	res.bookup()

	res.show()
}

