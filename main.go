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
	displayDirectory()
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
			selectDirectory(strings.Fields(line)[1])
		} else if strings.Fields(line)[0] == "create" && strings.Fields(line)[1] != "" {
			createDirectory(strings.Fields(line)[1])
		} else if strings.Fields(line)[0] == "delete" && strings.Fields(line)[1] != "" {
			deleteDirectory(strings.Fields(line)[1])
		} else if strings.Fields(line)[0] == "update" && strings.Fields(line)[1] != "" && strings.Fields(line)[2] != "" {
			updateDirectoryname(strings.Fields(line)[1], strings.Fields(line)[2])
		}  else if strings.Fields(line)[0] == "move" && strings.Fields(line)[1] != "" && strings.Fields(line)[2] != "" {
			moveDirectory(strings.Fields(line)[1], strings.Fields(line)[2])
		}  else if line[:len(line)-1] == "show" {
			showDirectories()
		} else {
			fmt.Println("Invalid command. Try again.")
		}
		displayDirectory()
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

func selectDirectory(path string) {
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

func createDirectory(name string) {
	// check if has extension if yes create that file else directory
	err := os.Mkdir(name, 0750)
	if err != nil && !os.IsExist(err) {
		log.Fatal(err)
	}
}

func deleteDirectory(directory string) {
	// add checking if file or directory
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory:", err)
		return
	}

	newPath := filepath.Join(currentDir, directory)
	if err := os.RemoveAll(newPath); err != nil {
		fmt.Println("Invalid directory:", err)
		return
	}
}

func updateDirectoryname(oldpath string, newPath string) {
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory:", err)
		return
	}

	updatedDirectory := filepath.Join(currentDir, newPath)
	oldDirectory := filepath.Join(currentDir, oldpath)
	 
	if err := os.Rename(oldDirectory, updatedDirectory); err != nil { 
			log.Fatal(err) 
	} 
}

func moveDirectory(name string, path string) { 
	// Get the absolute path of the current working directory
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory:", err)
		return
	}

	// Resolve the destination path to an absolute path
	selectedDirectory := filepath.Join(currentDir, name)
	updatedDirectory := filepath.Join(currentDir, path)
	
	// Move the directory
	if err = os.Rename(selectedDirectory, updatedDirectory); err != nil {
		fmt.Println("Error moving current directory:", err)
		return 
	}

	fmt.Printf("Moved directory %s to %s\n", name, updatedDirectory)
}

func displayDirectory() {
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory:", err)
		return
	}
	fmt.Println(currentDir)
}