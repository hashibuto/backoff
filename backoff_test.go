package backoff

import (
	"testing"
	"time"
)

func TestBackoff(t *testing.T) {
	err := Backoff(Config{}, time.Second*2, func() (bool, error) {
		return false, nil
	})
	if err == nil {
		t.Errorf("Expected to hit a timeout condition")
	}
}
