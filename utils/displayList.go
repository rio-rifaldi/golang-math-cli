package utils

import (
	"fmt"

	"github.com/fatih/color"
)


var Shapes = [5]string{"rectangle","square","circle","trapezoid","triangle"}


func DisplayList(){
	c := color.New(color.FgCyan).Add(color.Bold)
	c.Print("\n --------- MATH COUNTER --------\n\n",)	
	for index,val := range Shapes{
		fmt.Printf("%v) %v \n",string('a' + index),val)	
	}
	fmt.Println("q) Quit");
	
	
}