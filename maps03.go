package main
import "fmt"

func main() {
	var fileExtension = map[string]string {
		"Python": ".py",
		"Java": ".java",
		"C++": ".cpp",
		"GoLang": ".go",
		"Anisable": ".yml",
	}

	fmt.Println(fileExtension);

	delete(fileExtension,"C++")

	fileExtension["Julia"] = ".jil"

	fmt.Println("Now the Maps is ", fileExtension);


}
