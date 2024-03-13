package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)
func main() {
	for {
		line := readLine()
		// Remove white spaces and new line on user input
		line = strings.ReplaceAll(line, " ", "")
		line = strings.ReplaceAll(line, "\n", "")

		if line[:len(line)-1] == "close" {
			fmt.Println("Exiting the program.")
			break
		}

		if line[:len(line)-1] == "show" {
			showDirectories()
		} else {
			fmt.Println("Invalid command. Try again.")
		}
	}
}

func readLine() string {
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')

	if err != nil {
			log.Fatal(err)
	}
	return line
}

func showDirectories() {
	files, err := os.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}
	
	for _, f := range files {
		fmt.Println(f.Name())
	}
}

