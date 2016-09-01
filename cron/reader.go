package cron

import (
	"fmt"
	"time"
)

func ReadFeedsEach(p time.Duration) {
	ticker := time.NewTicker(p)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				readFeeds()
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
}

func readFeeds() {
	fmt.Printf("hello")
}
