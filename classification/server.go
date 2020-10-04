package classification

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

var tids jindu

func init() {
	b, err := ioutil.ReadFile(`tids.json`)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(b, &tids.jindu)
	if err != nil {
		panic(err)
	}
}

func Server() {
	mux := http.NewServeMux()
	mux.HandleFunc("/classification", index)
	mux.HandleFunc("/classification/add", add)
	server := http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		Addr:         ":8083",
		Handler:      mux,
	}
	log.Println(server.ListenAndServe())
}

type jindu struct {
	sync.RWMutex
	//[tid]bool
	jindu map[string]bool
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(indexhtml))
}

func add(w http.ResponseWriter, r *http.Request) {
	tids.Lock()
	for k, v := range tids.jindu {
		if !v {
			info, err := getidinfo(k)
			if err == nil {
				addtemp.Lock()
				defer addtemp.Unlock()
				for _, vv := range info.message {
					addtemp.i = len(info.message)
					if pidinfo := addtemp.temp[vv.pid]; pidinfo.category == "" {
						addtemp.temp[vv.pid] = pidtemp{text: vv.message}
						tids.Unlock()
						genhtml(w, vv)
						return
					}
				}
			} else {
				log.Println(err)
			}

			tids.jindu[k] = true
			tids.Unlock()

			addtemp.Lock()
			addtemp.temp = map[string]pidtemp{}
			addtemp.Unlock()

			tidsave()
			return
		}
		tids.Unlock()
	}
	w.Write([]byte(`全部分类完成`))
}

func tidsave() {
	f, err := os.Create(`tids.json`)
	if err != nil {
		panic(err)
	}
	tids.RLocker()
	b, err := json.Marshal(tids.jindu)
	tids.RUnlock()
	if err != nil {
		panic(err)
	}
	f.Write(b)
}

type temp struct {
	sync.RWMutex
	//键为 pid
	temp map[string]pidtemp
	i    int
}

var addtemp temp

type pidtemp struct {
	category string
	text     string
}
