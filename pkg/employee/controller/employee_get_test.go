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
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ckolumbus/golangRestApiExampleWithDependencyInjection/pkg/employee/dto"
	"github.com/stretchr/testify/assert"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo"
)

func TestGetEmployee(t *testing.T) {
	// Setup
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockEmployeePersist := NewMockIEmployeePersist(ctrl)

	// create return object from persistence layer
	employee := dto.Employee{ID: "2", Name: "name", Salary: "1", Age: "2"}
	// .. and create JSON representation for later asset
	employeeJSON, _ := json.Marshal(employee)

	mockEmployeePersist.EXPECT().Get("2").Return(&employee, nil)

	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/", nil)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.SetPath("/employee/:id")
	c.SetParamNames("id")
	c.SetParamValues("2")
	h := NewEmployeeController(mockEmployeePersist)

	// Act & Assert
	if assert.NoError(t, h.GetEmployee(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, employeeJSON, rec.Body.Bytes())
	}
}
func TestGetEmployee_(t *testing.T) {
	// Setup
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockEmployeePersist := NewMockIEmployeePersist(ctrl)
	mockEmployeePersist.EXPECT().Get("").Return(nil, errors.New("xxx"))

	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/", nil)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.SetPath("/employee")
	h := NewEmployeeController(mockEmployeePersist)

	// Assertions
	r := h.GetEmployee(c)
	if assert.Error(t, r) {
		httpError, ok := r.(*echo.HTTPError)
		assert.Equal(t, true, ok) // is
		assert.Equal(t, http.StatusBadRequest, httpError.Code)
	}
}
