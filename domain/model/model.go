package model

import "time"

type Error struct {
	Error interface{} `json:"error,omitempty"`
}

type Response struct {
	Body interface{} `json:"body,omitempty"`
}

type CalendarParams struct {
	ID            int    `json:"id"`
	InternalApiID string `json:"googleid"`
}

type CalendarEvent struct {
	Title       string    `json:"title"`
	Description string    `json:"description,omitempty"`
	StartDate   time.Time `json:"startDate"`
	EndDate     time.Time `json:"endDate"`
	ClassID     int       `json:"classid"`
	ID          string    `json:"id,omitempty"`
}

type CalendarEvents struct {
	Events []CalendarEvent `json:"events"`
}
