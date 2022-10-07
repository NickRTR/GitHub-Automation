package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err
}

func getInfo() {
	reader := bufio.NewReader(os.Stdin)

	org, _ := getInput("Organization (leave empty to use default): ", reader)
	title, _ := getInput("Repo title (leave empty to use project folder name): ", reader)
	visibility, _ := getInput("Visibility [private / public] (default value is public): ", reader)

	if len(title) == 0 {
		title = getProjectName()
	}

	initRepo(title, org, visibility)
}

func getProjectName() string {
	workingDirectory, err := os.Getwd()

	if err != nil {
		panic(err)
	}

	folders := strings.Split(workingDirectory, "\\")
	return folders[len(folders)-1]
}

func initRepo(title string, org string, visibility string) {
	fmt.Println("initRepo:", title, org, visibility)
}

func main() {
	fmt.Println("GitHub Automation")
	fmt.Println("-----------------")
	getInfo()
}
