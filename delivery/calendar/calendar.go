package calendar

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	d "main/delivery"
	e "main/domain/errors"
	m "main/domain/model"
	utils "main/domain/utils"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
)

type GoogleCalendar struct {
	tokenFile       string
	credentialsFile string
}

func NewGoogleCalendar(tokenFile string, credentialsFile string) d.CalendarInterface {
	return &GoogleCalendar{
		tokenFile:       tokenFile,
		credentialsFile: credentialsFile,
	}
}

// Retrieves a token from a local file.
func (gc *GoogleCalendar) tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, e.StacktraceError(err)
	}
	defer f.Close()
	tok := &oauth2.Token{}
	if err := json.NewDecoder(f).Decode(tok); err != nil {
		return nil, e.StacktraceError(err)
	}
	return tok, nil
}

// Retrieve a token, saves the token, then returns the generated client.
func (gc *GoogleCalendar) getClient(tokFile string, config *oauth2.Config) (*http.Client, error) {
	tok, err := gc.tokenFromFile(tokFile)
	if err != nil {
		return nil, e.StacktraceError(err)
	}
	tokenSource := config.TokenSource(context.Background(), tok)
	newToken, err := tokenSource.Token()
	if err != nil {
		return nil, e.StacktraceError(err)
	}
	if newToken.AccessToken != tok.AccessToken {
		if err := utils.SaveFile(tokFile, newToken); err != nil {
			return nil, e.StacktraceError(err)
		}
		log.Println("Saved new token:", newToken.AccessToken)
	}
	return config.Client(context.Background(), tok), nil
}

func (gc *GoogleCalendar) getCalendarServiceClient() (*calendar.Service, error) {
	ctx := context.Background()
	b, err := os.ReadFile(gc.credentialsFile)
	if err != nil {
		log.Println("Unable to read client secret file: ", err)
		return nil, e.StacktraceError(err)
	}

	config, err := google.ConfigFromJSON(b, calendar.CalendarScope)
	if err != nil {
		log.Println("Unable to parse client secret file to config: ", err)
		return nil, e.StacktraceError(err)
	}
	client, err := gc.getClient(gc.tokenFile, config)
	if err != nil {
		log.Println("Unable to get client from token: ", err)
		return nil, e.StacktraceError(err)
	}

	return calendar.NewService(ctx, option.WithHTTPClient(client))
}

func (gc *GoogleCalendar) CreateCalendar(teacherID int, desc string) (string, error) {
	srv, err := gc.getCalendarServiceClient()
	if err != nil {
		return "", e.StacktraceError(err)
	}

	newCal := &calendar.Calendar{TimeZone: "Europe/Moscow", Summary: desc}
	cal, err := srv.Calendars.Insert(newCal).Do()
	if err != nil {
		return "", e.StacktraceError(err)
	}

	acl := &calendar.AclRule{Scope: &calendar.AclRuleScope{Type: "default"}, Role: "reader"}
	_, err = srv.Acl.Insert(cal.Id, acl).Do()
	if err != nil {
		return "", e.StacktraceError(err)
	}

	return cal.Id, nil
}

func (gc *GoogleCalendar) CreateEvent(ev *m.CalendarEvent, calendarID string) (string, error) {
	srv, err := gc.getCalendarServiceClient()
	if err != nil {
		return "", e.StacktraceError(err)
	}

	event := &calendar.Event{
		Summary:     ev.Title + " Class " + fmt.Sprintf("%d", ev.ClassID),
		Description: ev.Description,
		Start: &calendar.EventDateTime{
			DateTime: ev.StartDate.Format(time.RFC3339),
			//TimeZone: "Europe/Moscow",
		},
		End: &calendar.EventDateTime{
			DateTime: ev.EndDate.Format(time.RFC3339),
			//TimeZone: "Europe/Moscow",
		},
		Visibility: "public",
	}

	event, err = srv.Events.Insert(calendarID, event).Do()
	if err != nil {
		return "", e.StacktraceError(err)
	}

	return event.Id, nil
}

func (gc *GoogleCalendar) DeleteEvent(calendarID string, eventID string) error {
	srv, err := gc.getCalendarServiceClient()
	if err != nil {
		return e.StacktraceError(err)
	}

	err = srv.Events.Delete(calendarID, eventID).Do()
	if err != nil {
		return e.StacktraceError(err)
	}

	return nil
}

func (gc *GoogleCalendar) GetEvents(teacherID int, calendarParams *m.CalendarParams) ([]m.CalendarEvent, error) {
	srv, err := gc.getCalendarServiceClient()
	if err != nil {
		log.Println("Unable to retrieve calendar Client: ", err)
		return nil, e.StacktraceError(err)
	}

	t := time.Now().Format(time.RFC3339)
	events, err := srv.Events.List(calendarParams.InternalApiID).
		ShowDeleted(false).
		SingleEvents(true).
		TimeMin(t).
		MaxResults(100).
		OrderBy("startTime").
		Do()
	if err != nil {
		log.Println("Unable to retrieve next ten of the user's events: ", err)
		return nil, e.StacktraceError(err)
	}

	res := []m.CalendarEvent{}
	for _, item := range events.Items {
		s := strings.Split(item.Summary, " ")
		clID := 0
		if len(s) > 2 && s[len(s)-2] == "Class" {
			clIDs := s[len(s)-1]
			clID, err = strconv.Atoi(clIDs)
			if err != nil {
				log.Println("error: ", err)
				return nil, e.StacktraceError(err)
			}
		}

		startDate, err := time.Parse(time.RFC3339, item.Start.DateTime)
		if err != nil {
			return nil, e.StacktraceError(err)
		}

		endDate, err := time.Parse(time.RFC3339, item.End.DateTime)
		if err != nil {
			return nil, e.StacktraceError(err)
		}

		tmp := m.CalendarEvent{
			Title:       item.Summary,
			Description: item.Description,
			StartDate:   startDate,
			EndDate:     endDate,
			ClassID:     clID,
			ID:          item.Id,
		}

		res = append(res, tmp)
	}

	return res, nil
}

func (gc *GoogleCalendar) UpdateEvent(ev *m.CalendarEvent, calendarID string) error {
	srv, err := gc.getCalendarServiceClient()
	if err != nil {
		return e.StacktraceError(err)
	}

	event := &calendar.Event{
		Summary:     ev.Title,
		Description: ev.Description,
		Start: &calendar.EventDateTime{
			DateTime: ev.StartDate.Format(time.RFC3339),
			//TimeZone: "Europe/Moscow",
		},
		End: &calendar.EventDateTime{
			DateTime: ev.EndDate.Format(time.RFC3339),
			//TimeZone: "Europe/Moscow",
		},
		Visibility: "public",
	}

	_, err = srv.Events.Update(calendarID, ev.ID, event).Do()
	if err != nil {
		return e.StacktraceError(err)
	}

	return nil
}
