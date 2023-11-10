package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func cloneGitHubRepo(repoURL, localPath string) error {
	// Ensure localPath ends with the repository name
	if len(localPath) == 0 || localPath[len(localPath)-1] != '/' {
		localPath += "/"
	}
	localPath += repoURL[strings.LastIndex(repoURL, "/")+1 : strings.LastIndex(repoURL, ".")]

	cmd := exec.Command("git", "clone", repoURL, localPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("error cloning repository: %v", err)
	}

	fmt.Printf("Repository cloned successfully to %s\n", localPath)
	return nil
}

func main() {
	// Replace 'your_github_repo_url' with the actual GitHub repository URL
	githubRepoURL := "https://github.com/emea12/lol.git"
	// Specify 'C:\' as the local path
	localClonePath := "C:/"

	err := cloneGitHubRepo(githubRepoURL, localClonePath)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
