package main
import  (
	"fmt"
	"os/exec"
	"io/ioutil"
	"log"
)
func main() {
    cmd := exec.Command("nslookup google.com")

    // this is new -- our command will output through this pipe
    stdout, err := cmd.StdoutPipe()

    if err != nil {
        log.Fatal(err)
    }

    // if we use Start, we must use wait to ensure our command will finish
    if err := cmd.Start(); err != nil {
        log.Fatal(err)
    }

    // grab the stdout from our command and read it in as string data
    data, err := ioutil.ReadAll(stdout)

    if err != nil {
        log.Fatal(err)
    }

    if err := cmd.Wait(); err != nil {
        log.Fatal(err)
    }

    // display the results to the screen
    fmt.Printf("%s\n", string(data))
}
