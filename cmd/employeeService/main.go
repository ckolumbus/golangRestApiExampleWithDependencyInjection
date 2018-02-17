/**
 * File: main.go
 * Created Date: Tuesday February 13th 2018
 * Author: Chris Drexler, ckolumbus@ac-drexler.de
 * -----
 * Copyright (c) 2018 Chris Drexler
 *
 * MIT License : http://www.opensource.org/licenses/MIT
 */

package main

import (
	"github.com/labstack/echo"

	"github.com/ckolumbus/golangRestApiExampleWithDependencyInjection/pkg/employee/controller"
	"github.com/ckolumbus/golangRestApiExampleWithDependencyInjection/pkg/employee/db"
	"github.com/ckolumbus/golangRestApiExampleWithDependencyInjection/pkg/employee/persistence"

	// sqlite3
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	conn := db.SetupDB("sqlite3", "./db.sqlite")
	defer conn.Close()

	e := echo.New()

	employeePersist := persistence.NewEmployeePersist(conn)
	employeeController := controllers.NewEmployeeController(employeePersist)

	e.POST("/employee", employeeController.CreateEmployee)
	e.DELETE("/employee/:id", employeeController.DeleteEmployee)
	e.GET("/employee/:id", employeeController.GetEmployee)
	e.GET("/employee", employeeController.GetAllEmployees)

	e.Logger.Fatal(e.Start(":8080"))
}
