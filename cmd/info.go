package cmd

import (
	"fmt"
	"github.com/shirou/gopsutil/disk"
	"github.com/spf13/cobra"
	"runtime"
)

func init() {
	rootCMD.AddCommand(infoCMD)
}

var infoCMD = &cobra.Command{
	Use:   "info",
	Short: "fscan info",
	Run: func(cmd *cobra.Command, args []string) {
		os := runtime.GOOS
		if os == "windows" {
			diskStat,_ := disk.Usage("\\")
			partitions, err := disk.Partitions(true)
			if err != nil {
				return
			}
			fmt.Println(diskStat, partitions)
		}
	},
}
