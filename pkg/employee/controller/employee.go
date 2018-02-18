/**
 * File: employee.go
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

package controller

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"

	"github.com/ckolumbus/golangRestApiExampleWithDependencyInjection/pkg/employee/dto"
	"github.com/ckolumbus/golangRestApiExampleWithDependencyInjection/pkg/employee/persistence"
)

// EmployeeController handler for Employee DTO CRUD requests
type EmployeeController struct {
	EmployeePersist persistence.IEmployeePersist
}

// NewEmployeeController is the constructor
func NewEmployeeController(employeePersist persistence.IEmployeePersist) *EmployeeController {
	return &EmployeeController{employeePersist}
}

// CreateEmployee handles the POST endpoint for creating new employees
func (ec *EmployeeController) CreateEmployee(c echo.Context) error {
	emp := new(dto.Employee)
	if err := c.Bind(emp); err != nil {
		return err
	}

	response, err := ec.EmployeePersist.Save(emp)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, response)
}

// DeleteEmployee handles the DELETE request, taking an employee `id` as parameter
func (ec *EmployeeController) DeleteEmployee(c echo.Context) error {
	requestedID := c.Param("id")

	response, err := ec.EmployeePersist.Delete(requestedID)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, response)
}

// GetEmployee handles the GET request for one employee identified by `id`
func (ec *EmployeeController) GetEmployee(c echo.Context) error {
	requestedID := c.Param("id")
	fmt.Println(requestedID)
	response, err := ec.EmployeePersist.Get(requestedID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("ID '%s' not found", requestedID))
	}

	return c.JSON(http.StatusOK, response)
}

// GetEmployees GET  all employees
func (ec *EmployeeController) GetAllEmployees(c echo.Context) error {
	response, err := ec.EmployeePersist.GetAll()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Error requesting all employees"))
	}

	return c.JSON(http.StatusOK, response)
}
