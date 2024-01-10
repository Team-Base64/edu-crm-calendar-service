package main

// go run helper/gen-token/main.go credentials.json token.json

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"

	e "main/domain/errors"
	m "main/domain/model"
	utils "main/domain/utils"

	"github.com/gorilla/mux"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
)

var tokenFile string
var credentialsFile string

var BaseUrl = "/api"
var PathOAuthSetToken = BaseUrl + "/oauth"
var PathOAuthSaveToken = BaseUrl + "/oauth/savetoken"
var Port = ":8084"

func SetOAUTH2Token() error {
	b, err := os.ReadFile(credentialsFile)
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

	exec.Command("xdg-open", authURL).Start()
	return nil
}

func SaveOAUTH2Token(authCode string) error {
	b, err := os.ReadFile(credentialsFile)
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

	if err := utils.SaveFile(tokenFile, token); err != nil {
		return e.StacktraceError(err)
	}

	return nil
}

func ReturnErrorJSON(w http.ResponseWriter, err error) {
	errCode, errText := e.CheckError(err)
	w.WriteHeader(errCode)
	json.NewEncoder(w).Encode(&m.Error{Error: errText})
}

func SetOAUTH2TokenHandler(w http.ResponseWriter, r *http.Request) {
	err := SetOAUTH2Token()
	if err != nil {
		log.Println(e.StacktraceError(err))
		ReturnErrorJSON(w, e.ErrServerError500)
		return
	}
	json.NewEncoder(w).Encode(&m.Response{})
}

func SaveOAUTH2TokenToFileHandler(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	err := SaveOAUTH2Token(code)
	if err != nil {
		log.Println(e.StacktraceError(err))
		ReturnErrorJSON(w, e.ErrServerError500)
		return
	}
	json.NewEncoder(w).Encode(&m.Response{})
}

func main() {
	credentialsFile = os.Args[1]
	tokenFile = os.Args[2]

	router := mux.NewRouter()
	router.HandleFunc(PathOAuthSetToken, SetOAUTH2TokenHandler).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc(PathOAuthSaveToken, SaveOAUTH2TokenToFileHandler).Methods(http.MethodGet, http.MethodOptions)

	fmt.Println("http://127.0.0.1" + Port + PathOAuthSetToken)
	err := http.ListenAndServe(Port, router)
	if err != nil {
		fmt.Println("cant serve", err)
	}
}
