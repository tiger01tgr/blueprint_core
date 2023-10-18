package models

import (
	"fmt"
	"time"
)

type User struct {
	ID         int64     // int8 non-nullable
	FirstName  string    // varchar non-nullable
	MiddleName string    // varchar nullable
	LastName   string    // varchar non-nullable
	Email      string    // varchar non-nullable
	UserType   string    // varchar non-nullable
	School 	   string	  // varchar nullable
	Major	   string	  // varchar nullable
	Employer   string	  // varchar nullable
	Position   string	  // varchar nullable
	Phone	   string	  // varchar nullable
	Resume	   string	  // varchar nullable
	CreatedAt  time.Time // timestamp with time zone non-nullable
	LastLogin  time.Time // timestamp with time zone non-nullable
	Deleted    bool      // boolean non-nullable
}

func (u User) String() string {
	return "User{" +
		"id: " + fmt.Sprint(u.ID) +
		", firstName: " + u.FirstName +
		", middleName: " + u.MiddleName +
		", lastName: " + u.LastName +
		", email: " + u.Email +
		", userType: " + u.UserType +
		", school: " + u.School +
		", major: " + u.Major +
		", employer: " + u.Employer +
		", position: " + u.Position +
		", phone: " + u.Phone +
		", resume: " + u.Resume +
		", createdAt: " + u.CreatedAt.String() +
		", lastLogin: " + u.LastLogin.String() +
		", deleted: " + fmt.Sprint(u.Deleted) +
		"}"
}
