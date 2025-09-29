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
