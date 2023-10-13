package industries

import (
	"backend/db/models"
	"backend/db/dao/roles"
	"database/sql"
	"strconv"
)


func GetAllRoles() ([]models.Roles, error) {
	rows, err := dao.GetRoles()
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	var roles []models.Roles
	for rows.Next() {
		var i models.Roles
		if err := rows.Scan(&i.ID, &i.Name); err != nil {
			return nil, err
		}
		roles = append(roles, i)
	}
	if err := rows.Err(); err != nil {
        return nil, err
    }
	return roles, nil
}

func CreateRole(name string) error {
	_, err := dao.CreateRole(name)
	return err
}

func EditRole(id, name string) error {
	dbId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	err = dao.UpdateRole(dbId, name)
	return err
}

func DeleteRole(id string) error {
	dbId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	err = dao.DeleteRole(dbId)
	return err
}


func createRoleHelper(row *sql.Row) (*models.Roles, error) {
	var i models.Roles
	err := row.Scan(
		&i.ID,
		&i.Name,
	)
	if err != nil {
		return nil, err
	}
	return &i, nil
}
