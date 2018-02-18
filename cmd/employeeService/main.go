/**
 * File: main.go
 * Created Date: Tuesday February 13th 2018
 * Author: Chris Drexler <ckolumbus@ac-drexler.de>
 * -----
 * Copyright (c) 2018 Chris Drexler <ckolumbus@ac-drexler.de>
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

package main

import (
	"github.com/karlkfi/inject"
	"github.com/labstack/echo"

	"github.com/ckolumbus/golangRestApiExampleWithDependencyInjection/pkg/employee/controller"
	"github.com/ckolumbus/golangRestApiExampleWithDependencyInjection/pkg/employee/db"
	"github.com/ckolumbus/golangRestApiExampleWithDependencyInjection/pkg/employee/persistence"

	// sqlite3
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	var (
		employeePersist    persistence.IEmployeePersist
		employeeController *controller.EmployeeController
	)
	conn := db.SetupDB("sqlite3", "./db.sqlite")
	defer conn.Close()

	e := echo.New()

	graph := inject.NewGraph(
		inject.NewDefinition(&employeePersist, inject.NewProvider(persistence.NewEmployeePersist, &conn)),
		inject.NewDefinition(&employeeController, inject.NewAutoProvider(controller.NewEmployeeController)),
	)
	graph.Resolve(&employeeController)

	e.POST("/employee", employeeController.CreateEmployee)
	e.DELETE("/employee/:id", employeeController.DeleteEmployee)
	e.GET("/employee/:id", employeeController.GetEmployee)
	e.GET("/employee", employeeController.GetAllEmployees)

	e.Logger.Fatal(e.Start(":8080"))
}
