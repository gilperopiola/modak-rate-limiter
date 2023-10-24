package pkg

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// Might have done this with table-driven tests, I just thought this way it would be easier
func TestRateLimiter(t *testing.T) {
	rl := NewRateLimiter(DefaultRules)

	// First, test wrong notification type
	n := Notification{
		Type: "Invalid", Recipient: "modak",
		Message: ":)", Date: time.Now(),
	}

	exceededQuota, err := rl.HasExceededQuota(n)
	assert.True(t, exceededQuota)
	assert.Error(t, err)

	// Then, try with Status notification type for user modak
	n.Type = NotifTypeStatus
	exceededQuota, err = rl.HasExceededQuota(n)
	assert.False(t, exceededQuota)
	assert.NoError(t, err)

	// Let's add the notification date to the store and try again
	n.Date = time.Now().Add(-20 * time.Second)
	rl.AddNotificationDate(n)

	exceededQuota, err = rl.HasExceededQuota(n)
	assert.False(t, exceededQuota)
	assert.NoError(t, err)

	// Perfect, but with 1 more notification added now it should exceed the quota
	rl.AddNotificationDate(n)

	exceededQuota, err = rl.HasExceededQuota(n)
	assert.True(t, exceededQuota)
	assert.NoError(t, err)

	// Now try but with another user
	n.Recipient = "kadom"
	exceededQuota, err = rl.HasExceededQuota(n)
	assert.False(t, exceededQuota)
	assert.NoError(t, err)

	// And finally test that the old notifications are removed from the store
	rl = NewRateLimiter(DefaultRules)
	n.Date = time.Now().Add(-20 * time.Hour)
	rl.AddNotificationDate(n)
	assert.Equal(t, 1, len(rl.datesStore[n.Type][n.Recipient]))

	exceededQuota, err = rl.HasExceededQuota(n)
	assert.False(t, exceededQuota)
	assert.NoError(t, err)
	assert.Equal(t, 0, len(rl.datesStore[n.Type][n.Recipient]))
}
