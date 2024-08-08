package ascii_art

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/* This code reads an ASCII art template from a file, processes it, and generates ASCII art for a given input string based on the template. It includes handling for multi-line input and escaped newline characters.*/
func AsciiArt(textFromOutside string, artstylepath string) (string, error) {
	// 3 textstyles in a folder
	fileLines, Err := ReadStandardTxt(artstylepath)
	if Err != nil {
		return "", Err
	}
	asciiTemplates := return2dASCIIArray(fileLines)
	str, err := printAllStringASCII(textFromOutside, asciiTemplates)
	return str, err
}

func ReadStandardTxt(artstyle string) ([]string, error) {
	readFile, err := os.Open(artstyle)
	if err != nil {
		return nil, err
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	readFile.Close()
	return fileLines, nil
}

func return2dASCIIArray(fileLines []string) [][]string {
	var asciiTemplates [][]string
	counter := 0
	var tempAsciArray []string
	for _, line := range fileLines {
		counter++
		// fmt.Println(index, line)
		if counter != 1 {
			tempAsciArray = append(tempAsciArray, line)
		}
		if counter == 9 {
			// fmt.Println("add to list") // but dont include the first line because it is empty line
			asciiTemplates = append(asciiTemplates, tempAsciArray)
			counter = 0
			tempAsciArray = nil
		} else {
		}
	}
	return asciiTemplates
}

// problem '\n' logic when we have 2 '\n' or 1 '\n' is different
func printMultipleCharacter(s string, asciiTemplates [][]string) string {
	returnString := ""
	tempIntArrLetter := returnAsciiCodeInt(s)
	for i := 0; i < 8; i++ { // Iterate over the 8 lines of ASCII art for each character
		for _, v := range tempIntArrLetter {
			returnString += asciiTemplates[v][i] // Add each line of the ASCII art
		}
		returnString += "\n" // Ensure a line break after each row
	}
	return returnString
}


func returnAsciiCodeInt(s string) []int {
	var tempIntArrLetter []int
	for _, v := range s {
		tempIntArrLetter = append(tempIntArrLetter, (int(v) - 32))
	}
	return tempIntArrLetter
}
func printAllStringASCII(text string, asciiTemplates [][]string) (string, error) {
	for _, ch := range text {
		if ch < 32 || ch > 126 { // Only printable ASCII characters
			if ch != 10 && ch != 13 { // Allow newline and carriage return
				return "", fmt.Errorf("contains unprintable non-ASCII character: %v", ch)
			}
		}
	}

	returnString := ""

	// Split text into lines based on '\n'
	substrings := strings.Split(text, "\n")
	for _, line := range substrings {
		if line == "" {
			returnString += "\n" // Add a blank line for empty input line
		} else {
			returnString += printMultipleCharacter(line, asciiTemplates)
		}
	}

	return returnString, nil
}

func returnstring2EndlineArray(text string) []string {
	substrings := make([]string, 0)
	escapedN := "\\n"
	newline := "\n"

	for {
		idx := strings.Index(text, escapedN)
		if idx == -1 {
			substrings = append(substrings, text)
			break
		}

		substrings = append(substrings, text[:idx])

		if idx+len(escapedN) < len(text) && text[idx+len(escapedN)] == 'n' {
			substrings = append(substrings, newline)
			text = text[idx+len(escapedN)+1:]
		} else {
			substrings = append(substrings, escapedN)
			text = text[idx+len(escapedN):]
		}
	}
	// fmt.Printf("%#v\n", substrings)
	var mysubstring2 []string
	for _, mysub := range substrings {
		if mysub != "" {
			mysubstring2 = append(mysubstring2, mysub)
		}
	}
	// fmt.Printf("%#v\n", mysubstring2)
	return mysubstring2
}
