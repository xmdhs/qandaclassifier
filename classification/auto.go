package classification

import (
	"log"
	"sync"
	"time"
)

func Autoclassification() {
	limit := make(chan struct{}, 7)
	wait := sync.WaitGroup{}
	for k := range tidch {
		limit <- struct{}{}
		wait.Add(1)
		go tid2class(k, limit, &wait)
	}
	wait.Wait()
}

func tid2class(tid string, limit chan struct{}, wait *sync.WaitGroup) {
	defer func() {
		tosave <- tid
		<-limit
		wait.Done()
	}()
	time.Sleep(1 * time.Second)
	info, err := getidinfo(tid)
	if err != nil {
		log.Println(err)
		return
	}
	handler.Training(info.message[0].message, fids[info.tid])
}

var fids = map[string]string{
	"110":  "原版问答",
	"265":  "基岩版问答",
	"1566": "周边问答",
	"431":  "联机问答",
	"266":  "Mod问答",
}
