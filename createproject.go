package main

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const prefix = "/home/moon/dev/"

func main() {
	args := os.Args

	if len(args) < 2 {
		log.Fatal("You must provide a project name")
	}

	name := args[1]

	projectsPath := filepath.Join(getExecutablePath(), "projects.json")
	data, err := os.ReadFile(projectsPath)

	if err != nil {
		log.Fatal(err.Error())
	}

	var projects map[string]string
	err = json.Unmarshal(data, &projects)

	if err != nil {
		log.Fatal(err.Error())
	}

	projects[name] = prefix + name

	if _, err := os.Stat("./conf/app.ini"); err != nil {
	    if os.IsNotExist(err) {
        	err = os.MkdirAll(prefix+name, 0o755)

		if err != nil {
			log.Fatal(err.Error())
		}
    	    } else {
	        // other error
    	    }
	}

	data, err = json.Marshal(projects)

	if err != nil {
		log.Fatal(err.Error())
	}

	err = os.WriteFile(projectsPath, data, 0o644)

	if err != nil {
		log.Fatal(err.Error())
	}
}

func getExecutablePath() string {
	executable, _ := os.Executable()
	path := strings.Split(executable, "/")
	directories := path[:len(path)-1]
	return strings.Join(directories, "/")
}
