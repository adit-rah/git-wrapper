package git

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
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

// ReadBranchBase reads the .branch_base file if it exists
func ReadBranchBase() (string, error) {
	data, err := os.ReadFile(".branch_base")
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(data)), nil
}

// CountCommitsSinceBase returns number of commits on current branch since base
func CountCommitsSinceBase(base, current string) (int, error) {
	out, err := exec.Command("git", "rev-list", fmt.Sprintf("%s..%s", base, current), "--count").Output()
	if err != nil {
		return 0, err
	}
	countStr := strings.TrimSpace(string(out))
	count, err := strconv.Atoi(countStr)
	if err != nil {
		return 0, err
	}
	return count, nil
}

// UncommittedChanges returns a slice of uncommitted file paths
func UncommittedChanges() ([]string, error) {
	out, err := exec.Command("git", "status", "--porcelain").Output()
	if err != nil {
		return nil, err
	}
	lines := strings.Split(strings.TrimSpace(string(out)), "\n")
	if len(lines) == 1 && lines[0] == "" {
		return []string{}, nil
	}
	return lines, nil
}

// LastCommitMessage returns the last commit message on the current branch
func LastCommitMessage(branch string) (string, error) {
	out, err := exec.Command("git", "log", "-1", "--pretty=%s", branch).Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(out)), nil
}
