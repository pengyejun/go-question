package syncoperation

import "fmt"

func selectChan() {
	ch := make(chan int, 0)
	select {
	case <- ch:
		fmt.Println("<- ch")
	default:
		fmt.Println("default")
	}
	fmt.Println("456")
}
