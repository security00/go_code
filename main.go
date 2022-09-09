package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(3)
	for _, salutation := range []string{"hello", "greetings", "good day"} {
		go func(str string) {
			defer wg.Done()
			fmt.Println(str)
		}(salutation)
	}
	wg.Done()
}
