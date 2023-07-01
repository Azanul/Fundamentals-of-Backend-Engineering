package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	fmt.Println("1")
	wg.Add(1)
	go func() {
		defer wg.Done()
		absFilepath, err := filepath.Abs("test.txt")
		if err != nil {
			fmt.Println(err)
			return
		}
		content, err := os.ReadFile(absFilepath)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(content))
	}()
	fmt.Println("2")
	wg.Wait()
}
