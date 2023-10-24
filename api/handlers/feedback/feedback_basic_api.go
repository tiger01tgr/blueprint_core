package handlers

import (
	"backend/api/middleware"
	// "backend/api/types"
	feedbackService "backend/services/feedback"
	"backend/utils"

	// "database/sql"
	"encoding/json"
	"log"
	"strings"

	// "time"

	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func InitFeedbackRoutes(router chi.Router) {
	router.Route("/api/feedback", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			// middleware.SuperAdminAuth(r)
			r.Post("/{sessionId}", PostFeedback)
		})

		r.Group(func(r chi.Router) {
			middleware.UserAuth(r)
			r.Get("/", GetAllFeedback)
			r.Get("/{feedbackId}", GetFeedbackEntries)
		})
	})
}

func PostFeedback(w http.ResponseWriter, r *http.Request) {
	sessionIdParam := chi.URLParam(r, "sessionId")
	sessionId, err := strconv.ParseInt(sessionIdParam, 10, 64)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	// Parse the request body into a map
	var requestBody map[string][]string
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&requestBody); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error parsing request body"))
		return
	}

	// Check if the "submission_ids" and "feedback" keys exist
	submissionIds, submissionIdsExist := requestBody["submission_ids"]
	feedback, feedbackExist := requestBody["feedback"]

	if !submissionIdsExist || !feedbackExist {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Both submission_ids and feedback are required"))
		return
	}
	submissionIdsInt := utils.ConvertStringToInt64Array(strings.Join(submissionIds, ","))

	// Create feedback
	err = feedbackService.CreateFeedback(sessionId, submissionIdsInt, feedback)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func GetAllFeedback(w http.ResponseWriter, r *http.Request) {
	// Get user id
	userId := r.Context().Value("id").(int64)

	// Get feedback
	feedback, err := feedbackService.GetAllFeedback(userId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	// Marshal the response to JSON
	jsonData, err := json.Marshal(feedback)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	// Write the response
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func GetFeedbackEntries(w http.ResponseWriter, r *http.Request) {
	feedbackIdParam := chi.URLParam(r, "feedbackId")
	feedbackId, err := strconv.ParseInt(feedbackIdParam, 10, 64)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	// Get user id
	userId := r.Context().Value("id").(int64)

	// Get feedback
	feedbackEntries, err := feedbackService.GetFeedbackEntries(userId, feedbackId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	jsonData, err := json.Marshal(feedbackEntries)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
	go feedbackService.MarkFeedbackAsRead(feedbackId)
}
