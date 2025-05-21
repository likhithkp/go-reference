package handlers

import (
	"net/http"

	"github.com/likhithkp/practice/helpers"
	"github.com/likhithkp/practice/services"
	"github.com/likhithkp/practice/shared"
)

func AddStudent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		helpers.JsonEncoder(w, "Not a valid method", http.StatusMethodNotAllowed, nil)
		return
	}

	student := new(shared.Student)
	err := helpers.JsonDecoder(w, r, &student)
	if err != nil {
		helpers.JsonEncoder(w, "Internal server error", http.StatusInternalServerError, nil)
		return
	}

	isValid, err := helpers.ValidateStudent(student)
	if !isValid {
		helpers.JsonEncoder(w, err.Error(), http.StatusBadRequest, nil)
		return
	}

	err = services.AddStudent(student)
	if err != nil {
		helpers.JsonEncoder(w, "Error while inserting student", http.StatusInternalServerError, nil)
		return
	}

	helpers.JsonEncoder(w, "Student added", http.StatusOK, nil)
}
