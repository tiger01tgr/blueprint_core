package industries

import (
	"backend/db/models"
	"backend/db/dao/industries"
	"database/sql"
	"strconv"
)


func GetAllIndustries() ([]models.Industry, error) {
	rows, err := dao.GetIndustries()
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	var industries []models.Industry
	for rows.Next() {
		var i models.Industry
		if err := rows.Scan(&i.ID, &i.Name); err != nil {
			return nil, err
		}
		industries = append(industries, i)
	}
	if err := rows.Err(); err != nil {
        return nil, err
    }
	return industries, nil
}

func CreateIndustry(name string) error {
	_, err := dao.CreateIndustry(name)
	return err
}

func EditIndustry(id, name string) error {
	dbId, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	err = dao.UpdateIndustry(dbId, name)
	return err
}

func DeleteIndustry(id string) error {
	dbId, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	err = dao.DeleteIndustry(dbId)
	return err
}


func createIndustryHelper(row *sql.Row) (*models.Industry, error) {
	var i models.Industry
	err := row.Scan(
		&i.ID,
		&i.Name,
	)
	if err != nil {
		return nil, err
	}
	return &i, nil
}
