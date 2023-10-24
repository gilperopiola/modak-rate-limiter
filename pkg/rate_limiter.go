package pkg

import (
	"fmt"
	"time"
)

func NewRateLimiter(rules map[NotifType]LimiterRule) *rateLimiter {
	return &rateLimiter{
		datesStore: make(map[NotifType]map[Username][]time.Time),
		rules:      rules,
	}
}

type RateLimiter interface {
	HasExceededQuota(notification Notification) (bool, error)
	AddNotificationDate(notification Notification)
}

type rateLimiter struct {
	datesStore map[NotifType]map[Username][]time.Time
	rules      map[NotifType]LimiterRule
}

// HasExceededQuota returns an error if the notification type is not supported, then it returns true or false depending on the quota
func (rl *rateLimiter) HasExceededQuota(notification Notification) (bool, error) {

	// Get rule for notification type
	rule, ok := rl.rules[notification.Type]
	if !ok {
		return true, fmt.Errorf("%s notification type unsupported", notification.Type)
	}

	// Get user notification count according to the obtained rule
	notificationsInTimespan := rl.notificationsInTimespan(notification.Type, notification.Recipient, rule.in)

	// Check if quota is exceeded
	if notificationsInTimespan >= rule.maxNotifications {
		return true, nil
	}

	return false, nil
}

// notificationsInTimespan returns the number of notifications of a certain type for a certain user in a certain timespan
func (rl *rateLimiter) notificationsInTimespan(notifType NotifType, recipient Username, in time.Duration) int {

	// Create key if doesn't exist
	if rl.datesStore[notifType] == nil {
		rl.datesStore[notifType] = make(map[Username][]time.Time)
	}

	// Get array of dates
	recentNotificationDates := rl.datesStore[notifType][recipient]

	// Count notifications that match the rules' duration
	notificationsInTimespan := 0
	for i, notificationDate := range recentNotificationDates {
		if notificationDate.Add(in).After(time.Now()) {
			notificationsInTimespan++
		} else {
			// Remove old notifications from the slice
			recentNotificationDates = append(recentNotificationDates[:i], recentNotificationDates[i+1:]...)
		}
	}

	// Assign removed slice to store
	rl.datesStore[notifType][recipient] = recentNotificationDates

	return notificationsInTimespan
}

// AddNotificationDate adds a notification date to the store
func (rl *rateLimiter) AddNotificationDate(notification Notification) {
	nType, user := notification.Type, notification.Recipient

	// Create key if doesn't exist
	if rl.datesStore[nType] == nil {
		rl.datesStore[nType] = make(map[Username][]time.Time)
	}

	rl.datesStore[nType][user] = append(rl.datesStore[nType][user], notification.Date)
}
