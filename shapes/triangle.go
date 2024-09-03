package shapes

import (
	"bufio"
	"fmt"
	"go-cli-app/utils"
	"math"
	"time"
)

func Triangle(reader *bufio.Reader) {
		topA:
		fmt.Print("\nInput triangular base : ");
		baseLength,err := utils.InputNumber(reader)
		if err != nil {goto topA}
		
		fmt.Print("Input rectangle height : ");
		height,err := utils.InputNumber(reader)
		if err != nil {goto topA}
			

		area := (height * baseLength) /2
		hypotenuse := math.Sqrt(math.Pow(float64(baseLength/2),2) * math.Pow(float64(height),2)) 


		fmt.Println("area : ", area );
	
		fmt.Println("hypotenuse : ", hypotenuse);

		fmt.Println("\nwe will redirect you to main menu soon ....");
		
		time.Sleep(time.Millisecond * 4000)
}