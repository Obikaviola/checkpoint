package main

import "fmt"

// import "fmt"

// // composition
// type Address struct{
// 	Street string
// 	City string
// }

// type Employee struct{
// 	name string
// 	age	int
// 	isRemote bool
// }

// func (e *Employee) updateName(newName string){
// 	e.name = newName
// 	fmt.Println((e.name))
// }

// func main(){
// 	address := Address {
// 		Street: "123 Main Street",
// 		City: "New York",
// 	}

// 	employee1 := Employee {
// 		name: "Alice",
// 		age: 30,
// 		isRemote: true,
// 		Address: address,
// 	}

// 	employee1.updateName("Bob")

// 	employee1.PrintAddress()
// 	fmt.Println("address:", employee1.street, employee1.City)

// 	func (a Address) PrintAddress() {
// 		fmt.Println()
// 	}

// 	fmt.Println("Employee Name:", employee1.name)
// 	fmt.Println("Employee Name:", employee1.age)
// 	fmt.Println("Employee Name:", employee1.isRemote)

// 	// Annoymous function
// 	job := struct{
// 		title string
// 		salary int
// 	}{
// 		title: "software engineer",
// 		salary: 100000,
// 	}

// 	fmt.Println("job:", job.title)

// 	employeePtr := &employee1
// 	employeePtr.age = 31
// }

type Student struct{
	name string
	grade []int
	age int
}

func (s Student) getAge() int{ // The function acts on a student object
	return s.age
}

func (s *Student) setAge(age int){ // Change the age using a pointer
	s.age = age
}

func (s Student) getAverageGrade() float32{
	sum := 0
	for _, v := range s.grade {
		sum += v
	}
	return float32(sum) / float32(len(s.grade))
}

func (s Student) getMaxGrade() int {
	curMax := 0
	for _, v := range s.grade{
		if curMax < v{
			curMax = v
		}
	}
	return curMax
}

func main(){
	s1 := Student{"Tim", []int{70, 90, 80, 85}, 19}
	s2 := Student{"Joe", []int{50, 80, 90, 85, 45, 60}, 23}
	s1.getAge()
	fmt.Println(s1)
	s1.setAge(7)
	fmt.Println(s1)

	maxGrade1 := s1.getMaxGrade()
	maxGrade2 := s2.getMaxGrade()
	fmt.Println(maxGrade1)
	fmt.Println(maxGrade2)

	average2 := s2.getAverageGrade()
	average := s1.getAverageGrade()
	fmt.Println(average)
	fmt.Println(average2)
}