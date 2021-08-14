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
		//rdb := redis.GetRedis()
		//res:= rdb.Info(context.Background(), "Keyspace")
		//rule := `db[0-9]{0,2}\:keys\=([0-9]+)`
		//re, _ := regexp.Compile(rule)
		//found := re.FindStringSubmatch(res.String())[1]
		//kernal32,_ := syscall.LoadLibrary("kernal32.dll")
	},
}
