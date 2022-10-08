package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/google/go-github/v47/github"
	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
)

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err
}

type info struct {
	org        string
	title      string
	visibility string
}

func getInfo() *info {
	reader := bufio.NewReader(os.Stdin)

	org, _ := getInput("Organization (leave empty to use default): ", reader)
	title, _ := getInput("Repo title (leave empty to use project folder name): ", reader)
	visibility, _ := getInput("Visibility [private / public] (default value is public): ", reader)

	if len(title) == 0 {
		title = getProjectName()
	}

	i := info{
		title:      title,
		org:        org,
		visibility: visibility,
	}

	return &i
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
		&oauth2.Token{AccessToken: token()},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	fmt.Println("creating repository...")
	repo := &github.Repository{
		Name:    github.String(title),
		Private: github.Bool(visibility == "private"),
	}
	client.Repositories.Create(ctx, org, repo)

	fmt.Println("Successfully initialized Repository")
}

func token() string {
	error := godotenv.Load(".env")
	if error != nil {
		fmt.Println("Could not load .env file")
		os.Exit(1)
	}

	token := os.Getenv("GH_TOKEN")

	if len(token) <= 0 {
		fmt.Println("Please insert your GitHub Token into the prepared .env file.")
		os.Exit(1)
	}

	return token
}

func main() {
	fmt.Println("GitHub Automation")
	fmt.Println("-----------------")
	info := getInfo()
	initRepo(info.title, info.org, info.visibility)
}
