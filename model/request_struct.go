package model

type Request struct {
	AssignmentID string `json:"assignment_id"`
	SchoolID     string `json:"school_id"`
	ClassID      string `json:"class_id"`
	SubjectID    string `json:"subject_id"`
}
