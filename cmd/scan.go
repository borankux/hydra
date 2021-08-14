package cmd

import (
	"fmt"
	"github.com/borankux/hydra/scanner"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func init() {
	scanCommand.Flags().StringVarP(&directory, "dir", "d", "./", "Target directory to scan, default is the current directory")
	scanCommand.Flags().BoolVarP(&verbose, "verbose", "v", false, "Show details or not, default is false")
	rootCMD.AddCommand(scanCommand)
}

var directory string
var verbose bool

var scanCommand = &cobra.Command{
	Use:   "scan",
	Short: "scans stuff",
	Long:  "Logn Command",
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		color.Yellow("scanning %s", directory)
		files, err := scanner.ScanDeep(directory, verbose)
		if err != nil {
			fmt.Printf("faild to scan dir:%s, due to error:%v \n", directory, err)
		}
		fmt.Printf("found %d files from %s\n", len(files), directory)
	},
}
