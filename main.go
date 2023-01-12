package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/google/go-github/v47/github"
	"golang.org/x/oauth2"
)

func getDirectoryName() string {
	workingDirectory, err := os.Getwd()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return filepath.Base(workingDirectory)
}

func token() string {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter your GitHub personal access token: ")
	input, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Remove spaces
	return strings.TrimSpace(input)
}

func authenticate() (*github.Client, context.Context) {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token()},
	)
	tc := oauth2.NewClient(ctx, ts)
	return github.NewClient(tc), ctx
}

func createRepo(client *github.Client, ctx context.Context) string {
	fmt.Println("Creating Repository...")
	repo := &github.Repository{
		Name:    github.String(title),
		Private: github.Bool(private),
	}
	res, _, err := client.Repositories.Create(ctx, organization, repo)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Successfully created Repository!")

	return *res.CloneURL
}

func execute(command string) {
	cmd := exec.Command("bash", "-c", command)

	_, err := cmd.Output()
	if err != nil {
		fmt.Println("Could not run command: ", err)
	}
}

func createREADME() {
	fmt.Println("Create README")

	err := os.WriteFile("README.md", []byte("# "+title), os.ModeAppend)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Successfully created README")
}

func initRepo(url string) {
	fmt.Println("Initializing Repository...")

	// Check for README
	_, err := os.Stat("README.md")
	if err != nil && os.IsNotExist(err) {
		createREADME()
	}

	execute("git init")
	execute("git add .")
	execute("git commit -m \"initial commit\"")
	execute("git branch -M main")
	execute("git remote add origin " + url)
	execute("git push -u origin main")

	fmt.Println("Successfully Initialized Repository.")
}

// flags
var title string
var organization string
var private bool

func init() {
	var (
		defaultTitle     = getDirectoryName()
		titleDescription = "Change the repository title, the default is your current directory"
	)
	flag.StringVar(&title, "title", defaultTitle, titleDescription)
	flag.StringVar(&title, "t", title, titleDescription+" (shorthand)")

	var (
		defaultOrganization     = ""
		organizationDescription = "Change the organization of the repository (must be existent and accessible through your token)"
	)
	flag.StringVar(&organization, "organization", defaultOrganization, organizationDescription)
	flag.StringVar(&organization, "o", defaultOrganization, organizationDescription+" (shorthand)")

	const (
		privateDescription = "set the visibility of the repository to private"
	)
	flag.BoolVar(&private, "private", false, privateDescription)
	flag.BoolVar(&private, "p", false, privateDescription+" (shorthand)")

	flag.Parse()
}

func main() {
	fmt.Println("GitHub-Automation")
	fmt.Println("--------------------------------------")

	repoURL := createRepo(authenticate())

	initRepo(repoURL)

	fmt.Printf("\nVisit the initialized Repository at %s\n", repoURL)
}
