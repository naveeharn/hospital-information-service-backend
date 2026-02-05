package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"

	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert"
	"github.com/naveeharn/hospital-information-service-backend/helper"
	"github.com/naveeharn/hospital-information-service-backend/middleware"
)

var (
	r = gin.Default()
)

func SetUpRequest(method, relativePath, jsonBody string) *http.Request {
	req, _ := http.NewRequest(method, relativePath, bytes.NewBuffer([]byte(jsonBody)))
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	return req
}

func SetAuthorizationHeaderRequest(request *http.Request) {
	r.POST("/staff/login", authController.Login)
	req := SetUpRequest(http.MethodPost, "/staff/login", `{
		"username": "admintwo",
		"password": "admintwo",
		"hospital": "Bangkok Hospital"
	}`)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	responseBody, _ := io.ReadAll(w.Body)
	response := helper.Response{}
	json.Unmarshal(responseBody, &response)
	accessToken := (response.Data).(map[string]any)["accessToken"]
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %v", accessToken))
}

func TestStaffLoginStatusOK(t *testing.T) {
	r.POST("/staff/login", authController.Login)
	req := SetUpRequest(http.MethodPost, "/staff/login", `{
		"username": "admintwo",
		"password": "admintwo",
		"hospital": "Bangkok Hospital"
	}`)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	responseBody, _ := io.ReadAll(w.Body)
	log.Printf("responseBody: %v", string(responseBody))
	assert.Equal(t, w.Code, http.StatusOK)

	response := helper.Response{}
	json.Unmarshal(responseBody, &response)
	assert.Equal(t, response.Result.Code, "P001000")
}

func TestStaffLoginBadRequest(t *testing.T) {
	r.POST("/staff/login", authController.Login)
	req := SetUpRequest(http.MethodPost, "/staff/login", `{
		"username": "admintwoa",
		"password": "admintwo",
		"hospital": "Bangkok Hospital"
	}`)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	responseBody, _ := io.ReadAll(w.Body)
	log.Printf("responseBody: %v", string(responseBody))
	assert.Equal(t, w.Code, http.StatusBadRequest)

	response := helper.Response{}
	json.Unmarshal(responseBody, &response)
	assert.Equal(t, response.Result.Code, "P001999")
}

func TestStaffCreateBadRequest(t *testing.T) {
	r.POST("/staff/create", staffController.CreateStaff)
	req := SetUpRequest(http.MethodPost, "/staff/create", `{
		"username": "ad",
		"password": "ad",
		"hospital": "Ba"
	}`)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	responseBody, _ := io.ReadAll(w.Body)
	log.Printf("responseBody: %v", string(responseBody))
	assert.Equal(t, w.Code, http.StatusBadRequest)

	response := helper.Response{}
	json.Unmarshal(responseBody, &response)
	assert.Equal(t, response.Result.Code, "P002999")
}

func TestStaffCreateDuplicationStaffAndHospital(t *testing.T) {
	r.POST("/staff/create", staffController.CreateStaff)
	req := SetUpRequest(http.MethodPost, "/staff/create", `{
		"username": "admintwo",
		"password": "admintwo",
		"hospital": "Bangkok Hospital"
	}`)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	responseBody, _ := io.ReadAll(w.Body)
	log.Printf("responseBody: %v", string(responseBody))
	assert.Equal(t, w.Code, http.StatusBadRequest)

	response := helper.Response{}
	json.Unmarshal(responseBody, &response)
	assert.Equal(t, response.Result.Code, "P002998")
}

func TestPatientSearchOK(t *testing.T) {
	r.POST("/patient/search", middleware.AuthorizeJWT(jwtService), patientController.SearchPatient)
	req := SetUpRequest(http.MethodPost, "/patient/search", `{
		"first_name": "N"
	}`)
	SetAuthorizationHeaderRequest(req)
	// authorization := req.Header.Get("Authorization")
	// log.Printf("authorization: %v", authorization)
	// assert.Equal(t, strings.HasPrefix(authorization, "Bearer "), true)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	responseBody, _ := io.ReadAll(w.Body)
	log.Printf("responseBody: %v", string(responseBody))
	assert.Equal(t, w.Code, http.StatusOK)

	response := helper.Response{}
	json.Unmarshal(responseBody, &response)
	assert.Equal(t, response.Result.Code, "P003000")
}

func TestPatientSearchByNationalIdInTheSameHospital(t *testing.T) {
	nationalId := "3100500123456"
	uri := fmt.Sprintf("/patient/search/%v", nationalId)
	r.GET("/patient/search/:id", middleware.AuthorizeJWT(jwtService), patientController.FindPatientByNationalIdOrPassportId)
	req := SetUpRequest(http.MethodGet, uri, "")
	SetAuthorizationHeaderRequest(req)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	responseBody, _ := io.ReadAll(w.Body)
	log.Printf("responseBody: %v", string(responseBody))
	assert.Equal(t, w.Code, http.StatusOK)

	response := helper.Response{}
	json.Unmarshal(responseBody, &response)
	assert.Equal(t, response.Result.Code, "P004000")
}

func TestPatientSearchByNationalIdDoesNotInTheSameHospital(t *testing.T) {
	nationalId := "1209900876543"
	uri := fmt.Sprintf("/patient/search/%v", nationalId)
	r.GET("/patient/search/:id", middleware.AuthorizeJWT(jwtService), patientController.FindPatientByNationalIdOrPassportId)
	req := SetUpRequest(http.MethodGet, uri, "")
	SetAuthorizationHeaderRequest(req)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	responseBody, _ := io.ReadAll(w.Body)
	log.Printf("responseBody: %v", string(responseBody))
	assert.Equal(t, w.Code, http.StatusBadRequest)

	response := helper.Response{}
	json.Unmarshal(responseBody, &response)
	assert.Equal(t, response.Result.Code, "P004001")
}

func TestPatientSearchByPassportIdInTheSameHospital(t *testing.T) {
	passportId := "AA1234567"
	uri := fmt.Sprintf("/patient/search/%v", passportId)
	r.GET("/patient/search/:id", middleware.AuthorizeJWT(jwtService), patientController.FindPatientByNationalIdOrPassportId)
	req := SetUpRequest(http.MethodGet, uri, "")
	SetAuthorizationHeaderRequest(req)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	responseBody, _ := io.ReadAll(w.Body)
	log.Printf("responseBody: %v", string(responseBody))
	assert.Equal(t, w.Code, http.StatusOK)

	response := helper.Response{}
	json.Unmarshal(responseBody, &response)
	assert.Equal(t, response.Result.Code, "P004000")
}

func TestPatientSearchByPassportIdDoesNotInTheSameHospital(t *testing.T) {
	passportId := "P99887766"
	uri := fmt.Sprintf("/patient/search/%v", passportId)
	r.GET("/patient/search/:id", middleware.AuthorizeJWT(jwtService), patientController.FindPatientByNationalIdOrPassportId)
	req := SetUpRequest(http.MethodGet, uri, "")
	SetAuthorizationHeaderRequest(req)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	responseBody, _ := io.ReadAll(w.Body)
	log.Printf("responseBody: %v", string(responseBody))
	assert.Equal(t, w.Code, http.StatusBadRequest)

	response := helper.Response{}
	json.Unmarshal(responseBody, &response)
	assert.Equal(t, response.Result.Code, "P004001")
}
