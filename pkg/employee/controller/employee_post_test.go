/**
 * File: employee_post_test.go
 * Created Date: Sunday February 13th 2018
 * Author: Chris Drexler, ckolumbus@ac-drexler.de
 * -----
 * Copyright (c) 2018 Chris Drexler
 *
 * Licensed under the Apache License, Version 2.0 (the "LICENSE");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *       http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
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
