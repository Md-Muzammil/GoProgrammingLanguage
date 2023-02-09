//Go - Program Structure
// This is my first sample Go program.
//Let us take a look at the various parts of the above program

package main

/*The first line of the program package main defines the package name in which this program should lie.
It is a mandatory statement, as Go programs run in packages.
The main package is the starting point to run the program.
Each package has a path and name associated with it.*/

import "fmt"

/* import "fmt" is a preprocessor command which tells the Go compiler to include the files lying in the package fmt. */
/* func main() is the main function where the program execution begins. */
func main() {
	fmt.Print("Hello! Muzammel Hoque\n\tWelcome to Golang Programming")
}

/* fmt.Println, Printf, Print (...) is another function available in Go which causes the message "Hello, World!" to be displayed on the screen.
Here fmt package has exported Println method which is used to display the message on the screen. */
