package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func gitac(commit string) {
	gitAdd := exec.Command("git", "add", ".")
	err := gitAdd.Run()
	if err != nil {
		log.Fatal(err)
	}
	gitCommit := exec.Command("git", "commit", "-m", commit)
	commitErr := gitCommit.Run()
	if commitErr != nil {
		log.Fatal(commitErr)
	}
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Files Committed\nProceed to push files?\n(y/n): ")
	response, _ := reader.ReadString('\n')
	response = strings.TrimSpace(response)
	if response == "y" {

		gitP := exec.Command("git", "push", "-u", "origin", "main")
		gitPErr := gitP.Run()
		if gitPErr != nil {
			log.Fatal(gitPErr)
		}
		fmt.Println("Files pushed to main")

	}

}
