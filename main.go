package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func getDirectoryName() string {
	workingDirectory, err := os.Getwd()

	if err != nil {
		panic(err)
	}

	return filepath.Base(workingDirectory)
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
	fmt.Println(title)
	fmt.Println(organization)
	fmt.Println(private)
}
