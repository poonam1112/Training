package main

import "fmt"

// A Month specifies a month of the year (January = 1, ...).
type Month int

const (
	January Month = 1 + iota
	February
	March
	April
	May
	June
	July
	August
	September
	October
	November
	December
)

var months = [12]string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}

// String returns the English name of the month ("January", "February", ...).
func (m Month) String() string { return months[m-1] }

func main() {

	var month Month
	fmt.Println("enter month from 1-12")
	fmt.Scanf("%d", &month)
	//month = 3
	fmt.Printf("%s \n", month)

}
