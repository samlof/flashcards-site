// Package model has gengql generated types
package model

// LogModifiers maps result into a modifier to be used with scheduled_next
var LogModifiers = map[CardResult]float64{
	CardResultRetry: 0,
	CardResultBad:   .5,
	CardResultGood:  3,
	CardResultEasy:  10,
}
