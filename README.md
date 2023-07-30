# GitHub-Automation

## Installation

(Recent installation of Go(lang) required)

```
go install github.com/NickRTR/GitHub-Automation@latest
```

**On macOS:**

```
sudo ln -s <Path to Executable> /usr/local/bin
```

(Path to executable should be similar to the following path: `Users/name/go/bin/GitHub-Automation`)

## Usage

1. Create a GitHub access token with repo access
1. Navigate into the folder you want to initialize a repository in.
1. Run `GitHub-Automation` with the arguments you wish:
    1. `-title`/`-t`: Change the repo title
    2. `-private`/`-p`: Set repo visibility to private (default ist public)
    3. `-organization`/`-o`: Set organization
    4. `-reset`/`r`: Reset stored GitHub Access Token
