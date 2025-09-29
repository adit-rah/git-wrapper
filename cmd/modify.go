package cmd

import (
	"fmt"
	"io/ioutil"
	"strconv"

	"github.com/adit-rah/git-wrapper/internal/git"

	"github.com/spf13/cobra"
)

var ModifyCmd = &cobra.Command{
	Use:   "modify",
	Short: "Add changes as the next commit on the current branch",
	Run: func(cmd *cobra.Command, args []string) {
		// Read commit counter
		countBytes, _ := ioutil.ReadFile(".gh_branch_meta")
		count := 1
		if len(countBytes) > 0 {
			if c, err := strconv.Atoi(string(countBytes)); err == nil {
				count = c
			}
		}
		count++

		currentBranch, err := git.GetCurrentBranch()
		if err != nil {
			fmt.Println("Error getting current branch:", err)
			return
		}

		message := fmt.Sprintf("%s update %d", currentBranch, count)
		if err := git.AddAllAndCommit(message); err != nil {
			fmt.Println("Error committing changes:", err)
			return
		}

		// Update counter
		_ = ioutil.WriteFile(".gh_branch_meta", []byte(strconv.Itoa(count)), 0644)
		fmt.Printf("Committed: %s\n", message)
	},
}
