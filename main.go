package main

import (
	"fmt"
	"os"

	"github.com/adit-rah/git-wrapper/cmd"

	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "gw",
		Short: "gw is a Git workflow CLI wrapper",
		Long:  "gw helps manage branches, commits, PRs, and folding branches into base branches.",
	}

	// Add commands
	rootCmd.AddCommand(cmd.CreateCmd)
	rootCmd.AddCommand(cmd.ModifyCmd)
	rootCmd.AddCommand(cmd.SubmitCmd)
	rootCmd.AddCommand(cmd.FoldCmd)
	rootCmd.AddCommand(cmd.StatusCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
