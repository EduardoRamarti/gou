package main

import "testing"

func TestGetFullTimeEmploye(t *testing.T) {
	table := []struct {
		id               int
		dni              string
		mockFunc         func()
		expectedEmployee FullTimeEmployee
	}{
		{
			id:  1,
			dni: "1",
			mockFunc: func() {
				GetEmployeeById = func(id int) (Employee, error) {
					return Employee{
						id:       1,
						position: "CEO",
					}, nil
				}

				GetPersonByDNI = func(id string) (Person, error) {
					return Person{
						name: "John Doe",
						age:  35,
						dni:  "1",
					}, nil
				}
			},
			expectedEmployee: FullTimeEmployee{
				Person: Person{
					age:  35,
					dni:  "1",
					name: "John Doe",
				},
				Employee: Employee{
					id:       1,
					position: "CEO",
				},
			},
		},
	}

	originalGetEmployeeById := GetEmployeeById
	originalGetPersonByDNI := GetPersonByDNI

	for _, test := range table {
		test.mockFunc()
		ft, err := GetFullTimeEmployeeById(test.id, test.dni)
		if err != nil {
			t.Errorf("Error when getting Employee")
		}

		if ft.age != test.expectedEmployee.age {
			t.Errorf("Error, got %d expected %d", ft.age, test.expectedEmployee.age)
		}
	}

	GetEmployeeById = originalGetEmployeeById
	GetPersonByDNI = originalGetPersonByDNI
}
