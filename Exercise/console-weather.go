package main

import (
	"fmt"
	"strconv"
)

var char = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "-"}
var l1 = []string{" _ ", "   ", " _ ", " _ ", "   ", " _ ", "   ", " _ ", " _ ", " _ ", "   "}
var l2 = []string{"| |", "  |", " _|", " _|", "|_|", "|_ ", "|_ ", "  |", "|_|", "|_|", " _ "}
var l3 = []string{"|_|", "  |", "|_ ", " _|", "  |", " _|", "|_|", "  |", "|_|", "  |", "   "}


func main() {
	fmt.Println(weatherCelcius(25, "Bangkok few cloud"))
	fmt.Println(weatherCelcius(34, "Tak sunny"))
	fmt.Println(weatherCelcius(17, "Phuket rainy"))
	fmt.Println(weatherCelcius(9, "Chiang-mai cold"))
	fmt.Println(weatherCelcius(-84, "Chiang-mai cold"))
	fmt.Println(weatherCelcius(-1234567890, "Chiang-mai cold"))
}

func weatherCelcius(Temp int, place string) string {

	line := [3]string{"", "", ""}
	var number string = strconv.Itoa(Temp)

	for _, c := range number {

		d1, d2, d3 := getformat(string(c))
		line[0] = line[0] + d1
		line[1] = line[1] + d2
		line[2] = line[2] + d3
	}

	return line[0] + "\n" + line[1] + "\n" + line[2] + " C\n " + place

	// fmt.Println(line[0])
	// fmt.Println(line[1])
	// fmt.Println(line[2] + " C")
	// fmt.Println(place)

}

func getformat(str string) (string, string, string) {

	for i, c := range char {
		if c == str {
			return l1[i], l2[i], l3[i]
		}
	}
	return "", "", ""
}
