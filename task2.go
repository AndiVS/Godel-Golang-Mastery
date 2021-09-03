package main

import "fmt"

type Bus struct {
	driverName         string // name of bus driver
	busNumber          string // bus number like 19, A12 and so on
	registrationNumber string //bus registration Number like 1542AX-4
}

type Route struct {
	Bus
	arrivalPoint   string //arrival point
	departurePoint string //departure point
	departureTime  int    // departure time
	arrivalTime    int    // arrival time
}

// Change route tume
func (r *Route) CangeRouteTime(newDepartureTime int, newArrivalTime int) {
	r.departureTime = newDepartureTime
	r.arrivalTime = newArrivalTime
}

// counting time in route
func RouteTime(r *Route) int {
	return r.arrivalTime - r.departureTime
}

func main() {

	route := [3]Route{{Bus{"Ibragim", "A123", "0482AX-4"}, "start", "finish", 1200, 1300}, {Bus{"A", "D45", "1542AX-4"}, "start", "finish", 1540, 1700}, {Bus{"Ibragim", "A123", "0482AX-4"}, "start", "finish", 1200, 1300}}
	for i, value := range route {
		//fmt.Printf("%d | %s %s %s %s %s %d %d\n", i, value.driverName, value.busNumber, value.registrationNumber, value.departurePoint, value.arrivalPoint, value.arrivalTime, value.departureTime)
		fmt.Printf("%v | %v\n", i, value)
	}
}
