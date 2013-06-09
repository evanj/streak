package streak

import (
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
