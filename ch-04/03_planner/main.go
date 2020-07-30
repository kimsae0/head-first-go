package main

import (
	"fmt"
	"github.com/matthewbrahms/head-first-go/ch-04/dates"
)

func main() {
	days := 3
	fmt.Println("Your appointment is in", days, "days")
	fmt.Println("with a follow-up in", days + dates.DaysInWeek, "days")
}