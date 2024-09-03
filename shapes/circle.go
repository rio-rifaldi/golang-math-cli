package shapes

import (
	"bufio"
	"fmt"
	"go-cli-app/utils"
	"math"
	"time"
)

const PI = math.Pi

func Circle(reader *bufio.Reader) {
		topA:
		
		fmt.Print("Input circle radius : ");
		radius,err := utils.InputNumber(reader)
		if err != nil {goto topA}
			
		curcumference := PI * (2 *float64(radius)) 
		area := PI * math.Pow(float64(radius),2)


		fmt.Println("\ncircumference : ", curcumference );
		fmt.Println("area : ", area );
	

		fmt.Println("\nwe will redirect you to main menu soon ....");
		
		time.Sleep(time.Millisecond * 4000)
}