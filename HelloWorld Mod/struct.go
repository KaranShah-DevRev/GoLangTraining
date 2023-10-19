package main

import "fmt"

type Employee struct {
	Id    int
	Name  string
	dept  string
	email string
}

type Operations interface {
	set(Employee)
	updateEmail(string)
	print(Employee)
}

func (e Employee) set(emp Employee) {
	e.Id = emp.Id
	e.Name = emp.Name
	e.dept = emp.dept
	e.email = emp.email
}

func (e *Employee) updateEmail(email string) {
	e.email = email
}

func (e *Employee) print() {
	fmt.Println(*e)
}

func main() {
	emp := Employee{Id: 1, Name: "Karan", dept: "CS", email: "mailtokks290199@gmail.com"}
	emp.set(emp)
	emp.print()
	emp.updateEmail("i-karan.shah@devrev.ai")
	emp.print()
}
