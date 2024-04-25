package main

import (
	git "Github/github"
	"fmt"
	"os"
)

func main() {
	username, err := git.GetUserName()
	if err != nil {
		fmt.Println("Error getting username:", err)
		return
	}
	file, errr := os.OpenFile("file", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if errr != nil {
		fmt.Println("Fileda xatolik?")
	}
	_, er := file.WriteString(username)
	if er != nil {
		fmt.Println("Xatolik?")
	}
	fmt.Println("Git username is:", username)
}
