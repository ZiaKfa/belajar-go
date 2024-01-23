package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func createFile(fileName string, content string) error {
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer file.Close()
	file.WriteString(content)
	return nil
}

func readFile(fileName string) (string, error) {
	file, err := os.OpenFile(fileName, os.O_RDONLY, 0666)
	if err != nil {
		return "", err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	var content string
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		content += string(line) + "\n"
	}
	return content, nil
}

func appendFile(fileName string, content string) error {
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		return err
	}

	defer file.Close()
	file.WriteString(content)
	return nil
}

func main() {
	//createFile("pokemon.txt", "Pikachu\nBulbasaur\nCharmander\n") //uncomment this line to create the file

	appendFile("pokemon.txt", "Mewtwo\n")

	content, err := readFile("pokemon.txt")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(content)
	}

}
