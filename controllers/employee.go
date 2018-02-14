/**
 * File: employee.go
 * Created Date: Tuesday February 13th 2018
 * Author: Chris Drexler, ckolumbus@ac-drexler.de
 * -----
 * Copyright (c) 2018 Chris Drexler
 *
 * MIT License : http://www.opensource.org/licenses/MIT
 */

package controllers

import (
	"fmt"
	"net/http"

	"github.com/ckolumbus/golangRestApiExampleWithDependencyInjection/dto"

	"github.com/labstack/echo"

	"github.com/ckolumbus/golangRestApiExampleWithDependencyInjection/persistence"
)

// EmployeeController handler for Employee DTO CRUD requests
type EmployeeController struct {
	EmployeePersist persistence.IEmployeePersist
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
		return err
	}

	return c.JSON(http.StatusOK, response)
}
