package cmd

import (
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func init () {
	rootCMD.AddCommand(registerCMD)
}

var registory string
var token string

var registerCMD = &cobra.Command{
	Use: "register",
	Run: func(cmd *cobra.Command, args []string) {
		color.Green("registering...")
	},
}
