package main

import "fmt"

// handles printing to the console

// print the results of search --- specific for memberInfo
func PrintResults(sliceOfMemberInfo []MemberInfo){
	fmt.Println("\n---------------- results ----------------")
	fmt.Println("\nMember(s):")
	for _, result := range sliceOfMemberInfo {
		fmt.Printf("%s %s \n Address: %s, \n %s %s %s\n", result.FirstName, result.LastName, result.Address, result.City, result.State, result.ZipCode)
	}
}
