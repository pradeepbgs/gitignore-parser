package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)


func ParserGitignore(filepath string) (map[string]string,error) {
	if filepath == ""{
		filepath  = ".gitignore"
	}

	file,err := os.Open(filepath)
	if err != nil{
		return nil, fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	config := make(map[string] string)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		
		if strings.TrimSpace(line) == "" || strings.HasPrefix(line,"#"){
			continue
		}

		parts := strings.Split(line, "=")
		if len(parts) == 2{
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])

			config[key] = value
		}
	}

	if err := scanner.Err(); err != nil{
		return nil, fmt.Errorf("error reading file: %v", err)	
	}

	return config,nil
}