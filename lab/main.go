package main

import (
	"github.com/gosuri/uiprogress"
	"math/rand"
	"sync"
	"time"
)

func job() {

}

func worker() {

}

func main() {
	total := 10000
	uiprogress.Start()
	bar := uiprogress.AddBar(total)
	bar.AppendCompleted()
	bar.PrependElapsed()

	var wg sync.WaitGroup
	for i := 0; i < total; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			time.Sleep(time.Second * time.Duration(rand.Intn(10)))
			bar.Incr()
		}()
	}

	wg.Wait()
	uiprogress.Stop()

	//for i := 0; i < total; i++ {
	//	time.Sleep(time.Second * 1)
	//	bar.Incr()
	//}
	//uiprogress.Stop()
}
