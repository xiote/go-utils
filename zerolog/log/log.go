package log

import (
	"fmt"
	"github.com/xiote/go-utils/zerolog"
	"time"
)

type StderrWriter struct {
	//
}

func (w StderrWriter) Write(loginId string, ticketingId string, stepName string, message string) {

	// fmt.Fprintln(os.Stderr, "hello world")

	go println(fmt.Sprintf("%s | %s | %s | %s | %s",
		loginId,
		ticketingId,
		// realClock.Now().Format("0102_15:04:05.000"),
		time.Now().Format("0102_15:04:05.000"),
		stepName,
		message,
	))
}

// Logger is the global logger.
// var Logger = zerolog.New(os.Stderr).With().Logger()
var Logger = zerolog.New(StderrWriter{}).With().Logger()

func Log() *zerolog.Event {
	return nil
}

func Error() *zerolog.Event {
	return nil
}

func With() zerolog.Context {
	return Logger.With()
}

func Printf(format string, v ...interface{}) {
	// Logger.Debug().CallerSkipFrame(1).Msgf(format, v...)
}

// func Printf(format string, v ...interface{}) {
//
// }

// var logchan chan string
//
// func init() {
// 	logchan = make(chan string, 1000)
//
// 	go func() {
// 		// 채널이 닫힐 때까지 메시지 받으면 로깅
// 		for msg := range logchan {
// 			log.Print(msg)
// 		}
// 	}()
// }
//
// func SetChanSize(size int) {
// 	logchan = make(chan string, size)
// }
//
// func SetFlags(flag int) {
// 	log.SetFlags(flag)
// }
//
// func Print(v ...interface{}) {
// 	// log.Printf(v...)
// 	logchan <- fmt.Sprint(v...)
// }
//
// func Printf(format string, v ...interface{}) {
// 	// log.Printf(format, v...)
// 	logchan <- fmt.Sprintf(format, v...)
// }
//
// func Println(a ...interface{}) (n int, err error) {
// 	logchan <- fmt.Sprintln(a...)
// 	return
// }
