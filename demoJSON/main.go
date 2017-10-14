package main

import (
	"fmt"
	"os"
)

func main()  {
InitCLI()
}


func ReadDataFromFile()  {
	fmt.Println("reading from file...")
	s, err := ReadFromJSON("../demoData1.json")
	ExitIfError(err)
	fmt.Println(string(s))
}

func WriteDataToFile()  {
	fmt.Println("writing to file...")
	userInput := ProcessInputData(PromptUser())
	WriteToJSON(userInput)

}

func ExitIfError(err error)  {
	if err != nil {
		fmt.Errorf(err.Error())
		os.Exit(1)
	}
}