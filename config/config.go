package config

import "time"

var Timings = map[int]time.Duration{
	0: 1 * time.Hour,
	1: 12 * time.Hour,
	2: 24 * time.Hour,
	3: 36 * time.Hour,
	4: 48 * time.Hour,
	5: 72 * time.Hour,
}

// var Timings = map[int]time.Duration{
// 	0: 1 * time.Minute / 2,
// 	1: 1 * time.Minute,
// 	2: 1 * time.Minute,
// 	3: 1 * time.Minute,
// 	4: 1 * time.Minute,
// 	5: 1 * time.Minute,
// }
