package timer

import (
	"testing"
	"time"
)

func hello() []string {
	result := []string{"Hello"}
	return result
}

func TestExecuteForDuration(t *testing.T) {
	start := time.Now()
	result := ExecuteForDuration(func() []string { return hello() }, 5*time.Second)
	duration := time.Since(start)
	expectedDuration := 5.0
	margin := 1.0
	if duration.Seconds() < expectedDuration-margin || duration.Seconds() > expectedDuration+margin {
		t.Errorf("Expected execution duration is %.2f seconds, but got %.2f s", expectedDuration, duration.Seconds())
	}
	minExecutions := 4
	maxExecutions := 5

	if len(result) < minExecutions || len(result) > maxExecutions {
		t.Errorf("Expected number of executions should be between %d and %d but got %d", minExecutions, maxExecutions, len(result))
	}
	if len(result) > 0 && result[0] != "Hello" {
		t.Errorf("Expected results to contain Hello, but got %v", result)
	}
}
