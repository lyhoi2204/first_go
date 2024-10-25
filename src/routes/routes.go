package routes

import (
	"fmt"
)

func routing(routeX routeX) {
	routeX.HandleFunc("/", indexRoute())

	return routeX
}

func indexRoute() {
	fmt.Println("This is home")

}
