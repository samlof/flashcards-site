// Package model has gengql generated types
package model

// LogModifiers maps result into a modifier to be used with scheduled_next
var LogModifiers map[CardResult]float64 = map[CardResult]float64{
	CardResultRetry:   0,
	CardResultAverage: 1.5,
	CardResultBad:     0.75,
	CardResultGood:    2,
}
