package main

// define the type struct for MemberInfo
// define the type struct for Payload
// Method call to build out the database entry
// Method Call to build out the payload



// define the MemberInfo struct
type MemberInfo struct {
	Id string `json:"id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Address string `json:"address"`
	City string `json:"city"`
	State string `json:"state"`
	ZipCode string `json:"zip_code"`
}

func buildMemberInfoStruct(userInput []string, id string) MemberInfo  {
	var a = &MemberInfo{
		Id:id,
		FirstName: userInput[0],
		LastName: userInput[1],
		Address: userInput[2],
		City:userInput[3],
		State:userInput[4],
		ZipCode:userInput[5],
	}
	return *a
}


func buildPayload(mi MemberInfo, uid string) map[string]MemberInfo  {
	p := make(map[string]MemberInfo)
	p[uid] = mi
	return p
}

func newId() string {
	id, err := NewUUID()
	ExitIfError(err)
	return id
}

func ProcessInputData(input []string) map[string]MemberInfo {
	entryId := newId()
	memberStruct := buildMemberInfoStruct(input, entryId)
	return buildPayload(memberStruct, entryId)
}


