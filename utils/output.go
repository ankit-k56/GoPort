package utils

import (
	"encoding/json"
	"os"

	"github.com/ankit-k56/GoPort/scanport"
)

func GenerateOutput(Output scanport.Output)  {
	folderPath := "./Output"

	os.Mkdir(folderPath, os.ModePerm)

	file, err := os.Create(folderPath+ "/output.json")
	if err != nil{
		panic(err)
	}
	jsonOutput, err := json.MarshalIndent(Output, "", " ")
	if err != nil{
		panic(err)
	
	}
	file.Write(jsonOutput)

	

	defer file.Close()
}