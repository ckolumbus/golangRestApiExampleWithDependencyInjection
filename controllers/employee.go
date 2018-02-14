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

type EmployeeController struct {
	EmployeePersist persistence.IEmployeePersist
}

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

func (ec *EmployeeController) DeleteEmployee(c echo.Context) error {
	requestedID := c.Param("id")

	response, err := ec.EmployeePersist.Delete(requestedID)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, response)
}

func (ec *EmployeeController) GetEmployee(c echo.Context) error {
	requestedID := c.Param("id")
	fmt.Println(requestedID)
	response, err := ec.EmployeePersist.Get(requestedID)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, response)
}
