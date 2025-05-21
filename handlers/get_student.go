package handlers

import (
	"net/http"

	"github.com/likhithkp/practice/helpers"
	"github.com/likhithkp/practice/services"
)

func GetStudent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		helpers.JsonEncoder(w, "Not a valid method", http.StatusMethodNotAllowed, nil)
		return
	}

	id := r.PathValue("id")
	student, err := services.GetStudent(id)

	if err != nil {
		helpers.JsonEncoder(w, "Error while getting user by id", http.StatusInternalServerError, nil)
		return
	}

	helpers.JsonEncoder(w, "Student fetched by id", http.StatusOK, &student)
}
