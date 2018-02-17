/**
 * File: employee_test.go
 * Created Date: Saturday February 17th 2018
 * Author: Chris Drexler, ckolumbus@ac-drexler.de
 * -----
 * Copyright (c) 2018 Chris Drexler
 *
 * MIT License : http://www.opensource.org/licenses/MIT
 */

package persistence

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

// a successful case
func TestShouldUpdateStats(t *testing.T) {
	// ARRANGE
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	queryID := "1"
	rows := sqlmock.NewRows([]string{"id", "name", "age", "salary"}).
		AddRow(1, "name1", "10", "100")

	mock.ExpectQuery("^SELECT id, name, age, salary FROM employees WHERE id = \\?$").
		WithArgs(queryID).
		WillReturnRows(rows)

	// ACT
	persist := NewEmployeePersist(db)
	employee, errGet := persist.Get(queryID)

	// ASSERT
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	assert.NoError(t, errGet)
	assert.Equal(t, queryID, employee.ID)
	assert.Equal(t, "name1", employee.Name)
	assert.Equal(t, "10", employee.Age)
	assert.Equal(t, "100", employee.Salary)

}
