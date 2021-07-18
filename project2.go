package main

import (
	"fmt"
)

func main() {
	var publisher string = "DizzyBooks Publishing Inc."
	var writer string = "Tracey Hatchet"
	var artist string = "Jewel Tampson"
	var title string = "Mr. GoToSleep"
	var year int64 = 1997
	var pageNumber int64 = 14
	var grade float32 = 6.5

	fmt.Println(title, "written by", writer, "drawn by", artist, "published by", publisher, "in", year, "- pages:", pageNumber, "- grade:", grade)

	var publisher2 string = "DizzyBooks Publishing Inc."
	var writer2 string = "Ryan N. Shawn"
	var artist2 string = "Phoebe Paperclips"
	var title2 string = "Epic Vol. 1"
	var year2 int64 = 2013
	var pageNumber2 int64 = 160
	var grade2 float32 = 9.0

	fmt.Println(title2, "written by", writer2, "drawn by", artist2, "published by", publisher2, "in", year2, "- pages:", pageNumber2, "- grade:", grade2)
}
