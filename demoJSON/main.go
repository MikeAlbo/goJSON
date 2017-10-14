package main

import "fmt"

func main()  {

InitCLI()
}


func ReadDataFromFile()  {
	fmt.Println("reading from file...")
	data, err := ReadFromJSON("../demoData1.json")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}

func WriteDataToFile()  {
	fmt.Println("writing to file...")
	answers := PromptUser()
	x := BuildStruct(answers)
	if err := WriteToJSON(x); err != nil {
		panic(err)
	}

}

