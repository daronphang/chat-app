package util

import (
	"fmt"
	"message-service/internal"
	"time"

	"go.uber.org/zap"
)

var (
	logger, _ = internal.WireLogger()
)


func ExpBackoff(delay time.Duration, multiplier int, maxRetries int, errMsg string, f func() error) error {
	count := 0
	for {
		err := f()
		if err != nil {
			if count > maxRetries {
				return err
			}
			logger.Warn(
				fmt.Sprintf("%v, retrying after %v...", errMsg, delay),
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