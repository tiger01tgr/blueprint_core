package handlers

import (
	"net/http"

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

func GetQuestionSets(w http.ResponseWriter, r *http.Request) {

}

func GetQuestionSet(w http.ResponseWriter, r *http.Request) {

}

func GetQuestions(w http.ResponseWriter, r *http.Request) {

}

func GetQuestion(w http.ResponseWriter, r *http.Request) {

}

func CreateQuestionSet(w http.ResponseWriter, r *http.Request) {

}

func CreateQuestion(w http.ResponseWriter, r *http.Request) {

}

func PatchQuestionSet(w http.ResponseWriter, r *http.Request) {
	
}

func PatchQuestion(w http.ResponseWriter, r *http.Request) {

}

func SubmitQuestion(w http.ResponseWriter, r *http.Request) {
	
}