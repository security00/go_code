package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	for _, salutation := range []string{"hello", "greetings", "good day"} {
		wg.Add(1)
		go func(str string) {
			defer wg.Done()
			fmt.Println(str)
		}(salutation)
	}
	wg.Wait()
}
