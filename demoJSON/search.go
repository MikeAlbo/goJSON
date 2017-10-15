package main

import (
	"os"
)

// handles the search functionality of the application
// will check to see if there was an argument passed with the search message
// if argument, will search for that argument
// else, will prompt the user to enter what they are searching for
// will load the database
// will parse the database into a []MemberInfo ** handle in the marshallData file!
// will return a slice of the objects found
// will print these objects to the screen

func SearchInsideFile()  {
	if ok := isthirdArgument(); ok {
		// call search with arg
	}
	// prompt user 
}

// check to see if there is a third argument 
func isthirdArgument() bool  {
	if len(os.Args) < 3 {
		return false
	}
	return  true
}

// prompt the user for input from the cli
func promptUserForSearchTerm() []string  {
	questions := []Question{{"Please enter a search term\n"}}



}