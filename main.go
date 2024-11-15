package main

import (
	"fmt"
	"log"
	"os"
	"sync"
)

var wg sync.WaitGroup

func main() {
	osArgs := os.Args
	if len(osArgs) <= 1 {
		fmt.Println(fmt.Errorf("need at lest one argument"))
		return
	}
	args := osArgs[1:]
	for _, value := range args {
		wg.Add(1)
		go createFiles(value)
	}
	wg.Wait()

}

func createFiles(value string) {
	defer wg.Done()
	_, err := os.Create(value)
	if err != nil {
		log.Println(err)
	}

}
