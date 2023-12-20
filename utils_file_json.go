package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func saveWebHookFile(account string, data string, timestamp int, eventType string) {

	date := timestampIntegerToTime(timestamp)
	
	
	year := strconv.Itoa(date.Year())
	month := strconv.Itoa(int(date.Month()))
	day := strconv.Itoa(date.Day())

	//fileLocation := configuration.LogFileLocation + "/data/" + year + "/" + month + "/" + day + "/"
	
	fileLocation := configuration.DataStorageLocation + year + "/" + month + "/" + day + "/"
	
	fmt.Printf("fileLocation: %v\n", fileLocation)

	if _, err := os.Stat(fileLocation); os.IsNotExist(err) {
		//os.Mkdir(fileLocation, 0644)
		os.MkdirAll(fileLocation, os.ModePerm)
	}

	/*if err != nil {
		fmt.Printf("err: %v\n", err)
	}*/

	finalName := fileLocation + account + "-" +eventType + "-" + strconv.Itoa(timestamp) + "-" + generateRandomString(4) + ".json"

	_ = ioutil.WriteFile(finalName, []byte(data), 0644)

	if configuration.AwsEnabled {
		uploadFile(finalName)
	}

}

