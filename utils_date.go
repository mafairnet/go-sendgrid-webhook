package main

import (
    "fmt"
    "time"
    "strconv"
)

func timestampToTime(timestamp string) time.Time{
	i, err := strconv.ParseInt(timestamp, 10, 64)
    if err != nil {
        panic(err)
    }
    tm := time.Unix(i, 0)
    fmt.Println(tm)
	return tm
}

func timestampIntegerToTime(timestamp int) time.Time{
    tm := time.Unix(int64(timestamp), 0)
    fmt.Println(tm)
	return tm
}