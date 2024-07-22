package main

import "fmt"

func main() {
	const name = "Alex Brt"
	const timeRan = 20.78

	msg := fmt.Sprintf("%s , a world champion athlete, ran 2km in astonishing %.1f seconds",
	name,
	timeRan)

	fmt.Println(msg)
}