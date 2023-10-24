package pkg

import (
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
