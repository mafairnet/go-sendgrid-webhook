package main

import (
    "io"
    //"log"
    "net/http"
    "fmt"

    "github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var configuration = getProgramConfiguration()

func main() {

    r := gin.Default()

	r.POST("/webhook-test", func(c *gin.Context) {
        headers := c.Request.Header
        payload, err := io.ReadAll(c.Request.Body)

        fmt.Printf("\nHeaders: %v",headers)

        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            fmt.Printf("\nError: %v", err.Error())
            return
        }

        // Do something with the message...
        payloadString := string(payload)
		processWebHookEvent("test",payloadString)

        c.JSON(200, gin.H{})
    })

    r.POST("/webhook-pt-account-privaterelay", func(c *gin.Context) {
        headers := c.Request.Header
        payload, err := io.ReadAll(c.Request.Body)

        fmt.Printf("\nHeaders: %v",headers)

        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            fmt.Printf("\nError: %v", err.Error())
            return
        }

        // Do something with the message...
        payloadString := string(payload)
		processWebHookEvent("pt-account-privaterelay",payloadString)

        c.JSON(200, gin.H{})
    })

	r.POST("/webhook-pt-account-em", func(c *gin.Context) {
        headers := c.Request.Header
        payload, err := io.ReadAll(c.Request.Body)

        fmt.Printf("\nHeaders: %v",headers)

        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            fmt.Printf("\nError: %v", err.Error())
            return
        }

        // Do something with the message...
        payloadString := string(payload)
		processWebHookEvent("pt-account-em",payloadString)

        c.JSON(200, gin.H{})
    })

	r.POST("/webhook-pt-account-emtb", func(c *gin.Context) {
        headers := c.Request.Header
        payload, err := io.ReadAll(c.Request.Body)

        fmt.Printf("\nHeaders: %v",headers)

        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            fmt.Printf("\nError: %v", err.Error())
            return
        }

        // Do something with the message...
        payloadString := string(payload)
		processWebHookEvent("pt-account-emtb",payloadString)

        c.JSON(200, gin.H{})
    })

	r.POST("/webhook-pt-account-trans", func(c *gin.Context) {
        headers := c.Request.Header
        payload, err := io.ReadAll(c.Request.Body)

        fmt.Printf("\nHeaders: %v",headers)

        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            fmt.Printf("\nError: %v", err.Error())
            return
        }

        // Do something with the message...
        payloadString := string(payload)
		processWebHookEvent("pt-account-trans",payloadString)

        c.JSON(200, gin.H{})
    })

	r.POST("/webhook-pt-account-ptrv", func(c *gin.Context) {
        headers := c.Request.Header
        payload, err := io.ReadAll(c.Request.Body)

        fmt.Printf("\nHeaders: %v",headers)

        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            fmt.Printf("\nError: %v", err.Error())
            return
        }

        // Do something with the message...
        payloadString := string(payload)
		processWebHookEvent("pt-account-ptrv",payloadString)

        c.JSON(200, gin.H{})
    })

    r.Run("0.0.0.0:"+configuration.Port)
}
