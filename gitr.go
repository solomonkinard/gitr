package main

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func cwd() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return dir
}

func main() {
	clargs := os.Args[1:]
	curr := cwd()
	for i := len(clargs) - 1; i >= 0; i-- {
		if strings.Contains(clargs[i], "/") {
			if strings.Compare(string(clargs[i][0]), "/") == 0 || strings.Compare(string(clargs[i][0]), "~") == 0 {
				curr = clargs[i]
				println(0)
			} else {
				curr = curr + "/" + clargs[i]
				clargs[i] = curr
			}
			if info, err := os.Stat(curr); err == nil && info.IsDir() {
				break
			} else {
				curr = filepath.Dir(curr)
				break
			}
		}
	}

	for ok := true; ok; ok = true {
		if curr == "/" {
			break
		}
		if _, err := os.Stat(curr + "/.git"); os.IsNotExist(err) {
		} else {
			break
		}
		curr = filepath.Dir(curr)
	}

	command := []string{"git"}
	pieces := append(command, clargs...)
	name := pieces[0]
	args := pieces[1:]
	cmd := exec.Command(name, args...)
	cmd.Dir = curr
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}
