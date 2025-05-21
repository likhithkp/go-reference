package services

import (
	"context"
	"log"

	"github.com/likhithkp/practice/db"
	"github.com/likhithkp/practice/shared"
)

func AddStudent(student *shared.Student) error {
	const query = `INSERT INTO student ( name, usn, branch, gpa, birth_place) VALUES ($1, $2, $3, $4, $5)`
	cmdTag, err := db.DB.Exec(context.Background(), query, student.Name, student.USN, student.Branch, student.GPA, student.BirthPlace)

	if err != nil {
		log.Printf("Error while inserting: %v", err)
		return err
	}

	if cmdTag.Insert() {
		log.Println("Added new student")
	}

	return nil
}
