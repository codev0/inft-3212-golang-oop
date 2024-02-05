package main

import "fmt"

type BankAccount struct {
	Iban   string
	Amount uint
}

type Employee struct {
	Name string
	ID   string
	BankAccount
}

func (e Employee) Description() string {
	return fmt.Sprintf("%s (%s), Bank: %s, %d", e.Name, e.ID, e.BankAccount.Iban, e.BankAccount.Amount)
}

type Manager struct {
	Employee
	Reports []Employee
}

func (m Manager) FindNewEmployees() []Employee {
	// do business logic
	return nil
}

func composition() {
	m := Manager{
		Employee: Employee{
			Name: "Bob Bobson",
			ID:   "12345",
			BankAccount: BankAccount{
				Iban:   "DE123456789",
				Amount: 1000,
			},
		},
		Reports: []Employee{},
	}
	fmt.Println(m.ID)
	// prints 12345
	fmt.Println(m.Description())
	// prints Bob Bobson (12345)

	// var eFail Employee = m // compilation error!
	// var eOK Employee = m.Employee // ok!
}
