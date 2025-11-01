package main

import (
	"fmt"
	"time"
)

func main(){
	i := 1
	fmt.Print("Write ", i , " as ")

	switch i{
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3: 
		fmt.Println("Three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It is Weekend")
	default:
		fmt.Println("It is weekday")

	}
	
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It is before Noon")
	default:
		fmt.Println("It is after Noon")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm boolean")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Dont know the Type %T\n", t)
		}
	}

	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}