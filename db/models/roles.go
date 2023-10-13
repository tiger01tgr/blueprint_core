package models

import (
	"fmt"
)

type Roles struct {
	ID   int64
	Name string
}

func (i Roles) String() string {
	return "Roles{" +
		"id: " + fmt.Sprint(i.ID) +
		", name: " + i.Name +
		"}"
}
