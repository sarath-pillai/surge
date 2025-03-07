package timer

import (
	"fmt"
	"time"
)

func ExecuteForDuration(f func() []string, duration time.Duration) []string {
	var results []string
	timeout := time.After(duration)
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-timeout:
			fmt.Println("Execution duration complete")
			return results
		case <-ticker.C:
			results = append(results, f()...)
		}
	}
}
