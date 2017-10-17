package main

import (
	"os"
	"fmt"
	"strings"
)

// removeItem contains the methods needed to remove an item from the database
// will call the search functions
// will take the results fo the search functions
// will verify that the result is len 1
// if not, will prompt user to filter down
// will confirm the delete
// will remove the item from the database via ...
// will capture the id value, and then take in the whole payload
// will copy over the payload objects that do not match the delete

// *** there may be a Delete function

// Delete Item from File handles the method call from the cli
func DeleteItemFromFile()  {
	if ok := isThirdArgument(); ok {
		DeleteItem(getPayload(), os.Args[2:])
	} else {
		DeleteItem(getPayload(), promptUserForDelete())
	}
}

func getPayload() *[]Payload {
	payload, err := ConvertRawDataIntoStruct()
	ExitIfError(err)
	return &payload
}

func DeleteItem(payload *[]Payload, args []string){

	filteredResults := filterPayload(*payload, args)
	if ok := isSingleValue(filteredResults); ok {
		if ok := confirmDelete(filteredResults[0]); ok {
			fmt.Println("demo: deleted item")
			updatedPayload := removeItemFromDataSet(payload, filteredResults[0].Id)
			if ok := writeAllPayloadsToFile(updatedPayload); ok {
				fmt.Println("database updated")
			}
		}
	} else {
		DeleteItem(payload, promptUserForDelete())
	}
}

func removeItemFromDataSet(payload *[]Payload, itemId string) []Payload {
	for i,p := range *payload {
		for _, m := range p {
			if m.Id == itemId {
				pay := *payload
				pay = append(pay[:i], pay[i + 1:]...)
				return pay
			}
		}
	}

	return nil
}

// Delete handles the loading and parsing of the data from the database
func isSingleValue(filteredResults []MemberInfo) bool  {
	if len(filteredResults) != 1  {
		fmt.Printf("results returned: %v", len(filteredResults))
		PrintStatusMessage("Please refine search")
		PrintResults(filteredResults)
		return false
	}
	return true
}

func promptUserForDelete() []string{
	return PromptUser(getSearchQuestions())
}

func confirmDelete(filteredItem MemberInfo) bool  {
	message := fmt.Sprintf("\nPlease confirm to delete %s %s from the database", filteredItem.FirstName, filteredItem.LastName)
	return dangerConfirmPrompt(message, strings.ToLower(filteredItem.FirstName))
}
