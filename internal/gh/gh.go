package gh

import (
	"os/exec"
	"strings"
)

func CreatePR(base, head string) (string, error) {
	out, err := exec.Command("gh", "pr", "create", "--base", base, "--head", head, "--fill", "--json", "url", "-q", ".url").Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(out)), nil
}

// GetPRInfo returns the PR URL and state for the given branch
func GetPRInfo(branch string) (string, string, error) {
	// Check if a PR exists
	out, err := exec.Command("gh", "pr", "view", branch, "--json", "url,state", "-q", ".url+\" \"+.state").Output()
	if err != nil {
		return "", "", err
	}
	parts := strings.SplitN(strings.TrimSpace(string(out)), " ", 2)
	if len(parts) != 2 {
		return parts[0], "", nil
	}
	return parts[0], parts[1], nil
}
