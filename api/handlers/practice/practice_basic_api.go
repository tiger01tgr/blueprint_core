package handlers

import (
	"backend/db/models"
	questionService "backend/services/questions"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func InitPracticeRoutes(router chi.Router) {
	router.Route("/api/practice", func(r chi.Router) {
		r.Get("/", GetQuestionSets)
		r.Post("/", CreateQuestionSet)

		r.Get("/{id}", GetQuestionSet)
		r.Patch("/{id}", PatchQuestionSet)

		r.Get("/{id}/questions", GetQuestions)
		r.Post("/{id}/questions", CreateQuestion)

		r.Get("/{id}/questions/{questionId}", GetQuestions)
		r.Patch("/{id}/questions/{questionId}", PatchQuestion)
		r.Post("/{id}/questions/{questionId}/submit", SubmitQuestion)
	})
}

type QuestionSetResponse struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Employer string `json:"employer"`
	EmployerId int `json:"employerId"`
	Role string `json:"role"`
	RoleId int `json:"roleId"`
	InterviewType string `json:"interviewType"`
	QuestionIds []int `json:"questionIds"`
	NumQuestions int `json:"numQuestions"`
}

type QuestionRequest struct {
    Text     string `json:"text"`
    Timelimit string `json:"timelimit"`
}

type QuestionSetRequest struct {
    Questions []QuestionRequest `json:"questions"`
    // Add other fields from your JSON request here
}


func GetQuestionSets(w http.ResponseWriter, r *http.Request) {
	// questionSets, err := questionService.GetAllQuestionSets()
	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	w.Write([]byte(err.Error()))
	// 	return
	// }

	// w.Write([]byte(questionSets.String()))
}

func GetQuestionSet(w http.ResponseWriter, r *http.Request) {

}

func GetQuestions(w http.ResponseWriter, r *http.Request) {

}

func GetQuestion(w http.ResponseWriter, r *http.Request) {

}

func CreateQuestionSet(w http.ResponseWriter, r *http.Request) {
	var request QuestionSetRequest

    // Parse the JSON data from the request body
    if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Access the questions slice in the request
    questions := request.Questions
	var questionList []models.Question
    // Now you can work with the questions slice
    for _, q := range questions {
		timelimit, err := strconv.Atoi(q.Timelimit)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		questionList = append(questionList, models.Question{
			Text: q.Text,
			TimeLimit: uint64(timelimit),
		})
    }

	name := r.FormValue("name")
	employerId := r.FormValue("employerId")
	employerIdInt, err := strconv.Atoi(employerId)
	if (err != nil) {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	roleId := r.FormValue("roleId")
	roleIdInt, err := strconv.Atoi(roleId)
	if (err != nil) {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	interviewType := r.FormValue("type")

	err = questionService.CreateQuestionSetAndQuestions(name, uint64(roleIdInt), uint64(employerIdInt), interviewType, questionList)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func CreateQuestion(w http.ResponseWriter, r *http.Request) {

}

func PatchQuestionSet(w http.ResponseWriter, r *http.Request) {
	
}

func PatchQuestion(w http.ResponseWriter, r *http.Request) {

}

func SubmitQuestion(w http.ResponseWriter, r *http.Request) {
	
}

// func makeQuestionSetResponseHelper(questionSet questionService.QuestionSet) QuestionSetResponse {
// 	return QuestionSetResponse{
// 		Id: questionSet.Id,
// 		Name: questionSet.Name,
// 		Employer: questionSet.Employer,
// 		EmployerId: questionSet.EmployerId,
// 		Role: questionSet.Role,
// 		RoleId: questionSet.RoleId,
// 		InterviewType: questionSet.InterviewType,
// 		QuestionIds: questionSet.QuestionIds,
// 		NumQuestions: questionSet.NumQuestions,
// 	}
// }
