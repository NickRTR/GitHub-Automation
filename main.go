package main

import (
	"bufio"
	"fmt"
	"http"
	"net/http"
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

	if visibility == "private" {
		payload := fmt.Sprintf("\"name\": \"Test\", \"private\": \"true\"")
		fmt.Println(payload)
    	// payload := '{"name": "' + title + '", "private": "true"}'
	} else {
		// payload := '{"name": "' + name + '", "private": "false"}'
	}

	fmt.Println("creating repository...")

	API_URL := "https://api.github.com"
	headers := {
		"Authorization": "token " + token(),
		"Accept": "application/vnd.github.v3+json",
	}

	if org == "" {
		res, err := http.PostForm(API_URL + "/user/repos", data=payload, headers=headers)

		if err != nil {
			panic(err)
		}

		var res map[string]interface{}

    json.NewDecoder(resp.Body).Decode(&res)

    fmt.Println(res["form"])
	}
}

func token() string {
	// TODO: ENV variable
}

func main() {
	fmt.Println("GitHub Automation")
	fmt.Println("-----------------")
	getInfo()
}
