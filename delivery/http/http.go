package http

import (
	"encoding/json"
	"log"
	"net/http"

	e "main/domain/errors"
	"main/domain/model"
	uc "main/usecase"
)

// @title TCRA API
// @version 1.0
// @description EDUCRM back calendar server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 127.0.0.1:8084
// @BasePath  /api

type Handler struct {
	uc uc.UsecaseInterface
}

func NewHandler(uc uc.UsecaseInterface) *Handler {
	return &Handler{
		uc: uc,
	}
}

func ReturnErrorJSON(w http.ResponseWriter, err error) {
	errCode, errText := e.CheckError(err)
	w.WriteHeader(errCode)
	json.NewEncoder(w).Encode(&model.Error{Error: errText})
}

// SetOAUTH2Token godoc
// @Summary Sets teacher's OAUTH2Token
// @Description Sets teacher's OAUTH2Token
// @ID SetOAUTH2Token
// @Accept  json
// @Produce  json
// @Tags Calendar
// @Success 200 {object} model.Response
// @Failure 500 {object} model.Error "internal server error - Request is valid but operation failed at server side"
// @Router /oauth [post]
func (api *Handler) SetOAUTH2Token(w http.ResponseWriter, r *http.Request) {
	err := api.uc.SetOAUTH2Token()
	if err != nil {
		log.Println(e.StacktraceError(err))
		ReturnErrorJSON(w, e.ErrServerError500)
		return
	}
	json.NewEncoder(w).Encode(&model.Response{})
}

func (api *Handler) SaveOAUTH2TokenToFile(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	err := api.uc.SaveOAUTH2Token(code)
	if err != nil {
		log.Println(e.StacktraceError(err))
		ReturnErrorJSON(w, e.ErrServerError500)
		return
	}
	json.NewEncoder(w).Encode(&model.Response{})
}
