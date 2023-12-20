package main

import (
        "fmt"
	"strings"
	"os/exec"
)


func uploadFile(filename string){

	//aws s3 cp c:\sync\logs\log1.xml s3://atasync1/
	//s3://sendgrid-event-webhooks/2023/3/
	//aws s3 cp configuration.StorageLocation+"/data/"+filename s3://sendgrid-event-webhooks/+filename

	prg := "/usr/bin/aws"

	arg1 := "s3"
    arg2 := "cp"
    arg3 := filename
	
	arg5 := "/usr/bin/rm"
	arg6 := "-f"

	fileNameForS3 := strings.Replace(filename, configuration.LogFileLocation+"/data/", "", -1)

	arg4 := "s3://sendgrid-event-webhooks/" + fileNameForS3

	fmt.Printf("Comand: %v\n", prg + " " + arg1 + " " + arg2 + " " + arg3 + " " + arg4)

	//Copiamos el archivo
	cmd := exec.Command(prg, arg1, arg2, arg3,arg4)
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Print(string(stdout))

	//Borramos el archivo
	cmd = exec.Command(prg, arg5, arg6, filename)
	stdout, err = cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Print(string(stdout))


}
