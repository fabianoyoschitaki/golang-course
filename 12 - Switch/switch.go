package main

import "fmt"

func dayOfWeek(number int) string {
	switch number % 7 {
	case 1:
		return "Sunday"
	case 2:
		return "Monday"
	case 3:
		return "Tuesday"
	case 4:
		return "Wednesday"
	case 5:
		return "Thursday"
	case 6:
		return "Friday"
	case 7:
		return "Saturday"
	default:
		return "Invalid"
	}
}

func dayOfWeek2(number int) string {
	var dayOfWeek string
	switch {
	// in case you want to validate other variables
	case number%7 == 1:
		dayOfWeek = "Sunday"
		fallthrough // makes code try run next condition, even if it's not true!
	case number%7 == 2:
		dayOfWeek = "Monday"
	case number%7 == 3:
		dayOfWeek = "Tuesday"
	case number%7 == 4:
		dayOfWeek = "Wednesday"
	case number%7 == 5:
		dayOfWeek = "Thursday"
	case number%7 == 6:
		dayOfWeek = "Friday"
	case number%7 == 7:
		dayOfWeek = "Saturday"
	default:
		dayOfWeek = "Invalid"
	}
	return dayOfWeek
}

func main() {
	day := dayOfWeek(9)
	fmt.Println(day)

	day2 := dayOfWeek2(9)
	fmt.Println(day2)
}
