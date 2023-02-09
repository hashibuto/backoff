package backoff

import (
	"fmt"
	"time"
)

type Config struct {
	Min  time.Duration
	Max  time.Duration
	Step float64
}

// Backoff calls the callback function, progressively backing off until either the function succeeds, and error condition
// is hit, or a timeout occurs.
func Backoff(config Config, timeout time.Duration, callback func() (bool, error)) error {
	if config.Max == 0 {
		config.Max = 10 * time.Second
	}

	if config.Min == 0 {
		config.Min = 250 * time.Millisecond
	}

	if config.Step == 0 {
		config.Step = 1.5
	}

	current := config.Min
	start := time.Now()
	for {
		complete, err := callback()
		if err != nil {
			return err
		}
		if complete {
			return nil
		}

		if timeout != 0 {
			total := time.Since(start)
			if total+current >= timeout {
				return fmt.Errorf("Timeout limit exceeded")
			}
		}

		time.Sleep(current)

		current = time.Duration(float64(current) * config.Step)
		if current > config.Max {
			current = config.Max
		}
	}
}
