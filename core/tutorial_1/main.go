package main

import (
	"fmt"
)

func welcomeUser(user string) {
	fmt.Printf("Welcome new user, %v \n", user)
}
func returnWelcomeUser(user []string, login func(string)) {
	for _, value := range user {
		login(value)
	}
}

func errUser(err string) {
	fmt.Printf("Error Handler: %v, no such user name exists\n", err)
}
func returnErrUser(err []string, exitUser func(string)) {
	for _, value := range err {
		exitUser(value)
	}
}

func exitUser(user string) {
	fmt.Printf("Exit User Handler: %v, you have been logged out\n", user)
}
func returnExitUser(user []string, exit func(string)) {
	for _, value := range user {
		exit(value)
	}
}

func main() {
	returnWelcomeUser([]string{
		"Daniel", "Robert", "Mr Smith", "Kodiak", "Verdana", "Paul", "Smith",
		"David", "Truman", "Trump", "George",
	}, welcomeUser)

	returnErrUser([]string{
		"Daniel", "Robert", "Mr Smith", "Kodiak", "Verdana", "Paul", "Smith", "David", "Truman", "Trump", "George",
	}, errUser)

	returnExitUser([]string{
		"Daniel", "Robert", "Mr Smith", "Kodiak", "Verdana", "Paul", "Smith", "David", "Truman", "Trump", "George",
	}, exitUser)
}
