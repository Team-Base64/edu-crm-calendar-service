package calendar

import (
	"context"
	"log"
	"time"

	proto "main/delivery/grpc/calendar/proto"
	e "main/domain/errors"
	m "main/domain/model"
	uc "main/usecase"
)

type CalendarGrpcHandler struct {
	proto.UnimplementedCalendarServer
	uc uc.UsecaseInterface
}

func NewCalendarGrpcHandler(uc uc.UsecaseInterface) proto.CalendarServer {
	return &CalendarGrpcHandler{
		uc: uc,
	}
}

func (cs *CalendarGrpcHandler) GetEventsCalendar(ctx context.Context, req *proto.GetEventsRequestCalendar) (*proto.GetEventsResponse, error) {
	log.Println("Called GetEventsCalendar")
	events, err := cs.uc.GetEvents(int(req.GetTeacherID()))
	if err != nil {
		log.Println(e.StacktraceError(err))
		return &proto.GetEventsResponse{Events: nil}, err
	}

	protoEvents := []*proto.EventData{}
	for _, event := range events {
		protoEvents = append(protoEvents, &proto.EventData{
			Id:          event.ID,
			Title:       event.Title,
			Description: event.Description,
			StartDate:   event.StartDate.Format(time.RFC3339),
			EndDate:     event.EndDate.Format(time.RFC3339),
			ClassID:     int32(event.ClassID),
		})
	}

	return &proto.GetEventsResponse{Events: protoEvents}, nil
}

func (cs *CalendarGrpcHandler) CreateEvent(ctx context.Context, req *proto.CreateEventRequest) (*proto.CreateEventResponse, error) {
	log.Println("Called CreateEvent")

	startDate, err := time.Parse(time.RFC3339, req.Event.StartDate)
	if err != nil {
		log.Println(e.StacktraceError(err))
		return &proto.CreateEventResponse{EventID: ""}, err
	}

	endDate, err := time.Parse(time.RFC3339, req.Event.EndDate)
	if err != nil {
		log.Println(e.StacktraceError(err))
		return &proto.CreateEventResponse{EventID: ""}, err
	}

	event := m.CalendarEvent{
		Title:       req.Event.Title,
		Description: req.Event.Description,
		StartDate:   startDate,
		EndDate:     endDate,
		ClassID:     int(req.Event.ClassID),
		ID:          req.Event.Id,
	}
	eventID, err := cs.uc.CreateEvent(&event, req.CalendarID)
	if err != nil {
		log.Println(e.StacktraceError(err))
		return &proto.CreateEventResponse{EventID: ""}, err
	}

	return &proto.CreateEventResponse{EventID: eventID}, nil
}

func (cs *CalendarGrpcHandler) DeleteEvent(ctx context.Context, req *proto.DeleteEventRequest) (*proto.Nothing, error) {
	log.Println("Called DeleteEvent")

	if err := cs.uc.DeleteEvent(req.CalendarID, req.Id); err != nil {
		log.Println(e.StacktraceError(err))
		return &proto.Nothing{}, err
	}

	return &proto.Nothing{}, nil
}

func (cs *CalendarGrpcHandler) UpdateEvent(ctx context.Context, req *proto.UpdateEventRequest) (*proto.Nothing, error) {
	log.Println("Called UpdateEvent")

	startDate, err := time.Parse(time.RFC3339, req.Event.StartDate)
	if err != nil {
		log.Println(e.StacktraceError(err))
		return &proto.Nothing{}, err
	}

	endDate, err := time.Parse(time.RFC3339, req.Event.EndDate)
	if err != nil {
		log.Println(e.StacktraceError(err))
		return &proto.Nothing{}, err
	}

	event := m.CalendarEvent{
		Title:       req.Event.Title,
		Description: req.Event.Description,
		StartDate:   startDate,
		EndDate:     endDate,
		ClassID:     int(req.Event.ClassID),
		ID:          req.Event.Id,
	}

	if err := cs.uc.UpdateEvent(&event, req.CalendarID); err != nil {
		log.Println(e.StacktraceError(err))
		return &proto.Nothing{}, err
	}

	return &proto.Nothing{}, nil
}

func (cs *CalendarGrpcHandler) CreateCalendar(ctx context.Context, req *proto.CreateCalendarRequest) (*proto.Nothing, error) {
	log.Println("Called CreateCalendar")

	if err := cs.uc.CreateCalendar(int(req.TeacherID)); err != nil {
		log.Println(e.StacktraceError(err))
		return &proto.Nothing{}, err
	}
	return &proto.Nothing{}, nil
}
