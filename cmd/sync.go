package cmd

import (
	"context"
	"github.com/borankux/hydra/api"
	"github.com/borankux/hydra/redis"
	"github.com/gosuri/uiprogress"
	"github.com/spf13/cobra"
	"strconv"
)

var key string

var syncCMD = &cobra.Command {
	Use: "sync",
	Short: "sycns local data to server",
	Run: func(cmd *cobra.Command, args []string) {
		rdb:= redis.GetRedis()
		res := rdb.Keys(context.Background(), key)
		uiprogress.Start()
		bar := uiprogress.AddBar(len(res.Val()))
		bar.AppendCompleted()
		bar.PrependElapsed()
		bar.AppendElapsed()
		for _,key := range res.Val() {
			res := rdb.HMGet(context.Background(), key, "size", "isdir")
			size, _ := strconv.ParseUint(res.Val()[0].(string), 10, 16)
			isDir, _ := strconv.ParseBool(res.Val()[1].(string))
			api.CreateFile(key, uint16(size), isDir)
			bar.Incr()
		}
		uiprogress.Stop()
	},
}


func init() {
	syncCMD.Flags().StringVarP(&key, "key", "k", "*", "key pattern to match the search before upload")
	rootCMD.AddCommand(syncCMD)
}