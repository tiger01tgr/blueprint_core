package handlers

import (
	"backend/api/middleware"
	"backend/api/types"
	sessionsService "backend/services/practice_sessions"
	"database/sql"
	"encoding/json"
	"log"
	"mime/multipart"
	"time"

	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func InitSessionsRoutes(router chi.Router) {
	router.Route("/api/sessions", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			// middleware.SuperAdminAuth(r)
			r.Get("/completed", GetCompletedPracticeSessions)
			r.Get("/completed/{sessionId}", GetCompletedPracticeSession)
		})

		r.Group(func(r chi.Router) {
			middleware.UserAuth(r)
			r.Get("/{questionSetId}", GetPracticeSession)
			r.Post("/{questionSetId}", CreatePracticeSession)
			r.Post("/{questionSetId}/{sessionId}/{questionId}", CreatePracticeSubmission)
		})
	})
}

type PracticeSessionResponse struct {
	ID                     int64     `json:"id"`
	QuestionSetId          int64     `json:"question_set_id"`
	CurrentQuestionId      int64     `json:"current_question_id"`
	LastAnsweredQuestionId int64     `json:"last_answered_question_id"`
	UserId                 int64     `json:"user_id"`
	Status                 string    `json:"status"`
	CompletedAt            time.Time `json:"completed_at"`
}

func CreatePracticeSession(w http.ResponseWriter, r *http.Request) {
	questionSetIdParam := chi.URLParam(r, "questionSetId")
	questionSetId, err := strconv.ParseInt(questionSetIdParam, 10, 64)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	// Get user id
	userId := r.Context().Value("id").(int64)

	// Create session
	err = sessionsService.CreatePracticeSession(userId, questionSetId)
	if err != nil && err.Error() == "User and Session pair already exists" {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusCreated)
	return
}

func CreatePracticeSubmission(w http.ResponseWriter, r *http.Request) {
	var video *multipart.File
	err := r.ParseMultipartForm(100 << 20) // 100 MB maximum file size
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Unable to parse form or form is too large"))
		return
	}
	videoFile, _, err := r.FormFile("video")
	defer videoFile.Close()
	if err != nil && err != http.ErrMissingFile {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	if err == http.ErrMissingFile {
		video = nil
	} else {
		video = &videoFile
	}

	questionSetIdParam := chi.URLParam(r, "questionSetId")
	questionSetId, err := strconv.ParseInt(questionSetIdParam, 10, 64)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	practiceSession := chi.URLParam(r, "sessionId")
	practiceSessionId, err := strconv.ParseInt(practiceSession, 10, 64)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	questionIdParam := chi.URLParam(r, "questionId")
	questionId, err := strconv.ParseInt(questionIdParam, 10, 64)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	userId := r.Context().Value("id").(int64)
	err = sessionsService.CreatePracticeSubmission(userId, questionSetId, practiceSessionId, questionId, video)
	if err != nil && (err.Error() == "Practice session does not exist" ||
		err.Error() == "Practice session id does not exist" ||
		err.Error() == "Question id does not exist in question set" ||
		err.Error() == "Practice session already completed" ||
		err.Error() == "Question id does not match current question id") {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusCreated)
	return
}

func GetPracticeSession(w http.ResponseWriter, r *http.Request) {
	questionSetIdParam := chi.URLParam(r, "questionSetId")
	questionSetId, err := strconv.ParseInt(questionSetIdParam, 10, 64)
	if err != nil {
		log.Println(questionSetIdParam)
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	log.Println(r.Context().Value("id"))
	userId := r.Context().Value("id").(int64)
	session, err := sessionsService.GetPracticeSession(userId, questionSetId)
	if err == sql.ErrNoRows {
		log.Println(err)
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Practice session not found"))
		return
	}
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	response := PracticeSessionResponse{
		ID:                     session.ID,
		QuestionSetId:          session.QuestionSetId,
		CurrentQuestionId:      session.LastAnsweredQuestionId,
		LastAnsweredQuestionId: session.LastAnsweredQuestionId,
		UserId:                 session.UserId,
		Status:                 session.Status,
		CompletedAt:            session.CompletedAt,
	}
	jsonData, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
	return
}

type CompletedPracticeSession struct {
	ID            int64     `json:"id"`
	UserId        int64     `json:"user_id"`
	QuestionSetId int64     `json:"question_set_id"`
	CompletedAt   time.Time `json:"completed_at"`
}

type CompletedPracticeSessionResponse struct {
	Data       []CompletedPracticeSession `json:"data"`
	Pagination types.Pagination           `json:"pagination"`
}

func GetCompletedPracticeSessions(w http.ResponseWriter, r *http.Request) {
	pageParam := r.FormValue("page")
	if pageParam == "" {
		pageParam = "1"
	}
	page, err := strconv.ParseInt(pageParam, 10, 64)
	if err != nil {
		log.Println(pageParam)
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	limitParam := r.FormValue("limit")
	if limitParam == "" {
		limitParam = "20"
	}
	limit, err := strconv.ParseInt(limitParam, 10, 64)
	if err != nil {
		log.Println(limitParam)
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	sessions, err := sessionsService.GetCompletedPracticeSessions(page, limit)
	completedSessions := make([]CompletedPracticeSession, len(sessions))
	for i, session := range sessions {
		completedSessions[i] = CompletedPracticeSession{
			ID:            session.ID,
			UserId:        session.UserId,
			QuestionSetId: session.QuestionSetId,
			CompletedAt:   session.CompletedAt,
		}
	}

	totalResults, totalPages, err := sessionsService.GetCompletedPracticeSessionsPagination(limit)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	response := CompletedPracticeSessionResponse{
		Data: completedSessions,
		Pagination: types.Pagination{
			TotalPages:  totalPages,
			CurrentPage: page,
			Limit:       limit,
			TotalResults: totalResults,
		},
	}
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	jsonData, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
	return
}

func GetCompletedPracticeSession(w http.ResponseWriter, r *http.Request) {
	sessionIdParam := chi.URLParam(r, "sessionId")
	sessionId, err := strconv.ParseInt(sessionIdParam, 10, 64)
	if err != nil {
		log.Println(sessionIdParam)
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	submissions, err := sessionsService.GetCompletedPracticeSessionSubmissions(sessionId)
	if err == sql.ErrNoRows {
		log.Println(err)
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Practice session not found"))
		return
	}
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	jsonData, err := json.Marshal(submissions)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
	return
}
