package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var releaseCmd = &cobra.Command{
	Use:   "release <assets>...",
	Short: "Automate the Git tagging and GitHub release process",
	Long: `Automate the Git tagging and GitHub release process.

Arguments:
  assets   One or more file paths to be uploaded as release assets`,
	Args: cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return release(args)
	},
}

func init() {
	rootCmd.AddCommand(releaseCmd)
}

func release(assets []string) error {
	dir, err := gitRootDir()
	if err != nil {
		return err
	}
	file := filepath.Join(dir, ".version")
	content, err := os.ReadFile(file)
	if err != nil {
		return err
	}
	version := strings.TrimSpace(string(content))
	// ignore error: may fail on first release when there are no tags yet
	cmd := exec.Command("git", "describe", "--tags", "--abbrev=0")
	tag, _ := cmd.Output()
	if strings.TrimSpace(string(tag)) == version {
		return nil
	}
	err = run("git", "add", file)
	if err != nil {
		return err
	}
	message := fmt.Sprintf("release: %s", version)
	err = run("git", "commit", "-m", message)
	if err != nil {
		return err
	}
	err = run("git", "tag", version)
	if err != nil {
		return err
	}
	err = run("git", "push")
	if err != nil {
		return err
	}
	err = run("git", "push", "--tags")
	if err != nil {
		return err
	}
	args := []string{"release", "create", version, "--generate-notes"}
	args = append(args, assets...)
	err = run("gh", args...)
	if err != nil {
		return err
	}
	return nil
}

func run(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
