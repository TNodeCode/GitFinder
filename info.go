package main

import (
	"os"
	"os/exec"
)

func isGitProject(files []os.FileInfo) bool {
	for _, file := range files {
		if file.IsDir() && file.Name() == ".git" {
			return true
		}
	}

	return false
}

func isNodeJsProject(files []os.FileInfo) bool {
	for _, file := range files {
		if !file.IsDir() && file.Name() == "package.json" {
			return true
		}

		if file.IsDir() && file.Name() == "node_modules" {
			return true
		}
	}

	return false
}

func isGoProject(files []os.FileInfo) bool {
	for _, file := range files {
		if !file.IsDir() && (file.Name() == "main.go" || file.Name() == "go.mod") {
			return true
		}
	}

	return false
}

func isPhpProject(files []os.FileInfo) bool {
	for _, file := range files {
		if file.IsDir() && file.Name() == "vendor" {
			return true
		}

		if !file.IsDir() && file.Name() == "composer.json" {
			return true
		}
	}

	return false
}

func isJavaMavenProject(files []os.FileInfo) bool {
	for _, file := range files {
		if file.IsDir() && file.Name() == "pom.xml" {
			return true
		}
	}

	return false
}

func getGitProjectInformation(dir string) []string {
	cmd := exec.Command("git", "remote", "get-url", "origin")
	cmd.Dir = dir
	out, err := cmd.Output()
	repo := ""

	if err == nil {
		repo = string(out[0 : len(out)-1]) // Remove '\n'
	}

	project := []string{dir, repo, ""}
	return project
}
