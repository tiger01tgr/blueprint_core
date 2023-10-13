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
		", createdAt: " + u.CreatedAt.String() +
		", lastLogin: " + u.LastLogin.String() +
		", deleted: " + fmt.Sprint(u.Deleted) +
		"}"
}
