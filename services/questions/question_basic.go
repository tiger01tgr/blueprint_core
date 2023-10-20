package questions

import (
	custom_dao "backend/db/dao/custom"
	questionSet_dao "backend/db/dao/question_sets"
	question_dao "backend/db/dao/questions"
	models "backend/db/models"
	"database/sql"
	"fmt"
)

func CreateQuestionSetAndQuestions(name string, roleId int64, employerId int64, interviewType string, questions []models.Question) error {
	id, err := CreateQuestionSet(name, employerId, roleId, interviewType)
	if err != nil {
		return err
	}
	fmt.Println("id: ", id)
	for _, q := range questions {
		err = CreateQuestion(id, q.Text, q.TimeLimit)
		if err != nil {
			return err
		}
	}
	return nil
}

func CreateQuestionSet(name string, employerId int64, roleId int64, interviewType string) (int64, error) {
	_, err := questionSet_dao.CreateQuestionSet(name, interviewType, employerId, roleId)
	if err != nil {
		return 0, err
	}
	row, err := custom_dao.CustomQueryForRow("SELECT id FROM QuestionSets WHERE name = $1 AND interviewType = $2", name, interviewType)
	if err != nil {
		fmt.Println("err: ", err)
		return 0, err
	}
	var id int64
	row.Scan(&id)
	return id, err
}

func CreateQuestion(questionSetId int64, question string, timelimit int64) error {
	_, err := question_dao.CreateQuestion(int64(questionSetId), question, int64(timelimit))
	return err
}

func GetAllQuestionSets(page, limit int64) ([]models.QuestionSet, error) {
	offset := (page - 1) * limit
	rows, err := questionSet_dao.GetAllQuestionSets(offset, limit)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	var questionSets []models.QuestionSet
	for rows.Next() {
		var qs models.QuestionSet
		if err := rows.Scan(&qs.ID, &qs.Name, &qs.InterviewType, &qs.EmployerId, &qs.RoleId, &qs.CreatedAt, &qs.Deleted, &qs.Logo, &qs.RoleName, &qs.EmployerName, &qs.IndustryId, &qs.IndustryName); err != nil {
			return nil, err
		}
		questionSets = append(questionSets, qs)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return questionSets, nil
}

func GetPaginationForAllQuestionSets(limit int64) (int64, error) {
	row, err := questionSet_dao.GetNumberOfQuestionSets()
	if err != nil {
		return 0, err
	}
	var count int64
	row.Scan(&count)
	return (count / limit) + 1, nil
}

func GetQuestionSetWithId(id int64) (*models.QuestionSet, error) {
	row, err := questionSet_dao.GetQuestionSetsByID(int64(id))
	if err != nil {
		return nil, err
	}
	qs, err := makeQuestionSetHelper(row)
	if err != nil {
		return nil, err
	}
	return qs, nil
}

func GetQuestionWithId(id int64) (*models.Question, error) {
	row, err := question_dao.GetQuestionByID(int64(id))
	if err != nil {
		return nil, err
	}
	q, err := makeQuestionHelper(row)
	if err != nil {
		return nil, err
	}
	return q, nil
}

func GetQuestionsWithQuestionSetId(id int64) (*[]models.Question, error) {
	rows, err := question_dao.GetQuestionsByQuestionSetID(id)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	var questions []models.Question
	for rows.Next() {
		var q models.Question
		if err := rows.Scan(&q.ID, &q.QuestionSetId, &q.Text, &q.TimeLimit); err != nil {
			return nil, err
		}
		questions = append(questions, q)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	fmt.Println("questions: ", questions)
	return &questions, nil
}

func GetFilteredQuestionSets(employers, industries, roles []int64, interviewTypes []string, page int64, limit int64) ([]models.QuestionSet, error) {
	offset := (page - 1) * limit
	questionSetFilter := CreateQuestionSetFilter()
	questionSetFilter.CreateGetBasicQuery()
	questionSetFilter.AddEmployersFilter(employers)
	questionSetFilter.AddIndustriesFilter(industries)
	questionSetFilter.AddRolesFilter(roles)
	questionSetFilter.AddInterviewTypesFilter(interviewTypes)
	questionSetFilter.AddLimitAndOffset(limit, offset)
	rows, err := custom_dao.CustomQueryForRows(questionSetFilter.Query)
	defer rows.Close()
	if err != nil {
		fmt.Println("err: ", err)
		return nil, err
	}
	var questionSets []models.QuestionSet
	for rows.Next() {
		var qs models.QuestionSet
		if err := rows.Scan(&qs.ID, &qs.Name, &qs.InterviewType, &qs.EmployerId, &qs.RoleId, &qs.CreatedAt, &qs.Deleted, &qs.Logo, &qs.RoleName, &qs.EmployerName, &qs.IndustryId, &qs.IndustryName); err != nil {
			return nil, err
		}
		questionSets = append(questionSets, qs)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return questionSets, nil

}

func GetPaginationForFilteredQuestionSets(employers, industries, roles []int64, interviewTypes []string, limit int64) (int64, error) {
	questionSetFilter := CreateQuestionSetFilter()
	questionSetFilter.CreateCountBasicQuery()
	questionSetFilter.AddEmployersFilter(employers)
	questionSetFilter.AddIndustriesFilter(industries)
	questionSetFilter.AddRolesFilter(roles)
	questionSetFilter.AddInterviewTypesFilter(interviewTypes)
	fmt.Println("questionSetFilter.Query: ", questionSetFilter.Query)
	row, err := custom_dao.CustomQueryForRow(questionSetFilter.Query)
	if err != nil {
		fmt.Println("err: ", err)
		return 0, err
	}
	var count int64
	row.Scan(&count)
	return (count / limit) + 1, nil
}

func EditQuestionSet(id int64, name, interviewType string, employerId, roleId int64) error {
	err := questionSet_dao.UpdateQuestionSet(int64(id), name, interviewType, employerId, roleId)
	return err
}

func EditQuestion(id int64, text string, timelimit int64) error {
	err := question_dao.UpdateQuestion(int64(id), text, int64(timelimit))
	return err
}

// func DeleteQuestionSetRecursively(id int64) error {
// 	qs, err := GetQuestionSetWithId(id)
// 	if err != nil {
// 		return err
// 	}
// 	for _, q := range qs.Questions {
// 		err = DeleteQuestion(q)
// 		if err != nil {
// 			return err
// 		}
// 	}
// 	err = questionSet_dao.DeleteQuestionSetRow(int64(id))
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

func DeleteQuestionSet(id int64) error {
	err := questionSet_dao.DeleteQuestionSet(int64(id))
	return err
}

func DeleteQuestion(id int64) error {
	err := question_dao.DeleteQuestion(int64(id))
	return err
}

func makeQuestionSetHelper(row *sql.Row) (*models.QuestionSet, error) {
	var qs models.QuestionSet
	err := row.Scan(&qs.ID, &qs.Name, &qs.InterviewType, &qs.EmployerId, &qs.RoleId, &qs.CreatedAt, &qs.Deleted)
	if err != nil {
		return nil, err
	}
	return &qs, nil
}

func makeQuestionHelper(row *sql.Row) (*models.Question, error) {
	var q models.Question
	err := row.Scan(&q.ID, &q.QuestionSetId, &q.Text, &q.TimeLimit)
	if err != nil {
		return nil, err
	}
	return &q, nil
}
