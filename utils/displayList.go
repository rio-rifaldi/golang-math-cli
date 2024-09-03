package utils

import "fmt"


var Shapes = [5]string{"rectangle","square","circle","trapezoid","triangle"}


func DisplayList(){

	fmt.Print("\n\t\tMATH COUNTER\n\n");
	
	for index,val := range Shapes{
		fmt.Printf("%v) %v \n",string('a' + index),val)	
	}
	fmt.Println("q) Quit");
	
	
}