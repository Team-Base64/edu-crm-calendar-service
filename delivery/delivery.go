package delivery

import (
	m "main/domain/model"
)

type CalendarInterface interface {
	GetEvents(teacherID int, calendarParams *m.CalendarParams) ([]m.CalendarEvent, error)
	CreateEvent(ev *m.CalendarEvent, calendarID string) (string, error)
	UpdateEvent(ev *m.CalendarEvent, calendarID string) error
	DeleteEvent(calendarID, eventID string) error
	CreateCalendar(teacherID int, desc string) (string, error)
}
