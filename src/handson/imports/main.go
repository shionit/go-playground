package main

import (
	"fmt"
	greeting "github.com/tenntenn/greeting/v2"
	"time"
)

func main() {
	fmt.Println(greeting.Do(time.Now()))
}
