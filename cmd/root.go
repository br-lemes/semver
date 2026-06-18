package cmd

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "semver",
	RunE: func(cmd *cobra.Command, args []string) error {
		dir, err := gitRootDir()
		if err != nil {
			return err
		}
		version := version()
		cmd.Println(version)
		file := filepath.Join(dir, ".version")
		return os.WriteFile(file, []byte(version), 0644)
	},
}

func Execute(version string) error {
	rootCmd.Version = version
	return rootCmd.Execute()
}

func gitRootDir() (string, error) {
	cmd := exec.Command("git", "rev-parse", "--show-toplevel")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", errors.New(string(output))
	}
	return strings.TrimSpace(string(output)), nil
}

func version() string {
	major := 0
	minor := 0
	patch := 0

	cmd := exec.Command("git", "log", "--pretty=format:%s")
	output, err := cmd.Output()
	if err != nil {
		return fmt.Sprintf("v%d.%d.%d", major, minor, patch)
	}

	logs := strings.Split(string(output), "\n")
	breakingChangeRegex := regexp.MustCompile(`^[a-z]+!:\s*`)

	for i := len(logs) - 1; i >= 0; i-- {
		line := strings.TrimSpace(logs[i])
		if line == "" {
			continue
		}

		if breakingChangeRegex.MatchString(line) {
			major++
			minor = 0
			patch = 0
			continue
		}

		if strings.HasPrefix(line, "feat") {
			minor++
			patch = 0
			continue
		}

		if strings.HasPrefix(line, "fix") {
			patch++
			continue
		}
	}

	return fmt.Sprintf("v%d.%d.%d", major, minor, patch)
}
