// Package functions - пакет вспомогательных функций.
package functions

import (
	"time"

	"go-app/internal/utils/dictionary"
)

// DateDiff - вспомогательная функция расчета разницы в датах.
func DateDiff(date1, date2 time.Time) int {
	duration := date2.Sub(date1).Hours() / dictionary.HourDivider
	return int(duration)
}
