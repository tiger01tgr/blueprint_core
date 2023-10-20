package models

type Job struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Employer string `json:"employer"`
	EmployerId int64 `json:"employer_id"`
	Role string `json:"role"`
	RoleId int64 `json:"role_id"`
	JobType string `json:"job_type"`
	JobTypeId int64 `json:"job_type_id"`
	Industry string `json:"industry"`
	IndustryId int64 `json:"industry_id"`
	Description string `json:"description"`
	QuestionSetId int64 `json:"question_set_id"`
	Logo string `json:"logo"`
}

type JobType struct {
	ID int64 `json:"id"`
	Name string `json:"name"`
}

