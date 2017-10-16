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
		returnedValue = Search(os.Args[2:])
	} else {
		searchAnswers := PromptUser(getSearchQuestions())
		returnedValue = Search(searchAnswers)
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
func Search(value []string) []MemberInfo {
	// unmarshal the data into structs
	payload, err := ConvertRawDataIntoStruct()
	ExitIfError(err)
	results := filterPayload(payload, value)

	return results
}

// take in a whole payload
// for each payload, extract the memberInfo
// for each field in the memberInfo, search for value
// if value is found, return the whole memberInfo object
func filterPayload(payload []Payload, searchValue []string) []MemberInfo {

	var preFilteredValues []MemberInfo
	var secondaryFilteredValues []MemberInfo

		for _, p := range payload {
			for _, mi := range p {
				v := reflect.ValueOf(mi)
				for i := 0; i < v.NumField(); i++ {
					if searchValue[0] == v.Field(i).Interface(){
						preFilteredValues = append(preFilteredValues, mi)
					}
				}
			}
		}
		if len(searchValue) > 1 {
			for _, sv := range searchValue[1:]{
				for _, preFilteredValue := range preFilteredValues {
					v := reflect.ValueOf(preFilteredValue)
					for i := 0; i < v.NumField(); i++ {
						if sv == v.Field(i).Interface(){
							secondaryFilteredValues = append(secondaryFilteredValues, preFilteredValue)
						}
					}
				}
			}
			return secondaryFilteredValues
		}

	return preFilteredValues
}



