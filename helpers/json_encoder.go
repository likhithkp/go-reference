package helpers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/likhithkp/practice/shared"
)

func JsonEncoder(w http.ResponseWriter, message string, statusCode uint, data any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(int(statusCode))

	err := json.NewEncoder(w).Encode(&shared.Response{
		Message:    message,
		StatusCode: statusCode,
		Data:       data,
	})

	if err != nil {
		log.Println("Error while encoding", err.Error())
		err := errors.New(err.Error())
		return err
	}

	return nil
}
