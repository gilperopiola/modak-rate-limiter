package pkg

import (
	"time"
)

type Username string
type NotifType string

const (
	NotifTypeStatus NotifType = "Status"
	NotifTypeNews   NotifType = "News"
	NotifTypeMarket NotifType = "Marketing"
)

type Notification struct {
	Type      NotifType
	Recipient Username
	Message   string
	Date      time.Time
}
