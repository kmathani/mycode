package main

import (
    "bytes"
    "encoding/json"
    "fmt"
    "log"
    "net/http"
)

func main() {

    // A map is serialized into JSON string with json.Marshal()
    values := map[string]string{"name": "John Doe", "occupation": "gardener"}
    
    json_data, err := json.Marshal(values)
    if err != nil {
	    log.Fatal(err)
    }
    resp, err := http.Post("http://httpbin.org/post", "application/json",bytes.NewBuffer(json_data))
    if err != nil {
        log.Fatal(err)
    }
    var res map[string] interface{}

    json.NewDecoder(resp.Body).Decode(&res)

    fmt.Println(res["json"])
}
