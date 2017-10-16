package main

import "fmt"

// handles printing to the console

// print the results of search --- specific for memberInfo
func PrintResults(sliceOfMemberInfo []MemberInfo){


	if len(sliceOfMemberInfo) >= 1 {
		fmt.Println("\n---------------- results ----------------")
		fmt.Println("\nMember(s):")
		for _, result := range sliceOfMemberInfo {
			fmt.Printf("\n%s %s \n%s, \n%s %s %s\n", result.FirstName, result.LastName, result.Address, result.City, result.State, result.ZipCode)
			fmt.Println("\n---------------------------")
		}
		fmt.Println("\n---------------- end ----------------")
	} else {
		fmt.Println("\n---------------- no results!! ----------------")
	}


}
