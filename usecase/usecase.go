package usecase

import (
	"context"
	"encoding/json"
	"log"
	rep "main/repository"
	"net/http"
	"os"
	"os/exec"

	e "main/domain/errors"
	"main/domain/model"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
)

type UsecaseInterface interface {
	getClient(tokFile string, config *oauth2.Config) (*http.Client, error)
	tokenFromFile(file string) (*oauth2.Token, error)
	saveToken(path string, token *oauth2.Token) error
	SetOAUTH2Token() error
	SaveOAUTH2Token(authCode string) error
	GetCalendarServiceClient() (*calendar.Service, error)
	GetCalendarDB(teacherID int) (*model.CalendarParams, error)
}

type Usecase struct {
	st              rep.StoreInterface
	tokenFile       string
	credentialsFile string
}

func NewUsecase(s rep.StoreInterface, tok string, cred string) UsecaseInterface {
	return &Usecase{
		st:              s,
		tokenFile:       tok,
		credentialsFile: cred,
	}
}

// Retrieve a token, saves the token, then returns the generated client.
func (uc *Usecase) getClient(tokFile string, config *oauth2.Config) (*http.Client, error) {
	tok, err := uc.tokenFromFile(tokFile)
	if err != nil {
		return nil, e.StacktraceError(err)
	}
	tokenSource := config.TokenSource(context.Background(), tok)
	newToken, err := tokenSource.Token()
	if err != nil {
		return nil, e.StacktraceError(err)
	}
	if newToken.AccessToken != tok.AccessToken {
		if err := uc.saveToken(tokFile, newToken); err != nil {
			return nil, e.StacktraceError(err)
		}
		log.Println("Saved new token:", newToken.AccessToken)
	}
	return config.Client(context.Background(), tok), nil
}

// Retrieves a token from a local file.
func (uc *Usecase) tokenFromFile(file string) (*oauth2.Token, error) {
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

// Saves a token to a file path.
func (uc *Usecase) saveToken(path string, token *oauth2.Token) error {
	log.Println("Saving credential file to: ", path)

	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return e.StacktraceError(err)
	}
	defer f.Close()

	if err := json.NewEncoder(f).Encode(token); err != nil {
		return e.StacktraceError(err)
	}

	return nil
}

func (uc *Usecase) SetOAUTH2Token() error {
	//ctx := context.Background()
	b, err := os.ReadFile(uc.credentialsFile)
	if err != nil {
		log.Println("Unable to read client secret file: ", err)
		return e.StacktraceError(err)
	}

	config, err := google.ConfigFromJSON(b, calendar.CalendarScope)
	if err != nil {
		log.Println("Unable to parse client secret file to config: ", err)
		return e.StacktraceError(err)
	}
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	//browser.OpenURL(authURL)
	exec.Command("xdg-open", authURL).Start()
	return nil
	// client, err := getClient(config)
	// if err != nil {
	// 	log.Println("Unable to get client from token: ", err)
	// 	return  e.StacktraceError(err)
	// }
}

func (uc *Usecase) SaveOAUTH2Token(authCode string) error {
	//ctx := context.Background()
	b, err := os.ReadFile(uc.credentialsFile)
	if err != nil {
		log.Println("Unable to read client secret file: ", err)
		return e.StacktraceError(err)
	}

	config, err := google.ConfigFromJSON(b, calendar.CalendarScope)
	if err != nil {
		log.Println("Unable to parse client secret file to config: ", err)
		return e.StacktraceError(err)
	}

	token, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		log.Println("Unable to retrieve token from web: ", err)
		return e.StacktraceError(err)
	}

	if err := uc.saveToken(uc.tokenFile, token); err != nil {
		return e.StacktraceError(err)
	}

	return nil
}

func (uc *Usecase) GetCalendarServiceClient() (*calendar.Service, error) {
	ctx := context.Background()
	b, err := os.ReadFile(uc.credentialsFile)
	if err != nil {
		log.Println("Unable to read client secret file: ", err)
		return nil, e.StacktraceError(err)
	}

	config, err := google.ConfigFromJSON(b, calendar.CalendarScope)
	if err != nil {
		log.Println("Unable to parse client secret file to config: ", err)
		return nil, e.StacktraceError(err)
	}
	client, err := uc.getClient(uc.tokenFile, config)
	if err != nil {
		log.Println("Unable to get client from token: ", err)
		return nil, e.StacktraceError(err)
	}

	return calendar.NewService(ctx, option.WithHTTPClient(client))

}

func (uc *Usecase) GetCalendarDB(teacherID int) (*model.CalendarParams, error) {
	return uc.st.GetCalendarDB(teacherID)
}
