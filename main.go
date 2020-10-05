package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"sync"

	"github.com/xmdhs/qandaclassifier/classification"
	"github.com/xmdhs/qandaclassifier/spider"
)

func main() {
	if len(os.Args) != 1 && os.Args[1] == "pa" {
		w := sync.WaitGroup{}
		suo := sync.Mutex{}
		tids := make(map[string]bool, 0)
		f, err := os.Open(`tids.json`)
		if err == nil {
			b, err := ioutil.ReadAll(f)
			if err != nil {
				panic(err)
			}
			err = json.Unmarshal(b, &tids)
			if err != nil {
				panic(err)
			}
		}
		f.Close()
		for v := range fids {
			v := v
			w.Add(1)
			go func() {
				m := spider.Fid2tids(v, 1, 500)
				suo.Lock()
				for k := range m {
					if _, ok := tids[k]; !ok {
						tids[k] = false
					}
				}
				suo.Unlock()
				w.Done()
			}()
		}
		w.Wait()
		f, err = os.Create(`tids.json`)
		defer f.Close()
		if err != nil {
			panic(err)
		}
		b, err := json.Marshal(tids)
		if err != nil {
			panic(err)
		}
		f.Write(b)
	} else if len(os.Args) != 1 && os.Args[1] == "fen" {
		classification.Server()
	} else if len(os.Args) != 1 && os.Args[1] == "auto" {
		classification.Autoclassification()
	}
}

var fids = map[string]string{
	"110":  "原版问答",
	"265":  "基岩版问答",
	"1566": "周边问答",
	"431":  "联机问答",
	"266":  "Mod问答",
}
