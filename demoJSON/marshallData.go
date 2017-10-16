package main

import (
	"io/ioutil"
	"encoding/json"
)



// method call to write the json to the file
func WriteToJSON(data Payload) error  {
	fileData, err := loadData("../demoData1.json")
	ExitIfError(err)

	var compositeData []Payload

	err = json.Unmarshal(*fileData, &compositeData)
	ExitIfError(err)
	
	compositeData = append(compositeData, data)

	exportData, err := json.MarshalIndent(compositeData, "", " ")
	ExitIfError(err)

	err = writeData("../demoData1.json", exportData)
	ExitIfError(err)

	return nil
}

// method call to read the entire JSON object from the file
func ReadFromJSON(filePath string) ([]byte, error)  {
	raw, err := loadData(filePath)
	if err != nil {
		return nil, err
	}
	return *raw, nil
}

// unmarshal data into data structs and return them 
func ConvertRawDataIntoStruct()([]Payload, error)  {

	// load raw data
	raw, err := loadData("../demoData1.json")
	ExitIfError(err)

	var p []Payload
	if err := json.Unmarshal(*raw, &p); err != nil {
		return nil, err
	}
	return p, nil
}

// load the data from the file
func loadData(filePath string) (*[]byte, error)  {
	f, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	return &f, nil
}

// write data to the file
func writeData(fileName string, data []byte) error  {
	if err := ioutil.WriteFile(fileName,data, 0644); err != nil {
		return err
	}
	return nil
}

