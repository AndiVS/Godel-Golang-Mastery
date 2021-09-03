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

// Stringers
func (r Route) String() string {
	return fmt.Sprintf("|%s %s %s %s %s %d %d", r.driverName, r.busNumber, r.registrationNumber, r.departurePoint, r.arrivalPoint, r.arrivalTime, r.departureTime)
}

// Change route tume
func (r *Route) CangeRouteTime(newDepartureTime int, newArrivalTime int) {
	r.departureTime = newDepartureTime
	r.arrivalTime = newArrivalTime
}

// counting time in route
func RouteTime(r *Route) int {
	return r.departureTime - r.arrivalTime
}

func main() {
	route := [3]Route{{Bus{"Ibragim", "A123", "0482AX-4"}, "start", "finish", 1200, 1300}, {Bus{"A", "D45", "1542AX-4"}, "start", "finish", 1540, 1700}, {Bus{"Ibragim", "A123", "0482AX-4"}, "start", "finish", 1200, 1300}}
	for i, value := range route {
		fmt.Println(i, value)
	}
}
