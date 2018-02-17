/**
 * File: employee.go
 * Created Date: Tuesday February 13th 2018
 * Author: Chris Drexler, ckolumbus@ac-drexler.de
 * -----
 * Copyright (c) 2018 Chris Drexler
 *
 * MIT License : http://www.opensource.org/licenses/MIT
 */

package persistence

import (
	"database/sql"
	"fmt"

	"github.com/ckolumbus/golangRestApiExampleWithDependencyInjection/dto"
)

type EmployeePersist struct {
	Db *sql.DB
}

func NewEmployeePersist(db *sql.DB) *EmployeePersist {
	return &EmployeePersist{db}
}

func (ep *EmployeePersist) Save(emp *dto.Employee) (string, error) {
	//
	sql := "INSERT INTO employees(name, age, salary) VALUES( ?, ?, ?)"
	stmt, err := ep.Db.Prepare(sql)

	if err != nil {
		fmt.Print(err.Error())
		return "", err
	}
	defer stmt.Close()
	result, err2 := stmt.Exec(emp.Name, emp.Salary, emp.Age)

	// Exit if we get an error
	if err2 != nil {
		fmt.Println(err)
		return "", err2
	}
	fmt.Println(result.LastInsertId())

	return emp.Name, nil
}

func (ep *EmployeePersist) Delete(requestedID string) (string, error) {

	sql := "Delete FROM employees Where id = ?"
	stmt, err := ep.Db.Prepare(sql)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	result, err2 := stmt.Exec(requestedID)
	if err2 != nil {
		fmt.Println(err)
		return "", err2
	}
	fmt.Println(result.RowsAffected())
	return "Deleted", nil
}

func (ep *EmployeePersist) Get(requestedID string) (*dto.Employee, error) {
	var err error
	fmt.Println(requestedID)
	var name string
	var id string
	var salary string
	var age string

	err = ep.Db.QueryRow("SELECT id, name, age, salary FROM employees WHERE id = ?", requestedID).Scan(&id, &name, &age, &salary)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	emp := dto.Employee{ID: id, Name: name, Salary: salary, Age: age}
	return &emp, nil
}
