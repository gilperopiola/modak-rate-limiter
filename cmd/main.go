package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/gilperopiola/modak-rate-limiter/pkg"
)

func main() {

	// Initialize rate limiter and notification service
	rateLimiter := pkg.NewRateLimiter(pkg.DefaultRules)
	notificationService := pkg.NewNotificationService(rateLimiter)

	// Send 15 notifications
	for i := 0; i < 15; i++ {
		notification := pkg.Notification{
			Type:      pkg.NotifSlice[rand.Intn(len(pkg.NotifSlice))],
			Recipient: pkg.Username("modak"),
			Message:   ":)",
			Date:      time.Now(),
		}

		err := notificationService.Send(notification)
		if err != nil {
			fmt.Println(err.Error() + "\n")
		}
	}

}
