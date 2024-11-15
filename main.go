package main

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	if len(os.Args) <= 0 {
		log.Fatal("need at lest one argument")
		return
	}
	args := os.Args[1:]
	t1 := time.Now()
	for _, value := range args {
		wg.Add(1)
		go createFiles(value)
	}
	wg.Wait()
	fmt.Println("completed in ", time.Since(t1))

}

func createFiles(value string) {
	defer wg.Done()
	// filename := fmt.Sprintf("type nul > %s", value)
	_, err := os.Create(value)
	// cmd := exec.Command("cmd", "/C", filename)
	if err != nil {
		log.Println(err)
	}

}
