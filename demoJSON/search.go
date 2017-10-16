package main

import (
	"os"
	"fmt"
	"reflect"
)

// handles the search functionality of the application
// will check to see if there was an argument passed with the search message
// if argument, will search for that argument
// else, will prompt the user to enter what they are searching for ** take place inside the prompt file!
// will load the database
// will parse the database into a []MemberInfo ** handle in the marshallData file!
// will return a slice of the objects found
// will print these objects to the screen

func SearchInsideFile()  {
	fmt.Println("Search inside file")
	var returnedValue []MemberInfo
	if ok := isThirdArgument(); ok {
		returnedValue = Search(os.Args[2])
	} else {
		searchAnswers := PromptUser(getSearchQuestions())
		returnedValue = Search(searchAnswers[0])
	}

	PrintResults(returnedValue)
	} // search inside file

// check to see if there is a third argument
func isThirdArgument() bool  {
	if len(os.Args) < 3 {
		return false
	}
	return  true
}

// search functions
func Search(value string) []MemberInfo {
	// unmarshal the data into structs
	payload, err := ConvertRawDataIntoStruct(value)
	ExitIfError(err)
	results := filterPayload(payload, value)

	return results
}

// take in a whole payload
// for each payload, extract the memberInfo
// for each field in the memberInfo, search for value
// if value is found, return the whole memberInfo object
func filterPayload(payload []Payload, searchValue string) []MemberInfo {

	var filteredValues []MemberInfo

	for _, p := range payload {
		for _, mi := range p {
			v := reflect.ValueOf(mi)
			for i := 0; i < v.NumField(); i++ {
				if searchValue == v.Field(i).Interface(){
					filteredValues = append(filteredValues, mi)
				}
			}
		}
	}

	return filteredValues
}



