package git

import "os/exec"

func GetUserName() (string, error) {
	cmd := exec.Command("git", "config", "--get", "user.name")
	s, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(s),nil
}
