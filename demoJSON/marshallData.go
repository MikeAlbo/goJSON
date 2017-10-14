package main

import (
	"io/ioutil"
	"encoding/json"
)

// define the struct
type MemberInfo struct {
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Address string `json:"address"`
	City string `json:"city"`
	State string `json:"state"`
	ZipCode string `json:"zip_code"`
}

func BuildStruct(input []string) MemberInfo  {
	var a = &MemberInfo{
		FirstName: input[0],
		LastName: input[1],
		Address: input[2],
		City:input[3],
		State:input[4],
		ZipCode:input[5],
	}

	return *a
}

func WriteToJSON(data MemberInfo) error  {
	fileData, err := loadData("../demoData1.json")
	if err != nil {
		return err
	}

	var compositeData []MemberInfo

	if err = json.Unmarshal(*fileData, &compositeData); err != nil {
		return err
	}
	
	compositeData = append(compositeData, data)

	exportData, err := json.MarshalIndent(compositeData, "", " ")
	if err != nil {
		return err
	}

	if err = writeData("../demoData1.json", exportData); err != nil {
		return err
	}

	return nil
}

func ReadFromJSON(filePath string) ([]byte, error)  {
	raw, err := loadData(filePath)
	if err != nil {
		return nil, err
	}
	return *raw, nil
}

// load the data from the file
func loadData(filePath string) (*[]byte, error)  {
	f, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	return &f, nil
}

func writeData(fileName string, data []byte) error  {
	if err := ioutil.WriteFile(fileName,data, 0644); err != nil {
		return err
	}
	return nil
}

