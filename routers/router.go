package routers

import (
	"net/http"

	"github.com/likhithkp/practice/handlers"
)

func Router(mux *http.ServeMux) {
	mux.HandleFunc("/addStudent", handlers.AddStudent)
	mux.HandleFunc("/getAllStudents", handlers.GetAllStudents)
	mux.HandleFunc("/getStudent/{id}", handlers.GetStudent)
	mux.HandleFunc("/deleteStudent/{id}", handlers.DeleteStudent)
	mux.HandleFunc("/updateStudent/{id}", handlers.UpdateStudent)
}
