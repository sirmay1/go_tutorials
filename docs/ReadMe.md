02/14/2024

For the past 3 weeks I've been trying to connect the backend and frontend using Django. But unfortunately connecting React to Django is one of the tuffest things to implement due Django's strict nature where you must do everything the Django way. So, after spending an entire day of research I decided to try out Golang because from what I read Golang is much easier to implement and it is better designed for restful api's. Where as you can create a restful api with a built in method of Django called "django-cors-headers" & adding into the MIDDLEWARE within the settings.py a corsheader middleware script which would allow such implementations.
So I spent 2 hours on GO and so far and within these 2 hours it is an amazing programming language. It reminds me of C programming language from the compilation process to the simplicity of the language, I really am enjoying GO!
NOTE: I only decided to pick up Golang because after 3 weeks Django felt like it wasn't the right tool for the job. But this doesn't mean I will abandon Django. I will continue to use Django BUT only for creating CRUD apps but not for connecting a restful api.

** NOTE: github Push from this directory location: **
Make sure you are within the "go_tutorials" folder...

"
cd go_tutorials
git commit
			"



Tutorial: Learn GO Fast: Full Tutorial
author: Alex Mux
Time: 1:07:52
Begin: (START-POINT)

mkdir go_tutorials

NOTE: In tutorial the project folder name was go_tutorials NOT core. Keep the name in mind. (3:40)
cd core
go mod init core
NOTE: Afterwards within your core file a go.mod file should appear. (3:42)
NEXT:
mkdir tutorial_1
cd tutorial_1
touch main.go (4:15)


Starting Template: (5:50)

package main
import "fmt"

func main()
{
	fmt.println("Holy-Shit! This is Golang!")
}

Next: To run our code we need to do the following:
go build main.go -(OR)-
go run main.go -(OR)-

NOTE: a main file has been created for compiling (6:00)
NOTE: you have 2 options for running your go code:
./main
"OR" (6:10)
go run main.go
---------------------------------------------------------
data types: (6:30)
const
var
INTS:
int
int16
int32
int64

(used for positive numbers) great for using rgba(255, 255, 255, 1) color ranges int not recommended:
uint
uint8
uint16
uint32
uint64



Operators int versus floats: (9:45)

package main

import "fmt"

func main() {
	var n1 float32 = 3.14
	var n2 int = 7
	var n3 float32 = n1 + float32(n2)
	fmt.Println(n3)
}


Concatenation in Golang: (10:15)

package main

import "fmt"

func main() {
	var name string = "Pumas"
	var sneakers string = `These are my
	favorite sneakers called `
	fmt.Println(sneakers + "\n" + name)
}


How Bytecode is translated in GO!: (10:40)

NOTE: when trying to run the length of a string to determine the numeric value do the following:
NOTE: Runes are a data type in Go! (10:50)

package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var myString string = "My name is BillDev from the planet Earth!"
	fmt.Println(myString)
	fmt.Println(utf8.RuneCountInString("BillDev"))
}

timeline(12:00)

var1, var2 := 1, 2


calling a function within a function by creating your own function: (14:15)

package main

import "fmt"

func main() {
	var printMe string = "Hi! MyFunctional call back!"
	MyFunction(printMe)
}

func MyFunction(printMe string) {
	fmt.Println(printMe)
}

using math in a function you create and calling that new function within the main function: (14:58)

package main

import "fmt"

func main() {
	var numerator int = 64
	var denominator int = 2
	var results int = intDivision(numerator, denominator)
	fmt.Println(results)
}

func intDivision(numerator int, denominator int) int {
	var results int = numerator / denominator
	return results
}

NOTE: Practing Go on my own from what I am learning: 15:06

package main

import "fmt"

func main() {
	var increment int = 600
	var decrement int = increment - 1100
	var result int = floatingMath(increment, decrement)
	fmt.Println(result)

}

func floatingMath(increment int, decrement int) int {
	var result int = increment + decrement
	return result
}

NOTE: Practing Go on my own from what I am learning: 15:06

package main

import "fmt"

func main() {
	var input int = 255
	var output int = 100
	var solution = display(input, output)
	fmt.Println(solution)
}

func display(input int, output int) int {
	var solution int = input / output
	return solution
}

NOTE: Practing more math methods with Golang: (15:20)

package main

import "fmt"

func main() {
	var lessMath int = 11
	var moreMath int = 2
	var result, remainder int = mathTeacher(lessMath, moreMath)
	fmt.Println(result, remainder)
}

func mathTeacher(lessMath int, moreMath int) (int, int) {
	var result int = lessMath / moreMath
	var remainder int = lessMath % moreMath
	return result, remainder
}

NOTE: Please Practice how to catch errors from this TIMELINE(17:00):

package main

import (
	"errors"
	"fmt"
)

func main() {
	var numerator int = 11
	var denominator int = 0
	var result, remainder, err = intDivision(numerator, denominator)
	if err != nil {
		fmt.Printf(err.Error())
	}
	fmt.Printf("The result of the integer division is %v with remainder %v", result, remainder)
}


func intDivision(numerator int, denominator int) (int, int, error) {
	var err error
	if denominator == 0 {
		err = errors.New("Cannot Divide by Zero!")
		return 0, 0, err
	}
	var result int = numerator / denominator
	var remainder int = numerator % denominator
	return result, remainder, err
}

PAUSED (17:50) Break till tomorrow.

package main

import (
	"fmt"
)

func main() {
	var tech [25]string
	tech[0] = "JavaScript\n"
	tech[1] = "Kotlin\n"
	tech[2] = "Swift\n"
	tech[3] = "PHP\n"
	tech[4] = "Ruby\n"
	tech[5] = "CoffeeScript\n"
	tech[6] = "TypeScript\n"
	tech[7] = "Golang\n"
	tech[8] = "Java\n"
	tech[9] = "C#\n"
	tech[10] = "C\n"
	tech[11] = "C++\n"
	tech[12] = "Rust\n"
	tech[13] = "Python\n"
	tech[14] = "Mojo\n"
	tech[15] = "Elixir\n"
	tech[16] = "Scala\n"
	tech[17] = "Dart\n"
	tech[18] = "R\n"
	tech[19] = "Haskell\n"
	tech[21] = "O'Camel\n"
	tech[22] = "Visual Basics\n"
	tech[23] = "Visual Basic.net\n"
	tech[24] = "COBOL\n"

	fmt.Println(tech[0:24])
	fmt.Println(len(tech[0:24]))
}

Timeline(21:00)

package main

import (
	"fmt"
)

func main() {
	var tech [5]string = [5]string{"PHP", "JavaScript", "Kotlin", "Golang"}
	fmt.Println(tech[0:5])

}

Timeline(21:20)

package main

import (
	"fmt"
)

func main() {
	intArr := [...]int32{0, 1, 2, 3, 4, 5}
	fmt.Println(intArr)

}

NOTE: Timeline(21:45) append() allows you to insert an element into the array or list, similar to JavaScript.

package main

import (
	"fmt"
)

func main() {
	var intSlice []int32 = []int32{4, 5, 6}
	fmt.Println(intSlice)
	intSlice = append(intSlice, 7)
	fmt.Println(intSlice)
}

Tutorial 13# Pass By Value
 netninja

 package main

import "fmt"

func newString(x string, y string) {
	x = "funky"
	y = "Potatoes"
}

func main() {
	fname := "William"
	lname := "Da Castrato"

	newString(fname, lname)

	fmt.Println(fname, lname)
}

package main

import "fmt"

func main() {
	numerator := 44
	denominator := 11
	remainder, result := intMath(numerator, denominator)
	fmt.Printf("Remainder: %v\n Result: %v\n", remainder, result)
}

func intMath(numerator int, denominator int) (int, int) {
	var remainder int = numerator / denominator
	var result int = numerator * denominator
	return remainder, result
}


NOTE: Tutorial 13#
we are able to carry the value from the newReturn function onto the main function. Only by adding the string type on the right side of the parameter of newReturn. Also, returning the x variable so that we are able to reuse the value within the main function.

package main

import "fmt"

func newReturn(x string) string {
	x = "super - wedges!"
	return x
}

func main() {
	name := "horse with no name!"

	name = newReturn(name)

	fmt.Println(name)
}

NOTE: tutorial 13#

package main

import "fmt"

func newName(name string) string {
	name = "I want taco bell!"
	return name
}

func main() {
	name := "Silly Me!"

	name = newName(name)

	fmt.Printf("Silly Me! %v", name)
}

NOTE: Tutorial 12#
author: The NetNinja

package main

import "fmt"

func main() {

	Programs := map[int]string{
		1: "JavaScript",
		2: "PHP",
		3: "Ruby",
		4: "C",
		5: "Golang",
		6: "Kotlin",
		7: "Python",
	}
	fmt.Printf("My favorite programming language is %v\n", Programs[6])
	fmt.Printf("Here is a list of all the Programming languages I had used:\n %v\n", Programs)

}


NOTE: Tutorial 12#
author: The NetNinja

package main

import "fmt"

func main() {

	languages := []string{
		"Python",
		"JavaScript",
		"C++",
		"C",
		"C#",
		"Objective-C",
		"Java",
		"Scala",
		"Kotlin",
		"Swift",
	}
	fmt.Println(languages)

	framework := map[string]string{
		"Python":     "Django",
		"JavaScript": "Next.js",
		"C":          "Raylib",
		"C++":        "Qt",
		"Golang":     "Standard Library",
	}
	framework["Golang"] = "TempL"

	fmt.Println(framework["Golang"])

	for k, v := range languages {
		fmt.Println(k, "--", v)
	}

}

NOTE: example of pointers

package main

import "fmt"

func main() {
	y := 3
	var x *int = &y
	*x = 5
	fmt.Println(y)
	// y is now 5
}


NOTE: GO(Golang) Tutorial #5 - Arrays & Slices
author: The NetNinja
package main

import "fmt"

func main() {
	var ages [10]int = [10]int{
		20, 30, 40, 50, 60, 70, 80, 90, 100,
	}
	fmt.Println(ages)

	GO(Golang) Tutorial #5

	package main

import "fmt"

func main() {
	ages := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}

	newAges := ages[1:4]

	fmt.Println(newAges, len(newAges))

	fmt.Println(ages, len(ages))

	names := []string{
		"William", "castro", "Alia", "Peaches", "Tania", "Daniella",
	}
	names[5] = "Princess Fart"
	names[1] = "Castrato"
	names[4] = "Terminator"
	names[0] = "Sir William"
	names[2] = "Apple Cheeks"
	fmt.Println(names, len(names))

	nameListing := names[2:4]

	fmt.Println(nameListing, len(nameListing))

	scores := []int{
		100, 92, 88, 77, 78, 65, 72, 82, 96,
	}
	scores = append(scores, 100000000)
	fmt.Println(scores, len(scores))
}


NOTE: GO(Golang) Tutorial#6 - The Standard Library
author: The NetNinja
package main

import (
	"fmt"
	"sort"
)

func main() {
	//greetings := "Hello there friends! How are you doing from the Golang community?"
	//fmt.Println(strings.Contains(greetings, "Golang community?\n"))
	//fmt.Println(strings.ReplaceAll(greetings, "friends!", "Mi Amigo's!"))
	//fmt.Println(strings.Index(greetings, "Golang"))
	//fmt.Println(strings.Split(greetings, " "))

	ages := []int{45, 20, 35, 75, 60, 50, 25}

	sort.Ints(ages)
	fmt.Println(ages)
}


NOTE: GO(Golang) Tutorial #7 - Loops
YouTube
author: The Net Ninja

package main

import "fmt"

func main() {
	x := 0

	for x < 5 {
		fmt.Println("value of x is:", x)
		x++
	}
}


package main

import "fmt"

func main() {
	for x := 0; x < 5; x++ {
		fmt.Println("value of 'i' is:", x)
	}
}


NOTE: GO(Golang) Tutorial #7 - Loops

package main

import "fmt"

func main() {
	programs := []string{
		"Java", "C#", "C", "C++", "Objective-C", "JavaScript", "PHP", "Ruby", "ODIN", "Matlab", "Clojure", "Perl", "Fortran", "Simula", "Haskell", "Dart", "Kotlin", "Golang", "R", "F#", "Scala", "ML", "APL", "Basics", "Visual Basics", "VB.NET", "SQL", "NO~SQL", "XML", "Flash", "Swift", "Assembler", "OCaml", "Pascal",
	}
	for i := 0; i < len(programs); i = i + 1 {
		fmt.Println(programs[i])
	}
	tech := map[int]string{
		1:  "Java",
		2:  "C#",
		3:  "C",
		4:  "C++",
		5:  "Objective-C",
		6:  "JavaScript",
		7:  "PHP",
		8:  "Ruby",
		9:  "ODIN",
		10: "COBOL",
		11: "Elixir",
		12: "Matlab",
		13: "Clojure",
		14: "Perl",
		15: "Fortran",
		16: "TypeScript",
		17: "CoffeeScript",
		18: "Moca",
		19: "Simula",
		20: "Haskell",
		21: "Dart",
		22: "Kotlin",
		23: "Golang",
		24: "R",
		25: "F#",
		26: "GDScript",
		27: "R",
		28: "Rust",
		29: "Scala",
		30: "ML",
		31: "APL",
		32: "Basics",
		33: "Visual Basics",
		34: "VB.NET",
		35: "SQL",
		36: "NO~SQL",
		37: "XML",
		38: "HTML5",
		39: "CSS3",
		40: "Python",
		41: "Flash",
		42: "Swift",
		43: "Assembler",
		44: "OCaml",
		45: "Pascal",
	}
	for x := 1; x < len(tech); x += 1 {
		fmt.Println("_", x, "_")
		fmt.Printf("Programming Languages: %v", tech[x])
	}

}

NOTE: Go(Golang) Tutorial #7 - Loops

package main

import "fmt"

func main() {
	names := []string{
		"Super Mario", "Peaches", "Luigi", "Bowser", "Mario", "Smash Brothers", "Fox",
	}

	count := 0
	for count < len(names) {
		fmt.Println(names[count])
		count = count + 1
	}
}


NOTE: Go (Golang) Tutorial #7 - Loops

package main

import "fmt"

func main() {
	names := []string{
		"mario", "luigi", "peaches", "toad", "yoshi", "bowser",
	}
	for index, value := range names {
		fmt.Printf("the value at index %v\n is %v\n", index, value)
	}
}

NOTE: GO(Golang) Tutorial #7 - Loops
"UNDERSCORE" if you don't want the index use a "underscore" as a place holder and the value of index won't be displayed as apart of the operation after it has ran. Example below:

package main

import "fmt"

func main() {
	names := []string{
		"mario", "luigi", "peaches", "toad", "yoshi", "bowser",
	}
	for _, value := range names {
		fmt.Printf("the value is %v\n", value)
	}
}


NOTE: Go (Golang) Tutorial 8# - Booleans & Conditionals

package main

import "fmt"

func main() {

	age := 49

	fmt.Println(age <= 50)
	fmt.Println(age >= 50)
	fmt.Println(age == 49)
	fmt.Println(age != 50)

	if age < 30 {
		fmt.Println("age is less than 30")
	} else if age < 40 {
		fmt.Println("age is less than 40")
	} else {
		fmt.Println("age is not less than 45")
	}
	names := []string{
		"mario", "luigi", "yoshi", "peach", "bowser",
	}
	for index, value := range names {
		if index == 1 {
			fmt.Println("\ncontinuing at pos", index)
			continue
		}
		if index > 2 {
			fmt.Println("\nbreaking at pos.", index)
			break
		}
		fmt.Printf("\nThe value at pos. %v is %v", index, value)
	}
}
STOP!!! PAUSED: Go (Golang) Tutorial #8 - Booleans & Conditionals...
Continue @ Tutorial #9 author: The NetNinja


NOTE: PRACTICE map & for loop key & value


package main

import "fmt"

func main() {
	tech := map[string]string{
		"React":          "JavaScript",
		"Next.js":        "TypeScript",
		"Nuxt.js":        "TypeScript",
		"FastAPI":        "Python",
		"Django":         "Python",
		"Fiber":          "Golang",
		"Springboot":     "Java",
		"Spring":         "Java",
		"Laravel":        "PHP",
		"Gorilla":        "Golang",
		"TempL":          "Golang",
		"Qt":             "C",
		"Raylib":         "C",
		"UnReal Engine":  "C++",
		"Unity":          "C#",
		"Pygame":         "Python",
		"Three.js":       "JavaScript",
		"Android Studio": "Kotlin",
		"Ktor":           "Kotlin",
		"Kaboom.js":      "JavaScript",
		"Godot":          "GDScript",
		"Love":           "Lua",
		"Rails":          "Ruby",
		"ASP.NET":        "C#",
		"WinForms":       "Visual Basics",
		"SvelteKit":      "JavaScript",
		"Svelte":         "JavaScript",
		"React Native":   "TypeScript",
	}
	for key, value := range tech {
		fmt.Printf("\nThe Framework Is: %v\nThe Programming Language is: %v\n", key, value)
	}
}


NOTE: GO(Golang) Tutorial #9 - Using Functions (02/25/2024)
Three step process:
create a function with a basic operation like "Hello World"
create a function which creates a range loop of the value "name" and a method
creates a method which invokes the function outside of the scope of the main function and runs an operation



package main

import "fmt"

func main() {
	//sayGreetings("Billy")
	//sayGreetings("Mario")
	//sayGreetings("Bill Bob")
	//sayGreetings("Luigi")
	//sayBye("Billy")
	//sayBye("Mario")
	//sayBye("Bill Bob")
	//sayBye("Luigi")
	cycleNames([]string{"Bob", "Mariposa", "Jessica"}, sayGreetings)
	cycleNames([]string{"Bob", "Mariposa", "Jessica"}, sayGreetings)
	cycleGoodBye([]string{"Bob", "Mariposa", "Jessica"}, sayBye)
	cycleGoodBye([]string{"Bob", "Mariposa", "Jessica"}, sayBye)
	cycleStop([]string{"Bob", "Mariposa", "Jessica"}, sayStop)
	cycleStop([]string{"Bob", "Mariposa", "Jessica"}, sayStop)
}

func sayGreetings(n string) {
	fmt.Printf("Welcome, and hello, %v \n", n)
}
func sayBye(n string) {
	fmt.Printf("Time to go, %v \n", n)
}
func sayStop(n string) {
	fmt.Printf("Stop! Is your name %v ?\n", n)
}
func cycleStop(n []string, fn func(string)) {
	for _, V := range n {
		fn(V)
	}
}
func cycleNames(n []string, f func(string)) {
	for _, v := range n {
		f(v)
	}
}
func cycleGoodBye(names []string, F func(string)) {
	for _, value := range names {
		F(value)
	}
}
