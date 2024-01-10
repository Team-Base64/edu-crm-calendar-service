package utils

import (
	"encoding/json"
	"log"
	"os"

	e "main/domain/errors"
)

func SaveFile(path string, token interface{}) error {
	log.Println("Saving file to: ", path)

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
