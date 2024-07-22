package main

import "fmt"

func concat(s1 string, s2 string ) string {
	return s1 + s2
}

func main() {
	fmt.Println(concat("Alex is ", "a unviersity student with a bright future"))
	fmt.Println(concat("My goal is to ", "be a proficient problem solver"))
}