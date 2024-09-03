package shapes

import (
	"bufio"
	"fmt"
	"go-cli-app/utils"
	"math"
	"time"
)


func Trapezoid(reader *bufio.Reader) {
		topA:
		
		fmt.Print("Input parallel line (1) : ");
		lineOne,err := utils.InputNumber(reader)
		if err != nil {goto topA}

		fmt.Print("Input parallel line (2) : ");
		lineTwo,err := utils.InputNumber(reader)
		if err != nil {goto topA}

		if(lineTwo < lineOne){
			fmt.Println("\nparallel line two must be bigger than parallel line one");
			goto topA
			
		}
		fmt.Print("Input height : ");
		height,err := utils.InputNumber(reader)
		if err != nil {goto topA}
			
		hypotenuse := math.Sqrt(float64(height) * float64((lineTwo - lineOne) / 2))
		curcumference := (2 * hypotenuse) + float64(lineOne) + float64(lineTwo) 
		area := height * (lineOne + lineTwo)/2


		fmt.Println("\ncircumference : ", curcumference );
		fmt.Println("area : ", area );
		fmt.Println("diagonal width (hypotenuse) : ", hypotenuse);
	

		fmt.Println("\nwe will redirect you to main menu soon ....");
		
		time.Sleep(time.Millisecond * 4000)
}