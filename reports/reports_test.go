package reports

import (
	"os"
	"testing"
)

func TestStats(t *testing.T) {
	results := []string{"0.04s     1256 200", "0.05s          1256 200", "0.06s          1256 200"}
	minResponse, maxResponse, notOkstatus := Stats(results)
	if minResponse != 0.04 {
		t.Errorf("Expected minimum response time is %.2f but got %.2f", 0.04, minResponse)
	}
	if maxResponse != 0.06 {
		t.Errorf("Expected maximum response time is %.2f but got %.2f", 0.06, maxResponse)
	}
	if notOkstatus != 0 {
		t.Errorf("Expected non 200 statuses were %d, but got %d", 0, notOkstatus)
	}
}

func TestGenerateCSV(t *testing.T) {
	results := []string{"0.04s     1256 200", "0.05s          1256 200", "0.06s          1256 200"}
	GenerateCSV("/tmp/reports_test_DELETE_AFTERTEST", results)
	_, err := os.Stat("/tmp/reports_test_DELETE_AFTERTEST")

	if err != nil {
		t.Errorf("Expected a csv file to be found at %s", "/tmp/reports_test_DELETE_AFTERTEST")
	}
	os.Remove("/tmp/reports_test_DELETE_AFTERTEST")
}
