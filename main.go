package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)
func main() {
	fmt.Println("input command:")

	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')

	if err != nil {
			log.Fatal(err)
	}

	// removed white spaces and new line on user input
	line = strings.ReplaceAll(line, " ", "")
	line = strings.ReplaceAll(line, "\n", "")

	if line[:len(line)-1] == "show" {
		showDirectories()
	}

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
