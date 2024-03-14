package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)
func main() {
	for {
		line := readLine()
		// Remove white spaces and new line on user input
		line = strings.ReplaceAll(line, "\n", "")

		if line[:len(line)-1] == "close" {
			fmt.Println("Exiting the program.")
			break
		} else if line[:len(line)-1] == "back" {
			changeDirectory()
		} else if strings.Fields(line)[0] == "select" && strings.Fields(line)[1] != "" {
			moveDirectory(strings.Fields(line)[1])
		} else if line[:len(line)-1] == "show" {
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

func changeDirectory() {
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory:", err)
		return
	}
	newPath := filepath.Dir(currentDir)
	os.Chdir(newPath)
}

func moveDirectory(path string) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		fmt.Println("Invalid directory:", err)
		return
	}

	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory:", err)
		return
	}

	if fileInfo.IsDir() {
		newPath := filepath.Join(currentDir, path)
		os.Chdir(newPath)
	} else {
		fmt.Println("Invalid directory:", err)
		return
	}
}