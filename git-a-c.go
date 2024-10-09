package main

import (
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Commit message required")
	}
	gitAdd := exec.Command("git", "add", ".")
	err := gitAdd.Run()
	if err != nil {
		log.Fatal(err)
	}
	commit := strings.Join(os.Args[1:], " ")
	gitCommit := exec.Command("git", "commit", "-m", commit)
	commitErr := gitCommit.Run()
	if commitErr != nil {
		log.Fatal(commitErr)
	}

	gitP := exec.Command("git", "push", "-u", "origin", "main")

	gitPErr := gitP.Run()
	if gitPErr != nil {
		log.Fatal(gitPErr)
	}

}
