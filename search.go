package main

import (
	"io/ioutil"
	"sync"
)

/**
 * Search for git projects recursively
 */
func searchGitProjects(dir string, projectDirs *[][]string, depth int, wg *sync.WaitGroup) {
	defer wg.Done()

	// Get files and directories of dir
	files, err := ioutil.ReadDir(dir)
	isProject := false
	project := []string{"", "", ""}

	// Check for errors during reading
	if err != nil {
		panic(err)
	}

	if isGitProject(files) {
		_project := getGitProjectInformation(dir)
		project[0] = _project[0]
		project[1] = _project[1]
	}

	if isNodeJsProject(files) {
		project[2] = "Node.JS"
	}

	if isGoProject(files) {
		project[2] = "Go"
	}

	if isPhpProject(files) {
		project[2] = "Php"
	}

	if isJavaMavenProject(files) {
		project[2] = "Java - Maven"
	}

	if project[0] != "" {
		*projectDirs = append(*projectDirs, project)
		isProject = true
	}

	// Recursive function call
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
