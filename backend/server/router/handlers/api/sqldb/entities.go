package sqldb

//Student struct
type Student struct {
	ID          int    `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Email       string `json:"email,omitempty"`
	CreatedTime string `json:"CreatedTime,omitempty"`
}

// StudentS is a slice of Student
type StudentS struct {
	StudentS []Student `json:"students"`
}
