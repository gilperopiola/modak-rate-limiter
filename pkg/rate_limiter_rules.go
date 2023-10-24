package pkg

import "time"

type LimiterRule struct {
	maxNotifications int
	in               time.Duration
}

var (
	DefaultRules = map[NotifType]LimiterRule{
		NotifTypeStatus: {
			maxNotifications: 2,
			in:               time.Minute,
		},
		NotifTypeNews: {
			maxNotifications: 1,
			in:               24 * time.Hour,
		},
		NotifTypeMarket: {
			maxNotifications: 3,
			in:               time.Hour,
		},
	}
)
