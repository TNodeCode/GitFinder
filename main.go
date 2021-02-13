package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"sync"
	"time"
)

var (
	flagRootDir = flag.String("dir", ".", "Root directory where to search for Git projects")
	flagDepth   = flag.Int("depth", -1, "Maximum subdirectory depth, -1 for infinity")
	flagCsv     = flag.String("csv", "", "CSV output file name, leave empty if you don't want a CSV file")
)

/**
 * Search for git projects recursively
 */
func searchGitProjects(dir string, projectDirs *[][]string, depth int, wg *sync.WaitGroup) {
	defer wg.Done()

	files, err := ioutil.ReadDir(dir)
	isProject := false

	if err != nil {
		panic(err)
	}

	for _, file := range files {
		if file.IsDir() && file.Name() == ".git" {
			cmd := exec.Command("git", "remote", "get-url", "origin")
			cmd.Dir = dir
			out, err := cmd.Output()
			repo := ""

			if err == nil {
				repo = string(out[0 : len(out)-1]) // Remove '\n'
			}

			project := []string{dir, repo}

			*projectDirs = append(*projectDirs, project)
			isProject = true
		}
	}

	if !isProject && (depth == -1 || depth > 1) {
		for _, file := range files {
			if file.IsDir() && file.Name()[:1] != "_" && file.Name()[:1] != "." {
				if depth != -1 {
					depth -= 1
				}

				wg.Add(1)
				searchGitProjects(dir+"/"+file.Name(), projectDirs, depth, wg)
			}
		}
	}
}

func main() {
	var wg sync.WaitGroup

	flag.Parse()

	ticker := time.NewTicker(5000 * time.Millisecond)
	done := make(chan bool)

	var projectDirs = [][]string{{"Directory", "Origin"}}

	fmt.Printf("Start searching Git projects in %s ...\n", *flagRootDir)

	// Run search in separate Go routin
	wg.Add(1)
	go searchGitProjects(*flagRootDir, &projectDirs, *flagDepth, &wg)

	// Check regulary how many Git projects were found
	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Printf("%v Found %v Git Projects\n", t, len(projectDirs))
			}
		}
	}()

	wg.Wait()

	ticker.Stop()
	done <- true

	if *flagCsv != "" {
		file, err := os.Create(*flagCsv)

		if err != nil {
			log.Fatalf("Error creating file: %s", err)
		}

		defer file.Close()

		writer := csv.NewWriter(file)
		defer writer.Flush()

		for _, project := range projectDirs {
			err = writer.Write(project)
			if err != nil {
				log.Fatalf("Error writing file: %s", err)
			}
		}

		fmt.Printf("Result written to %s\n", *flagCsv)
	}

	for i, project := range projectDirs {
		fmt.Println(i+1, project)
	}
}
