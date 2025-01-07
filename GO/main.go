package main

import "fmt"

func main(){
	// filepath := "./.gitignore"
	cfg, err := ParserGitignore("")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(cfg["API"])
	fmt.Println(cfg["PAI"])

}