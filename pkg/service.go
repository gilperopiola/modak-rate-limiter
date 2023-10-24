package pkg

import (
	"fmt"
)

func NewNotificationService(limiter RateLimiter) *notificationService {
	return &notificationService{
		limiter: limiter,
	}
}

type NotificationService interface {
	Send(notification Notification) error
}

type notificationService struct {
	limiter RateLimiter
}

// Send first checks the rate limiter to see if the quota is ok, then it sends the notification and adds it to the store
func (ns *notificationService) Send(notification Notification) error {

	// Check if quota exceeded for user & notifType
	exceeded, err := ns.limiter.HasExceededQuota(notification)
	if err != nil {
		return err
	}

	if exceeded {
		return fmt.Errorf("quota exceeded for user %s and notification type %s", notification.Recipient, notification.Type)
	}

	// Send notification
	fmt.Printf("%s notification sent to user %s", notification.Type, notification.Recipient)

	// Add notification date to rate limiter store
	ns.limiter.AddNotificationDate(notification)

	return nil
}
