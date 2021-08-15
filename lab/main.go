package main

import (
	"context"
	"fmt"
	"github.com/borankux/hydra/redis"
)

func job() {

}

func worker() {

}

func main() {
	rdb := redis.GetRedis()
	res := rdb.HMSet(context.Background(), "random", "a", "a1", "b", "b1")
	if res.Val() {
		fmt.Println("success")
	} else {
		fmt.Println("failed")
	}
	//for i := 0; i < total; i++ {
	//	time.Sleep(time.Second * 1)
	//	bar.Incr()
	//}
	//uiprogress.Stop()
}
