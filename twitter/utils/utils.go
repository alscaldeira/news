package utils

import (
	"fmt"
	"time"
)

func Log(message string) {
	t := "[ LOG ]"
	var now = time.Now().UTC().Format("2006-01-02 15:04:05")
	fmt.Println(now, t, message)
}

func Logeadline(message string) {
	t := "[ DEADLINE ]"
	var now = time.Now().UTC().Format("2006-01-02 15:04:05")
	fmt.Println(now, t, message)
}
