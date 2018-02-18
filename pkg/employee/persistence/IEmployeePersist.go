/**
 * File: IPersistEmployee.go
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
	"github.com/ckolumbus/golangRestApiExampleWithDependencyInjection/pkg/employee/dto"
)

// IEmployeePersist defines the interface to persist Employee
type IEmployeePersist interface {
	Save(*dto.Employee) (string, error)
	Delete(string) (string, error)
	Get(string) (*dto.Employee, error)
	GetAll() (dto.Employees, error)
}
