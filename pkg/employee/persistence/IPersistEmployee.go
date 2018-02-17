/**
 * File: IPersistEmployee.go
 * Created Date: Tuesday February 13th 2018
 * Author: Chris Drexler, ckolumbus@ac-drexler.de
 * -----
 * Copyright (c) 2018 Chris Drexler
 *
 * MIT License : http://www.opensource.org/licenses/MIT
 */

package persistence

import (
	"github.com/ckolumbus/golangRestApiExampleWithDependencyInjection/pkg/employee/dto"
)

// IPersistEmployee defines the interface to persist Employee
type IEmployeePersist interface {
	Save(*dto.Employee) (string, error)
	Delete(string) (string, error)
	Get(string) (*dto.Employee, error)
}
