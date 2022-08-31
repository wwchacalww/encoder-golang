package domain

import "time"

type Video struct {
	ID         string
	ResourceID string
	filepath   string
	createdAt  time.Time
}
