package main

import ("fmt"
"strings"
"bufio"
 "os"
)

func userMessage(users []string, printer func(string, string)) {
    for _, value := range users { // Use _ to ignore the index
        printer(value, " Welcome") // Add a space before "Welcome"
    }
}

func greeting(user string, message string) {
    fmt.Println(user + message) // Use + to concatenate strings
}

func welcome() (string, int) {
    users := []string{"ABHIL", "AKHIL", "BINIL", "ROHIT"}
    userMessage(users, greeting)
    return "SUCCESS", 10
}

func variablesAndFunctions() {
    fmt.Print("\n----------------     VARIABLES    ---------------- \n ")
    operationStatus, operationValue := welcome()
    fmt.Println("value after calling functions", operationStatus, operationValue)

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
    fmt.Print("\n----------------     ARRAY AND SLICES   ----------------  \n ")
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
    fmt.Println(test, test2, test3)
}

func sliceMethods() {
    fmt.Print("\n----------------     SLICE METHODS   ----------------  \n ")
    makeSlice := make([]int, 5, 10) // Creates a slice with length 5 and capacity 10
    mySlice := []int{1, 2, 3}
    appendData := append(mySlice, 4, 5) // "appendData" variable was missing

    source := []int{1, 2, 3}
    copyDestination := make([]int, len(source))
    copy(copyDestination, source) // Corrected variable name and added "copy" function

    original := []int{1, 2, 3, 4, 5}
    subSlice := original[1:4]
    length := len(original)

    fmt.Println(makeSlice, appendData, copyDestination, subSlice, length)
}

func stringMethods() {
    fmt.Print("\n----------------     STRING METHODS   ----------------  \n ")
    s := "Hello, world!   "
    length := len(s)
    hasPrefix := strings.HasPrefix(s, "Hello")
    hasSuffix := strings.HasSuffix(s, "world")
    contains := strings.Contains(s, "world") // true / false
    index := strings.Index(s, "world")       // -1 if not found
    lastOccurrence := strings.LastIndex(s, "l") // returns 10, Returns -1 if not found
    replace := strings.Replace(s, "!", " Abhil", 1) // last is count -> how many times need to be replaced
    tolower := strings.ToLower("HELLO")
    toupper := strings.ToUpper("world")
    trim := strings.TrimSpace(s)          // removes white space

    split := strings.Split(s, ",")       // []string{"Hello", " world!   "}
    join := strings.Join(split, ", ")

    fmt.Println(length, hasPrefix, hasSuffix, contains, index, lastOccurrence, replace, tolower, toupper, trim, split, join)
}

func loops() {
    fmt.Print("\n----------------     LOOPS   ----------------  \n ")
    for i := 0; i < 5; i++ {
        fmt.Println(i)
    }

    fmt.Println("------------")
    i := 6
    for i < 10 {
        if i == 8 {
            break
        }
          // if i == 7 {
        //     continue
        // }
        fmt.Println(i)
        i++
    }
    fmt.Println("------------")

    numbers := []int{11, 22, 33, 44, 55, 66, 77, 88, 99}
    for index, value := range numbers {
        fmt.Println(index, value)
    }

    fmt.Println("------------")

    // goto
    j := 10
loopLabel:
    if j < 15 {
        fmt.Println(j)
        j++
        goto loopLabel
    }
}

func updater(value *int) *int {
    *value = *value + 1
    return value
}

func pointers() {
    fmt.Print("\n--------------    POINTERS      --------------\n")
    var a int = 10
    var ptr *int
    ptr = &a
    result := updater(&a)
    fmt.Println(*ptr, *result)
}

func pointersSub(ptr3 *int) *int {
    y := *ptr3 + 1
    return &y
}

func pointersMain() {
    x := 10
    ptr := &x
    ptr2 := pointersSub(ptr)
    
    fmt.Println("before", *ptr, *ptr2)
}


// struct 
type Address struct {
    pin   int
    house string
}

type Person struct {
    firstName string
    lastName  string
    address   Address
}

func (p Person) StructMethod() string {
    return p.firstName + " " + p.lastName
}


func structLogics() {
    person := Person{
        firstName: "John",
        lastName:  "Doe",
        address: Address{
            pin:   12345,
            house: "123 Main Street",
        },
    }

    fmt.Println("inner", person.address.pin)
    fullName := person.StructMethod()
    fmt.Println("Full Name:", fullName)
}

func ReadInput() {
    var x int
    var y int
    fmt.Print("ENter Number :")
    fmt.Scan(&x)
    fmt.Println(x)

    // error handling on reading 
    _, err := fmt.Scan(&y)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("You entered:", y)
    }

    // looping 
    var slice []int
    for {
        fmt.Print("Enter a positive integer: ")
        _, err := fmt.Scan(&num)
        if err == nil && num > 0 {
            break
        }
        slice = append(slice, num) 
        fmt.Println("Invalid input. Please try again.")
    }
    fmt.Println("You entered:", num)


    // reading line of input 
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter a line of text: ")
    userInput, _ := reader.ReadString('\n')
    fmt.Println("You entered:", userInput)


}


// Interfacesss 

type shape interface {
    area() float64
}

type square struct {
    length float64
}

type circle struct {
    radius float64
}

func (a circle) area() float64 {
    return a.radius * a.radius * 3.14
}

func (a square) area() float64 {
    return a.length * a.length
}

func PrintData(s shape) {
    fmt.Printf("%T\t%v\n", s, s.area())
}

// empty interface 
func describe(val interface{}) {  // based on the passed value it dynamically adjust the type 
	fmt.Printf("Type: %T, Value: %v\n", val, val)
}


func InterFace() {
    r := circle{ 4 }
    s := square{ 5 }
    shapes := []shape{r, s}  // initialising interface with instances of circle and square struct 
    for _, a := range shapes {
        PrintData(a)
    }

    // empty interface 
    describe(42)
	describe("Hello, Go!")
	describe(3.14)

    // type Assertion
    var interfaceVar interface{} = 42
    val, ok := interfaceVar.(int)
    if ok {
        fmt.Println("Value is an int:", val)
    } else {
        fmt.Println("Value is not an int")
    }


    // type swtich
    switch val := interfaceVar.(type) {
    case int:
        fmt.Println("Value is an int:", val)
    case string:
        fmt.Println("Value is a string:", val)
    default:
        fmt.Println("Value is of an unknown type")
    }
    

}




func main() {
    printing()
    variablesAndFunctions()
    arrayAndSlice()
    sliceMethods()
    stringMethods()
    loops()
    pointersMain()
    structLogics()
    ReadInput()
    InterFace()
}
