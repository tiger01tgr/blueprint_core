package models

import (
	"fmt"
)

type Industry struct {
	ID        uint64
	Name      string
}

func (i Industry) String() string {
	return "Industry{" +
		"id: " + fmt.Sprint(i.ID) +
		", name: " + i.Name +
		"}"
}

