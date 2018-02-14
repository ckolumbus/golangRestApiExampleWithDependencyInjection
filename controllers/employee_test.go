/**
 * File: emplolyee_test.go
 * Created Date: Tuesday February 13th 2018
 * Author: Chris Drexler, ckolumbus@ac-drexler.de
 * -----
 * Copyright (c) 2018 Chris Drexler
 *
 * MIT License : http://www.opensource.org/licenses/MIT
 */

package controllers

import (
	"errors"

	"github.com/ckolumbus/golangRestApiExampleWithDependencyInjection/dto"
)

type MyEmployeePersistMock struct {
}

func (epm *MyEmployeePersistMock) Save(*dto.Employee) (string, error) {
	return "", nil
}

func (epm *MyEmployeePersistMock) Delete(string) (string, error) {
	return "", nil

}

func (epm *MyEmployeePersistMock) Get(id string) (*dto.Employee, error) {
	if id == "" {
		return nil, errors.New("error")
	}
	return &dto.Employee{}, nil
}
