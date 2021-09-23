package main

import (
    "fmt"
    "log"
    "net/http"
)

func process(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.Error(w, "404 - your requested page not found", http.StatusNotFound)
		return
	}

	//GET and POST request handling
	switch r.Method {
	case "GET":
		http.ServeFile(w, r,"form.html")

	case "POST":
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w,"ParseForm() error: %v", err)
			return
		}
		name := r.FormValue("name")
		occupation := r.FormValue("occupation")

		fmt.Fprintf(w, "%s is an %s\n", name, occupation)

	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported")
	}
}

func main() {

	http.HandleFunc("/", process)
	fmt.Printf("Starting the server at port 2224\n")
	log.Fatal(http.ListenAndServe(":2224",nil))
}


