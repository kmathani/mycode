package main
import  (
	"database/sql"
	"fmt"
	"log"

	// when a package is imported prefix with a blank identifier
	// the init function of the package is called
	// the function registers the driver
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	// open the datavase
	// db name, connection information
	db, err :=  sql.Open("mysql", "student:playstation5@tcp(127.0.0.1:3306)/testdb")

	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	var version string

	err2 := db.QueryRow("SELECT VERSION()").Scan(&version)

	if err2 != nil {
		log.Fatal(err2)
	}

	fmt.Println(version)
}

