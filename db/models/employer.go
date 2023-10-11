package models

import (
	"fmt"
	"time"
)

type Employer struct {
	ID        uint64
	Name      string
	Logo      string
	Industry  string
	IndustryId uint64
	CreatedAt time.Time
	Deleted   bool
}

func (e Employer) String() string {
	return "Employer{" +
		"id: " + fmt.Sprint(e.ID) +
		", name: " + e.Name +
		", logo: " + e.Logo +
		", industry: " + e.Industry +
		", industryId: " + fmt.Sprint(e.IndustryId) +
		", created_at: " + e.CreatedAt.String() +
		", deleted: " + fmt.Sprint(e.Deleted) +
		"}"
}

