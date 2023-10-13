package models

import (
	"fmt"
)

type Industry struct {
	ID   int64
	Name string
}

func (i Industry) String() string {
	return "Industry{" +
		"id: " + fmt.Sprint(i.ID) +
		", name: " + i.Name +
		"}"
}
