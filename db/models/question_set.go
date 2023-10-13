package models

import (
	"fmt"
)

type QuestionSet struct {
	ID            int64
	Name          string
	EmployerId    int64
	RoleId        int64
	InterviewType string
	CreatedAt     string
	Deleted       bool

	// Optional fields
	Logo         string
	EmployerName string
	RoleName     string
	IndustryName string
	IndustryId   int64
	NumQuestions int64
}

func NewQuestionSet(name string, employerId int64, roleId int64, interviewType string) QuestionSet {
	return QuestionSet{
		Name:          name,
		EmployerId:    employerId,
		RoleId:        roleId,
		InterviewType: interviewType,
	}
}

func (qs QuestionSet) WithEmployerName(name string) QuestionSet {
	qs.EmployerName = name
	return qs
}

func (qs QuestionSet) WithRoleName(name string) QuestionSet {
	qs.RoleName = name
	return qs
}

func (qs QuestionSet) WithIndustryName(name string) QuestionSet {
	qs.IndustryName = name
	return qs
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
