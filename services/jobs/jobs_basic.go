package jobs

import (
	custom_dao "backend/db/dao/custom"
	dao "backend/db/dao/jobs"
	"backend/db/models"
	"fmt"
)

func CreateJob(name, employerId, jobTypeId, roleId, questionSetId, description string) error {
	return dao.CreateJob(name, employerId, jobTypeId, roleId, questionSetId, description)
}

func CreateJobType(name string) error {
	return dao.CreateJobType(name)
}

func GetJobs(employers []int64, industries []int64, roles []int64, jobTypes []int64, page int64, limit int64, shouldFilter bool) ([]models.Job, error) {
	offset := (page - 1) * limit
	jobFilter := MakeJobFilter()
	jobFilter.BasicJobsFilter()
	if shouldFilter {
		fmt.Println("Filtering")
		jobFilter.AddEmployersFilter(employers)
		jobFilter.AddIndustriesFilter(industries)
		jobFilter.AddRolesFilter(roles)
		jobFilter.AddJobTypesFilter(jobTypes)
	}
	jobFilter.AddLimitAndOffset(limit, offset)
	rows, err := custom_dao.CustomQueryForRows(jobFilter.Query)
	defer rows.Close()
	if err != nil {
		fmt.Println("err: ", err)
		return nil, err
	}
	var jobs []models.Job
	for rows.Next() {
		var j models.Job
		if err := rows.Scan(&j.ID, &j.Name, &j.EmployerId, &j.JobTypeId, &j.RoleId, &j.QuestionSetId, &j.Description, &j.Logo, &j.Role, &j.Employer, &j.IndustryId, &j.Industry, &j.JobType); err != nil {
			return nil, err
		}
		jobs = append(jobs, j)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return jobs, nil
}

func GetPaginationForJobs(employers, industries, roles, jobTypes []int64, limit int64, shouldFilter bool) (int64, error) {
	jobFilter := MakeJobFilter()
	if shouldFilter {
		jobFilter.BasicFilteredCountFilter()
		jobFilter.AddEmployersFilter(employers)
		jobFilter.AddIndustriesFilter(industries)
		jobFilter.AddRolesFilter(roles)
		jobFilter.AddJobTypesFilter(jobTypes)
	} else {
		jobFilter.BasicCountFilter()
	}
	row, err := custom_dao.CustomQueryForRow(jobFilter.Query)
	if err != nil {
		fmt.Println("err: ", err)
		return 0, err
	}
	var count int64
	row.Scan(&count)
	return (count / limit) + 1, nil
}

func GetJobTypes() ([]models.JobType, error) {
	rows, err := dao.GetJobTypes()
	defer rows.Close()
	if err != nil {
		fmt.Println("err: ", err)
		return nil, err
	}
	var jobTypes []models.JobType
	for rows.Next() {
		var jt models.JobType
		if err := rows.Scan(&jt.ID, &jt.Name); err != nil {
			return nil, err
		}
		jobTypes = append(jobTypes, jt)
	}
	return jobTypes, nil
}