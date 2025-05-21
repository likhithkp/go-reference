package services

import (
	"context"
	"log"

	"github.com/likhithkp/practice/db"
	"github.com/likhithkp/practice/shared"
)

func GetStudent(id string) (*shared.Student, error) {
	const query = `SELECT id, name, usn, branch, gpa, birth_place, created_at FROM student WHERE id = $1`
	row := db.DB.QueryRow(context.Background(), query, id)

	var student shared.Student
	err := row.Scan(&student.ID, &student.Name, &student.USN, &student.Branch, &student.GPA, &student.BirthPlace, &student.CreatedAt)
	if err != nil {
		log.Printf("Error while scanning the student from row: %v", err)
		return nil, err
	}

	return &student, nil
}
