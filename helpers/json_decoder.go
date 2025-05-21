package helpers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

func JsonDecoder(w http.ResponseWriter, r *http.Request, decodeVariable any) error {
	if err := json.NewDecoder(r.Body).Decode(decodeVariable); err != nil {
		err := errors.New(err.Error())
		log.Println("Error while decoding", err)
		return err
	}

	return nil
}
