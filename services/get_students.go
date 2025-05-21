package services

import (
	"context"
	"log"

	"github.com/likhithkp/practice/db"
	"github.com/likhithkp/practice/shared"
)

func GetAllStudents() (*[]shared.Student, error) {
	const query = `SELECT id, name, usn, branch, gpa, birth_place, created_at FROM student`
	rows, err := db.DB.Query(context.Background(), query)

	if err != nil {
		log.Printf("Error while getting/querying the rows: %v", err)
		return nil, err
	}

	var students []shared.Student
	for rows.Next() {
		var student shared.Student
		err := rows.Scan(&student.ID, &student.Name, &student.USN, &student.Branch, &student.GPA, &student.BirthPlace, &student.CreatedAt)
		if err != nil {
			log.Printf("Error while scanning the rows to `student`: %v", err)
			return nil, err
		}

		students = append(students, student)
	}

	defer rows.Close()
	return &students, nil
}
