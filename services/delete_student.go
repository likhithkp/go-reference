package services

import (
	"context"
	"log"

	"github.com/likhithkp/practice/db"
)

func DeleteStudent(id string) error {
	const query = `DELETE FROM student WHERE id = $1`
	cmdTag, err := db.DB.Exec(context.Background(), query, id)

	if err != nil {
		log.Printf("Error while executing the delete: %v", err)
		return err
	}

	if cmdTag.Delete() {
		log.Printf("Deleted student successfully")
		return nil
	}

	return nil
}
