package admin

import (
	dao "backend/db/dao/admin"
	"backend/services/jwt"
	"errors"
)

type SuperAdmin struct {
	ID       int64
	Email    string
	Password string
}

func LoginSuperAdmin(email, password string) (string, error) {
	row, err := dao.GetSuperAdminByEmail(email)
	if err != nil {
		return "", err
	}
	var sa SuperAdmin
	err = row.Scan(&sa.ID, &sa.Email, &sa.Password)
	if err != nil {
		return "", err
	}
	if sa.Password == password {
		_, tokenString, _ := jwt.GetJWT().Encode(map[string]interface{}{"email": email})
		return tokenString, nil
	}
	return "", errors.New("invalid password")
}
