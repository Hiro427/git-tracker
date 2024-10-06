package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/tabwriter"

	"github.com/fatih/color"
	"github.com/go-git/go-git/v5"
	"github.com/pkg/browser"
	"gopkg.in/ini.v1"
)

func gitStatus(path string) bool {
	// Open the Git repository
	repo, err := git.PlainOpen(path)
	if err != nil {
		fmt.Println("Error opening repository:", err)
		return false
	}

	// Get the working tree
	wt, err := repo.Worktree()
	if err != nil {
		fmt.Println("Error getting working tree:", err)
		return false
	}

	// Get the status of the files
	status, err := wt.Status()
	if err != nil {
		fmt.Println("Error getting status:", err)
		return false
	}

	// Declare the status flag before the loop
	var stat bool

	// Loop through the files and check if any of them meet the condition
	for _, fileStatus := range status {
		// If any file is modified (staged or unstaged), untracked, or staged for commit
		if fileStatus.Worktree == git.Modified || // unstaged modifications
			fileStatus.Worktree == git.Untracked || // untracked files
			fileStatus.Staging != git.Unmodified && fileStatus.Staging != git.Untracked { // staged files

			// Set stat to true if any of these are true
			stat = true
			break // Exit the loop early, since we found a matching condition
		}
	}

	// Return the result after checking all files
	return stat
}

func readRepoPaths(filePath string) ([]string, error) {
	// Open the file

	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)
	var paths []string

	// Loop through each line
	for scanner.Scan() {
		// Read the current line
		line := strings.TrimSpace(scanner.Text()) // Trim any extra whitespace
		if line != "" {                           // Ignore empty lines
			// prefix, _ := os.UserHomeDir()
			// gitRepo := prefix + line
			paths = append(paths, line) // Add the line (path) to the slice
		}
	}

	// Check for any scanning errors
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	return paths, nil
}

func appendCwdToFile(filePath string) error {
	// Get the current working directory (CWD)
	cwd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get current directory: %w", err)
	}

	// Open the repos.txt file in append mode (or create it if it doesn't exist)
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to open repos.txt file: %w", err)
	}
	defer file.Close()

	// Append the CWD to repos.txt with a newline
	if _, err := file.WriteString(cwd + "\n"); err != nil {
		return fmt.Errorf("failed to write CWD to repos.txt: %w", err)
	}

	fmt.Println("Successfully appended current directory to repos.txt")
	return nil
}
func printTable() {

	// Path to the .txt file that contains repository paths
	repoFilePath := "/home/jacobrambarran/.dotfiles/github/repos.txt"

	// Read the repository paths from the file
	repoPaths, err := readRepoPaths(repoFilePath)
	if err != nil {
		fmt.Println("Error reading repo paths:", err)
		return
	}

	red := color.New(color.FgRed).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()
	blue := color.New(color.FgBlue).SprintFunc()
	yellow := color.New(color.FgYellow).SprintFunc()
	underline := color.New(color.Underline).SprintFunc()

	writer := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	heading := "Repository\tStatus"
	fmt.Fprintln(writer, underline(blue(heading)))
	// Process each repository path
	for _, repoPath := range repoPaths {
		projectName := filepath.Base(repoPath)

		var tableStat string

		status := gitStatus(repoPath)
		if status {
			tableStat = red("Dirty")
		} else {
			tableStat = green("Clean")
		}

		fmt.Fprintf(writer, "%s\t%s\t\n", yellow(projectName), tableStat)

	}

	writer.Flush()
}

func openRepo() {

	gitDir := filepath.Join(".git", "config")
	cfg, _ := ini.Load(gitDir)

	remoteOrigin := cfg.Section(`remote "origin"`)
	originURL := remoteOrigin.Key("url").String()
	err := browser.OpenURL(originURL)
	if err != nil {
		fmt.Println("Failed to Open URL:", err)

	}

}

func main() {

	appendCwdFlag := flag.Bool("append-cwd", false, "Append the current working directory to repos.txt")
	listReposFlag := flag.Bool("list-repos", false, "List all repositories in repos.txt")
	openGit := flag.Bool("open-repo", false, "Open Repo in Current Directory on github.com")

	flag.Parse()

	if *appendCwdFlag {
		appendCwdToFile("/home/jacobrambarran/.dotfiles/github/repos.txt")
	}

	if *listReposFlag {
		printTable()

	}
	if *openGit {
		openRepo()
	}

}
