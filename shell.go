//https://medium.com/@BugDiver/lets-write-a-shell-from-scratch-using-golang-part-i-e528ff6ee6f8
// source ^^
package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	fmt.Print("gosh> ")
	scanner := bufio.NewScanner(os.Stdin)
	for {
		if !scanner.Scan() {
			continue
		}
		command := scanner.Text()
		if strings.TrimSpace(command) == "" {
			continue
		}
		switch command {
		case "pwd":
			wd, _ := os.Getwd()
			fmt.Println(wd)
		case "ls":
			wd, _ := os.Getwd()
			fileInfos, _ := ioutil.ReadDir(wd)
			for _, f := range fileInfos {
				fmt.Println(f.Name())
			}
		case "cd":
			fmt.Println("Action to change dir.")
		case "exit":
			os.Exit(0)
		default:
			fmt.Printf("The command {%s} is not available yet.\n", command)
		}
		fmt.Print("gosh> ")
	}
}
