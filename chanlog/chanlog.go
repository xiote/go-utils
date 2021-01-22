package chanlog

import (
	"fmt"
	"log"
)

var logchan chan string

func init() {
	logchan = make(chan string, 1000)

	go func() {
		// 채널이 닫힐 때까지 메시지 받으면 로깅
		for msg := range logchan {
			log.Print(msg)
		}
	}()
}

func SetChanSize(size int) {
	logchan = make(chan string, size)
}

func SetFlags(flag int) {
	log.SetFlags(flag)
}

func Print(v ...interface{}) {
	// log.Printf(v...)
	logchan <- fmt.Sprint(v...)
}

func Printf(format string, v ...interface{}) {
	// log.Printf(format, v...)
	logchan <- fmt.Sprintf(format, v...)
}

func Println(a ...interface{}) (n int, err error) {
	logchan <- fmt.Sprintln(a...)
	return
}
