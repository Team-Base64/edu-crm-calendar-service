package calendar

import (
	d "main/delivery"
	e "main/domain/errors"
	m "main/domain/model"
	rep "main/repository"
	uc "main/usecase"
)

type CalendarUsecase struct {
	store    rep.StoreInterface
	calendar d.CalendarInterface
}

func NewCalendarUsecase(s rep.StoreInterface, c d.CalendarInterface) uc.UsecaseInterface {
	return &CalendarUsecase{
		store:    s,
		calendar: c,
	}
}

func (uc *CalendarUsecase) GetEvents(teacherID int) ([]m.CalendarEvent, error) {
	calendarParams, err := uc.store.GetCalendar(teacherID)
	if err != nil {
		return nil, e.StacktraceError(err)
	}

	events, err := uc.calendar.GetEvents(teacherID, calendarParams)
	if err != nil {
		return nil, e.StacktraceError(err)
	}
	return events, nil
}

func (uc *CalendarUsecase) CreateEvent(ev *m.CalendarEvent, calendarID string) (string, error) {
	eventID, err := uc.calendar.CreateEvent(ev, calendarID)
	if err != nil {
		return "", e.StacktraceError(err)
	}
	return eventID, nil
}

func (uc *CalendarUsecase) DeleteEvent(internalApiID string, eventID string) error {
	if err := uc.calendar.DeleteEvent(internalApiID, eventID); err != nil {
		return e.StacktraceError(err)
	}
	return nil
}

func (uc *CalendarUsecase) UpdateEvent(ev *m.CalendarEvent, calendarID string) error {
	if err := uc.calendar.UpdateEvent(ev, calendarID); err != nil {
		return e.StacktraceError(err)
	}
	return nil
}

func (uc *CalendarUsecase) CreateCalendar(teacherID int) error {
	calendarID, err := uc.calendar.CreateCalendar(teacherID, "EDUCRM Calendar")
	if err != nil {
		return e.StacktraceError(err)
	}

	_, err = uc.store.AddCalendar(teacherID, calendarID)
	if err != nil {
		return e.StacktraceError(err)
	}

	return nil
}
