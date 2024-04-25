package main

import (
	git "Github/github"
	"fmt"
	"os"
)

func main() {
	username, err := git.GetUserName()
	if err != nil {
		fmt.Println("Usernamni o'qishda xatolik?")
		return
	}

	useremail, ok := git.GetUserEmail()
	if ok != nil {
		fmt.Println("Emailni o'qishda xatolik?")
		return
	}

	file, errr := os.OpenFile("file", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if errr != nil {
		fmt.Println("File ni ochishda xatolik?")
	}

	_, er := file.WriteString(username)
	if er != nil {
		fmt.Println("Filega username ni yozishda xatolik?")
	}

	_, okk := file.WriteString(useremail)
	if okk != nil {
		fmt.Println("Filega useremailni yozishda xatolik?")
	}
}
