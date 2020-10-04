package main

import (
	"encoding/json"
	"os"
	"sync"

	"github.com/xmdhs/qandaclassifier/spider"
)

func main() {
	if len(os.Args) != 1 {
		w := sync.WaitGroup{}
		for v := range fids {
			v := v
			w.Add(1)
			go func() {
				m := spider.Fid2tids(v, 100)
				f, err := os.Create(v + `.json`)
				defer f.Close()
				if err != nil {
					panic(err)
				}
				b, err := json.Marshal(m)
				if err != nil {
					panic(err)
				}
				f.Write(b)
				w.Done()
			}()
		}
		w.Wait()
	} else {

	}
}

var fids = map[string]string{
	"110":  "原版问答",
	"265":  "基岩版问答",
	"1566": "周边问答",
	"431":  "联机问答",
	"266":  "Mod问答",
}
