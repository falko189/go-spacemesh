package config

import (
	"fmt"
	"gopkg.in/urfave/cli.v1"
	"runtime"
)

var (
// todo: add all command vars here (commands w no need for factory)
)

// NewVersionCommand is a command factory function to output the app's version.
// This is an example of an app cli command. To test use `/go-spacemesh version`.
func NewVersionCommand(appVersion, branch, commit string) cli.Command {
	return cli.Command{
		Name:      "version",
		Aliases:   []string{"v"},
		Usage:     "print versions",
		ArgsUsage: " ",
		Category:  "General commands",
		Action: func(c *cli.Context) error {
			fmt.Println("App Version:", appVersion)
			fmt.Printf("Git branch: %s. Commit: %s\n", branch, commit)
			fmt.Println("Go Version:", runtime.Version())
			fmt.Println("OS:", runtime.GOOS)
			fmt.Println("Arch:", runtime.GOARCH)
			return nil
		},
	}
}
