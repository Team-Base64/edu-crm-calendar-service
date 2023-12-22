package chat

import (
	"encoding/json"
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

// @host 127.0.0.1:8083
// @BasePath  /

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
