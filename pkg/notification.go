package pkg

import (
	"math/rand"
	"time"
)

type Username string
type NotifType string

var (
	NotifTypeStatus NotifType = "Status"
	NotifTypeNews   NotifType = "News"
	NotifTypeMarket NotifType = "Marketing"
	NotifSlice                = []NotifType{
		NotifTypeStatus, NotifTypeNews, NotifTypeMarket,
	}
)

type Notification struct {
	Type      NotifType
	Recipient Username
	Message   string
	Date      time.Time
}

func NewRandomNotification(username Username) Notification {
	return Notification{
		Type:      NotifSlice[rand.Intn(len(NotifSlice))],
		Recipient: username,
		Message:   ":)",
		Date:      time.Now(),
	}
}
