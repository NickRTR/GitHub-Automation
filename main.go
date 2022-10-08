package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/google/go-github/v47/github"
	"golang.org/x/oauth2"
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
	fmt.Println("Authenticating...")
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: "... your access token ..."},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	// list all repositories for the authenticated user
	repos, _, err := client.Repositories.List(ctx, "", nil)
	fmt.Println(repos, err)

	fmt.Println("creating repository...")

	// repo := &github.Repository{
	// 	Name:    github.String(title),
	// 	Private: github.Bool(visibility == "private"),
	// 	// Org:	 github.String(org)
	// }
	// client.Repositories.Create(ctx, "", repo)
}

// func token() string {
// 	// TODO: ENV variable
// }

func main() {
	fmt.Println("GitHub Automation")
	fmt.Println("-----------------")
	getInfo()
}
