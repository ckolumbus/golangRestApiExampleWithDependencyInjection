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

package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestGetEmployee(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/employe/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")
	h := &EmployeeController{EmployeePersistMock}

	// Assertions
	if assert.NoError(t, h.GetEmployee(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, userJSON, rec.Body.String())
	}
}
