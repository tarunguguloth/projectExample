
package main

// Importing fmt and time
import "fmt"
import "time"

// Calling main
func main() {

	// Declaring layout constant
	//const layout = "Jan 2, 2006 at 3:04pm (MST)"

	// Calling Parse() method with its parameters
	tm, _ := time.Parse("020106150405", "110399000000")

	// Returns output
	fmt.Println(tm)
}