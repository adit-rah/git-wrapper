package cmd

import (
	"fmt"
	"io/ioutil"

	"github.com/adit-rah/git-wrapper/internal/gh"
	"github.com/adit-rah/git-wrapper/internal/git"

	"github.com/spf13/cobra"
)

var SubmitCmd = &cobra.Command{
	Use:   "submit",
	Short: "Push the current branch and create a pull request",
	Run: func(cmd *cobra.Command, args []string) {
		currentBranch, err := git.GetCurrentBranch()
		if err != nil {
			fmt.Println("Error getting current branch:", err)
			return
		}

		baseBranch := "main"
		if data, err := ioutil.ReadFile(".branch_base"); err == nil {
			baseBranch = string(data)
		}

		if err := git.PushBranch(currentBranch); err != nil {
			fmt.Println("Error pushing branch:", err)
			return
		}

		prURL, err := gh.CreatePR(baseBranch, currentBranch)
		if err != nil {
			fmt.Println("Error creating PR:", err)
			return
		}

		fmt.Printf("Pull Request created: %s\n", prURL)
	},
}
