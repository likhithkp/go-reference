package helpers

import (
	"errors"

	"github.com/likhithkp/practice/shared"
)

func ValidateStudent(student *shared.Student) (bool, error) {
	if student.Name == "" || student.USN == "" || student.Branch == "" || student.BirthPlace == "" {
		err := errors.New("missing name/usn/branch/birthplace")
		return false, err
	}

	if student.GPA == 0 || student.GPA > 10 || student.GPA < 0 {
		err := errors.New("GPC must be between 0 and 10")
		return false, err
	}

	return true, nil
}
