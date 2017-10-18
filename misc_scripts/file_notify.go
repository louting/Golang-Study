package main

import (
	"log"

	"github.com/howeyc/fsnotify"
)

func main() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}

	done := make(chan bool)

	// Process events

	go func() {
		for {
			select {
			case ev := <-watcher.Event:
				log.Println("event:", ev)
			case err := <-watcher.Error:
				log.Println("error:", err)
			}
		}
	}()

	err = watcher.Watch("/Users/tony/go")
	if err != nil {
		log.Fatal(err)
	}

	// Hang so program doesn't exit.其实这个地方用个for{}效果一样。
	<-done

	/* ... do stuff ... */
	watcher.Close()
}
