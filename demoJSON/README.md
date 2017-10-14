# JSON with GO

## Goals
1. demo marshalling of data from structs to JSON
2. demo un-marhsalling data from JSON to structs
3. demo filtering/ searching for data within the JSON 


## setup

v0.1

change directory into the demoJSON folder inside the src folder

`cd demoJSON `

call the GO install command, this will compile all of the files (all are under package main)

`go install`

## run the application

v 0.1

cd into the bin folder *from the demoJSON folder*

`cd ../../bin`

run the application with a flag

1. `demoJSON -write` : will setup the application to write to the JSON file
2. `demoJSON -read` : will setup the application to read the JSON file

