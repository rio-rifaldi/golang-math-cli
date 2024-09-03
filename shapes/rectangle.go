package shapes

import (
	"bufio"
	"fmt"
	"go-cli-app/utils"
	"math"
	"time"
)

func Rectangle(reader *bufio.Reader) {
		topA:
		fmt.Print("\nInput rectangle length : ");
		length,err := utils.InputNumber(reader)
		if err != nil {goto topA}
		
		fmt.Print("Input rectangle width : ");
		width,err := utils.InputNumber(reader)
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