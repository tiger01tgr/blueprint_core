package models

import (
	"fmt"
)

type QuestionSet struct {
	ID         uint64
	Name	   string
	EmployerId uint64
	RoleId       uint64
	InterviewType string
	CreatedAt string
	Deleted    bool
}

func (qs QuestionSet) String() string {
	return "QuestionSet{" +
		"id: " + fmt.Sprint(qs.ID) +
		", employerId: " + fmt.Sprint(qs.EmployerId) +
		", roleId: " + fmt.Sprint(qs.RoleId) +
		", interviewType: " + qs.InterviewType +
		", createdAt: " + qs.CreatedAt +
		", deleted: " + fmt.Sprint(qs.Deleted) +
		"}"
}