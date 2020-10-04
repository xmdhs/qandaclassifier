package spider

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

var c = http.Client{Timeout: 5 * time.Second}

var cookie string

func init() {
	b, err := ioutil.ReadFile("cookie.txt")
	if err != nil {
		panic(err)
	}
	cookie = string(b)
}

func httpget(aurl string) ([]byte, error) {
	reqs, err := http.NewRequest("GET", aurl, nil)
	if err != nil {
		return nil, err
	}
	reqs.Header.Set("Accept", "*/*")
	reqs.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.105 Safari/537.36")
	reqs.Header.Set("Cookie", cookie)
	rep, err := c.Do(reqs)
	if rep != nil {
		defer rep.Body.Close()
	}
	if err != nil {
		return nil, err
	}
	if rep.StatusCode != 200 {
		return nil, NotOk
	}
	return ioutil.ReadAll(rep.Body)
}

var NotOk = errors.New("not 200")

func httpGetWithRetry(aurl string) []byte {
	for {
		b, err := httpget(aurl)
		if err != nil {
			log.Println(err)
			continue
		}
		return b
	}
}

func Getfidtids(fid, page string) ([]string, error) {
	b := httpGetWithRetry(`https://www.mcbbs.net/api/mobile/index.php?version=4&module=forumdisplay&fid=` + fid + `&filter=author&orderby=dateline&page=` + page)
	f, err := jsondecoder(b)
	if err != nil {
		return nil, err
	}
	tids := make([]string, 0, len(f.Variables.ForumThreadlist))
	for _, v := range f.Variables.ForumThreadlist {
		tids = append(tids, v.Tid)
	}
	return tids, nil
}

func Fid2tids(fid string, endpage int) map[string]bool {
	tidmap := make(map[string]bool, 0)
	for i := 1; i < endpage; i++ {
		time.Sleep(1 * time.Second)
		tids, err := Getfidtids(fid, strconv.Itoa(i))
		if err != nil {
			continue
		}
		for _, v := range tids {
			tidmap[v] = false
		}
		log.Println(fid, i)
	}
	return tidmap
}
