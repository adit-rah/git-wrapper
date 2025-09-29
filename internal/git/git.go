package git

import (
	"os/exec"
	"strings"
)

func GetCurrentBranch() (string, error) {
	out, err := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD").Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(out)), nil
}

func CreateBranch(branch string) error {
	cmd := exec.Command("git", "checkout", "-b", branch)
	return cmd.Run()
}

func AddAllAndCommit(message string) error {
	if err := exec.Command("git", "add", ".").Run(); err != nil {
		return err
	}
	return exec.Command("git", "commit", "-m", message).Run()
}

func CheckoutBranch(branch string) error {
	return exec.Command("git", "checkout", branch).Run()
}

func PullBranch(branch string) error {
	return exec.Command("git", "pull", "origin", branch).Run()
}

func MergeBranch(branch string) error {
	return exec.Command("git", "merge", "--no-ff", branch).Run()
}

func PushBranch(branch string) error {
	return exec.Command("git", "push", "-u", "origin", branch).Run()
}

func DeleteBranch(branch string) error {
	_ = exec.Command("git", "branch", "-d", branch).Run()
	_ = exec.Command("git", "push", "origin", "--delete", branch).Run()
	return nil
}
