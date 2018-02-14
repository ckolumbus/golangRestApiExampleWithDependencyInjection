/**
 * File: employee.go
 * Created Date: Tuesday February 13th 2018
 * Author: Chris Drexler, ckolumbus@ac-drexler.de
 * -----
 * Copyright (c) 2018 Chris Drexler
 *
 * MIT License : http://www.opensource.org/licenses/MIT
 */

package dto

// Employee definition
type Employee struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Salary string `json:"salary"`
	Age    string `json:"age"`
}

// Employees list of Employees
type Employees struct {
	Employees []Employee `json:"employees"`
}
