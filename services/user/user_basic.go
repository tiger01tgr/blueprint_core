package user

import (
	dao "backend/db/dao/users"
	models "backend/db/models"
	"backend/services/s3"
	"database/sql"
	"fmt"
	"mime/multipart"

	"github.com/google/uuid"
)

// will support pagination: should go by id?
func GetUsers() {

}

func GetUserWithId(id int64) (*models.User, error) {
	row, err := dao.ReadUserWithId(id)
	if err != nil {
		return nil, err
	}
	u, err := fillUserHelper(row)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func GetUserIdWithEmail(email string) (int64, error) {
	row, err := dao.ReadUserIdWithEmail(email)
	if err != nil {
		return 0, err
	}
	var id int64
	err = row.Scan(&id)
	fmt.Println("Serivces: ", id)
	return id, err
}

func GetUserProfile(id int64) (*models.User, error) {
	row, err := dao.ReadUserProfile(id)
	if err != nil {
		return nil, err
	}
	return fillUserProfileHelper(row)
}

func CreateUser(firstname string, middlename string, lastname string, email string, userType string) error {
	_, err := dao.CreateUser(firstname, middlename, lastname, email, userType)
	if err != nil {
		return err
	}
	return nil
}

func UpdateUserProfile(id int64, firstname, middlename, lastname, school, major, employer, position, phone string, resume *multipart.File) error {
	uuid := uuid.New()
	resumeUrl, err := s3.UploadResume(uuid.String(), resume)
	if err != nil {
		return err
	}
	row, err := dao.ReadUserProfile(id)
	if err != nil {
		return err
	}
	u, err := fillUserProfileHelper(row)
	if err != nil {
		return err
	}
	patchUserProfileModel(u, firstname, middlename, lastname, school, major, employer, position, phone, resumeUrl)
	_, err = dao.UpdateUserProfile(u)
	return err
}

func fillUserHelper(row *sql.Row) (*models.User, error) {
	var u models.User
	err := row.Scan(
		&u.ID,
		&u.FirstName,
		&u.MiddleName,
		&u.LastName,
		&u.Email,
		&u.UserType,
		&u.CreatedAt,
		&u.LastLogin,
		&u.Deleted,
	)
	return &u, err
}

func fillUserProfileHelper(row *sql.Row) (*models.User, error) {
	var u models.User
	var school sql.NullString
	var major sql.NullString
	var employer sql.NullString
	var position sql.NullString
	var phone sql.NullString
	var resume sql.NullString
	fmt.Println(row)
	err := row.Scan(
		&u.ID,
		&u.FirstName,
		&u.MiddleName,
		&u.LastName,
		&u.Email,
		&u.UserType,
		&school,
		&major,
		&employer,
		&position,
		&phone,
		&resume,
	)
	if school.Valid {
		u.School = school.String
	} else {
		u.School = ""
	}
	if major.Valid {
		u.Major = major.String
	} else {
		u.Major = ""
	}
	if employer.Valid {
		u.Employer = employer.String
	} else {
		u.Employer = ""
	}
	if position.Valid {
		u.Position = position.String
	} else {
		u.Position = ""
	}
	if phone.Valid {
		u.Phone = phone.String
	} else {
		u.Phone = ""
	}
	if resume.Valid {
		u.Resume = resume.String
	} else {
		u.Resume = ""
	}

	return &u, err
}

func patchUserProfileModel(u *models.User, firstname, middlename, lastname, school, major, employer, position, phone, resume string) {
	if firstname != "" {
		u.FirstName = firstname
	}
	if middlename != "" {
		u.MiddleName = middlename
	}
	if lastname != "" {
		u.LastName = lastname
	}
	if school != "" {
		u.School = school
	}
	if major != "" {
		u.Major = major
	}
	if employer != "" {
		u.Employer = employer
	}
	if position != "" {
		u.Position = position
	}
	if phone != "" {
		u.Phone = phone
	}
	if resume != "" {
		u.Resume = resume
	}
	fmt.Println(u)
}
