package shared

import "time"

type Student struct {
	ID         string    `json:"id"`
	Name       string    `json:"name"`
	USN        string    `json:"usn"`
	Branch     string    `json:"branch"`
	GPA        float64   `json:"gpa"`
	BirthPlace string    `json:"birth_place"`
	CreatedAt  time.Time `json:"created_at"`
}
