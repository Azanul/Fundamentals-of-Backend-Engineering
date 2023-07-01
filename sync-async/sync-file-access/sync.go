package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	fmt.Println("1")
	absFilepath, err := filepath.Abs("test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	content, err := os.ReadFile(absFilepath)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(content))
	fmt.Println("2")
}
