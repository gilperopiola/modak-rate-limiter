package main

import (
	"fmt"

	"github.com/gilperopiola/modak-rate-limiter/pkg"
)

func main() {
	rateLimiter := pkg.NewRateLimiter(pkg.DefaultRules)
	notificationService := pkg.NewNotificationService(rateLimiter)

	err := notificationService.Send(pkg.Notification{})
	if err != nil {
		fmt.Println(err.Error())
	}
}
