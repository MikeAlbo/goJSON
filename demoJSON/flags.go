package main

import (
	"flag"
	"os"
	"fmt"
)

// flags to set either reading from db or writing to

var read = flag.Bool("read", false, "Read the database file")
var write = flag.Bool("write", false, "Add an Item to the database file")
var search = flag.Bool("search", false, "Search for a name in the database file")
var delete = flag.Bool("delete", false, "Delete an item form the database file")

func InitCLI(){

	if len(os.Args) < 2 {
		fmt.Println("Please indicate wether you want to -read or -write to database")
		os.Exit(1)
	}

	flag.Parse()

	switch {
	case *read: ReadDataFromFile()
	case *write: WriteDataToFile()
	case *search: SearchInsideFile()
	case *delete: DeleteItemFromFile()
	default:
		invalidCase()
	}


}

// handles invalid case, usually caught in flag.Parse(), this is just a backup
func invalidCase()  {
	fmt.Printf("%s, is an invalid argument, please use -read or -write", os.Args[2])
	os.Exit(1)
}