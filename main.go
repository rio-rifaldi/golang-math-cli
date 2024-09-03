package main

import (
	"bufio"
	"fmt"
	"go-cli-app/function"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var Shapes = [6]string{"rectangle","square","circle","hexagon","triangle","trapesium"}



func inputNumber(reader *bufio.Reader) (value int, err error ){
		length,_ := reader.ReadString('\n')
		input := strings.TrimSpace(length)		
		newInput,err := strconv.Atoi(input)
		
		if err != nil  {
			fmt.Println("pelase input number corrently, instead of character");
			return 0,err
		}
		return newInput,err
}

func rectangle(reader *bufio.Reader) {
		topA:
		fmt.Print("\nInput rectangle length : ");
		length,err := inputNumber(reader)
		if err != nil {goto topA}
		
		fmt.Print("Input rectangle width : ");
		width,err := inputNumber(reader)
		if err != nil {goto topA}
			
		curcumference := 2 * (width + length) 
		area := width * length
		hypotenuse := math.Sqrt(math.Pow(float64(length),2) * math.Pow(float64(width),2)) 

		fmt.Println("\ncircumference : ", curcumference );
		fmt.Println("area : ", area );
		fmt.Println("diagonal count : ", 2 );
		fmt.Println("diagonal width (hypotenuse) : ", hypotenuse);

		fmt.Println("\nwe will redirect you to main menu soon ....");
		
		time.Sleep(time.Millisecond * 4000)
}
func main() {
	function.ClearScreen()
	top:
	var value string;

	fmt.Print("\n\t\tMATH COUNTER\n\n");
	
	for index,val := range Shapes{
		fmt.Printf("%v) %v \n",string('a' + index),val)
		
	}
	

	fmt.Print("\nenter option [a-f] : ");
	fmt.Scan(&value)
	reader := bufio.NewReader(os.Stdin)

	matcher := regexp.MustCompile(`^[a-f]{1}$`)
	isMatch := matcher.MatchString(value)

	if !isMatch {
		log.Println("please input corrently")
		goto top
	}

	switch value {
	case "a" :
		function.ClearScreen()
			rectangle(reader)

		goto top
	case "b" :
		fmt.Println("you choose b");
	case "c" :
		fmt.Println("you choose c");
	case "d" :
		fmt.Println("you choose d");
	case "e" :
		fmt.Println("you choose e");
	case "f" :
		fmt.Println("you choose f");
	}

		
}

	
	

	
	
	


