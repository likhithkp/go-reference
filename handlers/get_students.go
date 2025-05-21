package handlers

import (
	"net/http"

	"github.com/likhithkp/practice/helpers"
	"github.com/likhithkp/practice/services"
)

func GetAllStudents(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		helpers.JsonEncoder(w, "Not a valid method", http.StatusMethodNotAllowed, nil)
		return
	}

	students, err := services.GetAllStudents()
	if err != nil {
		helpers.JsonEncoder(w, "Error while getting the students", http.StatusInternalServerError, nil)
		return
	}

	helpers.JsonEncoder(w, "Students fetched", http.StatusOK, &students)
}
