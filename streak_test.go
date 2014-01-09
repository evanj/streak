package streak

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func expectEquals(t *testing.T, expected time.Time, actual time.Time) {
	if actual != expected {
		t.Errorf("Expected %v; actual %v", expected, actual)
	}
}

func TestTimestampRoundDown(t *testing.T) {
	roundDown, err := time.Parse(time.RFC3339Nano, "2013-06-09T14:56:59.001499999Z")
	if err != nil {
		t.Fatal("bad time constant", err)
	}
	roundTripped := TimestampToTime(TimeToTimestamp(roundDown))
	expectEquals(t, roundDown.Truncate(time.Millisecond), roundTripped)
}

func TestTimestampRoundUp(t *testing.T) {
	roundUp, err := time.Parse(time.RFC3339Nano, "2013-06-09T14:56:59.001500000Z")
	if err != nil {
		t.Fatal("bad time constant", err)
	}
	roundTripped := TimestampToTime(TimeToTimestamp(roundUp))
	expectEquals(t, roundUp.Truncate(time.Millisecond).Add(time.Millisecond), roundTripped)
}

// streak occasionally returns a 500 error with HTML from AppEngine
func Test500Error(t *testing.T) {
	// start test server
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)
	defer server.Close()

	// handle the API call
	mux.HandleFunc("/pipelines", func(w http.ResponseWriter, r *http.Request) {
		// Return a 500 server error
		w.Header().Set("Content-Type", "text/html; charset=UTF-8")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("<html><head></head><body>Error: Server Error</body></html>"))
	})

	client := newWithbaseUrl("apiKey", server.URL)
	pipelines, err := client.GetPipelines()
	if err == nil {
		t.Errorf("Expected error, got pipelines: %v", pipelines)
	} else if !strings.Contains(err.Error(), "failed code: 500") {
		t.Errorf("Unexpected error: %v", err)
	}
}
