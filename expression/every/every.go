package every

import (
	"time"
)

// Expression represents a cycle expression.
type Expression struct {
	Interval time.Duration
}

// Next implements jobs.Expression interface.
func (e *Expression) Next(from time.Time) time.Time {
	return from.Add(e.Interval)
}

// NewExpression returns a new cycle expression (every).
// A duration string supported, check time.ParseDuration.
// Less than a second are not supported (will round up to 1 second).
func NewExpression(dur time.Duration) *Expression {
	if dur < time.Second {
		dur = time.Second
	}

	return &Expression{
		Interval: dur,
	}
}
