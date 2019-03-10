package cmd

import (
	"fmt"
	"os"

	"github.com/tukaelu/opsworks-cli/credentials"
	"gopkg.in/AlecAivazis/survey.v1"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(timebasedCmd)

	timebasedCmd.Flags().StringVarP(&profile, "profile", "p", "", "AWS Profile name on your ~/.aws/credentilas")
	timebasedCmd.Flags().StringVarP(&region, "region", "r", "us-east-1", "AWS Region")
}

var timebasedCmd = &cobra.Command{
	Use:   "time",
	Short: "time-based instance configuration",
	Run:   doSubCommand,
}

var region string
var profile string

func doSubCommand(cmd *cobra.Command, args []string) {
	fmt.Println("run time sub-command")

	credsfile := os.Getenv("HOME") + "/.aws/credentials"

	creds, err := credentials.NewCredentials(credsfile)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err = doInteractive(creds); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func doInteractive(creds *credentials.Credentials) error {
	profiles, err := creds.GetAwsProfileNames()
	if err != nil {
		return err
	}

	profile, err := selected("Choose an AWS Profile", profiles)
	if err != nil {
		return err
	}

	fmt.Printf("selected profile: %s\n", profile)

	return nil
}

func selected(message string, items []string) (string, error) {
	selected := ""
	prompt := &survey.Select{
		Message: fmt.Sprintf("%s: ", message),
		Options: items,
	}

	err := survey.AskOne(prompt, &selected, nil)
	return selected, err
}
