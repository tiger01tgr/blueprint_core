package models

import (
	"fmt"
	"strconv"
	"time"
)

type User struct {
	ID         uint      // int8 non-nullable
	FirstName  string    // varchar non-nullable
	MiddleName string    // varchar nullable
	LastName   string    // varchar non-nullable
	Email      string    // varchar non-nullable
	TypeOfUser string    // "user" | "admin" non-nullable
	CreatedAt  time.Time // timestamptz non-nullable
	LastLogin  time.Time // timestamptz non-nullable
	Deleted    bool      // bool non-nullable
}

func NewUser(firstName string, middleName string, lastName string, email string, typeOfUser string) User {
	return User{
		FirstName:  firstName,
		MiddleName: middleName,
		LastName:   lastName,
		Email:      email,
		TypeOfUser: typeOfUser,
		CreatedAt:  time.Now(),
		LastLogin:  time.Now(),
		Deleted:    false,
	}
}

func (u User) String() string {
	return "User{" +
		"ID: " + fmt.Sprint(u.ID) +
		", FirstName: " + u.FirstName +
		", MiddleName: " + u.MiddleName +
		", LastName: " + u.LastName +
		", Email: " + u.Email +
		", TypeOfUser: " + u.TypeOfUser +
		", CreatedAt: " + u.CreatedAt.String() +
		", LastLogin: " + u.LastLogin.String() +
		", Deleted: " + strconv.FormatBool(u.Deleted) +
		"}"
}