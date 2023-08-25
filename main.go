// Online Go compiler to run Golang program online
// Print "Hello World!" message

package main

import "fmt"

func variables() {
       fmt.Print("\n----------------     VARIABLES    ---------------- \n ")
    // strings
    var nameOne string = "ABHIL"

	var nameThree string // declaring the variable first and initialising it later
	nameThree = "Sobhana"

	var nameTwo = "K V" // type will be asigned dynamically

	nameFour := "Ahul" //  short hand of writing without var , can't be used outside of functions
	
		fmt.Println("string ", nameOne, nameTwo, nameThree, nameFour)
		
	// integer
	var age int = 10000
	
	var age2 int
	age2 = 20
	
	var age3 = 30
	
	age4 := 40
	
		fmt.Println("integer", age, age2, age3, age4)
		
	// boolean
	var trueValue bool = true
	    fmt.Println(trueValue)

}

func printing() {
    fmt.Print("\n----------------     PRINTING   ----------------  \n ")
    fmt.Print("Hello, ")
    fmt.Print("World, \n")
    fmt.Print("New Line Printed , \n") // Print wont add \n at the end by default were Println will add \n at the end by default
    age := 27
    name:= "ABhil"
    fmt.Print("\n My name is ", name,  " and my age ", age );
    fmt.Printf("\n My name is %v and my age is %v \n", name, age); // template printing / formatted string printing 
    fmt.Printf("\n Type of age is %T", age)
    
    fmt.Printf("\n Your score is %0.3f", 23.4367)  // to print formatted flot values  will only have 3 decimal points 
    fmt.Printf("\n Your score is %f \n", 23.4367)
    
    var formatted = fmt.Sprint("\n My name is ", name,  " and my age ", age ); // returns the string and can be stored a variable 
    fmt.Println(formatted)
}

func arrayAndSlice() {
    var ages [3]int = [3]int{20, 30, 25}
    var age2 [2]int = [2]int{30}
    age3 := [2]string{"abhil", "kv"}
    fmt.Println(ages, len(ages), age2)
    age3[0] = "Test"
    fmt.Println(age3[0])
    
    // slices
    scores := []int{20, 22, 32, 12}
    scores = append(scores, 34)
    fmt.Println("slices", scores[4])
    
    // slice ranges
    test := scores[1:2]
    test2 := scores[2:]
    test3 := scores[:3]
    
    
    
    
}

func main() {
    printing()
    variables()
    arrayAndSlice()
}
