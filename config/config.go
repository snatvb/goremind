package config

import "time"

// var timings = map[int]time.Duration{
// 	0: 1 * time.Hour,
// 	1: 12 * time.Hour,
// 	2: 24 * time.Hour,
// 	3: 36 * time.Hour,
// 	4: 48 * time.Hour,
// 	5: 72 * time.Hour,
// }

var timings = map[int]time.Duration{
	0: 1 * time.Minute / 2,
	1: 12 * time.Minute / 2,
	2: 24 * time.Minute / 2,
	3: 36 * time.Minute / 2,
	4: 48 * time.Minute / 2,
	5: 72 * time.Minute / 2,
}
