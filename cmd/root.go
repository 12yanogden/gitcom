package cmd

import (
	"fmt"
	"os"

	"github.com/12yanogden/shell"
	sb "github.com/12yanogden/statusbar"
	"github.com/spf13/cobra"
)

// Base command
var rootCmd = &cobra.Command{
	Use:   "gitcom",
	Short: "git add, commit, and push.",
	Long: `Add, commit, and push changes in a local repo to its remote counterpart.
	
	For example: gitcom "some custom message"`,

	Run: gitcom,
}

func gitcom(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		fmt.Println("gitcom: must pass a commit message")
		os.Exit(1)
	}

	var addBar sb.StatusBar
	var commitBar sb.StatusBar
	var pushBar sb.StatusBar

	addBar.Start("Stage local changes")
	shell.Run("git", []string{"add", "."})
	addBar.Pass()

	commitBar.Start("Commit changes to local repository")
	shell.Run("git", []string{"commit", "-m", args[0]})
	commitBar.Pass()

	pushBar.Start("Push changes to remote repository")
	shell.Run("git", []string{"push"})
	pushBar.Pass()
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()

	if err != nil {
		os.Exit(1)
	}
}
