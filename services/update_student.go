package services

import (
	"context"
	"log"

	"github.com/likhithkp/practice/db"
	"github.com/likhithkp/practice/shared"
)

func UpdateStudent(updatedStudent *shared.Student, id string) error {
	const query = `UPDATE student SET
		name = COALESCE(NULLIF($1, ''), name),
		usn = COALESCE(NULLIF($2, ''), usn),
		branch = COALESCE(NULLIF($3, ''), branch),
		gpa = COALESCE(NULLIF($4, 0), gpa),
		birth_place = COALESCE(NULLIF($5, ''), birth_place)
		WHERE id = $6`

	cmdTag, err := db.DB.Exec(context.Background(), query,
		&updatedStudent.Name,
		&updatedStudent.USN,
		&updatedStudent.Branch,
		&updatedStudent.GPA,
		&updatedStudent.BirthPlace,
		&id,
	)

	if err != nil {
		log.Printf("Error while updating the student: %v", err)
		return err
	}

	if cmdTag.Update() {
		log.Println("Student updated successfully")
	}

	return nil
}
