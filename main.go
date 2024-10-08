package main

import (
	"bufio"
	"fmt"
	"go-cli-app/shapes"
	"go-cli-app/utils"
	"os"
	"regexp"
)



func main() {
	utils.ClearScreen()
	top:
	var value string;

	utils.DisplayList()

	fmt.Print("\nenter option : ");
	fmt.Scan(&value)
	reader := bufio.NewReader(os.Stdin)

	matcher := regexp.MustCompile(`^[a-e]{1}$|^q{1}$`)
	isMatch := matcher.MatchString(value)

	
	if !isMatch {
		utils.Error("Please input corrently !!")
		goto top
	}
	if(value == "q"){
		utils.ClearScreen()
		utils.Success("Thank you for using us :)")
		os.Exit(1)
	}

	switch value {
	case "a" :
		utils.ClearScreen()
		shapes.Rectangle(reader)
		goto top
	case "b" :
		utils.ClearScreen()
		shapes.Square(reader)
		goto top
	case "c" :
		utils.ClearScreen()
		shapes.Circle(reader)
		goto top
	case "d" :
	utils.ClearScreen()
		shapes.Trapezoid(reader)
		goto top
	case "e" :
		utils.ClearScreen()
		shapes.Triangle(reader)
		goto top
		default : 
	fmt.Println("you choose :",value);

	}
		
}

	
	

	
	
	


