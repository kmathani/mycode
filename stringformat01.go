package main
import "fmt"
func main() {
	fmt.Printf("bool: %t\n", true);

	fmt.Printf("int: %d\n", 12344)

	fmt.Printf("binary: %b\n",14)

	fmt.Printf("char: %c\n", 33)

	fmt.Printf("hex: %x\n",456)

	fmt.Printf("float: %f\n",78.99)

	fmt.Printf("float2: %e\n" ,1232323232323.4343)
	fmt.Printf("float3: %E\n" ,23233434344.3434)

	// basic string printing use %s
        fmt.Printf("str1: %s\n", "\"string\"")

	fmt.Printf("str2: %q\n", "\"Quoted String \"")

	// %x renders the string in base-16
        //  with two output characters per byte of input
	
	fmt.Printf("str3:  %x\n", "hex this");

	// specify the width of an integer, use a number after the % in the verb
       // By default the result will be 
        // right-justified and padded with spaces
	fmt.Printf("width1: |%6d|%6d|\n", 12,456)

	    // specify the width of printed floats
    fmt.Printf("width2: |%6.2f|%6.2f|\n", 1.2, 3.45)

    // left-justify, use the - flag
    fmt.Printf("width3: |%-6.2f|%-6.2f|\n", 1.2, 3.45)

    // left-justify use the - flag as with numbers
    fmt.Printf("width4: |%6s|%6s|\n", "foo", "b")


    fmt.Printf("width5: |%-6s|%-6s|\n", "foo", "b")

    //  Sprintf formats and returns a string without printing it anywhere
    s := fmt.Sprintf("sprintf: a %s", "string")
    fmt.Println(s)

    // format+print to io.Writers other than os.Stdout using Fprintf.
    //fmt.Fprintf(os.Stderr, "io: an %s\n", "error")
 



}

