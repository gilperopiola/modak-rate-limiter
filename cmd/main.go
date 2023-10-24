package main

import (
	"fmt"

	"github.com/gilperopiola/modak-rate-limiter/pkg"
)

func main() {

	// Initialize rate limiter and notification service
	rateLimiter := pkg.NewRateLimiter(pkg.DefaultRules)
	notificationService := pkg.NewNotificationService(rateLimiter)
	sendTestingNotifications(notificationService)
}

func sendTestingNotifications(notificationService pkg.NotificationService) {

	// Send 15 notifications to 2 users
	for i := 0; i < 15; i++ {
		// User 1
		notification := pkg.NewRandomNotification("modak")
		if err := notificationService.Send(notification); err != nil {
			fmt.Println(err.Error() + "\n")
		}

		// User 2
		notification2 := pkg.NewRandomNotification("kadom")
		if err := notificationService.Send(notification2); err != nil {
			fmt.Println(err.Error() + "\n")
		}
	}
}
