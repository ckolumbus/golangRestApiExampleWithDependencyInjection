/**
 * File: employee_test.go
 * Created Date: Tuesday February 13th 2018
 * Author: Chris Drexler, ckolumbus@ac-drexler.de
 * -----
 * Copyright (c) 2018 Chris Drexler
 *
 * MIT License : http://www.opensource.org/licenses/MIT
 */

// adopted from : https://echo.labstack.com/guide/testing

package controller

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	dto "github.com/ckolumbus/golangRestApiExampleWithDependencyInjection/pkg/employee/dto"
	gomock "github.com/golang/mock/gomock"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestCreateEmployee(t *testing.T) {
	// Setup
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockEmployeePersist := NewMockIEmployeePersist(ctrl)

	// create return object from persistence layer
	employee := dto.Employee{ID: "2", Name: "name", Salary: "1", Age: "2"}
	// .. and create JSON representation for later asset
	employeeJSON, _ := json.Marshal(employee)
	mockEmployeePersist.EXPECT().Save(&employee).Return(employee.ID, nil)

	e := echo.New()
	req := httptest.NewRequest(echo.POST, "/", strings.NewReader(string(employeeJSON)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	h := NewEmployeeController(mockEmployeePersist)

	// Assertions
	if assert.NoError(t, h.CreateEmployee(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, "\""+employee.ID+"\"", rec.Body.String())
	}
}
