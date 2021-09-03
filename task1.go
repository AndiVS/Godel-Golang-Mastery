package main

import "fmt"

type Buss struct {
	driverName         string
	bussNumber         string
	registrationNumber string
}

type Route struct {
	Buss
	departurePoint string
	arrivalPoint   string
	departureTime  int
	arrivalTime    int
}

func (r *Route) CangeRouteTime(newDepartureTime int, newArrivalTime int) {
	r.departureTime = newDepartureTime
	r.arrivalTime = newArrivalTime
}

func RouteTime(r *Route) int {
	return r.arrivalTime - r.departureTime
}

func main() {

	route := [1]Route{{Buss{"Ibragim", "A123", "0482AX-4"}, "start", "finish", 1200, 1300}}
	first := RouteTime(&route[0])
	route[0].CangeRouteTime(1230, 1300)
	second := RouteTime(&route[0])
	fmt.Printf("%d %d\n", first, second)
	for i, value := range route {
		//	fmt.Printf("%d | %s %s %s %s %s %d %d\n", i, value.driverName, value.bussNumber, value.registrationNumber, value.departurePoint, value.arrivalPoint, value.arrivalTime, value.departureTime)
		fmt.Printf("%v | %v\n", i, value)
	}
}
