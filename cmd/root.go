package cmd

import (
	"fmt"
	"github.com/borankux/fscan/helpers"
	"github.com/spf13/cobra"
	"os"
)

var rootCMD = &cobra.Command{
	Use:  "scan",
	Long: helpers.GetBanner("clear"),
}

func Execute() {
	if err := rootCMD.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
