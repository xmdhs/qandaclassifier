package classification

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"
)

func init() {
	go getid()
}

func Server() {
	mux := http.NewServeMux()
	mux.HandleFunc("/classification", index)
	mux.HandleFunc("/classification/add", add)
	mux.HandleFunc("/classification/fenlei", fenlei)
	mux.HandleFunc("/classification/classification", classification)
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
	tid      string
	i        int
	infotemp info
}

var tids jindu

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(indexhtml))
}

func add(w http.ResponseWriter, r *http.Request) {
	var tid string
	tids.RLock()
	if tids.tid == "" {
		var info info
		for k := range tidch {
			tid = k
			var err error
			if tid == "" {
				w.Write([]byte(`全部分类完成`))
				return
			}
			info, err = getidinfo(tid)
			if err != nil {
				log.Println(err)
				tosave <- tid
			} else {
				break
			}
		}
		tids.RUnlock()
		tids.Lock()
		tids.infotemp = info
		tids.i = len(info.message)
		tids.tid = tid
		tids.Unlock()
	} else {
		tid = tids.tid
		tids.RUnlock()
	}
	tids.RLock()
	for _, vv := range tids.infotemp.message {
		addtemp.Lock()
		addtemp.i = tids.i
		if pidinfo := addtemp.temp[vv.pid]; pidinfo.category == "" {
			addtemp.temp[vv.pid] = pidtemp{text: vv.message}
			addtemp.Unlock()
			genhtml(w, vv)
			tids.RUnlock()
			return
		}
		addtemp.Unlock()
	}
	addtemp.Lock()
	addtemp.ii = 0
	addtemp.temp = map[string]pidtemp{}
	addtemp.Unlock()

	tids.RUnlock()
	tids.Lock()
	tids.tid = ""
	tids.Unlock()
	tosave <- tid
	http.Redirect(w, r, `/classification/add`, 302)
}

type temp struct {
	sync.RWMutex
	//键为 pid
	temp map[string]pidtemp
	i    int
	ii   int
}

var addtemp temp

func init() {
	addtemp.temp = make(map[string]pidtemp)
}

type pidtemp struct {
	category string
	text     string
}

func classification(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	if len(q["q"]) == 0 {
		w.Write([]byte(pinfen))
	} else {
		text := q["q"][0]
		scores := handler.Categorize(text)
		if len(scores) == 0 {
			e(w, "未知")
		} else {
			for _, v := range scores {
				w.Write([]byte(v.Category + ": " + strconv.FormatFloat(v.Score, 'f', 5, 64) + "\n"))
			}
		}
	}
}

var tidch = make(chan string, 1)
var tosave = make(chan string, 1)

func getid() {
	b, err := ioutil.ReadFile(`tids.json`)
	if err != nil {
		panic(err)
	}
	m := map[string]bool{}
	err = json.Unmarshal(b, &m)
	if err != nil {
		panic(err)
	}
	var tid string
	for {
		for k, ok := range m {
			if !ok {
				tid = k
				break
			}
		}
		if tid == "" {
			close(tidch)
		}
		select {
		case tidch <- tid:
		case tid := <-tosave:
			m[tid] = true
			f, err := os.Create(`tids.json`)
			if err != nil {
				panic(err)
			}
			b, err := json.Marshal(m)
			if err != nil {
				panic(err)
			}
			f.Write(b)
			f.Close()
		}
	}

}
