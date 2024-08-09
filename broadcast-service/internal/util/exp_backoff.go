package util

import (
	"broadcast-service/internal"
	"fmt"
	"time"

	"go.uber.org/zap"
)

var (
	logger, _ = internal.WireLogger()
)


func ExpBackoff(delay time.Duration, multiplier int, maxRetries int, f func() error) error {
	count := 0
	for {
		err := f()
		if err != nil {
			if count > maxRetries {
				return err
			}
			logger.Warn(
				fmt.Sprintf("failed to execute, retrying after %v...", delay),
				zap.String("trace", err.Error()),
			)
			<-time.After(delay)
			delay *= time.Duration(multiplier)
			count += 1
			continue
		}
		return nil
	}
}