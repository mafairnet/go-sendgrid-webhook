package main

import (
        "log"
        "os"
)

func saveIntoLog(text string){
        f, err := os.OpenFile(configuration.LogFileLocation + "sendgrid.log",os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

        if err != nil {
                log.Println(err)
        }

        defer f.Close()

        if _, err := f.WriteString(text + "\n"); err != nil {
                log.Println(err)
        }
}
