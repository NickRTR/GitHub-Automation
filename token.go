package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

var path string = ".gh-automation-config.json"

type Configuration struct {
	GitHub_Token string
}

func GetToken() string {
	t := GetTokenFromConfiguration()
	if len(t) > 0 {
		return t
	}

	// Prompt Token
	reader := bufio.NewReader(os.Stdin)

	brint("Enter your GitHub personal access token: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		brintErr(err.Error())
		os.Exit(1)
	}
	// Remove spaces
	t = strings.TrimSpace(input)

	StoreToken(t)

	return t
}

func GetTokenFromConfiguration() string {
	var config Configuration
	configFile, err := os.Open(path)
	if err != nil {
		return ""
	}

	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&config)
	if err != nil {
		brintErr(fmt.Sprintf("An error ocurred while retrieving the token from configuration: %s\n", err))
		return ""
	}

	return config.GitHub_Token
}

func StoreToken(token string) {
	config := Configuration{
		GitHub_Token: token,
	}

	bytes, err := json.Marshal(config)
	if err != nil {
		brintErr(fmt.Sprintf("An error ocurred while writing the token to configuration: %s\n", err))
	}

	file, err := os.Create(path)
	if err != nil {
		brintErr(fmt.Sprintf("An error ocurred while writing the token to configuration: %s\n", err))
	}
	file.Write(bytes)
	file.Close()
}

func reset() {
	brint("Resetting stored token...")
	err := os.Remove(path)
	if err != nil {
		brintErr(fmt.Sprintf("An error ocurred while resetting the stored token: %s\n", err))
	}
}
