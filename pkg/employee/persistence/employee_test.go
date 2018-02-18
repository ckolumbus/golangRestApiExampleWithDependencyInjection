/**
 * File: employee_test.go
 * Created Date: Sunday February 17th 2018
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
