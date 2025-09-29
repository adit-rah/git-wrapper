package cmd

import (
	"fmt"
	"io/ioutil"

	"github.com/adit-rah/git-wrapper/internal/git"

	"github.com/spf13/cobra"
)

var FoldCmd = &cobra.Command{
	Use:   "fold",
	Short: "Merge the current branch into its base and delete it locally and remotely",
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

		fmt.Printf("Merging %s into %s...\n", currentBranch, baseBranch)

		if err := git.CheckoutBranch(baseBranch); err != nil {
			fmt.Println("Error checking out base branch:", err)
			return
		}

		if err := git.PullBranch(baseBranch); err != nil {
			fmt.Println("Error pulling base branch:", err)
			return
		}

		if err := git.MergeBranch(currentBranch); err != nil {
			fmt.Println("Error merging branch:", err)
			return
		}

		if err := git.PushBranch(baseBranch); err != nil {
			fmt.Println("Error pushing base branch:", err)
			return
		}

		if err := git.DeleteBranch(currentBranch); err != nil {
			fmt.Println("Error deleting branch:", err)
			return
		}

		fmt.Printf("✅ Branch %s merged into %s and deleted.\n", currentBranch, baseBranch)
	},
}
