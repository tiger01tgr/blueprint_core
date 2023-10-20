package handlers

import (
	// "backend/api/middleware"
	"backend/api/types"
	"backend/db/models"
	jobsService "backend/services/jobs"
	"backend/utils"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
)

func InitJobsRoutes(router chi.Router) {
	router.Route("/api/jobs", func(r chi.Router) {
		// Middlewares
		//r.Use(middleware.GoogleAuth)

		// Routes
		r.Get("/", GetJobs)
		r.Post("/", CreateJob)
		r.Patch("/", EditJob)
		r.Delete("/", DeleteJob)

		r.Get("/types", GetJobTypes)
		r.Post("/types", CreateJobType)
		r.Patch("/types", EditJobType)
		r.Delete("/types", DeleteJobType)
	})
}

type JobsResponse struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Employer    string `json:"employer"`
	EmployerId  int64  `json:"employer_id"`
	Role        string `json:"role"`
	RoleId      int64  `json:"role_id"`
	JobType     string `json:"job_type"`
	JobTypeId   int64  `json:"job_type_id"`
	Industry	string `json:"industry"`
	IndustryId	int64  `json:"industry_id"`
	Description string `json:"description"`
	QuestionSetId int64 `json:"question_set_id"`
	Logo 	  string `json:"logo"`
}

type GetJobsResponse struct {
	Data       []JobsResponse `json:"data"`
	Pagination types.Pagination     `json:"pagination"`
}

func GetJobs(w http.ResponseWriter, r *http.Request) {
	// Set limit
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
	var page int64
	if r.FormValue("page") != "" {
		page, err = strconv.ParseInt(r.FormValue("page"), 10, 64)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
	} else {
		page = 1
	}
	// fmt.Println("page: ", page)
	// fmt.Println("limit: ", limit)
	var shouldFilter bool
	switch r.FormValue("query") {
		case "filter":
			shouldFilter = true
		default:
			shouldFilter = false
	}
	

	employersStr := r.FormValue("employers")
	industriesStr := r.FormValue("industries")
	rolesStr := r.FormValue("roles")
	jobTypesStr := r.FormValue("jobTypes")
	var employers, industries, roles, jobTypes []int64
	if employersStr != "" {
		employers = utils.ConvertStringToInt64Array(employersStr)
	}
	if industriesStr != "" {
		industries = utils.ConvertStringToInt64Array(industriesStr)
	}
	if rolesStr != "" {
		roles = utils.ConvertStringToInt64Array(rolesStr)
	}
	if jobTypesStr != "" {
		jobTypes = utils.ConvertStringToInt64Array(jobTypesStr)
	}
	fmt.Println("employers: ", employers)
	fmt.Println("industries: ", industries)
	fmt.Println("roles: ", roles)
	fmt.Println("jobTypes: ", jobTypes)
	// Getting data for response
	jobs, err := jobsService.GetJobs(employers, industries, roles, jobTypes, page, limit, shouldFilter)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(jobs), func(i, j int) {
		jobs[i], jobs[j] = jobs[j], jobs[i]
	})
	var responses []JobsResponse
	for _, job := range jobs {
		responses = append(responses, makeJobsResponseHelper(job))
	}
	totalPages, err := jobsService.GetPaginationForJobs(employers, industries, roles, jobTypes, limit, shouldFilter)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	pagination := types.Pagination{
		TotalPages:  totalPages,
		CurrentPage: page,
		Limit:       limit,
	}
	response := GetJobsResponse{responses, pagination}
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

func CreateJob(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	employerId := r.FormValue("employerId")
	jobTypeId := r.FormValue("jobTypeId")
	roleId := r.FormValue("roleId")
	questionSetId := r.FormValue("questionSetId")
	description := r.FormValue("description")

	if name == "" || employerId == "" || jobTypeId == "" || roleId == "" || questionSetId == "" || description == "" {
		log.Println("Missing required fields")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Name cannot be empty"))
		return
	}
	err := jobsService.CreateJob(name, employerId, jobTypeId, roleId, questionSetId, description)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func EditJob(w http.ResponseWriter, r *http.Request) {

}

func DeleteJob(w http.ResponseWriter, r *http.Request) {

}

func GetJobTypes(w http.ResponseWriter, r *http.Request) {
	jobTypes, err := jobsService.GetJobTypes()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	// Marshal the response to JSON
	jsonData, err := json.Marshal(jobTypes)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	// Write the response
	w.Write(jsonData)
}

func CreateJobType(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	if name == "" {
		log.Println("Name cannot be empty")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Name cannot be empty"))
		return
	}
	err := jobsService.CreateJobType(name)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func EditJobType(w http.ResponseWriter, r *http.Request) {

}

func DeleteJobType(w http.ResponseWriter, r *http.Request) {

}


func makeJobsResponseHelper(job models.Job) JobsResponse {
	return JobsResponse{
		ID:          job.ID,
		Name:        job.Name,
		Employer:    job.Employer,
		EmployerId:  job.EmployerId,
		Role:        job.Role,
		RoleId:      job.RoleId,
		JobType:     job.JobType,
		JobTypeId:   job.JobTypeId,
		Industry:	job.Industry,
		IndustryId:	job.IndustryId,
		Description: job.Description,
		QuestionSetId: job.QuestionSetId,
		Logo: job.Logo,
	}
}