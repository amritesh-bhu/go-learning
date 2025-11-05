package main

import (
	"fmt"
	"time"
)

func main(){

	i := 2
	fmt.Print("write",i," as ")
	switch i{
		
		case 1: 
			fmt.Println("One")
		case 2:
			fmt.Println("Two")
		case 3: 
			fmt.Println("Three")
	}

	switch time.Now().Weekday(){
		case time.Saturday, time.Sunday:
			fmt.Println("it's the weekend")
		default: 
			fmt.Println("it's the weekday")
	}
	
	t := time.Now()
	switch {
		case t.Hour() < 12:
			fmt.Println("it's befor noon")
		default:
			fmt.Println("it's after noon")
	}

	whatAmI := func(i interface{}){
		switch t := i.(type){
			case bool:
				fmt.Println("I am bool")
			case int:
				fmt.Println("I am int")
			default:
				fmt.Println("Don't know the type %T\n",t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}
	
