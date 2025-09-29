package cmd

import (
	"fmt"

	"github.com/adit-rah/git-wrapper/internal/gh"
	"github.com/adit-rah/git-wrapper/internal/git"

	"github.com/spf13/cobra"
)

var StatusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show the current branch status and GitHub PR info",
	Run: func(cmd *cobra.Command, args []string) {
		// Current branch
		currentBranch, err := git.GetCurrentBranch()
		if err != nil {
			fmt.Println("Error getting current branch:", err)
			return
		}

		// Base branch
		baseBranch := "main"
		if data, err := git.ReadBranchBase(); err == nil {
			baseBranch = data
		}

		// Commits on this branch
		count, err := git.CountCommitsSinceBase(baseBranch, currentBranch)
		if err != nil {
			fmt.Println("Error counting commits:", err)
			return
		}

		// Uncommitted changes
		uncommitted, err := git.UncommittedChanges()
		if err != nil {
			fmt.Println("Error getting uncommitted changes:", err)
			return
		}

		// Last commit message
		lastCommit, err := git.LastCommitMessage(currentBranch)
		if err != nil {
			fmt.Println("Error getting last commit:", err)
			return
		}

		// GitHub PR info
		prURL, prState, err := gh.GetPRInfo(currentBranch)
		if err != nil {
			prURL = "None"
			prState = ""
		}

		// Print status
		fmt.Printf("Current branch: %s\n", currentBranch)
		fmt.Printf("Base branch: %s\n", baseBranch)
		fmt.Printf("Commits on this branch: %d\n", count)
		fmt.Printf("Uncommitted changes: %d files\n", len(uncommitted))
		fmt.Printf("Last commit: %s\n", lastCommit)
		fmt.Printf("GitHub PR: %s", prURL)
		if prState != "" {
			fmt.Printf(" (%s)", prState)
		}
		fmt.Println()

		// Next steps hint
		fmt.Println("\nNext steps:")
		fmt.Println("  - gw modify   # add another commit")
		fmt.Println("  - gw submit   # push & create/update PR")
		fmt.Println("  - gw fold     # merge into base branch")
	},
}
