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

package persistence

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/ckolumbus/golangRestApiExampleWithDependencyInjection/pkg/employee/dto"
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

func (ep *EmployeePersist) GetAll() (dto.Employees, error) {
	var (
		name    string
		id      string
		salary  string
		age     string
		empList dto.Employees
	)

	// http://go-database-sql.org/retrieving.html
	rows, err := ep.Db.Query("SELECT id, name, age, salary FROM employees")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&id, &name, &age, &salary)
		if err != nil {
			log.Fatal(err)
		}
		empList.Employees = append(empList.Employees, dto.Employee{ID: id, Name: name, Salary: salary, Age: age})
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return empList, nil
}
