package placeholder

import (
	"fmt"
	"platform/logging"
	"time"
)

type DayHandler struct {
	logging.Logger
}

func (dh DayHandler) GetDay() string {
	return fmt.Sprintf("Day: %v", time.Now().Day())
}
