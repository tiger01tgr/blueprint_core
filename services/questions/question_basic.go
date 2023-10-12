package employers

import (
	questionSet_dao "backend/db/dao/question_sets"
	question_dao "backend/db/dao/questions"
	models "backend/db/models"
	"database/sql"
	"fmt"
)

func CreateQuestionSetAndQuestions(name string, roleId uint64, employerId uint64, interviewType string, questions []models.Question) error {
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

func CreateQuestionSet(name string, employerId uint64, roleId uint64, interviewType string) (uint64, error) {
	_, err := questionSet_dao.CreateQuestionSet(name, interviewType, employerId, roleId)
	if err != nil {
		return 0, err
	}
	row, err := questionSet_dao.GetQuestionSetsByName(name)
	var qs models.QuestionSet
	row.Scan(&qs.ID, &qs.Name, &qs.InterviewType, &qs.EmployerId, &qs.RoleId, &qs.CreatedAt, &qs.Deleted)
	return uint64(qs.ID), nil
}

func CreateQuestion(questionSetId uint64, question string, timelimit uint64) error {
	_, err := question_dao.CreateQuestion(int(questionSetId), question, int(timelimit))
	return err
}

func GetAllQuestionSets() ([]models.QuestionSet, error) {
	rows, err := questionSet_dao.GetAllQuestionSets()
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	var questionSets []models.QuestionSet
	for rows.Next() {
		var qs models.QuestionSet
		if err := rows.Scan(&qs.ID, &qs.Name, &qs.InterviewType, &qs.EmployerId, &qs.RoleId, &qs.CreatedAt, &qs.Deleted); err != nil {
			return nil, err
		}
		questionSets = append(questionSets, qs)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return questionSets, nil
}

func GetQuestionSetWithId(id uint64) (*models.QuestionSet, error) {
	row, err := questionSet_dao.GetQuestionSetsByID(int(id))
	if err != nil {
		return nil, err
	}
	qs, err := makeQuestionSetHelper(row)
	if err != nil {
		return nil, err
	}
	return qs, nil
}

func GetQuestionWithId(id uint64) (*models.Question, error) {
	row, err := question_dao.GetQuestionByID(int(id))
	if err != nil {
		return nil, err
	}
	q, err := makeQuestionHelper(row)
	if err != nil {
		return nil, err
	}
	return q, nil
}

func GetQuestionsWithQuestionSetId(id uint64) ([]models.Question, error) {
	rows, err := question_dao.GetQuestionsByQuestionSetID(int(id))
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
	return questions, nil
}

func EditQuestionSet(id uint64, name, interviewType string, employerId, roleId uint64) error {
	err := questionSet_dao.UpdateQuestionSet(int(id), name, interviewType, employerId, roleId)
	return err
}

func EditQuestion(id uint64, text string, timelimit uint64) error {
	err := question_dao.UpdateQuestion(int(id), text, int(timelimit))
	return err
}

// func DeleteQuestionSetRecursively(id uint64) error {
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
// 	err = questionSet_dao.DeleteQuestionSetRow(int(id))
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

func DeleteQuestionSet(id uint64) error {
	err := questionSet_dao.DeleteQuestionSet(int(id))
	return err
}

func DeleteQuestion(id uint64) error {
	err := question_dao.DeleteQuestion(int(id))
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
