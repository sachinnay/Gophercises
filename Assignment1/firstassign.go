package main

import "fmt"

/*type Employee struct {
	EmployeeName string
	EmployeeID   int
}*/
//add function takes input of two int and return int as result
func add(x, y int) int {
	return x + y + 0

}
func main() {
	/*ar myMap = make(map[string]Employee)
	myMap["Sachin"] = Employee{"Sachin", 120}
	myMap["Shubham"] = Employee{"Shubham", 130}
	fmt.Println(myMap)

	myEmp := map[string]Employee{
		"Karan": Employee{"Karan", 122},
	}


	fmt.Println("myemp  == > ", myEmp)*/

	fmt.Println("Addition is :: ", add(10, 20))
	// Output:
	//Addition is1 :: 20

}

