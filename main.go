package main

import "fmt"

func main() {
	//FIRST WAY
	//a map called colors with keys stored as strings and the values also stored as strings
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "fffffff",
	}

	//var colors map[string]string //SECOND WAY

	//colors := make(map[string]string) //THIRD WAY

	//fmt.Println(colors)

	printMap(colors)
}

//iterating over a map
func printMap(c map[string]string) {
	for key, value := range c {
		fmt.Println(key, value)
	}
}
