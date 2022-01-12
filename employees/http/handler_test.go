package http_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	fuzz "github.com/google/gofuzz"
	"github.com/stretchr/testify/assert"

	empHTTP "algogrit.com/httpex/employees/http"
	"algogrit.com/httpex/employees/service"
	"algogrit.com/httpex/entities"
)

func TestIndex(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSvc := service.NewMockEmployeeService(ctrl)

	sut := empHTTP.NewHandler(mockSvc)

	expectedEmployees := []entities.Employee{
		{0, "Gaurav", "LnD", 1001},
	}

	mockSvc.EXPECT().Index().Return(expectedEmployees, nil)

	req := httptest.NewRequest("GET", "/v1/employees", nil)
	resRec := httptest.NewRecorder()

	sut.ServeHTTP(resRec, req)

	respRes := resRec.Result()

	assert.Equal(t, http.StatusOK, respRes.StatusCode)

	var actualEmployees []entities.Employee

	json.NewDecoder(respRes.Body).Decode(&actualEmployees)

	assert.Equal(t, expectedEmployees, actualEmployees)
}

func TestCreate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSvc := service.NewMockEmployeeService(ctrl)

	sut := empHTTP.NewHandler(mockSvc)

	newEmployee := entities.Employee{0, "Alok", "DBA", 1001}

	// createdEmp := newEmployee
	// createdEmp.ID = 1

	// mockSvc.EXPECT().Create(gomock.Any()).Return(&newEmployee, nil)
	mockSvc.EXPECT().Create(newEmployee).Return(&newEmployee, nil)

	reqBody := strings.NewReader(`{"name": "Alok", "division": "DBA", "project": 1001}`)

	req := httptest.NewRequest("POST", "/v1/employees", reqBody)
	resRec := httptest.NewRecorder()

	sut.ServeHTTP(resRec, req)

	respRes := resRec.Result()

	assert.Equal(t, http.StatusOK, respRes.StatusCode)
}

func TestFuzzCreate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSvc := service.NewMockEmployeeService(ctrl)

	sut := empHTTP.NewHandler(mockSvc)

	mockSvc.EXPECT().Create(gomock.Any()).AnyTimes().Return(nil, nil)

	f := fuzz.NewWithSeed(time.Now().UnixNano())
	for i := 0; i < 50_000; i++ {
		var body []byte
		f.Fuzz(&body)

		// t.Log("Len:", len(body), "Cap:", cap(body))

		reqBody := bytes.NewReader(body)

		req := httptest.NewRequest("POST", "/v1/employees", reqBody)
		resRec := httptest.NewRecorder()

		sut.ServeHTTP(resRec, req)

		respRes := resRec.Result()

		assert.Equal(t, http.StatusBadRequest, respRes.StatusCode)
	}
}
