package scheduler

import (
	"fmt"
	"time"
)

func StartScheduler() {
	go func() {
		for {
			now := time.Now()
			next := time.Date(now.Year(), now.Month(), now.Day(), 12, 0, 0, 0, now.Location())
			if next.Before(now) {
				next = next.Add(24 * time.Hour)
			}
			time.Sleep(time.Until(next))
			executeDailyTask()
		}
	}()
}

// executeDailyTask 执行每天的任务
func executeDailyTask() {
	fmt.Printf("Executing daily task at %s\n", time.Now())
}
