package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"syscall"

	"github.com/google/go-github/v47/github"
	"golang.org/x/oauth2"
)

func getDirectoryName() string {
	workingDirectory, err := os.Getwd()

	if err != nil {
		brintErr(err.Error())
		os.Exit(1)
	}

	return filepath.Base(workingDirectory)
}

func authenticate() (*github.Client, context.Context) {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: GetToken()},
	)
	tc := oauth2.NewClient(ctx, ts)
	return github.NewClient(tc), ctx
}

func formatUmlauts(text string) string {
	// Replace ä with ae, ö with oe, ü with ue, ß with ss
	text = strings.ReplaceAll(text, "ä", "ae")
	text = strings.ReplaceAll(text, "ö", "oe")
	text = strings.ReplaceAll(text, "ü", "ue")
	text = strings.ReplaceAll(text, "ß", "ss")

	return text
}

func createRepo(client *github.Client, ctx context.Context) string {
	brint("Creating Repository...")
	repo := &github.Repository{
		Name:    github.String(formatUmlauts(title)),
		Private: github.Bool(private),
	}
	res, _, err := client.Repositories.Create(ctx, organization, repo)

	if err != nil {
		brintErr(err.Error())
		os.Exit(1)
	}

	brint("Successfully created Repository!")

	return *res.CloneURL
}

func execute(command string) error {
	splitted := strings.Split(command, "#")
	name := splitted[0]
	splitted = splitted[1:]

	cmd := exec.Command(name, splitted...)

	_, err := cmd.Output()
	if err != nil {
		brintErr(fmt.Sprintf("Could not run command (%s) %s", command, err.Error()))
	}

	return err
}

func createREADME() {
	brint("Create README")

	err := os.WriteFile("README.md", []byte("# "+title), os.ModeAppend)
	if err != nil {
		brintErr(err.Error())
		os.Exit(1)
	}

	// Set file mode explicitly to 0644
	err = syscall.Chmod("README.md", 0644)
	if err != nil {
		brintErr(err.Error())
		os.Exit(1)
	}

	brint("Successfully created README")
}

func initRepo(url string) {
	brint("Initializing Repository...")

	// Check for README
	_, err := os.Stat("README.md")
	if err != nil && os.IsNotExist(err) {
		createREADME()
	}

	commands := []string{"git#init", "git#add#.", "git#commit#-m#initial commit", "git#branch#-M#main", "git#remote#add#origin#" + url, "git#push#-u#origin#main"}

	for _, c := range commands {
		err := execute(c)
		if err != nil {
			brintErr("Failed initializing Repository")
			return
		}
	}

	brint("Successfully initialized Repository.")

}

// flags
var title string
var organization string
var private bool
var resetToken bool

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

	const (
		tokenDescription = "Reset the saved GitHub Access Token"
	)
	flag.BoolVar(&resetToken, "reset", false, tokenDescription)
	flag.BoolVar(&resetToken, "r", false, tokenDescription)

	flag.Parse()
}

func main() {
	if resetToken {
		reset()
	}

	repoURL := createRepo(authenticate())
	initRepo(repoURL)

	fmt.Printf("🟡 Visit the Repository at %s\n", repoURL)
}
