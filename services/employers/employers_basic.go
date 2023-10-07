package employers

import (
	"fmt"
	"strconv"
	"backend/services/s3"
	dao "backend/db/dao/employers"
	models "backend/db/models"
	"database/sql"
	"log"
	"mime/multipart"
)


func GetAllEmployers() ([]models.Employer, error) {
	rows, err := dao.GetAllEmployers()
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

    var employers []models.Employer
    for rows.Next() {
        var e models.Employer
        if err := rows.Scan(&e.ID, &e.Name, &e.Logo, &e.CreatedAt, &e.Deleted); err != nil {
            return nil, err
        }
        employers = append(employers, e)
    }

    if err := rows.Err(); err != nil {
        return nil, err
    }

    return employers, nil
}	

func GetAllEmployersFilter(companyIds, industries []string) {

}

func GetEmployerWithId(id string) (*models.Employer, error) {
	row, err := dao.GetEmployerByID(id)
	if err != nil {
		return nil, err
	}
	e, err := createEmployerHelper(row)
	if err != nil {
		return nil, err
	}
	return e, nil
}

func CreateEmployer(name, industryId string, logo multipart.File) error {

	if name == "" || industryId == "" {
		return fmt.Errorf("Invalid request payload")
	}

	industryIdInt, err := strconv.Atoi(industryId)
	if err != nil {
		return err
	}

	// upload logo to s3
	logoUrl, err := s3.UploadCompanyLogo(name, &logo)
	if err != nil {
		log.Println(err)
		return err
	}

	_, err = dao.CreateEmployer(name, logoUrl, industryIdInt)

	if err != nil {
		return err
	}

	return nil
}

func UpdateEmployer(id, name, logo, industry string) error {
	err := dao.UpdateEmployer(id, name, logo, industry)
	if err != nil {
		return err
	}
	return nil
}

func createEmployerHelper(row *sql.Row) (*models.Employer, error) {
	var e models.Employer
	err := row.Scan(
		&e.ID,
		&e.Name,
		&e.Logo,
		&e.Industry,
		&e.CreatedAt,
		&e.Deleted,
	)
	if err != nil {
		return nil, err
	}
	return &e, nil
}