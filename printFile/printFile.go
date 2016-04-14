package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	printFile("main.go")
}

func printFile(path string) {
	file, err := os.Open(path)

	defer file.Close()

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}

}
