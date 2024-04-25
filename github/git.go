package git

import (
	"os/exec"
)

func GetUserName() (string, error) {
	name := exec.Command("git", "config", "--get", "user.name")
	s, err := name.Output()
	if err != nil {
		return "", err
	}
	return string(s), nil
}

func GetUserEmail() (string, error) {
	email := exec.Command("git", "config", "--get", "user.email")
	d, er := email.Output()
	if er != nil {
		return "", er
	}
	return string(d), nil
}
