// logminute.go -- log something every minute, ad infinutum.
//
// logs to STDOUT.
//
// (where "minute" is whatever the system thinks it is)

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

	ticker := time.NewTicker(time.Minute)
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
