package handlers

import (
	"backend/db/models"
	questionService "backend/services/questions"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

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

type QuestionResponse struct {
	Id 			int64 `json:"id"`
	Text 		string `json:"text"`
	TimeLimit 	int64 `json:"timelimit"`
}

type QuestionSetResponse struct {
	Id            int64  `json:"id"`
	Name          string `json:"name"`
	Logo          string `json:"logo"`
	Employer      string `json:"employer"`
	EmployerId    int64  `json:"employerId"`
	Role          string `json:"role"`
	RoleId        int64  `json:"roleId"`
	Industry      string `json:"industry"`
	IndustryId    int64  `json:"industryId"`
	InterviewType string `json:"interviewType"`
}

type QuestionSetWithQuestionsResponse struct {
	Id            int64  `json:"id"`
	Name 		string `json:"name"`
	Employer 	string `json:"employer"`
	EmployerId 	int64 `json:"employerId"`
	Role 		string `json:"role"`
	RoleId 		int64 `json:"roleId"`
	InterviewType string `json:"interviewType"`
	Questions 	[]QuestionResponse `json:"questions"`
}

type QuestionRequest struct {
	Text      string `json:"text"`
	Timelimit int64  `json:"timelimit"`
}

type PostQuestionSetRequest struct {
	Questions []QuestionRequest `json:"questions"`
	// Add other fields from your JSON request here
}

type GetQuestionSetsRequest struct {
	CompanyId     []int64  `json:"companyId"`
	RoleId        []int64  `json:"roleId"`
	IndustryId    []int64  `json:"industryId"`
	InterviewType []string `json:"interviewType"`
}

func GetQuestionSets(w http.ResponseWriter, r *http.Request) {
	var limit int64
	var err error
	if r.FormValue("limit") != "" {
		limit, err = strconv.ParseInt(r.FormValue("limit"), 10, 64)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
	} else {
		limit = 40
	}

	switch r.FormValue("query") {
		case "all": {
			questionSets, err := questionService.GetAllQuestionSets(limit)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
				return
			}
			rand.Seed(time.Now().UnixNano())
            rand.Shuffle(len(questionSets), func(i, j int) {
                questionSets[i], questionSets[j] = questionSets[j], questionSets[i]
            })
			var responses []QuestionSetResponse
			for _, questionSet := range questionSets {
				responses = append(responses, makeQuestionSetResponseHelper(questionSet))
			}
			jsonData, err := json.Marshal(responses)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(jsonData)
			break
		}
		
		case "filter": {
			break
		}

		default: {
			break
		}
	}


	// var request GetQuestionSetsRequest
	// // Parse the JSON data from the request body
	// if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// 	fmt.Println(err)
	// 	return
	// }
	// companies := request.CompanyId
	// roles := request.RoleId
	// industries := request.IndustryId
	// interviewTypes := request.InterviewType

	// if len(companies) == 0 && len(roles) == 0 && len(industries) == 0 && len(interviewTypes) == 0 {

	// }

}

func makeQuestionSetResponseHelper(questionSet models.QuestionSet) QuestionSetResponse {
	return QuestionSetResponse{
		Id:            int64(questionSet.ID),
		Name:          questionSet.Name,
		Logo:          questionSet.Logo,
		Employer:      questionSet.EmployerName,
		EmployerId:    int64(questionSet.EmployerId),
		Role:          questionSet.RoleName,
		RoleId:        int64(questionSet.RoleId),
		Industry:      questionSet.IndustryName,
		IndustryId:    int64(questionSet.IndustryId),
		InterviewType: questionSet.InterviewType,
	}
}

func makeQuestionSetWithQuestionsResponseHelper(questionSet *models.QuestionSet, questions *[]models.Question) QuestionSetWithQuestionsResponse {
	fmt.Println("questionSet: ", questions)
	var questionResponses []QuestionResponse
	for _, question := range *questions {
		questionResponses = append(questionResponses, QuestionResponse{
			Id: question.ID,
			Text: question.Text,
			TimeLimit: question.TimeLimit,
		})
	}
	return QuestionSetWithQuestionsResponse{
		Id:            int64(questionSet.ID),
		Name:          questionSet.Name,
		Employer:      questionSet.EmployerName,
		EmployerId:    int64(questionSet.EmployerId),
		Role:          questionSet.RoleName,
		RoleId:        int64(questionSet.RoleId),
		InterviewType: questionSet.InterviewType,
		Questions: questionResponses,
	}
}

func GetQuestionSet(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	questionSet, err := questionService.GetQuestionSetWithId(idInt)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	questions, err := questionService.GetQuestionsWithQuestionSetId(idInt)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	response := makeQuestionSetWithQuestionsResponseHelper(questionSet, questions)
	jsonData, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func GetQuestions(w http.ResponseWriter, r *http.Request) {

}

func GetQuestion(w http.ResponseWriter, r *http.Request) {

}

func CreateQuestionSet(w http.ResponseWriter, r *http.Request) {
	var request PostQuestionSetRequest
	// Parse the JSON data from the request body
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		fmt.Println(err)
		return
	}

	// Access the questions slice in the request
	questions := request.Questions
	var questionList []models.Question
	// Now you can work with the questions slice
	for _, q := range questions {
		questionList = append(questionList, models.Question{
			Text:      q.Text,
			TimeLimit: int64(q.Timelimit),
		})
	}

	name := r.FormValue("name")
	employerId := r.FormValue("employerId")
	employerIdInt, err := strconv.ParseInt(employerId, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	roleId := r.FormValue("roleId")
	roleIdInt, err :=strconv.ParseInt(roleId, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	interviewType := r.FormValue("type")

	err = questionService.CreateQuestionSetAndQuestions(name, int64(roleIdInt), int64(employerIdInt), interviewType, questionList)
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
