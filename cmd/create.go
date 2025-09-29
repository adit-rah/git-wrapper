package cmd

import (
	"fmt"

	"github.com/adit-rah/git-wrapper/internal/git"

	"github.com/spf13/cobra"
)

var CreateCmd = &cobra.Command{
	Use:   "create [type] [name]",
	Short: "Create a new branch with an initial commit",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		branchType := args[0]
		branchName := args[1]

		baseBranch, err := git.GetCurrentBranch()
		if err != nil {
			fmt.Println("Error getting current branch:", err)
			return
		}

		newBranch := fmt.Sprintf("%s/%s", branchType, branchName)
		if err := git.CreateBranch(newBranch); err != nil {
			fmt.Println("Error creating branch:", err)
			return
		}

		if err := git.AddAllAndCommit(fmt.Sprintf("%s update 1", newBranch)); err != nil {
			fmt.Println("Error committing changes:", err)
			return
		}

		// TODO: Save metadata (branchType, branchName, baseBranch, commit count)
		fmt.Printf("Created branch %s from %s and committed initial changes.\n", newBranch, baseBranch)
	},
}
