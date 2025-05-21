package handlers

import (
	"net/http"

	"github.com/likhithkp/practice/helpers"
	"github.com/likhithkp/practice/services"
)

func DeleteStudent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		helpers.JsonEncoder(w, "Not a valid method", http.StatusMethodNotAllowed, nil)
		return
	}

	id := r.PathValue("id")
	if id == "" {
		helpers.JsonEncoder(w, "Missing params", http.StatusBadRequest, nil)
		return
	}

	err := services.DeleteStudent(id)
	if err != nil {
		helpers.JsonEncoder(w, "Error while deleting the student", http.StatusInternalServerError, nil)
		return
	}

	helpers.JsonEncoder(w, "Student with id"+id+"deleted successfuly", http.StatusOK, nil)
}
