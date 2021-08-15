package scanner

import (
	"context"
	"fmt"
	"github.com/borankux/hydra/redis"
	"io/ioutil"
	"os"
)

func ScanDir(dir string, verbose bool) ([]File, error) {

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return nil, err
	}

	files, _ := ioutil.ReadDir(dir)
	var ret []File
	for _, v := range files {
		file := File{}.BuildFromFileInfo(v, dir)

		if verbose {
			fmt.Println(file.AbsolutePath, file.Size())
		}
		ret = append(ret, file)
	}
	return ret, nil
}

func ScanDeep(root string, verbose bool) ([]File, error) {
	var ret []File
	rdb := redis.GetRedis()
	scanned, err := ScanDir(root, verbose)

	if err != nil {
		return nil, err
	}
	for _, v := range scanned {
		if !verbose {
			//rdb.Set(context.Background(), v.AbsolutePath, 1, 0)
			rdb.HMSet(context.Background(), v.AbsolutePath, "size",v.Size(), "isdir", v.IsDir())
		}

		ret = append(ret, v)
		if v.IsDir() {
			found, err := ScanDeep(v.AbsolutePath, verbose)
			if err == nil {
				ret = append(ret, found...)
			}
		}
	}
	return ret, nil
}
