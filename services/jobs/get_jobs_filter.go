package jobs

import (
	"strings"
	"fmt"
	"reflect"
)

type JobFilter struct {
	Query         string
	hasAddedWhere bool
}

func MakeJobFilter() *JobFilter {
	return &JobFilter{
		Query:         "",
		hasAddedWhere: false,
	}
}

func (j *JobFilter) BasicJobsFilter() {
	j.Query = `SELECT
				j.id,
				j.name,
				j.employerId,
				j.jobTypeId,
				j.roleId,
				j.questionSetId,
				j.description,
				e.logo as logo,
				r.name AS role_name,
				e.name AS employer_name,
				i.id AS industry_id,
				i.name AS industry_name,
				jt.name AS job_type_name
			FROM Jobs AS j
			JOIN Roles AS r ON j.roleId = r.id
			JOIN Employers AS e ON j.employerId = e.id
			JOIN Industries AS i ON e.industryId = i.id
			JOIN JobTypes AS jt ON j.jobTypeId = jt.id
		`
}

func (q *JobFilter) BasicCountFilter() {
	q.Query = `SELECT COUNT(*)
			FROM Jobs AS j`
}

func (q *JobFilter) BasicFilteredCountFilter() {
	q.Query = `SELECT COUNT(*)
			FROM Jobs AS j
			JOIN Roles AS r ON j.roleId = r.id
			JOIN Employers AS e ON j.employerId = e.id
			JOIN Industries AS i ON e.industryId = i.id
			JOIN JobTypes AS jt ON j.jobTypeId = jt.id
		`
}

func (q *JobFilter) AddEmployersFilter(employerIds []int64) {
	if len(employerIds) == 0 {
		return
	}
	if !q.hasAddedWhere {
		q.Query += " WHERE "
		q.hasAddedWhere = true
		q.Query += "e.id IN (" + stringsArrayBuilder(employerIds) + ")"
	} else {
		q.Query += " AND e.id IN (" + stringsArrayBuilder(employerIds) + ")"
	}
}

func (q *JobFilter) AddRolesFilter(roleIds []int64) {
	if len(roleIds) == 0 {
		return
	}
	if !q.hasAddedWhere {
		q.Query += " WHERE "
		q.hasAddedWhere = true
		q.Query += "r.id IN (" + stringsArrayBuilder(roleIds) + ")"
	} else {
		q.Query += " AND r.id IN (" + stringsArrayBuilder(roleIds) + ")"
	}
}

func (q *JobFilter) AddIndustriesFilter(industryIds []int64) {
	if len(industryIds) == 0 {
		return
	}
	if !q.hasAddedWhere {
		q.Query += " WHERE "
		q.hasAddedWhere = true
		q.Query += "i.id IN (" + stringsArrayBuilder(industryIds) + ")"
	} else {
		q.Query += " AND i.id IN (" + stringsArrayBuilder(industryIds) + ")"
	}
}

func (q *JobFilter) AddJobTypesFilter(jobTypes []int64) {
	if len(jobTypes) == 0 {
		return
	}
	if !q.hasAddedWhere {
		q.Query += " WHERE "
		q.hasAddedWhere = true
		q.Query += "jt.id IN (" + stringsArrayBuilder(jobTypes) + ")"
	} else {
		q.Query += " AND jt.id IN (" + stringsArrayBuilder(jobTypes) + ")"
	}
}

func (q *JobFilter) AddLimitAndOffset(limit, offset int64) {
	q.Query += " LIMIT " + fmt.Sprint(limit) + " OFFSET " + fmt.Sprint(offset)
}

func makeString(s any, isString bool) string {
	if isString {
		return fmt.Sprintf("'%s'", s)
	}
	return fmt.Sprintf("%v", s)
}

func stringsArrayBuilder(arr interface{}) string {
	val := reflect.ValueOf(arr)
	if val.Kind() != reflect.Array && val.Kind() != reflect.Slice {
		return ""
	}

	// Convert the elements to a comma-separated string
	var strBuilder strings.Builder
	for i := 0; i < val.Len(); i++ {
		if i > 0 {
			strBuilder.WriteString(", ")
		}
		strBuilder.WriteString(makeString(val.Index(i).Interface(), val.Index(i).Kind() == reflect.String))
	}

	return strBuilder.String()
}