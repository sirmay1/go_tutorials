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
go build main.go
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

