package models

import (
	"fmt"
)

type Roles struct {
	ID        uint64
	Name      string
}

func (i Roles) String() string {
	return "Roles{" +
		"id: " + fmt.Sprint(i.ID) +
		", name: " + i.Name +
		"}"
}

