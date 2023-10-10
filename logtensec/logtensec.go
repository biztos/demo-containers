// logtensec.go -- log something every ten seconds, ad infinutum.
//
// logs to STDOUT.
//
// This is less boring than waiting for logminute but less annoying than
// logging every second.  This way you get "only" 8640 logs per day. :-/

package main

import (
	"log"
	"os"
	"sync"
	"time"
)

var Lock = &sync.Mutex{}
var Count = 0

func main() {

	// hostname is nice to have...
	hostname := os.Getenv("HOSTNAME")
	if hostname == "" {
		hostname = "<UNKOWN HOST>"
	}

	ticker := time.NewTicker(time.Second * 10)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			Lock.Lock()
			Count++
			Lock.Unlock()
			log.Println(hostname, "Log number", Count)
		}
	}
}
