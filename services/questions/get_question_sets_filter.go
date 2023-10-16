package questions

import (
	"strings"
	"fmt"
	"reflect"
)

type QuestionSetFilter struct {
	Query         string
	hasAddedWhere bool
}

func CreateQuestionSetFilter() *QuestionSetFilter {
	return &QuestionSetFilter{
		Query:         "",
		hasAddedWhere: false,
	}
}

func (q *QuestionSetFilter) CreateGetBasicQuery() {
	q.Query = `SELECT
				qs.id,
				qs.name,
				qs.interviewType,
				qs.employerId,
				qs.roleId,
				qs.created_at,
				qs.deleted,
				e.logo as logo,
				r.name AS role_name,
				e.name AS employer_name,
				i.id AS industry_id,
				i.name AS industry_name
			FROM QuestionSets AS qs
			JOIN Roles AS r ON qs.roleId = r.id
			JOIN Employers AS e ON qs.employerId = e.id
			JOIN Industries AS i ON e.industryId = i.id
		`
}

func (q *QuestionSetFilter) CreateCountBasicQuery() {
	q.Query = `SELECT COUNT(*)
			FROM QuestionSets AS qs
			JOIN Roles AS r ON qs.roleId = r.id
			JOIN Employers AS e ON qs.employerId = e.id
			JOIN Industries AS i ON e.industryId = i.id
		`
}

func (q *QuestionSetFilter) AddEmployersFilter(employerIds []int64) {
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

func (q *QuestionSetFilter) AddRolesFilter(roleIds []int64) {
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

func (q *QuestionSetFilter) AddIndustriesFilter(industryIds []int64) {
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

func (q *QuestionSetFilter) AddInterviewTypesFilter(interviewTypes []string) {
	if len(interviewTypes) == 0 {
		return
	}
	if !q.hasAddedWhere {
		q.Query += " WHERE "
		q.hasAddedWhere = true
		q.Query += "qs.interviewType IN (" + stringsArrayBuilder(interviewTypes) + ")"
	} else {
		q.Query += " AND qs.interviewType IN (" + stringsArrayBuilder(interviewTypes) + ")"
	}
}

func (q *QuestionSetFilter) AddLimitAndOffset(limit, offset int64) {
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