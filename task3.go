package main

import "fmt"

type Buss struct {
	driverName         string
	bussNumber         string
	registrationNumber string
}

type Route struct {
	Buss
	arrivalPoint   string
	departurePoint string
	departureTime  int
	arrivalTime    int
}

func (r Route) String() string {
	return fmt.Sprintf("|%s %s %s %s %s %d %d", r.driverName, r.bussNumber, r.registrationNumber, r.departurePoint, r.arrivalPoint, r.arrivalTime, r.departureTime)
}

func (r *Route) CangeRouteTime(newDepartureTime int, newArrivalTime int) {
	r.departureTime = newDepartureTime
	r.arrivalTime = newArrivalTime
}

func RouteTime(r *Route) int {
	return r.departureTime - r.arrivalTime
}

func main() {
	route := [3]Route{{Buss{"Ibragim", "A123", "0482AX-4"}, "start", "finish", 1200, 1300}, {Buss{"A", "D45", "1542AX-4"}, "start", "finish", 1540, 1700}, {Buss{"Ibragim", "A123", "0482AX-4"}, "start", "finish", 1200, 1300}}
	for i, value := range route {
		fmt.Println(i, value)
	}
}
