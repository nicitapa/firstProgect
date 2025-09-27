package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	mock_contracts "github.com/nicitapa/firstProgect/internal/contracts/mocks"
	"github.com/nicitapa/firstProgect/internal/errs"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestController_DeleteEmployeesByID(t *testing.T) {
	type mockBehaviour func(s *mock_contracts.MockServiceI, id int)

	testTable := []struct {
		name                 string
		paramID              string
		mockBehaviour        mockBehaviour
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:    "OK",
			paramID: "1",
			mockBehaviour: func(s *mock_contracts.MockServiceI, id int) {
				s.EXPECT().DeleteEmployeesByID(id).Return(nil)
			},
			expectedStatusCode:   http.StatusOK,
			expectedResponseBody: `{"message":"Employees deleted successfully"}`,
		},
		{
			name:                 "Invalid ID",
			paramID:              "abc",
			mockBehaviour:        func(s *mock_contracts.MockServiceI, id int) {},
			expectedStatusCode:   http.StatusBadRequest,
			expectedResponseBody: `{"error":"invalid employees id"}`,
		},
		{
			name:    "Service error",
			paramID: "2",
			mockBehaviour: func(s *mock_contracts.MockServiceI, id int) {
				s.EXPECT().DeleteEmployeesByID(id).Return(errs.ErrEmployeesNotfound)
			},
			expectedStatusCode:   http.StatusNotFound,
			expectedResponseBody: `{"error":"users not found"}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			svc := mock_contracts.NewMockServiceI(ctrl)
			id, _ := strconv.Atoi(testCase.paramID)
			testCase.mockBehaviour(svc, id)

			handler := NewController(svc)

			r := gin.New()
			r.DELETE("/employees/:id", handler.DeleteEmployeesByID)

			w := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodDelete, "/employees/"+testCase.paramID, nil)

			r.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedResponseBody, w.Body.String())
		})
	}
}
