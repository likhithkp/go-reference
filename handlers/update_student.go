package handlers

import (
	"net/http"

	"github.com/likhithkp/practice/helpers"
	"github.com/likhithkp/practice/services"
	"github.com/likhithkp/practice/shared"
)

func UpdateStudent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPatch {
		helpers.JsonEncoder(w, "Not a valid method", http.StatusMethodNotAllowed, nil)
		return
	}

	id := r.PathValue("id")
	if id == "" {
		helpers.JsonEncoder(w, "Missing params", http.StatusBadRequest, nil)
		return
	}

	updatedStudent := new(shared.Student)
	helpers.JsonDecoder(w, r, &updatedStudent)

	err := services.UpdateStudent(updatedStudent, id)
	if err != nil {
		helpers.JsonEncoder(w, "Error while updating student", http.StatusInternalServerError, nil)
		return
	}

	helpers.JsonEncoder(w, "Student with id "+id+" Updated successfully", http.StatusOK, nil)
}
