package usecase

import (
	m "main/domain/model"
)

type UsecaseInterface interface {
	GetEvents(teacherID int) ([]m.CalendarEvent, error)
	CreateEvent(ev *m.CalendarEvent, calendarID string) (string, error)
	DeleteEvent(calendarID string, eventID string) error
	UpdateEvent(ev *m.CalendarEvent, calendarID string) error
	CreateCalendar(teacherID int) error
}
